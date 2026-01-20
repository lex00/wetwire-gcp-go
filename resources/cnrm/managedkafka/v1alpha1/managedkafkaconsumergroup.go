package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ManagedKafkaConsumerGroup is the Schema for the ManagedKafkaConsumerGroup API
type ManagedKafkaConsumerGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ManagedKafkaConsumerGroupSpec   `json:"spec,omitempty"`
	Status ManagedKafkaConsumerGroupStatus `json:"status,omitempty"`
}

// ManagedKafkaConsumerGroupSpec defines the desired state of ManagedKafkaConsumerGroup.
type ManagedKafkaConsumerGroupSpec struct {
	// ClusterRef defines the resource reference to ManagedKafkaCluster, which "External" field holds the GCP identifier for the KRM object.
	ClusterRef map[string]interface{} `json:"clusterRef,omitempty" yaml:"clusterRef,omitempty"`
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// The ManagedKafkaConsumerGroup name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// ManagedKafkaConsumerGroupStatus defines the observed state of ManagedKafkaConsumerGroup.
type ManagedKafkaConsumerGroupStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
