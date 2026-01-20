package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ContainerNodePool represents a container.cnrm.cloud.google.com ContainerNodePool resource.
type ContainerNodePool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ContainerNodePoolSpec   `json:"spec,omitempty"`
	Status ContainerNodePoolStatus `json:"status,omitempty"`
}

// ContainerNodePoolSpec defines the desired state of ContainerNodePool.
type ContainerNodePoolSpec struct {
	// Configuration required by cluster autoscaler to adjust the size of the node pool to the current cluster usage. To disable autoscaling, set minNodeCount and maxNodeCount to 0.
	Autoscaling map[string]interface{} `json:"autoscaling,omitempty" yaml:"autoscaling,omitempty"`
	ClusterRef map[string]interface{} `json:"clusterRef,omitempty" yaml:"clusterRef,omitempty"`
	// Immutable. The initial number of nodes for the pool. In regional or multi-zonal clusters, this is the number of nodes per zone. Changing this will force recreation of the resource.
	InitialNodeCount int32 `json:"initialNodeCount,omitempty" yaml:"initialNodeCount,omitempty"`
	// Immutable. The location (region or zone) of the cluster.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Node management configuration, wherein auto-repair and auto-upgrade is configured.
	Management map[string]interface{} `json:"management,omitempty" yaml:"management,omitempty"`
	// Immutable. The maximum number of pods per node in this node pool. Note that this does not work on node pools which are "route-based" - that is, node pools belonging to clusters that do not have IP Aliasing enabled.
	MaxPodsPerNode int32 `json:"maxPodsPerNode,omitempty" yaml:"maxPodsPerNode,omitempty"`
	// Immutable. Creates a unique name for the node pool beginning with the specified prefix. Conflicts with name.
	NamePrefix string `json:"namePrefix,omitempty" yaml:"namePrefix,omitempty"`
	// Networking configuration for this NodePool. If specified, it overrides the cluster-level defaults.
	NetworkConfig map[string]interface{} `json:"networkConfig,omitempty" yaml:"networkConfig,omitempty"`
	// Immutable. The configuration of the nodepool.
	NodeConfig map[string]interface{} `json:"nodeConfig,omitempty" yaml:"nodeConfig,omitempty"`
	// The number of nodes per instance group. This field can be used to update the number of nodes per instance group but should not be used alongside autoscaling.
	NodeCount int32 `json:"nodeCount,omitempty" yaml:"nodeCount,omitempty"`
	// The list of zones in which the node pool's nodes should be located. Nodes must be in the region of their regional cluster or in the same region as their cluster's zone for zonal clusters. If unspecified, the cluster-level node_locations will be used.
	NodeLocations []string `json:"nodeLocations,omitempty" yaml:"nodeLocations,omitempty"`
	// Immutable. Specifies the node placement policy.
	PlacementPolicy map[string]interface{} `json:"placementPolicy,omitempty" yaml:"placementPolicy,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Specify node upgrade settings to change how many nodes GKE attempts to upgrade at once. The number of nodes upgraded simultaneously is the sum of max_surge and max_unavailable. The maximum number of nodes upgraded simultaneously is limited to 20.
	UpgradeSettings map[string]interface{} `json:"upgradeSettings,omitempty" yaml:"upgradeSettings,omitempty"`
	Version string `json:"version,omitempty" yaml:"version,omitempty"`
}

// ContainerNodePoolStatus defines the observed state of ContainerNodePool.
type ContainerNodePoolStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
