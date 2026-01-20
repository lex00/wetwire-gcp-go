package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EventarcTrigger represents a eventarc.cnrm.cloud.google.com EventarcTrigger resource.
type EventarcTrigger struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EventarcTriggerSpec   `json:"spec,omitempty"`
	Status EventarcTriggerStatus `json:"status,omitempty"`
}

// EventarcTriggerSpec defines the desired state of EventarcTrigger.
type EventarcTriggerSpec struct {
	// Immutable.
	ChannelRef map[string]interface{} `json:"channelRef,omitempty" yaml:"channelRef,omitempty"`
	// Required. Destination specifies where the events should be sent to.
	Destination map[string]interface{} `json:"destination,omitempty" yaml:"destination,omitempty"`
	// Optional. EventDataContentType specifies the type of payload in MIME format that is expected from the CloudEvent data field. This is set to `application/json` if the value is not defined.
	EventDataContentType string `json:"eventDataContentType,omitempty" yaml:"eventDataContentType,omitempty"`
	// Immutable. The location for the resource
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Required. null The list of filters that applies to event attributes. Only events that match all the provided filters will be sent to the destination.
	MatchingCriteria []map[string]interface{} `json:"matchingCriteria,omitempty" yaml:"matchingCriteria,omitempty"`
	// Immutable. The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	ServiceAccountRef map[string]interface{} `json:"serviceAccountRef,omitempty" yaml:"serviceAccountRef,omitempty"`
	// Immutable. Optional. In order to deliver messages, Eventarc may use other GCP products as transport intermediary. This field contains a reference to that transport intermediary. This information can be used for debugging purposes.
	Transport map[string]interface{} `json:"transport,omitempty" yaml:"transport,omitempty"`
}

// EventarcTriggerStatus defines the observed state of EventarcTrigger.
type EventarcTriggerStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
