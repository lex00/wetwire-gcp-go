package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ApigeeInstanceAttachment is the Schema for the ApigeeInstanceAttachment API
type ApigeeInstanceAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ApigeeInstanceAttachmentSpec   `json:"spec,omitempty"`
	Status ApigeeInstanceAttachmentStatus `json:"status,omitempty"`
}

// ApigeeInstanceAttachmentSpec defines the desired state of ApigeeInstanceAttachment.
type ApigeeInstanceAttachmentSpec struct {
	// ID of the attached environment.
	EnvironmentRef map[string]interface{} `json:"environmentRef,omitempty" yaml:"environmentRef,omitempty"`
	// Reference to parent Apigee Instance.
	InstanceRef map[string]interface{} `json:"instanceRef,omitempty" yaml:"instanceRef,omitempty"`
	// The ApigeeInstanceAttachment name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// ApigeeInstanceAttachmentStatus defines the observed state of ApigeeInstanceAttachment.
type ApigeeInstanceAttachmentStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
