package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// VertexAIFeaturestore is the Schema for the VertexAIFeaturestore API
type VertexAIFeaturestore struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VertexAIFeaturestoreSpec   `json:"spec,omitempty"`
	Status VertexAIFeaturestoreStatus `json:"status,omitempty"`
}

// VertexAIFeaturestoreSpec defines the desired state of VertexAIFeaturestore.
type VertexAIFeaturestoreSpec struct {
	// Optional. Customer-managed encryption key spec for data storage. If set, both of the online and offline data storage will be secured by this key.
	EncryptionSpec map[string]interface{} `json:"encryptionSpec,omitempty" yaml:"encryptionSpec,omitempty"`
	// Optional. The labels with user-defined metadata to organize your
	// Featurestore.
	// Label keys and values can be no longer than 64 characters
	// (Unicode codepoints), can only contain lowercase letters, numeric
	// characters, underscores and dashes. International characters are allowed.
	// See https://goo.gl/xmQnxf for more information on and examples of labels.
	// No more than 64 user labels can be associated with one Featurestore(System
	// labels are excluded)."
	// System reserved label keys are prefixed with "aiplatform.googleapis.com/"
	// and are immutable.
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	// The location for the resource.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Optional. Config for online storage resources. The field should not co-exist with the field of `OnlineStoreReplicationConfig`. If both of it and OnlineStoreReplicationConfig are unset, the feature store will not have an online store and cannot be used for online serving.
	OnlineServingConfig map[string]interface{} `json:"onlineServingConfig,omitempty" yaml:"onlineServingConfig,omitempty"`
	// Optional. TTL in days for feature values that will be stored in online serving storage. The Feature Store online storage periodically removes obsolete feature values older than `online_storage_ttl_days` since the feature generation time. Note that `online_storage_ttl_days` should be less than or equal to `offline_storage_ttl_days` for each EntityType under a featurestore. If not set, default to 4000 days
	OnlineStorageTTLDays int32 `json:"onlineStorageTTLDays,omitempty" yaml:"onlineStorageTTLDays,omitempty"`
	// The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// The VertexAIFeaturestore name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// VertexAIFeaturestoreStatus defines the observed state of VertexAIFeaturestore.
type VertexAIFeaturestoreStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
