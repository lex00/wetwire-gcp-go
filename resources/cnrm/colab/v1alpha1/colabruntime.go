package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ColabRuntime is the Schema for the ColabRuntime API
type ColabRuntime struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ColabRuntimeSpec   `json:"spec,omitempty"`
	Status ColabRuntimeStatus `json:"status,omitempty"`
}

// ColabRuntimeSpec defines the desired state of ColabRuntime.
type ColabRuntimeSpec struct {
	// The pointer to NotebookRuntimeTemplate this NotebookRuntime is created from.
	ColabRuntimeTemplateRef map[string]interface{} `json:"colabRuntimeTemplateRef,omitempty" yaml:"colabRuntimeTemplateRef,omitempty"`
	// The description of the NotebookRuntime.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Required. The display name of the NotebookRuntime. The name can be up to 128 characters long and can consist of any UTF-8 characters.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// The labels with user-defined metadata to organize your
	// NotebookRuntime.
	// Label keys and values can be no longer than 64 characters
	// (Unicode codepoints), can only contain lowercase letters, numeric
	// characters, underscores and dashes. International characters are allowed.
	// No more than 64 user labels can be associated with one NotebookRuntime
	// (System labels are excluded).
	// See https://goo.gl/xmQnxf for more information and examples of labels.
	// System reserved label keys are prefixed with "aiplatform.googleapis.com/"
	// and are immutable. Following system labels exist for NotebookRuntime:
	// * "aiplatform.googleapis.com/notebook_runtime_gce_instance_id": output
	// only, its value is the Compute Engine instance id.
	// * "aiplatform.googleapis.com/colab_enterprise_entry_service": its value is
	// either "bigquery" or "vertex"; if absent, it should be "vertex". This is to
	// describe the entry service, either BigQuery or Vertex.
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	// Immutable. The name of the location where the Runtime will be created. Required.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Optional. The Compute Engine tags to add to runtime (see [Tagging instances](https://cloud.google.com/vpc/docs/add-remove-network-tags)).
	NetworkTags []string `json:"networkTags,omitempty" yaml:"networkTags,omitempty"`
	// Immutable. The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// The ColabRuntime name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Required. The user email of the NotebookRuntime.
	RuntimeUser string `json:"runtimeUser,omitempty" yaml:"runtimeUser,omitempty"`
}

// ColabRuntimeStatus defines the observed state of ColabRuntime.
type ColabRuntimeStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
