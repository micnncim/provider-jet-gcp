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

type TunnelIAMBindingConditionObservation struct {
}

type TunnelIAMBindingConditionParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	Expression *string `json:"expression" tf:"expression,omitempty"`

	// +kubebuilder:validation:Required
	Title *string `json:"title" tf:"title,omitempty"`
}

type TunnelIAMBindingObservation struct {
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type TunnelIAMBindingParameters struct {

	// +kubebuilder:validation:Optional
	Condition []TunnelIAMBindingConditionParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// +kubebuilder:validation:Required
	Members []*string `json:"members" tf:"members,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// +kubebuilder:validation:Required
	Role *string `json:"role" tf:"role,omitempty"`
}

// TunnelIAMBindingSpec defines the desired state of TunnelIAMBinding
type TunnelIAMBindingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TunnelIAMBindingParameters `json:"forProvider"`
}

// TunnelIAMBindingStatus defines the observed state of TunnelIAMBinding.
type TunnelIAMBindingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TunnelIAMBindingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TunnelIAMBinding is the Schema for the TunnelIAMBindings API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type TunnelIAMBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TunnelIAMBindingSpec   `json:"spec"`
	Status            TunnelIAMBindingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TunnelIAMBindingList contains a list of TunnelIAMBindings
type TunnelIAMBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TunnelIAMBinding `json:"items"`
}

// Repository type metadata.
var (
	TunnelIAMBinding_Kind             = "TunnelIAMBinding"
	TunnelIAMBinding_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TunnelIAMBinding_Kind}.String()
	TunnelIAMBinding_KindAPIVersion   = TunnelIAMBinding_Kind + "." + CRDGroupVersion.String()
	TunnelIAMBinding_GroupVersionKind = CRDGroupVersion.WithKind(TunnelIAMBinding_Kind)
)

func init() {
	SchemeBuilder.Register(&TunnelIAMBinding{}, &TunnelIAMBindingList{})
}
