package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ApigeeSyncAuthorization represents a apigee.cnrm.cloud.google.com ApigeeSyncAuthorization resource.
type ApigeeSyncAuthorization struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ApigeeSyncAuthorizationSpec   `json:"spec,omitempty"`
	Status ApigeeSyncAuthorizationStatus `json:"status,omitempty"`
}

// ApigeeSyncAuthorizationSpec defines the desired state of ApigeeSyncAuthorization.
type ApigeeSyncAuthorizationSpec struct {
	// Array of service accounts to grant access to control plane resources, each specified using the following format: 'serviceAccount:service-account-name'.
	// The 'service-account-name' is formatted like an email address. For example: my-synchronizer-manager-serviceAccount@my_project_id.iam.gserviceaccount.com
	// You might specify multiple service accounts, for example, if you have multiple environments and wish to assign a unique service account to each one.
	// The service accounts must have **Apigee Synchronizer Manager** role. See also [Create service accounts](https://cloud.google.com/apigee/docs/hybrid/v1.8/sa-about#create-the-service-accounts).
	Identities []string `json:"identities,omitempty" yaml:"identities,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// ApigeeSyncAuthorizationStatus defines the observed state of ApigeeSyncAuthorization.
type ApigeeSyncAuthorizationStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
