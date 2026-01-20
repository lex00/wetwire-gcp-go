package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ComputeSSLPolicy represents a compute.cnrm.cloud.google.com ComputeSSLPolicy resource.
type ComputeSSLPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeSSLPolicySpec   `json:"spec,omitempty"`
	Status ComputeSSLPolicyStatus `json:"status,omitempty"`
}

// ComputeSSLPolicySpec defines the desired state of ComputeSSLPolicy.
type ComputeSSLPolicySpec struct {
	// Profile specifies the set of SSL features that can be used by the
	// load balancer when negotiating SSL with clients. This can be one of
	// 'COMPATIBLE', 'MODERN', 'RESTRICTED', or 'CUSTOM'. If using 'CUSTOM',
	// the set of SSL features to enable must be specified in the
	// 'customFeatures' field.
	// See the [official documentation](https://cloud.google.com/compute/docs/load-balancing/ssl-policies#profilefeaturesupport)
	// for which ciphers are available to use. **Note**: this argument
	// *must* be present when using the 'CUSTOM' profile. This argument
	// *must not* be present when using any other profile.
	CustomFeatures []string `json:"customFeatures,omitempty" yaml:"customFeatures,omitempty"`
	// Immutable. An optional description of this resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// The minimum version of SSL protocol that can be used by the clients
	// to establish a connection with the load balancer. Default value: "TLS_1_0" Possible values: ["TLS_1_0", "TLS_1_1", "TLS_1_2"].
	MinTlsVersion string `json:"minTlsVersion,omitempty" yaml:"minTlsVersion,omitempty"`
	// Profile specifies the set of SSL features that can be used by the
	// load balancer when negotiating SSL with clients. If using 'CUSTOM',
	// the set of SSL features to enable must be specified in the
	// 'customFeatures' field.
	// See the [official documentation](https://cloud.google.com/compute/docs/load-balancing/ssl-policies#profilefeaturesupport)
	// for information on what cipher suites each profile provides. If
	// 'CUSTOM' is used, the 'custom_features' attribute **must be set**. Default value: "COMPATIBLE" Possible values: ["COMPATIBLE", "MODERN", "RESTRICTED", "CUSTOM"].
	Profile string `json:"profile,omitempty" yaml:"profile,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// ComputeSSLPolicyStatus defines the observed state of ComputeSSLPolicy.
type ComputeSSLPolicyStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
