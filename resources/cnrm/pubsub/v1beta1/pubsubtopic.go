package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PubSubTopic represents a pubsub.cnrm.cloud.google.com PubSubTopic resource.
type PubSubTopic struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PubSubTopicSpec   `json:"spec,omitempty"`
	Status PubSubTopicStatus `json:"status,omitempty"`
}

// PubSubTopicSpec defines the desired state of PubSubTopic.
type PubSubTopicSpec struct {
	// The KMSCryptoKey to be used to protect access to messages published
	// on this topic. Your project's Pub/Sub service account
	// ('service-{{PROJECT_NUMBER}}@gcp-sa-pubsub.iam.gserviceaccount.com')
	// must have 'roles/cloudkms.cryptoKeyEncrypterDecrypter' to use this
	// feature.
	KmsKeyRef map[string]interface{} `json:"kmsKeyRef,omitempty" yaml:"kmsKeyRef,omitempty"`
	// Indicates the minimum duration to retain a message after it is published
	// to the topic. If this field is set, messages published to the topic in
	// the last messageRetentionDuration are always available to subscribers.
	// For instance, it allows any attached subscription to seek to a timestamp
	// that is up to messageRetentionDuration in the past. If this field is not
	// set, message retention is controlled by settings on individual subscriptions.
	// Cannot be more than 31 days or less than 10 minutes.
	MessageRetentionDuration string `json:"messageRetentionDuration,omitempty" yaml:"messageRetentionDuration,omitempty"`
	// Policy constraining the set of Google Cloud Platform regions where
	// messages published to the topic may be stored. If not present, then no
	// constraints are in effect.
	MessageStoragePolicy map[string]interface{} `json:"messageStoragePolicy,omitempty" yaml:"messageStoragePolicy,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Settings for validating messages published against a schema.
	SchemaSettings map[string]interface{} `json:"schemaSettings,omitempty" yaml:"schemaSettings,omitempty"`
}

// PubSubTopicStatus defines the observed state of PubSubTopic.
type PubSubTopicStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
