package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EdgeContainerMachine is the Schema for the EdgeContainerMachine API
type EdgeContainerMachine struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EdgeContainerMachineSpec   `json:"spec,omitempty"`
	Status EdgeContainerMachineStatus `json:"status,omitempty"`
}

// EdgeContainerMachineSpec defines the desired state of EdgeContainerMachine.
type EdgeContainerMachineSpec struct {
	// Labels associated with this resource.
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	// Required. The location of the machine.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Required. The host project of the machine.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// The EdgeContainerMachine name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// The Google Distributed Cloud Edge zone of this machine.
	Zone string `json:"zone,omitempty" yaml:"zone,omitempty"`
}

// EdgeContainerMachineStatus defines the observed state of EdgeContainerMachine.
type EdgeContainerMachineStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
