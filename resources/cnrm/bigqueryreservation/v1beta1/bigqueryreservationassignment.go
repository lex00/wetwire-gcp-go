package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// BigQueryReservationAssignment is the Schema for the BigQueryReservationAssignment API
type BigQueryReservationAssignment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BigQueryReservationAssignmentSpec   `json:"spec,omitempty"`
	Status BigQueryReservationAssignmentStatus `json:"status,omitempty"`
}

// BigQueryReservationAssignmentSpec defines the desired state of BigQueryReservationAssignment.
type BigQueryReservationAssignmentSpec struct {
	// Immutable. Required. The resource which will use the reservation. E.g. `projects/myproject`, `folders/123`, or `organizations/456`.
	Assignee map[string]interface{} `json:"assignee,omitempty" yaml:"assignee,omitempty"`
	// Immutable. Which type of jobs will use the reservation.
	JobType string `json:"jobType,omitempty" yaml:"jobType,omitempty"`
	// The name of reservation to create a new assignment in, or to move the assignment to.
	ReservationRef map[string]interface{} `json:"reservationRef,omitempty" yaml:"reservationRef,omitempty"`
	// Immutable. Optional. The BigQueryReservationAssignment ID used for resource creation or acquisition. Service-generated.Can be set only if resource acquisition . For acquisition: This field must be provided to identify the Reservation resource to acquire.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// BigQueryReservationAssignmentStatus defines the observed state of BigQueryReservationAssignment.
type BigQueryReservationAssignmentStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
