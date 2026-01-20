package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CloudFunctions2Function represents a cloudfunctions2.cnrm.cloud.google.com CloudFunctions2Function resource.
type CloudFunctions2Function struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CloudFunctions2FunctionSpec   `json:"spec,omitempty"`
	Status CloudFunctions2FunctionStatus `json:"status,omitempty"`
}

// CloudFunctions2FunctionSpec defines the desired state of CloudFunctions2Function.
type CloudFunctions2FunctionSpec struct {
	// Describes the Build step of the function that builds a container
	// from the given source.
	BuildConfig map[string]interface{} `json:"buildConfig,omitempty" yaml:"buildConfig,omitempty"`
	// User-provided description of a function.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// An Eventarc trigger managed by Google Cloud Functions that fires events in
	// response to a condition in another service.
	EventTrigger map[string]interface{} `json:"eventTrigger,omitempty" yaml:"eventTrigger,omitempty"`
	// Resource name of a KMS crypto key (managed by the user) used to encrypt/decrypt function resources.
	// It must match the pattern projects/{project}/locations/{location}/keyRings/{key_ring}/cryptoKeys/{crypto_key}.
	KmsKeyName string `json:"kmsKeyName,omitempty" yaml:"kmsKeyName,omitempty"`
	// Immutable. The location of this cloud function.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Describes the Service being deployed.
	ServiceConfig map[string]interface{} `json:"serviceConfig,omitempty" yaml:"serviceConfig,omitempty"`
}

// CloudFunctions2FunctionStatus defines the observed state of CloudFunctions2Function.
type CloudFunctions2FunctionStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
