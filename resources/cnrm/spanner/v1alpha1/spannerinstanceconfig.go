package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// SpannerInstanceConfig is the Schema for the SpannerInstanceConfig API
type SpannerInstanceConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SpannerInstanceConfigSpec   `json:"spec,omitempty"`
	Status SpannerInstanceConfigStatus `json:"status,omitempty"`
}

// SpannerInstanceConfigSpec defines the desired state of SpannerInstanceConfig.
type SpannerInstanceConfigSpec struct {
	// Base configuration name, e.g. projects/<project_name>/instanceConfigs/nam3, based on which this configuration is created. Only set for user-managed configurations. `base_config` must refer to a configuration of type `GOOGLE_MANAGED` in the same project as this configuration.
	BaseConfigRef map[string]interface{} `json:"baseConfigRef,omitempty" yaml:"baseConfigRef,omitempty"`
	// The name of this instance configuration as it appears in UIs.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// Cloud Labels are a flexible and lightweight mechanism for organizing cloud
	// resources into groups that reflect a customer's organizational needs and
	// deployment strategies. Cloud Labels can be used to filter collections of
	// resources. They can be used to control how resource metrics are aggregated.
	// And they can be used as arguments to policy management rules (e.g. route,
	// firewall, load balancing, etc.).
	// * Label keys must be between 1 and 63 characters long and must conform to
	// the following regular expression: `[a-z][a-z0-9_-]{0,62}`.
	// * Label values must be between 0 and 63 characters long and must conform
	// to the regular expression: `[a-z0-9_-]{0,63}`.
	// * No more than 64 labels can be associated with a given resource.
	// See https://goo.gl/xmQnxf for more information on and examples of labels.
	// If you plan to use labels in your own code, please note that additional
	// characters may be allowed in the future. Therefore, you are advised to use
	// an internal label representation, such as JSON, which doesn't rely upon
	// specific characters being disallowed.  For example, representing labels
	// as the string:  name + "_" + value  would prove problematic if we were to
	// allow "_" in a future release.
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	// Allowed values of the "default_leader" schema option for databases in instances that use this instance configuration.
	LeaderOptions []string `json:"leaderOptions,omitempty" yaml:"leaderOptions,omitempty"`
	// The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// The geographic placement of nodes in this instance configuration and their
	// replication properties.
	// To create user-managed configurations, input
	// `replicas` must include all replicas in `replicas` of the `base_config`
	// and include one or more replicas in the `optional_replicas` of the
	// `base_config`.
	Replicas []map[string]interface{} `json:"replicas,omitempty" yaml:"replicas,omitempty"`
	// The SpannerInstanceConfig name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// SpannerInstanceConfigStatus defines the observed state of SpannerInstanceConfig.
type SpannerInstanceConfigStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
