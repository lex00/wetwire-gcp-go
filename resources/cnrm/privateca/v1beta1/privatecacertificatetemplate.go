package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PrivateCACertificateTemplate represents a privateca.cnrm.cloud.google.com PrivateCACertificateTemplate resource.
type PrivateCACertificateTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PrivateCACertificateTemplateSpec   `json:"spec,omitempty"`
	Status PrivateCACertificateTemplateStatus `json:"status,omitempty"`
}

// PrivateCACertificateTemplateSpec defines the desired state of PrivateCACertificateTemplate.
type PrivateCACertificateTemplateSpec struct {
	// Optional. A human-readable description of scenarios this template is intended for.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Optional. Describes constraints on identities that may be appear in Certificates issued using this template. If this is omitted, then this template will not add restrictions on a certificate's identity.
	IdentityConstraints map[string]interface{} `json:"identityConstraints,omitempty" yaml:"identityConstraints,omitempty"`
	// Immutable. The location for the resource
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Optional. Describes the set of X.509 extensions that may appear in a Certificate issued using this CertificateTemplate. If a certificate request sets extensions that don't appear in the passthrough_extensions, those extensions will be dropped. If the issuing CaPool's IssuancePolicy defines baseline_values that don't appear here, the certificate issuance request will fail. If this is omitted, then this template will not add restrictions on a certificate's X.509 extensions. These constraints do not apply to X.509 extensions set in this CertificateTemplate's predefined_values.
	PassthroughExtensions map[string]interface{} `json:"passthroughExtensions,omitempty" yaml:"passthroughExtensions,omitempty"`
	// Optional. A set of X.509 values that will be applied to all issued certificates that use this template. If the certificate request includes conflicting values for the same properties, they will be overwritten by the values defined here. If the issuing CaPool's IssuancePolicy defines conflicting baseline_values for the same properties, the certificate issuance request will fail.
	PredefinedValues map[string]interface{} `json:"predefinedValues,omitempty" yaml:"predefinedValues,omitempty"`
	// Immutable. The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// PrivateCACertificateTemplateStatus defines the observed state of PrivateCACertificateTemplate.
type PrivateCACertificateTemplateStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
