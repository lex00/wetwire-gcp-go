package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PubSubLiteReservation represents a pubsublite.cnrm.cloud.google.com PubSubLiteReservation resource.
type PubSubLiteReservation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PubSubLiteReservationSpec   `json:"spec,omitempty"`
	Status PubSubLiteReservationStatus `json:"status,omitempty"`
}

// PubSubLiteReservationSpec defines the desired state of PubSubLiteReservation.
type PubSubLiteReservationSpec struct {
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// The region of the pubsub lite reservation.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// The reserved throughput capacity. Every unit of throughput capacity is
	// equivalent to 1 MiB/s of published messages or 2 MiB/s of subscribed
	// messages.
	ThroughputCapacity int32 `json:"throughputCapacity,omitempty" yaml:"throughputCapacity,omitempty"`
}

// PubSubLiteReservationStatus defines the observed state of PubSubLiteReservation.
type PubSubLiteReservationStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
