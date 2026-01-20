package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// VertexAIFeaturestoreEntityTypeFeature represents a vertexai.cnrm.cloud.google.com VertexAIFeaturestoreEntityTypeFeature resource.
type VertexAIFeaturestoreEntityTypeFeature struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VertexAIFeaturestoreEntityTypeFeatureSpec   `json:"spec,omitempty"`
	Status VertexAIFeaturestoreEntityTypeFeatureStatus `json:"status,omitempty"`
}

// VertexAIFeaturestoreEntityTypeFeatureSpec defines the desired state of VertexAIFeaturestoreEntityTypeFeature.
type VertexAIFeaturestoreEntityTypeFeatureSpec struct {
	// Description of the feature.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Immutable. The name of the Featurestore to use, in the format projects/{project}/locations/{location}/featurestores/{featurestore}/entityTypes/{entitytype}.
	Entitytype string `json:"entitytype,omitempty" yaml:"entitytype,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Immutable. Type of Feature value. Immutable. https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.featurestores.entityTypes.features#ValueType.
	ValueType string `json:"valueType,omitempty" yaml:"valueType,omitempty"`
}

// VertexAIFeaturestoreEntityTypeFeatureStatus defines the observed state of VertexAIFeaturestoreEntityTypeFeature.
type VertexAIFeaturestoreEntityTypeFeatureStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
