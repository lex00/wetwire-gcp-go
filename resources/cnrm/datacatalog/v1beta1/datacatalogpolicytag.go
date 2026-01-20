package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DataCatalogPolicyTag is the Schema for the DataCatalogPolicyTag API
type DataCatalogPolicyTag struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DataCatalogPolicyTagSpec   `json:"spec,omitempty"`
	Status DataCatalogPolicyTagStatus `json:"status,omitempty"`
}

// DataCatalogPolicyTagSpec defines the desired state of DataCatalogPolicyTag.
type DataCatalogPolicyTagSpec struct {
	// Description of this policy tag. It must: contain only unicode characters, tabs, newlines, carriage returns and page breaks; and be at most 2000 bytes long when encoded in UTF-8. If not set, defaults to an empty description. If not set, defaults to an empty description.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// User defined name of this policy tag. It must: be unique within the parent taxonomy; contain only unicode letters, numbers, underscores, dashes and spaces; not start or end with spaces; and be at most 200 bytes long when encoded in UTF-8.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// PolicyTagRef defines the resource reference to DataCatalogPolicyTag, which "External" field holds the GCP identifier for the KRM object.
	ParentPolicyTagRef map[string]interface{} `json:"parentPolicyTagRef,omitempty" yaml:"parentPolicyTagRef,omitempty"`
	// Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// TaxonomyRef defines the resource reference to DataCatalogTaxonomy, which "External" field holds the GCP identifier for the KRM object.
	TaxonomyRef map[string]interface{} `json:"taxonomyRef,omitempty" yaml:"taxonomyRef,omitempty"`
}

// DataCatalogPolicyTagStatus defines the observed state of DataCatalogPolicyTag.
type DataCatalogPolicyTagStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
