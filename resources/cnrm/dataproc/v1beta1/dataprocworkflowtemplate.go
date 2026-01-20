package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DataprocWorkflowTemplate represents a dataproc.cnrm.cloud.google.com DataprocWorkflowTemplate resource.
type DataprocWorkflowTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DataprocWorkflowTemplateSpec   `json:"spec,omitempty"`
	Status DataprocWorkflowTemplateStatus `json:"status,omitempty"`
}

// DataprocWorkflowTemplateSpec defines the desired state of DataprocWorkflowTemplate.
type DataprocWorkflowTemplateSpec struct {
	// Immutable. Optional. Timeout duration for the DAG of jobs, expressed in seconds (see [JSON representation of duration](https://developers.google.com/protocol-buffers/docs/proto3#json)). The timeout duration must be from 10 minutes ("600s") to 24 hours ("86400s"). The timer begins when the first job is submitted. If the workflow is running at the end of the timeout period, any remaining jobs are cancelled, the workflow is ended, and if the workflow was running on a [managed cluster](/dataproc/docs/concepts/workflows/using-workflows#configuring_or_selecting_a_cluster), the cluster is deleted.
	DagTimeout string `json:"dagTimeout,omitempty" yaml:"dagTimeout,omitempty"`
	// Immutable. Required. The Directed Acyclic Graph of Jobs to submit.
	Jobs []map[string]interface{} `json:"jobs,omitempty" yaml:"jobs,omitempty"`
	// Immutable. The location for the resource
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Immutable. Optional. Template parameters whose values are substituted into the template. Values for parameters must be provided when the template is instantiated.
	Parameters []map[string]interface{} `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	// Immutable. Required. WorkflowTemplate scheduling information.
	Placement map[string]interface{} `json:"placement,omitempty" yaml:"placement,omitempty"`
	// Immutable. The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// DataprocWorkflowTemplateStatus defines the observed state of DataprocWorkflowTemplate.
type DataprocWorkflowTemplateStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
