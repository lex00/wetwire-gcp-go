package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ComputeReservation represents a compute.cnrm.cloud.google.com ComputeReservation resource.
type ComputeReservation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeReservationSpec   `json:"spec,omitempty"`
	Status ComputeReservationStatus `json:"status,omitempty"`
}

// ComputeReservationSpec defines the desired state of ComputeReservation.
type ComputeReservationSpec struct {
	// Immutable. An optional description of this resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Reservation for instances with specific machine shapes.
	SpecificReservation map[string]interface{} `json:"specificReservation,omitempty" yaml:"specificReservation,omitempty"`
	// Immutable. When set to true, only VMs that target this reservation by name can
	// consume this reservation. Otherwise, it can be consumed by VMs with
	// affinity for any reservation. Defaults to false.
	SpecificReservationRequired bool `json:"specificReservationRequired,omitempty" yaml:"specificReservationRequired,omitempty"`
	// Immutable. The zone where the reservation is made.
	Zone string `json:"zone,omitempty" yaml:"zone,omitempty"`
}

// ComputeReservationStatus defines the observed state of ComputeReservation.
type ComputeReservationStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
