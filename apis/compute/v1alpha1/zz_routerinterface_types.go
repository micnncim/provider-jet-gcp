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

type RouterInterfaceObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type RouterInterfaceParameters struct {

	// IP address and range of the interface. The IP range must be in the RFC3927 link-local IP space. Changing this forces a new interface to be created.
	// +kubebuilder:validation:Optional
	IPRange *string `json:"ipRange,omitempty" tf:"ip_range,omitempty"`

	// The name or resource link to the VLAN interconnect for this interface. Changing this forces a new interface to be created. Only one of vpn_tunnel and interconnect_attachment can be specified.
	// +kubebuilder:validation:Optional
	InterconnectAttachment *string `json:"interconnectAttachment,omitempty" tf:"interconnect_attachment,omitempty"`

	// A unique name for the interface, required by GCE. Changing this forces a new interface to be created.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The ID of the project in which this interface's router belongs. If it is not provided, the provider project is used. Changing this forces a new interface to be created.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The region this interface's router sits in. If not specified, the project region will be used. Changing this forces a new interface to be created.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The name of the router this interface will be attached to. Changing this forces a new interface to be created.
	// +kubebuilder:validation:Required
	Router *string `json:"router" tf:"router,omitempty"`

	// The name or resource link to the VPN tunnel this interface will be linked to. Changing this forces a new interface to be created. Only one of vpn_tunnel and interconnect_attachment can be specified.
	// +kubebuilder:validation:Optional
	VPNTunnel *string `json:"vpnTunnel,omitempty" tf:"vpn_tunnel,omitempty"`
}

// RouterInterfaceSpec defines the desired state of RouterInterface
type RouterInterfaceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RouterInterfaceParameters `json:"forProvider"`
}

// RouterInterfaceStatus defines the observed state of RouterInterface.
type RouterInterfaceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RouterInterfaceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RouterInterface is the Schema for the RouterInterfaces API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type RouterInterface struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RouterInterfaceSpec   `json:"spec"`
	Status            RouterInterfaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RouterInterfaceList contains a list of RouterInterfaces
type RouterInterfaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RouterInterface `json:"items"`
}

// Repository type metadata.
var (
	RouterInterface_Kind             = "RouterInterface"
	RouterInterface_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RouterInterface_Kind}.String()
	RouterInterface_KindAPIVersion   = RouterInterface_Kind + "." + CRDGroupVersion.String()
	RouterInterface_GroupVersionKind = CRDGroupVersion.WithKind(RouterInterface_Kind)
)

func init() {
	SchemeBuilder.Register(&RouterInterface{}, &RouterInterfaceList{})
}
