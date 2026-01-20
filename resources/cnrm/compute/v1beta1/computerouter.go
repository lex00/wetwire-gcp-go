package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ComputeRouter represents a compute.cnrm.cloud.google.com ComputeRouter resource.
type ComputeRouter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeRouterSpec   `json:"spec,omitempty"`
	Status ComputeRouterStatus `json:"status,omitempty"`
}

// ComputeRouterSpec defines the desired state of ComputeRouter.
type ComputeRouterSpec struct {
	// BGP information specific to this router.
	Bgp map[string]interface{} `json:"bgp,omitempty" yaml:"bgp,omitempty"`
	// An optional description of this resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Immutable. Indicates if a router is dedicated for use with encrypted VLAN
	// attachments (interconnectAttachments).
	EncryptedInterconnectRouter bool `json:"encryptedInterconnectRouter,omitempty" yaml:"encryptedInterconnectRouter,omitempty"`
	// A reference to the network to which this router belongs.
	NetworkRef map[string]interface{} `json:"networkRef,omitempty" yaml:"networkRef,omitempty"`
	// Immutable. Region where the router resides.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// ComputeRouterStatus defines the observed state of ComputeRouter.
type ComputeRouterStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
