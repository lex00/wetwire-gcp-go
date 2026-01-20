package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NetworkConnectivityHub represents a networkconnectivity.cnrm.cloud.google.com NetworkConnectivityHub resource.
type NetworkConnectivityHub struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NetworkConnectivityHubSpec   `json:"spec,omitempty"`
	Status NetworkConnectivityHubStatus `json:"status,omitempty"`
}

// NetworkConnectivityHubSpec defines the desired state of NetworkConnectivityHub.
type NetworkConnectivityHubSpec struct {
	// An optional description of the hub.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Immutable. The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// NetworkConnectivityHubStatus defines the observed state of NetworkConnectivityHub.
type NetworkConnectivityHubStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
