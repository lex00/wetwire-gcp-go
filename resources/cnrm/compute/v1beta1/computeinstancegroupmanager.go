package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ComputeInstanceGroupManager represents a compute.cnrm.cloud.google.com ComputeInstanceGroupManager resource.
type ComputeInstanceGroupManager struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeInstanceGroupManagerSpec   `json:"spec,omitempty"`
	Status ComputeInstanceGroupManagerStatus `json:"status,omitempty"`
}

// ComputeInstanceGroupManagerSpec defines the desired state of ComputeInstanceGroupManager.
type ComputeInstanceGroupManagerSpec struct {
	// The autohealing policy for this managed instance group. You can specify only one value.
	AutoHealingPolicies []map[string]interface{} `json:"autoHealingPolicies,omitempty" yaml:"autoHealingPolicies,omitempty"`
	// The base instance name to use for instances in this group. The value must be 1-58 characters long. Instances are named by appending a hyphen and a random four-character string to the base instance name. The base instance name must comply with [RFC1035](https://www.ietf.org/rfc/rfc1035.txt).
	BaseInstanceName string `json:"baseInstanceName,omitempty" yaml:"baseInstanceName,omitempty"`
	// Immutable. An optional description of this resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Policy specifying the intended distribution of managed instances across zones in a regional managed instance group.
	DistributionPolicy map[string]interface{} `json:"distributionPolicy,omitempty" yaml:"distributionPolicy,omitempty"`
	// The action to perform in case of zone failure. Only one value is supported, `NO_FAILOVER`. The default is `NO_FAILOVER`. Possible values: UNKNOWN, NO_FAILOVER
	FailoverAction string `json:"failoverAction,omitempty" yaml:"failoverAction,omitempty"`
	InstanceTemplateRef map[string]interface{} `json:"instanceTemplateRef,omitempty" yaml:"instanceTemplateRef,omitempty"`
	// Immutable. The location of this resource.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Immutable. Named ports configured for the Instance Groups complementary to this Instance Group Manager.
	NamedPorts []map[string]interface{} `json:"namedPorts,omitempty" yaml:"namedPorts,omitempty"`
	// Immutable. The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	ServiceAccountRef map[string]interface{} `json:"serviceAccountRef,omitempty" yaml:"serviceAccountRef,omitempty"`
	// Stateful configuration for this Instanced Group Manager
	StatefulPolicy map[string]interface{} `json:"statefulPolicy,omitempty" yaml:"statefulPolicy,omitempty"`
	TargetPools []map[string]interface{} `json:"targetPools,omitempty" yaml:"targetPools,omitempty"`
	// The target number of running instances for this managed instance group. You can reduce this number by using the instanceGroupManager deleteInstances or abandonInstances methods. Resizing the group also changes this number.
	TargetSize int64 `json:"targetSize,omitempty" yaml:"targetSize,omitempty"`
	// The update policy for this managed instance group.
	UpdatePolicy map[string]interface{} `json:"updatePolicy,omitempty" yaml:"updatePolicy,omitempty"`
	// Specifies the instance templates used by this managed instance group to create instances. Each version is defined by an `instanceTemplate` and a `name`. Every version can appear at most once per instance group. This field overrides the top-level `instanceTemplate` field. Read more about the [relationships between these fields](/compute/docs/instance-groups/rolling-out-updates-to-managed-instance-groups#relationship_between_versions_and_instancetemplate_properties_for_a_managed_instance_group). Exactly one `version` must leave the `targetSize` field unset. That version will be applied to all remaining instances. For more information, read about [canary updates](/compute/docs/instance-groups/rolling-out-updates-to-managed-instance-groups#starting_a_canary_update).
	Versions []map[string]interface{} `json:"versions,omitempty" yaml:"versions,omitempty"`
}

// ComputeInstanceGroupManagerStatus defines the observed state of ComputeInstanceGroupManager.
type ComputeInstanceGroupManagerStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
