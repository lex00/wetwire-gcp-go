package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ComputeSnapshot represents a compute.cnrm.cloud.google.com ComputeSnapshot resource.
type ComputeSnapshot struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeSnapshotSpec   `json:"spec,omitempty"`
	Status ComputeSnapshotStatus `json:"status,omitempty"`
}

// ComputeSnapshotSpec defines the desired state of ComputeSnapshot.
type ComputeSnapshotSpec struct {
	// Immutable. Creates the new snapshot in the snapshot chain labeled with the
	// specified name. The chain name must be 1-63 characters long and
	// comply with RFC1035. This is an uncommon option only for advanced
	// service owners who needs to create separate snapshot chains, for
	// example, for chargeback tracking.  When you describe your snapshot
	// resource, this field is visible only if it has a non-empty value.
	ChainName string `json:"chainName,omitempty" yaml:"chainName,omitempty"`
	// Immutable. An optional description of this resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Immutable. Encrypts the snapshot using a customer-supplied encryption key.
	// After you encrypt a snapshot using a customer-supplied key, you must
	// provide the same key if you use the snapshot later. For example, you
	// must provide the encryption key when you create a disk from the
	// encrypted snapshot in a future request.
	// Customer-supplied encryption keys do not protect access to metadata of
	// the snapshot.
	// If you do not provide an encryption key when creating the snapshot,
	// then the snapshot will be encrypted using an automatically generated
	// key and you do not need to provide a key to use the snapshot later.
	SnapshotEncryptionKey map[string]interface{} `json:"snapshotEncryptionKey,omitempty" yaml:"snapshotEncryptionKey,omitempty"`
	// Immutable. The customer-supplied encryption key of the source snapshot. Required
	// if the source snapshot is protected by a customer-supplied encryption
	// key.
	SourceDiskEncryptionKey map[string]interface{} `json:"sourceDiskEncryptionKey,omitempty" yaml:"sourceDiskEncryptionKey,omitempty"`
	// A reference to the disk used to create this snapshot.
	SourceDiskRef map[string]interface{} `json:"sourceDiskRef,omitempty" yaml:"sourceDiskRef,omitempty"`
	// Immutable. Cloud Storage bucket storage location of the snapshot (regional or multi-regional).
	StorageLocations []string `json:"storageLocations,omitempty" yaml:"storageLocations,omitempty"`
	// Immutable. A reference to the zone where the disk is hosted.
	Zone string `json:"zone,omitempty" yaml:"zone,omitempty"`
}

// ComputeSnapshotStatus defines the observed state of ComputeSnapshot.
type ComputeSnapshotStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
