package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ComputePacketMirroring represents a compute.cnrm.cloud.google.com ComputePacketMirroring resource.
type ComputePacketMirroring struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputePacketMirroringSpec   `json:"spec,omitempty"`
	Status ComputePacketMirroringStatus `json:"status,omitempty"`
}

// ComputePacketMirroringSpec defines the desired state of ComputePacketMirroring.
type ComputePacketMirroringSpec struct {
	// The Forwarding Rule resource of type `loadBalancingScheme=INTERNAL` that will be used as collector for mirrored traffic. The specified forwarding rule must have `isMirroringCollector` set to true.
	CollectorIlb map[string]interface{} `json:"collectorIlb,omitempty" yaml:"collectorIlb,omitempty"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Indicates whether or not this packet mirroring takes effect. If set to FALSE, this packet mirroring policy will not be enforced on the network. The default is TRUE.
	Enable string `json:"enable,omitempty" yaml:"enable,omitempty"`
	// Filter for mirrored traffic. If unspecified, all traffic is mirrored.
	Filter map[string]interface{} `json:"filter,omitempty" yaml:"filter,omitempty"`
	// Immutable. The location for the resource
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// PacketMirroring mirroredResourceInfos. MirroredResourceInfo specifies a set of mirrored VM instances, subnetworks and/or tags for which traffic from/to all VM instances will be mirrored.
	MirroredResources map[string]interface{} `json:"mirroredResources,omitempty" yaml:"mirroredResources,omitempty"`
	// Immutable. Specifies the mirrored VPC network. Only packets in this network will be mirrored. All mirrored VMs should have a NIC in the given network. All mirrored subnetworks should belong to the given network.
	Network map[string]interface{} `json:"network,omitempty" yaml:"network,omitempty"`
	// The priority of applying this configuration. Priority is used to break ties in cases where there is more than one matching rule. In the case of two rules that apply for a given Instance, the one with the lowest-numbered priority value wins. Default value is 1000. Valid range is 0 through 65535.
	Priority int64 `json:"priority,omitempty" yaml:"priority,omitempty"`
	// Immutable. The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// ComputePacketMirroringStatus defines the observed state of ComputePacketMirroring.
type ComputePacketMirroringStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
