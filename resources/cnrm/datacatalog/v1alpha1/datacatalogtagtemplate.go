package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DataCatalogTagTemplate is the Schema for the DataCatalogTagTemplate API
type DataCatalogTagTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DataCatalogTagTemplateSpec   `json:"spec,omitempty"`
	Status DataCatalogTagTemplateStatus `json:"status,omitempty"`
}

// DataCatalogTagTemplateSpec defines the desired state of DataCatalogTagTemplate.
type DataCatalogTagTemplateSpec struct {
	// Display name for this template. Defaults to an empty string.
	// The name must contain only Unicode letters, numbers (0-9), underscores (_),
	// dashes (-), spaces ( ), and can't start or end with spaces.
	// The maximum length is 200 characters.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// Fields used to create a Tag
	Fields map[string]map[string]interface{} `json:"fields,omitempty" yaml:"fields,omitempty"`
	// Indicates whether tags created with this template are public. Public tags
	// do not require tag template access to appear in
	// [ListTags][google.cloud.datacatalog.v1.DataCatalog.ListTags] API response.
	// Additionally, you can search for a public tag by value with a
	// simple search query in addition to using a ``tag:`` predicate.
	IsPubliclyReadable bool `json:"isPubliclyReadable,omitempty" yaml:"isPubliclyReadable,omitempty"`
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// The DataCatalogTagTemplate name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// DataCatalogTagTemplateStatus defines the observed state of DataCatalogTagTemplate.
type DataCatalogTagTemplateStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
