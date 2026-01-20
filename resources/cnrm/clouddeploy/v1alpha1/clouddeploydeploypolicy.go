package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CloudDeployDeployPolicy is the Schema for the CloudDeployDeployPolicy API
type CloudDeployDeployPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CloudDeployDeployPolicySpec   `json:"spec,omitempty"`
	Status CloudDeployDeployPolicyStatus `json:"status,omitempty"`
}

// CloudDeployDeployPolicySpec defines the desired state of CloudDeployDeployPolicy.
type CloudDeployDeployPolicySpec struct {
	// Description of the `DeployPolicy`. Max length is 255 characters.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Immutable.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// The DeployDeployPolicy name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Required. Rules to apply. At least one rule must be present.
	Rules []map[string]interface{} `json:"rules,omitempty" yaml:"rules,omitempty"`
	// Required. Selected resources to which the policy will be applied. At least one selector is required. If one selector matches the resource the policy applies. For example, if there are two selectors and the action being attempted matches one of them, the policy will apply to that action.
	Selectors []map[string]interface{} `json:"selectors,omitempty" yaml:"selectors,omitempty"`
	// When suspended, the policy will not prevent actions from occurring, even if the action violates the policy.
	Suspended bool `json:"suspended,omitempty" yaml:"suspended,omitempty"`
}

// CloudDeployDeployPolicyStatus defines the observed state of CloudDeployDeployPolicy.
type CloudDeployDeployPolicyStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
