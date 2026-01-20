package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EventarcGoogleChannelConfig is the Schema for the EventarcGoogleChannelConfig API
type EventarcGoogleChannelConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EventarcGoogleChannelConfigSpec   `json:"spec,omitempty"`
	Status EventarcGoogleChannelConfigStatus `json:"status,omitempty"`
}

// EventarcGoogleChannelConfigSpec defines the desired state of EventarcGoogleChannelConfig.
type EventarcGoogleChannelConfigSpec struct {
	// Optional. Resource name of a KMS crypto key (managed by the user) used to encrypt/decrypt their event data.
	CryptoKeyRef map[string]interface{} `json:"cryptoKeyRef,omitempty" yaml:"cryptoKeyRef,omitempty"`
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// The user-provided name of the EventarcGoogleChannelConfig. If not specified, the name of the KRM resource will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// EventarcGoogleChannelConfigStatus defines the observed state of EventarcGoogleChannelConfig.
type EventarcGoogleChannelConfigStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
