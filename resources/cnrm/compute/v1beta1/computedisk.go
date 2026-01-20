package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ComputeDisk represents a compute.cnrm.cloud.google.com ComputeDisk resource.
type ComputeDisk struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeDiskSpec   `json:"spec,omitempty"`
	Status ComputeDiskStatus `json:"status,omitempty"`
}

// ComputeDiskSpec defines the desired state of ComputeDisk.
type ComputeDiskSpec struct {
	// Immutable. A nested object resource.
	AsyncPrimaryDisk map[string]interface{} `json:"asyncPrimaryDisk,omitempty" yaml:"asyncPrimaryDisk,omitempty"`
	// Immutable. An optional description of this resource. Provide this property when
	// you create the resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Immutable. Encrypts the disk using a customer-supplied encryption key.
	// After you encrypt a disk with a customer-supplied key, you must
	// provide the same key if you use the disk later (e.g. to create a disk
	// snapshot or an image, or to attach the disk to a virtual machine).
	// Customer-supplied encryption keys do not protect access to metadata of
	// the disk.
	// If you do not provide an encryption key when creating the disk, then
	// the disk will be encrypted using an automatically generated key and
	// you do not need to provide a key to use the disk later.
	DiskEncryptionKey map[string]interface{} `json:"diskEncryptionKey,omitempty" yaml:"diskEncryptionKey,omitempty"`
	// Immutable. Whether this disk is using confidential compute mode.
	// Note: Only supported on hyperdisk skus, disk_encryption_key is required when setting to true.
	EnableConfidentialCompute bool `json:"enableConfidentialCompute,omitempty" yaml:"enableConfidentialCompute,omitempty"`
	// Immutable. A list of features to enable on the guest operating system.
	// Applicable only for bootable disks.
	GuestOsFeatures []map[string]interface{} `json:"guestOsFeatures,omitempty" yaml:"guestOsFeatures,omitempty"`
	// The image from which to initialize this disk.
	ImageRef map[string]interface{} `json:"imageRef,omitempty" yaml:"imageRef,omitempty"`
	// DEPRECATED. `interface` is deprecated. This field is no longer used and can be safely removed from your configurations; disk interfaces are automatically determined on attachment. Immutable. Specifies the disk interface to use for attaching this disk, which is either SCSI or NVME. The default is SCSI.
	Interface string `json:"interface,omitempty" yaml:"interface,omitempty"`
	// Immutable. Any applicable license URI.
	Licenses []string `json:"licenses,omitempty" yaml:"licenses,omitempty"`
	// Location represents the geographical location of the ComputeDisk. Specify a region name or a zone name. Reference: GCP definition of regions/zones (https://cloud.google.com/compute/docs/regions-zones/)
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Immutable. Indicates whether or not the disk can be read/write attached to more than one instance.
	MultiWriter bool `json:"multiWriter,omitempty" yaml:"multiWriter,omitempty"`
	// Immutable. Physical block size of the persistent disk, in bytes. If not present
	// in a request, a default value is used. Currently supported sizes
	// are 4096 and 16384, other sizes may be added in the future.
	// If an unsupported value is requested, the error message will list
	// the supported values for the caller's project.
	PhysicalBlockSizeBytes int32 `json:"physicalBlockSizeBytes,omitempty" yaml:"physicalBlockSizeBytes,omitempty"`
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Indicates how many IOPS must be provisioned for the disk.
	// Note: Updating currently is only supported by hyperdisk skus without the need to delete and recreate the disk, hyperdisk
	// allows for an update of IOPS every 4 hours. To update your hyperdisk more frequently, you'll need to manually delete and recreate it.
	ProvisionedIops int32 `json:"provisionedIops,omitempty" yaml:"provisionedIops,omitempty"`
	// Indicates how much Throughput must be provisioned for the disk.
	// Note: Updating currently is only supported by hyperdisk skus without the need to delete and recreate the disk, hyperdisk
	// allows for an update of Throughput every 4 hours. To update your hyperdisk more frequently, you'll need to manually delete and recreate it.
	ProvisionedThroughput int32 `json:"provisionedThroughput,omitempty" yaml:"provisionedThroughput,omitempty"`
	// Immutable. URLs of the zones where the disk should be replicated to.
	ReplicaZones []string `json:"replicaZones,omitempty" yaml:"replicaZones,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	ResourcePolicies []map[string]interface{} `json:"resourcePolicies,omitempty" yaml:"resourcePolicies,omitempty"`
	// Size of the persistent disk, specified in GB. You can specify this
	// field when creating a persistent disk using the 'image' or
	// 'snapshot' parameter, or specify it alone to create an empty
	// persistent disk.
	// If you specify this field along with 'image' or 'snapshot',
	// the value must not be less than the size of the image
	// or the size of the snapshot.
	// Upsizing the disk is mutable, but downsizing the disk
	// requires re-creating the resource.
	Size int32 `json:"size,omitempty" yaml:"size,omitempty"`
	// The source snapshot used to create this disk.
	SnapshotRef map[string]interface{} `json:"snapshotRef,omitempty" yaml:"snapshotRef,omitempty"`
	// The source disk used to create this disk.
	SourceDiskRef map[string]interface{} `json:"sourceDiskRef,omitempty" yaml:"sourceDiskRef,omitempty"`
	// Immutable. The customer-supplied encryption key of the source image. Required if
	// the source image is protected by a customer-supplied encryption key.
	SourceImageEncryptionKey map[string]interface{} `json:"sourceImageEncryptionKey,omitempty" yaml:"sourceImageEncryptionKey,omitempty"`
	// Immutable. The customer-supplied encryption key of the source snapshot. Required
	// if the source snapshot is protected by a customer-supplied encryption
	// key.
	SourceSnapshotEncryptionKey map[string]interface{} `json:"sourceSnapshotEncryptionKey,omitempty" yaml:"sourceSnapshotEncryptionKey,omitempty"`
	// Immutable. URL of the disk type resource describing which disk type to use to
	// create the disk. Provide this when creating the disk.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}

// ComputeDiskStatus defines the observed state of ComputeDisk.
type ComputeDiskStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
