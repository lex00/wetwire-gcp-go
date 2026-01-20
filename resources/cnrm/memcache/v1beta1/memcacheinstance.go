package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MemcacheInstance represents a memcache.cnrm.cloud.google.com MemcacheInstance resource.
type MemcacheInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MemcacheInstanceSpec   `json:"spec,omitempty"`
	Status MemcacheInstanceStatus `json:"status,omitempty"`
}

// MemcacheInstanceSpec defines the desired state of MemcacheInstance.
type MemcacheInstanceSpec struct {
	// A user-visible name for the instance.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// Maintenance policy for an instance.
	MaintenancePolicy map[string]interface{} `json:"maintenancePolicy,omitempty" yaml:"maintenancePolicy,omitempty"`
	// Immutable. User-specified parameters for this memcache instance.
	MemcacheParameters map[string]interface{} `json:"memcacheParameters,omitempty" yaml:"memcacheParameters,omitempty"`
	// The major version of Memcached software. If not provided, latest supported version will be used.
	// Currently the latest supported major version is MEMCACHE_1_5. The minor version will be automatically
	// determined by our system based on the latest supported minor version. Default value: "MEMCACHE_1_5" Possible values: ["MEMCACHE_1_5"].
	MemcacheVersion string `json:"memcacheVersion,omitempty" yaml:"memcacheVersion,omitempty"`
	// The full name of the network to connect the instance to.
	NetworkRef map[string]interface{} `json:"networkRef,omitempty" yaml:"networkRef,omitempty"`
	// Immutable. Configuration for memcache nodes.
	NodeConfig map[string]interface{} `json:"nodeConfig,omitempty" yaml:"nodeConfig,omitempty"`
	// Number of nodes in the memcache instance.
	NodeCount int32 `json:"nodeCount,omitempty" yaml:"nodeCount,omitempty"`
	// Immutable. The region of the Memcache instance. If it is not provided, the provider region is used.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Immutable. Zones where memcache nodes should be provisioned.  If not
	// provided, all zones will be used.
	Zones []string `json:"zones,omitempty" yaml:"zones,omitempty"`
}

// MemcacheInstanceStatus defines the observed state of MemcacheInstance.
type MemcacheInstanceStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
