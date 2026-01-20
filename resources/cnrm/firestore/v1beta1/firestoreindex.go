package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// FirestoreIndex is the Schema for the FirestoreIndex API
type FirestoreIndex struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FirestoreIndexSpec   `json:"spec,omitempty"`
	Status FirestoreIndexStatus `json:"status,omitempty"`
}

// FirestoreIndexSpec defines the desired state of FirestoreIndex.
type FirestoreIndexSpec struct {
	// Immutable. The collection being indexed.
	Collection string `json:"collection,omitempty" yaml:"collection,omitempty"`
	// Immutable. The Firestore database id. Defaults to '"(default)"'.
	Database string `json:"database,omitempty" yaml:"database,omitempty"`
	// Immutable. The fields supported by this index. The last field entry is always for the field path '__name__'. If, on creation, '__name__' was not specified as the last field, it will be added automatically with the same direction as that of the last field defined. If the final field in a composite index is not directional, the '__name__' will be ordered '"ASCENDING"' (unless explicitly specified otherwise).
	Fields []map[string]interface{} `json:"fields,omitempty" yaml:"fields,omitempty"`
	// Immutable. The scope at which a query is run. Default value: "COLLECTION" Possible values: ["COLLECTION", "COLLECTION_GROUP"].
	QueryScope string `json:"queryScope,omitempty" yaml:"queryScope,omitempty"`
}

// FirestoreIndexStatus defines the observed state of FirestoreIndex.
type FirestoreIndexStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
