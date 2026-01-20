package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ComputePerInstanceConfig represents a compute.cnrm.cloud.google.com ComputePerInstanceConfig resource.
type ComputePerInstanceConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputePerInstanceConfigSpec   `json:"spec,omitempty"`
	Status ComputePerInstanceConfigStatus `json:"status,omitempty"`
}

// ComputePerInstanceConfigSpec defines the desired state of ComputePerInstanceConfig.
type ComputePerInstanceConfigSpec struct {
	InstanceGroupManagerRef map[string]interface{} `json:"instanceGroupManagerRef,omitempty" yaml:"instanceGroupManagerRef,omitempty"`
	// The minimal action to perform on the instance during an update.
	// Default is 'NONE'. Possible values are:
	// * REPLACE
	// * RESTART
	// * REFRESH
	// * NONE.
	MinimalAction string `json:"minimalAction,omitempty" yaml:"minimalAction,omitempty"`
	// The most disruptive action to perform on the instance during an update.
	// Default is 'REPLACE'. Possible values are:
	// * REPLACE
	// * RESTART
	// * REFRESH
	// * NONE.
	MostDisruptiveAllowedAction string `json:"mostDisruptiveAllowedAction,omitempty" yaml:"mostDisruptiveAllowedAction,omitempty"`
	// The preserved state for this instance.
	PreservedState map[string]interface{} `json:"preservedState,omitempty" yaml:"preservedState,omitempty"`
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// When true, deleting this config will immediately remove any specified state from the underlying instance.
	// When false, deleting this config will *not* immediately remove any state from the underlying instance.
	// State will be removed on the next instance recreation or update.
	RemoveInstanceStateOnDestroy bool `json:"removeInstanceStateOnDestroy,omitempty" yaml:"removeInstanceStateOnDestroy,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Immutable. Zone where the containing instance group manager is located.
	Zone string `json:"zone,omitempty" yaml:"zone,omitempty"`
}

// ComputePerInstanceConfigStatus defines the observed state of ComputePerInstanceConfig.
type ComputePerInstanceConfigStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
