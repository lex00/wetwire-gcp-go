package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// FilestoreSnapshot represents a filestore.cnrm.cloud.google.com FilestoreSnapshot resource.
type FilestoreSnapshot struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FilestoreSnapshotSpec   `json:"spec,omitempty"`
	Status FilestoreSnapshotStatus `json:"status,omitempty"`
}

// FilestoreSnapshotSpec defines the desired state of FilestoreSnapshot.
type FilestoreSnapshotSpec struct {
	// A description of the snapshot with 2048 characters or less. Requests with longer descriptions will be rejected.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Immutable. The resource name of the filestore instance.
	Instance string `json:"instance,omitempty" yaml:"instance,omitempty"`
	// Immutable. The name of the location of the instance. This can be a region for ENTERPRISE tier instances.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// FilestoreSnapshotStatus defines the observed state of FilestoreSnapshot.
type FilestoreSnapshotStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
