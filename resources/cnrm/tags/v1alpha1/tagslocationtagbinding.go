package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// TagsLocationTagBinding is the Schema for the TagsLocationTagBinding API
type TagsLocationTagBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TagsLocationTagBindingSpec   `json:"spec,omitempty"`
	Status TagsLocationTagBindingStatus `json:"status,omitempty"`
}

// TagsLocationTagBindingSpec defines the desired state of TagsLocationTagBinding.
type TagsLocationTagBindingSpec struct {
	// The location for the resource being tagged.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// ParentRef is a reference to a parent resource.
	ParentRef map[string]interface{} `json:"parentRef,omitempty" yaml:"parentRef,omitempty"`
	// The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// TagsTagValueRef is a reference to a TagsTagValue resource.
	TagValueRef map[string]interface{} `json:"tagValueRef,omitempty" yaml:"tagValueRef,omitempty"`
}

// TagsLocationTagBindingStatus defines the observed state of TagsLocationTagBinding.
type TagsLocationTagBindingStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
