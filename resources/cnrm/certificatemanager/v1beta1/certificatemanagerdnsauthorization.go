package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CertificateManagerDNSAuthorization is the Schema for the CertificateManagerDNSAuthorization API
type CertificateManagerDNSAuthorization struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CertificateManagerDNSAuthorizationSpec   `json:"spec,omitempty"`
	Status CertificateManagerDNSAuthorizationStatus `json:"status,omitempty"`
}

// CertificateManagerDNSAuthorizationSpec defines the desired state of CertificateManagerDNSAuthorization.
type CertificateManagerDNSAuthorizationSpec struct {
	// A human-readable description of the resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Immutable. A domain which is being authorized. A DnsAuthorization resource covers a single domain and its wildcard, e.g. authorization for "example.com" can be used to issue certificates for "example.com" and "*.example.com".
	Domain string `json:"domain,omitempty" yaml:"domain,omitempty"`
	// Immutable. Optional. Location represents the geographical location of the DnsAuthorization. If not specified, "global" is used.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// CertificateManagerDNSAuthorizationStatus defines the observed state of CertificateManagerDNSAuthorization.
type CertificateManagerDNSAuthorizationStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
