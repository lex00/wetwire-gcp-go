package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// IdentityPlatformTenant represents a identityplatform.cnrm.cloud.google.com IdentityPlatformTenant resource.
type IdentityPlatformTenant struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IdentityPlatformTenantSpec   `json:"spec,omitempty"`
	Status IdentityPlatformTenantStatus `json:"status,omitempty"`
}

// IdentityPlatformTenantSpec defines the desired state of IdentityPlatformTenant.
type IdentityPlatformTenantSpec struct {
	// Whether to allow email/password user authentication.
	AllowPasswordSignup bool `json:"allowPasswordSignup,omitempty" yaml:"allowPasswordSignup,omitempty"`
	// Whether authentication is disabled for the tenant. If true, the users under the disabled tenant are not allowed to sign-in. Admins of the disabled tenant are not able to manage its users.
	DisableAuth bool `json:"disableAuth,omitempty" yaml:"disableAuth,omitempty"`
	// Display name of the tenant.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// Whether to enable anonymous user authentication.
	EnableAnonymousUser bool `json:"enableAnonymousUser,omitempty" yaml:"enableAnonymousUser,omitempty"`
	// Whether to enable email link user authentication.
	EnableEmailLinkSignin bool `json:"enableEmailLinkSignin,omitempty" yaml:"enableEmailLinkSignin,omitempty"`
	// The tenant-level configuration of MFA options.
	MfaConfig map[string]interface{} `json:"mfaConfig,omitempty" yaml:"mfaConfig,omitempty"`
	// Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// A map of <test phone number, fake code> pairs that can be used for MFA. The phone number should be in E.164 format (https://www.itu.int/rec/T-REC-E.164/) and a maximum of 10 pairs can be added (error will be thrown once exceeded).
	TestPhoneNumbers map[string]string `json:"testPhoneNumbers,omitempty" yaml:"testPhoneNumbers,omitempty"`
}

// IdentityPlatformTenantStatus defines the observed state of IdentityPlatformTenant.
type IdentityPlatformTenantStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
