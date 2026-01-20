package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// IAPSettings is the Schema for the IAPSettings API
type IAPSettings struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IAPSettingsSpec   `json:"spec,omitempty"`
	Status IAPSettingsStatus `json:"status,omitempty"`
}

// IAPSettingsSpec defines the desired state of IAPSettings.
type IAPSettingsSpec struct {
	// Top level wrapper for all access related setting in IAP
	AccessSettings map[string]interface{} `json:"accessSettings,omitempty" yaml:"accessSettings,omitempty"`
	// Project-wide App Engine service settings
	AppEngineRef map[string]interface{} `json:"appEngineRef,omitempty" yaml:"appEngineRef,omitempty"`
	// Top level wrapper for all application related settings in IAP
	ApplicationSettings map[string]interface{} `json:"applicationSettings,omitempty" yaml:"applicationSettings,omitempty"`
	// Project-wide Compute service settings
	ComputeServiceRef map[string]interface{} `json:"computeServiceRef,omitempty" yaml:"computeServiceRef,omitempty"`
	// Folder-level settings
	FolderRef map[string]interface{} `json:"folderRef,omitempty" yaml:"folderRef,omitempty"`
	// Organization-level settings
	OrganizationRef map[string]interface{} `json:"organizationRef,omitempty" yaml:"organizationRef,omitempty"`
	// Project-level settings
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Project-wide web service settings
	ProjectWebRef map[string]interface{} `json:"projectWebRef,omitempty" yaml:"projectWebRef,omitempty"`
	// The IAPSettings name.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// IAPSettingsStatus defines the observed state of IAPSettings.
type IAPSettingsStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
