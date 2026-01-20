package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PrivateCACertificate represents a privateca.cnrm.cloud.google.com PrivateCACertificate resource.
type PrivateCACertificate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PrivateCACertificateSpec   `json:"spec,omitempty"`
	Status PrivateCACertificateStatus `json:"status,omitempty"`
}

// PrivateCACertificateSpec defines the desired state of PrivateCACertificate.
type PrivateCACertificateSpec struct {
	// Immutable.
	CaPoolRef map[string]interface{} `json:"caPoolRef,omitempty" yaml:"caPoolRef,omitempty"`
	// Immutable.
	CertificateAuthorityRef map[string]interface{} `json:"certificateAuthorityRef,omitempty" yaml:"certificateAuthorityRef,omitempty"`
	// Immutable.
	CertificateTemplateRef map[string]interface{} `json:"certificateTemplateRef,omitempty" yaml:"certificateTemplateRef,omitempty"`
	// Immutable. Immutable. A description of the certificate and key that does not require X.509 or ASN.1.
	Config map[string]interface{} `json:"config,omitempty" yaml:"config,omitempty"`
	// Immutable. Required. Immutable. The desired lifetime of a certificate. Used to create the "not_before_time" and "not_after_time" fields inside an X.509 certificate. Note that the lifetime may be truncated if it would extend past the life of any certificate authority in the issuing chain.
	Lifetime string `json:"lifetime,omitempty" yaml:"lifetime,omitempty"`
	// Immutable. The location for the resource
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Immutable. Immutable. A pem-encoded X.509 certificate signing request (CSR).
	PemCsr string `json:"pemCsr,omitempty" yaml:"pemCsr,omitempty"`
	// Immutable. The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Immutable. Immutable. Specifies how the Certificate's identity fields are to be decided. If this is omitted, the `DEFAULT` subject mode will be used. Possible values: SUBJECT_REQUEST_MODE_UNSPECIFIED, DEFAULT, REFLECTED_SPIFFE
	SubjectMode string `json:"subjectMode,omitempty" yaml:"subjectMode,omitempty"`
}

// PrivateCACertificateStatus defines the observed state of PrivateCACertificate.
type PrivateCACertificateStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
