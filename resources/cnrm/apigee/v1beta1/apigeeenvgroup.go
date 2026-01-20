package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ApigeeEnvgroup is the Schema for the ApigeeEnvgroup API
type ApigeeEnvgroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ApigeeEnvgroupSpec   `json:"spec,omitempty"`
	Status ApigeeEnvgroupStatus `json:"status,omitempty"`
}

// ApigeeEnvgroupSpec defines the desired state of ApigeeEnvgroup.
type ApigeeEnvgroupSpec struct {
	// Host names for this environment group.
	Hostnames []string `json:"hostnames,omitempty" yaml:"hostnames,omitempty"`
	// ApigeeOrganizationRef is a reference to a ApigeeOrganization resource.
	OrganizationRef map[string]interface{} `json:"organizationRef,omitempty" yaml:"organizationRef,omitempty"`
	// The ApigeeEnvgroup name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// ApigeeEnvgroupStatus defines the observed state of ApigeeEnvgroup.
type ApigeeEnvgroupStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
