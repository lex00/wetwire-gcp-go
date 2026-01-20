package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// BigQueryReservationReservation is the Schema for the BigQueryReservationReservation API
type BigQueryReservationReservation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BigQueryReservationReservationSpec   `json:"spec,omitempty"`
	Status BigQueryReservationReservationStatus `json:"status,omitempty"`
}

// BigQueryReservationReservationSpec defines the desired state of BigQueryReservationReservation.
type BigQueryReservationReservationSpec struct {
	// Optional. The configuration parameters for the auto scaling feature.
	Autoscale map[string]interface{} `json:"autoscale,omitempty" yaml:"autoscale,omitempty"`
	// Job concurrency target which sets a soft upper bound on the number of jobs that can run concurrently in this reservation. This is a soft target due to asynchronous nature of the system and various optimizations for small queries. Default value is 0 which means that concurrency target will be automatically computed by the system. NOTE: this field is exposed as target job concurrency in the Information Schema, DDL and BigQuery CLI.
	Concurrency int64 `json:"concurrency,omitempty" yaml:"concurrency,omitempty"`
	// Immutable. Optional. Edition of the reservation. Valid values are STANDARD, ENTERPRISE, ENTERPRISE_PLUS
	Edition string `json:"edition,omitempty" yaml:"edition,omitempty"`
	// Optional. This field is only set for reservations using the managed disaster recovery feature. Users can set this to create a failover reservation.
	Failover map[string]interface{} `json:"failover,omitempty" yaml:"failover,omitempty"`
	// If false, any query or pipeline job using this reservation will use idle slots from other reservations within the same admin project. If true, a query or pipeline job using this reservation will execute with the slot capacity specified in the slot_capacity field at most.
	IgnoreIdleSlots bool `json:"ignoreIdleSlots,omitempty" yaml:"ignoreIdleSlots,omitempty"`
	// Immutable. You can configure spec.secondaryLocation to enable the reservation fail-over to a secondary location, in which case the primary location could be different from the spec.location.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The BigQuery Reservation ID used for resource creation or acquisition. It must only contain lower case alphanumeric characters or dashes. It must start with a letter and must not end with a dash. Its maximum length is 64 characters. For creation: If specified, this value is used as the Reservation ID. If not provided, a UUID will be generated and assigned as the Reservation ID. For acquisition: This field must be provided to identify the Reservation resource to acquire.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Optional. Baseline slots available to this reservation. A slot is a unit of
	// computational power in BigQuery, and serves as the unit of parallelism.
	// Queries using this reservation might use more slots during runtime if
	// ignore_idle_slots is set to false, or autoscaling is enabled.
	SlotCapacity int64 `json:"slotCapacity,omitempty" yaml:"slotCapacity,omitempty"`
}

// BigQueryReservationReservationStatus defines the observed state of BigQueryReservationReservation.
type BigQueryReservationReservationStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
