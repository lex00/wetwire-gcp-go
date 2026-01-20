package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DataplexLake is the Schema for the DataplexLake API
type DataplexLake struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DataplexLakeSpec   `json:"spec,omitempty"`
	Status DataplexLakeStatus `json:"status,omitempty"`
}

// DataplexLakeSpec defines the desired state of DataplexLake.
type DataplexLakeSpec struct {
	// Optional. Description of the lake.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Optional. User friendly display name.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Optional. Settings to manage lake and Dataproc Metastore service instance association.
	Metastore map[string]interface{} `json:"metastore,omitempty" yaml:"metastore,omitempty"`
	// The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// The DataplexLake name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// DataplexLakeStatus defines the observed state of DataplexLake.
type DataplexLakeStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
