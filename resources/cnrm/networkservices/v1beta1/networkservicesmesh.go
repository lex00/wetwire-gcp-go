package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NetworkServicesMesh represents a networkservices.cnrm.cloud.google.com NetworkServicesMesh resource.
type NetworkServicesMesh struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NetworkServicesMeshSpec   `json:"spec,omitempty"`
	Status NetworkServicesMeshStatus `json:"status,omitempty"`
}

// NetworkServicesMeshSpec defines the desired state of NetworkServicesMesh.
type NetworkServicesMeshSpec struct {
	// Optional. A free-text description of the resource. Max length 1024 characters.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Optional. If set to a valid TCP port (1-65535), instructs the SIDECAR proxy to listen on the specified port of localhost (127.0.0.1) address. The SIDECAR proxy will expect all traffic to be redirected to this port regardless of its actual ip:port destination. If unset, a port '15001' is used as the interception port. This field is only valid if the type of Mesh is SIDECAR.
	InterceptionPort int64 `json:"interceptionPort,omitempty" yaml:"interceptionPort,omitempty"`
	// Immutable. The location for the resource
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Immutable. The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// NetworkServicesMeshStatus defines the observed state of NetworkServicesMesh.
type NetworkServicesMeshStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
