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

type ReservationAssignmentObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	State *string `json:"state,omitempty" tf:"state,omitempty"`
}

type ReservationAssignmentParameters struct {

	// The resource which will use the reservation. E.g. projects/myproject, folders/123, organizations/456.
	// +kubebuilder:validation:Required
	Assignee *string `json:"assignee" tf:"assignee,omitempty"`

	// Types of job, which could be specified when using the reservation. Possible values: JOB_TYPE_UNSPECIFIED, PIPELINE, QUERY
	// +kubebuilder:validation:Required
	JobType *string `json:"jobType" tf:"job_type,omitempty"`

	// The location for the resource
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The project for the resource
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The reservation for the resource
	// +kubebuilder:validation:Required
	Reservation *string `json:"reservation" tf:"reservation,omitempty"`
}

// ReservationAssignmentSpec defines the desired state of ReservationAssignment
type ReservationAssignmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ReservationAssignmentParameters `json:"forProvider"`
}

// ReservationAssignmentStatus defines the observed state of ReservationAssignment.
type ReservationAssignmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ReservationAssignmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ReservationAssignment is the Schema for the ReservationAssignments API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type ReservationAssignment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ReservationAssignmentSpec   `json:"spec"`
	Status            ReservationAssignmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ReservationAssignmentList contains a list of ReservationAssignments
type ReservationAssignmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ReservationAssignment `json:"items"`
}

// Repository type metadata.
var (
	ReservationAssignment_Kind             = "ReservationAssignment"
	ReservationAssignment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ReservationAssignment_Kind}.String()
	ReservationAssignment_KindAPIVersion   = ReservationAssignment_Kind + "." + CRDGroupVersion.String()
	ReservationAssignment_GroupVersionKind = CRDGroupVersion.WithKind(ReservationAssignment_Kind)
)

func init() {
	SchemeBuilder.Register(&ReservationAssignment{}, &ReservationAssignmentList{})
}
