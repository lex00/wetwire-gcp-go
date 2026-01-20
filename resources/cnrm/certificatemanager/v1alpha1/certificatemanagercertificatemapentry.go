package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CertificateManagerCertificateMapEntry represents a certificatemanager.cnrm.cloud.google.com CertificateManagerCertificateMapEntry resource.
type CertificateManagerCertificateMapEntry struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CertificateManagerCertificateMapEntrySpec   `json:"spec,omitempty"`
	Status CertificateManagerCertificateMapEntryStatus `json:"status,omitempty"`
}

// CertificateManagerCertificateMapEntrySpec defines the desired state of CertificateManagerCertificateMapEntry.
type CertificateManagerCertificateMapEntrySpec struct {
	CertificatesRefs []map[string]interface{} `json:"certificatesRefs,omitempty" yaml:"certificatesRefs,omitempty"`
	// A human-readable description of the resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Immutable. A Hostname (FQDN, e.g. example.com) or a wildcard hostname expression (*.example.com)
	// for a set of hostnames with common suffix. Used as Server Name Indication (SNI) for
	// selecting a proper certificate.
	Hostname string `json:"hostname,omitempty" yaml:"hostname,omitempty"`
	// A map entry that is inputted into the certificate map.
	MapRef map[string]interface{} `json:"mapRef,omitempty" yaml:"mapRef,omitempty"`
	// Immutable. A predefined matcher for particular cases, other than SNI selection.
	Matcher string `json:"matcher,omitempty" yaml:"matcher,omitempty"`
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// CertificateManagerCertificateMapEntryStatus defines the observed state of CertificateManagerCertificateMapEntry.
type CertificateManagerCertificateMapEntryStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
