package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// VMwareEngineExternalAccessRule is the Schema for the VMwareEngineExternalAccessRule API
type VMwareEngineExternalAccessRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VMwareEngineExternalAccessRuleSpec   `json:"spec,omitempty"`
	Status VMwareEngineExternalAccessRuleStatus `json:"status,omitempty"`
}

// VMwareEngineExternalAccessRuleSpec defines the desired state of VMwareEngineExternalAccessRule.
type VMwareEngineExternalAccessRuleSpec struct {
	// The action that the external access rule performs.
	Action string `json:"action,omitempty" yaml:"action,omitempty"`
	// User-provided description for this external access rule.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// If destination ranges are specified, the external access rule applies only to the traffic that has a destination IP address in these ranges. The specified IP addresses must have reserved external IP addresses in the scope of the parent network policy. To match all external IP addresses in the scope of the parent network policy, specify `0.0.0.0/0`. To match a specific external IP address, specify it using the `IpRange.external_address` property.
	DestinationIPRanges []map[string]interface{} `json:"destinationIPRanges,omitempty" yaml:"destinationIPRanges,omitempty"`
	// A list of destination ports to which the external access rule applies. This field is only applicable for the UDP or TCP protocol. Each entry must be either an integer or a range. For example: `["22"]`, `["80","443"]`, or `["12345-12349"]`. To match all destination ports, specify `["0-65535"]`.
	DestinationPorts []string `json:"destinationPorts,omitempty" yaml:"destinationPorts,omitempty"`
	// The IP protocol to which the external access rule applies. This value can be one of the following three protocol strings (not case-sensitive): `tcp`, `udp`, or `icmp`.
	IPProtocol string `json:"ipProtocol,omitempty" yaml:"ipProtocol,omitempty"`
	// Required. The resource name of the network policy to create a new external access firewall rule in.
	NetworkPolicyRef map[string]interface{} `json:"networkPolicyRef,omitempty" yaml:"networkPolicyRef,omitempty"`
	// External access rule priority, which determines the external access rule to use when multiple rules apply. If multiple rules have the same priority, their ordering is non-deterministic. If specific ordering is required, assign unique priorities to enforce such ordering. The external access rule priority is an integer from 100 to 4096, both inclusive. Lower integers indicate higher precedence. For example, a rule with priority `100` has higher precedence than a rule with priority `101`.
	Priority int32 `json:"priority,omitempty" yaml:"priority,omitempty"`
	// The VMwareEngineExternalAccessRule name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// If source ranges are specified, the external access rule applies only to traffic that has a source IP address in these ranges. These ranges can either be expressed in the CIDR format or as an IP address. As only inbound rules are supported, `ExternalAddress` resources cannot be the source IP addresses of an external access rule. To match all source addresses, specify `0.0.0.0/0`.
	SourceIPRanges []map[string]interface{} `json:"sourceIPRanges,omitempty" yaml:"sourceIPRanges,omitempty"`
	// A list of source ports to which the external access rule applies. This field is only applicable for the UDP or TCP protocol. Each entry must be either an integer or a range. For example: `["22"]`, `["80","443"]`, or `["12345-12349"]`. To match all source ports, specify `["0-65535"]`.
	SourcePorts []string `json:"sourcePorts,omitempty" yaml:"sourcePorts,omitempty"`
}

// VMwareEngineExternalAccessRuleStatus defines the observed state of VMwareEngineExternalAccessRule.
type VMwareEngineExternalAccessRuleStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
