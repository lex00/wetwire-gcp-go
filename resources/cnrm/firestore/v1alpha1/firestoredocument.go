package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// FirestoreDocument is the Schema for the FirestoreDocument API
type FirestoreDocument struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FirestoreDocumentSpec   `json:"spec,omitempty"`
	Status FirestoreDocumentStatus `json:"status,omitempty"`
}

// FirestoreDocumentSpec defines the desired state of FirestoreDocument.
type FirestoreDocumentSpec struct {
	// Collection is the identity of the firestore collection in which to create the document.
	Collection string `json:"collection,omitempty" yaml:"collection,omitempty"`
	// DatabaseRef references the FirestoreDatabase in which to create the document.
	DatabaseRef map[string]interface{} `json:"databaseRef,omitempty" yaml:"databaseRef,omitempty"`
	// Fields holds the field values; values follow JSON typing conventions.
	Fields map[string]interface{} `json:"fields,omitempty" yaml:"fields,omitempty"`
	// The FirestoreDocument name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// FirestoreDocumentStatus defines the observed state of FirestoreDocument.
type FirestoreDocumentStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
