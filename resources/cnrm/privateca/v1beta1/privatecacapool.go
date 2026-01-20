package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PrivateCACAPool represents a privateca.cnrm.cloud.google.com PrivateCACAPool resource.
type PrivateCACAPool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PrivateCACAPoolSpec   `json:"spec,omitempty"`
	Status PrivateCACAPoolStatus `json:"status,omitempty"`
}

// PrivateCACAPoolSpec defines the desired state of PrivateCACAPool.
type PrivateCACAPoolSpec struct {
	// Optional. The IssuancePolicy to control how Certificates will be issued from this CaPool.
	IssuancePolicy map[string]interface{} `json:"issuancePolicy,omitempty" yaml:"issuancePolicy,omitempty"`
	// Immutable. The location for the resource
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Immutable. The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Optional. The PublishingOptions to follow when issuing Certificates from any CertificateAuthority in this CaPool.
	PublishingOptions map[string]interface{} `json:"publishingOptions,omitempty" yaml:"publishingOptions,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Immutable. Required. Immutable. The Tier of this CaPool. Possible values: TIER_UNSPECIFIED, ENTERPRISE, DEVOPS
	Tier string `json:"tier,omitempty" yaml:"tier,omitempty"`
}

// PrivateCACAPoolStatus defines the observed state of PrivateCACAPool.
type PrivateCACAPoolStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
