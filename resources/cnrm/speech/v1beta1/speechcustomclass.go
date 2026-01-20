package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// SpeechCustomClass is the Schema for the SpeechCustomClass API
type SpeechCustomClass struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SpeechCustomClassSpec   `json:"spec,omitempty"`
	Status SpeechCustomClassStatus `json:"status,omitempty"`
}

// SpeechCustomClassSpec defines the desired state of SpeechCustomClass.
type SpeechCustomClassSpec struct {
	// Optional. Allows users to store small amounts of arbitrary data. Both the key and the value must be 63 characters or less each. At most 100 annotations.
	Annotations map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	// Optional. User-settable, human-readable name for the CustomClass. Must be 63 characters or less.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// A collection of class items.
	Items []map[string]interface{} `json:"items,omitempty" yaml:"items,omitempty"`
	// Immutable.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// The SpeechCustomClass name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// SpeechCustomClassStatus defines the observed state of SpeechCustomClass.
type SpeechCustomClassStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
