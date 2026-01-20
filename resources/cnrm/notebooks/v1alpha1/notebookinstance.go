package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NotebookInstance is the Schema for the NotebookInstance API
type NotebookInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NotebookInstanceSpec   `json:"spec,omitempty"`
	Status NotebookInstanceStatus `json:"status,omitempty"`
}

// NotebookInstanceSpec defines the desired state of NotebookInstance.
type NotebookInstanceSpec struct {
	// The hardware accelerator used on this instance. If you use accelerators, make sure that your configuration has [enough vCPUs and memory to support the `machine_type` you have selected](https://cloud.google.com/compute/docs/gpus/#gpus-list).
	AcceleratorConfig map[string]interface{} `json:"acceleratorConfig,omitempty" yaml:"acceleratorConfig,omitempty"`
	// Input only. The size of the boot disk in GB attached to this instance, up to a maximum of 64000 GB (64 TB). The minimum recommended value is 100 GB. If not specified, this defaults to 100.
	BootDiskSizeGB int64 `json:"bootDiskSizeGB,omitempty" yaml:"bootDiskSizeGB,omitempty"`
	// Input only. The type of the boot disk attached to this instance, defaults to standard persistent disk (`PD_STANDARD`).
	BootDiskType string `json:"bootDiskType,omitempty" yaml:"bootDiskType,omitempty"`
	// Optional. Flag to enable ip forwarding or not, default false/off. https://cloud.google.com/vpc/docs/using-routes#canipforward
	CanIPForward bool `json:"canIPForward,omitempty" yaml:"canIPForward,omitempty"`
	// Use a container image to start the notebook instance.
	ContainerImage map[string]interface{} `json:"containerImage,omitempty" yaml:"containerImage,omitempty"`
	// Specify a custom Cloud Storage path where the GPU driver is stored. If not specified, we'll automatically choose from official GPU drivers.
	CustomGpuDriverPath string `json:"customGpuDriverPath,omitempty" yaml:"customGpuDriverPath,omitempty"`
	// Input only. The size of the data disk in GB attached to this instance, up to a maximum of 64000 GB (64 TB). You can choose the size of the data disk based on how big your notebooks and data are. If not specified, this defaults to 100.
	DataDiskSizeGB int64 `json:"dataDiskSizeGB,omitempty" yaml:"dataDiskSizeGB,omitempty"`
	// Input only. The type of the data disk attached to this instance, defaults to standard persistent disk (`PD_STANDARD`).
	DataDiskType string `json:"dataDiskType,omitempty" yaml:"dataDiskType,omitempty"`
	// Input only. Disk encryption method used on the boot and data disks, defaults to GMEK.
	DiskEncryption string `json:"diskEncryption,omitempty" yaml:"diskEncryption,omitempty"`
	// Whether the end user authorizes Google Cloud to install GPU driver on this instance. If this field is empty or set to false, the GPU driver won't be installed. Only applicable to instances with GPUs.
	InstallGpuDriver bool `json:"installGpuDriver,omitempty" yaml:"installGpuDriver,omitempty"`
	// Input only. The owner of this instance after creation. Format: `alias@example.com`
	// Currently supports one owner only. If not specified, all of the service
	// account users of your VM instance's service account can use
	// the instance.
	InstanceOwners []string `json:"instanceOwners,omitempty" yaml:"instanceOwners,omitempty"`
	// Input only. The KMS key used to encrypt the disks, only applicable if disk_encryption is CMEK. Learn more about [using your own encryption keys](/kms/docs/quickstart).
	KmsKeyRef map[string]interface{} `json:"kmsKeyRef,omitempty" yaml:"kmsKeyRef,omitempty"`
	// Labels to apply to this instance. These can be later modified by the setLabels method.
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	// Required. The [Compute Engine machine type](https://cloud.google.com/compute/docs/machine-types) of this instance.
	MachineType string `json:"machineType,omitempty" yaml:"machineType,omitempty"`
	// Custom metadata to apply to this instance.
	Metadata map[string]string `json:"metadata,omitempty" yaml:"metadata,omitempty"`
	// The name of the VPC that this instance is in.
	NetworkRef map[string]interface{} `json:"networkRef,omitempty" yaml:"networkRef,omitempty"`
	// Optional. The type of vNIC to be used on this interface. This may be gVNIC or VirtioNet.
	NicType string `json:"nicType,omitempty" yaml:"nicType,omitempty"`
	// If true, the notebook instance will not register with the proxy.
	NoProxyAccess bool `json:"noProxyAccess,omitempty" yaml:"noProxyAccess,omitempty"`
	// If true, no public IP will be assigned to this instance.
	NoPublicIP bool `json:"noPublicIP,omitempty" yaml:"noPublicIP,omitempty"`
	// Input only. If true, the data disk will not be auto deleted when deleting the instance.
	NoRemoveDataDisk bool `json:"noRemoveDataDisk,omitempty" yaml:"noRemoveDataDisk,omitempty"`
	// Path to a Bash script that automatically runs after a notebook instance fully boots up. The path must be a URL or Cloud Storage path (`gs://path-to-file/file-name`).
	PostStartupScript string `json:"postStartupScript,omitempty" yaml:"postStartupScript,omitempty"`
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Optional. The optional reservation affinity. Setting this field will apply the specified [Zonal Compute Reservation](https://cloud.google.com/compute/docs/instances/reserving-zonal-resources) to this notebook instance.
	ReservationAffinity map[string]interface{} `json:"reservationAffinity,omitempty" yaml:"reservationAffinity,omitempty"`
	// The NotebookInstance name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// The service account on this instance, giving access to other Google
	// Cloud services.
	// You can use any service account within the same project, but you
	// must have the service account user permission to use the instance.
	// If not specified, the [Compute Engine default service
	// account](https://cloud.google.com/compute/docs/access/service-accounts#default_service_account)
	// is used.
	ServiceAccountRef map[string]interface{} `json:"serviceAccountRef,omitempty" yaml:"serviceAccountRef,omitempty"`
	// Optional. The URIs of service account scopes to be included in
	// Compute Engine instances.
	// If not specified, the following
	// [scopes](https://cloud.google.com/compute/docs/access/service-accounts#accesscopesiam)
	// are defined:
	// - https://www.googleapis.com/auth/cloud-platform
	// - https://www.googleapis.com/auth/userinfo.email
	// If not using default scopes, you need at least:
	// https://www.googleapis.com/auth/compute
	ServiceAccountScopes []string `json:"serviceAccountScopes,omitempty" yaml:"serviceAccountScopes,omitempty"`
	// Optional. Shielded VM configuration. [Images using supported Shielded VM features](https://cloud.google.com/compute/docs/instances/modifying-shielded-vm).
	ShieldedInstanceConfig map[string]interface{} `json:"shieldedInstanceConfig,omitempty" yaml:"shieldedInstanceConfig,omitempty"`
	// The name of the subnet that this instance is in.
	SubnetRef map[string]interface{} `json:"subnetRef,omitempty" yaml:"subnetRef,omitempty"`
	// Optional. The Compute Engine tags to add to runtime (see [Tagging instances](https://cloud.google.com/compute/docs/label-or-tag-resources#tags)).
	Tags []string `json:"tags,omitempty" yaml:"tags,omitempty"`
	// The upgrade history of this instance.
	UpgradeHistory []map[string]interface{} `json:"upgradeHistory,omitempty" yaml:"upgradeHistory,omitempty"`
	// Use a Compute Engine VM image to start the notebook instance.
	VmImage map[string]interface{} `json:"vmImage,omitempty" yaml:"vmImage,omitempty"`
	// Immutable. The location where the notebook instance should reside.
	Zone string `json:"zone,omitempty" yaml:"zone,omitempty"`
}

// NotebookInstanceStatus defines the observed state of NotebookInstance.
type NotebookInstanceStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
