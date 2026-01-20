package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// KMSSecretCiphertext represents a kms.cnrm.cloud.google.com KMSSecretCiphertext resource.
type KMSSecretCiphertext struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KMSSecretCiphertextSpec   `json:"spec,omitempty"`
	Status KMSSecretCiphertextStatus `json:"status,omitempty"`
}

// KMSSecretCiphertextSpec defines the desired state of KMSSecretCiphertext.
type KMSSecretCiphertextSpec struct {
	// Immutable. The additional authenticated data used for integrity checks during encryption and decryption.
	AdditionalAuthenticatedData map[string]interface{} `json:"additionalAuthenticatedData,omitempty" yaml:"additionalAuthenticatedData,omitempty"`
	// Immutable. The full name of the CryptoKey that will be used to encrypt the provided plaintext.
	// Format: ''projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}/cryptoKeys/{{cryptoKey}}''.
	CryptoKey string `json:"cryptoKey,omitempty" yaml:"cryptoKey,omitempty"`
	// Immutable. The plaintext to be encrypted.
	Plaintext map[string]interface{} `json:"plaintext,omitempty" yaml:"plaintext,omitempty"`
	// Immutable. Optional. The service-generated ciphertext of the resource. Used for acquisition only. Leave unset to create a new resource.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// KMSSecretCiphertextStatus defines the observed state of KMSSecretCiphertext.
type KMSSecretCiphertextStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
