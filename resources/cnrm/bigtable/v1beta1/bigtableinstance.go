package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// BigtableInstance is the Schema for the BigtableInstance API
type BigtableInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BigtableInstanceSpec   `json:"spec,omitempty"`
	Status BigtableInstanceStatus `json:"status,omitempty"`
}

// BigtableInstanceSpec defines the desired state of BigtableInstance.
type BigtableInstanceSpec struct {
	// A block of cluster configuration options. This can be specified at least once.
	Cluster []map[string]interface{} `json:"cluster,omitempty" yaml:"cluster,omitempty"`
	// DEPRECATED. This field no longer serves any function and is intended to be dropped in a later version of the resource.
	DeletionProtection bool `json:"deletionProtection,omitempty" yaml:"deletionProtection,omitempty"`
	// Required. The descriptive name for this instance as it appears in UIs. Can be changed at any time, but should be kept globally unique to avoid confusion.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// DEPRECATED. It is recommended to leave this field unspecified since the distinction between "DEVELOPMENT" and "PRODUCTION" instances is going away, and all instances will become "PRODUCTION" instances. This means that new and existing "DEVELOPMENT" instances will be converted to "PRODUCTION" instances. It is recommended for users to use "PRODUCTION" instances in any case, since a 1-node "PRODUCTION" instance is functionally identical to a "DEVELOPMENT" instance, but without the accompanying restrictions. The instance type to create. One of "DEVELOPMENT" or "PRODUCTION". Defaults to "PRODUCTION".
	InstanceType string `json:"instanceType,omitempty" yaml:"instanceType,omitempty"`
	// The Instance name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// BigtableInstanceStatus defines the observed state of BigtableInstance.
type BigtableInstanceStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
