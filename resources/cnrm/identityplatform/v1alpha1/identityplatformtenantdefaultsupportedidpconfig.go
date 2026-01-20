package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// IdentityPlatformTenantDefaultSupportedIDPConfig represents a identityplatform.cnrm.cloud.google.com IdentityPlatformTenantDefaultSupportedIDPConfig resource.
type IdentityPlatformTenantDefaultSupportedIDPConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IdentityPlatformTenantDefaultSupportedIDPConfigSpec   `json:"spec,omitempty"`
	Status IdentityPlatformTenantDefaultSupportedIDPConfigStatus `json:"status,omitempty"`
}

// IdentityPlatformTenantDefaultSupportedIDPConfigSpec defines the desired state of IdentityPlatformTenantDefaultSupportedIDPConfig.
type IdentityPlatformTenantDefaultSupportedIDPConfigSpec struct {
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
	// Immutable. The name of the tenant where this DefaultSupportedIdpConfig resource exists.
	Tenant string `json:"tenant,omitempty" yaml:"tenant,omitempty"`
}

// IdentityPlatformTenantDefaultSupportedIDPConfigStatus defines the observed state of IdentityPlatformTenantDefaultSupportedIDPConfig.
type IdentityPlatformTenantDefaultSupportedIDPConfigStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
