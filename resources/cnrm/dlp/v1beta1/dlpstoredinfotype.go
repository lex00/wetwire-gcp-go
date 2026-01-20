package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DLPStoredInfoType represents a dlp.cnrm.cloud.google.com DLPStoredInfoType resource.
type DLPStoredInfoType struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DLPStoredInfoTypeSpec   `json:"spec,omitempty"`
	Status DLPStoredInfoTypeStatus `json:"status,omitempty"`
}

// DLPStoredInfoTypeSpec defines the desired state of DLPStoredInfoType.
type DLPStoredInfoTypeSpec struct {
	// Description of the StoredInfoType (max 256 characters).
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Store dictionary-based CustomInfoType.
	Dictionary map[string]interface{} `json:"dictionary,omitempty" yaml:"dictionary,omitempty"`
	// Display name of the StoredInfoType (max 256 characters).
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// StoredInfoType where findings are defined by a dictionary of phrases.
	LargeCustomDictionary map[string]interface{} `json:"largeCustomDictionary,omitempty" yaml:"largeCustomDictionary,omitempty"`
	// Immutable. The location of the resource
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Immutable. The Organization that this resource belongs to. Only one of [organizationRef, projectRef] may be specified.
	OrganizationRef map[string]interface{} `json:"organizationRef,omitempty" yaml:"organizationRef,omitempty"`
	// Immutable. The Project that this resource belongs to. Only one of [organizationRef, projectRef] may be specified.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Store regular expression-based StoredInfoType.
	Regex map[string]interface{} `json:"regex,omitempty" yaml:"regex,omitempty"`
	// Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// DLPStoredInfoTypeStatus defines the observed state of DLPStoredInfoType.
type DLPStoredInfoTypeStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
