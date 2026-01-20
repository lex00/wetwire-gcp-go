package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CloudDMSConversionWorkspace is the Schema for the CloudDMSConversionWorkspace API
type CloudDMSConversionWorkspace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CloudDMSConversionWorkspaceSpec   `json:"spec,omitempty"`
	Status CloudDMSConversionWorkspaceStatus `json:"status,omitempty"`
}

// CloudDMSConversionWorkspaceSpec defines the desired state of CloudDMSConversionWorkspace.
type CloudDMSConversionWorkspaceSpec struct {
	// Required. The destination engine details.
	Destination map[string]interface{} `json:"destination,omitempty" yaml:"destination,omitempty"`
	// Optional. The display name for the workspace.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// Optional. A generic list of settings for the workspace. The settings are database pair dependant and can indicate default behavior for the mapping rules engine or turn on or off specific features. Such examples can be: convert_foreign_key_to_interleave=true, skip_triggers=false, ignore_non_table_synonyms=true
	GlobalSettings map[string]string `json:"globalSettings,omitempty" yaml:"globalSettings,omitempty"`
	// Immutable. The location where the alloydb cluster should reside.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// The CloudDMSConversionWorkspace name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Required. The source engine details.
	Source map[string]interface{} `json:"source,omitempty" yaml:"source,omitempty"`
}

// CloudDMSConversionWorkspaceStatus defines the observed state of CloudDMSConversionWorkspace.
type CloudDMSConversionWorkspaceStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
