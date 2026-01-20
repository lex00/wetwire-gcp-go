package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// HealthcareDICOMStore represents a healthcare.cnrm.cloud.google.com HealthcareDICOMStore resource.
type HealthcareDICOMStore struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HealthcareDICOMStoreSpec   `json:"spec,omitempty"`
	Status HealthcareDICOMStoreStatus `json:"status,omitempty"`
}

// HealthcareDICOMStoreSpec defines the desired state of HealthcareDICOMStore.
type HealthcareDICOMStoreSpec struct {
	// Immutable. Identifies the dataset addressed by this request. Must be in the format
	// 'projects/{project}/locations/{location}/datasets/{dataset}'.
	Dataset string `json:"dataset,omitempty" yaml:"dataset,omitempty"`
	// A nested object resource.
	NotificationConfig map[string]interface{} `json:"notificationConfig,omitempty" yaml:"notificationConfig,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// To enable streaming to BigQuery, configure the streamConfigs object in your DICOM store.
	// streamConfigs is an array, so you can specify multiple BigQuery destinations. You can stream metadata from a single DICOM store to up to five BigQuery tables in a BigQuery dataset.
	StreamConfigs []map[string]interface{} `json:"streamConfigs,omitempty" yaml:"streamConfigs,omitempty"`
}

// HealthcareDICOMStoreStatus defines the observed state of HealthcareDICOMStore.
type HealthcareDICOMStoreStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
