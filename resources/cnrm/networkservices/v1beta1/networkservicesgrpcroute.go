package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NetworkServicesGRPCRoute represents a networkservices.cnrm.cloud.google.com NetworkServicesGRPCRoute resource.
type NetworkServicesGRPCRoute struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NetworkServicesGRPCRouteSpec   `json:"spec,omitempty"`
	Status NetworkServicesGRPCRouteStatus `json:"status,omitempty"`
}

// NetworkServicesGRPCRouteSpec defines the desired state of NetworkServicesGRPCRoute.
type NetworkServicesGRPCRouteSpec struct {
	// Optional. A free-text description of the resource. Max length 1024 characters.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	Gateways []map[string]interface{} `json:"gateways,omitempty" yaml:"gateways,omitempty"`
	// Required. Service hostnames with an optional port for which this route describes traffic. Format: [:] Hostname is the fully qualified domain name of a network host. This matches the RFC 1123 definition of a hostname with 2 notable exceptions: - IPs are not allowed. - A hostname may be prefixed with a wildcard label (*.). The wildcard label must appear by itself as the first label. Hostname can be “precise” which is a domain name without the terminating dot of a network host (e.g. “foo.example.com”) or “wildcard”, which is a domain name prefixed with a single wildcard label (e.g. *.example.com). Note that as per RFC1035 and RFC1123, a label must consist of lower case alphanumeric characters or ‘-’, and must start and end with an alphanumeric character. No other punctuation is allowed. The routes associated with a Router must have unique hostnames. If you attempt to attach multiple routes with conflicting hostnames, the configuration will be rejected. For example, while it is acceptable for routes for the hostnames "*.foo.bar.com" and "*.bar.com" to be associated with the same route, it is not possible to associate two routes both with "*.bar.com" or both with "bar.com". In the case that multiple routes match the hostname, the most specific match will be selected. For example, "foo.bar.baz.com" will take precedence over "*.bar.baz.com" and "*.bar.baz.com" will take precedence over "*.baz.com". If a port is specified, then gRPC clients must use the channel URI with the port to match this rule (i.e. "xds:///service:123"), otherwise they must supply the URI without a port (i.e. "xds:///service").
	Hostnames []string `json:"hostnames,omitempty" yaml:"hostnames,omitempty"`
	// Immutable. The location for the resource
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	Meshes []map[string]interface{} `json:"meshes,omitempty" yaml:"meshes,omitempty"`
	// Immutable. The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Required. A list of detailed rules defining how to route traffic. Within a single GrpcRoute, the GrpcRoute.RouteAction associated with the first matching GrpcRoute.RouteRule will be executed. At least one rule must be supplied.
	Rules []map[string]interface{} `json:"rules,omitempty" yaml:"rules,omitempty"`
}

// NetworkServicesGRPCRouteStatus defines the observed state of NetworkServicesGRPCRoute.
type NetworkServicesGRPCRouteStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
