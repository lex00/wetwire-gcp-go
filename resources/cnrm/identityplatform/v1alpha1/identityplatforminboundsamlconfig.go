package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// IdentityPlatformInboundSAMLConfig represents a identityplatform.cnrm.cloud.google.com IdentityPlatformInboundSAMLConfig resource.
type IdentityPlatformInboundSAMLConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IdentityPlatformInboundSAMLConfigSpec   `json:"spec,omitempty"`
	Status IdentityPlatformInboundSAMLConfigStatus `json:"status,omitempty"`
}

// IdentityPlatformInboundSAMLConfigSpec defines the desired state of IdentityPlatformInboundSAMLConfig.
type IdentityPlatformInboundSAMLConfigSpec struct {
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
}

// IdentityPlatformInboundSAMLConfigStatus defines the observed state of IdentityPlatformInboundSAMLConfig.
type IdentityPlatformInboundSAMLConfigStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
