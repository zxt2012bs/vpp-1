// +build !ignore_autogenerated

/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// This file was autogenerated by conversion-gen. Do not edit it manually!

package v1

import (
	unsafe "unsafe"

	v1 "k8s.io/api/rbac/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	rbac "k8s.io/kubernetes/pkg/apis/rbac"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1_ClusterRole_To_rbac_ClusterRole,
		Convert_rbac_ClusterRole_To_v1_ClusterRole,
		Convert_v1_ClusterRoleBinding_To_rbac_ClusterRoleBinding,
		Convert_rbac_ClusterRoleBinding_To_v1_ClusterRoleBinding,
		Convert_v1_ClusterRoleBindingList_To_rbac_ClusterRoleBindingList,
		Convert_rbac_ClusterRoleBindingList_To_v1_ClusterRoleBindingList,
		Convert_v1_ClusterRoleList_To_rbac_ClusterRoleList,
		Convert_rbac_ClusterRoleList_To_v1_ClusterRoleList,
		Convert_v1_PolicyRule_To_rbac_PolicyRule,
		Convert_rbac_PolicyRule_To_v1_PolicyRule,
		Convert_v1_Role_To_rbac_Role,
		Convert_rbac_Role_To_v1_Role,
		Convert_v1_RoleBinding_To_rbac_RoleBinding,
		Convert_rbac_RoleBinding_To_v1_RoleBinding,
		Convert_v1_RoleBindingList_To_rbac_RoleBindingList,
		Convert_rbac_RoleBindingList_To_v1_RoleBindingList,
		Convert_v1_RoleList_To_rbac_RoleList,
		Convert_rbac_RoleList_To_v1_RoleList,
		Convert_v1_RoleRef_To_rbac_RoleRef,
		Convert_rbac_RoleRef_To_v1_RoleRef,
		Convert_v1_Subject_To_rbac_Subject,
		Convert_rbac_Subject_To_v1_Subject,
	)
}

func autoConvert_v1_ClusterRole_To_rbac_ClusterRole(in *v1.ClusterRole, out *rbac.ClusterRole, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Rules = *(*[]rbac.PolicyRule)(unsafe.Pointer(&in.Rules))
	return nil
}

// Convert_v1_ClusterRole_To_rbac_ClusterRole is an autogenerated conversion function.
func Convert_v1_ClusterRole_To_rbac_ClusterRole(in *v1.ClusterRole, out *rbac.ClusterRole, s conversion.Scope) error {
	return autoConvert_v1_ClusterRole_To_rbac_ClusterRole(in, out, s)
}

func autoConvert_rbac_ClusterRole_To_v1_ClusterRole(in *rbac.ClusterRole, out *v1.ClusterRole, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Rules = *(*[]v1.PolicyRule)(unsafe.Pointer(&in.Rules))
	return nil
}

// Convert_rbac_ClusterRole_To_v1_ClusterRole is an autogenerated conversion function.
func Convert_rbac_ClusterRole_To_v1_ClusterRole(in *rbac.ClusterRole, out *v1.ClusterRole, s conversion.Scope) error {
	return autoConvert_rbac_ClusterRole_To_v1_ClusterRole(in, out, s)
}

