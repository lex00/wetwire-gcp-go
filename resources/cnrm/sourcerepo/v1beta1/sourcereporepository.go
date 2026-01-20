package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// SourceRepoRepository represents a sourcerepo.cnrm.cloud.google.com SourceRepoRepository resource.
type SourceRepoRepository struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SourceRepoRepositorySpec   `json:"spec,omitempty"`
	Status SourceRepoRepositoryStatus `json:"status,omitempty"`
}

// SourceRepoRepositorySpec defines the desired state of SourceRepoRepository.
type SourceRepoRepositorySpec struct {
	// How this repository publishes a change in the repository through Cloud Pub/Sub.
	// Keyed by the topic names.
	PubsubConfigs []map[string]interface{} `json:"pubsubConfigs,omitempty" yaml:"pubsubConfigs,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// SourceRepoRepositoryStatus defines the observed state of SourceRepoRepository.
type SourceRepoRepositoryStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
