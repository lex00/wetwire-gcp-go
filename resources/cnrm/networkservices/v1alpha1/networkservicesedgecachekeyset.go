package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NetworkServicesEdgeCacheKeyset represents a networkservices.cnrm.cloud.google.com NetworkServicesEdgeCacheKeyset resource.
type NetworkServicesEdgeCacheKeyset struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NetworkServicesEdgeCacheKeysetSpec   `json:"spec,omitempty"`
	Status NetworkServicesEdgeCacheKeysetStatus `json:"status,omitempty"`
}

// NetworkServicesEdgeCacheKeysetSpec defines the desired state of NetworkServicesEdgeCacheKeyset.
type NetworkServicesEdgeCacheKeysetSpec struct {
	// A human-readable description of the resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// An ordered list of Ed25519 public keys to use for validating signed requests.
	// You must specify 'public_keys' or 'validation_shared_keys' (or both). The keys in 'public_keys' are checked first.
	// You may specify no more than one Google-managed public key.
	// If you specify 'public_keys', you must specify at least one (1) key and may specify up to three (3) keys.
	// Ed25519 public keys are not secret, and only allow Google to validate a request was signed by your corresponding private key.
	// Ensure that the private key is kept secret, and that only authorized users can add public keys to a keyset.
	PublicKey []map[string]interface{} `json:"publicKey,omitempty" yaml:"publicKey,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// An ordered list of shared keys to use for validating signed requests.
	// Shared keys are secret.  Ensure that only authorized users can add 'validation_shared_keys' to a keyset.
	// You can rotate keys by appending (pushing) a new key to the list of 'validation_shared_keys' and removing any superseded keys.
	// You must specify 'public_keys' or 'validation_shared_keys' (or both). The keys in 'public_keys' are checked first.
	ValidationSharedKeys []map[string]interface{} `json:"validationSharedKeys,omitempty" yaml:"validationSharedKeys,omitempty"`
}

// NetworkServicesEdgeCacheKeysetStatus defines the observed state of NetworkServicesEdgeCacheKeyset.
type NetworkServicesEdgeCacheKeysetStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
