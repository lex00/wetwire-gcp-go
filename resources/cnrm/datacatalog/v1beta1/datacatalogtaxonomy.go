package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DataCatalogTaxonomy is the Schema for the datacatalog API
type DataCatalogTaxonomy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DataCatalogTaxonomySpec   `json:"spec,omitempty"`
	Status DataCatalogTaxonomyStatus `json:"status,omitempty"`
}

// DataCatalogTaxonomySpec defines the desired state of DataCatalogTaxonomy.
type DataCatalogTaxonomySpec struct {
	// A list of policy types that are activated for this taxonomy. If not set, defaults to an empty list. Possible values: ["POLICY_TYPE_UNSPECIFIED", "FINE_GRAINED_ACCESS_CONTROL"].
	ActivatedPolicyTypes []string `json:"activatedPolicyTypes,omitempty" yaml:"activatedPolicyTypes,omitempty"`
	// Description of this taxonomy. It must: contain only unicode characters, tabs, newlines, carriage returns and page breaks; and be at most 2000 bytes long when encoded in UTF-8. If not set, defaults to an empty description.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// User defined name of this taxonomy. It must: contain only unicode letters, numbers, underscores, dashes and spaces; not start or end with spaces; and be at most 200 bytes long when encoded in UTF-8.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Taxonomy location region.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`
	// Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// DataCatalogTaxonomyStatus defines the observed state of DataCatalogTaxonomy.
type DataCatalogTaxonomyStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
