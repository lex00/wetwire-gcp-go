package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ComputeInstanceTemplate represents a compute.cnrm.cloud.google.com ComputeInstanceTemplate resource.
type ComputeInstanceTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeInstanceTemplateSpec   `json:"spec,omitempty"`
	Status ComputeInstanceTemplateStatus `json:"status,omitempty"`
}

// ComputeInstanceTemplateSpec defines the desired state of ComputeInstanceTemplate.
type ComputeInstanceTemplateSpec struct {
	// Immutable. Controls for advanced machine-related behavior features.
	AdvancedMachineFeatures map[string]interface{} `json:"advancedMachineFeatures,omitempty" yaml:"advancedMachineFeatures,omitempty"`
	// Immutable. Whether to allow sending and receiving of packets with non-matching source or destination IPs. This defaults to false.
	CanIPForward bool `json:"canIpForward,omitempty" yaml:"canIpForward,omitempty"`
	// Immutable. The Confidential VM config being used by the instance. on_host_maintenance has to be set to TERMINATE or this will fail to create.
	ConfidentialInstanceConfig map[string]interface{} `json:"confidentialInstanceConfig,omitempty" yaml:"confidentialInstanceConfig,omitempty"`
	// Immutable. A brief description of this resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Immutable. Disks to attach to instances created from this template. This can be specified multiple times for multiple disks.
	Disk []map[string]interface{} `json:"disk,omitempty" yaml:"disk,omitempty"`
	// Immutable. Enable Virtual Displays on this instance. Note: allow_stopping_for_update must be set to true in order to update this field.
	EnableDisplay bool `json:"enableDisplay,omitempty" yaml:"enableDisplay,omitempty"`
	// Immutable. List of the type and count of accelerator cards attached to the instance.
	GuestAccelerator []map[string]interface{} `json:"guestAccelerator,omitempty" yaml:"guestAccelerator,omitempty"`
	// Immutable. A description of the instance.
	InstanceDescription string `json:"instanceDescription,omitempty" yaml:"instanceDescription,omitempty"`
	// Immutable. The machine type to create. To create a machine with a custom type (such as extended memory), format the value like custom-VCPUS-MEM_IN_MB like custom-6-20480 for 6 vCPU and 20GB of RAM.
	MachineType string `json:"machineType,omitempty" yaml:"machineType,omitempty"`
	Metadata []map[string]interface{} `json:"metadata,omitempty" yaml:"metadata,omitempty"`
	// Immutable. An alternative to using the startup-script metadata key, mostly to match the compute_instance resource. This replaces the startup-script metadata key on the created instance and thus the two mechanisms are not allowed to be used simultaneously.
	MetadataStartupScript string `json:"metadataStartupScript,omitempty" yaml:"metadataStartupScript,omitempty"`
	// Immutable. Specifies a minimum CPU platform. Applicable values are the friendly names of CPU platforms, such as Intel Haswell or Intel Skylake.
	MinCPUPlatform string `json:"minCpuPlatform,omitempty" yaml:"minCpuPlatform,omitempty"`
	// Immutable. Creates a unique name beginning with the specified prefix. Conflicts with name.
	NamePrefix string `json:"namePrefix,omitempty" yaml:"namePrefix,omitempty"`
	// Immutable. Networks to attach to instances created from this template. This can be specified multiple times for multiple networks.
	NetworkInterface []map[string]interface{} `json:"networkInterface,omitempty" yaml:"networkInterface,omitempty"`
	// Immutable. Configures network performance settings for the instance. If not specified, the instance will be created with its default network performance configuration.
	NetworkPerformanceConfig map[string]interface{} `json:"networkPerformanceConfig,omitempty" yaml:"networkPerformanceConfig,omitempty"`
	// Immutable. An instance template is a global resource that is not bound to a zone or a region. However, you can still specify some regional resources in an instance template, which restricts the template to the region where that resource resides. For example, a custom subnetwork resource is tied to a specific region. Defaults to the region of the Provider if no value is given.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`
	// Immutable. Specifies the reservations that this instance can consume from.
	ReservationAffinity map[string]interface{} `json:"reservationAffinity,omitempty" yaml:"reservationAffinity,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	ResourcePolicies []map[string]interface{} `json:"resourcePolicies,omitempty" yaml:"resourcePolicies,omitempty"`
	// Immutable. The scheduling strategy to use.
	Scheduling map[string]interface{} `json:"scheduling,omitempty" yaml:"scheduling,omitempty"`
	// Immutable. Service account to attach to the instance.
	ServiceAccount map[string]interface{} `json:"serviceAccount,omitempty" yaml:"serviceAccount,omitempty"`
	// Immutable. Enable Shielded VM on this instance. Shielded VM provides verifiable integrity to prevent against malware and rootkits. Defaults to disabled. Note: shielded_instance_config can only be used with boot images with shielded vm support.
	ShieldedInstanceConfig map[string]interface{} `json:"shieldedInstanceConfig,omitempty" yaml:"shieldedInstanceConfig,omitempty"`
	// Immutable. Tags to attach to the instance.
	Tags []string `json:"tags,omitempty" yaml:"tags,omitempty"`
}

// ComputeInstanceTemplateStatus defines the observed state of ComputeInstanceTemplate.
type ComputeInstanceTemplateStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
