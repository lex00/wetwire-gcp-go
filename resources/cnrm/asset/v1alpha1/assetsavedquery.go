package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// AssetSavedQuery is the Schema for the AssetSavedQuery API
type AssetSavedQuery struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AssetSavedQuerySpec   `json:"spec,omitempty"`
	Status AssetSavedQueryStatus `json:"status,omitempty"`
}

// AssetSavedQuerySpec defines the desired state of AssetSavedQuery.
type AssetSavedQuerySpec struct {
	// The query content.
	Content map[string]interface{} `json:"content,omitempty" yaml:"content,omitempty"`
	// The description of this saved query. This value should be fewer than 255 characters.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// FolderRef represents the Folder that this resource belongs to.
	FolderRef map[string]interface{} `json:"folderRef,omitempty" yaml:"folderRef,omitempty"`
	// Labels applied on the resource. This value should not contain more than 10 entries. The key and value of each entry must be non-empty and fewer than 64 characters.
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	// OrganizationRef represents the Organization that this resource belongs to.
	OrganizationRef map[string]interface{} `json:"organizationRef,omitempty" yaml:"organizationRef,omitempty"`
	// The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// The AssetSavedQuery name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// AssetSavedQueryStatus defines the observed state of AssetSavedQuery.
type AssetSavedQueryStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
