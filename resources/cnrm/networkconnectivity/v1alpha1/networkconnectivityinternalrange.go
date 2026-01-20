package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NetworkConnectivityInternalRange is the Schema for the NetworkConnectivityInternalRange API
type NetworkConnectivityInternalRange struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NetworkConnectivityInternalRangeSpec   `json:"spec,omitempty"`
	Status NetworkConnectivityInternalRangeStatus `json:"status,omitempty"`
}

// NetworkConnectivityInternalRangeSpec defines the desired state of NetworkConnectivityInternalRange.
type NetworkConnectivityInternalRangeSpec struct {
	// A description of this resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// The IP range that this internal range defines. NOTE: IPv6 ranges are limited to usage=EXTERNAL_TO_VPC and peering=FOR_SELF. NOTE: For IPv6 Ranges this field is compulsory, i.e. the address range must be specified explicitly.
	IPCIDRRange string `json:"ipCIDRRange,omitempty" yaml:"ipCIDRRange,omitempty"`
	// User-defined labels.
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	// Required. The location of the application.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Optional. Must be present if usage is set to FOR_MIGRATION. This field is for internal use.
	Migration map[string]interface{} `json:"migration,omitempty" yaml:"migration,omitempty"`
	// The network in which to reserve the internal range. The network cannot be deleted if there are any reserved internal ranges referring to it. Legacy networks are not supported. For example: https://www.googleapis.com/compute/v1/projects/{project}/locations/global/networks/{network} projects/{project}/locations/global/networks/{network} {network}
	NetworkRef map[string]interface{} `json:"networkRef,omitempty" yaml:"networkRef,omitempty"`
	// Optional. Types of resources that are allowed to overlap with the current internal range.
	Overlaps []string `json:"overlaps,omitempty" yaml:"overlaps,omitempty"`
	// The type of peering set for this internal range.
	Peering string `json:"peering,omitempty" yaml:"peering,omitempty"`
	// An alternate to ip_cidr_range. Can be set when trying to create an IPv4 reservation that automatically finds a free range of the given size. If both ip_cidr_range and prefix_length are set, there is an error if the range sizes do not match. Can also be used during updates to change the range size. NOTE: For IPv6 this field only works if ip_cidr_range is set as well, and both fields must match. In other words, with IPv6 this field only works as a redundant parameter.
	PrefixLength int32 `json:"prefixLength,omitempty" yaml:"prefixLength,omitempty"`
	// Required. The host project of the application.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// The NetworkConnectivityInternalRange name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Optional. Can be set to narrow down or pick a different address space while searching for a free range. If not set, defaults to the "10.0.0.0/8" address space. This can be used to search in other rfc-1918 address spaces like "172.16.0.0/12" and "192.168.0.0/16" or non-rfc-1918 address spaces used in the VPC.
	TargetCIDRRange []string `json:"targetCIDRRange,omitempty" yaml:"targetCIDRRange,omitempty"`
	// The type of usage set for this InternalRange.
	Usage string `json:"usage,omitempty" yaml:"usage,omitempty"`
}

// NetworkConnectivityInternalRangeStatus defines the observed state of NetworkConnectivityInternalRange.
type NetworkConnectivityInternalRangeStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
