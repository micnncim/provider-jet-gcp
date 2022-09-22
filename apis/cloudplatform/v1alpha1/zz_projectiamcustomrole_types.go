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

type ProjectIAMCustomRoleObservation struct {
	Deleted *bool `json:"deleted,omitempty" tf:"deleted,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type ProjectIAMCustomRoleParameters struct {

	// A human-readable description for the role.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The names of the permissions this role grants when bound in an IAM policy. At least one permission must be specified.
	// +kubebuilder:validation:Required
	Permissions []*string `json:"permissions" tf:"permissions,omitempty"`

	// The project that the service account will be created in. Defaults to the provider project configuration.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The camel case role id to use for this role. Cannot contain - characters.
	// +kubebuilder:validation:Required
	RoleID *string `json:"roleId" tf:"role_id,omitempty"`

	// The current launch stage of the role. Defaults to GA.
	// +kubebuilder:validation:Optional
	Stage *string `json:"stage,omitempty" tf:"stage,omitempty"`

	// A human-readable title for the role.
	// +kubebuilder:validation:Required
	Title *string `json:"title" tf:"title,omitempty"`
}

// ProjectIAMCustomRoleSpec defines the desired state of ProjectIAMCustomRole
type ProjectIAMCustomRoleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProjectIAMCustomRoleParameters `json:"forProvider"`
}

// ProjectIAMCustomRoleStatus defines the observed state of ProjectIAMCustomRole.
type ProjectIAMCustomRoleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProjectIAMCustomRoleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ProjectIAMCustomRole is the Schema for the ProjectIAMCustomRoles API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type ProjectIAMCustomRole struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProjectIAMCustomRoleSpec   `json:"spec"`
	Status            ProjectIAMCustomRoleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProjectIAMCustomRoleList contains a list of ProjectIAMCustomRoles
type ProjectIAMCustomRoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProjectIAMCustomRole `json:"items"`
}

// Repository type metadata.
var (
	ProjectIAMCustomRole_Kind             = "ProjectIAMCustomRole"
	ProjectIAMCustomRole_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ProjectIAMCustomRole_Kind}.String()
	ProjectIAMCustomRole_KindAPIVersion   = ProjectIAMCustomRole_Kind + "." + CRDGroupVersion.String()
	ProjectIAMCustomRole_GroupVersionKind = CRDGroupVersion.WithKind(ProjectIAMCustomRole_Kind)
)

func init() {
	SchemeBuilder.Register(&ProjectIAMCustomRole{}, &ProjectIAMCustomRoleList{})
}
