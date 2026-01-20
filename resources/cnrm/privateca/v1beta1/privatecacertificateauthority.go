package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PrivateCACertificateAuthority represents a privateca.cnrm.cloud.google.com PrivateCACertificateAuthority resource.
type PrivateCACertificateAuthority struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PrivateCACertificateAuthoritySpec   `json:"spec,omitempty"`
	Status PrivateCACertificateAuthorityStatus `json:"status,omitempty"`
}

// PrivateCACertificateAuthoritySpec defines the desired state of PrivateCACertificateAuthority.
type PrivateCACertificateAuthoritySpec struct {
	// Immutable.
	CaPoolRef map[string]interface{} `json:"caPoolRef,omitempty" yaml:"caPoolRef,omitempty"`
	// Immutable. Required. Immutable. The config used to create a self-signed X.509 certificate or CSR.
	Config map[string]interface{} `json:"config,omitempty" yaml:"config,omitempty"`
	// Immutable.
	GcsBucketRef map[string]interface{} `json:"gcsBucketRef,omitempty" yaml:"gcsBucketRef,omitempty"`
	// Immutable. Required. Immutable. Used when issuing certificates for this CertificateAuthority. If this CertificateAuthority is a self-signed CertificateAuthority, this key is also used to sign the self-signed CA certificate. Otherwise, it is used to sign a CSR.
	KeySpec map[string]interface{} `json:"keySpec,omitempty" yaml:"keySpec,omitempty"`
	// Immutable. Required. The desired lifetime of the CA certificate. Used to create the "not_before_time" and "not_after_time" fields inside an X.509 certificate.
	Lifetime string `json:"lifetime,omitempty" yaml:"lifetime,omitempty"`
	// Immutable. The location for the resource
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Immutable. The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Immutable. Required. Immutable. The Type of this CertificateAuthority. Possible values: SELF_SIGNED, SUBORDINATE
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}

// PrivateCACertificateAuthorityStatus defines the observed state of PrivateCACertificateAuthority.
type PrivateCACertificateAuthorityStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
