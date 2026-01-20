package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DatastreamStream represents a datastream.cnrm.cloud.google.com DatastreamStream resource.
type DatastreamStream struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DatastreamStreamSpec   `json:"spec,omitempty"`
	Status DatastreamStreamStatus `json:"status,omitempty"`
}

// DatastreamStreamSpec defines the desired state of DatastreamStream.
type DatastreamStreamSpec struct {
	// Backfill strategy to automatically backfill the Stream's objects. Specific objects can be excluded.
	BackfillAll map[string]interface{} `json:"backfillAll,omitempty" yaml:"backfillAll,omitempty"`
	// Backfill strategy to disable automatic backfill for the Stream's objects.
	BackfillNone map[string]interface{} `json:"backfillNone,omitempty" yaml:"backfillNone,omitempty"`
	// Immutable. A reference to a KMS encryption key. If provided, it will be used to encrypt the data. If left blank, data
	// will be encrypted using an internal Stream-specific encryption key provisioned through KMS.
	CustomerManagedEncryptionKey string `json:"customerManagedEncryptionKey,omitempty" yaml:"customerManagedEncryptionKey,omitempty"`
	// Desired state of the Stream. Set this field to 'RUNNING' to start the stream, and 'PAUSED' to pause the stream.
	DesiredState string `json:"desiredState,omitempty" yaml:"desiredState,omitempty"`
	// Destination connection profile configuration.
	DestinationConfig map[string]interface{} `json:"destinationConfig,omitempty" yaml:"destinationConfig,omitempty"`
	// Display name.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// Immutable. The name of the location this stream is located in.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The streamId of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Source connection profile configuration.
	SourceConfig map[string]interface{} `json:"sourceConfig,omitempty" yaml:"sourceConfig,omitempty"`
}

// DatastreamStreamStatus defines the observed state of DatastreamStream.
type DatastreamStreamStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
