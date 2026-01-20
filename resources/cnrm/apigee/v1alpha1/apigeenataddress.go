package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ApigeeNATAddress represents a apigee.cnrm.cloud.google.com ApigeeNATAddress resource.
type ApigeeNATAddress struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ApigeeNATAddressSpec   `json:"spec,omitempty"`
	Status ApigeeNATAddressStatus `json:"status,omitempty"`
}

// ApigeeNATAddressSpec defines the desired state of ApigeeNATAddress.
type ApigeeNATAddressSpec struct {
	// Immutable. The Apigee instance associated with the Apigee environment,
	// in the format 'organizations/{{org_name}}/instances/{{instance_name}}'.
	InstanceID string `json:"instanceId,omitempty" yaml:"instanceId,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// ApigeeNATAddressStatus defines the observed state of ApigeeNATAddress.
type ApigeeNATAddressStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
