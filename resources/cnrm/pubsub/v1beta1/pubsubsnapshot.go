package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PubSubSnapshot is the Schema for the PubSubSnapshot API
type PubSubSnapshot struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PubSubSnapshotSpec   `json:"spec,omitempty"`
	Status PubSubSnapshotStatus `json:"status,omitempty"`
}

// PubSubSnapshotSpec defines the desired state of PubSubSnapshot.
type PubSubSnapshotSpec struct {
	// Optional. See [Creating and managing labels] (https://cloud.google.com/pubsub/docs/labels).
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	// The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// The subscription whose backlog the snapshot retains. Specifically, the created snapshot is guaranteed to retain: (a) The existing backlog on the subscription. More precisely, this is defined as the messages in the subscription's backlog that are unacknowledged upon the successful completion of the snapshots.create request; as well as: (b) Any messages published to the subscription's topic following the successful completion of the snapshots.create request. Format is projects/{project}/subscriptions/{sub}.
	PubSubSubscriptionRef map[string]interface{} `json:"pubSubSubscriptionRef,omitempty" yaml:"pubSubSubscriptionRef,omitempty"`
	// The PubSubSnapshot name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Optional. The name of the topic from which this snapshot is retaining messages.
	TopicRef map[string]interface{} `json:"topicRef,omitempty" yaml:"topicRef,omitempty"`
}

// PubSubSnapshotStatus defines the observed state of PubSubSnapshot.
type PubSubSnapshotStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
