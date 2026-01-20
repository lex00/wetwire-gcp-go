package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ComputeNetworkEdgeSecurityService is the Schema for the ComputeNetworkEdgeSecurityService API
type ComputeNetworkEdgeSecurityService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeNetworkEdgeSecurityServiceSpec   `json:"spec,omitempty"`
	Status ComputeNetworkEdgeSecurityServiceStatus `json:"status,omitempty"`
}

// ComputeNetworkEdgeSecurityServiceSpec defines the desired state of ComputeNetworkEdgeSecurityService.
type ComputeNetworkEdgeSecurityServiceSpec struct {
	// An optional description of this resource. Provide this property when you create the resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a NetworkEdgeSecurityService. An up-to-date fingerprint must be provided in order to update the NetworkEdgeSecurityService, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve a NetworkEdgeSecurityService.
	Fingerprint string `json:"fingerprint,omitempty" yaml:"fingerprint,omitempty"`
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// The ComputeNetworkEdgeSecurityService name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// The resource URL for the network edge security service associated with this network edge security service.
	SecurityPolicyRef map[string]interface{} `json:"securityPolicyRef,omitempty" yaml:"securityPolicyRef,omitempty"`
}

// ComputeNetworkEdgeSecurityServiceStatus defines the observed state of ComputeNetworkEdgeSecurityService.
type ComputeNetworkEdgeSecurityServiceStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
