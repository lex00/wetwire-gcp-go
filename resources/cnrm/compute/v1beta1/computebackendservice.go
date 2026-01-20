package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ComputeBackendService represents a compute.cnrm.cloud.google.com ComputeBackendService resource.
type ComputeBackendService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeBackendServiceSpec   `json:"spec,omitempty"`
	Status ComputeBackendServiceStatus `json:"status,omitempty"`
}

// ComputeBackendServiceSpec defines the desired state of ComputeBackendService.
type ComputeBackendServiceSpec struct {
	// Lifetime of cookies in seconds if session_affinity is
	// GENERATED_COOKIE. If set to 0, the cookie is non-persistent and lasts
	// only until the end of the browser session (or equivalent). The
	// maximum allowed value for TTL is one day.
	// When the load balancing scheme is INTERNAL, this field is not used.
	AffinityCookieTtlSec int32 `json:"affinityCookieTtlSec,omitempty" yaml:"affinityCookieTtlSec,omitempty"`
	// The set of backends that serve this BackendService.
	Backend []map[string]interface{} `json:"backend,omitempty" yaml:"backend,omitempty"`
	// Cloud CDN configuration for this BackendService.
	CdnPolicy map[string]interface{} `json:"cdnPolicy,omitempty" yaml:"cdnPolicy,omitempty"`
	// Settings controlling the volume of connections to a backend service. This field
	// is applicable only when the load_balancing_scheme is set to INTERNAL_SELF_MANAGED.
	CircuitBreakers map[string]interface{} `json:"circuitBreakers,omitempty" yaml:"circuitBreakers,omitempty"`
	// Compress text responses using Brotli or gzip compression, based on the client's Accept-Encoding header. Possible values: ["AUTOMATIC", "DISABLED"].
	CompressionMode string `json:"compressionMode,omitempty" yaml:"compressionMode,omitempty"`
	// Time for which instance will be drained (not accept new
	// connections, but still work to finish started).
	ConnectionDrainingTimeoutSec int32 `json:"connectionDrainingTimeoutSec,omitempty" yaml:"connectionDrainingTimeoutSec,omitempty"`
	// Connection Tracking configuration for this BackendService.
	// This is available only for Layer 4 Internal Load Balancing and
	// Network Load Balancing.
	ConnectionTrackingPolicy map[string]interface{} `json:"connectionTrackingPolicy,omitempty" yaml:"connectionTrackingPolicy,omitempty"`
	// Consistent Hash-based load balancing can be used to provide soft session
	// affinity based on HTTP headers, cookies or other properties. This load balancing
	// policy is applicable only for HTTP connections. The affinity to a particular
	// destination host will be lost when one or more hosts are added/removed from the
	// destination service. This field specifies parameters that control consistent
	// hashing. This field only applies if the load_balancing_scheme is set to
	// INTERNAL_SELF_MANAGED. This field is only applicable when locality_lb_policy is
	// set to MAGLEV or RING_HASH.
	ConsistentHash map[string]interface{} `json:"consistentHash,omitempty" yaml:"consistentHash,omitempty"`
	// Headers that the HTTP/S load balancer should add to proxied
	// requests.
	CustomRequestHeaders []string `json:"customRequestHeaders,omitempty" yaml:"customRequestHeaders,omitempty"`
	// Headers that the HTTP/S load balancer should add to proxied
	// responses.
	CustomResponseHeaders []string `json:"customResponseHeaders,omitempty" yaml:"customResponseHeaders,omitempty"`
	// An optional description of this resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// The resource URL for the edge security policy associated with this
	// backend service.
	EdgeSecurityPolicyRef map[string]interface{} `json:"edgeSecurityPolicyRef,omitempty" yaml:"edgeSecurityPolicyRef,omitempty"`
	// If true, enable Cloud CDN for this BackendService.
	EnableCdn bool `json:"enableCdn,omitempty" yaml:"enableCdn,omitempty"`
	// Policy for failovers.
	FailoverPolicy map[string]interface{} `json:"failoverPolicy,omitempty" yaml:"failoverPolicy,omitempty"`
	HealthChecks []map[string]interface{} `json:"healthChecks,omitempty" yaml:"healthChecks,omitempty"`
	// Settings for enabling Cloud Identity Aware Proxy.
	Iap map[string]interface{} `json:"iap,omitempty" yaml:"iap,omitempty"`
	// Immutable. Indicates whether the backend service will be used with internal or
	// external load balancing. A backend service created for one type of
	// load balancing cannot be used with the other. For more information, refer to
	// [Choosing a load balancer](https://cloud.google.com/load-balancing/docs/backend-service). Default value: "EXTERNAL" Possible values: ["EXTERNAL", "INTERNAL_SELF_MANAGED", "INTERNAL_MANAGED", "EXTERNAL_MANAGED"].
	LoadBalancingScheme string `json:"loadBalancingScheme,omitempty" yaml:"loadBalancingScheme,omitempty"`
	// A list of locality load balancing policies to be used in order of
	// preference. Either the policy or the customPolicy field should be set.
	// Overrides any value set in the localityLbPolicy field.
	// localityLbPolicies is only supported when the BackendService is referenced
	// by a URL Map that is referenced by a target gRPC proxy that has the
	// validateForProxyless field set to true.
	LocalityLbPolicies []map[string]interface{} `json:"localityLbPolicies,omitempty" yaml:"localityLbPolicies,omitempty"`
	// The load balancing algorithm used within the scope of the locality.
	// The possible values are:
	// * 'ROUND_ROBIN': This is a simple policy in which each healthy backend
	// is selected in round robin order.
	// * 'LEAST_REQUEST': An O(1) algorithm which selects two random healthy
	// hosts and picks the host which has fewer active requests.
	// * 'RING_HASH': The ring/modulo hash load balancer implements consistent
	// hashing to backends. The algorithm has the property that the
	// addition/removal of a host from a set of N hosts only affects
	// 1/N of the requests.
	// * 'RANDOM': The load balancer selects a random healthy host.
	// * 'ORIGINAL_DESTINATION': Backend host is selected based on the client
	// connection metadata, i.e., connections are opened
	// to the same address as the destination address of
	// the incoming connection before the connection
	// was redirected to the load balancer.
	// * 'MAGLEV': used as a drop in replacement for the ring hash load balancer.
	// Maglev is not as stable as ring hash but has faster table lookup
	// build times and host selection times. For more information about
	// Maglev, refer to https://ai.google/research/pubs/pub44824
	// * 'WEIGHTED_MAGLEV': Per-instance weighted Load Balancing via health check
	// reported weights. If set, the Backend Service must
	// configure a non legacy HTTP-based Health Check, and
	// health check replies are expected to contain
	// non-standard HTTP response header field
	// X-Load-Balancing-Endpoint-Weight to specify the
	// per-instance weights. If set, Load Balancing is weight
	// based on the per-instance weights reported in the last
	// processed health check replies, as long as every
	// instance either reported a valid weight or had
	// UNAVAILABLE_WEIGHT. Otherwise, Load Balancing remains
	// equal-weight.
	// This field is applicable to either:
	// * A regional backend service with the service_protocol set to HTTP, HTTPS, or HTTP2,
	// and loadBalancingScheme set to INTERNAL_MANAGED.
	// * A global backend service with the load_balancing_scheme set to INTERNAL_SELF_MANAGED.
	// * A regional backend service with loadBalancingScheme set to EXTERNAL (External Network
	// Load Balancing). Only MAGLEV and WEIGHTED_MAGLEV values are possible for External
	// Network Load Balancing. The default is MAGLEV.
	// If session_affinity is not NONE, and this field is not set to MAGLEV, WEIGHTED_MAGLEV,
	// or RING_HASH, session affinity settings will not take effect.
	// Only ROUND_ROBIN and RING_HASH are supported when the backend service is referenced
	// by a URL map that is bound to target gRPC proxy that has validate_for_proxyless
	// field set to true. Possible values: ["ROUND_ROBIN", "LEAST_REQUEST", "RING_HASH", "RANDOM", "ORIGINAL_DESTINATION", "MAGLEV", "WEIGHTED_MAGLEV"].
	LocalityLbPolicy string `json:"localityLbPolicy,omitempty" yaml:"localityLbPolicy,omitempty"`
	// Location represents the geographical location of the ComputeBackendService. Specify a region name or "global" for global resources. Reference: GCP definition of regions/zones (https://cloud.google.com/compute/docs/regions-zones/)
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// This field denotes the logging options for the load balancer traffic served by this backend service.
	// If logging is enabled, logs will be exported to Stackdriver.
	LogConfig map[string]interface{} `json:"logConfig,omitempty" yaml:"logConfig,omitempty"`
	// The network to which this backend service belongs.  This field can
	// only be specified when the load balancing scheme is set to
	// INTERNAL.
	NetworkRef map[string]interface{} `json:"networkRef,omitempty" yaml:"networkRef,omitempty"`
	// Settings controlling eviction of unhealthy hosts from the load balancing pool.
	// This field is applicable only when the load_balancing_scheme is set
	// to INTERNAL_SELF_MANAGED.
	OutlierDetection map[string]interface{} `json:"outlierDetection,omitempty" yaml:"outlierDetection,omitempty"`
	// Name of backend port. The same name should appear in the instance
	// groups referenced by this service. Required when the load balancing
	// scheme is EXTERNAL.
	PortName string `json:"portName,omitempty" yaml:"portName,omitempty"`
	// The protocol this BackendService uses to communicate with backends.
	// The default is HTTP. **NOTE**: HTTP2 is only valid for beta HTTP/2 load balancer
	// types and may result in errors if used with the GA API. **NOTE**: With protocol “UNSPECIFIED”,
	// the backend service can be used by Layer 4 Internal Load Balancing or Network Load Balancing
	// with TCP/UDP/L3_DEFAULT Forwarding Rule protocol. Possible values: ["HTTP", "HTTPS", "HTTP2", "TCP", "SSL", "GRPC", "UNSPECIFIED"].
	Protocol string `json:"protocol,omitempty" yaml:"protocol,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// The security policy associated with this backend service.
	SecurityPolicy string `json:"securityPolicy,omitempty" yaml:"securityPolicy,omitempty"`
	// The security policy associated with this backend service.
	SecurityPolicyRef map[string]interface{} `json:"securityPolicyRef,omitempty" yaml:"securityPolicyRef,omitempty"`
	// The security settings that apply to this backend service. This field is applicable to either
	// a regional backend service with the service_protocol set to HTTP, HTTPS, or HTTP2, and
	// load_balancing_scheme set to INTERNAL_MANAGED; or a global backend service with the
	// load_balancing_scheme set to INTERNAL_SELF_MANAGED.
	SecuritySettings map[string]interface{} `json:"securitySettings,omitempty" yaml:"securitySettings,omitempty"`
	// Type of session affinity to use. The default is NONE. Session affinity is
	// not applicable if the protocol is UDP. Possible values: ["NONE", "CLIENT_IP", "CLIENT_IP_PORT_PROTO", "CLIENT_IP_PROTO", "GENERATED_COOKIE", "HEADER_FIELD", "HTTP_COOKIE"].
	SessionAffinity string `json:"sessionAffinity,omitempty" yaml:"sessionAffinity,omitempty"`
	// Subsetting configuration for this BackendService. Currently this is applicable only for Internal TCP/UDP load balancing and Internal HTTP(S) load balancing.
	Subsetting map[string]interface{} `json:"subsetting,omitempty" yaml:"subsetting,omitempty"`
	// How many seconds to wait for the backend before considering it a
	// failed request. Default is 30 seconds. Valid range is [1, 86400].
	TimeoutSec int32 `json:"timeoutSec,omitempty" yaml:"timeoutSec,omitempty"`
}

// ComputeBackendServiceStatus defines the observed state of ComputeBackendService.
type ComputeBackendServiceStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
