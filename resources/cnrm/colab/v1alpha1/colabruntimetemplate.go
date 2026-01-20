package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ColabRuntimeTemplate is the Schema for the ColabRuntimeTemplate API
type ColabRuntimeTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ColabRuntimeTemplateSpec   `json:"spec,omitempty"`
	Status ColabRuntimeTemplateStatus `json:"status,omitempty"`
}

// ColabRuntimeTemplateSpec defines the desired state of ColabRuntimeTemplate.
type ColabRuntimeTemplateSpec struct {
	// Optional. The specification of [persistent disk][https://cloud.google.com/compute/docs/disks/persistent-disks] attached to the runtime as data disk storage.
	DataPersistentDiskSpec map[string]interface{} `json:"dataPersistentDiskSpec,omitempty" yaml:"dataPersistentDiskSpec,omitempty"`
	// The description of the NotebookRuntimeTemplate.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Required. The display name of the NotebookRuntimeTemplate. The name can be up to 128 characters long and can consist of any UTF-8 characters.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// Customer-managed encryption key spec for the notebook runtime.
	EncryptionSpec map[string]interface{} `json:"encryptionSpec,omitempty" yaml:"encryptionSpec,omitempty"`
	// EUC configuration of the NotebookRuntimeTemplate.
	EucConfig map[string]interface{} `json:"eucConfig,omitempty" yaml:"eucConfig,omitempty"`
	// The idle shutdown configuration of NotebookRuntimeTemplate. This config will only be set when idle shutdown is enabled.
	IdleShutdownConfig map[string]interface{} `json:"idleShutdownConfig,omitempty" yaml:"idleShutdownConfig,omitempty"`
	// The labels with user-defined metadata to organize the
	// NotebookRuntimeTemplates.
	// Label keys and values can be no longer than 64 characters
	// (Unicode codepoints), can only contain lowercase letters, numeric
	// characters, underscores and dashes. International characters are allowed.
	// See https://goo.gl/xmQnxf for more information and examples of labels.
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	// Immutable. The name of the location where the RuntimeTemplate will be created. Required.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Optional. Immutable. The specification of a single machine for the template.
	MachineSpec map[string]interface{} `json:"machineSpec,omitempty" yaml:"machineSpec,omitempty"`
	// Optional. Network spec.
	NetworkSpec map[string]interface{} `json:"networkSpec,omitempty" yaml:"networkSpec,omitempty"`
	// Optional. The Compute Engine tags to add to runtime (see [Tagging instances](https://cloud.google.com/vpc/docs/add-remove-network-tags)).
	NetworkTags []string `json:"networkTags,omitempty" yaml:"networkTags,omitempty"`
	// Optional. Immutable. The type of the notebook runtime template.
	NotebookRuntimeType string `json:"notebookRuntimeType,omitempty" yaml:"notebookRuntimeType,omitempty"`
	// Immutable. The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// The ColabRuntimeTemplate name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// The service account that the runtime workload runs as.
	// You can use any service account within the same project, but you
	// must have the service account user permission to use the instance.
	// If not specified, the [Compute Engine default service
	// account](https://cloud.google.com/compute/docs/access/service-accounts#default_service_account)
	// is used.
	ServiceAccountRef map[string]interface{} `json:"serviceAccountRef,omitempty" yaml:"serviceAccountRef,omitempty"`
	// Optional. Immutable. Runtime Shielded VM spec.
	ShieldedVMConfig map[string]interface{} `json:"shieldedVMConfig,omitempty" yaml:"shieldedVMConfig,omitempty"`
}

// ColabRuntimeTemplateStatus defines the observed state of ColabRuntimeTemplate.
type ColabRuntimeTemplateStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
