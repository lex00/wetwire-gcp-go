package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DocumentAIProcessor is the Schema for the DocumentAIProcessor API
type DocumentAIProcessor struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DocumentAIProcessorSpec   `json:"spec,omitempty"`
	Status DocumentAIProcessorStatus `json:"status,omitempty"`
}

// DocumentAIProcessorSpec defines the desired state of DocumentAIProcessor.
type DocumentAIProcessorSpec struct {
	// The display name of the processor.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// The [KMS key](https://cloud.google.com/security-key-management) used for encryption and decryption in CMEK scenarios.
	KmsKeyRef map[string]interface{} `json:"kmsKeyRef,omitempty" yaml:"kmsKeyRef,omitempty"`
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// The GCP resource identifier. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// The processor type, such as: `OCR_PROCESSOR`, `INVOICE_PROCESSOR`. To get a list of processor types, see [FetchProcessorTypes][google.cloud.documentai.v1.DocumentProcessorService.FetchProcessorTypes].
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}

// DocumentAIProcessorStatus defines the observed state of DocumentAIProcessor.
type DocumentAIProcessorStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
