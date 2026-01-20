package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// VertexAIMetadataStore is the Schema for the VertexAIMetadataStore API
type VertexAIMetadataStore struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VertexAIMetadataStoreSpec   `json:"spec,omitempty"`
	Status VertexAIMetadataStoreStatus `json:"status,omitempty"`
}

// VertexAIMetadataStoreSpec defines the desired state of VertexAIMetadataStore.
type VertexAIMetadataStoreSpec struct {
	// Optional. Dataplex integration settings.
	DataplexConfig map[string]interface{} `json:"dataplexConfig,omitempty" yaml:"dataplexConfig,omitempty"`
	// Description of the MetadataStore.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Customer-managed encryption key spec for a Metadata Store. If set, this Metadata Store and all sub-resources of this Metadata Store are secured using this key.
	EncryptionSpec map[string]interface{} `json:"encryptionSpec,omitempty" yaml:"encryptionSpec,omitempty"`
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// The region of the Metadata Store. eg us-central1.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`
	// The VertexAIMetadataStore name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// VertexAIMetadataStoreStatus defines the observed state of VertexAIMetadataStore.
type VertexAIMetadataStoreStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
