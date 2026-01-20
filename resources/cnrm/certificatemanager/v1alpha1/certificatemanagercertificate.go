package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CertificateManagerCertificate represents a certificatemanager.cnrm.cloud.google.com CertificateManagerCertificate resource.
type CertificateManagerCertificate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CertificateManagerCertificateSpec   `json:"spec,omitempty"`
	Status CertificateManagerCertificateStatus `json:"status,omitempty"`
}

// CertificateManagerCertificateSpec defines the desired state of CertificateManagerCertificate.
type CertificateManagerCertificateSpec struct {
	// A human-readable description of the resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Immutable. The Certificate Manager location. If not specified, "global" is used.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Immutable. Configuration and state of a Managed Certificate.
	// Certificate Manager provisions and renews Managed Certificates
	// automatically, for as long as it's authorized to do so.
	Managed map[string]interface{} `json:"managed,omitempty" yaml:"managed,omitempty"`
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Immutable. The scope of the certificate.
	// DEFAULT: Certificates with default scope are served from core Google data centers.
	// If unsure, choose this option.
	// EDGE_CACHE: Certificates with scope EDGE_CACHE are special-purposed certificates,
	// served from non-core Google data centers.
	// ALL_REGIONS: Certificates with ALL_REGIONS scope are served from all GCP regions (You can only use ALL_REGIONS with global certs).
	// see https://cloud.google.com/compute/docs/regions-zones.
	Scope string `json:"scope,omitempty" yaml:"scope,omitempty"`
	// Immutable. Certificate data for a SelfManaged Certificate.
	// SelfManaged Certificates are uploaded by the user. Updating such
	// certificates before they expire remains the user's responsibility.
	SelfManaged map[string]interface{} `json:"selfManaged,omitempty" yaml:"selfManaged,omitempty"`
}

// CertificateManagerCertificateStatus defines the observed state of CertificateManagerCertificate.
type CertificateManagerCertificateStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
