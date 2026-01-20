package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// VertexAIIndexEndpoint represents a vertexai.cnrm.cloud.google.com VertexAIIndexEndpoint resource.
type VertexAIIndexEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VertexAIIndexEndpointSpec   `json:"spec,omitempty"`
	Status VertexAIIndexEndpointStatus `json:"status,omitempty"`
}

// VertexAIIndexEndpointSpec defines the desired state of VertexAIIndexEndpoint.
type VertexAIIndexEndpointSpec struct {
	// The description of the Index.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// The display name of the Index. The name can be up to 128 characters long and can consist of any UTF-8 characters.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// Immutable. The full name of the Google Compute Engine [network](https://cloud.google.com//compute/docs/networks-and-firewalls#networks) to which the index endpoint should be peered.
	// Private services access must already be configured for the network. If left unspecified, the index endpoint is not peered with any network.
	// [Format](https://cloud.google.com/compute/docs/reference/rest/v1/networks/insert): 'projects/{project}/global/networks/{network}'.
	// Where '{project}' is a project number, as in '12345', and '{network}' is network name.
	Network string `json:"network,omitempty" yaml:"network,omitempty"`
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. If true, the deployed index will be accessible through public endpoint.
	PublicEndpointEnabled bool `json:"publicEndpointEnabled,omitempty" yaml:"publicEndpointEnabled,omitempty"`
	// Immutable. The region of the index endpoint. eg us-central1.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`
	// Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// VertexAIIndexEndpointStatus defines the observed state of VertexAIIndexEndpoint.
type VertexAIIndexEndpointStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
