package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NotebooksEnvironment is the Schema for the NotebooksEnvironment API
type NotebooksEnvironment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NotebooksEnvironmentSpec   `json:"spec,omitempty"`
	Status NotebooksEnvironmentStatus `json:"status,omitempty"`
}

// NotebooksEnvironmentSpec defines the desired state of NotebooksEnvironment.
type NotebooksEnvironmentSpec struct {
	// Use a container image to start the notebook instance.
	ContainerImage map[string]interface{} `json:"containerImage,omitempty" yaml:"containerImage,omitempty"`
	// A brief description of this environment.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Display name of this environment for the UI.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// The location for the resource.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Path to a Bash script that automatically runs after a notebook instance fully boots up. The path must be a URL or Cloud Storage path. Example: `"gs://path-to-file/file-name"`
	PostStartupScript string `json:"postStartupScript,omitempty" yaml:"postStartupScript,omitempty"`
	// The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// The NotebooksEnvironment name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Use a Compute Engine VM image to start the notebook instance.
	VmImage map[string]interface{} `json:"vmImage,omitempty" yaml:"vmImage,omitempty"`
}

// NotebooksEnvironmentStatus defines the observed state of NotebooksEnvironment.
type NotebooksEnvironmentStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
