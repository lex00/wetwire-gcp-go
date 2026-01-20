package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DataCatalogTag is the Schema for the DataCatalogTag API
type DataCatalogTag struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DataCatalogTagSpec   `json:"spec,omitempty"`
	Status DataCatalogTagStatus `json:"status,omitempty"`
}

// DataCatalogTagSpec defines the desired state of DataCatalogTag.
type DataCatalogTagSpec struct {
	// Resources like entry can have schemas associated with them. This scope
	// allows you to attach tags to an individual column based on that schema.
	// To attach a tag to a nested column, separate column names with a dot
	// (`.`). Example: `column.nested_column`.
	Column string `json:"column,omitempty" yaml:"column,omitempty"`
	// Required. Reference to the DataCatalogEntry that owns this Tag. The entry must be in the same project and location as the tag.
	EntryRef map[string]interface{} `json:"entryRef,omitempty" yaml:"entryRef,omitempty"`
	Fields map[string]map[string]interface{} `json:"fields,omitempty" yaml:"fields,omitempty"`
	// The DataCatalogTag name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Required. The resource name of the tag template this tag uses.
	// This field cannot be modified after creation.
	TemplateRef map[string]interface{} `json:"templateRef,omitempty" yaml:"templateRef,omitempty"`
}

// DataCatalogTagStatus defines the observed state of DataCatalogTag.
type DataCatalogTagStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
