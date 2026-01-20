package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DialogflowEntityType represents a dialogflow.cnrm.cloud.google.com DialogflowEntityType resource.
type DialogflowEntityType struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DialogflowEntityTypeSpec   `json:"spec,omitempty"`
	Status DialogflowEntityTypeStatus `json:"status,omitempty"`
}

// DialogflowEntityTypeSpec defines the desired state of DialogflowEntityType.
type DialogflowEntityTypeSpec struct {
	// The name of this entity type to be displayed on the console.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// Enables fuzzy entity extraction during classification.
	EnableFuzzyExtraction bool `json:"enableFuzzyExtraction,omitempty" yaml:"enableFuzzyExtraction,omitempty"`
	// The collection of entity entries associated with the entity type.
	Entities []map[string]interface{} `json:"entities,omitempty" yaml:"entities,omitempty"`
	// Indicates the kind of entity type.
	// * KIND_MAP: Map entity types allow mapping of a group of synonyms to a reference value.
	// * KIND_LIST: List entity types contain a set of entries that do not map to reference values. However, list entity
	// types can contain references to other entity types (with or without aliases).
	// * KIND_REGEXP: Regexp entity types allow to specify regular expressions in entries values. Possible values: ["KIND_MAP", "KIND_LIST", "KIND_REGEXP"].
	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// DialogflowEntityTypeStatus defines the observed state of DialogflowEntityType.
type DialogflowEntityTypeStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
