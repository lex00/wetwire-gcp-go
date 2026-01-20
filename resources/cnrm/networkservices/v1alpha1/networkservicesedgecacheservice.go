package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NetworkServicesEdgeCacheService represents a networkservices.cnrm.cloud.google.com NetworkServicesEdgeCacheService resource.
type NetworkServicesEdgeCacheService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NetworkServicesEdgeCacheServiceSpec   `json:"spec,omitempty"`
	Status NetworkServicesEdgeCacheServiceStatus `json:"status,omitempty"`
}

// NetworkServicesEdgeCacheServiceSpec defines the desired state of NetworkServicesEdgeCacheService.
type NetworkServicesEdgeCacheServiceSpec struct {
	// A human-readable description of the resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Disables HTTP/2.
	// HTTP/2 (h2) is enabled by default and recommended for performance. HTTP/2 improves connection re-use and reduces connection setup overhead by sending multiple streams over the same connection.
	// Some legacy HTTP clients may have issues with HTTP/2 connections due to broken HTTP/2 implementations. Setting this to true will prevent HTTP/2 from being advertised and negotiated.
	DisableHttp2 bool `json:"disableHttp2,omitempty" yaml:"disableHttp2,omitempty"`
	// HTTP/3 (IETF QUIC) and Google QUIC are enabled by default.
	DisableQuic bool `json:"disableQuic,omitempty" yaml:"disableQuic,omitempty"`
	// Resource URL that points at the Cloud Armor edge security policy that is applied on each request against the EdgeCacheService.
	EdgeSecurityPolicy string `json:"edgeSecurityPolicy,omitempty" yaml:"edgeSecurityPolicy,omitempty"`
	// URLs to sslCertificate resources that are used to authenticate connections between users and the EdgeCacheService.
	// Note that only "global" certificates with a "scope" of "EDGE_CACHE" can be attached to an EdgeCacheService.
	EdgeSslCertificates []string `json:"edgeSslCertificates,omitempty" yaml:"edgeSslCertificates,omitempty"`
	// Specifies the logging options for the traffic served by this service. If logging is enabled, logs will be exported to Cloud Logging.
	LogConfig map[string]interface{} `json:"logConfig,omitempty" yaml:"logConfig,omitempty"`
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Require TLS (HTTPS) for all clients connecting to this service.
	// Clients who connect over HTTP (port 80) will receive a HTTP 301 to the same URL over HTTPS (port 443).
	// You must have at least one (1) edgeSslCertificate specified to enable this.
	RequireTls bool `json:"requireTls,omitempty" yaml:"requireTls,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Defines how requests are routed, modified, cached and/or which origin content is filled from.
	Routing map[string]interface{} `json:"routing,omitempty" yaml:"routing,omitempty"`
	// URL of the SslPolicy resource that will be associated with the EdgeCacheService.
	// If not set, the EdgeCacheService has no SSL policy configured, and will default to the "COMPATIBLE" policy.
	SslPolicy string `json:"sslPolicy,omitempty" yaml:"sslPolicy,omitempty"`
}

// NetworkServicesEdgeCacheServiceStatus defines the observed state of NetworkServicesEdgeCacheService.
type NetworkServicesEdgeCacheServiceStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
