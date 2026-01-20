package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// TasksQueue is the Schema for the TasksQueue API
type TasksQueue struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TasksQueueSpec   `json:"spec,omitempty"`
	Status TasksQueueStatus `json:"status,omitempty"`
}

// TasksQueueSpec defines the desired state of TasksQueue.
type TasksQueueSpec struct {
	// Overrides for
	// [task-level
	// app_engine_routing][google.cloud.tasks.v2.AppEngineHttpRequest.app_engine_routing].
	// These settings apply only to
	// [App Engine tasks][google.cloud.tasks.v2.AppEngineHttpRequest] in this
	// queue. [Http tasks][google.cloud.tasks.v2.HttpRequest] are not affected.
	// If set, `app_engine_routing_override` is used for all
	// [App Engine tasks][google.cloud.tasks.v2.AppEngineHttpRequest] in the
	// queue, no matter what the setting is for the [task-level
	// app_engine_routing][google.cloud.tasks.v2.AppEngineHttpRequest.app_engine_routing].
	AppEngineRoutingOverride map[string]interface{} `json:"appEngineRoutingOverride,omitempty" yaml:"appEngineRoutingOverride,omitempty"`
	// Required. The location of the queue.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Required. The host project of the queue.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Rate limits for task dispatches.
	// [rate_limits][google.cloud.tasks.v2.Queue.rate_limits] and
	// [retry_config][google.cloud.tasks.v2.Queue.retry_config] are related
	// because they both control task attempts. However they control task attempts
	// in different ways:
	// * [rate_limits][google.cloud.tasks.v2.Queue.rate_limits] controls the total
	// rate of
	// dispatches from a queue (i.e. all traffic dispatched from the
	// queue, regardless of whether the dispatch is from a first
	// attempt or a retry).
	// * [retry_config][google.cloud.tasks.v2.Queue.retry_config] controls what
	// happens to
	// particular a task after its first attempt fails. That is,
	// [retry_config][google.cloud.tasks.v2.Queue.retry_config] controls task
	// retries (the second attempt, third attempt, etc).
	// The queue's actual dispatch rate is the result of:
	// * Number of tasks in the queue
	// * User-specified throttling:
	// [rate_limits][google.cloud.tasks.v2.Queue.rate_limits],
	// [retry_config][google.cloud.tasks.v2.Queue.retry_config], and the
	// [queue's state][google.cloud.tasks.v2.Queue.state].
	// * System throttling due to `429` (Too Many Requests) or `503` (Service
	// Unavailable) responses from the worker, high error rates, or to smooth
	// sudden large traffic spikes.
	RateLimits map[string]interface{} `json:"rateLimits,omitempty" yaml:"rateLimits,omitempty"`
	// The TasksQueue name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Settings that determine the retry behavior.
	// * For tasks created using Cloud Tasks: the queue-level retry settings
	// apply to all tasks in the queue that were created using Cloud Tasks.
	// Retry settings cannot be set on individual tasks.
	// * For tasks created using the App Engine SDK: the queue-level retry
	// settings apply to all tasks in the queue which do not have retry settings
	// explicitly set on the task and were created by the App Engine SDK. See
	// [App Engine
	// documentation](https://cloud.google.com/appengine/docs/standard/python/taskqueue/push/retrying-tasks).
	RetryConfig map[string]interface{} `json:"retryConfig,omitempty" yaml:"retryConfig,omitempty"`
	// Configuration options for writing logs to [Stackdriver Logging](https://cloud.google.com/logging/docs/). If this field is unset, then no logs are written.
	StackdriverLoggingConfig map[string]interface{} `json:"stackdriverLoggingConfig,omitempty" yaml:"stackdriverLoggingConfig,omitempty"`
}

// TasksQueueStatus defines the observed state of TasksQueue.
type TasksQueueStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
