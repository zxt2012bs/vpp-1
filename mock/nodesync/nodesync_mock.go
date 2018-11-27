package nodesync

import (
	"github.com/contiv/vpp/plugins/nodesync"
)

// MockNodeSync is a mock implementation of nodesync plugin.
type MockNodeSync struct {
	thisNodeName string

	nodes    nodesync.Nodes
	thisNode *nodesync.Node
}

// NewMockNodeSync is a constructor for MockNodeSync.
func NewMockNodeSync(thisNodeName string) *MockNodeSync {
	return &MockNodeSync{
		thisNodeName: thisNodeName,
		nodes:        make(nodesync.Nodes),
	}
}

// GetNodeID return this node ID as set via UpdateNode() method.
func (m *MockNodeSync) GetNodeID() uint32 {
	if m.thisNode == nil {
		return 0
	}
	return m.thisNode.ID
}

// PublishNodeIPs does nothing here.
func (m *MockNodeSync) PublishNodeIPs(addresses []*nodesync.IPWithNetwork, version nodesync.IPVersion) error {
	return nil
}

// GetAllNodes returns mock node data as set via UpdateNode() method.
func (m *MockNodeSync) GetAllNodes() nodesync.Nodes {
	return m.nodes
}

// UpdateNode allows to set mock node data to test against.
func (m *MockNodeSync) UpdateNode(node *nodesync.Node) *nodesync.NodeUpdate {
	if node.Name == m.thisNodeName {
		m.thisNode = node
	}
	prev, _ := m.nodes[node.Name]
	m.nodes[node.Name] = node
	return &nodesync.NodeUpdate{
		NodeName:  node.Name,
		PrevState: prev,
		NewState:  node,
	}
}

// DeleteNode allows to delete node data.
func (m *MockNodeSync) DeleteNode(nodeName string) *nodesync.NodeUpdate {
	if nodeName == m.thisNodeName {
		m.thisNode = nil
	}
	prev, hasPrev := m.nodes[nodeName]
	if !hasPrev {
		return nil
	}
	delete(m.nodes, nodeName)
	return &nodesync.NodeUpdate{
		NodeName:  nodeName,
		PrevState: prev,
		NewState:  nil,
	}
}