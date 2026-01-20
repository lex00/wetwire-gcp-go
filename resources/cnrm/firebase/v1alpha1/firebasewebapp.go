package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// FirebaseWebApp represents a firebase.cnrm.cloud.google.com FirebaseWebApp resource.
type FirebaseWebApp struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FirebaseWebAppSpec   `json:"spec,omitempty"`
	Status FirebaseWebAppStatus `json:"status,omitempty"`
}

// FirebaseWebAppSpec defines the desired state of FirebaseWebApp.
type FirebaseWebAppSpec struct {
	// The globally unique, Google-assigned identifier (UID) for the Firebase API key associated with the WebApp.
	// If apiKeyId is not set during creation, then Firebase automatically associates an apiKeyId with the WebApp.
	// This auto-associated key may be an existing valid key or, if no valid key exists, a new one will be provisioned.
	APIKeyID string `json:"apiKeyId,omitempty" yaml:"apiKeyId,omitempty"`
	DeletionPolicy string `json:"deletionPolicy,omitempty" yaml:"deletionPolicy,omitempty"`
	// The user-assigned display name of the App.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// Immutable.
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
	// Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// FirebaseWebAppStatus defines the observed state of FirebaseWebApp.
type FirebaseWebAppStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
