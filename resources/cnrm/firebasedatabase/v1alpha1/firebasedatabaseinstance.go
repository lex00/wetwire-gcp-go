package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// FirebaseDatabaseInstance represents a firebasedatabase.cnrm.cloud.google.com FirebaseDatabaseInstance resource.
type FirebaseDatabaseInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FirebaseDatabaseInstanceSpec   `json:"spec,omitempty"`
	Status FirebaseDatabaseInstanceStatus `json:"status,omitempty"`
}

// FirebaseDatabaseInstanceSpec defines the desired state of FirebaseDatabaseInstance.
type FirebaseDatabaseInstanceSpec struct {
	// The intended database state.
	DesiredState string `json:"desiredState,omitempty" yaml:"desiredState,omitempty"`
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. A reference to the region where the Firebase Realtime database resides.
	// Check all [available regions](https://firebase.google.com/docs/projects/locations#rtdb-locations).
	Region string `json:"region,omitempty" yaml:"region,omitempty"`
	// Immutable. Optional. The instanceId of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Immutable. The database type.
	// Each project can create one default Firebase Realtime Database, which cannot be deleted once created.
	// Creating user Databases is only available for projects on the Blaze plan.
	// Projects can be upgraded using the Cloud Billing API https://cloud.google.com/billing/reference/rest/v1/projects/updateBillingInfo. Default value: "USER_DATABASE" Possible values: ["DEFAULT_DATABASE", "USER_DATABASE"].
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}

// FirebaseDatabaseInstanceStatus defines the observed state of FirebaseDatabaseInstance.
type FirebaseDatabaseInstanceStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
