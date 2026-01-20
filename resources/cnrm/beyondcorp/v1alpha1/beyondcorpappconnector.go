package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// BeyondCorpAppConnector represents a beyondcorp.cnrm.cloud.google.com BeyondCorpAppConnector resource.
type BeyondCorpAppConnector struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BeyondCorpAppConnectorSpec   `json:"spec,omitempty"`
	Status BeyondCorpAppConnectorStatus `json:"status,omitempty"`
}

// BeyondCorpAppConnectorSpec defines the desired state of BeyondCorpAppConnector.
type BeyondCorpAppConnectorSpec struct {
	// An arbitrary user-provided name for the AppConnector.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// Principal information about the Identity of the AppConnector.
	PrincipalInfo map[string]interface{} `json:"principalInfo,omitempty" yaml:"principalInfo,omitempty"`
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. The region of the AppConnector.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// BeyondCorpAppConnectorStatus defines the observed state of BeyondCorpAppConnector.
type BeyondCorpAppConnectorStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
