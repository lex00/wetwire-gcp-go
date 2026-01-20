package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DialogflowCXEntityType represents a dialogflowcx.cnrm.cloud.google.com DialogflowCXEntityType resource.
type DialogflowCXEntityType struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DialogflowCXEntityTypeSpec   `json:"spec,omitempty"`
	Status DialogflowCXEntityTypeStatus `json:"status,omitempty"`
}

// DialogflowCXEntityTypeSpec defines the desired state of DialogflowCXEntityType.
type DialogflowCXEntityTypeSpec struct {
	// Represents kinds of entities.
	// * AUTO_EXPANSION_MODE_UNSPECIFIED: Auto expansion disabled for the entity.
	// * AUTO_EXPANSION_MODE_DEFAULT: Allows an agent to recognize values that have not been explicitly listed in the entity. Possible values: ["AUTO_EXPANSION_MODE_DEFAULT", "AUTO_EXPANSION_MODE_UNSPECIFIED"].
	AutoExpansionMode string `json:"autoExpansionMode,omitempty" yaml:"autoExpansionMode,omitempty"`
	// The human-readable name of the entity type, unique within the agent.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// Enables fuzzy entity extraction during classification.
	EnableFuzzyExtraction bool `json:"enableFuzzyExtraction,omitempty" yaml:"enableFuzzyExtraction,omitempty"`
	// The collection of entity entries associated with the entity type.
	Entities []map[string]interface{} `json:"entities,omitempty" yaml:"entities,omitempty"`
	// Collection of exceptional words and phrases that shouldn't be matched. For example, if you have a size entity type with entry giant(an adjective), you might consider adding giants(a noun) as an exclusion.
	// If the kind of entity type is KIND_MAP, then the phrases specified by entities and excluded phrases should be mutually exclusive.
	ExcludedPhrases []map[string]interface{} `json:"excludedPhrases,omitempty" yaml:"excludedPhrases,omitempty"`
	// Indicates whether the entity type can be automatically expanded.
	// * KIND_MAP: Map entity types allow mapping of a group of synonyms to a canonical value.
	// * KIND_LIST: List entity types contain a set of entries that do not map to canonical values. However, list entity types can contain references to other entity types (with or without aliases).
	// * KIND_REGEXP: Regexp entity types allow to specify regular expressions in entries values. Possible values: ["KIND_MAP", "KIND_LIST", "KIND_REGEXP"].
	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`
	// Immutable. The language of the following fields in entityType:
	// EntityType.entities.value
	// EntityType.entities.synonyms
	// EntityType.excluded_phrases.value
	// If not specified, the agent's default language is used. Many languages are supported. Note: languages must be enabled in the agent before they can be used.
	LanguageCode string `json:"languageCode,omitempty" yaml:"languageCode,omitempty"`
	// Immutable. The agent to create a entity type for.
	// Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>.
	Parent string `json:"parent,omitempty" yaml:"parent,omitempty"`
	// Indicates whether parameters of the entity type should be redacted in log. If redaction is enabled, page parameters and intent parameters referring to the entity type will be replaced by parameter name when logging.
	Redact bool `json:"redact,omitempty" yaml:"redact,omitempty"`
	// Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// DialogflowCXEntityTypeStatus defines the observed state of DialogflowCXEntityType.
type DialogflowCXEntityTypeStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
