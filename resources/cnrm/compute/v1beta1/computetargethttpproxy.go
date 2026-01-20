package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ComputeTargetHTTPProxy represents a compute.cnrm.cloud.google.com ComputeTargetHTTPProxy resource.
type ComputeTargetHTTPProxy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeTargetHTTPProxySpec   `json:"spec,omitempty"`
	Status ComputeTargetHTTPProxyStatus `json:"status,omitempty"`
}

// ComputeTargetHTTPProxySpec defines the desired state of ComputeTargetHTTPProxy.
type ComputeTargetHTTPProxySpec struct {
	// Immutable. An optional description of this resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Immutable. Specifies how long to keep a connection open, after completing a response,
	// while there is no matching traffic (in seconds). If an HTTP keepalive is
	// not specified, a default value (610 seconds) will be used. For Global
	// external HTTP(S) load balancer, the minimum allowed value is 5 seconds and
	// the maximum allowed value is 1200 seconds. For Global external HTTP(S)
	// load balancer (classic), this option is not available publicly.
	HTTPKeepAliveTimeoutSec int32 `json:"httpKeepAliveTimeoutSec,omitempty" yaml:"httpKeepAliveTimeoutSec,omitempty"`
	// Location represents the geographical location of the ComputeTargetHTTPProxy. Specify a region name or "global" for global resources. Reference: GCP definition of regions/zones (https://cloud.google.com/compute/docs/regions-zones/)
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Immutable. This field only applies when the forwarding rule that references
	// this target proxy has a loadBalancingScheme set to INTERNAL_SELF_MANAGED.
	ProxyBind bool `json:"proxyBind,omitempty" yaml:"proxyBind,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// A reference to the ComputeURLMap resource that defines the mapping
	// from URL to the BackendService.
	URLMapRef map[string]interface{} `json:"urlMapRef,omitempty" yaml:"urlMapRef,omitempty"`
}

// ComputeTargetHTTPProxyStatus defines the observed state of ComputeTargetHTTPProxy.
type ComputeTargetHTTPProxyStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
