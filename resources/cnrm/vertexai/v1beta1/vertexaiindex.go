package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// VertexAIIndex represents a vertexai.cnrm.cloud.google.com VertexAIIndex resource.
type VertexAIIndex struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VertexAIIndexSpec   `json:"spec,omitempty"`
	Status VertexAIIndexStatus `json:"status,omitempty"`
}

// VertexAIIndexSpec defines the desired state of VertexAIIndex.
type VertexAIIndexSpec struct {
	// The description of the Index.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// The display name of the Index. The name can be up to 128 characters long and can consist of any UTF-8 characters.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// Immutable. The update method to use with this Index. The value must be the followings. If not set, BATCH_UPDATE will be used by default.
	// * BATCH_UPDATE: user can call indexes.patch with files on Cloud Storage of datapoints to update.
	// * STREAM_UPDATE: user can call indexes.upsertDatapoints/DeleteDatapoints to update the Index and the updates will be applied in corresponding DeployedIndexes in nearly real-time.
	IndexUpdateMethod string `json:"indexUpdateMethod,omitempty" yaml:"indexUpdateMethod,omitempty"`
	// An additional information about the Index.
	Metadata map[string]interface{} `json:"metadata,omitempty" yaml:"metadata,omitempty"`
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. The region of the index. eg us-central1.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`
	// Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// VertexAIIndexStatus defines the observed state of VertexAIIndex.
type VertexAIIndexStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
