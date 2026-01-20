package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// LoggingLogMetric is the Schema for the logging API
type LoggingLogMetric struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LoggingLogMetricSpec   `json:"spec,omitempty"`
	Status LoggingLogMetricStatus `json:"status,omitempty"`
}

// LoggingLogMetricSpec defines the desired state of LoggingLogMetric.
type LoggingLogMetricSpec struct {
	// Optional. The `bucket_options` are required when the logs-based metric is using a DISTRIBUTION value type and it describes the bucket boundaries used to create a histogram of the extracted values.
	BucketOptions map[string]interface{} `json:"bucketOptions,omitempty" yaml:"bucketOptions,omitempty"`
	// Optional. A description of this metric, which is used in documentation. The maximum length of the description is 8000 characters.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Optional. If set to True, then this metric is disabled and it does not generate any points.
	Disabled bool `json:"disabled,omitempty" yaml:"disabled,omitempty"`
	// Required. An [advanced logs filter](https://cloud.google.com/logging/docs/view/advanced_filters) which is used to match log entries. Example: "resource.type=gae_app AND severity>=ERROR" The maximum length of the filter is 20000 characters.
	Filter string `json:"filter,omitempty" yaml:"filter,omitempty"`
	// Optional. A map from a label key string to an extractor expression which is used to extract data from a log entry field and assign as the label value. Each label key specified in the LabelDescriptor must have an associated extractor expression in this map. The syntax of the extractor expression is the same as for the `value_extractor` field. The extracted value is converted to the type defined in the label descriptor. If the either the extraction or the type conversion fails, the label will have a default value. The default value for a string label is an empty string, for an integer label its 0, and for a boolean label its `false`. Note that there are upper bounds on the maximum number of labels and the number of active time series that are allowed in a project.
	LabelExtractors map[string]string `json:"labelExtractors,omitempty" yaml:"labelExtractors,omitempty"`
	// The reference to the Log Bucket that owns the Log Metric. Only Log Buckets in projects are supported. The bucket has to be in the same project as the metric. For example:projects/my-project/locations/global/buckets/my-bucket If empty, then the Log Metric is considered a non-Bucket Log Metric. Only `external` field is supported to configure the reference for now.
	LoggingLogBucketRef map[string]interface{} `json:"loggingLogBucketRef,omitempty" yaml:"loggingLogBucketRef,omitempty"`
	// Optional. The metric descriptor associated with the logs-based metric. If unspecified, it uses a default metric descriptor with a DELTA metric kind, INT64 value type, with no labels and a unit of "1". Such a metric counts the number of log entries matching the `filter` expression. The `name`, `type`, and `description` fields in the `metric_descriptor` are output only, and is constructed using the `name` and `description` field in the LogMetric. To create a logs-based metric that records a distribution of log values, a DELTA metric kind with a DISTRIBUTION value type must be used along with a `value_extractor` expression in the LogMetric. Each label in the metric descriptor must have a matching label name as the key and an extractor expression as the value in the `label_extractors` map. The `metric_kind` and `value_type` fields in the `metric_descriptor` cannot be updated once initially configured. New labels can be added in the `metric_descriptor`, but existing labels cannot be modified except for their description.
	MetricDescriptor map[string]interface{} `json:"metricDescriptor,omitempty" yaml:"metricDescriptor,omitempty"`
	// Immutable. The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Optional. A `value_extractor` is required when using a distribution logs-based metric to extract the values to record from a log entry. Two functions are supported for value extraction: `EXTRACT(field)` or `REGEXP_EXTRACT(field, regex)`. The argument are: 1. field: The name of the log entry field from which the value is to be extracted. 2. regex: A regular expression using the Google RE2 syntax (https://github.com/google/re2/wiki/Syntax) with a single capture group to extract data from the specified log entry field. The value of the field is converted to a string before applying the regex. It is an error to specify a regex that does not include exactly one capture group. The result of the extraction must be convertible to a double type, as the distribution always records double values. If either the extraction or the conversion to double fails, then those values are not recorded in the distribution. Example: `REGEXP_EXTRACT(jsonPayload.request, ".*quantity=(d+).*")`
	ValueExtractor string `json:"valueExtractor,omitempty" yaml:"valueExtractor,omitempty"`
}

// LoggingLogMetricStatus defines the observed state of LoggingLogMetric.
type LoggingLogMetricStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
