package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DialogflowFulfillment represents a dialogflow.cnrm.cloud.google.com DialogflowFulfillment resource.
type DialogflowFulfillment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DialogflowFulfillmentSpec   `json:"spec,omitempty"`
	Status DialogflowFulfillmentStatus `json:"status,omitempty"`
}

// DialogflowFulfillmentSpec defines the desired state of DialogflowFulfillment.
type DialogflowFulfillmentSpec struct {
	// The human-readable name of the fulfillment, unique within the agent.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// Whether fulfillment is enabled.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	// The field defines whether the fulfillment is enabled for certain features.
	Features []map[string]interface{} `json:"features,omitempty" yaml:"features,omitempty"`
	// Represents configuration for a generic web service. Dialogflow supports two mechanisms for authentications: - Basic authentication with username and password. - Authentication with additional authentication headers.
	GenericWebService map[string]interface{} `json:"genericWebService,omitempty" yaml:"genericWebService,omitempty"`
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// DialogflowFulfillmentStatus defines the observed state of DialogflowFulfillment.
type DialogflowFulfillmentStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
