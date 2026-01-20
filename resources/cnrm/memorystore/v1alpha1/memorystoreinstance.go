package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MemorystoreInstance is the Schema for the MemorystoreInstance API
type MemorystoreInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MemorystoreInstanceSpec   `json:"spec,omitempty"`
	Status MemorystoreInstanceStatus `json:"status,omitempty"`
}

// MemorystoreInstanceSpec defines the desired state of MemorystoreInstance.
type MemorystoreInstanceSpec struct {
	// Optional. Immutable. Authorization mode of the instance.
	AuthorizationMode string `json:"authorizationMode,omitempty" yaml:"authorizationMode,omitempty"`
	// Optional. If set to true deletion of the instance will fail.
	DeletionProtectionEnabled bool `json:"deletionProtectionEnabled,omitempty" yaml:"deletionProtectionEnabled,omitempty"`
	// Optional. Endpoints for the instance.
	Endpoints []map[string]interface{} `json:"endpoints,omitempty" yaml:"endpoints,omitempty"`
	// Optional. User-provided engine configurations for the instance.
	EngineConfigs map[string]string `json:"engineConfigs,omitempty" yaml:"engineConfigs,omitempty"`
	// Optional. Immutable. Engine version of the instance.
	EngineVersion string `json:"engineVersion,omitempty" yaml:"engineVersion,omitempty"`
	// Optional. Labels to represent user-provided metadata.
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	// Immutable.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Optional. The mode config for the instance.
	Mode string `json:"mode,omitempty" yaml:"mode,omitempty"`
	// Optional. Immutable. Machine type for individual nodes of the instance.
	NodeType string `json:"nodeType,omitempty" yaml:"nodeType,omitempty"`
	// Optional. Persistence configuration of the instance.
	PersistenceConfig map[string]interface{} `json:"persistenceConfig,omitempty" yaml:"persistenceConfig,omitempty"`
	// The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Optional. Number of replica nodes per shard. If omitted the default is 0 replicas.
	ReplicaCount int32 `json:"replicaCount,omitempty" yaml:"replicaCount,omitempty"`
	// The MemorystoreInstance name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Optional. Number of shards for the instance.
	ShardCount int32 `json:"shardCount,omitempty" yaml:"shardCount,omitempty"`
	// Optional. Immutable. In-transit encryption mode of the instance.
	TransitEncryptionMode string `json:"transitEncryptionMode,omitempty" yaml:"transitEncryptionMode,omitempty"`
	// Optional. Immutable. Zone distribution configuration of the instance for node allocation.
	ZoneDistributionConfig map[string]interface{} `json:"zoneDistributionConfig,omitempty" yaml:"zoneDistributionConfig,omitempty"`
}

// MemorystoreInstanceStatus defines the observed state of MemorystoreInstance.
type MemorystoreInstanceStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
