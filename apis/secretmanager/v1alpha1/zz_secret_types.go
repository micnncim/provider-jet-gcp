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

type CustomerManagedEncryptionObservation struct {
}

type CustomerManagedEncryptionParameters struct {

	// Describes the Cloud KMS encryption key that will be used to protect destination secret.
	// +kubebuilder:validation:Required
	KMSKeyName *string `json:"kmsKeyName" tf:"kms_key_name,omitempty"`
}

type ReplicasObservation struct {
}

type ReplicasParameters struct {

	// Customer Managed Encryption for the secret.
	// +kubebuilder:validation:Optional
	CustomerManagedEncryption []CustomerManagedEncryptionParameters `json:"customerManagedEncryption,omitempty" tf:"customer_managed_encryption,omitempty"`

	// The canonical IDs of the location to replicate data. For example: "us-east1".
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`
}

type ReplicationObservation struct {
}

type ReplicationParameters struct {

	// The Secret will automatically be replicated without any restrictions.
	// +kubebuilder:validation:Optional
	Automatic *bool `json:"automatic,omitempty" tf:"automatic,omitempty"`

	// The Secret will automatically be replicated without any restrictions.
	// +kubebuilder:validation:Optional
	UserManaged []UserManagedParameters `json:"userManaged,omitempty" tf:"user_managed,omitempty"`
}

type RotationObservation struct {
}

type RotationParameters struct {

	// Timestamp in UTC at which the Secret is scheduled to rotate.
	// A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	// +kubebuilder:validation:Optional
	NextRotationTime *string `json:"nextRotationTime,omitempty" tf:"next_rotation_time,omitempty"`

	// The Duration between rotation notifications. Must be in seconds and at least 3600s (1h) and at most 3153600000s (100 years).
	// If rotationPeriod is set, 'next_rotation_time' must be set. 'next_rotation_time' will be advanced by this period when the service automatically sends rotation notifications.
	// +kubebuilder:validation:Optional
	RotationPeriod *string `json:"rotationPeriod,omitempty" tf:"rotation_period,omitempty"`
}

type SecretObservation struct {
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type SecretParameters struct {

	// Timestamp in UTC when the Secret is scheduled to expire. This is always provided on output, regardless of what was sent on input.
	// A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	// +kubebuilder:validation:Optional
	ExpireTime *string `json:"expireTime,omitempty" tf:"expire_time,omitempty"`

	// The labels assigned to this Secret.
	//
	// Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes,
	// and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}
	//
	// Label values must be between 0 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes,
	// and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63}
	//
	// No more than 64 labels can be assigned to a given resource.
	//
	// An object containing a list of "key": value pairs. Example:
	// { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The replication policy of the secret data attached to the Secret. It cannot be changed
	// after the Secret has been created.
	// +kubebuilder:validation:Required
	Replication []ReplicationParameters `json:"replication" tf:"replication,omitempty"`

	// The rotation time and period for a Secret. At 'next_rotation_time', Secret Manager will send a Pub/Sub notification to the topics configured on the Secret. 'topics' must be set to configure rotation.
	// +kubebuilder:validation:Optional
	Rotation []RotationParameters `json:"rotation,omitempty" tf:"rotation,omitempty"`

	// This must be unique within the project.
	// +kubebuilder:validation:Required
	SecretID *string `json:"secretId" tf:"secret_id,omitempty"`

	// The TTL for the Secret.
	// A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
	// +kubebuilder:validation:Optional
	TTL *string `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// A list of up to 10 Pub/Sub topics to which messages are published when control plane operations are called on the secret or its versions.
	// +kubebuilder:validation:Optional
	Topics []TopicsParameters `json:"topics,omitempty" tf:"topics,omitempty"`
}

type TopicsObservation struct {
}

type TopicsParameters struct {

	// The resource name of the Pub/Sub topic that will be published to, in the following format: projects/*/topics/*.
	// For publication to succeed, the Secret Manager Service Agent service account must have pubsub.publisher permissions on the topic.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

type UserManagedObservation struct {
}

type UserManagedParameters struct {

	// The list of Replicas for this Secret. Cannot be empty.
	// +kubebuilder:validation:Required
	Replicas []ReplicasParameters `json:"replicas" tf:"replicas,omitempty"`
}

// SecretSpec defines the desired state of Secret
type SecretSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecretParameters `json:"forProvider"`
}

// SecretStatus defines the observed state of Secret.
type SecretStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecretObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Secret is the Schema for the Secrets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type Secret struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SecretSpec   `json:"spec"`
	Status            SecretStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecretList contains a list of Secrets
type SecretList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Secret `json:"items"`
}

// Repository type metadata.
var (
	Secret_Kind             = "Secret"
	Secret_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Secret_Kind}.String()
	Secret_KindAPIVersion   = Secret_Kind + "." + CRDGroupVersion.String()
	Secret_GroupVersionKind = CRDGroupVersion.WithKind(Secret_Kind)
)

func init() {
	SchemeBuilder.Register(&Secret{}, &SecretList{})
}
