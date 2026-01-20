package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Workflow is the Schema for the Workflow API
type WorkflowsWorkflow struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WorkflowsWorkflowSpec   `json:"spec,omitempty"`
	Status WorkflowsWorkflowStatus `json:"status,omitempty"`
}

// WorkflowsWorkflowSpec defines the desired state of WorkflowsWorkflow.
type WorkflowsWorkflowSpec struct {
	// Optional. Describes the level of platform logging to apply to calls and call responses during executions of this workflow. If both the workflow and the execution specify a logging level, the execution level takes precedence.
	CallLogLevel string `json:"callLogLevel,omitempty" yaml:"callLogLevel,omitempty"`
	// Description of the workflow provided by the user. Must be at most 1000 unicode characters long.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Optional. The resource name of a KMS crypto key used to encrypt or decrypt the data associated with the workflow. If not provided, data associated with the workflow will not be CMEK-encrypted.
	KmsCryptoKeyRef map[string]interface{} `json:"kmsCryptoKeyRef,omitempty" yaml:"kmsCryptoKeyRef,omitempty"`
	// Labels associated with this workflow. Labels can contain at most 64 entries. Keys and values can be no longer than 63 characters and can only contain lowercase letters, numeric characters, underscores, and dashes. Label keys must start with a letter. International characters are allowed.
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	// Immutable.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// The Workflow name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// The service account associated with the latest workflow version. This service account represents the identity of the workflow and determines what permissions the workflow has. If not provided, workflow will use the project's default service account. Modifying this field for an existing workflow results in a new workflow revision.
	ServiceAccountRef map[string]interface{} `json:"serviceAccountRef,omitempty" yaml:"serviceAccountRef,omitempty"`
	// Required. Workflow code to be executed. The size limit is 128KB.
	SourceContents string `json:"sourceContents,omitempty" yaml:"sourceContents,omitempty"`
	// Optional.User-defined environment variables associated with this workflow revision. This map has a maximum length of 20. Each string can take up to 40KiB. Keys cannot be empty strings and cannot start with “GOOGLE” or “WORKFLOWS".
	UserEnvVars map[string]string `json:"userEnvVars,omitempty" yaml:"userEnvVars,omitempty"`
}

// WorkflowsWorkflowStatus defines the observed state of WorkflowsWorkflow.
type WorkflowsWorkflowStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
