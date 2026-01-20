package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// SpannerInstance is the Schema for the SpannerInstance API
type SpannerInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SpannerInstanceSpec   `json:"spec,omitempty"`
	Status SpannerInstanceStatus `json:"status,omitempty"`
}

// SpannerInstanceSpec defines the desired state of SpannerInstance.
type SpannerInstanceSpec struct {
	// Optional. The autoscaling configuration. Autoscaling is enabled if this field is set. When autoscaling is enabled, node_count and processing_units are treated as OUTPUT_ONLY fields and reflect the current compute capacity allocated to the instance.
	AutoscalingConfig map[string]interface{} `json:"autoscalingConfig,omitempty" yaml:"autoscalingConfig,omitempty"`
	// Immutable. The name of the instance's configuration (similar but not quite the same as a region) which defines the geographic placement and replication of your databases in this instance. It determines where your data is stored. Values are typically of the form 'regional-europe-west1' , 'us-central' etc. In order to obtain a valid list please consult the [Configuration section of the docs](https://cloud.google.com/spanner/docs/instances).
	Config string `json:"config,omitempty" yaml:"config,omitempty"`
	// Optional. Controls the default backup schedule behavior for new databases
	// within the instance. By default, a backup schedule is created automatically
	// when a new database is created in a new instance.
	// Note that the `AUTOMATIC` value isn't permitted for free instances,
	// as backups and backup schedules aren't supported for free instances.
	// In the `GetInstance` or `ListInstances` response, if the value of
	// `default_backup_schedule_type` isn't set, or set to `NONE`, Spanner doesn't
	// create a default backup schedule for new databases in the instance.
	DefaultBackupScheduleType string `json:"defaultBackupScheduleType,omitempty" yaml:"defaultBackupScheduleType,omitempty"`
	// The descriptive name for this instance as it appears in UIs. Must be unique per project and between 4 and 30 characters in length.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// Optional. The `Edition` of the current instance. Currently accepted values are STANDARD, ENTERPRISE, ENTERPRISE_PLUS. If edition is unspecified, it has automatically upgraded to the lowest edition that matches your usage pattern.
	Edition string `json:"edition,omitempty" yaml:"edition,omitempty"`
	// Cloud Labels are a flexible and lightweight mechanism for organizing cloud
	// resources into groups that reflect a customer's organizational needs and
	// deployment strategies. Cloud Labels can be used to filter collections of
	// resources. They can be used to control how resource metrics are aggregated.
	// And they can be used as arguments to policy management rules (e.g. route,
	// firewall, load balancing, etc.).
	// - Label keys must be between 1 and 63 characters long and must conform to
	// the following regular expression: `[a-z][a-z0-9_-]{0,62}`.
	// - Label values must be between 0 and 63 characters long and must conform
	// to the regular expression `[a-z0-9_-]{0,63}`.
	// - No more than 64 labels can be associated with a given resource.
	// See https://goo.gl/xmQnxf for more information on and examples of labels.
	// If you plan to use labels in your own code, please note that additional
	// characters may be allowed in the future. And so you are advised to use an
	// internal label representation, such as JSON, which doesn't rely upon
	// specific characters being disallowed.  For example, representing labels
	// as the string:  name + "_" + value  would prove problematic if we were to
	// allow "_" in a future release.
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	NumNodes int32 `json:"numNodes,omitempty" yaml:"numNodes,omitempty"`
	ProcessingUnits int32 `json:"processingUnits,omitempty" yaml:"processingUnits,omitempty"`
	// The SpannerInstance name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// SpannerInstanceStatus defines the observed state of SpannerInstance.
type SpannerInstanceStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
