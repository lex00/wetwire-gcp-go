package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ArtifactRegistryRepository represents a artifactregistry.cnrm.cloud.google.com ArtifactRegistryRepository resource.
type ArtifactRegistryRepository struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ArtifactRegistryRepositorySpec   `json:"spec,omitempty"`
	Status ArtifactRegistryRepositoryStatus `json:"status,omitempty"`
}

// ArtifactRegistryRepositorySpec defines the desired state of ArtifactRegistryRepository.
type ArtifactRegistryRepositorySpec struct {
	// Cleanup policies for this repository. Cleanup policies indicate when
	// certain package versions can be automatically deleted.
	// Map keys are policy IDs supplied by users during policy creation. They must
	// unique within a repository and be under 128 characters in length.
	CleanupPolicies []map[string]interface{} `json:"cleanupPolicies,omitempty" yaml:"cleanupPolicies,omitempty"`
	// If true, the cleanup pipeline is prevented from deleting versions in this
	// repository.
	CleanupPolicyDryRun bool `json:"cleanupPolicyDryRun,omitempty" yaml:"cleanupPolicyDryRun,omitempty"`
	// The user-provided description of the repository.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Docker repository config contains repository level configuration for the repositories of docker type.
	DockerConfig map[string]interface{} `json:"dockerConfig,omitempty" yaml:"dockerConfig,omitempty"`
	// Immutable. The format of packages that are stored in the repository. Supported formats
	// can be found [here](https://cloud.google.com/artifact-registry/docs/supported-formats).
	// You can only create alpha formats if you are a member of the
	// [alpha user group](https://cloud.google.com/artifact-registry/docs/supported-formats#alpha-access).
	Format string `json:"format,omitempty" yaml:"format,omitempty"`
	// The customer managed encryption key that’s used to encrypt the
	// contents of the Repository.
	KmsKeyRef map[string]interface{} `json:"kmsKeyRef,omitempty" yaml:"kmsKeyRef,omitempty"`
	// Immutable. The name of the location this repository is located in.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// MavenRepositoryConfig is maven related repository details.
	// Provides additional configuration details for repositories of the maven
	// format type.
	MavenConfig map[string]interface{} `json:"mavenConfig,omitempty" yaml:"mavenConfig,omitempty"`
	// Immutable. The mode configures the repository to serve artifacts from different sources. Default value: "STANDARD_REPOSITORY" Possible values: ["STANDARD_REPOSITORY", "VIRTUAL_REPOSITORY", "REMOTE_REPOSITORY"].
	Mode string `json:"mode,omitempty" yaml:"mode,omitempty"`
	// Immutable. Configuration specific for a Remote Repository.
	RemoteRepositoryConfig map[string]interface{} `json:"remoteRepositoryConfig,omitempty" yaml:"remoteRepositoryConfig,omitempty"`
	// Immutable. Optional. The repositoryId of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Configuration specific for a Virtual Repository.
	VirtualRepositoryConfig map[string]interface{} `json:"virtualRepositoryConfig,omitempty" yaml:"virtualRepositoryConfig,omitempty"`
}

// ArtifactRegistryRepositoryStatus defines the observed state of ArtifactRegistryRepository.
type ArtifactRegistryRepositoryStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
