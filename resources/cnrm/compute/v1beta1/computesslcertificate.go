package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ComputeSSLCertificate represents a compute.cnrm.cloud.google.com ComputeSSLCertificate resource.
type ComputeSSLCertificate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeSSLCertificateSpec   `json:"spec,omitempty"`
	Status ComputeSSLCertificateStatus `json:"status,omitempty"`
}

// ComputeSSLCertificateSpec defines the desired state of ComputeSSLCertificate.
type ComputeSSLCertificateSpec struct {
	// Immutable. The certificate in PEM format.
	// The certificate chain must be no greater than 5 certs long.
	// The chain must include at least one intermediate cert.
	Certificate map[string]interface{} `json:"certificate,omitempty" yaml:"certificate,omitempty"`
	// Immutable. An optional description of this resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Location represents the geographical location of the ComputeSSLCertificate. Specify a region name or "global" for global resources. Reference: GCP definition of regions/zones (https://cloud.google.com/compute/docs/regions-zones/)
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Immutable. The write-only private key in PEM format.
	PrivateKey map[string]interface{} `json:"privateKey,omitempty" yaml:"privateKey,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// ComputeSSLCertificateStatus defines the observed state of ComputeSSLCertificate.
type ComputeSSLCertificateStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
