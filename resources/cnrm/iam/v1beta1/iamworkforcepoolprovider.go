package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// IAMWorkforcePoolProvider represents a iam.cnrm.cloud.google.com IAMWorkforcePoolProvider resource.
type IAMWorkforcePoolProvider struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IAMWorkforcePoolProviderSpec   `json:"spec,omitempty"`
	Status IAMWorkforcePoolProviderStatus `json:"status,omitempty"`
}

// IAMWorkforcePoolProviderSpec defines the desired state of IAMWorkforcePoolProvider.
type IAMWorkforcePoolProviderSpec struct {
	// A [Common Expression Language](https://opensource.google/projects/cel) expression, in plain text, to restrict what otherwise valid authentication credentials issued by the provider should not be accepted. The expression must output a boolean representing whether to allow the federation. The following keywords may be referenced in the expressions: * `assertion`: JSON representing the authentication credential issued by the provider. * `google`: The Google attributes mapped from the assertion in the `attribute_mappings`. `google.profile_photo` and `google.display_name` are not supported. * `attribute`: The custom attributes mapped from the assertion in the `attribute_mappings`. The maximum length of the attribute condition expression is 4096 characters. If unspecified, all valid authentication credentials will be accepted. The following example shows how to only allow credentials with a mapped `google.groups` value of `admins`: ``` "'admins' in google.groups" ```
	AttributeCondition string `json:"attributeCondition,omitempty" yaml:"attributeCondition,omitempty"`
	// Required. Maps attributes from the authentication credentials issued by an external identity provider to Google Cloud attributes, such as `subject` and `segment`. Each key must be a string specifying the Google Cloud IAM attribute to map to. The following keys are supported: * `google.subject`: The principal IAM is authenticating. You can reference this value in IAM bindings. This is also the subject that appears in Cloud Logging logs. This is a required field and the mapped subject cannot exceed 127 bytes. * `google.groups`: Groups the authenticating user belongs to. You can grant groups access to resources using an IAM `principalSet` binding; access applies to all members of the group. * `google.display_name`: The name of the authenticated user. This is an optional field and the mapped display name cannot exceed 100 bytes. If not set, `google.subject` will be displayed instead. This attribute cannot be referenced in IAM bindings. * `google.profile_photo`: The URL that specifies the authenticated user's thumbnail photo. This is an optional field. When set, the image will be visible as the user's profile picture. If not set, a generic user icon will be displayed instead. This attribute cannot be referenced in IAM bindings. You can also provide custom attributes by specifying `attribute.{custom_attribute}`, where {custom_attribute} is the name of the custom attribute to be mapped. You can define a maximum of 50 custom attributes. The maximum length of a mapped attribute key is 100 characters, and the key may only contain the characters [a-z0-9_]. You can reference these attributes in IAM policies to define fine-grained access for a workforce pool to Google Cloud resources. For example:
	AttributeMapping map[string]string `json:"attributeMapping,omitempty" yaml:"attributeMapping,omitempty"`
	// A user-specified description of the provider. Cannot exceed 256 characters.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Whether the provider is disabled. You cannot use a disabled provider to exchange tokens. However, existing tokens still grant access.
	Disabled bool `json:"disabled,omitempty" yaml:"disabled,omitempty"`
	// A user-specified display name for the provider. Cannot exceed 32 characters.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// Immutable. The location for the resource
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// An OpenId Connect 1.0 identity provider configuration.
	Oidc map[string]interface{} `json:"oidc,omitempty" yaml:"oidc,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// A SAML identity provider configuration.
	Saml map[string]interface{} `json:"saml,omitempty" yaml:"saml,omitempty"`
	// Immutable.
	WorkforcePoolRef map[string]interface{} `json:"workforcePoolRef,omitempty" yaml:"workforcePoolRef,omitempty"`
}

// IAMWorkforcePoolProviderStatus defines the observed state of IAMWorkforcePoolProvider.
type IAMWorkforcePoolProviderStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
