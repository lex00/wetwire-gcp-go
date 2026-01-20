package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// FirebaseHostingSite represents a firebasehosting.cnrm.cloud.google.com FirebaseHostingSite resource.
type FirebaseHostingSite struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FirebaseHostingSiteSpec   `json:"spec,omitempty"`
	Status FirebaseHostingSiteStatus `json:"status,omitempty"`
}

// FirebaseHostingSiteSpec defines the desired state of FirebaseHostingSite.
type FirebaseHostingSiteSpec struct {
	// Optional. The [ID of a Web App](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects.webApps#WebApp.FIELDS.app_id)
	// associated with the Hosting site.
	AppID string `json:"appId,omitempty" yaml:"appId,omitempty"`
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The siteId of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// FirebaseHostingSiteStatus defines the observed state of FirebaseHostingSite.
type FirebaseHostingSiteStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
