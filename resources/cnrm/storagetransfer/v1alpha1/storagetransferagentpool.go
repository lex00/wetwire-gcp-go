package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// StorageTransferAgentPool represents a storagetransfer.cnrm.cloud.google.com StorageTransferAgentPool resource.
type StorageTransferAgentPool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   StorageTransferAgentPoolSpec   `json:"spec,omitempty"`
	Status StorageTransferAgentPoolStatus `json:"status,omitempty"`
}

// StorageTransferAgentPoolSpec defines the desired state of StorageTransferAgentPool.
type StorageTransferAgentPoolSpec struct {
	// Specifies the bandwidth limit details. If this field is unspecified, the default value is set as 'No Limit'.
	BandwidthLimit map[string]interface{} `json:"bandwidthLimit,omitempty" yaml:"bandwidthLimit,omitempty"`
	// Specifies the client-specified AgentPool description.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// StorageTransferAgentPoolStatus defines the observed state of StorageTransferAgentPool.
type StorageTransferAgentPoolStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
