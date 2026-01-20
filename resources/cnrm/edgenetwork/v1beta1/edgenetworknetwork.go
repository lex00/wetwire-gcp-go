package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EdgeNetworkNetwork represents a edgenetwork.cnrm.cloud.google.com EdgeNetworkNetwork resource.
type EdgeNetworkNetwork struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EdgeNetworkNetworkSpec   `json:"spec,omitempty"`
	Status EdgeNetworkNetworkStatus `json:"status,omitempty"`
}

// EdgeNetworkNetworkSpec defines the desired state of EdgeNetworkNetwork.
type EdgeNetworkNetworkSpec struct {
	// Immutable. A free-text description of the resource. Max length 1024 characters.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Immutable. The Google Cloud region to which the target Distributed Cloud Edge zone belongs.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Immutable. IP (L3) MTU value of the network. Default value is '1500'. Possible values are: '1500', '9000'.
	Mtu int32 `json:"mtu,omitempty" yaml:"mtu,omitempty"`
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The networkId of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Immutable. The name of the target Distributed Cloud Edge zone.
	Zone string `json:"zone,omitempty" yaml:"zone,omitempty"`
}

// EdgeNetworkNetworkStatus defines the observed state of EdgeNetworkNetwork.
type EdgeNetworkNetworkStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
