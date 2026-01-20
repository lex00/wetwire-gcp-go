package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// TagsTagBinding is the Schema for the TagsTagBinding API
type TagsTagBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TagsTagBindingSpec   `json:"spec,omitempty"`
	Status TagsTagBindingStatus `json:"status,omitempty"`
}

// TagsTagBindingSpec defines the desired state of TagsTagBinding.
type TagsTagBindingSpec struct {
	// ParentRef is a reference to a parent resource.
	ParentRef map[string]interface{} `json:"parentRef,omitempty" yaml:"parentRef,omitempty"`
	// The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// TagsTagValueRef is a reference to a TagsTagValue resource.
	TagValueRef map[string]interface{} `json:"tagValueRef,omitempty" yaml:"tagValueRef,omitempty"`
}

// TagsTagBindingStatus defines the observed state of TagsTagBinding.
type TagsTagBindingStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
