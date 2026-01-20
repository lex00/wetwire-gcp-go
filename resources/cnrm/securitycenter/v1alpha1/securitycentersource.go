package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// SecurityCenterSource represents a securitycenter.cnrm.cloud.google.com SecurityCenterSource resource.
type SecurityCenterSource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SecurityCenterSourceSpec   `json:"spec,omitempty"`
	Status SecurityCenterSourceStatus `json:"status,omitempty"`
}

// SecurityCenterSourceSpec defines the desired state of SecurityCenterSource.
type SecurityCenterSourceSpec struct {
	// The description of the source (max of 1024 characters).
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// The source’s display name. A source’s display name must be unique
	// amongst its siblings, for example, two sources with the same parent
	// can't share the same display name. The display name must start and end
	// with a letter or digit, may contain letters, digits, spaces, hyphens,
	// and underscores, and can be no longer than 32 characters.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// The organization that this resource belongs to.
	OrganizationRef map[string]interface{} `json:"organizationRef,omitempty" yaml:"organizationRef,omitempty"`
	// Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// SecurityCenterSourceStatus defines the observed state of SecurityCenterSource.
type SecurityCenterSourceStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
