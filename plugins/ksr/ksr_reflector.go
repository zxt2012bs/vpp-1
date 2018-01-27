// Copyright (c) 2017 Cisco and/or its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ksr

import (
	"fmt"
	"reflect"
	"sync"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/ligato/cn-infra/logging"

	"k8s.io/apimachinery/pkg/fields"
	k8sRuntime "k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
)

// ReflectorStats defines the usage statistics for K8s State Reflectors
type ReflectorStats struct {
	NumAdds    int
	NumDeletes int
	NumUpdates int
	NumResyncs int

	NumArgErrors int
	NumAddErrors int
	NumDelErrors int
	NumUpdErrors int
	NumResErrors int
}

// Reflector holds data that is common to all KSR reflectors.
type Reflector struct {
	// Each reflector gets a separate child logger.
	Log logging.Logger
	// A K8s client gets the appropriate REST client.
	K8sClientset *kubernetes.Clientset
	// K8s List-Watch watches for Kubernetes config changes.
	K8sListWatch K8sListWatcher
	// Writer propagates changes into a data store.
	Writer KeyProtoValWriter
	// Lister lists values from a data store.
	Lister KeyProtoValLister
	// objType defines the type of the object handled by a particular reflector
	objType string
	// ksrStopCh is used to gracefully shutdown the Reflector
	ksrStopCh <-chan struct{}
	wg        *sync.WaitGroup
	// K8s cache
	k8sStore cache.Store
	// K8s controller
	k8sController cache.Controller
	// Reflector statistics
	stats ReflectorStats

	prefix string
	pa     ProtoAllocator
	kpc    K8sToProtoConverter

	// Data store sync status and the mutex that protects access to it
	dsSynced bool
	dsMutex  sync.Mutex

	syncStopCh chan bool
}

// reflectors is the reflector registry
var reflectors = make(map[string]*Reflector)

// DsItems defines the structure holding items listed from the data store.
type DsItems map[string]interface{}

// ProtoAllocator defines the signature for a protobuf message allocation
// function
type ProtoAllocator func() proto.Message

// K8sToProtoConverter defines the signature for a function converting k8s
// objects to KSR protobuf objects.
type K8sToProtoConverter func(interface{}) (interface{}, string, bool)

// K8sClientGetter defines the signature for a function that allocates
// a REST client for a given K8s data type
type K8sClientGetter func(*kubernetes.Clientset) rest.Interface

// ReflectorFunctions defines the function types required in the KSR reflector
type ReflectorFunctions struct {
	EventHdlrFunc cache.ResourceEventHandlerFuncs

	ProtoAllocFunc ProtoAllocator
	K8s2ProtoFunc  K8sToProtoConverter
	K8sClntGetFunc K8sClientGetter
}

// dataStoreDownEvent signals to all registered reflectors that the data store
// is down. Reflectors should stop updates from the  respective K8s caches.
// Optionally, if data store resync with the K8s cache is in progress, it will
// be abort.
func dataStoreDownEvent() {
	for _, r := range reflectors {
		select {
		case r.syncStopCh <- true:
			r.Log.Infof("%s: sent dataStoreResyncAbort signal", r.objType)
		default:
			r.Log.Infof("%s: syncStopCh full", r.objType)
		}
		r.stopDataStoreUpdates()
	}
}

// dataStoreUpEvent signals to all registered reflectors that the data store
// is back up. Reflectors should start the resync procedure between their
// respective data stores with their respective K8s caches.
func dataStoreUpEvent() {
	for _, r := range reflectors {
		select {
		case <-r.syncStopCh:
			r.Log.Infof("%s: syncStopCh flushed", r.objType)
		default:
			r.Log.Infof("%s: syncStopCh not flushed", r.objType)
		}
		r.startDataStoreResync()
	}
}

// GetStats returns the Service Reflector usage statistics
func (r *Reflector) GetStats() *ReflectorStats {
	return &r.stats
}

// Start activates the K8s subscription.
func (r *Reflector) Start() {
	r.wg.Add(1)
	go r.ksrRun()
}

// Close does nothing for this particular reflector.
func (r *Reflector) Close() error {
	if _, objExists := reflectors[r.objType]; !objExists {
		return fmt.Errorf("%s reflector type does not exist", r.objType)
	}
	delete(reflectors, r.objType)

	return nil
}

// HasSynced returns the KSR Reflector's sync status.
func (r *Reflector) HasSynced() bool {
	r.dsMutex.Lock()
	defer r.dsMutex.Unlock()
	return r.dsSynced
}

// stopDataStoreUpdates marks the data store to be out of sync with the
// K8s cache, which will stop any updates to the data store until proper
// reconciliation is finished.
func (r *Reflector) stopDataStoreUpdates() {
	r.dsMutex.Lock()
	defer r.dsMutex.Unlock()
	r.dsSynced = false
}

