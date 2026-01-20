package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DatastreamRoute is the Schema for the DatastreamRoute API
type DatastreamRoute struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DatastreamRouteSpec   `json:"spec,omitempty"`
	Status DatastreamRouteStatus `json:"status,omitempty"`
}

// DatastreamRouteSpec defines the desired state of DatastreamRoute.
type DatastreamRouteSpec struct {
	// Required. Destination address for connection
	DestinationAddress string `json:"destinationAddress,omitempty" yaml:"destinationAddress,omitempty"`
	// Destination port for connection
	DestinationPort int32 `json:"destinationPort,omitempty" yaml:"destinationPort,omitempty"`
	// Required. Display name.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// Labels.
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	// Required. The parent that owns the collection of Routes.
	PrivateConnectionRef map[string]interface{} `json:"privateConnectionRef,omitempty" yaml:"privateConnectionRef,omitempty"`
	// The DatastreamRoute name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// DatastreamRouteStatus defines the observed state of DatastreamRoute.
type DatastreamRouteStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
