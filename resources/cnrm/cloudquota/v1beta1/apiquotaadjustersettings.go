package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// APIQuotaAdjusterSettings is the Schema for the APIQuotaAdjusterSettings API
type APIQuotaAdjusterSettings struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   APIQuotaAdjusterSettingsSpec   `json:"spec,omitempty"`
	Status APIQuotaAdjusterSettingsStatus `json:"status,omitempty"`
}

// APIQuotaAdjusterSettingsSpec defines the desired state of APIQuotaAdjusterSettings.
type APIQuotaAdjusterSettingsSpec struct {
	// Required. The configured value of the enablement at the given resource.
	Enablement string `json:"enablement,omitempty" yaml:"enablement,omitempty"`
	// The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// The APIQuotaAdjusterSettings name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// APIQuotaAdjusterSettingsStatus defines the observed state of APIQuotaAdjusterSettings.
type APIQuotaAdjusterSettingsStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
