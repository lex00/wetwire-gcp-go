package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// SpeechPhraseSet is the Schema for the SpeechPhraseSet API
type SpeechPhraseSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SpeechPhraseSetSpec   `json:"spec,omitempty"`
	Status SpeechPhraseSetStatus `json:"status,omitempty"`
}

// SpeechPhraseSetSpec defines the desired state of SpeechPhraseSet.
type SpeechPhraseSetSpec struct {
	// Allows users to store small amounts of arbitrary data. Both the key and the value must be 63 characters or less each. At most 100 annotations.
	Annotations map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	// User-settable, human-readable name for the PhraseSet. Must be 63 characters or less.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// Immutable.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// A list of word and phrases.
	Phrases []map[string]interface{} `json:"phrases,omitempty" yaml:"phrases,omitempty"`
	// The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// The SpeechPhraseSet name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// SpeechPhraseSetStatus defines the observed state of SpeechPhraseSet.
type SpeechPhraseSetStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
