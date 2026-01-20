package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ComputeInstance represents a compute.cnrm.cloud.google.com ComputeInstance resource.
type ComputeInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeInstanceSpec   `json:"spec,omitempty"`
	Status ComputeInstanceStatus `json:"status,omitempty"`
}

// ComputeInstanceSpec defines the desired state of ComputeInstance.
type ComputeInstanceSpec struct {
	// Controls for advanced machine-related behavior features.
	AdvancedMachineFeatures map[string]interface{} `json:"advancedMachineFeatures,omitempty" yaml:"advancedMachineFeatures,omitempty"`
	// List of disks attached to the instance.
	AttachedDisk []map[string]interface{} `json:"attachedDisk,omitempty" yaml:"attachedDisk,omitempty"`
	// Immutable. The boot disk for the instance.
	BootDisk map[string]interface{} `json:"bootDisk,omitempty" yaml:"bootDisk,omitempty"`
	// Whether sending and receiving of packets with non-matching source or destination IPs is allowed.
	CanIPForward bool `json:"canIpForward,omitempty" yaml:"canIpForward,omitempty"`
	// Immutable. The Confidential VM config being used by the instance.  on_host_maintenance has to be set to TERMINATE or this will fail to create.
	ConfidentialInstanceConfig map[string]interface{} `json:"confidentialInstanceConfig,omitempty" yaml:"confidentialInstanceConfig,omitempty"`
	// Whether deletion protection is enabled on this instance.
	DeletionProtection bool `json:"deletionProtection,omitempty" yaml:"deletionProtection,omitempty"`
	// Immutable. A brief description of the resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Desired status of the instance. Either "RUNNING" or "TERMINATED".
	DesiredStatus string `json:"desiredStatus,omitempty" yaml:"desiredStatus,omitempty"`
	// Whether the instance has virtual displays enabled.
	EnableDisplay bool `json:"enableDisplay,omitempty" yaml:"enableDisplay,omitempty"`
	// Immutable. List of the type and count of accelerator cards attached to the instance.
	GuestAccelerator []map[string]interface{} `json:"guestAccelerator,omitempty" yaml:"guestAccelerator,omitempty"`
	// Immutable. A custom hostname for the instance. Must be a fully qualified DNS name and RFC-1035-valid. Valid format is a series of labels 1-63 characters long matching the regular expression [a-z]([-a-z0-9]*[a-z0-9]), concatenated with periods. The entire hostname must not exceed 253 characters. Changing this forces a new resource to be created.
	Hostname string `json:"hostname,omitempty" yaml:"hostname,omitempty"`
	InstanceTemplateRef map[string]interface{} `json:"instanceTemplateRef,omitempty" yaml:"instanceTemplateRef,omitempty"`
	// The machine type to create.
	MachineType string `json:"machineType,omitempty" yaml:"machineType,omitempty"`
	Metadata []map[string]interface{} `json:"metadata,omitempty" yaml:"metadata,omitempty"`
	// Immutable. Metadata startup scripts made available within the instance.
	MetadataStartupScript string `json:"metadataStartupScript,omitempty" yaml:"metadataStartupScript,omitempty"`
	// The minimum CPU platform specified for the VM instance.
	MinCPUPlatform string `json:"minCpuPlatform,omitempty" yaml:"minCpuPlatform,omitempty"`
	// Immutable. The networks attached to the instance.
	NetworkInterface []map[string]interface{} `json:"networkInterface,omitempty" yaml:"networkInterface,omitempty"`
	// Immutable. Configures network performance settings for the instance. If not specified, the instance will be created with its default network performance configuration.
	NetworkPerformanceConfig map[string]interface{} `json:"networkPerformanceConfig,omitempty" yaml:"networkPerformanceConfig,omitempty"`
	// Immutable. Stores additional params passed with the request, but not persisted as part of resource payload.
	Params map[string]interface{} `json:"params,omitempty" yaml:"params,omitempty"`
	// Immutable. Specifies the reservations that this instance can consume from.
	ReservationAffinity map[string]interface{} `json:"reservationAffinity,omitempty" yaml:"reservationAffinity,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	ResourcePolicies []map[string]interface{} `json:"resourcePolicies,omitempty" yaml:"resourcePolicies,omitempty"`
	// The scheduling strategy being used by the instance.
	Scheduling map[string]interface{} `json:"scheduling,omitempty" yaml:"scheduling,omitempty"`
	// Immutable. The scratch disks attached to the instance.
	ScratchDisk []map[string]interface{} `json:"scratchDisk,omitempty" yaml:"scratchDisk,omitempty"`
	// The service account to attach to the instance.
	ServiceAccount map[string]interface{} `json:"serviceAccount,omitempty" yaml:"serviceAccount,omitempty"`
	// The shielded vm config being used by the instance.
	ShieldedInstanceConfig map[string]interface{} `json:"shieldedInstanceConfig,omitempty" yaml:"shieldedInstanceConfig,omitempty"`
	// The list of tags attached to the instance.
	Tags []string `json:"tags,omitempty" yaml:"tags,omitempty"`
	// Immutable. The zone of the instance. If self_link is provided, this value is ignored. If neither self_link nor zone are provided, the provider zone is used.
	Zone string `json:"zone,omitempty" yaml:"zone,omitempty"`
}

// ComputeInstanceStatus defines the observed state of ComputeInstance.
type ComputeInstanceStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
