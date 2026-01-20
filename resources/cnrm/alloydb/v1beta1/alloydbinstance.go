package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// AlloyDBInstance is the Schema for the AlloyDBInstance API
type AlloyDBInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AlloyDBInstanceSpec   `json:"spec,omitempty"`
	Status AlloyDBInstanceStatus `json:"status,omitempty"`
}

// AlloyDBInstanceSpec defines the desired state of AlloyDBInstance.
type AlloyDBInstanceSpec struct {
	// Annotations to allow client tools to store small amount of arbitrary data. This is distinct from labels.
	Annotations map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	// Availability type of an Instance. If empty, defaults to REGIONAL for primary instances.
	// For read pools, availabilityType is always UNSPECIFIED. Instances in the
	// read pools are evenly distributed across available zones within the region
	// (i.e. read pools with more than one node will have a node in at least two zones).
	// Possible values: ["AVAILABILITY_TYPE_UNSPECIFIED", "ZONAL", "REGIONAL"].
	AvailabilityType string `json:"availabilityType,omitempty" yaml:"availabilityType,omitempty"`
	// The AlloyDBInstance cluster that this resource belongs to.
	ClusterRef map[string]interface{} `json:"clusterRef,omitempty" yaml:"clusterRef,omitempty"`
	// Database flags. Set at instance level. * They are copied from primary instance on read instance creation. * Read instances can set new or override existing flags that are relevant for reads, e.g. for enabling columnar cache on a read instance. Flags set on read instance may or may not be present on primary.
	DatabaseFlags map[string]string `json:"databaseFlags,omitempty" yaml:"databaseFlags,omitempty"`
	// User-settable and human-readable display name for the Instance.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// The Compute Engine zone that the instance should serve from, per https://cloud.google.com/compute/docs/regions-zones This can ONLY be specified for ZONAL instances. If present for a REGIONAL instance, an error will be thrown. If this is absent for a ZONAL instance, instance is created in a random zone with available capacity.
	GceZone string `json:"gceZone,omitempty" yaml:"gceZone,omitempty"`
	// Not recommended. We recommend that you use `instanceTypeRef` instead. The type of the instance. Possible values: [PRIMARY, READ_POOL, SECONDARY]
	InstanceType string `json:"instanceType,omitempty" yaml:"instanceType,omitempty"`
	// The type of instance.
	// Possible values: ["PRIMARY", "READ_POOL", "SECONDARY"]
	// For PRIMARY and SECONDARY instances, set the value to refer to the name of the associated cluster.
	// This is recommended because the instance type of primary and secondary instances is tied to the cluster type of the associated cluster.
	// If the secondary cluster is promoted to primary cluster, then the associated secondary instance also becomes primary instance.
	// Example:
	// instanceTypeRef:
	// name: clusterName
	// For instances of type READ_POOL, set the value using external keyword.
	// Example:
	// instanceTypeRef:
	// external: READ_POOL
	// If the instance type is SECONDARY, the delete instance operation does not delete the secondary instance but abandons it instead.
	// Use deletionPolicy = "FORCE" in the associated secondary cluster and delete the cluster forcefully to delete the secondary cluster as well its associated secondary instance.
	InstanceTypeRef map[string]interface{} `json:"instanceTypeRef,omitempty" yaml:"instanceTypeRef,omitempty"`
	// Configurations for the machines that host the underlying database engine.
	MachineConfig map[string]interface{} `json:"machineConfig,omitempty" yaml:"machineConfig,omitempty"`
	// Instance level network configuration.
	NetworkConfig map[string]interface{} `json:"networkConfig,omitempty" yaml:"networkConfig,omitempty"`
	ObservabilityConfig map[string]interface{} `json:"observabilityConfig,omitempty" yaml:"observabilityConfig,omitempty"`
	QueryInsightsConfig map[string]interface{} `json:"queryInsightsConfig,omitempty" yaml:"queryInsightsConfig,omitempty"`
	// Read pool specific config. If the instance type is READ_POOL, this configuration must be provided.
	ReadPoolConfig map[string]interface{} `json:"readPoolConfig,omitempty" yaml:"readPoolConfig,omitempty"`
	// Optional. The instanceId of the resource. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// AlloyDBInstanceStatus defines the observed state of AlloyDBInstance.
type AlloyDBInstanceStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
