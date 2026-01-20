package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DLPInspectTemplate represents a dlp.cnrm.cloud.google.com DLPInspectTemplate resource.
type DLPInspectTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DLPInspectTemplateSpec   `json:"spec,omitempty"`
	Status DLPInspectTemplateStatus `json:"status,omitempty"`
}

// DLPInspectTemplateSpec defines the desired state of DLPInspectTemplate.
type DLPInspectTemplateSpec struct {
	// Short description (max 256 chars).
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Display name (max 256 chars).
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// The core content of the template. Configuration of the scanning process.
	InspectConfig map[string]interface{} `json:"inspectConfig,omitempty" yaml:"inspectConfig,omitempty"`
	// Immutable. The location of the resource
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Immutable. The Organization that this resource belongs to. Only one of [organizationRef, projectRef] may be specified.
	OrganizationRef map[string]interface{} `json:"organizationRef,omitempty" yaml:"organizationRef,omitempty"`
	// Immutable. The Project that this resource belongs to. Only one of [organizationRef, projectRef] may be specified.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// DLPInspectTemplateStatus defines the observed state of DLPInspectTemplate.
type DLPInspectTemplateStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
