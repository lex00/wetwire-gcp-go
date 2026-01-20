package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ComputeNodeTemplate represents a compute.cnrm.cloud.google.com ComputeNodeTemplate resource.
type ComputeNodeTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeNodeTemplateSpec   `json:"spec,omitempty"`
	Status ComputeNodeTemplateStatus `json:"status,omitempty"`
}

// ComputeNodeTemplateSpec defines the desired state of ComputeNodeTemplate.
type ComputeNodeTemplateSpec struct {
	// Immutable. CPU overcommit. Default value: "NONE" Possible values: ["ENABLED", "NONE"].
	CPUOvercommitType string `json:"cpuOvercommitType,omitempty" yaml:"cpuOvercommitType,omitempty"`
	// Immutable. An optional textual description of the resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Immutable. Node type to use for nodes group that are created from this template.
	// Only one of nodeTypeFlexibility and nodeType can be specified.
	NodeType string `json:"nodeType,omitempty" yaml:"nodeType,omitempty"`
	// Immutable. Flexible properties for the desired node type. Node groups that
	// use this node template will create nodes of a type that matches
	// these properties. Only one of nodeTypeFlexibility and nodeType can
	// be specified.
	NodeTypeFlexibility map[string]interface{} `json:"nodeTypeFlexibility,omitempty" yaml:"nodeTypeFlexibility,omitempty"`
	// Immutable. Region where nodes using the node template will be created.
	// If it is not provided, the provider region is used.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Immutable. The server binding policy for nodes using this template. Determines
	// where the nodes should restart following a maintenance event.
	ServerBinding map[string]interface{} `json:"serverBinding,omitempty" yaml:"serverBinding,omitempty"`
}

// ComputeNodeTemplateStatus defines the observed state of ComputeNodeTemplate.
type ComputeNodeTemplateStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
