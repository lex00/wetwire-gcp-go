package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EdgeContainerCluster represents a edgecontainer.cnrm.cloud.google.com EdgeContainerCluster resource.
type EdgeContainerCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EdgeContainerClusterSpec   `json:"spec,omitempty"`
	Status EdgeContainerClusterStatus `json:"status,omitempty"`
}

// EdgeContainerClusterSpec defines the desired state of EdgeContainerCluster.
type EdgeContainerClusterSpec struct {
	// Immutable. RBAC policy that will be applied and managed by GEC.
	Authorization map[string]interface{} `json:"authorization,omitempty" yaml:"authorization,omitempty"`
	// The configuration of the cluster control plane.
	ControlPlane map[string]interface{} `json:"controlPlane,omitempty" yaml:"controlPlane,omitempty"`
	// Remote control plane disk encryption options. This field is only used when
	// enabling CMEK support.
	ControlPlaneEncryption map[string]interface{} `json:"controlPlaneEncryption,omitempty" yaml:"controlPlaneEncryption,omitempty"`
	// The default maximum number of pods per node used if a maximum value is not
	// specified explicitly for a node pool in this cluster. If unspecified, the
	// Kubernetes default value will be used.
	DefaultMaxPodsPerNode int32 `json:"defaultMaxPodsPerNode,omitempty" yaml:"defaultMaxPodsPerNode,omitempty"`
	// Address pools for cluster data plane external load balancing.
	ExternalLoadBalancerIpv4AddressPools []string `json:"externalLoadBalancerIpv4AddressPools,omitempty" yaml:"externalLoadBalancerIpv4AddressPools,omitempty"`
	// Immutable. Fleet related configuration.
	// Fleets are a Google Cloud concept for logically organizing clusters,
	// letting you use and manage multi-cluster capabilities and apply
	// consistent policies across your systems.
	Fleet map[string]interface{} `json:"fleet,omitempty" yaml:"fleet,omitempty"`
	// Immutable. The location of the resource.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Cluster-wide maintenance policy configuration.
	MaintenancePolicy map[string]interface{} `json:"maintenancePolicy,omitempty" yaml:"maintenancePolicy,omitempty"`
	// Fleet related configuration.
	// Fleets are a Google Cloud concept for logically organizing clusters,
	// letting you use and manage multi-cluster capabilities and apply
	// consistent policies across your systems.
	Networking map[string]interface{} `json:"networking,omitempty" yaml:"networking,omitempty"`
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// The release channel a cluster is subscribed to. Possible values: ["RELEASE_CHANNEL_UNSPECIFIED", "NONE", "REGULAR"].
	ReleaseChannel string `json:"releaseChannel,omitempty" yaml:"releaseChannel,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Config that customers are allowed to define for GDCE system add-ons.
	SystemAddonsConfig map[string]interface{} `json:"systemAddonsConfig,omitempty" yaml:"systemAddonsConfig,omitempty"`
	// The target cluster version. For example: "1.5.0".
	TargetVersion string `json:"targetVersion,omitempty" yaml:"targetVersion,omitempty"`
}

// EdgeContainerClusterStatus defines the observed state of EdgeContainerCluster.
type EdgeContainerClusterStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
