package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EventarcChannel is the Schema for the EventarcChannel API
type EventarcChannel struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EventarcChannelSpec   `json:"spec,omitempty"`
	Status EventarcChannelStatus `json:"status,omitempty"`
}

// EventarcChannelSpec defines the desired state of EventarcChannel.
type EventarcChannelSpec struct {
	// Resource name of a KMS crypto key (managed by the user) used to
	// encrypt/decrypt their event data.
	// It must match the pattern
	// `projects/*/locations/*/keyRings/*/cryptoKeys/*`.
	KmsKeyRef map[string]interface{} `json:"kmsKeyRef,omitempty" yaml:"kmsKeyRef,omitempty"`
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// The name of the event provider (e.g. Eventarc SaaS partner) associated with the channel. This provider will be granted permissions to publish events to the channel.
	ProviderRef map[string]interface{} `json:"providerRef,omitempty" yaml:"providerRef,omitempty"`
	// The EventarcChannel name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// EventarcChannelStatus defines the observed state of EventarcChannel.
type EventarcChannelStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
