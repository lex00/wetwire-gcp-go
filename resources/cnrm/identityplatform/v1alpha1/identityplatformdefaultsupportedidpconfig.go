package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// IdentityPlatformDefaultSupportedIDPConfig represents a identityplatform.cnrm.cloud.google.com IdentityPlatformDefaultSupportedIDPConfig resource.
type IdentityPlatformDefaultSupportedIDPConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IdentityPlatformDefaultSupportedIDPConfigSpec   `json:"spec,omitempty"`
	Status IdentityPlatformDefaultSupportedIDPConfigStatus `json:"status,omitempty"`
}

// IdentityPlatformDefaultSupportedIDPConfigSpec defines the desired state of IdentityPlatformDefaultSupportedIDPConfig.
type IdentityPlatformDefaultSupportedIDPConfigSpec struct {
	// OAuth client ID.
	ClientID string `json:"clientId,omitempty" yaml:"clientId,omitempty"`
	// OAuth client secret.
	ClientSecret string `json:"clientSecret,omitempty" yaml:"clientSecret,omitempty"`
	// If this IDP allows the user to sign in.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The idpId of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// IdentityPlatformDefaultSupportedIDPConfigStatus defines the observed state of IdentityPlatformDefaultSupportedIDPConfig.
type IdentityPlatformDefaultSupportedIDPConfigStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
