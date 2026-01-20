package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// BigtableCluster is the Schema for the BigtableCluster API
type BigtableCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BigtableClusterSpec   `json:"spec,omitempty"`
	Status BigtableClusterStatus `json:"status,omitempty"`
}

// BigtableClusterSpec defines the desired state of BigtableCluster.
type BigtableClusterSpec struct {
	// Configuration for this cluster.
	ClusterConfig map[string]interface{} `json:"clusterConfig,omitempty" yaml:"clusterConfig,omitempty"`
	// Immutable. The type of storage used by this cluster to serve its parent instance's tables, unless explicitly overridden.
	DefaultStorageType string `json:"defaultStorageType,omitempty" yaml:"defaultStorageType,omitempty"`
	// Immutable. The encryption configuration for CMEK-protected clusters.
	EncryptionConfig map[string]interface{} `json:"encryptionConfig,omitempty" yaml:"encryptionConfig,omitempty"`
	// InstanceRef defines the resource reference to BigtableInstance, which "External" field holds the GCP identifier for the KRM object.
	InstanceRef map[string]interface{} `json:"instanceRef,omitempty" yaml:"instanceRef,omitempty"`
	// Immutable. The location where this cluster's nodes and storage reside. For best performance, clients should be located as close as possible to this cluster. Currently only zones are supported, so values should be of the form `projects/{project}/locations/{zone}`.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Immutable. The node scaling factor of this cluster.
	NodeScalingFactor string `json:"nodeScalingFactor,omitempty" yaml:"nodeScalingFactor,omitempty"`
	// The BigtableCluster name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// The number of nodes allocated to this cluster. More nodes enable higher throughput and more consistent performance.
	ServeNodes int32 `json:"serveNodes,omitempty" yaml:"serveNodes,omitempty"`
}

// BigtableClusterStatus defines the observed state of BigtableCluster.
type BigtableClusterStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
