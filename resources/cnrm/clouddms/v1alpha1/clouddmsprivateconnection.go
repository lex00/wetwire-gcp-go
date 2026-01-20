package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CloudDMSPrivateConnection is the Schema for the CloudDMSPrivateConnection API
type CloudDMSPrivateConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CloudDMSPrivateConnectionSpec   `json:"spec,omitempty"`
	Status CloudDMSPrivateConnectionStatus `json:"status,omitempty"`
}

// CloudDMSPrivateConnectionSpec defines the desired state of CloudDMSPrivateConnection.
type CloudDMSPrivateConnectionSpec struct {
	// The private connection display name.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// The resource labels for private connections to use to annotate any related
	// underlying resources such as Compute Engine VMs. An object containing a
	// list of "key": "value" pairs.
	// Example: `{ "name": "wrench", "mass": "1.3kg", "count": "3" }`.
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	// Required. The location of the application.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Required. The host project of the application.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// The CloudDMSPrivateConnection name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// VPC peering configuration.
	VpcPeeringConfig map[string]interface{} `json:"vpcPeeringConfig,omitempty" yaml:"vpcPeeringConfig,omitempty"`
}

// CloudDMSPrivateConnectionStatus defines the observed state of CloudDMSPrivateConnection.
type CloudDMSPrivateConnectionStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
