package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ComputeTargetHTTPSProxy represents a compute.cnrm.cloud.google.com ComputeTargetHTTPSProxy resource.
type ComputeTargetHTTPSProxy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeTargetHTTPSProxySpec   `json:"spec,omitempty"`
	Status ComputeTargetHTTPSProxyStatus `json:"status,omitempty"`
}

// ComputeTargetHTTPSProxySpec defines the desired state of ComputeTargetHTTPSProxy.
type ComputeTargetHTTPSProxySpec struct {
	CertificateManagerCertificates []map[string]interface{} `json:"certificateManagerCertificates,omitempty" yaml:"certificateManagerCertificates,omitempty"`
	// A reference to the CertificateMap resource uri that identifies a
	// certificate map associated with the given target proxy. This field
	// can only be set for global target proxies. This field is only supported
	// for EXTERNAL and EXTERNAL_MANAGED load balancing schemes.
	// For INTERNAL_MANAGED, use certificateManagerCertificates instead.
	// sslCertificates and certificateMap fields cannot be defined together.
	CertificateMapRef map[string]interface{} `json:"certificateMapRef,omitempty" yaml:"certificateMapRef,omitempty"`
	// Immutable. An optional description of this resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Immutable. Specifies how long to keep a connection open, after completing a response,
	// while there is no matching traffic (in seconds). If an HTTP keepalive is
	// not specified, a default value (610 seconds) will be used. For Global
	// external HTTP(S) load balancer, the minimum allowed value is 5 seconds and
	// the maximum allowed value is 1200 seconds. For Global external HTTP(S)
	// load balancer (classic), this option is not available publicly.
	HTTPKeepAliveTimeoutSec int32 `json:"httpKeepAliveTimeoutSec,omitempty" yaml:"httpKeepAliveTimeoutSec,omitempty"`
	// Location represents the geographical location of the ComputeTargetHTTPSProxy. Specify a region name or "global" for global resources. Reference: GCP definition of regions/zones (https://cloud.google.com/compute/docs/regions-zones/)
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Immutable. This field only applies when the forwarding rule that references
	// this target proxy has a loadBalancingScheme set to INTERNAL_SELF_MANAGED.
	ProxyBind bool `json:"proxyBind,omitempty" yaml:"proxyBind,omitempty"`
	// Specifies the QUIC override policy for this resource. This determines
	// whether the load balancer will attempt to negotiate QUIC with clients
	// or not. Can specify one of NONE, ENABLE, or DISABLE. If NONE is
	// specified, Google manages whether QUIC is used. Default value: "NONE" Possible values: ["NONE", "ENABLE", "DISABLE"].
	QuicOverride string `json:"quicOverride,omitempty" yaml:"quicOverride,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Immutable. A URL referring to a networksecurity.ServerTlsPolicy
	// resource that describes how the proxy should authenticate inbound
	// traffic. serverTlsPolicy only applies to a global TargetHttpsProxy
	// attached to globalForwardingRules with the loadBalancingScheme
	// set to INTERNAL_SELF_MANAGED or EXTERNAL or EXTERNAL_MANAGED.
	// For details which ServerTlsPolicy resources are accepted with
	// INTERNAL_SELF_MANAGED and which with EXTERNAL, EXTERNAL_MANAGED
	// loadBalancingScheme consult ServerTlsPolicy documentation.
	// If left blank, communications are not encrypted.
	ServerTlsPolicyRef map[string]interface{} `json:"serverTlsPolicyRef,omitempty" yaml:"serverTlsPolicyRef,omitempty"`
	SslCertificates []map[string]interface{} `json:"sslCertificates,omitempty" yaml:"sslCertificates,omitempty"`
	// A reference to the ComputeSSLPolicy resource that will be
	// associated with the ComputeTargetHTTPSProxy resource. If not set,
	// the ComputeTargetHTTPSProxy resource will not have any SSL policy
	// configured.
	SslPolicyRef map[string]interface{} `json:"sslPolicyRef,omitempty" yaml:"sslPolicyRef,omitempty"`
	// A reference to the ComputeURLMap resource that defines the mapping
	// from URL to the BackendService.
	URLMapRef map[string]interface{} `json:"urlMapRef,omitempty" yaml:"urlMapRef,omitempty"`
}

// ComputeTargetHTTPSProxyStatus defines the observed state of ComputeTargetHTTPSProxy.
type ComputeTargetHTTPSProxyStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
