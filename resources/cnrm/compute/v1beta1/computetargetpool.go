package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ComputeTargetPool represents a compute.cnrm.cloud.google.com ComputeTargetPool resource.
type ComputeTargetPool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeTargetPoolSpec   `json:"spec,omitempty"`
	Status ComputeTargetPoolStatus `json:"status,omitempty"`
}

// ComputeTargetPoolSpec defines the desired state of ComputeTargetPool.
type ComputeTargetPoolSpec struct {
	BackupTargetPoolRef map[string]interface{} `json:"backupTargetPoolRef,omitempty" yaml:"backupTargetPoolRef,omitempty"`
	// Immutable. Textual description field.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Immutable. Ratio (0 to 1) of failed nodes before using the backup pool (which must also be set).
	FailoverRatio float64 `json:"failoverRatio,omitempty" yaml:"failoverRatio,omitempty"`
	HealthChecks []map[string]interface{} `json:"healthChecks,omitempty" yaml:"healthChecks,omitempty"`
	Instances []map[string]interface{} `json:"instances,omitempty" yaml:"instances,omitempty"`
	// Immutable. Where the target pool resides. Defaults to project region.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// The resource URL for the security policy associated with this target pool.
	SecurityPolicyRef map[string]interface{} `json:"securityPolicyRef,omitempty" yaml:"securityPolicyRef,omitempty"`
	// Immutable. How to distribute load. Options are "NONE" (no affinity). "CLIENT_IP" (hash of the source/dest addresses / ports), and "CLIENT_IP_PROTO" also includes the protocol (default "NONE").
	SessionAffinity string `json:"sessionAffinity,omitempty" yaml:"sessionAffinity,omitempty"`
}

// ComputeTargetPoolStatus defines the observed state of ComputeTargetPool.
type ComputeTargetPoolStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
