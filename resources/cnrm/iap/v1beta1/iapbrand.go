package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// IAPBrand represents a iap.cnrm.cloud.google.com IAPBrand resource.
type IAPBrand struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IAPBrandSpec   `json:"spec,omitempty"`
	Status IAPBrandStatus `json:"status,omitempty"`
}

// IAPBrandSpec defines the desired state of IAPBrand.
type IAPBrandSpec struct {
	// Immutable. Application name displayed on OAuth consent screen.
	ApplicationTitle string `json:"applicationTitle,omitempty" yaml:"applicationTitle,omitempty"`
	// Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Immutable. Support email displayed on the OAuth consent screen.
	SupportEmail string `json:"supportEmail,omitempty" yaml:"supportEmail,omitempty"`
}

// IAPBrandStatus defines the observed state of IAPBrand.
type IAPBrandStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
