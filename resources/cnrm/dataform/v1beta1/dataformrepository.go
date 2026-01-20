package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DataformRepository is the Schema for the dataform API
type DataformRepository struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DataformRepositorySpec   `json:"spec,omitempty"`
	Status DataformRepositoryStatus `json:"status,omitempty"`
}

// DataformRepositorySpec defines the desired state of DataformRepository.
type DataformRepositorySpec struct {
	// Optional. The repository's user-friendly name.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// Optional. If set, configures this repository to be linked to a Git remote.
	GitRemoteSettings map[string]interface{} `json:"gitRemoteSettings,omitempty" yaml:"gitRemoteSettings,omitempty"`
	// Optional. The name of the Secret Manager secret version to be used to interpolate variables into the .npmrc file for package installation operations.
	NpmrcEnvironmentVariablesSecretVersionRef map[string]interface{} `json:"npmrcEnvironmentVariablesSecretVersionRef,omitempty" yaml:"npmrcEnvironmentVariablesSecretVersionRef,omitempty"`
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. A reference to the region.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Optional. The service account reference to run workflow invocations under.
	ServiceAccountRef map[string]interface{} `json:"serviceAccountRef,omitempty" yaml:"serviceAccountRef,omitempty"`
	// Optional. Input only. If set to true, the authenticated user will be granted the roles/dataform.admin role on the created repository.
	SetAuthenticatedUserAdmin bool `json:"setAuthenticatedUserAdmin,omitempty" yaml:"setAuthenticatedUserAdmin,omitempty"`
	// Optional. If set, fields of workspaceCompilationOverrides override the default compilation settings that are specified in dataform.json when creating workspace-scoped compilation results.
	WorkspaceCompilationOverrides map[string]interface{} `json:"workspaceCompilationOverrides,omitempty" yaml:"workspaceCompilationOverrides,omitempty"`
}

// DataformRepositoryStatus defines the observed state of DataformRepository.
type DataformRepositoryStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
