package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// VMwareEngineExternalAddress is the Schema for the VMwareEngineExternalAddress API
type VMwareEngineExternalAddress struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VMwareEngineExternalAddressSpec   `json:"spec,omitempty"`
	Status VMwareEngineExternalAddressStatus `json:"status,omitempty"`
}

// VMwareEngineExternalAddressSpec defines the desired state of VMwareEngineExternalAddress.
type VMwareEngineExternalAddressSpec struct {
	// User-provided description for this resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// The internal IP address of a workload VM.
	InternalIP string `json:"internalIP,omitempty" yaml:"internalIP,omitempty"`
	// Required. The resource name of the private cloud to create a new external IP address in.
	PrivateCloudRef map[string]interface{} `json:"privateCloudRef,omitempty" yaml:"privateCloudRef,omitempty"`
	// The VMwareEngineExternalAddress name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// VMwareEngineExternalAddressStatus defines the observed state of VMwareEngineExternalAddress.
type VMwareEngineExternalAddressStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
