package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// OrgPolicyCustomConstraint is the Schema for the OrgPolicyCustomConstraint API
type OrgPolicyCustomConstraint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OrgPolicyCustomConstraintSpec   `json:"spec,omitempty"`
	Status OrgPolicyCustomConstraintStatus `json:"status,omitempty"`
}

// OrgPolicyCustomConstraintSpec defines the desired state of OrgPolicyCustomConstraint.
type OrgPolicyCustomConstraintSpec struct {
	// Allow or deny type.
	ActionType string `json:"actionType,omitempty" yaml:"actionType,omitempty"`
	// Org policy condition/expression. For example:
	// `resource.instanceName.matches("[production|test]_.*_(\d)+")` or,
	// `resource.management.auto_upgrade == true`
	// The max length of the condition is 1000 characters.
	Condition string `json:"condition,omitempty" yaml:"condition,omitempty"`
	// Detailed information about this custom policy constraint. The max length of the description is 2000 characters.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// One line display name for the UI. The max length of the display_name is 200 characters.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// All the operations being applied for this constraint.
	MethodTypes []string `json:"methodTypes,omitempty" yaml:"methodTypes,omitempty"`
	// The Organization that this resource belongs to.
	OrganizationRef map[string]interface{} `json:"organizationRef,omitempty" yaml:"organizationRef,omitempty"`
	// The OrgPolicyCustomConstraint name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Immutable. The resource instance type on which this policy applies. Format
	// will be of the form : `<canonical service name>/<type>` Example:
	// * `compute.googleapis.com/Instance`.
	ResourceTypes []string `json:"resourceTypes,omitempty" yaml:"resourceTypes,omitempty"`
}

// OrgPolicyCustomConstraintStatus defines the observed state of OrgPolicyCustomConstraint.
type OrgPolicyCustomConstraintStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
