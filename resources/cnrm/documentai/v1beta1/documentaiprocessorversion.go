package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DocumentAIProcessorVersion is the Schema for the DocumentAIProcessorVersion API
type DocumentAIProcessorVersion struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DocumentAIProcessorVersionSpec   `json:"spec,omitempty"`
	Status DocumentAIProcessorVersionStatus `json:"status,omitempty"`
}

// DocumentAIProcessorVersionSpec defines the desired state of DocumentAIProcessorVersion.
type DocumentAIProcessorVersionSpec struct {
	// If set, information about the eventual deprecation of this version.
	DeprecationInfo map[string]interface{} `json:"deprecationInfo,omitempty" yaml:"deprecationInfo,omitempty"`
	// The display name of the processor version.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// The KMS key name used for encryption.
	KmsKeyNameRef map[string]interface{} `json:"kmsKeyNameRef,omitempty" yaml:"kmsKeyNameRef,omitempty"`
	// The KMS key version with which data is encrypted.
	KmsKeyVersionNameRef map[string]interface{} `json:"kmsKeyVersionNameRef,omitempty" yaml:"kmsKeyVersionNameRef,omitempty"`
	// ProcessorRef defines the resource reference to DocumentAIProcessor, which "External" field holds the GCP identifier for the KRM object.
	ProcessorRef map[string]interface{} `json:"processorRef,omitempty" yaml:"processorRef,omitempty"`
	// The DocumentAIProcessorVersion name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// DocumentAIProcessorVersionStatus defines the observed state of DocumentAIProcessorVersion.
type DocumentAIProcessorVersionStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
