package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// HealthcareHL7V2Store represents a healthcare.cnrm.cloud.google.com HealthcareHL7V2Store resource.
type HealthcareHL7V2Store struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HealthcareHL7V2StoreSpec   `json:"spec,omitempty"`
	Status HealthcareHL7V2StoreStatus `json:"status,omitempty"`
}

// HealthcareHL7V2StoreSpec defines the desired state of HealthcareHL7V2Store.
type HealthcareHL7V2StoreSpec struct {
	// Immutable. Identifies the dataset addressed by this request. Must be in the format
	// 'projects/{project}/locations/{location}/datasets/{dataset}'.
	Dataset string `json:"dataset,omitempty" yaml:"dataset,omitempty"`
	// DEPRECATED. `notification_config` is deprecated. Use `notification_configs` instead. A nested object resource.
	NotificationConfig map[string]interface{} `json:"notificationConfig,omitempty" yaml:"notificationConfig,omitempty"`
	// A list of notification configs. Each configuration uses a filter to determine whether to publish a
	// message (both Ingest & Create) on the corresponding notification destination. Only the message name
	// is sent as part of the notification. Supplied by the client.
	NotificationConfigs []map[string]interface{} `json:"notificationConfigs,omitempty" yaml:"notificationConfigs,omitempty"`
	// A nested object resource.
	ParserConfig map[string]interface{} `json:"parserConfig,omitempty" yaml:"parserConfig,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// HealthcareHL7V2StoreStatus defines the observed state of HealthcareHL7V2Store.
type HealthcareHL7V2StoreStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
