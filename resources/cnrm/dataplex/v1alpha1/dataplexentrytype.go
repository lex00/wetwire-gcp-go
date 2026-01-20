package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DataplexEntryType is the Schema for the DataplexEntryType API
type DataplexEntryType struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DataplexEntryTypeSpec   `json:"spec,omitempty"`
	Status DataplexEntryTypeStatus `json:"status,omitempty"`
}

// DataplexEntryTypeSpec defines the desired state of DataplexEntryType.
type DataplexEntryTypeSpec struct {
	// Authorization contains constraints on the visibility of Entries that conform to the EntryType.
	Authorization map[string]interface{} `json:"authorization,omitempty" yaml:"authorization,omitempty"`
	// Optional. Description of the EntryType.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Optional. User friendly display name.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Optional. The platform that Entries of this type belongs to.
	Platform string `json:"platform,omitempty" yaml:"platform,omitempty"`
	// The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// AspectInfo contains overriding configuration for aspects.
	RequiredAspects []map[string]interface{} `json:"requiredAspects,omitempty" yaml:"requiredAspects,omitempty"`
	// The DataplexEntryType name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Optional. The system that Entries of this type belongs to. Examples include CloudSQL, MariaDB etc
	System string `json:"system,omitempty" yaml:"system,omitempty"`
	// Optional. Indicates the classes this Entry Type belongs to, for example, TABLE, DATABASE, MODEL.
	TypeAliases []string `json:"typeAliases,omitempty" yaml:"typeAliases,omitempty"`
}

// DataplexEntryTypeStatus defines the observed state of DataplexEntryType.
type DataplexEntryTypeStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
