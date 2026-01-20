package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ComputeTargetGRPCProxy represents a compute.cnrm.cloud.google.com ComputeTargetGRPCProxy resource.
type ComputeTargetGRPCProxy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeTargetGRPCProxySpec   `json:"spec,omitempty"`
	Status ComputeTargetGRPCProxyStatus `json:"status,omitempty"`
}

// ComputeTargetGRPCProxySpec defines the desired state of ComputeTargetGRPCProxy.
type ComputeTargetGRPCProxySpec struct {
	// An optional description of this resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// The UrlMap resource that defines the mapping from URL to the BackendService.
	// The protocol field in the BackendService must be set to GRPC.
	URLMapRef map[string]interface{} `json:"urlMapRef,omitempty" yaml:"urlMapRef,omitempty"`
	// Immutable. If true, indicates that the BackendServices referenced by
	// the urlMap may be accessed by gRPC applications without using
	// a sidecar proxy. This will enable configuration checks on urlMap
	// and its referenced BackendServices to not allow unsupported features.
	// A gRPC application must use "xds:///" scheme in the target URI
	// of the service it is connecting to. If false, indicates that the
	// BackendServices referenced by the urlMap will be accessed by gRPC
	// applications via a sidecar proxy. In this case, a gRPC application
	// must not use "xds:///" scheme in the target URI of the service
	// it is connecting to.
	ValidateForProxyless bool `json:"validateForProxyless,omitempty" yaml:"validateForProxyless,omitempty"`
}

// ComputeTargetGRPCProxyStatus defines the observed state of ComputeTargetGRPCProxy.
type ComputeTargetGRPCProxyStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
