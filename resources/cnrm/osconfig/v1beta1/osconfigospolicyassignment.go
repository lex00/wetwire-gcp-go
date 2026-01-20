package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// OSConfigOSPolicyAssignment represents a osconfig.cnrm.cloud.google.com OSConfigOSPolicyAssignment resource.
type OSConfigOSPolicyAssignment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OSConfigOSPolicyAssignmentSpec   `json:"spec,omitempty"`
	Status OSConfigOSPolicyAssignmentStatus `json:"status,omitempty"`
}

// OSConfigOSPolicyAssignmentSpec defines the desired state of OSConfigOSPolicyAssignment.
type OSConfigOSPolicyAssignmentSpec struct {
	// OS policy assignment description. Length of the description is limited to 1024 characters.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Required. Filter to select VMs.
	InstanceFilter map[string]interface{} `json:"instanceFilter,omitempty" yaml:"instanceFilter,omitempty"`
	// Immutable. The location for the resource
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Required. List of OS policies to be applied to the VMs.
	OsPolicies []map[string]interface{} `json:"osPolicies,omitempty" yaml:"osPolicies,omitempty"`
	// Immutable. The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Required. Rollout to deploy the OS policy assignment. A rollout is triggered in the following situations: 1) OSPolicyAssignment is created. 2) OSPolicyAssignment is updated and the update contains changes to one of the following fields: - instance_filter - os_policies 3) OSPolicyAssignment is deleted.
	Rollout map[string]interface{} `json:"rollout,omitempty" yaml:"rollout,omitempty"`
	// Set to true to skip awaiting rollout during resource creation and update.
	SkipAwaitRollout bool `json:"skipAwaitRollout,omitempty" yaml:"skipAwaitRollout,omitempty"`
}

// OSConfigOSPolicyAssignmentStatus defines the observed state of OSConfigOSPolicyAssignment.
type OSConfigOSPolicyAssignmentStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
