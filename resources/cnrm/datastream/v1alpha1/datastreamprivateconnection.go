package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DatastreamPrivateConnection is the Schema for the DatastreamPrivateConnection API
type DatastreamPrivateConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DatastreamPrivateConnectionSpec   `json:"spec,omitempty"`
	Status DatastreamPrivateConnectionStatus `json:"status,omitempty"`
}

// DatastreamPrivateConnectionSpec defines the desired state of DatastreamPrivateConnection.
type DatastreamPrivateConnectionSpec struct {
	// Required. Display name.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// Labels.
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	// Immutable.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// The DatastreamPrivateConnection name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// VPC Peering Config.
	VpcPeeringConfig map[string]interface{} `json:"vpcPeeringConfig,omitempty" yaml:"vpcPeeringConfig,omitempty"`
}

// DatastreamPrivateConnectionStatus defines the observed state of DatastreamPrivateConnection.
type DatastreamPrivateConnectionStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
