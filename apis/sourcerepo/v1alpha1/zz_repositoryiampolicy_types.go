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

type RepositoryIAMPolicyObservation struct {
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type RepositoryIAMPolicyParameters struct {

	// +kubebuilder:validation:Required
	PolicyData *string `json:"policyData" tf:"policy_data,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// +kubebuilder:validation:Required
	Repository *string `json:"repository" tf:"repository,omitempty"`
}

// RepositoryIAMPolicySpec defines the desired state of RepositoryIAMPolicy
type RepositoryIAMPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RepositoryIAMPolicyParameters `json:"forProvider"`
}

// RepositoryIAMPolicyStatus defines the observed state of RepositoryIAMPolicy.
type RepositoryIAMPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RepositoryIAMPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RepositoryIAMPolicy is the Schema for the RepositoryIAMPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type RepositoryIAMPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RepositoryIAMPolicySpec   `json:"spec"`
	Status            RepositoryIAMPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RepositoryIAMPolicyList contains a list of RepositoryIAMPolicys
type RepositoryIAMPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RepositoryIAMPolicy `json:"items"`
}

// Repository type metadata.
var (
	RepositoryIAMPolicy_Kind             = "RepositoryIAMPolicy"
	RepositoryIAMPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RepositoryIAMPolicy_Kind}.String()
	RepositoryIAMPolicy_KindAPIVersion   = RepositoryIAMPolicy_Kind + "." + CRDGroupVersion.String()
	RepositoryIAMPolicy_GroupVersionKind = CRDGroupVersion.WithKind(RepositoryIAMPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&RepositoryIAMPolicy{}, &RepositoryIAMPolicyList{})
}
