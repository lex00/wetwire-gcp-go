package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// TPUNode represents a tpu.cnrm.cloud.google.com TPUNode resource.
type TPUNode struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TPUNodeSpec   `json:"spec,omitempty"`
	Status TPUNodeStatus `json:"status,omitempty"`
}

// TPUNodeSpec defines the desired state of TPUNode.
type TPUNodeSpec struct {
	// Immutable. The type of hardware accelerators associated with this node.
	AcceleratorType string `json:"acceleratorType,omitempty" yaml:"acceleratorType,omitempty"`
	// Immutable. The CIDR block that the TPU node will use when selecting an IP
	// address. This CIDR block must be a /29 block; the Compute Engine
	// networks API forbids a smaller block, and using a larger block would
	// be wasteful (a node can only consume one IP address).
	// Errors will occur if the CIDR block has already been used for a
	// currently existing TPU node, the CIDR block conflicts with any
	// subnetworks in the user's provided network, or the provided network
	// is peered with another network that is using that CIDR block.
	CidrBlock string `json:"cidrBlock,omitempty" yaml:"cidrBlock,omitempty"`
	// Immutable. The user-supplied description of the TPU. Maximum of 512 characters.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Immutable. The name of a network to peer the TPU node to. It must be a
	// preexisting Compute Engine network inside of the project on which
	// this API has been activated. If none is provided, "default" will be
	// used.
	Network string `json:"network,omitempty" yaml:"network,omitempty"`
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Immutable. Sets the scheduling options for this TPU instance.
	SchedulingConfig map[string]interface{} `json:"schedulingConfig,omitempty" yaml:"schedulingConfig,omitempty"`
	// The version of Tensorflow running in the Node.
	TensorflowVersion string `json:"tensorflowVersion,omitempty" yaml:"tensorflowVersion,omitempty"`
	// Immutable. Whether the VPC peering for the node is set up through Service Networking API.
	// The VPC Peering should be set up before provisioning the node. If this field is set,
	// cidr_block field should not be specified. If the network that you want to peer the
	// TPU Node to is a Shared VPC network, the node must be created with this this field enabled.
	UseServiceNetworking bool `json:"useServiceNetworking,omitempty" yaml:"useServiceNetworking,omitempty"`
	// Immutable. The GCP location for the TPU. If it is not provided, the provider zone is used.
	Zone string `json:"zone,omitempty" yaml:"zone,omitempty"`
}

// TPUNodeStatus defines the observed state of TPUNode.
type TPUNodeStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
