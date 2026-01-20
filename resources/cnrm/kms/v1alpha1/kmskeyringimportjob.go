package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// KMSKeyRingImportJob represents a kms.cnrm.cloud.google.com KMSKeyRingImportJob resource.
type KMSKeyRingImportJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KMSKeyRingImportJobSpec   `json:"spec,omitempty"`
	Status KMSKeyRingImportJobStatus `json:"status,omitempty"`
}

// KMSKeyRingImportJobSpec defines the desired state of KMSKeyRingImportJob.
type KMSKeyRingImportJobSpec struct {
	// Immutable. It must be unique within a KeyRing and match the regular expression [a-zA-Z0-9_-]{1,63}.
	ImportJobID string `json:"importJobId,omitempty" yaml:"importJobId,omitempty"`
	// Immutable. The wrapping method to be used for incoming key material. Possible values: ["RSA_OAEP_3072_SHA1_AES_256", "RSA_OAEP_4096_SHA1_AES_256"].
	ImportMethod string `json:"importMethod,omitempty" yaml:"importMethod,omitempty"`
	// Immutable. The KeyRing that this import job belongs to.
	// Format: ''projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}''.
	KeyRing string `json:"keyRing,omitempty" yaml:"keyRing,omitempty"`
	// Immutable. The protection level of the ImportJob. This must match the protectionLevel of the
	// versionTemplate on the CryptoKey you attempt to import into. Possible values: ["SOFTWARE", "HSM", "EXTERNAL"].
	ProtectionLevel string `json:"protectionLevel,omitempty" yaml:"protectionLevel,omitempty"`
	// Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// KMSKeyRingImportJobStatus defines the observed state of KMSKeyRingImportJob.
type KMSKeyRingImportJobStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
