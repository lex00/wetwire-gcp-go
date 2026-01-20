package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MonitoringNotificationChannel is the Schema for the MonitoringNotificationChannel API
type MonitoringNotificationChannel struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MonitoringNotificationChannelSpec   `json:"spec,omitempty"`
	Status MonitoringNotificationChannelStatus `json:"status,omitempty"`
}

// MonitoringNotificationChannelSpec defines the desired state of MonitoringNotificationChannel.
type MonitoringNotificationChannelSpec struct {
	// An optional human-readable description of this notification channel. This description may provide additional details, beyond the display name, for the channel. This may not exceed 1024 Unicode characters.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Whether notifications are forwarded to the described channel. This makes it possible to disable delivery of notifications to a particular channel without removing the channel from all alerting policies that reference the channel. This is a more convenient approach when the change is temporary and you want to receive notifications from the same set of alerting policies on the channel at some point in the future.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	// If true, the notification channel will be deleted regardless of its use in alert policies (the policies will be updated to remove the channel). If false, channels that are still referenced by an existing alerting policy will fail to be deleted in a delete operation.
	ForceDelete bool `json:"forceDelete,omitempty" yaml:"forceDelete,omitempty"`
	// Configuration fields that define the channel and its behavior. The permissible and required labels are specified in the NotificationChannelDescriptor.labels of the NotificationChannelDescriptor corresponding to the type field.
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	// Immutable. Optional. The service-generated name of theresource. Used for acquisition only. Leave unset to create a new resource.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Different notification type behaviors are configured primarily using the the 'labels' field on this
	// resource. This block contains the labels which contain secrets or passwords so that they can be marked
	// sensitive and hidden from plan output. The name of the field, eg: password, will be the key
	// in the 'labels' map in the api request.
	// Credentials may not be specified in both locations and will cause an error. Changing from one location
	// to a different credential configuration in the config will require an apply to update state.
	SensitiveLabels map[string]interface{} `json:"sensitiveLabels,omitempty" yaml:"sensitiveLabels,omitempty"`
	// The type of the notification channel. This field matches the value of the NotificationChannelDescriptor.type field. See https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.notificationChannelDescriptors/list to get the list of valid values such as "email", "slack", etc...
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}

// MonitoringNotificationChannelStatus defines the observed state of MonitoringNotificationChannel.
type MonitoringNotificationChannelStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
