package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// KMSImportJob is the Schema for the KMSImportJob API
type KMSImportJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KMSImportJobSpec   `json:"spec,omitempty"`
	Status KMSImportJobStatus `json:"status,omitempty"`
}

// KMSImportJobSpec defines the desired state of KMSImportJob.
type KMSImportJobSpec struct {
	// Required. Immutable. The wrapping method to be used for incoming key material.
	ImportMethod string `json:"importMethod,omitempty" yaml:"importMethod,omitempty"`
	// KMSKeyRingRef defines the resource reference to KMSKeyRing, which "External" field holds the GCP identifier for the KRM object.
	KmsKeyRingRef map[string]interface{} `json:"kmsKeyRingRef,omitempty" yaml:"kmsKeyRingRef,omitempty"`
	// Required. Immutable. The protection level of the [ImportJob][google.cloud.kms.v1.ImportJob]. This must match the [protection_level][google.cloud.kms.v1.CryptoKeyVersionTemplate.protection_level] of the [version_template][google.cloud.kms.v1.CryptoKey.version_template] on the [CryptoKey][google.cloud.kms.v1.CryptoKey] you attempt to import into.
	ProtectionLevel string `json:"protectionLevel,omitempty" yaml:"protectionLevel,omitempty"`
	// The KMSImportJob name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// KMSImportJobStatus defines the observed state of KMSImportJob.
type KMSImportJobStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
