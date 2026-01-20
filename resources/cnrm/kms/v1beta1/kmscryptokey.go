package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// KMSCryptoKey represents a kms.cnrm.cloud.google.com KMSCryptoKey resource.
type KMSCryptoKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KMSCryptoKeySpec   `json:"spec,omitempty"`
	Status KMSCryptoKeyStatus `json:"status,omitempty"`
}

// KMSCryptoKeySpec defines the desired state of KMSCryptoKey.
type KMSCryptoKeySpec struct {
	// Immutable. The period of time that versions of this key spend in the DESTROY_SCHEDULED state before transitioning to DESTROYED.
	// If not specified at creation time, the default duration is 24 hours.
	DestroyScheduledDuration string `json:"destroyScheduledDuration,omitempty" yaml:"destroyScheduledDuration,omitempty"`
	// Immutable. Whether this key may contain imported versions only.
	ImportOnly bool `json:"importOnly,omitempty" yaml:"importOnly,omitempty"`
	// The KMSKeyRing that this key belongs to.
	KeyRingRef map[string]interface{} `json:"keyRingRef,omitempty" yaml:"keyRingRef,omitempty"`
	// Immutable. The immutable purpose of this CryptoKey. See the
	// [purpose reference](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys#CryptoKeyPurpose)
	// for possible inputs.
	// Default value is "ENCRYPT_DECRYPT".
	Purpose string `json:"purpose,omitempty" yaml:"purpose,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Every time this period passes, generate a new CryptoKeyVersion and set it as the primary.
	// The first rotation will take place after the specified period. The rotation period has
	// the format of a decimal number with up to 9 fractional digits, followed by the
	// letter 's' (seconds). It must be greater than a day (ie, 86400).
	RotationPeriod string `json:"rotationPeriod,omitempty" yaml:"rotationPeriod,omitempty"`
	// Immutable. If set to true, the request will create a CryptoKey without any CryptoKeyVersions.
	// You must use the 'google_kms_key_ring_import_job' resource to import the CryptoKeyVersion.
	SkipInitialVersionCreation bool `json:"skipInitialVersionCreation,omitempty" yaml:"skipInitialVersionCreation,omitempty"`
	// A template describing settings for new crypto key versions.
	VersionTemplate map[string]interface{} `json:"versionTemplate,omitempty" yaml:"versionTemplate,omitempty"`
}

// KMSCryptoKeyStatus defines the observed state of KMSCryptoKey.
type KMSCryptoKeyStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
