package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CloudSchedulerJob represents a cloudscheduler.cnrm.cloud.google.com CloudSchedulerJob resource.
type CloudSchedulerJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CloudSchedulerJobSpec   `json:"spec,omitempty"`
	Status CloudSchedulerJobStatus `json:"status,omitempty"`
}

// CloudSchedulerJobSpec defines the desired state of CloudSchedulerJob.
type CloudSchedulerJobSpec struct {
	// App Engine HTTP target.
	AppEngineHTTPTarget map[string]interface{} `json:"appEngineHttpTarget,omitempty" yaml:"appEngineHttpTarget,omitempty"`
	// The deadline for job attempts. If the request handler does not respond by this deadline then the request is cancelled and the attempt is marked as a `DEADLINE_EXCEEDED` failure. The failed attempt can be viewed in execution logs. Cloud Scheduler will retry the job according to the RetryConfig. The allowed duration for this deadline is: * For HTTP targets, between 15 seconds and 30 minutes. * For App Engine HTTP targets, between 15 seconds and 24 hours.
	AttemptDeadline string `json:"attemptDeadline,omitempty" yaml:"attemptDeadline,omitempty"`
	// Optionally caller-specified in CreateJob or UpdateJob. A human-readable description for the job. This string must not contain more than 500 characters.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// HTTP target.
	HTTPTarget map[string]interface{} `json:"httpTarget,omitempty" yaml:"httpTarget,omitempty"`
	// Immutable. The location for the resource
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Pub/Sub target.
	PubsubTarget map[string]interface{} `json:"pubsubTarget,omitempty" yaml:"pubsubTarget,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Settings that determine the retry behavior.
	RetryConfig map[string]interface{} `json:"retryConfig,omitempty" yaml:"retryConfig,omitempty"`
	// Required, except when used with UpdateJob. Describes the schedule on which the job will be executed. The schedule can be either of the following types: * [Crontab](http://en.wikipedia.org/wiki/Cron#Overview) * English-like [schedule](https://cloud.google.com/scheduler/docs/configuring/cron-job-schedules) As a general rule, execution `n + 1` of a job will not begin until execution `n` has finished. Cloud Scheduler will never allow two simultaneously outstanding executions. For example, this implies that if the `n+1`th execution is scheduled to run at 16:00 but the `n`th execution takes until 16:15, the `n+1`th execution will not start until `16:15`. A scheduled start time will be delayed if the previous execution has not ended when its scheduled time occurs. If retry_count > 0 and a job attempt fails, the job will be tried a total of retry_count times, with exponential backoff, until the next scheduled start time.
	Schedule string `json:"schedule,omitempty" yaml:"schedule,omitempty"`
	// Specifies the time zone to be used in interpreting schedule. The value of this field must be a time zone name from the [tz database](http://en.wikipedia.org/wiki/Tz_database). Note that some time zones include a provision for daylight savings time. The rules for daylight saving time are determined by the chosen tz. For UTC use the string "utc". If a time zone is not specified, the default will be in UTC (also known as GMT).
	TimeZone string `json:"timeZone,omitempty" yaml:"timeZone,omitempty"`
}

// CloudSchedulerJobStatus defines the observed state of CloudSchedulerJob.
type CloudSchedulerJobStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
