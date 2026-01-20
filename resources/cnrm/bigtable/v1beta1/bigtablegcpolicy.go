package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// BigtableGCPolicy represents a bigtable.cnrm.cloud.google.com BigtableGCPolicy resource.
type BigtableGCPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BigtableGCPolicySpec   `json:"spec,omitempty"`
	Status BigtableGCPolicyStatus `json:"status,omitempty"`
}

// BigtableGCPolicySpec defines the desired state of BigtableGCPolicy.
type BigtableGCPolicySpec struct {
	// Immutable. The name of the column family.
	ColumnFamily string `json:"columnFamily,omitempty" yaml:"columnFamily,omitempty"`
	// The deletion policy for the GC policy. Setting ABANDON allows the resource
	// to be abandoned rather than deleted. This is useful for GC policy as it cannot be deleted
	// in a replicated instance. Possible values are: "ABANDON".
	DeletionPolicy string `json:"deletionPolicy,omitempty" yaml:"deletionPolicy,omitempty"`
	// Serialized JSON string for garbage collection policy. Conflicts with "mode", "max_age" and "max_version".
	GcRules string `json:"gcRules,omitempty" yaml:"gcRules,omitempty"`
	// The name of the Bigtable instance.
	InstanceRef map[string]interface{} `json:"instanceRef,omitempty" yaml:"instanceRef,omitempty"`
	// Immutable. NOTE: 'gc_rules' is more flexible, and should be preferred over this field for new resources. This field may be deprecated in the future. GC policy that applies to all cells older than the given age.
	MaxAge []map[string]interface{} `json:"maxAge,omitempty" yaml:"maxAge,omitempty"`
	// Immutable. NOTE: 'gc_rules' is more flexible, and should be preferred over this field for new resources. This field may be deprecated in the future. GC policy that applies to all versions of a cell except for the most recent.
	MaxVersion []map[string]interface{} `json:"maxVersion,omitempty" yaml:"maxVersion,omitempty"`
	// Immutable. NOTE: 'gc_rules' is more flexible, and should be preferred over this field for new resources. This field may be deprecated in the future. If multiple policies are set, you should choose between UNION OR INTERSECTION.
	Mode string `json:"mode,omitempty" yaml:"mode,omitempty"`
	// The name of the table.
	TableRef map[string]interface{} `json:"tableRef,omitempty" yaml:"tableRef,omitempty"`
}

// BigtableGCPolicyStatus defines the observed state of BigtableGCPolicy.
type BigtableGCPolicyStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
