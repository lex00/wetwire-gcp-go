package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// IAPIdentityAwareProxyClient represents a iap.cnrm.cloud.google.com IAPIdentityAwareProxyClient resource.
type IAPIdentityAwareProxyClient struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IAPIdentityAwareProxyClientSpec   `json:"spec,omitempty"`
	Status IAPIdentityAwareProxyClientStatus `json:"status,omitempty"`
}

// IAPIdentityAwareProxyClientSpec defines the desired state of IAPIdentityAwareProxyClient.
type IAPIdentityAwareProxyClientSpec struct {
	// Immutable.
	BrandRef map[string]interface{} `json:"brandRef,omitempty" yaml:"brandRef,omitempty"`
	// Immutable. Human-friendly name given to the OAuth client.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// IAPIdentityAwareProxyClientStatus defines the observed state of IAPIdentityAwareProxyClient.
type IAPIdentityAwareProxyClientStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
