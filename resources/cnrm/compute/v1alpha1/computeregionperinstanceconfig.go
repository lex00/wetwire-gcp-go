package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ComputeRegionPerInstanceConfig represents a compute.cnrm.cloud.google.com ComputeRegionPerInstanceConfig resource.
type ComputeRegionPerInstanceConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeRegionPerInstanceConfigSpec   `json:"spec,omitempty"`
	Status ComputeRegionPerInstanceConfigStatus `json:"status,omitempty"`
}

// ComputeRegionPerInstanceConfigSpec defines the desired state of ComputeRegionPerInstanceConfig.
type ComputeRegionPerInstanceConfigSpec struct {
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
	// Immutable. Region where the containing instance group manager is located.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`
	RegionInstanceGroupManagerRef map[string]interface{} `json:"regionInstanceGroupManagerRef,omitempty" yaml:"regionInstanceGroupManagerRef,omitempty"`
	// When true, deleting this config will immediately remove any specified state from the underlying instance.
	// When false, deleting this config will *not* immediately remove any state from the underlying instance.
	// State will be removed on the next instance recreation or update.
	RemoveInstanceStateOnDestroy bool `json:"removeInstanceStateOnDestroy,omitempty" yaml:"removeInstanceStateOnDestroy,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// ComputeRegionPerInstanceConfigStatus defines the observed state of ComputeRegionPerInstanceConfig.
type ComputeRegionPerInstanceConfigStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
