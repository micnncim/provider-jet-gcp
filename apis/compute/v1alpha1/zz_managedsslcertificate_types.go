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

type ManagedObservation struct {
}

type ManagedParameters struct {

	// Domains for which a managed SSL certificate will be valid.  Currently,
	// there can be up to 100 domains in this list.
	// +kubebuilder:validation:Required
	Domains []*string `json:"domains" tf:"domains,omitempty"`
}

type ManagedSSLCertificateObservation struct {
	CreationTimestamp *string `json:"creationTimestamp,omitempty" tf:"creation_timestamp,omitempty"`

	ExpireTime *string `json:"expireTime,omitempty" tf:"expire_time,omitempty"`

	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`

	SubjectAlternativeNames []*string `json:"subjectAlternativeNames,omitempty" tf:"subject_alternative_names,omitempty"`
}

type ManagedSSLCertificateParameters struct {

	// The unique identifier for the resource.
	// +kubebuilder:validation:Optional
	CertificateID *int64 `json:"certificateId,omitempty" tf:"certificate_id,omitempty"`

	// An optional description of this resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Properties relevant to a managed certificate.  These will be used if the
	// certificate is managed (as indicated by a value of 'MANAGED' in 'type').
	// +kubebuilder:validation:Optional
	Managed []ManagedParameters `json:"managed,omitempty" tf:"managed,omitempty"`

	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	//
	//
	// These are in the same namespace as the managed SSL certificates.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Enum field whose value is always 'MANAGED' - used to signal to the API
	// which type this is. Default value: "MANAGED" Possible values: ["MANAGED"]
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// ManagedSSLCertificateSpec defines the desired state of ManagedSSLCertificate
type ManagedSSLCertificateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ManagedSSLCertificateParameters `json:"forProvider"`
}

// ManagedSSLCertificateStatus defines the observed state of ManagedSSLCertificate.
type ManagedSSLCertificateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ManagedSSLCertificateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ManagedSSLCertificate is the Schema for the ManagedSSLCertificates API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfgcp}
type ManagedSSLCertificate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ManagedSSLCertificateSpec   `json:"spec"`
	Status            ManagedSSLCertificateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ManagedSSLCertificateList contains a list of ManagedSSLCertificates
type ManagedSSLCertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ManagedSSLCertificate `json:"items"`
}

// Repository type metadata.
var (
	ManagedSSLCertificate_Kind             = "ManagedSSLCertificate"
	ManagedSSLCertificate_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ManagedSSLCertificate_Kind}.String()
	ManagedSSLCertificate_KindAPIVersion   = ManagedSSLCertificate_Kind + "." + CRDGroupVersion.String()
	ManagedSSLCertificate_GroupVersionKind = CRDGroupVersion.WithKind(ManagedSSLCertificate_Kind)
)

func init() {
	SchemeBuilder.Register(&ManagedSSLCertificate{}, &ManagedSSLCertificateList{})
}
