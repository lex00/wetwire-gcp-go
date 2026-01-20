package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ServiceDirectoryEndpoint represents a servicedirectory.cnrm.cloud.google.com ServiceDirectoryEndpoint resource.
type ServiceDirectoryEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ServiceDirectoryEndpointSpec   `json:"spec,omitempty"`
	Status ServiceDirectoryEndpointStatus `json:"status,omitempty"`
}

// ServiceDirectoryEndpointSpec defines the desired state of ServiceDirectoryEndpoint.
type ServiceDirectoryEndpointSpec struct {
	AddressRef map[string]interface{} `json:"addressRef,omitempty" yaml:"addressRef,omitempty"`
	// Only the `external` field is supported to configure the reference.
	// Immutable. The Google Compute Engine network (VPC) of the endpoint in the format
	// projects/<project number>/locations/global/networks/*.
	// The project must be specified by project number (project id is rejected). Incorrectly formatted networks are
	// rejected, but no other validation is performed on this field (ex. network or project existence,
	// reachability, or permissions).
	NetworkRef map[string]interface{} `json:"networkRef,omitempty" yaml:"networkRef,omitempty"`
	// Port that the endpoint is running on, must be in the
	// range of [0, 65535]. If unspecified, the default is 0.
	Port int32 `json:"port,omitempty" yaml:"port,omitempty"`
	// Immutable. Optional. The endpointId of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// The ServiceDirectoryService that this endpoint belongs to.
	ServiceRef map[string]interface{} `json:"serviceRef,omitempty" yaml:"serviceRef,omitempty"`
}

// ServiceDirectoryEndpointStatus defines the observed state of ServiceDirectoryEndpoint.
type ServiceDirectoryEndpointStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
