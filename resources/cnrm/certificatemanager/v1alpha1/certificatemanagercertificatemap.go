package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CertificateManagerCertificateMap represents a certificatemanager.cnrm.cloud.google.com CertificateManagerCertificateMap resource.
type CertificateManagerCertificateMap struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CertificateManagerCertificateMapSpec   `json:"spec,omitempty"`
	Status CertificateManagerCertificateMapStatus `json:"status,omitempty"`
}

// CertificateManagerCertificateMapSpec defines the desired state of CertificateManagerCertificateMap.
type CertificateManagerCertificateMapSpec struct {
	// A human-readable description of the resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// CertificateManagerCertificateMapStatus defines the observed state of CertificateManagerCertificateMap.
type CertificateManagerCertificateMapStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
