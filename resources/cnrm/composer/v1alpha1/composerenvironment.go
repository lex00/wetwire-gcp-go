package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ComposerEnvironment is the Schema for the ComposerEnvironment API
type ComposerEnvironment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComposerEnvironmentSpec   `json:"spec,omitempty"`
	Status ComposerEnvironmentStatus `json:"status,omitempty"`
}

// ComposerEnvironmentSpec defines the desired state of ComposerEnvironment.
type ComposerEnvironmentSpec struct {
	// Optional. Configuration parameters for this environment.
	Config map[string]interface{} `json:"config,omitempty" yaml:"config,omitempty"`
	// Optional. User-defined labels for this environment.
	// The labels map can contain no more than 64 entries. Entries of the labels
	// map are UTF8 strings that comply with the following restrictions:
	// * Keys must conform to regexp: [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}
	// * Values must conform to regexp:  [\p{Ll}\p{Lo}\p{N}_-]{0,63}
	// * Both keys and values are additionally constrained to be <= 128 bytes in
	// size.
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	// Immutable. The name of the location where the Environment will be created. Required.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Immutable. The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// The ComposerEnvironment name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Optional. Storage configuration for this environment.
	StorageConfig map[string]interface{} `json:"storageConfig,omitempty" yaml:"storageConfig,omitempty"`
}

// ComposerEnvironmentStatus defines the observed state of ComposerEnvironment.
type ComposerEnvironmentStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
