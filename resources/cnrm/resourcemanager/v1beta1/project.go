package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Project represents a resourcemanager.cnrm.cloud.google.com Project resource.
type Project struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ProjectSpec   `json:"spec,omitempty"`
	Status ProjectStatus `json:"status,omitempty"`
}

// ProjectSpec defines the desired state of Project.
type ProjectSpec struct {
	BillingAccountRef map[string]interface{} `json:"billingAccountRef,omitempty" yaml:"billingAccountRef,omitempty"`
	// The folder that this resource belongs to. Changing this forces the
	// resource to be migrated to the newly specified folder. Only one of
	// folderRef or organizationRef may be specified.
	FolderRef map[string]interface{} `json:"folderRef,omitempty" yaml:"folderRef,omitempty"`
	// The display name of the project.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	// The organization that this resource belongs to. Changing this
	// forces the resource to be migrated to the newly specified
	// organization. Only one of folderRef or organizationRef may be
	// specified.
	OrganizationRef map[string]interface{} `json:"organizationRef,omitempty" yaml:"organizationRef,omitempty"`
	// Immutable. Optional. The projectId of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// ProjectStatus defines the observed state of Project.
type ProjectStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
