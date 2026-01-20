package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// VertexAIEndpoint represents a vertexai.cnrm.cloud.google.com VertexAIEndpoint resource.
type VertexAIEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VertexAIEndpointSpec   `json:"spec,omitempty"`
	Status VertexAIEndpointStatus `json:"status,omitempty"`
}

// VertexAIEndpointSpec defines the desired state of VertexAIEndpoint.
type VertexAIEndpointSpec struct {
	// The description of the Endpoint.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Required. The display name of the Endpoint. The name can be up to 128 characters long and can consist of any UTF-8 characters.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// Immutable. Customer-managed encryption key spec for an Endpoint. If set, this Endpoint and all sub-resources of this Endpoint will be secured by this key.
	EncryptionSpec map[string]interface{} `json:"encryptionSpec,omitempty" yaml:"encryptionSpec,omitempty"`
	// Optional. The full name of the Google Compute Engine network to which the Endpoint should be peered.
	// Private services access must already be configured for the network. If left unspecified, the Endpoint is not peered with any network.
	// Only one of the fields, network or enablePrivateServiceConnect, can be set.
	// Format: projects/{project_id}/global/networks/{network_name}.
	NetworkRef map[string]interface{} `json:"networkRef,omitempty" yaml:"networkRef,omitempty"`
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. The region for the resource.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// VertexAIEndpointStatus defines the observed state of VertexAIEndpoint.
type VertexAIEndpointStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
