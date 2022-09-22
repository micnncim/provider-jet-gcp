/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type BucketAccessControlObservation struct {
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type BucketAccessControlParameters struct {

	// The name of the bucket.
	// +kubebuilder:validation:Required
	Bucket *string `json:"bucket" tf:"bucket,omitempty"`

	// The entity holding the permission, in one of the following forms:
	// user-userId
	// user-email
	// group-groupId
	// group-email
	// domain-domain
	// project-team-projectId
	// allUsers
	// allAuthenticatedUsers
	// Examples:
	// The user liz@example.com would be user-liz@example.com.
	// The group example@googlegroups.com would be
	// group-example@googlegroups.com.
	// To refer to all members of the Google Apps for Business domain
	// example.com, the entity would be domain-example.com.
	// +kubebuilder:validation:Required
	Entity *string `json:"entity" tf:"entity,omitempty"`

	// The access permission for the entity. Possible values: ["OWNER", "READER", "WRITER"]
	// +kubebuilder:validation:Optional
	Role *string `json:"role,omitempty" tf:"role,omitempty"`
}

// BucketAccessControlSpec defines the desired state of BucketAccessControl
type BucketAccessControlSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BucketAccessControlParameters `json:"forProvider"`
}

// BucketAccessControlStatus defines the observed state of BucketAccessControl.
type BucketAccessControlStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BucketAccessControlObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BucketAccessControl is the Schema for the BucketAccessControls API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type BucketAccessControl struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BucketAccessControlSpec   `json:"spec"`
	Status            BucketAccessControlStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BucketAccessControlList contains a list of BucketAccessControls
type BucketAccessControlList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BucketAccessControl `json:"items"`
}

// Repository type metadata.
var (
	BucketAccessControl_Kind             = "BucketAccessControl"
	BucketAccessControl_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BucketAccessControl_Kind}.String()
	BucketAccessControl_KindAPIVersion   = BucketAccessControl_Kind + "." + CRDGroupVersion.String()
	BucketAccessControl_GroupVersionKind = CRDGroupVersion.WithKind(BucketAccessControl_Kind)
)

func init() {
	SchemeBuilder.Register(&BucketAccessControl{}, &BucketAccessControlList{})
}
