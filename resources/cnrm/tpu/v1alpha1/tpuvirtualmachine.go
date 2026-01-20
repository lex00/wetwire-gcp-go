package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// TPUVirtualMachine is the Schema for the TPUVirtualMachine API
type TPUVirtualMachine struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TPUVirtualMachineSpec   `json:"spec,omitempty"`
	Status TPUVirtualMachineStatus `json:"status,omitempty"`
}

// TPUVirtualMachineSpec defines the desired state of TPUVirtualMachine.
type TPUVirtualMachineSpec struct {
	// The AccleratorConfig for the TPU Node.
	AcceleratorConfig map[string]interface{} `json:"acceleratorConfig,omitempty" yaml:"acceleratorConfig,omitempty"`
	// Optional. The type of hardware accelerators associated with this node.
	AcceleratorType string `json:"acceleratorType,omitempty" yaml:"acceleratorType,omitempty"`
	// The CIDR block that the TPU node will use when selecting an IP address. This CIDR block must be a /29 block; the Compute Engine networks API forbids a smaller block, and using a larger block would be wasteful (a node can only consume one IP address). Errors will occur if the CIDR block has already been used for a currently existing TPU node, the CIDR block conflicts with any subnetworks in the user's provided network, or the provided network is peered with another network that is using that CIDR block.
	CidrBlock string `json:"cidrBlock,omitempty" yaml:"cidrBlock,omitempty"`
	// The additional data disks for the Node.
	DataDisks []map[string]interface{} `json:"dataDisks,omitempty" yaml:"dataDisks,omitempty"`
	// The user-supplied description of the TPU. Maximum of 512 characters.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Immutable. The location where the TPU virtual machine should reside.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Custom metadata to apply to the TPU Node. Can set startup-script and shutdown-script
	Metadata map[string]string `json:"metadata,omitempty" yaml:"metadata,omitempty"`
	// Network configurations for the TPU node. network_config and network_configs are mutually exclusive, you can only specify one of them. If both are specified, an error will be returned.
	NetworkConfig map[string]interface{} `json:"networkConfig,omitempty" yaml:"networkConfig,omitempty"`
	// Optional. Repeated network configurations for the TPU node. This field is used to specify multiple networks configs for the TPU node. network_config and network_configs are mutually exclusive, you can only specify one of them. If both are specified, an error will be returned.
	NetworkConfigs []map[string]interface{} `json:"networkConfigs,omitempty" yaml:"networkConfigs,omitempty"`
	// The project that the TPU virtual machine belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// The TPUVirtualMachine name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Required. The runtime version running in the Node.
	RuntimeVersion string `json:"runtimeVersion,omitempty" yaml:"runtimeVersion,omitempty"`
	// The scheduling options for this node.
	SchedulingConfig map[string]interface{} `json:"schedulingConfig,omitempty" yaml:"schedulingConfig,omitempty"`
	// The Google Cloud Platform Service Account to be used by the TPU node VMs. If None is specified, the default compute service account will be used.
	ServiceAccount map[string]interface{} `json:"serviceAccount,omitempty" yaml:"serviceAccount,omitempty"`
	// Shielded Instance options.
	ShieldedInstanceConfig map[string]interface{} `json:"shieldedInstanceConfig,omitempty" yaml:"shieldedInstanceConfig,omitempty"`
	// Tags to apply to the TPU Node. Tags are used to identify valid sources or targets for network firewalls.
	Tags []string `json:"tags,omitempty" yaml:"tags,omitempty"`
}

// TPUVirtualMachineStatus defines the observed state of TPUVirtualMachine.
type TPUVirtualMachineStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
