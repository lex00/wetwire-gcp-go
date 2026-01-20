package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PubSubLiteTopic represents a pubsublite.cnrm.cloud.google.com PubSubLiteTopic resource.
type PubSubLiteTopic struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PubSubLiteTopicSpec   `json:"spec,omitempty"`
	Status PubSubLiteTopicStatus `json:"status,omitempty"`
}

// PubSubLiteTopicSpec defines the desired state of PubSubLiteTopic.
type PubSubLiteTopicSpec struct {
	// The settings for this topic's partitions.
	PartitionConfig map[string]interface{} `json:"partitionConfig,omitempty" yaml:"partitionConfig,omitempty"`
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// The region of the pubsub lite topic.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`
	// The settings for this topic's Reservation usage.
	ReservationConfig map[string]interface{} `json:"reservationConfig,omitempty" yaml:"reservationConfig,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// The settings for a topic's message retention.
	RetentionConfig map[string]interface{} `json:"retentionConfig,omitempty" yaml:"retentionConfig,omitempty"`
	// The zone of the pubsub lite topic.
	Zone string `json:"zone,omitempty" yaml:"zone,omitempty"`
}

// PubSubLiteTopicStatus defines the observed state of PubSubLiteTopic.
type PubSubLiteTopicStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