func autoConvert_v1_ClusterRoleBinding_To_rbac_ClusterRoleBinding(in *v1.ClusterRoleBinding, out *rbac.ClusterRoleBinding, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Subjects = *(*[]rbac.Subject)(unsafe.Pointer(&in.Subjects))
	if err := Convert_v1_RoleRef_To_rbac_RoleRef(&in.RoleRef, &out.RoleRef, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_ClusterRoleBinding_To_rbac_ClusterRoleBinding is an autogenerated conversion function.
func Convert_v1_ClusterRoleBinding_To_rbac_ClusterRoleBinding(in *v1.ClusterRoleBinding, out *rbac.ClusterRoleBinding, s conversion.Scope) error {
	return autoConvert_v1_ClusterRoleBinding_To_rbac_ClusterRoleBinding(in, out, s)
}

func autoConvert_rbac_ClusterRoleBinding_To_v1_ClusterRoleBinding(in *rbac.ClusterRoleBinding, out *v1.ClusterRoleBinding, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Subjects = *(*[]v1.Subject)(unsafe.Pointer(&in.Subjects))
	if err := Convert_rbac_RoleRef_To_v1_RoleRef(&in.RoleRef, &out.RoleRef, s); err != nil {
		return err
	}
	return nil
}

// Convert_rbac_ClusterRoleBinding_To_v1_ClusterRoleBinding is an autogenerated conversion function.
func Convert_rbac_ClusterRoleBinding_To_v1_ClusterRoleBinding(in *rbac.ClusterRoleBinding, out *v1.ClusterRoleBinding, s conversion.Scope) error {
	return autoConvert_rbac_ClusterRoleBinding_To_v1_ClusterRoleBinding(in, out, s)
}

func autoConvert_v1_ClusterRoleBindingList_To_rbac_ClusterRoleBindingList(in *v1.ClusterRoleBindingList, out *rbac.ClusterRoleBindingList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]rbac.ClusterRoleBinding)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1_ClusterRoleBindingList_To_rbac_ClusterRoleBindingList is an autogenerated conversion function.
func Convert_v1_ClusterRoleBindingList_To_rbac_ClusterRoleBindingList(in *v1.ClusterRoleBindingList, out *rbac.ClusterRoleBindingList, s conversion.Scope) error {
	return autoConvert_v1_ClusterRoleBindingList_To_rbac_ClusterRoleBindingList(in, out, s)
}

func autoConvert_rbac_ClusterRoleBindingList_To_v1_ClusterRoleBindingList(in *rbac.ClusterRoleBindingList, out *v1.ClusterRoleBindingList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]v1.ClusterRoleBinding)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_rbac_ClusterRoleBindingList_To_v1_ClusterRoleBindingList is an autogenerated conversion function.
func Convert_rbac_ClusterRoleBindingList_To_v1_ClusterRoleBindingList(in *rbac.ClusterRoleBindingList, out *v1.ClusterRoleBindingList, s conversion.Scope) error {
	return autoConvert_rbac_ClusterRoleBindingList_To_v1_ClusterRoleBindingList(in, out, s)
}

func autoConvert_v1_ClusterRoleList_To_rbac_ClusterRoleList(in *v1.ClusterRoleList, out *rbac.ClusterRoleList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]rbac.ClusterRole)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1_ClusterRoleList_To_rbac_ClusterRoleList is an autogenerated conversion function.
func Convert_v1_ClusterRoleList_To_rbac_ClusterRoleList(in *v1.ClusterRoleList, out *rbac.ClusterRoleList, s conversion.Scope) error {
	return autoConvert_v1_ClusterRoleList_To_rbac_ClusterRoleList(in, out, s)
}

func autoConvert_rbac_ClusterRoleList_To_v1_ClusterRoleList(in *rbac.ClusterRoleList, out *v1.ClusterRoleList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]v1.ClusterRole)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_rbac_ClusterRoleList_To_v1_ClusterRoleList is an autogenerated conversion function.
func Convert_rbac_ClusterRoleList_To_v1_ClusterRoleList(in *rbac.ClusterRoleList, out *v1.ClusterRoleList, s conversion.Scope) error {
	return autoConvert_rbac_ClusterRoleList_To_v1_ClusterRoleList(in, out, s)
}

func autoConvert_v1_PolicyRule_To_rbac_PolicyRule(in *v1.PolicyRule, out *rbac.PolicyRule, s conversion.Scope) error {
	out.Verbs = *(*[]string)(unsafe.Pointer(&in.Verbs))
	out.APIGroups = *(*[]string)(unsafe.Pointer(&in.APIGroups))
	out.Resources = *(*[]string)(unsafe.Pointer(&in.Resources))
	out.ResourceNames = *(*[]string)(unsafe.Pointer(&in.ResourceNames))
	out.NonResourceURLs = *(*[]string)(unsafe.Pointer(&in.NonResourceURLs))
	return nil
}

