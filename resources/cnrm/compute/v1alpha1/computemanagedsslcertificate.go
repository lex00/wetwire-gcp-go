package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ComputeManagedSSLCertificate represents a compute.cnrm.cloud.google.com ComputeManagedSSLCertificate resource.
type ComputeManagedSSLCertificate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeManagedSSLCertificateSpec   `json:"spec,omitempty"`
	Status ComputeManagedSSLCertificateStatus `json:"status,omitempty"`
}

// ComputeManagedSSLCertificateSpec defines the desired state of ComputeManagedSSLCertificate.
type ComputeManagedSSLCertificateSpec struct {
	// Immutable. An optional description of this resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Immutable. Properties relevant to a managed certificate.  These will be used if the
	// certificate is managed (as indicated by a value of 'MANAGED' in 'type').
	Managed map[string]interface{} `json:"managed,omitempty" yaml:"managed,omitempty"`
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Immutable. Enum field whose value is always 'MANAGED' - used to signal to the API
	// which type this is. Default value: "MANAGED" Possible values: ["MANAGED"].
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}

// ComputeManagedSSLCertificateStatus defines the observed state of ComputeManagedSSLCertificate.
type ComputeManagedSSLCertificateStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
