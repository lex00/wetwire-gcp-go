package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// BigQueryReservationCapacityCommitment represents a bigqueryreservation.cnrm.cloud.google.com BigQueryReservationCapacityCommitment resource.
type BigQueryReservationCapacityCommitment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BigQueryReservationCapacityCommitmentSpec   `json:"spec,omitempty"`
	Status BigQueryReservationCapacityCommitmentStatus `json:"status,omitempty"`
}

// BigQueryReservationCapacityCommitmentSpec defines the desired state of BigQueryReservationCapacityCommitment.
type BigQueryReservationCapacityCommitmentSpec struct {
	// Immutable. The edition type. Valid values are STANDARD, ENTERPRISE, ENTERPRISE_PLUS.
	Edition string `json:"edition,omitempty" yaml:"edition,omitempty"`
	// Immutable. If true, fail the request if another project in the organization has a capacity commitment.
	EnforceSingleAdminProjectPerOrg string `json:"enforceSingleAdminProjectPerOrg,omitempty" yaml:"enforceSingleAdminProjectPerOrg,omitempty"`
	// Immutable. The geographic location where the transfer config should reside.
	// Examples: US, EU, asia-northeast1. The default value is US.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Capacity commitment plan. Valid values are at https://cloud.google.com/bigquery/docs/reference/reservations/rpc/google.cloud.bigquery.reservation.v1#commitmentplan.
	Plan string `json:"plan,omitempty" yaml:"plan,omitempty"`
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// The plan this capacity commitment is converted to after commitmentEndTime passes. Once the plan is changed, committed period is extended according to commitment plan. Only applicable some commitment plans.
	RenewalPlan string `json:"renewalPlan,omitempty" yaml:"renewalPlan,omitempty"`
	// Immutable. Optional. The capacityCommitmentId of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Immutable. Number of slots in this commitment.
	SlotCount int32 `json:"slotCount,omitempty" yaml:"slotCount,omitempty"`
}

// BigQueryReservationCapacityCommitmentStatus defines the observed state of BigQueryReservationCapacityCommitment.
type BigQueryReservationCapacityCommitmentStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
