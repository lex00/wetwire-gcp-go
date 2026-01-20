package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ComputeRouterNAT represents a compute.cnrm.cloud.google.com ComputeRouterNAT resource.
type ComputeRouterNAT struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeRouterNATSpec   `json:"spec,omitempty"`
	Status ComputeRouterNATStatus `json:"status,omitempty"`
}

// ComputeRouterNATSpec defines the desired state of ComputeRouterNAT.
type ComputeRouterNATSpec struct {
	DrainNatIps []map[string]interface{} `json:"drainNatIps,omitempty" yaml:"drainNatIps,omitempty"`
	// Enable Dynamic Port Allocation.
	// If minPortsPerVm is set, minPortsPerVm must be set to a power of two greater than or equal to 32.
	// If minPortsPerVm is not set, a minimum of 32 ports will be allocated to a VM from this NAT config.
	// If maxPortsPerVm is set, maxPortsPerVm must be set to a power of two greater than minPortsPerVm.
	// If maxPortsPerVm is not set, a maximum of 65536 ports will be allocated to a VM from this NAT config.
	// Mutually exclusive with enableEndpointIndependentMapping.
	EnableDynamicPortAllocation bool `json:"enableDynamicPortAllocation,omitempty" yaml:"enableDynamicPortAllocation,omitempty"`
	// Specifies if endpoint independent mapping is enabled. This is enabled by default. For more information
	// see the [official documentation](https://cloud.google.com/nat/docs/overview#specs-rfcs).
	EnableEndpointIndependentMapping bool `json:"enableEndpointIndependentMapping,omitempty" yaml:"enableEndpointIndependentMapping,omitempty"`
	// Timeout (in seconds) for ICMP connections. Defaults to 30s if not set.
	IcmpIdleTimeoutSec int32 `json:"icmpIdleTimeoutSec,omitempty" yaml:"icmpIdleTimeoutSec,omitempty"`
	// Configuration for logging on NAT.
	LogConfig map[string]interface{} `json:"logConfig,omitempty" yaml:"logConfig,omitempty"`
	// Maximum number of ports allocated to a VM from this NAT.
	// This field can only be set when enableDynamicPortAllocation is enabled.
	MaxPortsPerVm int32 `json:"maxPortsPerVm,omitempty" yaml:"maxPortsPerVm,omitempty"`
	// Minimum number of ports allocated to a VM from this NAT.
	MinPortsPerVm int32 `json:"minPortsPerVm,omitempty" yaml:"minPortsPerVm,omitempty"`
	// How external IPs should be allocated for this NAT. Valid values are
	// 'AUTO_ONLY' for only allowing NAT IPs allocated by Google Cloud
	// Platform, or 'MANUAL_ONLY' for only user-allocated NAT IP addresses. Possible values: ["MANUAL_ONLY", "AUTO_ONLY"].
	NatIPAllocateOption string `json:"natIpAllocateOption,omitempty" yaml:"natIpAllocateOption,omitempty"`
	NatIps []map[string]interface{} `json:"natIps,omitempty" yaml:"natIps,omitempty"`
	// Immutable. Region where the router and NAT reside.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// The Cloud Router in which this NAT will be configured.
	RouterRef map[string]interface{} `json:"routerRef,omitempty" yaml:"routerRef,omitempty"`
	// A list of rules associated with this NAT.
	Rules []map[string]interface{} `json:"rules,omitempty" yaml:"rules,omitempty"`
	// How NAT should be configured per Subnetwork.
	// If 'ALL_SUBNETWORKS_ALL_IP_RANGES', all of the
	// IP ranges in every Subnetwork are allowed to Nat.
	// If 'ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES', all of the primary IP
	// ranges in every Subnetwork are allowed to Nat.
	// 'LIST_OF_SUBNETWORKS': A list of Subnetworks are allowed to Nat
	// (specified in the field subnetwork below). Note that if this field
	// contains ALL_SUBNETWORKS_ALL_IP_RANGES or
	// ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES, then there should not be any
	// other RouterNat section in any Router for this network in this region. Possible values: ["ALL_SUBNETWORKS_ALL_IP_RANGES", "ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES", "LIST_OF_SUBNETWORKS"].
	SourceSubnetworkIPRangesToNat string `json:"sourceSubnetworkIpRangesToNat,omitempty" yaml:"sourceSubnetworkIpRangesToNat,omitempty"`
	// One or more subnetwork NAT configurations. Only used if
	// 'source_subnetwork_ip_ranges_to_nat' is set to 'LIST_OF_SUBNETWORKS'.
	Subnetwork []map[string]interface{} `json:"subnetwork,omitempty" yaml:"subnetwork,omitempty"`
	// Timeout (in seconds) for TCP established connections.
	// Defaults to 1200s if not set.
	TCPEstablishedIdleTimeoutSec int32 `json:"tcpEstablishedIdleTimeoutSec,omitempty" yaml:"tcpEstablishedIdleTimeoutSec,omitempty"`
	// Timeout (in seconds) for TCP connections that are in TIME_WAIT state.
	// Defaults to 120s if not set.
	TCPTimeWaitTimeoutSec int32 `json:"tcpTimeWaitTimeoutSec,omitempty" yaml:"tcpTimeWaitTimeoutSec,omitempty"`
	// Timeout (in seconds) for TCP transitory connections.
	// Defaults to 30s if not set.
	TCPTransitoryIdleTimeoutSec int32 `json:"tcpTransitoryIdleTimeoutSec,omitempty" yaml:"tcpTransitoryIdleTimeoutSec,omitempty"`
	// Timeout (in seconds) for UDP connections. Defaults to 30s if not set.
	UDPIdleTimeoutSec int32 `json:"udpIdleTimeoutSec,omitempty" yaml:"udpIdleTimeoutSec,omitempty"`
}

// ComputeRouterNATStatus defines the observed state of ComputeRouterNAT.
type ComputeRouterNATStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
