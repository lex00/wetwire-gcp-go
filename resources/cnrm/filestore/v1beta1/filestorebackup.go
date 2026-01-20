package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// FilestoreBackup represents a filestore.cnrm.cloud.google.com FilestoreBackup resource.
type FilestoreBackup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FilestoreBackupSpec   `json:"spec,omitempty"`
	Status FilestoreBackupStatus `json:"status,omitempty"`
}

// FilestoreBackupSpec defines the desired state of FilestoreBackup.
type FilestoreBackupSpec struct {
	// A description of the backup with 2048 characters or less. Requests with longer descriptions will be rejected.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Immutable. The location for the resource
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Immutable. The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Immutable. Name of the file share in the source Cloud Filestore instance that the backup is created from.
	SourceFileShare string `json:"sourceFileShare,omitempty" yaml:"sourceFileShare,omitempty"`
	// Immutable.
	SourceInstanceRef map[string]interface{} `json:"sourceInstanceRef,omitempty" yaml:"sourceInstanceRef,omitempty"`
}

// FilestoreBackupStatus defines the observed state of FilestoreBackup.
type FilestoreBackupStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
