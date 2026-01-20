package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DLPJobTrigger represents a dlp.cnrm.cloud.google.com DLPJobTrigger resource.
type DLPJobTrigger struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DLPJobTriggerSpec   `json:"spec,omitempty"`
	Status DLPJobTriggerStatus `json:"status,omitempty"`
}

// DLPJobTriggerSpec defines the desired state of DLPJobTrigger.
type DLPJobTriggerSpec struct {
	// User provided description (max 256 chars)
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Display name (max 100 chars)
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// For inspect jobs, a snapshot of the configuration.
	InspectJob map[string]interface{} `json:"inspectJob,omitempty" yaml:"inspectJob,omitempty"`
	// Immutable. The location of the resource
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Immutable. The Project that this resource belongs to. Only one of [projectRef] may be specified.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Immutable. Required. A status for this trigger. Possible values: STATUS_UNSPECIFIED, HEALTHY, PAUSED, CANCELLED
	Status string `json:"status,omitempty" yaml:"status,omitempty"`
	// A list of triggers which will be OR'ed together. Only one in the list needs to trigger for a job to be started. The list may contain only a single Schedule trigger and must have at least one object.
	Triggers []map[string]interface{} `json:"triggers,omitempty" yaml:"triggers,omitempty"`
}

// DLPJobTriggerStatus defines the observed state of DLPJobTrigger.
type DLPJobTriggerStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
