package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// IdentityPlatformConfig represents a identityplatform.cnrm.cloud.google.com IdentityPlatformConfig resource.
type IdentityPlatformConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IdentityPlatformConfigSpec   `json:"spec,omitempty"`
	Status IdentityPlatformConfigStatus `json:"status,omitempty"`
}

// IdentityPlatformConfigSpec defines the desired state of IdentityPlatformConfig.
type IdentityPlatformConfigSpec struct {
	// List of domains authorized for OAuth redirects
	AuthorizedDomains []string `json:"authorizedDomains,omitempty" yaml:"authorizedDomains,omitempty"`
	// Configuration related to blocking functions.
	BlockingFunctions map[string]interface{} `json:"blockingFunctions,omitempty" yaml:"blockingFunctions,omitempty"`
	// Options related to how clients making requests on behalf of a project should be configured.
	Client map[string]interface{} `json:"client,omitempty" yaml:"client,omitempty"`
	// Configuration for this project's multi-factor authentication, including whether it is active and what factors can be used for the second factor
	Mfa map[string]interface{} `json:"mfa,omitempty" yaml:"mfa,omitempty"`
	// Configuration related to monitoring project activity.
	Monitoring map[string]interface{} `json:"monitoring,omitempty" yaml:"monitoring,omitempty"`
	// Configuration related to multi-tenant functionality.
	MultiTenant map[string]interface{} `json:"multiTenant,omitempty" yaml:"multiTenant,omitempty"`
	// Configuration related to sending notifications to users.
	Notification map[string]interface{} `json:"notification,omitempty" yaml:"notification,omitempty"`
	// Immutable. The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Configuration related to quotas.
	Quota map[string]interface{} `json:"quota,omitempty" yaml:"quota,omitempty"`
	// Configuration related to local sign in methods.
	SignIn map[string]interface{} `json:"signIn,omitempty" yaml:"signIn,omitempty"`
}

// IdentityPlatformConfigStatus defines the observed state of IdentityPlatformConfig.
type IdentityPlatformConfigStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
