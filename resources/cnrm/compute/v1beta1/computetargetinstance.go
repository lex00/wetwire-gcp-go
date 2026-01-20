package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ComputeTargetInstance represents a compute.cnrm.cloud.google.com ComputeTargetInstance resource.
type ComputeTargetInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeTargetInstanceSpec   `json:"spec,omitempty"`
	Status ComputeTargetInstanceStatus `json:"status,omitempty"`
}

// ComputeTargetInstanceSpec defines the desired state of ComputeTargetInstance.
type ComputeTargetInstanceSpec struct {
	// Immutable. An optional description of this resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// The ComputeInstance handling traffic for this target instance.
	InstanceRef map[string]interface{} `json:"instanceRef,omitempty" yaml:"instanceRef,omitempty"`
	// Immutable. NAT option controlling how IPs are NAT'ed to the instance.
	// Currently only NO_NAT (default value) is supported. Default value: "NO_NAT" Possible values: ["NO_NAT"].
	NatPolicy string `json:"natPolicy,omitempty" yaml:"natPolicy,omitempty"`
	// The network this target instance uses to forward
	// traffic. If not specified, the traffic will be forwarded to the network
	// that the default network interface belongs to.
	NetworkRef map[string]interface{} `json:"networkRef,omitempty" yaml:"networkRef,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// The resource URL for the security policy associated with this target instance.
	SecurityPolicyRef map[string]interface{} `json:"securityPolicyRef,omitempty" yaml:"securityPolicyRef,omitempty"`
	// Immutable. URL of the zone where the target instance resides.
	Zone string `json:"zone,omitempty" yaml:"zone,omitempty"`
}

// ComputeTargetInstanceStatus defines the observed state of ComputeTargetInstance.
type ComputeTargetInstanceStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
