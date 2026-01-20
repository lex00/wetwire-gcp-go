package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ManagedKafkaTopic is the Schema for the ManagedKafkaTopic API
type ManagedKafkaTopic struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ManagedKafkaTopicSpec   `json:"spec,omitempty"`
	Status ManagedKafkaTopicStatus `json:"status,omitempty"`
}

// ManagedKafkaTopicSpec defines the desired state of ManagedKafkaTopic.
type ManagedKafkaTopicSpec struct {
	// Required. Reference to the Kafka cluster to create the topic in.
	ClusterRef map[string]interface{} `json:"clusterRef,omitempty" yaml:"clusterRef,omitempty"`
	// Optional. Configurations for the topic that are overridden from the cluster defaults. The key of the map is a Kafka topic property name, for example: `cleanup.policy`, `compression.type`.
	Configs map[string]string `json:"configs,omitempty" yaml:"configs,omitempty"`
	// Required. the location of the Kafka resource. See https://cloud.google.com/managed-kafka/docs/locations for a list of supported locations.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Required. The number of partitions this topic has. The partition count can only be increased, not decreased. Please note that if partitions are increased for a topic that has a key, the partitioning logic or the ordering of the messages will be affected.
	PartitionCount int32 `json:"partitionCount,omitempty" yaml:"partitionCount,omitempty"`
	// The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Required. Immutable. The number of replicas of each partition. A replication factor of 3 is recommended for high availability.
	ReplicationFactor int32 `json:"replicationFactor,omitempty" yaml:"replicationFactor,omitempty"`
	// The GCP resource identifier. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// ManagedKafkaTopicStatus defines the observed state of ManagedKafkaTopic.
type ManagedKafkaTopicStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
