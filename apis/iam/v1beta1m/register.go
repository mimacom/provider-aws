/*
Copyright 2024 The Crossplane Authors.

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
// +kubebuilder:object:generate=true
// +groupName=iam.aws.m.crossplane.io
// +versionName=v1beta1

// Package v1beta1m contains the namespaced (Crossplane v2) IAM managed
// resources: User, Policy, UserPolicyAttachment and AccessKey. Their
// cluster-scoped siblings (Role, Group, ...) remain in the
// iam.aws.crossplane.io package.
package v1beta1m

import (
	"reflect"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

// Package type metadata.
const (
	CRDGroup   = "iam.aws.m.crossplane.io"
	CRDVersion = "v1beta1"
)

var (
	// SchemeGroupVersion is group version used to register these objects
	SchemeGroupVersion = schema.GroupVersion{Group: CRDGroup, Version: CRDVersion}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	SchemeBuilder = &scheme.Builder{GroupVersion: SchemeGroupVersion}
)

// User type metadata.
var (
	UserKind             = reflect.TypeOf(User{}).Name()
	UserGroupKind        = schema.GroupKind{Group: CRDGroup, Kind: UserKind}.String()
	UserKindAPIVersion   = UserKind + "." + SchemeGroupVersion.String()
	UserGroupVersionKind = SchemeGroupVersion.WithKind(UserKind)
)

// UserPolicyAttachment type metadata.
var (
	UserPolicyAttachmentKind             = reflect.TypeOf(UserPolicyAttachment{}).Name()
	UserPolicyAttachmentGroupKind        = schema.GroupKind{Group: CRDGroup, Kind: UserPolicyAttachmentKind}.String()
	UserPolicyAttachmentKindAPIVersion   = UserPolicyAttachmentKind + "." + SchemeGroupVersion.String()
	UserPolicyAttachmentGroupVersionKind = SchemeGroupVersion.WithKind(UserPolicyAttachmentKind)
)

// Policy type metadata.
var (
	PolicyKind             = reflect.TypeOf(Policy{}).Name()
	PolicyGroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PolicyKind}.String()
	PolicyKindAPIVersion   = PolicyKind + "." + SchemeGroupVersion.String()
	PolicyGroupVersionKind = SchemeGroupVersion.WithKind(PolicyKind)
)

// AccessKey type metadata.
var (
	AccessKeyKind             = reflect.TypeOf(AccessKey{}).Name()
	AccessKeyGroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AccessKeyKind}.String()
	AccessKeyKindAPIVersion   = AccessKeyKind + "." + SchemeGroupVersion.String()
	AccessKeyGroupVersionKind = SchemeGroupVersion.WithKind(AccessKeyKind)
)

func init() {
	SchemeBuilder.Register(&User{}, &UserList{})
	SchemeBuilder.Register(&Policy{}, &PolicyList{})
	SchemeBuilder.Register(&UserPolicyAttachment{}, &UserPolicyAttachmentList{})
	SchemeBuilder.Register(&AccessKey{}, &AccessKeyList{})
}
