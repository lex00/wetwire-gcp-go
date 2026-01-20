package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// IdentityPlatformTenantInboundSAMLConfig represents a identityplatform.cnrm.cloud.google.com IdentityPlatformTenantInboundSAMLConfig resource.
type IdentityPlatformTenantInboundSAMLConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IdentityPlatformTenantInboundSAMLConfigSpec   `json:"spec,omitempty"`
	Status IdentityPlatformTenantInboundSAMLConfigStatus `json:"status,omitempty"`
}

// IdentityPlatformTenantInboundSAMLConfigSpec defines the desired state of IdentityPlatformTenantInboundSAMLConfig.
type IdentityPlatformTenantInboundSAMLConfigSpec struct {
	// Human friendly display name.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// If this config allows users to sign in with the provider.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	// SAML IdP configuration when the project acts as the relying party.
	IdpConfig map[string]interface{} `json:"idpConfig,omitempty" yaml:"idpConfig,omitempty"`
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// SAML SP (Service Provider) configuration when the project acts as the relying party to receive
	// and accept an authentication assertion issued by a SAML identity provider.
	SpConfig map[string]interface{} `json:"spConfig,omitempty" yaml:"spConfig,omitempty"`
	// Immutable. The name of the tenant where this inbound SAML config resource exists.
	Tenant string `json:"tenant,omitempty" yaml:"tenant,omitempty"`
}

// IdentityPlatformTenantInboundSAMLConfigStatus defines the observed state of IdentityPlatformTenantInboundSAMLConfig.
type IdentityPlatformTenantInboundSAMLConfigStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
