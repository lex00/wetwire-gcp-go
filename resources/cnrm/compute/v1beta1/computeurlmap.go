package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ComputeURLMap represents a compute.cnrm.cloud.google.com ComputeURLMap resource.
type ComputeURLMap struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeURLMapSpec   `json:"spec,omitempty"`
	Status ComputeURLMapStatus `json:"status,omitempty"`
}

// ComputeURLMapSpec defines the desired state of ComputeURLMap.
type ComputeURLMapSpec struct {
	// defaultRouteAction takes effect when none of the hostRules match. The load balancer performs advanced routing actions, such as URL rewrites and header transformations, before forwarding the request to the selected backend. If defaultRouteAction specifies any weightedBackendServices, defaultService must not be set. Conversely if defaultService is set, defaultRouteAction cannot contain any weightedBackendServices.
	// Only one of defaultRouteAction or defaultUrlRedirect must be set.
	// URL maps for Classic external HTTP(S) load balancers only support the urlRewrite action within defaultRouteAction.
	// defaultRouteAction has no effect when the URL map is bound to a target gRPC proxy that has the validateForProxyless field set to true.
	DefaultRouteAction map[string]interface{} `json:"defaultRouteAction,omitempty" yaml:"defaultRouteAction,omitempty"`
	// The defaultService resource to which traffic is directed if none of
	// the hostRules match.
	// For the Global URL Map, it should be a reference to the backend
	// service or backend bucket.
	// For the Regional URL Map, it should be a reference to the backend
	// service.
	// If defaultRouteAction is additionally specified, advanced routing
	// actions like URL Rewrites, etc. take effect prior to sending the
	// request to the backend. However, if defaultService is specified,
	// defaultRouteAction cannot contain any weightedBackendServices.
	// Conversely, if routeAction specifies any weightedBackendServices,
	// service must not be specified. Only one of defaultService,
	// defaultUrlRedirect or defaultRouteAction.weightedBackendService
	// must be set.
	DefaultService map[string]interface{} `json:"defaultService,omitempty" yaml:"defaultService,omitempty"`
	// When none of the specified hostRules match, the request is redirected to a URL specified
	// by defaultUrlRedirect. If defaultUrlRedirect is specified, defaultService or
	// defaultRouteAction must not be set.
	DefaultURLRedirect map[string]interface{} `json:"defaultUrlRedirect,omitempty" yaml:"defaultUrlRedirect,omitempty"`
	// An optional description of this resource. Provide this property when
	// you create the resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Specifies changes to request and response headers that need to take effect for
	// the selected backendService. The headerAction specified here take effect after
	// headerAction specified under pathMatcher.
	HeaderAction map[string]interface{} `json:"headerAction,omitempty" yaml:"headerAction,omitempty"`
	// The list of HostRules to use against the URL.
	HostRule []map[string]interface{} `json:"hostRule,omitempty" yaml:"hostRule,omitempty"`
	// Location represents the geographical location of the ComputeURLMap. Specify a region name or "global" for global resources. Reference: GCP definition of regions/zones (https://cloud.google.com/compute/docs/regions-zones/)
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// The list of named PathMatchers to use against the URL.
	PathMatcher []map[string]interface{} `json:"pathMatcher,omitempty" yaml:"pathMatcher,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// The list of expected URL mappings. Requests to update this UrlMap will
	// succeed only if all of the test cases pass.
	Test []map[string]interface{} `json:"test,omitempty" yaml:"test,omitempty"`
}

// ComputeURLMapStatus defines the observed state of ComputeURLMap.
type ComputeURLMapStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
