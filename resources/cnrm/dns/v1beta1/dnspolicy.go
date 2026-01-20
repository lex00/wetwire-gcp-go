package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DNSPolicy represents a dns.cnrm.cloud.google.com DNSPolicy resource.
type DNSPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DNSPolicySpec   `json:"spec,omitempty"`
	Status DNSPolicyStatus `json:"status,omitempty"`
}

// DNSPolicySpec defines the desired state of DNSPolicy.
type DNSPolicySpec struct {
	// Sets an alternative name server for the associated networks.
	// When specified, all DNS queries are forwarded to a name server that you choose.
	// Names such as .internal are not available when an alternative name server is specified.
	AlternativeNameServerConfig map[string]interface{} `json:"alternativeNameServerConfig,omitempty" yaml:"alternativeNameServerConfig,omitempty"`
	// A textual description field. Defaults to 'Managed by Config Connector'.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Allows networks bound to this policy to receive DNS queries sent
	// by VMs or applications over VPN connections. When enabled, a
	// virtual IP address will be allocated from each of the sub-networks
	// that are bound to this policy.
	EnableInboundForwarding bool `json:"enableInboundForwarding,omitempty" yaml:"enableInboundForwarding,omitempty"`
	// Controls whether logging is enabled for the networks bound to this policy.
	// Defaults to no logging if not set.
	EnableLogging bool `json:"enableLogging,omitempty" yaml:"enableLogging,omitempty"`
	// List of network names specifying networks to which this policy is applied.
	Networks []map[string]interface{} `json:"networks,omitempty" yaml:"networks,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// DNSPolicyStatus defines the observed state of DNSPolicy.
type DNSPolicyStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
