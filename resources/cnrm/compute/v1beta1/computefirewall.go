package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ComputeFirewall represents a compute.cnrm.cloud.google.com ComputeFirewall resource.
type ComputeFirewall struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeFirewallSpec   `json:"spec,omitempty"`
	Status ComputeFirewallStatus `json:"status,omitempty"`
}

// ComputeFirewallSpec defines the desired state of ComputeFirewall.
type ComputeFirewallSpec struct {
	// The list of ALLOW rules specified by this firewall. Each rule
	// specifies a protocol and port-range tuple that describes a permitted
	// connection.
	Allow []map[string]interface{} `json:"allow,omitempty" yaml:"allow,omitempty"`
	// The list of DENY rules specified by this firewall. Each rule specifies
	// a protocol and port-range tuple that describes a denied connection.
	Deny []map[string]interface{} `json:"deny,omitempty" yaml:"deny,omitempty"`
	// An optional description of this resource. Provide this property when
	// you create the resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// If destination ranges are specified, the firewall will apply only to
	// traffic that has destination IP address in these ranges. These ranges
	// must be expressed in CIDR format. IPv4 or IPv6 ranges are supported.
	DestinationRanges []string `json:"destinationRanges,omitempty" yaml:"destinationRanges,omitempty"`
	// Immutable. Direction of traffic to which this firewall applies; default is
	// INGRESS. Note: For INGRESS traffic, one of 'source_ranges',
	// 'source_tags' or 'source_service_accounts' is required. Possible values: ["INGRESS", "EGRESS"].
	Direction string `json:"direction,omitempty" yaml:"direction,omitempty"`
	// Denotes whether the firewall rule is disabled, i.e not applied to the
	// network it is associated with. When set to true, the firewall rule is
	// not enforced and the network behaves as if it did not exist. If this
	// is unspecified, the firewall rule will be enabled.
	Disabled bool `json:"disabled,omitempty" yaml:"disabled,omitempty"`
	// DEPRECATED. Deprecated in favor of log_config. This field denotes whether to enable logging for a particular firewall rule. If logging is enabled, logs will be exported to Stackdriver.
	EnableLogging bool `json:"enableLogging,omitempty" yaml:"enableLogging,omitempty"`
	// This field denotes the logging options for a particular firewall rule.
	// If defined, logging is enabled, and logs will be exported to Cloud Logging.
	LogConfig map[string]interface{} `json:"logConfig,omitempty" yaml:"logConfig,omitempty"`
	// The network to attach this firewall to.
	NetworkRef map[string]interface{} `json:"networkRef,omitempty" yaml:"networkRef,omitempty"`
	// Priority for this rule. This is an integer between 0 and 65535, both
	// inclusive. When not specified, the value assumed is 1000. Relative
	// priorities determine precedence of conflicting rules. Lower value of
	// priority implies higher precedence (eg, a rule with priority 0 has
	// higher precedence than a rule with priority 1). DENY rules take
	// precedence over ALLOW rules having equal priority.
	Priority int32 `json:"priority,omitempty" yaml:"priority,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// If source ranges are specified, the firewall will apply only to
	// traffic that has source IP address in these ranges. These ranges must
	// be expressed in CIDR format. One or both of sourceRanges and
	// sourceTags may be set. If both properties are set, the firewall will
	// apply to traffic that has source IP address within sourceRanges OR the
	// source IP that belongs to a tag listed in the sourceTags property. The
	// connection does not need to match both properties for the firewall to
	// apply. IPv4 or IPv6 ranges are supported. For INGRESS traffic, one of
	// 'source_ranges', 'source_tags' or 'source_service_accounts' is required.
	SourceRanges []string `json:"sourceRanges,omitempty" yaml:"sourceRanges,omitempty"`
	SourceServiceAccounts []map[string]interface{} `json:"sourceServiceAccounts,omitempty" yaml:"sourceServiceAccounts,omitempty"`
	// If source tags are specified, the firewall will apply only to traffic
	// with source IP that belongs to a tag listed in source tags. Source
	// tags cannot be used to control traffic to an instance's external IP
	// address. Because tags are associated with an instance, not an IP
	// address. One or both of sourceRanges and sourceTags may be set. If
	// both properties are set, the firewall will apply to traffic that has
	// source IP address within sourceRanges OR the source IP that belongs to
	// a tag listed in the sourceTags property. The connection does not need
	// to match both properties for the firewall to apply. For INGRESS traffic,
	// one of 'source_ranges', 'source_tags' or 'source_service_accounts' is required.
	SourceTags []string `json:"sourceTags,omitempty" yaml:"sourceTags,omitempty"`
	TargetServiceAccounts []map[string]interface{} `json:"targetServiceAccounts,omitempty" yaml:"targetServiceAccounts,omitempty"`
	// A list of instance tags indicating sets of instances located in the
	// network that may make network connections as specified in allowed[].
	// If no targetTags are specified, the firewall rule applies to all
	// instances on the specified network.
	TargetTags []string `json:"targetTags,omitempty" yaml:"targetTags,omitempty"`
}

// ComputeFirewallStatus defines the observed state of ComputeFirewall.
type ComputeFirewallStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
