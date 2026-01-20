package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ComputeRegionSSLPolicy represents a compute.cnrm.cloud.google.com ComputeRegionSSLPolicy resource.
type ComputeRegionSSLPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeRegionSSLPolicySpec   `json:"spec,omitempty"`
	Status ComputeRegionSSLPolicyStatus `json:"status,omitempty"`
}

// ComputeRegionSSLPolicySpec defines the desired state of ComputeRegionSSLPolicy.
type ComputeRegionSSLPolicySpec struct {
	// A list of features enabled when the selected profile is CUSTOM. The
	// method returns the set of features that can be specified in this
	// list. This field must be empty if the profile is not CUSTOM.
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
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. The region where the regional SSL policy resides.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// ComputeRegionSSLPolicyStatus defines the observed state of ComputeRegionSSLPolicy.
type ComputeRegionSSLPolicyStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
