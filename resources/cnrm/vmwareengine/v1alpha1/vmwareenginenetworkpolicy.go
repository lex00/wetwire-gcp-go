package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// VMwareEngineNetworkPolicy is the Schema for the VMwareEngineNetworkPolicy API
type VMwareEngineNetworkPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VMwareEngineNetworkPolicySpec   `json:"spec,omitempty"`
	Status VMwareEngineNetworkPolicyStatus `json:"status,omitempty"`
}

// VMwareEngineNetworkPolicySpec defines the desired state of VMwareEngineNetworkPolicy.
type VMwareEngineNetworkPolicySpec struct {
	// Optional. User-provided description for this network policy.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Required. IP address range in CIDR notation used to create internet access and external IP access. An RFC 1918 CIDR block, with a "/26" prefix, is required. The range cannot overlap with any prefixes either in the consumer VPC network or in use by the private clouds attached to that VPC network.
	EdgeServicesCIDR string `json:"edgeServicesCIDR,omitempty" yaml:"edgeServicesCIDR,omitempty"`
	// Network service that allows External IP addresses to be assigned to VMware workloads. This service can only be enabled when `internet_access` is also enabled.
	ExternalIP map[string]interface{} `json:"externalIP,omitempty" yaml:"externalIP,omitempty"`
	// Network service that allows VMware workloads to access the internet.
	InternetAccess map[string]interface{} `json:"internetAccess,omitempty" yaml:"internetAccess,omitempty"`
	// Immutable.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// The VMwareEngineNetworkPolicy name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Optional. The relative resource name of the VMware Engine network.
	VmwareEngineNetworkRef map[string]interface{} `json:"vmwareEngineNetworkRef,omitempty" yaml:"vmwareEngineNetworkRef,omitempty"`
}

// VMwareEngineNetworkPolicyStatus defines the observed state of VMwareEngineNetworkPolicy.
type VMwareEngineNetworkPolicyStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
