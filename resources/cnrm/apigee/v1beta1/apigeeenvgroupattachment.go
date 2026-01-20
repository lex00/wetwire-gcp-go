package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ApigeeEnvgroupAttachment is the Schema for the ApigeeEnvgroupAttachment API
type ApigeeEnvgroupAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ApigeeEnvgroupAttachmentSpec   `json:"spec,omitempty"`
	Status ApigeeEnvgroupAttachmentStatus `json:"status,omitempty"`
}

// ApigeeEnvgroupAttachmentSpec defines the desired state of ApigeeEnvgroupAttachment.
type ApigeeEnvgroupAttachmentSpec struct {
	// Reference to parent Environment Group
	EnvgroupRef map[string]interface{} `json:"envgroupRef,omitempty" yaml:"envgroupRef,omitempty"`
	// Required. ID of the attached environment.
	EnvironmentRef map[string]interface{} `json:"environmentRef,omitempty" yaml:"environmentRef,omitempty"`
	// The ApigeeEnvgroupAttachment name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// ApigeeEnvgroupAttachmentStatus defines the observed state of ApigeeEnvgroupAttachment.
type ApigeeEnvgroupAttachmentStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
