package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NetworkManagementConnectivityTest is the Schema for the NetworkManagementConnectivityTest API
type NetworkManagementConnectivityTest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NetworkManagementConnectivityTestSpec   `json:"spec,omitempty"`
	Status NetworkManagementConnectivityTestStatus `json:"status,omitempty"`
}

// NetworkManagementConnectivityTestSpec defines the desired state of NetworkManagementConnectivityTest.
type NetworkManagementConnectivityTestSpec struct {
	// Whether the test should skip firewall checking. If not provided, we assume false.
	BypassFirewallChecks bool `json:"bypassFirewallChecks,omitempty" yaml:"bypassFirewallChecks,omitempty"`
	// The user-supplied description of the Connectivity Test. Maximum of 512 characters.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Required. Destination specification of the Connectivity Test.
	// You can use a combination of destination IP address, Compute Engine
	// VM instance, or VPC network to uniquely identify the destination
	// location.
	// Even if the destination IP address is not unique, the source IP
	// location is unique. Usually, the analysis can infer the destination
	// endpoint from route information.
	// If the destination you specify is a VM instance and the instance has
	// multiple network interfaces, then you must also specify either
	// a destination IP address  or VPC network to identify the destination
	// interface.
	// A reachability analysis proceeds even if the destination location is
	// ambiguous. However, the result can include endpoints that you don't
	// intend to test.
	Destination map[string]interface{} `json:"destination,omitempty" yaml:"destination,omitempty"`
	// Resource labels to represent user-provided metadata.
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	// Immutable. The name of the location where the RuntimeTemplate will be created.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Immutable. The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// IP Protocol of the test. When not provided, "TCP" is assumed.
	Protocol string `json:"protocol,omitempty" yaml:"protocol,omitempty"`
	// Other projects that may be relevant for reachability analysis. This is applicable to scenarios where a test can cross project boundaries.
	RelatedProjects []map[string]interface{} `json:"relatedProjects,omitempty" yaml:"relatedProjects,omitempty"`
	// The NetworkManagementConnectivityTest name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Whether run analysis for the return path from destination to source. Default value is false.
	RoundTrip bool `json:"roundTrip,omitempty" yaml:"roundTrip,omitempty"`
	// Required. Source specification of the Connectivity Test.
	// You can use a combination of source IP address, virtual machine
	// (VM) instance, or Compute Engine network to uniquely identify
	// the source location.
	// Examples:
	// If the source IP address is an internal IP address within a Google Cloud
	// Virtual Private Cloud (VPC) network, then you must also specify the VPC
	// network. Otherwise, specify the VM instance, which already contains its
	// internal IP address and VPC network information.
	// If the source of the test is within an on-premises network, then you must
	// provide the destination VPC network.
	// If the source endpoint is a Compute Engine VM instance with multiple
	// network interfaces, the instance itself is not sufficient to identify the
	// endpoint. So, you must also specify the source IP address or VPC network.
	// A reachability analysis proceeds even if the source location is
	// ambiguous. However, the test result may include endpoints that you don't
	// intend to test.
	Source map[string]interface{} `json:"source,omitempty" yaml:"source,omitempty"`
}

// NetworkManagementConnectivityTestStatus defines the observed state of NetworkManagementConnectivityTest.
type NetworkManagementConnectivityTestStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
