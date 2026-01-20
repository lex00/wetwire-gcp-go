package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// SpeechRecognizer is the Schema for the SpeechRecognizer API
type SpeechRecognizer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SpeechRecognizerSpec   `json:"spec,omitempty"`
	Status SpeechRecognizerStatus `json:"status,omitempty"`
}

// SpeechRecognizerSpec defines the desired state of SpeechRecognizer.
type SpeechRecognizerSpec struct {
	// Allows users to store small amounts of arbitrary data. Both the key and the value must be 63 characters or less each. At most 100 annotations.
	Annotations map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	// Default configuration to use for requests with this Recognizer. This can be overwritten by inline configuration in the [RecognizeRequest.config][google.cloud.speech.v2.RecognizeRequest.config] field.
	DefaultRecognitionConfig map[string]interface{} `json:"defaultRecognitionConfig,omitempty" yaml:"defaultRecognitionConfig,omitempty"`
	// User-settable, human-readable name for the Recognizer. Must be 63 characters or less.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// Immutable.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// The SpeechRecognizer name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// SpeechRecognizerStatus defines the observed state of SpeechRecognizer.
type SpeechRecognizerStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
