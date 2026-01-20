package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// VertexAIFeaturestoreEntityType represents a vertexai.cnrm.cloud.google.com VertexAIFeaturestoreEntityType resource.
type VertexAIFeaturestoreEntityType struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VertexAIFeaturestoreEntityTypeSpec   `json:"spec,omitempty"`
	Status VertexAIFeaturestoreEntityTypeStatus `json:"status,omitempty"`
}

// VertexAIFeaturestoreEntityTypeSpec defines the desired state of VertexAIFeaturestoreEntityType.
type VertexAIFeaturestoreEntityTypeSpec struct {
	// Optional. Description of the EntityType.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Immutable. The name of the Featurestore to use, in the format projects/{project}/locations/{location}/featurestores/{featurestore}.
	Featurestore string `json:"featurestore,omitempty" yaml:"featurestore,omitempty"`
	// The default monitoring configuration for all Features under this EntityType.
	// If this is populated with [FeaturestoreMonitoringConfig.monitoring_interval] specified, snapshot analysis monitoring is enabled. Otherwise, snapshot analysis monitoring is disabled.
	MonitoringConfig map[string]interface{} `json:"monitoringConfig,omitempty" yaml:"monitoringConfig,omitempty"`
	// Config for data retention policy in offline storage. TTL in days for feature values that will be stored in offline storage. The Feature Store offline storage periodically removes obsolete feature values older than offlineStorageTtlDays since the feature generation time. If unset (or explicitly set to 0), default to 4000 days TTL.
	OfflineStorageTtlDays int32 `json:"offlineStorageTtlDays,omitempty" yaml:"offlineStorageTtlDays,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// VertexAIFeaturestoreEntityTypeStatus defines the observed state of VertexAIFeaturestoreEntityType.
type VertexAIFeaturestoreEntityTypeStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
