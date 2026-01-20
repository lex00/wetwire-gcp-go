package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EdgeContainerNodePool represents a edgecontainer.cnrm.cloud.google.com EdgeContainerNodePool resource.
type EdgeContainerNodePool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EdgeContainerNodePoolSpec   `json:"spec,omitempty"`
	Status EdgeContainerNodePoolStatus `json:"status,omitempty"`
}

// EdgeContainerNodePoolSpec defines the desired state of EdgeContainerNodePool.
type EdgeContainerNodePoolSpec struct {
	ClusterRef map[string]interface{} `json:"clusterRef,omitempty" yaml:"clusterRef,omitempty"`
	// Local disk encryption options. This field is only used when enabling CMEK support.
	LocalDiskEncryption map[string]interface{} `json:"localDiskEncryption,omitempty" yaml:"localDiskEncryption,omitempty"`
	// Immutable. The location of the resource.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Only machines matching this filter will be allowed to join the node pool.
	// The filtering language accepts strings like "name=<name>", and is
	// documented in more detail in [AIP-160](https://google.aip.dev/160).
	MachineFilter string `json:"machineFilter,omitempty" yaml:"machineFilter,omitempty"`
	// Configuration for each node in the NodePool.
	NodeConfig map[string]interface{} `json:"nodeConfig,omitempty" yaml:"nodeConfig,omitempty"`
	// The number of nodes in the pool.
	NodeCount int32 `json:"nodeCount,omitempty" yaml:"nodeCount,omitempty"`
	// Immutable. Name of the Google Distributed Cloud Edge zone where this node pool will be created. For example: 'us-central1-edge-customer-a'.
	NodeLocation string `json:"nodeLocation,omitempty" yaml:"nodeLocation,omitempty"`
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// EdgeContainerNodePoolStatus defines the observed state of EdgeContainerNodePool.
type EdgeContainerNodePoolStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
