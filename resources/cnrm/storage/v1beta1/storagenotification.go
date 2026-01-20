package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// StorageNotification represents a storage.cnrm.cloud.google.com StorageNotification resource.
type StorageNotification struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   StorageNotificationSpec   `json:"spec,omitempty"`
	Status StorageNotificationStatus `json:"status,omitempty"`
}

// StorageNotificationSpec defines the desired state of StorageNotification.
type StorageNotificationSpec struct {
	BucketRef map[string]interface{} `json:"bucketRef,omitempty" yaml:"bucketRef,omitempty"`
	// Immutable.  A set of key/value attribute pairs to attach to each Cloud Pub/Sub message published for this notification subscription.
	CustomAttributes map[string]string `json:"customAttributes,omitempty" yaml:"customAttributes,omitempty"`
	// Immutable. List of event type filters for this notification config. If not specified, Cloud Storage will send notifications for all event types. The valid types are: "OBJECT_FINALIZE", "OBJECT_METADATA_UPDATE", "OBJECT_DELETE", "OBJECT_ARCHIVE".
	EventTypes []string `json:"eventTypes,omitempty" yaml:"eventTypes,omitempty"`
	// Immutable. Specifies a prefix path filter for this notification config. Cloud Storage will only send notifications for objects in this bucket whose names begin with the specified prefix.
	ObjectNamePrefix string `json:"objectNamePrefix,omitempty" yaml:"objectNamePrefix,omitempty"`
	// Immutable. The desired content of the Payload. One of "JSON_API_V1" or "NONE".
	PayloadFormat string `json:"payloadFormat,omitempty" yaml:"payloadFormat,omitempty"`
	// Immutable. Optional. The service-generated notificationId of the resource. Used for acquisition only. Leave unset to create a new resource.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	TopicRef map[string]interface{} `json:"topicRef,omitempty" yaml:"topicRef,omitempty"`
}

// StorageNotificationStatus defines the observed state of StorageNotification.
type StorageNotificationStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
