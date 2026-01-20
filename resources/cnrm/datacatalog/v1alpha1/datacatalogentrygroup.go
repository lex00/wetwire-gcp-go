package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DataCatalogEntryGroup is the Schema for the DataCatalogEntryGroup API
type DataCatalogEntryGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DataCatalogEntryGroupSpec   `json:"spec,omitempty"`
	Status DataCatalogEntryGroupStatus `json:"status,omitempty"`
}

// DataCatalogEntryGroupSpec defines the desired state of DataCatalogEntryGroup.
type DataCatalogEntryGroupSpec struct {
	// Entry group description. Can consist of several sentences or paragraphs that describe the entry group contents. Default value is an empty string.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// A short name to identify the entry group, for example, "analytics data - jan 2011". Default value is an empty string.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// The DataCatalogEntryGroup name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Optional. When set to [true], it means DataCatalog EntryGroup was transferred to Dataplex Catalog Service. It makes EntryGroup and its Entries to be read-only in DataCatalog. However, new Tags on EntryGroup and its Entries can be created. After setting the flag to [true] it cannot be unset.
	TransferredToDataplex bool `json:"transferredToDataplex,omitempty" yaml:"transferredToDataplex,omitempty"`
}

// DataCatalogEntryGroupStatus defines the observed state of DataCatalogEntryGroup.
type DataCatalogEntryGroupStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
