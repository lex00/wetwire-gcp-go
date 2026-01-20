package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// IdentityPlatformProjectDefaultConfig represents a identityplatform.cnrm.cloud.google.com IdentityPlatformProjectDefaultConfig resource.
type IdentityPlatformProjectDefaultConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IdentityPlatformProjectDefaultConfigSpec   `json:"spec,omitempty"`
	Status IdentityPlatformProjectDefaultConfigStatus `json:"status,omitempty"`
}

// IdentityPlatformProjectDefaultConfigSpec defines the desired state of IdentityPlatformProjectDefaultConfig.
type IdentityPlatformProjectDefaultConfigSpec struct {
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Configuration related to local sign in methods.
	SignIn map[string]interface{} `json:"signIn,omitempty" yaml:"signIn,omitempty"`
}

// IdentityPlatformProjectDefaultConfigStatus defines the observed state of IdentityPlatformProjectDefaultConfig.
type IdentityPlatformProjectDefaultConfigStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
