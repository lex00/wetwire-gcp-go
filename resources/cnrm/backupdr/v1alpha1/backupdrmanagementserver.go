package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// BackupDRManagementServer is the Schema for the BackupDRManagementServer API
type BackupDRManagementServer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BackupDRManagementServerSpec   `json:"spec,omitempty"`
	Status BackupDRManagementServerStatus `json:"status,omitempty"`
}

// BackupDRManagementServerSpec defines the desired state of BackupDRManagementServer.
type BackupDRManagementServerSpec struct {
	// Optional. The description of the ManagementServer instance (2048 characters or less).
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Optional. Resource labels to represent user provided metadata. Labels currently defined: 1. migrate_from_go=<false|true> If set to true, the MS is created in migration ready mode.
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	// Immutable.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Optional. VPC networks to which the ManagementServer instance is connected. For this version, only a single network is supported. This field is optional if MS is created without PSA
	Networks []map[string]interface{} `json:"networks,omitempty" yaml:"networks,omitempty"`
	// The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// The BackupDRManagementServer name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Optional. The type of the ManagementServer resource.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}

// BackupDRManagementServerStatus defines the observed state of BackupDRManagementServer.
type BackupDRManagementServerStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
