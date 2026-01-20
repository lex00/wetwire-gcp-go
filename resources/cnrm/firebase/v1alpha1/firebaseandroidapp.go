package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// FirebaseAndroidApp represents a firebase.cnrm.cloud.google.com FirebaseAndroidApp resource.
type FirebaseAndroidApp struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FirebaseAndroidAppSpec   `json:"spec,omitempty"`
	Status FirebaseAndroidAppStatus `json:"status,omitempty"`
}

// FirebaseAndroidAppSpec defines the desired state of FirebaseAndroidApp.
type FirebaseAndroidAppSpec struct {
	// The globally unique, Google-assigned identifier (UID) for the Firebase API key associated with the AndroidApp.
	// If apiKeyId is not set during creation, then Firebase automatically associates an apiKeyId with the AndroidApp.
	// This auto-associated key may be an existing valid key or, if no valid key exists, a new one will be provisioned.
	APIKeyID string `json:"apiKeyId,omitempty" yaml:"apiKeyId,omitempty"`
	DeletionPolicy string `json:"deletionPolicy,omitempty" yaml:"deletionPolicy,omitempty"`
	// The user-assigned display name of the AndroidApp.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// Immutable. The canonical package name of the Android app as would appear in the Google Play
	// Developer Console.
	PackageName string `json:"packageName,omitempty" yaml:"packageName,omitempty"`
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The service-generated appId of the resource. Used for acquisition only. Leave unset to create a new resource.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// The SHA1 certificate hashes for the AndroidApp.
	Sha1Hashes []string `json:"sha1Hashes,omitempty" yaml:"sha1Hashes,omitempty"`
	// The SHA256 certificate hashes for the AndroidApp.
	Sha256Hashes []string `json:"sha256Hashes,omitempty" yaml:"sha256Hashes,omitempty"`
}

// FirebaseAndroidAppStatus defines the observed state of FirebaseAndroidApp.
type FirebaseAndroidAppStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
