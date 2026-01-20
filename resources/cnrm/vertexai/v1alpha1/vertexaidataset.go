package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// VertexAIDataset represents a vertexai.cnrm.cloud.google.com VertexAIDataset resource.
type VertexAIDataset struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VertexAIDatasetSpec   `json:"spec,omitempty"`
	Status VertexAIDatasetStatus `json:"status,omitempty"`
}

// VertexAIDatasetSpec defines the desired state of VertexAIDataset.
type VertexAIDatasetSpec struct {
	// The user-defined name of the Dataset. The name can be up to 128 characters long and can be consist of any UTF-8 characters.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// Immutable. Customer-managed encryption key spec for a Dataset. If set, this Dataset and all sub-resources of this Dataset will be secured by this key.
	EncryptionSpec map[string]interface{} `json:"encryptionSpec,omitempty" yaml:"encryptionSpec,omitempty"`
	// Immutable. Points to a YAML file stored on Google Cloud Storage describing additional information about the Dataset. The schema is defined as an OpenAPI 3.0.2 Schema Object. The schema files that can be used here are found in gs://google-cloud-aiplatform/schema/dataset/metadata/.
	MetadataSchemaURI string `json:"metadataSchemaUri,omitempty" yaml:"metadataSchemaUri,omitempty"`
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. The region of the dataset. eg us-central1.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`
	// Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// VertexAIDatasetStatus defines the observed state of VertexAIDataset.
type VertexAIDatasetStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
