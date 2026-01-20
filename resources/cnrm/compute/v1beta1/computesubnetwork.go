package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ComputeSubnetwork represents a compute.cnrm.cloud.google.com ComputeSubnetwork resource.
type ComputeSubnetwork struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeSubnetworkSpec   `json:"spec,omitempty"`
	Status ComputeSubnetworkStatus `json:"status,omitempty"`
}

// ComputeSubnetworkSpec defines the desired state of ComputeSubnetwork.
type ComputeSubnetworkSpec struct {
	// Immutable. An optional description of this resource. Provide this property when
	// you create the resource. This field can be set only at resource
	// creation time.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// The range of internal addresses that are owned by this subnetwork.
	// Provide this property when you create the subnetwork. For example,
	// 10.0.0.0/8 or 192.168.0.0/16. Ranges must be unique and
	// non-overlapping within a network. Only IPv4 is supported.
	IPCidrRange string `json:"ipCidrRange,omitempty" yaml:"ipCidrRange,omitempty"`
	// The access type of IPv6 address this subnet holds. It's immutable and can only be specified during creation
	// or the first time the subnet is updated into IPV4_IPV6 dual stack. If the ipv6_type is EXTERNAL then this subnet
	// cannot enable direct path. Possible values: ["EXTERNAL", "INTERNAL"].
	Ipv6AccessType string `json:"ipv6AccessType,omitempty" yaml:"ipv6AccessType,omitempty"`
	// This field denotes the VPC flow logging options for this subnetwork. If
	// logging is enabled, logs are exported to Cloud Logging. Flow logging
	// isn't supported if the subnet 'purpose' field is set to subnetwork is
	// 'REGIONAL_MANAGED_PROXY' or 'GLOBAL_MANAGED_PROXY'.
	LogConfig map[string]interface{} `json:"logConfig,omitempty" yaml:"logConfig,omitempty"`
	// The network this subnet belongs to. Only networks that are in the
	// distributed mode can have subnetworks.
	NetworkRef map[string]interface{} `json:"networkRef,omitempty" yaml:"networkRef,omitempty"`
	// When enabled, VMs in this subnetwork without external IP addresses can
	// access Google APIs and services by using Private Google Access.
	PrivateIPGoogleAccess bool `json:"privateIpGoogleAccess,omitempty" yaml:"privateIpGoogleAccess,omitempty"`
	// The private IPv6 google access type for the VMs in this subnet.
	PrivateIpv6GoogleAccess string `json:"privateIpv6GoogleAccess,omitempty" yaml:"privateIpv6GoogleAccess,omitempty"`
	// Immutable. The purpose of the resource. This field can be either 'PRIVATE_RFC_1918', 'REGIONAL_MANAGED_PROXY', 'GLOBAL_MANAGED_PROXY', or 'PRIVATE_SERVICE_CONNECT'.
	// A subnet with purpose set to 'REGIONAL_MANAGED_PROXY' is a user-created subnetwork that is reserved for regional Envoy-based load balancers.
	// A subnetwork in a given region with purpose set to 'GLOBAL_MANAGED_PROXY' is a proxy-only subnet and is shared between all the cross-regional Envoy-based load balancers.
	// A subnetwork with purpose set to 'PRIVATE_SERVICE_CONNECT' reserves the subnet for hosting a Private Service Connect published service.
	// Note that 'REGIONAL_MANAGED_PROXY' is the preferred setting for all regional Envoy load balancers.
	// If unspecified, the purpose defaults to 'PRIVATE_RFC_1918'.
	Purpose string `json:"purpose,omitempty" yaml:"purpose,omitempty"`
	// Immutable. The GCP region for this subnetwork.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// The role of subnetwork.
	// Currently, this field is only used when 'purpose' is 'REGIONAL_MANAGED_PROXY'.
	// The value can be set to 'ACTIVE' or 'BACKUP'.
	// An 'ACTIVE' subnetwork is one that is currently being used for Envoy-based load balancers in a region.
	// A 'BACKUP' subnetwork is one that is ready to be promoted to 'ACTIVE' or is currently draining. Possible values: ["ACTIVE", "BACKUP"].
	Role string `json:"role,omitempty" yaml:"role,omitempty"`
	SecondaryIPRange []map[string]interface{} `json:"secondaryIpRange,omitempty" yaml:"secondaryIpRange,omitempty"`
	// The stack type for this subnet to identify whether the IPv6 feature is enabled or not.
	// If not specified IPV4_ONLY will be used. Possible values: ["IPV4_ONLY", "IPV4_IPV6"].
	StackType string `json:"stackType,omitempty" yaml:"stackType,omitempty"`
}

// ComputeSubnetworkStatus defines the observed state of ComputeSubnetwork.
type ComputeSubnetworkStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
