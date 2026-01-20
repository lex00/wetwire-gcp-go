package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// HealthcareConsentStore represents a healthcare.cnrm.cloud.google.com HealthcareConsentStore resource.
type HealthcareConsentStore struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HealthcareConsentStoreSpec   `json:"spec,omitempty"`
	Status HealthcareConsentStoreStatus `json:"status,omitempty"`
}

// HealthcareConsentStoreSpec defines the desired state of HealthcareConsentStore.
type HealthcareConsentStoreSpec struct {
	// Immutable. Identifies the dataset addressed by this request. Must be in the format
	// 'projects/{project}/locations/{location}/datasets/{dataset}'.
	Dataset string `json:"dataset,omitempty" yaml:"dataset,omitempty"`
	// Default time to live for consents in this store. Must be at least 24 hours. Updating this field will not affect the expiration time of existing consents.
	// A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
	DefaultConsentTtl string `json:"defaultConsentTtl,omitempty" yaml:"defaultConsentTtl,omitempty"`
	// If true, [consents.patch] [google.cloud.healthcare.v1.consent.UpdateConsent] creates the consent if it does not already exist.
	EnableConsentCreateOnUpdate bool `json:"enableConsentCreateOnUpdate,omitempty" yaml:"enableConsentCreateOnUpdate,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// HealthcareConsentStoreStatus defines the observed state of HealthcareConsentStore.
type HealthcareConsentStoreStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
