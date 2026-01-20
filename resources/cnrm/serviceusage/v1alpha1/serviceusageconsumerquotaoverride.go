package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ServiceUsageConsumerQuotaOverride represents a serviceusage.cnrm.cloud.google.com ServiceUsageConsumerQuotaOverride resource.
type ServiceUsageConsumerQuotaOverride struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ServiceUsageConsumerQuotaOverrideSpec   `json:"spec,omitempty"`
	Status ServiceUsageConsumerQuotaOverrideStatus `json:"status,omitempty"`
}

// ServiceUsageConsumerQuotaOverrideSpec defines the desired state of ServiceUsageConsumerQuotaOverride.
type ServiceUsageConsumerQuotaOverrideSpec struct {
	// Immutable. If this map is nonempty, then this override applies only to specific values for dimensions defined in the limit unit.
	Dimensions map[string]string `json:"dimensions,omitempty" yaml:"dimensions,omitempty"`
	// If the new quota would decrease the existing quota by more than 10%, the request is rejected.
	// If 'force' is 'true', that safety check is ignored.
	Force bool `json:"force,omitempty" yaml:"force,omitempty"`
	// Immutable. The limit on the metric, e.g. '/project/region'.
	// ~> Make sure that 'limit' is in a format that doesn't start with '1/' or contain curly braces.
	// E.g. use '/project/user' instead of '1/{project}/{user}'.
	Limit string `json:"limit,omitempty" yaml:"limit,omitempty"`
	// Immutable. The metric that should be limited, e.g. 'compute.googleapis.com/cpus'.
	Metric string `json:"metric,omitempty" yaml:"metric,omitempty"`
	// The overriding quota limit value. Can be any nonnegative integer, or -1 (unlimited quota).
	OverrideValue string `json:"overrideValue,omitempty" yaml:"overrideValue,omitempty"`
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Immutable. The service that the metrics belong to, e.g. 'compute.googleapis.com'.
	Service string `json:"service,omitempty" yaml:"service,omitempty"`
}

// ServiceUsageConsumerQuotaOverrideStatus defines the observed state of ServiceUsageConsumerQuotaOverride.
type ServiceUsageConsumerQuotaOverrideStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
