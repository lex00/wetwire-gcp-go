package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ConfigControllerInstance represents a configcontroller.cnrm.cloud.google.com ConfigControllerInstance resource.
type ConfigControllerInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ConfigControllerInstanceSpec   `json:"spec,omitempty"`
	Status ConfigControllerInstanceStatus `json:"status,omitempty"`
}

// ConfigControllerInstanceSpec defines the desired state of ConfigControllerInstance.
type ConfigControllerInstanceSpec struct {
	// Immutable. The location for the resource
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Immutable. Configuration of the cluster management
	ManagementConfig map[string]interface{} `json:"managementConfig,omitempty" yaml:"managementConfig,omitempty"`
	// Immutable. The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Immutable. Only allow access to the master's private endpoint IP.
	UsePrivateEndpoint bool `json:"usePrivateEndpoint,omitempty" yaml:"usePrivateEndpoint,omitempty"`
}

// ConfigControllerInstanceStatus defines the observed state of ConfigControllerInstance.
type ConfigControllerInstanceStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
