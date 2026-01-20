package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ComputeImage represents a compute.cnrm.cloud.google.com ComputeImage resource.
type ComputeImage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeImageSpec   `json:"spec,omitempty"`
	Status ComputeImageStatus `json:"status,omitempty"`
}

// ComputeImageSpec defines the desired state of ComputeImage.
type ComputeImageSpec struct {
	// Immutable. An optional description of this resource. Provide this property when
	// you create the resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// The source disk to create this image based on.
	// You must provide either this property or the
	// rawDisk.source property but not both to create an image.
	DiskRef map[string]interface{} `json:"diskRef,omitempty" yaml:"diskRef,omitempty"`
	// Immutable. Size of the image when restored onto a persistent disk (in GB).
	DiskSizeGb int32 `json:"diskSizeGb,omitempty" yaml:"diskSizeGb,omitempty"`
	// Immutable. The name of the image family to which this image belongs. You can
	// create disks by specifying an image family instead of a specific
	// image name. The image family always returns its latest image that is
	// not deprecated. The name of the image family must comply with
	// RFC1035.
	Family string `json:"family,omitempty" yaml:"family,omitempty"`
	// Immutable. A list of features to enable on the guest operating system.
	// Applicable only for bootable images.
	GuestOsFeatures []map[string]interface{} `json:"guestOsFeatures,omitempty" yaml:"guestOsFeatures,omitempty"`
	// Immutable. Encrypts the image using a customer-supplied encryption key.
	// After you encrypt an image with a customer-supplied key, you must
	// provide the same key if you use the image later (e.g. to create a
	// disk from the image).
	ImageEncryptionKey map[string]interface{} `json:"imageEncryptionKey,omitempty" yaml:"imageEncryptionKey,omitempty"`
	// Immutable. Any applicable license URI.
	Licenses []string `json:"licenses,omitempty" yaml:"licenses,omitempty"`
	// Immutable. The parameters of the raw disk image.
	RawDisk map[string]interface{} `json:"rawDisk,omitempty" yaml:"rawDisk,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// The source image used to create this image.
	SourceImageRef map[string]interface{} `json:"sourceImageRef,omitempty" yaml:"sourceImageRef,omitempty"`
	// The source snapshot used to create this image.
	SourceSnapshotRef map[string]interface{} `json:"sourceSnapshotRef,omitempty" yaml:"sourceSnapshotRef,omitempty"`
	// Immutable. Cloud Storage bucket storage location of the image
	// (regional or multi-regional).
	// Reference link: https://cloud.google.com/compute/docs/reference/rest/v1/images.
	StorageLocations []string `json:"storageLocations,omitempty" yaml:"storageLocations,omitempty"`
}

// ComputeImageStatus defines the observed state of ComputeImage.
type ComputeImageStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
