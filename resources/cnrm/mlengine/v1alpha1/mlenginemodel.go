package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MLEngineModel represents a mlengine.cnrm.cloud.google.com MLEngineModel resource.
type MLEngineModel struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MLEngineModelSpec   `json:"spec,omitempty"`
	Status MLEngineModelStatus `json:"status,omitempty"`
}

// MLEngineModelSpec defines the desired state of MLEngineModel.
type MLEngineModelSpec struct {
	// Immutable. The default version of the model. This version will be used to handle
	// prediction requests that do not specify a version.
	DefaultVersion map[string]interface{} `json:"defaultVersion,omitempty" yaml:"defaultVersion,omitempty"`
	// Immutable. The description specified for the model when it was created.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Immutable. If true, online prediction nodes send stderr and stdout streams to Stackdriver Logging.
	OnlinePredictionConsoleLogging bool `json:"onlinePredictionConsoleLogging,omitempty" yaml:"onlinePredictionConsoleLogging,omitempty"`
	// Immutable. If true, online prediction access logs are sent to StackDriver Logging.
	OnlinePredictionLogging bool `json:"onlinePredictionLogging,omitempty" yaml:"onlinePredictionLogging,omitempty"`
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. The list of regions where the model is going to be deployed.
	// Currently only one region per model is supported.
	Regions []string `json:"regions,omitempty" yaml:"regions,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// MLEngineModelStatus defines the observed state of MLEngineModel.
type MLEngineModelStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
