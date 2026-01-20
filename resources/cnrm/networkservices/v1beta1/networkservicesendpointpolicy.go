package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NetworkServicesEndpointPolicy represents a networkservices.cnrm.cloud.google.com NetworkServicesEndpointPolicy resource.
type NetworkServicesEndpointPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NetworkServicesEndpointPolicySpec   `json:"spec,omitempty"`
	Status NetworkServicesEndpointPolicyStatus `json:"status,omitempty"`
}

// NetworkServicesEndpointPolicySpec defines the desired state of NetworkServicesEndpointPolicy.
type NetworkServicesEndpointPolicySpec struct {
	AuthorizationPolicyRef map[string]interface{} `json:"authorizationPolicyRef,omitempty" yaml:"authorizationPolicyRef,omitempty"`
	ClientTlsPolicyRef map[string]interface{} `json:"clientTlsPolicyRef,omitempty" yaml:"clientTlsPolicyRef,omitempty"`
	// Optional. A free-text description of the resource. Max length 1024 characters.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Required. A matcher that selects endpoints to which the policies should be applied.
	EndpointMatcher map[string]interface{} `json:"endpointMatcher,omitempty" yaml:"endpointMatcher,omitempty"`
	// Immutable. The location for the resource
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Immutable. The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	ServerTlsPolicyRef map[string]interface{} `json:"serverTlsPolicyRef,omitempty" yaml:"serverTlsPolicyRef,omitempty"`
	// Optional. Port selector for the (matched) endpoints. If no port selector is provided, the matched config is applied to all ports.
	TrafficPortSelector map[string]interface{} `json:"trafficPortSelector,omitempty" yaml:"trafficPortSelector,omitempty"`
	// Required. The type of endpoint config. This is primarily used to validate the configuration. Possible values: ENDPOINT_CONFIG_SELECTOR_TYPE_UNSPECIFIED, SIDECAR_PROXY, GRPC_SERVER
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}

// NetworkServicesEndpointPolicyStatus defines the observed state of NetworkServicesEndpointPolicy.
type NetworkServicesEndpointPolicyStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
