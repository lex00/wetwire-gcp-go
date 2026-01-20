package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// BigLakeCatalog is the Schema for the BigLakeCatalog API
type BigLakeCatalog struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BigLakeCatalogSpec   `json:"spec,omitempty"`
	Status BigLakeCatalogStatus `json:"status,omitempty"`
}

// BigLakeCatalogSpec defines the desired state of BigLakeCatalog.
type BigLakeCatalogSpec struct {
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// The BigLakeCatalog name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// BigLakeCatalogStatus defines the observed state of BigLakeCatalog.
type BigLakeCatalogStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
