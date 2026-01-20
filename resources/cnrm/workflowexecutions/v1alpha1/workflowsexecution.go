package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// WorkflowsExecution is the Schema for the WorkflowsExecution API
type WorkflowsExecution struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WorkflowsExecutionSpec   `json:"spec,omitempty"`
	Status WorkflowsExecutionStatus `json:"status,omitempty"`
}

// WorkflowsExecutionSpec defines the desired state of WorkflowsExecution.
type WorkflowsExecutionSpec struct {
	// Input parameters of the execution represented as a JSON string.
	// The size limit is 32KB.
	// *Note*: If you are using the REST API directly to run your workflow, you
	// must escape any JSON string value of `argument`. Example:
	// `'{"argument":"{\"firstName\":\"FIRST\",\"lastName\":\"LAST\"}"}'`
	Argument string `json:"argument,omitempty" yaml:"argument,omitempty"`
	// The call logging level associated to this execution.
	CallLogLevel string `json:"callLogLevel,omitempty" yaml:"callLogLevel,omitempty"`
	// Labels associated with this execution. Labels can contain at most 64 entries. Keys and values can be no longer than 63 characters and can only contain lowercase letters, numeric characters, underscores, and dashes. Label keys must start with a letter. International characters are allowed. By default, labels are inherited from the workflow but are overridden by any labels associated with the execution.
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	// Required. The location of the application.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Required. The host project of the application.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// The WorkflowsExecution name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Required.
	WorkflowRef map[string]interface{} `json:"workflowRef,omitempty" yaml:"workflowRef,omitempty"`
}

// WorkflowsExecutionStatus defines the observed state of WorkflowsExecution.
type WorkflowsExecutionStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
