package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// KMSCryptoKeyVersion represents a kms.cnrm.cloud.google.com KMSCryptoKeyVersion resource.
type KMSCryptoKeyVersion struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KMSCryptoKeyVersionSpec   `json:"spec,omitempty"`
	Status KMSCryptoKeyVersionStatus `json:"status,omitempty"`
}

// KMSCryptoKeyVersionSpec defines the desired state of KMSCryptoKeyVersion.
type KMSCryptoKeyVersionSpec struct {
	// Immutable. The name of the cryptoKey associated with the CryptoKeyVersions.
	// Format: ''projects/{{project}}/locations/{{location}}/keyRings/{{keyring}}/cryptoKeys/{{cryptoKey}}''.
	CryptoKey string `json:"cryptoKey,omitempty" yaml:"cryptoKey,omitempty"`
	// Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// The current state of the CryptoKeyVersion. Possible values: ["PENDING_GENERATION", "ENABLED", "DISABLED", "DESTROYED", "DESTROY_SCHEDULED", "PENDING_IMPORT", "IMPORT_FAILED"].
	State string `json:"state,omitempty" yaml:"state,omitempty"`
}

// KMSCryptoKeyVersionStatus defines the observed state of KMSCryptoKeyVersion.
type KMSCryptoKeyVersionStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