// ksrRun runs k8s subscription in a separate go routine.
func (r *Reflector) ksrRun() {
	defer r.wg.Done()
	r.Log.Infof("%s reflector is now running", r.objType)
	r.k8sController.Run(r.ksrStopCh)
	r.Log.Infof("%s reflector stopped", r.objType)
}

// listDataStoreItems gets all items of a given type from Etcd
func (r *Reflector) listDataStoreItems(pfx string, iaf func() proto.Message) (DsItems, error) {
	dsDump := make(map[string]interface{})

	// Retrieve all data items for a given data type (i.e. key prefix)
	kvi, err := r.Lister.ListValues(pfx)
	if err != nil {
		return dsDump, fmt.Errorf("%s reflector can not get kv iterator, error: %s", r.objType, err)
	}

	// Put the retrieved items to a map where an item can be addressed
	// by its key
	for {
		kv, stop := kvi.GetNext()
		if stop {
			break
		}
		key := kv.GetKey()
		item := iaf()
		err := kv.GetValue(item)
		if err != nil {
			r.Log.WithField("Key", key).
				Errorf("%s reflector failed to get object from data store, error %s", r.objType, err)
		} else {
			dsDump[key] = item
		}
	}

	return dsDump, nil
}

// markAndSweep performs the mark-and-sweep reconciliation between data in
// the k8s cache and data in Etcd. This function must be called with dsMutex
// locked, because it manipulates dsFlag and because no updates to the data
// store can happen while the resync is in progress.
func (r *Reflector) markAndSweep(dsItems DsItems, oc K8sToProtoConverter) error {
	for _, obj := range r.k8sStore.List() {
		k8sProtoObj, key, ok := oc(obj)
		if ok {
			dsProtoObj, exists := dsItems[key]
			if exists {
				if !reflect.DeepEqual(k8sProtoObj, dsProtoObj) {
					// Object exists in the data store, but it changed in the
					// K8s cache; overwrite the data store
					err := r.Writer.Put(key, k8sProtoObj.(proto.Message))
					if err != nil {
						r.stats.NumUpdErrors++
						return fmt.Errorf("update for key '%s' failed", key)
					}
					r.stats.NumUpdates++
				}
			} else {
				// Object does not exist in the data store, but it exists in
				// the K8s cache; create object in the data store
				err := r.Writer.Put(key, k8sProtoObj.(proto.Message))
				if err != nil {
					r.stats.NumAddErrors++
					return fmt.Errorf("add for key '%s' failed", key)
				}
				r.stats.NumAdds++
			}
			delete(dsItems, key)
		}
	}

	// Delete from data store all objects that no longer exist in the K8s
	// cache.
	for key := range dsItems {
		_, err := r.Writer.Delete(key)
		if err != nil {
			r.stats.NumDelErrors++
			return fmt.Errorf("delete for key '%s' failed", key)
		}
		r.stats.NumDeletes++

		delete(dsItems, key)
	}

	// This must be performed with r.dsMutex locked
	r.dsSynced = true
	return nil
}

// syncDataStoreWithK8sCache syncs data in etcd with data in KSR's
// k8s cache. Returns ok if reconciliation is successful, error otherwise.
func (r *Reflector) syncDataStoreWithK8sCache(dsItems DsItems) error {
	r.dsMutex.Lock()
	defer r.dsMutex.Unlock()

	// don't do anything unless the K8s cache itself is synced
	if !r.k8sController.HasSynced() {
		return fmt.Errorf("%s data sync: k8sController not synced", r.objType)
	}

	// Reconcile data store with k8s cache using mark-and-sweep
	err := r.markAndSweep(dsItems, r.kpc)
	if err != nil {
		return fmt.Errorf("%s data sync: mark-and-sweep failed, '%s'", r.objType, err)
	}
	return nil
}

// startDataStoreResync starts the synchronization of the data store with
// the reflector's K8s cache. The resync will only stop if it's successful,
// or until it's aborted because of a data store failure or a KSR process
// termination notification.
func (r *Reflector) startDataStoreResync() {
	go func(r *Reflector) {
		r.Log.Debug("%s: starting data sync", r.objType)
		r.stats.NumResyncs++

		// Keep trying to reconcile until data sync succeeds.
		for {
			// Try to get a snapshot of the data store.
			dsItems, err := r.listDataStoreItems(r.prefix, r.pa)
			if err == nil {
				// Now that we have a data store snapshot, keep trying to
				// resync the cache with it
				for {
					// Try to perform mark-and-sweep data sync
					err := r.syncDataStoreWithK8sCache(dsItems)
					if err == nil {
						r.Log.Infof("%s: data sync done, stats %+v", r.objType, r.stats)
						return
					}

					r.stats.NumResErrors++
					// Wait before attempting the resync again
					select {
					case <-r.ksrStopCh: // KSR is being terminated
						r.Log.Debug("Resync aborted due to KSR process termination")
						return
					case <-r.syncStopCh: // Data Store resync is aborted
						r.Log.Debug("Resync aborted due to data store down")
						return
					case <-time.After(100 * time.Millisecond):
					}
				}
			} else {
				r.Log.Debugf("%s data sync: error listing data store items, '%s'", r.objType, err)
			}

			r.stats.NumResErrors++
			// Wait before attempting to list data store items again
			select {
			case <-r.ksrStopCh: // KSR is being aborted
				r.Log.Debug("Resync aborted due to KSR process termination")
				return
			case <-r.syncStopCh: // Data Store resync is aborted
				r.Log.Debug("Resync aborted due to data store down")
				return
			case <-time.After(100 * time.Millisecond):
			}
		}
	}(r)
}

