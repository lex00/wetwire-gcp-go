package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// SQLSSLCert represents a sql.cnrm.cloud.google.com SQLSSLCert resource.
type SQLSSLCert struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SQLSSLCertSpec   `json:"spec,omitempty"`
	Status SQLSSLCertStatus `json:"status,omitempty"`
}

// SQLSSLCertSpec defines the desired state of SQLSSLCert.
type SQLSSLCertSpec struct {
	// Immutable. The common name to be used in the certificate to identify the client. Constrained to [a-zA-Z.-_ ]+. Changing this forces a new resource to be created.
	CommonName string `json:"commonName,omitempty" yaml:"commonName,omitempty"`
	// The Cloud SQL instance.
	InstanceRef map[string]interface{} `json:"instanceRef,omitempty" yaml:"instanceRef,omitempty"`
	// Immutable. Optional. The service-generated sha1Fingerprint of the resource. Used for acquisition only. Leave unset to create a new resource.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// SQLSSLCertStatus defines the observed state of SQLSSLCert.
type SQLSSLCertStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
