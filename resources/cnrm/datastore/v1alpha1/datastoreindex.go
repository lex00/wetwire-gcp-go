package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DatastoreIndex represents a datastore.cnrm.cloud.google.com DatastoreIndex resource.
type DatastoreIndex struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DatastoreIndexSpec   `json:"spec,omitempty"`
	Status DatastoreIndexStatus `json:"status,omitempty"`
}

// DatastoreIndexSpec defines the desired state of DatastoreIndex.
type DatastoreIndexSpec struct {
	// Immutable. Policy for including ancestors in the index. Default value: "NONE" Possible values: ["NONE", "ALL_ANCESTORS"].
	Ancestor string `json:"ancestor,omitempty" yaml:"ancestor,omitempty"`
	// Immutable. The entity kind which the index applies to.
	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. An ordered list of properties to index on.
	Properties []map[string]interface{} `json:"properties,omitempty" yaml:"properties,omitempty"`
	// Immutable. Optional. The service-generated indexId of the resource. Used for acquisition only. Leave unset to create a new resource.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// DatastoreIndexStatus defines the observed state of DatastoreIndex.
type DatastoreIndexStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
