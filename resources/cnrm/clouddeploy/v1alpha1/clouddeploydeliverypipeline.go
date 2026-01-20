package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CloudDeployDeliveryPipeline is the Schema for the CloudDeployDeliveryPipeline API
type CloudDeployDeliveryPipeline struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CloudDeployDeliveryPipelineSpec   `json:"spec,omitempty"`
	Status CloudDeployDeliveryPipelineStatus `json:"status,omitempty"`
}

// CloudDeployDeliveryPipelineSpec defines the desired state of CloudDeployDeliveryPipeline.
type CloudDeployDeliveryPipelineSpec struct {
	// User annotations. These attributes can only be set and used by the user, and not by Cloud Deploy.
	Annotations map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	// Description of the `DeliveryPipeline`. Max length is 255 characters.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Immutable. The location where the DeliveryPipeline should reside.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// The GCP resource identifier. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// SerialPipeline defines a sequential set of stages for a `DeliveryPipeline`.
	SerialPipeline map[string]interface{} `json:"serialPipeline,omitempty" yaml:"serialPipeline,omitempty"`
	// When suspended, no new releases or rollouts can be created, but in-progress ones will complete.
	Suspended bool `json:"suspended,omitempty" yaml:"suspended,omitempty"`
}

// CloudDeployDeliveryPipelineStatus defines the observed state of CloudDeployDeliveryPipeline.
type CloudDeployDeliveryPipelineStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