// Convert_v1_PolicyRule_To_rbac_PolicyRule is an autogenerated conversion function.
func Convert_v1_PolicyRule_To_rbac_PolicyRule(in *v1.PolicyRule, out *rbac.PolicyRule, s conversion.Scope) error {
	return autoConvert_v1_PolicyRule_To_rbac_PolicyRule(in, out, s)
}

func autoConvert_rbac_PolicyRule_To_v1_PolicyRule(in *rbac.PolicyRule, out *v1.PolicyRule, s conversion.Scope) error {
	out.Verbs = *(*[]string)(unsafe.Pointer(&in.Verbs))
	out.APIGroups = *(*[]string)(unsafe.Pointer(&in.APIGroups))
	out.Resources = *(*[]string)(unsafe.Pointer(&in.Resources))
	out.ResourceNames = *(*[]string)(unsafe.Pointer(&in.ResourceNames))
	out.NonResourceURLs = *(*[]string)(unsafe.Pointer(&in.NonResourceURLs))
	return nil
}

// Convert_rbac_PolicyRule_To_v1_PolicyRule is an autogenerated conversion function.
func Convert_rbac_PolicyRule_To_v1_PolicyRule(in *rbac.PolicyRule, out *v1.PolicyRule, s conversion.Scope) error {
	return autoConvert_rbac_PolicyRule_To_v1_PolicyRule(in, out, s)
}

func autoConvert_v1_Role_To_rbac_Role(in *v1.Role, out *rbac.Role, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Rules = *(*[]rbac.PolicyRule)(unsafe.Pointer(&in.Rules))
	return nil
}

// Convert_v1_Role_To_rbac_Role is an autogenerated conversion function.
func Convert_v1_Role_To_rbac_Role(in *v1.Role, out *rbac.Role, s conversion.Scope) error {
	return autoConvert_v1_Role_To_rbac_Role(in, out, s)
}

func autoConvert_rbac_Role_To_v1_Role(in *rbac.Role, out *v1.Role, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Rules = *(*[]v1.PolicyRule)(unsafe.Pointer(&in.Rules))
	return nil
}

// Convert_rbac_Role_To_v1_Role is an autogenerated conversion function.
func Convert_rbac_Role_To_v1_Role(in *rbac.Role, out *v1.Role, s conversion.Scope) error {
	return autoConvert_rbac_Role_To_v1_Role(in, out, s)
}

func autoConvert_v1_RoleBinding_To_rbac_RoleBinding(in *v1.RoleBinding, out *rbac.RoleBinding, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Subjects = *(*[]rbac.Subject)(unsafe.Pointer(&in.Subjects))
	if err := Convert_v1_RoleRef_To_rbac_RoleRef(&in.RoleRef, &out.RoleRef, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_RoleBinding_To_rbac_RoleBinding is an autogenerated conversion function.
func Convert_v1_RoleBinding_To_rbac_RoleBinding(in *v1.RoleBinding, out *rbac.RoleBinding, s conversion.Scope) error {
	return autoConvert_v1_RoleBinding_To_rbac_RoleBinding(in, out, s)
}

func autoConvert_rbac_RoleBinding_To_v1_RoleBinding(in *rbac.RoleBinding, out *v1.RoleBinding, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Subjects = *(*[]v1.Subject)(unsafe.Pointer(&in.Subjects))
	if err := Convert_rbac_RoleRef_To_v1_RoleRef(&in.RoleRef, &out.RoleRef, s); err != nil {
		return err
	}
	return nil
}

// Convert_rbac_RoleBinding_To_v1_RoleBinding is an autogenerated conversion function.
func Convert_rbac_RoleBinding_To_v1_RoleBinding(in *rbac.RoleBinding, out *v1.RoleBinding, s conversion.Scope) error {
	return autoConvert_rbac_RoleBinding_To_v1_RoleBinding(in, out, s)
}

func autoConvert_v1_RoleBindingList_To_rbac_RoleBindingList(in *v1.RoleBindingList, out *rbac.RoleBindingList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]rbac.RoleBinding)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1_RoleBindingList_To_rbac_RoleBindingList is an autogenerated conversion function.
func Convert_v1_RoleBindingList_To_rbac_RoleBindingList(in *v1.RoleBindingList, out *rbac.RoleBindingList, s conversion.Scope) error {
	return autoConvert_v1_RoleBindingList_To_rbac_RoleBindingList(in, out, s)
}

func autoConvert_rbac_RoleBindingList_To_v1_RoleBindingList(in *rbac.RoleBindingList, out *v1.RoleBindingList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]v1.RoleBinding)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_rbac_RoleBindingList_To_v1_RoleBindingList is an autogenerated conversion function.
func Convert_rbac_RoleBindingList_To_v1_RoleBindingList(in *rbac.RoleBindingList, out *v1.RoleBindingList, s conversion.Scope) error {
	return autoConvert_rbac_RoleBindingList_To_v1_RoleBindingList(in, out, s)
}

