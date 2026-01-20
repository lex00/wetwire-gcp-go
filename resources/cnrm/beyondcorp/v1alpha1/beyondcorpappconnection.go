package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// BeyondCorpAppConnection represents a beyondcorp.cnrm.cloud.google.com BeyondCorpAppConnection resource.
type BeyondCorpAppConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BeyondCorpAppConnectionSpec   `json:"spec,omitempty"`
	Status BeyondCorpAppConnectionStatus `json:"status,omitempty"`
}

// BeyondCorpAppConnectionSpec defines the desired state of BeyondCorpAppConnection.
type BeyondCorpAppConnectionSpec struct {
	// Address of the remote application endpoint for the BeyondCorp AppConnection.
	ApplicationEndpoint map[string]interface{} `json:"applicationEndpoint,omitempty" yaml:"applicationEndpoint,omitempty"`
	// List of AppConnectors that are authorised to be associated with this AppConnection.
	Connectors []string `json:"connectors,omitempty" yaml:"connectors,omitempty"`
	// An arbitrary user-provided name for the AppConnection.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// Gateway used by the AppConnection.
	Gateway map[string]interface{} `json:"gateway,omitempty" yaml:"gateway,omitempty"`
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. The region of the AppConnection.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Immutable. The type of network connectivity used by the AppConnection. Refer to
	// https://cloud.google.com/beyondcorp/docs/reference/rest/v1/projects.locations.appConnections#type
	// for a list of possible values.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}

// BeyondCorpAppConnectionStatus defines the observed state of BeyondCorpAppConnection.
type BeyondCorpAppConnectionStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
