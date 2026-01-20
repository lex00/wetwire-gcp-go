package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DLPDeidentifyTemplate represents a dlp.cnrm.cloud.google.com DLPDeidentifyTemplate resource.
type DLPDeidentifyTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DLPDeidentifyTemplateSpec   `json:"spec,omitempty"`
	Status DLPDeidentifyTemplateStatus `json:"status,omitempty"`
}

// DLPDeidentifyTemplateSpec defines the desired state of DLPDeidentifyTemplate.
type DLPDeidentifyTemplateSpec struct {
	// The core content of the template.
	DeidentifyConfig map[string]interface{} `json:"deidentifyConfig,omitempty" yaml:"deidentifyConfig,omitempty"`
	// Short description (max 256 chars).
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Display name (max 256 chars).
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// Immutable. The location of the resource
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Immutable. The Organization that this resource belongs to. Only one of [organizationRef, projectRef] may be specified.
	OrganizationRef map[string]interface{} `json:"organizationRef,omitempty" yaml:"organizationRef,omitempty"`
	// Immutable. The Project that this resource belongs to. Only one of [organizationRef, projectRef] may be specified.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// DLPDeidentifyTemplateStatus defines the observed state of DLPDeidentifyTemplate.
type DLPDeidentifyTemplateStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
