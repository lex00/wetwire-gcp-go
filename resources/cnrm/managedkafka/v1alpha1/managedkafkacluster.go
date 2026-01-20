package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ManagedKafkaCluster is the Schema for the ManagedKafkaCluster API
type ManagedKafkaCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ManagedKafkaClusterSpec   `json:"spec,omitempty"`
	Status ManagedKafkaClusterStatus `json:"status,omitempty"`
}

// ManagedKafkaClusterSpec defines the desired state of ManagedKafkaCluster.
type ManagedKafkaClusterSpec struct {
	// Required. Capacity configuration for the Kafka cluster.
	CapacityConfig map[string]interface{} `json:"capacityConfig,omitempty" yaml:"capacityConfig,omitempty"`
	// Required. Configuration properties for a Kafka cluster deployed to Google Cloud Platform.
	GcpConfig map[string]interface{} `json:"gcpConfig,omitempty" yaml:"gcpConfig,omitempty"`
	// Optional. Labels as key value pairs.
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Optional. Rebalance configuration for the Kafka cluster.
	RebalanceConfig map[string]interface{} `json:"rebalanceConfig,omitempty" yaml:"rebalanceConfig,omitempty"`
	// The GCP resource identifier. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// ManagedKafkaClusterStatus defines the observed state of ManagedKafkaCluster.
type ManagedKafkaClusterStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
