package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// TagsTagValue is the Schema for the TagsTagValue API
type TagsTagValue struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TagsTagValueSpec   `json:"spec,omitempty"`
	Status TagsTagValueStatus `json:"status,omitempty"`
}

// TagsTagValueSpec defines the desired state of TagsTagValue.
type TagsTagValueSpec struct {
	// Optional. User-assigned description of the TagValue.
	// Must not exceed 256 characters.
	// Read-write.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Immutable. The TagValue's parent TagKey.
	ParentRef map[string]interface{} `json:"parentRef,omitempty" yaml:"parentRef,omitempty"`
	// Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Required. Immutable. User-assigned short name for TagValue. The short name
	// should be unique for TagValues within the same parent TagKey.
	// The short name must be 63 characters or less, beginning and ending with
	// an alphanumeric character ([a-z0-9A-Z]) with dashes (-), underscores (_),
	// dots (.), and alphanumerics between.
	ShortName string `json:"shortName,omitempty" yaml:"shortName,omitempty"`
}

// TagsTagValueStatus defines the observed state of TagsTagValue.
type TagsTagValueStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
