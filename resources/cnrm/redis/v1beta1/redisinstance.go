package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// RedisInstance represents a redis.cnrm.cloud.google.com RedisInstance resource.
type RedisInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RedisInstanceSpec   `json:"spec,omitempty"`
	Status RedisInstanceStatus `json:"status,omitempty"`
}

// RedisInstanceSpec defines the desired state of RedisInstance.
type RedisInstanceSpec struct {
	// Immutable. Only applicable to STANDARD_HA tier which protects the instance
	// against zonal failures by provisioning it across two zones.
	// If provided, it must be a different zone from the one provided in
	// [locationId].
	AlternativeLocationID string `json:"alternativeLocationId,omitempty" yaml:"alternativeLocationId,omitempty"`
	// Optional. Indicates whether OSS Redis AUTH is enabled for the
	// instance. If set to "true" AUTH is enabled on the instance.
	// Default value is "false" meaning AUTH is disabled.
	AuthEnabled bool `json:"authEnabled,omitempty" yaml:"authEnabled,omitempty"`
	// Output only. AUTH String set on the instance. This field will only be populated if auth_enabled is true.
	AuthString string `json:"authString,omitempty" yaml:"authString,omitempty"`
	// The network to which the instance is connected. If left
	// unspecified, the default network will be used.
	AuthorizedNetworkRef map[string]interface{} `json:"authorizedNetworkRef,omitempty" yaml:"authorizedNetworkRef,omitempty"`
	// Immutable. The connection mode of the Redis instance. Default value: "DIRECT_PEERING" Possible values: ["DIRECT_PEERING", "PRIVATE_SERVICE_ACCESS"].
	ConnectMode string `json:"connectMode,omitempty" yaml:"connectMode,omitempty"`
	// Immutable. Optional. The KMS key reference that you want to use to
	// encrypt the data at rest for this Redis instance. If this is
	// provided, CMEK is enabled.
	CustomerManagedKeyRef map[string]interface{} `json:"customerManagedKeyRef,omitempty" yaml:"customerManagedKeyRef,omitempty"`
	// An arbitrary and optional user-provided name for the instance.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// Immutable. The zone where the instance will be provisioned. If not provided,
	// the service will choose a zone for the instance. For STANDARD_HA tier,
	// instances will be created across two zones for protection against
	// zonal failures. If [alternativeLocationId] is also provided, it must
	// be different from [locationId].
	LocationID string `json:"locationId,omitempty" yaml:"locationId,omitempty"`
	// Maintenance policy for an instance.
	MaintenancePolicy map[string]interface{} `json:"maintenancePolicy,omitempty" yaml:"maintenancePolicy,omitempty"`
	// Upcoming maintenance schedule.
	MaintenanceSchedule []map[string]interface{} `json:"maintenanceSchedule,omitempty" yaml:"maintenanceSchedule,omitempty"`
	// Redis memory size in GiB.
	MemorySizeGb int32 `json:"memorySizeGb,omitempty" yaml:"memorySizeGb,omitempty"`
	// Persistence configuration for an instance.
	PersistenceConfig map[string]interface{} `json:"persistenceConfig,omitempty" yaml:"persistenceConfig,omitempty"`
	// Optional. Read replica mode. Can only be specified when trying to create the instance.
	// If not set, Memorystore Redis backend will default to READ_REPLICAS_DISABLED.
	// - READ_REPLICAS_DISABLED: If disabled, read endpoint will not be provided and the
	// instance cannot scale up or down the number of replicas.
	// - READ_REPLICAS_ENABLED: If enabled, read endpoint will be provided and the instance
	// can scale up and down the number of replicas. Possible values: ["READ_REPLICAS_DISABLED", "READ_REPLICAS_ENABLED"].
	ReadReplicasMode string `json:"readReplicasMode,omitempty" yaml:"readReplicasMode,omitempty"`
	// Redis configuration parameters, according to http://redis.io/topics/config.
	// Please check Memorystore documentation for the list of supported parameters:
	// https://cloud.google.com/memorystore/docs/redis/reference/rest/v1/projects.locations.instances#Instance.FIELDS.redis_configs.
	RedisConfigs map[string]string `json:"redisConfigs,omitempty" yaml:"redisConfigs,omitempty"`
	// The version of Redis software. If not provided, latest supported
	// version will be used. Please check the API documentation linked
	// at the top for the latest valid values.
	RedisVersion string `json:"redisVersion,omitempty" yaml:"redisVersion,omitempty"`
	// Immutable. The name of the Redis region of the instance.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`
	// Optional. The number of replica nodes. The valid range for the Standard Tier with
	// read replicas enabled is [1-5] and defaults to 2. If read replicas are not enabled
	// for a Standard Tier instance, the only valid value is 1 and the default is 1.
	// The valid value for basic tier is 0 and the default is also 0.
	ReplicaCount int32 `json:"replicaCount,omitempty" yaml:"replicaCount,omitempty"`
	// Immutable. The CIDR range of internal addresses that are reserved for this
	// instance. If not provided, the service will choose an unused /29
	// block, for example, 10.0.0.0/29 or 192.168.0.0/29. Ranges must be
	// unique and non-overlapping with existing subnets in an authorized
	// network.
	ReservedIPRange string `json:"reservedIpRange,omitempty" yaml:"reservedIpRange,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Optional. Additional IP range for node placement. Required when enabling read replicas on
	// an existing instance. For DIRECT_PEERING mode value must be a CIDR range of size /28, or
	// "auto". For PRIVATE_SERVICE_ACCESS mode value must be the name of an allocated address
	// range associated with the private service access connection, or "auto".
	SecondaryIPRange string `json:"secondaryIpRange,omitempty" yaml:"secondaryIpRange,omitempty"`
	// Immutable. The service tier of the instance. Must be one of these values:
	// - BASIC: standalone instance
	// - STANDARD_HA: highly available primary/replica instances Default value: "BASIC" Possible values: ["BASIC", "STANDARD_HA"].
	Tier string `json:"tier,omitempty" yaml:"tier,omitempty"`
	// Immutable. The TLS mode of the Redis instance, If not provided, TLS is disabled for the instance.
	// - SERVER_AUTHENTICATION: Client to Server traffic encryption enabled with server authentication Default value: "DISABLED" Possible values: ["SERVER_AUTHENTICATION", "DISABLED"].
	TransitEncryptionMode string `json:"transitEncryptionMode,omitempty" yaml:"transitEncryptionMode,omitempty"`
}

// RedisInstanceStatus defines the observed state of RedisInstance.
type RedisInstanceStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
