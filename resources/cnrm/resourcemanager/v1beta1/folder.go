package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Folder represents a resourcemanager.cnrm.cloud.google.com Folder resource.
type Folder struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FolderSpec   `json:"spec,omitempty"`
	Status FolderStatus `json:"status,omitempty"`
}

// FolderSpec defines the desired state of Folder.
type FolderSpec struct {
	// The folder's display name. A folder's display name must be unique amongst its siblings, e.g. no two folders with the same parent can share the same display name. The display name must start and end with a letter or digit, may contain letters, digits, spaces, hyphens and underscores and can be no longer than 30 characters.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// The folder that this resource belongs to. Changing this forces the
	// resource to be migrated to the newly specified folder. Only one of
	// folderRef or organizationRef may be specified.
	FolderRef map[string]interface{} `json:"folderRef,omitempty" yaml:"folderRef,omitempty"`
	// The organization that this resource belongs to. Changing this
	// forces the resource to be migrated to the newly specified
	// organization. Only one of folderRef or organizationRef may be
	// specified.
	OrganizationRef map[string]interface{} `json:"organizationRef,omitempty" yaml:"organizationRef,omitempty"`
	// Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// FolderStatus defines the observed state of Folder.
type FolderStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
