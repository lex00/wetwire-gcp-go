package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// VPCAccessConnector represents a vpcaccess.cnrm.cloud.google.com VPCAccessConnector resource.
type VPCAccessConnector struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VPCAccessConnectorSpec   `json:"spec,omitempty"`
	Status VPCAccessConnectorStatus `json:"status,omitempty"`
}

// VPCAccessConnectorSpec defines the desired state of VPCAccessConnector.
type VPCAccessConnectorSpec struct {
	// Immutable. The range of internal addresses that follows RFC 4632 notation. Example: '10.132.0.0/28'.
	IPCidrRange string `json:"ipCidrRange,omitempty" yaml:"ipCidrRange,omitempty"`
	// Location represents the geographical location of the VPCAccessConnector. Specify a region name. Reference: GCP definition of regions/zones (https://cloud.google.com/compute/docs/regions-zones/)
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Immutable. Machine type of VM Instance underlying connector. Default is e2-micro.
	MachineType string `json:"machineType,omitempty" yaml:"machineType,omitempty"`
	// Immutable. Maximum value of instances in autoscaling group underlying the connector.
	MaxInstances int32 `json:"maxInstances,omitempty" yaml:"maxInstances,omitempty"`
	// Immutable. Maximum throughput of the connector in Mbps, must be greater than 'min_throughput'. Default is 300.
	MaxThroughput int32 `json:"maxThroughput,omitempty" yaml:"maxThroughput,omitempty"`
	// Immutable. Minimum value of instances in autoscaling group underlying the connector.
	MinInstances int32 `json:"minInstances,omitempty" yaml:"minInstances,omitempty"`
	// Immutable. Minimum throughput of the connector in Mbps. Default and min is 200.
	MinThroughput int32 `json:"minThroughput,omitempty" yaml:"minThroughput,omitempty"`
	// Immutable. Name or self_link of the VPC network. Required if 'ip_cidr_range' is set.
	NetworkRef map[string]interface{} `json:"networkRef,omitempty" yaml:"networkRef,omitempty"`
	// Immutable. The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Immutable. The subnet in which to house the connector.
	Subnet map[string]interface{} `json:"subnet,omitempty" yaml:"subnet,omitempty"`
}

// VPCAccessConnectorStatus defines the observed state of VPCAccessConnector.
type VPCAccessConnectorStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
