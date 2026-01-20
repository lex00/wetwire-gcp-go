package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// VertexAITensorboard represents a vertexai.cnrm.cloud.google.com VertexAITensorboard resource.
type VertexAITensorboard struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VertexAITensorboardSpec   `json:"spec,omitempty"`
	Status VertexAITensorboardStatus `json:"status,omitempty"`
}

// VertexAITensorboardSpec defines the desired state of VertexAITensorboard.
type VertexAITensorboardSpec struct {
	// Description of this Tensorboard.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// User provided name of this Tensorboard.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// Immutable. Customer-managed encryption key spec for a Tensorboard. If set, this Tensorboard and all sub-resources of this Tensorboard will be secured by this key.
	EncryptionSpec map[string]interface{} `json:"encryptionSpec,omitempty" yaml:"encryptionSpec,omitempty"`
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. The region of the tensorboard. eg us-central1.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`
	// Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// VertexAITensorboardStatus defines the observed state of VertexAITensorboard.
type VertexAITensorboardStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
