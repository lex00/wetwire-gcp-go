package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ConfigDeliveryResourceBundle is the Schema for the ConfigDeliveryResourceBundle API
type ConfigDeliveryResourceBundle struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ConfigDeliveryResourceBundleSpec   `json:"spec,omitempty"`
	Status ConfigDeliveryResourceBundleStatus `json:"status,omitempty"`
}

// ConfigDeliveryResourceBundleSpec defines the desired state of ConfigDeliveryResourceBundle.
type ConfigDeliveryResourceBundleSpec struct {
	// Optional. Human readable description of the `ResourceBundle`.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// The ConfigDeliveryResourceBundle name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// ConfigDeliveryResourceBundleStatus defines the observed state of ConfigDeliveryResourceBundle.
type ConfigDeliveryResourceBundleStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
