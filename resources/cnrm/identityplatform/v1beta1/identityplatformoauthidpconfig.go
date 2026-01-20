package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// IdentityPlatformOAuthIDPConfig represents a identityplatform.cnrm.cloud.google.com IdentityPlatformOAuthIDPConfig resource.
type IdentityPlatformOAuthIDPConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IdentityPlatformOAuthIDPConfigSpec   `json:"spec,omitempty"`
	Status IdentityPlatformOAuthIDPConfigStatus `json:"status,omitempty"`
}

// IdentityPlatformOAuthIDPConfigSpec defines the desired state of IdentityPlatformOAuthIDPConfig.
type IdentityPlatformOAuthIDPConfigSpec struct {
	// The client id of an OAuth client.
	ClientID string `json:"clientId,omitempty" yaml:"clientId,omitempty"`
	// The client secret of the OAuth client, to enable OIDC code flow.
	ClientSecret map[string]interface{} `json:"clientSecret,omitempty" yaml:"clientSecret,omitempty"`
	// The config's display name set by developers.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// True if allows the user to sign in with the provider.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	// For OIDC Idps, the issuer identifier.
	Issuer string `json:"issuer,omitempty" yaml:"issuer,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// The multiple response type to request for in the OAuth authorization flow. This can possibly be a combination of set bits (e.g.: {id\_token, token}).
	ResponseType map[string]interface{} `json:"responseType,omitempty" yaml:"responseType,omitempty"`
}

// IdentityPlatformOAuthIDPConfigStatus defines the observed state of IdentityPlatformOAuthIDPConfig.
type IdentityPlatformOAuthIDPConfigStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