// ksrAdd adds an item to the Etcd data store. This function must be called
// with dsMutex locked, since it manipulates the dsSynced flag.
func (r *Reflector) ksrAdd(key string, item proto.Message) {
	err := r.Writer.Put(key, item)
	if err != nil {
		r.Log.WithField("err", err).Warnf("%s: failed to add item to data store", r.objType)
		r.stats.NumAddErrors++
		r.dsSynced = false
		r.startDataStoreResync()
		return
	}
	r.stats.NumAdds++
}

// ksrUpdate updates an item to the Etcd data store. This function must be
// called with dsMutex locked, since it manipulates the dsSynced flag.
func (r *Reflector) ksrUpdate(key string, itemOld, itemNew proto.Message) {
	if !reflect.DeepEqual(itemOld, itemNew) {

		r.Log.WithField("key", key).Debugf("%s: updating item in data store", r.objType)

		err := r.Writer.Put(key, itemNew)
		if err != nil {
			r.Log.WithField("err", err).
				Warnf("%s: failed to update item in data store", r.objType)
			r.stats.NumUpdErrors++
			r.dsSynced = false
			r.startDataStoreResync()
			return
		}
		r.stats.NumUpdates++
	}
}

// ksrDelete deletes an item from the Etcd data store. This function must be
// called with dsMutex locked, since it manipulates the dsSynced flag.
func (r *Reflector) ksrDelete(key string) {
	_, err := r.Writer.Delete(key)
	if err != nil {
		r.Log.WithField("err", err).
			Warnf("%s: Failed to remove item from data store", r.objType)
		r.stats.NumDelErrors++
		r.dsSynced = false
		r.startDataStoreResync()
		return
	}
	r.stats.NumDeletes++
}

// Init subscribes to K8s cluster to watch for changes in the configuration
// of k8s services. The subscription does not become active until Start()
// is called.
func (r *Reflector) ksrInit(stopCh <-chan struct{}, wg *sync.WaitGroup, prefix string,
	objType string, k8sObjType k8sRuntime.Object, ksrFuncs ReflectorFunctions) error {

	if _, objExists := reflectors[objType]; objExists {
		return fmt.Errorf("%s reflector type already exists", r.objType)
	}

	r.syncStopCh = make(chan bool, 1)

	r.ksrStopCh = stopCh
	r.wg = wg

	r.prefix = prefix
	r.pa = ksrFuncs.ProtoAllocFunc
	r.kpc = ksrFuncs.K8s2ProtoFunc

	var restClient rest.Interface
	if ksrFuncs.K8sClntGetFunc != nil {
		restClient = ksrFuncs.K8sClntGetFunc(r.K8sClientset)
	} else {
		// If API version getter not specified, use CoreV1 by default
		restClient = r.K8sClientset.CoreV1().RESTClient()
	}

	listWatch := r.K8sListWatch.NewListWatchFromClient(restClient, objType, "", fields.Everything())
	r.k8sStore, r.k8sController = r.K8sListWatch.NewInformer(
		listWatch,
		k8sObjType,
		0,
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
				r.dsMutex.Lock()
				defer r.dsMutex.Unlock()

				if !r.dsSynced {
					return
				}
				ksrFuncs.EventHdlrFunc.AddFunc(obj)
			},
			DeleteFunc: func(obj interface{}) {
				r.dsMutex.Lock()
				defer r.dsMutex.Unlock()

				if !r.dsSynced {
					return
				}
				ksrFuncs.EventHdlrFunc.DeleteFunc(obj)
			},
			UpdateFunc: func(oldObj, newObj interface{}) {
				r.dsMutex.Lock()
				defer r.dsMutex.Unlock()

				if !r.dsSynced {
					return
				}
				ksrFuncs.EventHdlrFunc.UpdateFunc(oldObj, newObj)
			},
		},
	)
	reflectors[objType] = r
	return nil
}
