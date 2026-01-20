package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// RedisCluster is the Schema for the RedisCluster API
type RedisCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RedisClusterSpec   `json:"spec,omitempty"`
	Status RedisClusterStatus `json:"status,omitempty"`
}

// RedisClusterSpec defines the desired state of RedisCluster.
type RedisClusterSpec struct {
	// Optional. The authorization mode of the Redis cluster. If not provided, auth feature is disabled for the cluster.
	AuthorizationMode string `json:"authorizationMode,omitempty" yaml:"authorizationMode,omitempty"`
	// Optional. The delete operation will fail when the value is set to true.
	DeletionProtectionEnabled bool `json:"deletionProtectionEnabled,omitempty" yaml:"deletionProtectionEnabled,omitempty"`
	// Immutable. Location of the resource.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Optional. The type of a redis node in the cluster. NodeType determines the underlying machine-type of a redis node.
	NodeType string `json:"nodeType,omitempty" yaml:"nodeType,omitempty"`
	// Optional. Persistence config (RDB, AOF) for the cluster.
	PersistenceConfig map[string]interface{} `json:"persistenceConfig,omitempty" yaml:"persistenceConfig,omitempty"`
	// Immutable. The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Required. Each PscConfig configures the consumer network where IPs will be designated to the cluster for client access through Private Service Connect Automation. Currently, only one PscConfig is supported.
	PscConfigs []map[string]interface{} `json:"pscConfigs,omitempty" yaml:"pscConfigs,omitempty"`
	// Optional. Key/Value pairs of customer overrides for mutable Redis Configs
	RedisConfigs map[string]string `json:"redisConfigs,omitempty" yaml:"redisConfigs,omitempty"`
	// Optional. The number of replica nodes per shard.
	ReplicaCount int32 `json:"replicaCount,omitempty" yaml:"replicaCount,omitempty"`
	// The RedisCluster name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Required. Number of shards for the Redis cluster.
	ShardCount int32 `json:"shardCount,omitempty" yaml:"shardCount,omitempty"`
	// Optional. The in-transit encryption for the Redis cluster. If not provided, encryption  is disabled for the cluster.
	TransitEncryptionMode string `json:"transitEncryptionMode,omitempty" yaml:"transitEncryptionMode,omitempty"`
	// Optional. This config will be used to determine how the customer wants us to distribute cluster resources within the region.
	ZoneDistributionConfig map[string]interface{} `json:"zoneDistributionConfig,omitempty" yaml:"zoneDistributionConfig,omitempty"`
}

// RedisClusterStatus defines the observed state of RedisCluster.
type RedisClusterStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
