package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CloudIDSEndpoint represents a cloudids.cnrm.cloud.google.com CloudIDSEndpoint resource.
type CloudIDSEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CloudIDSEndpointSpec   `json:"spec,omitempty"`
	Status CloudIDSEndpointStatus `json:"status,omitempty"`
}

// CloudIDSEndpointSpec defines the desired state of CloudIDSEndpoint.
type CloudIDSEndpointSpec struct {
	// Immutable. An optional description of the endpoint.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Immutable. The location for the endpoint.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Immutable. Name of the VPC network that is connected
	// to the IDS endpoint. This can either contain the VPC network name
	// itself (like "src-net") or the full URL to the network (like "projects/{project_id}/global/networks/src-net").
	NetworkRef map[string]interface{} `json:"networkRef,omitempty" yaml:"networkRef,omitempty"`
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Immutable. The minimum alert severity level that is reported by the endpoint. Possible values: ["INFORMATIONAL", "LOW", "MEDIUM", "HIGH", "CRITICAL"].
	Severity string `json:"severity,omitempty" yaml:"severity,omitempty"`
	// Configuration for threat IDs excluded from generating alerts. Limit: 99 IDs.
	ThreatExceptions []string `json:"threatExceptions,omitempty" yaml:"threatExceptions,omitempty"`
}

// CloudIDSEndpointStatus defines the observed state of CloudIDSEndpoint.
type CloudIDSEndpointStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
