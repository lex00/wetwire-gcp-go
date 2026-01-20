package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ResourceManagerPolicy represents a resourcemanager.cnrm.cloud.google.com ResourceManagerPolicy resource.
type ResourceManagerPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ResourceManagerPolicySpec   `json:"spec,omitempty"`
	Status ResourceManagerPolicyStatus `json:"status,omitempty"`
}

// ResourceManagerPolicySpec defines the desired state of ResourceManagerPolicy.
type ResourceManagerPolicySpec struct {
	// A boolean policy is a constraint that is either enforced or not.
	BooleanPolicy map[string]interface{} `json:"booleanPolicy,omitempty" yaml:"booleanPolicy,omitempty"`
	// Immutable. The name of the Constraint the Policy is configuring, for example, serviceuser.services.
	Constraint string `json:"constraint,omitempty" yaml:"constraint,omitempty"`
	// The folder on which to configure the constraint. Only one of
	// projectRef, folderRef, or organizationRef may be specified.
	FolderRef map[string]interface{} `json:"folderRef,omitempty" yaml:"folderRef,omitempty"`
	// A policy that can define specific values that are allowed or denied for the given constraint. It can also be used to allow or deny all values. .
	ListPolicy map[string]interface{} `json:"listPolicy,omitempty" yaml:"listPolicy,omitempty"`
	// The organization on which to configure the constraint. Only one of
	// projectRef, folderRef, or organizationRef may be specified.
	OrganizationRef map[string]interface{} `json:"organizationRef,omitempty" yaml:"organizationRef,omitempty"`
	// The project on which to configure the constraint. Only one of
	// projectRef, folderRef, or organizationRef may be specified.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// A restore policy is a constraint to restore the default policy.
	RestorePolicy map[string]interface{} `json:"restorePolicy,omitempty" yaml:"restorePolicy,omitempty"`
	// Version of the Policy. Default version is 0.
	Version int32 `json:"version,omitempty" yaml:"version,omitempty"`
}

// ResourceManagerPolicyStatus defines the observed state of ResourceManagerPolicy.
type ResourceManagerPolicyStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
