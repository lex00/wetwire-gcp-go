package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ComputeTargetTCPProxy is the Schema for the ComputeTargetTCPProxy API
type ComputeTargetTCPProxy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeTargetTCPProxySpec   `json:"spec,omitempty"`
	Status ComputeTargetTCPProxyStatus `json:"status,omitempty"`
}

// ComputeTargetTCPProxySpec defines the desired state of ComputeTargetTCPProxy.
type ComputeTargetTCPProxySpec struct {
	// A reference to the ComputeBackendService resource.
	BackendServiceRef map[string]interface{} `json:"backendServiceRef,omitempty" yaml:"backendServiceRef,omitempty"`
	// Immutable. An optional description of this resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// The geographical location of the ComputeTargetTCPProxy. Reference: GCP definition of regions/zones (https://cloud.google.com/compute/docs/regions-zones/)
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Immutable. This field only applies when the forwarding rule that references this target proxy has a loadBalancingScheme set to INTERNAL_SELF_MANAGED.
	ProxyBind bool `json:"proxyBind,omitempty" yaml:"proxyBind,omitempty"`
	// Specifies the type of proxy header to append before sending data to the backend. Default value: "NONE" Possible values: ["NONE", "PROXY_V1"].
	ProxyHeader string `json:"proxyHeader,omitempty" yaml:"proxyHeader,omitempty"`
	// The ComputeTargetTCPProxy name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// ComputeTargetTCPProxyStatus defines the observed state of ComputeTargetTCPProxy.
type ComputeTargetTCPProxyStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
