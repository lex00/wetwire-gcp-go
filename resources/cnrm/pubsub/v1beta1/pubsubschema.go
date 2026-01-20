package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PubSubSchema represents a pubsub.cnrm.cloud.google.com PubSubSchema resource.
type PubSubSchema struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PubSubSchemaSpec   `json:"spec,omitempty"`
	Status PubSubSchemaStatus `json:"status,omitempty"`
}

// PubSubSchemaSpec defines the desired state of PubSubSchema.
type PubSubSchemaSpec struct {
	// The definition of the schema.
	// This should contain a string representing the full definition of the schema
	// that is a valid schema definition of the type specified in type.
	Definition string `json:"definition,omitempty" yaml:"definition,omitempty"`
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// The type of the schema definition Default value: "TYPE_UNSPECIFIED" Possible values: ["TYPE_UNSPECIFIED", "PROTOCOL_BUFFER", "AVRO"].
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}

// PubSubSchemaStatus defines the observed state of PubSubSchema.
type PubSubSchemaStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
