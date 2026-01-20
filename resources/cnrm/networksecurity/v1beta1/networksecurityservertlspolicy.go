package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NetworkSecurityServerTLSPolicy represents a networksecurity.cnrm.cloud.google.com NetworkSecurityServerTLSPolicy resource.
type NetworkSecurityServerTLSPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NetworkSecurityServerTLSPolicySpec   `json:"spec,omitempty"`
	Status NetworkSecurityServerTLSPolicyStatus `json:"status,omitempty"`
}

// NetworkSecurityServerTLSPolicySpec defines the desired state of NetworkSecurityServerTLSPolicy.
type NetworkSecurityServerTLSPolicySpec struct {
	// Optional. Determines if server allows plaintext connections. If set to true, server allows plain text connections. By default, it is set to false. This setting is not exclusive of other encryption modes. For example, if allow_open and mtls_policy are set, server allows both plain text and mTLS connections. See documentation of other encryption modes to confirm compatibility.
	AllowOpen bool `json:"allowOpen,omitempty" yaml:"allowOpen,omitempty"`
	// Optional. Free-text description of the resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Immutable. The location for the resource
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Optional. Defines a mechanism to provision peer validation certificates for peer to peer authentication (Mutual TLS - mTLS). If not specified, client certificate will not be requested. The connection is treated as TLS and not mTLS. If allow_open and mtls_policy are set, server allows both plain text and mTLS connections.
	MtlsPolicy map[string]interface{} `json:"mtlsPolicy,omitempty" yaml:"mtlsPolicy,omitempty"`
	// Immutable. The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Optional. Defines a mechanism to provision server identity (public and private keys). Cannot be combined with allow_open as a permissive mode that allows both plain text and TLS is not supported.
	ServerCertificate map[string]interface{} `json:"serverCertificate,omitempty" yaml:"serverCertificate,omitempty"`
}

// NetworkSecurityServerTLSPolicyStatus defines the observed state of NetworkSecurityServerTLSPolicy.
type NetworkSecurityServerTLSPolicyStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
