package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NetworkSecurityClientTLSPolicy represents a networksecurity.cnrm.cloud.google.com NetworkSecurityClientTLSPolicy resource.
type NetworkSecurityClientTLSPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NetworkSecurityClientTLSPolicySpec   `json:"spec,omitempty"`
	Status NetworkSecurityClientTLSPolicyStatus `json:"status,omitempty"`
}

// NetworkSecurityClientTLSPolicySpec defines the desired state of NetworkSecurityClientTLSPolicy.
type NetworkSecurityClientTLSPolicySpec struct {
	// Optional. Defines a mechanism to provision client identity (public and private keys) for peer to peer authentication. The presence of this dictates mTLS.
	ClientCertificate map[string]interface{} `json:"clientCertificate,omitempty" yaml:"clientCertificate,omitempty"`
	// Optional. Free-text description of the resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Immutable. The location for the resource
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Immutable. The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Required. Defines the mechanism to obtain the Certificate Authority certificate to validate the server certificate.
	ServerValidationCa []map[string]interface{} `json:"serverValidationCa,omitempty" yaml:"serverValidationCa,omitempty"`
	// Optional. Server Name Indication string to present to the server during TLS handshake. E.g: "secure.example.com".
	Sni string `json:"sni,omitempty" yaml:"sni,omitempty"`
}

// NetworkSecurityClientTLSPolicyStatus defines the observed state of NetworkSecurityClientTLSPolicy.
type NetworkSecurityClientTLSPolicyStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