func autoConvert_v1_RoleList_To_rbac_RoleList(in *v1.RoleList, out *rbac.RoleList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]rbac.Role)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1_RoleList_To_rbac_RoleList is an autogenerated conversion function.
func Convert_v1_RoleList_To_rbac_RoleList(in *v1.RoleList, out *rbac.RoleList, s conversion.Scope) error {
	return autoConvert_v1_RoleList_To_rbac_RoleList(in, out, s)
}

func autoConvert_rbac_RoleList_To_v1_RoleList(in *rbac.RoleList, out *v1.RoleList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]v1.Role)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_rbac_RoleList_To_v1_RoleList is an autogenerated conversion function.
func Convert_rbac_RoleList_To_v1_RoleList(in *rbac.RoleList, out *v1.RoleList, s conversion.Scope) error {
	return autoConvert_rbac_RoleList_To_v1_RoleList(in, out, s)
}

func autoConvert_v1_RoleRef_To_rbac_RoleRef(in *v1.RoleRef, out *rbac.RoleRef, s conversion.Scope) error {
	out.APIGroup = in.APIGroup
	out.Kind = in.Kind
	out.Name = in.Name
	return nil
}

// Convert_v1_RoleRef_To_rbac_RoleRef is an autogenerated conversion function.
func Convert_v1_RoleRef_To_rbac_RoleRef(in *v1.RoleRef, out *rbac.RoleRef, s conversion.Scope) error {
	return autoConvert_v1_RoleRef_To_rbac_RoleRef(in, out, s)
}

func autoConvert_rbac_RoleRef_To_v1_RoleRef(in *rbac.RoleRef, out *v1.RoleRef, s conversion.Scope) error {
	out.APIGroup = in.APIGroup
	out.Kind = in.Kind
	out.Name = in.Name
	return nil
}

// Convert_rbac_RoleRef_To_v1_RoleRef is an autogenerated conversion function.
func Convert_rbac_RoleRef_To_v1_RoleRef(in *rbac.RoleRef, out *v1.RoleRef, s conversion.Scope) error {
	return autoConvert_rbac_RoleRef_To_v1_RoleRef(in, out, s)
}

func autoConvert_v1_Subject_To_rbac_Subject(in *v1.Subject, out *rbac.Subject, s conversion.Scope) error {
	out.Kind = in.Kind
	out.APIGroup = in.APIGroup
	out.Name = in.Name
	out.Namespace = in.Namespace
	return nil
}

// Convert_v1_Subject_To_rbac_Subject is an autogenerated conversion function.
func Convert_v1_Subject_To_rbac_Subject(in *v1.Subject, out *rbac.Subject, s conversion.Scope) error {
	return autoConvert_v1_Subject_To_rbac_Subject(in, out, s)
}

func autoConvert_rbac_Subject_To_v1_Subject(in *rbac.Subject, out *v1.Subject, s conversion.Scope) error {
	out.Kind = in.Kind
	out.APIGroup = in.APIGroup
	out.Name = in.Name
	out.Namespace = in.Namespace
	return nil
}

// Convert_rbac_Subject_To_v1_Subject is an autogenerated conversion function.
func Convert_rbac_Subject_To_v1_Subject(in *rbac.Subject, out *v1.Subject, s conversion.Scope) error {
	return autoConvert_rbac_Subject_To_v1_Subject(in, out, s)
}
