package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ServiceNetworkingConnection represents a servicenetworking.cnrm.cloud.google.com ServiceNetworkingConnection resource.
type ServiceNetworkingConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ServiceNetworkingConnectionSpec   `json:"spec,omitempty"`
	Status ServiceNetworkingConnectionStatus `json:"status,omitempty"`
}

// ServiceNetworkingConnectionSpec defines the desired state of ServiceNetworkingConnection.
type ServiceNetworkingConnectionSpec struct {
	NetworkRef map[string]interface{} `json:"networkRef,omitempty" yaml:"networkRef,omitempty"`
	ReservedPeeringRanges []map[string]interface{} `json:"reservedPeeringRanges,omitempty" yaml:"reservedPeeringRanges,omitempty"`
	// Immutable. Provider peering service that is managing peering connectivity for a service provider organization. For Google services that support this functionality it is 'servicenetworking.googleapis.com'.
	Service string `json:"service,omitempty" yaml:"service,omitempty"`
}

// ServiceNetworkingConnectionStatus defines the observed state of ServiceNetworkingConnection.
type ServiceNetworkingConnectionStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
