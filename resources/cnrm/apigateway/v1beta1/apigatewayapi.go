package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// APIGatewayAPI is the Schema for the APIGatewayAPI API
type APIGatewayAPI struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   APIGatewayAPISpec   `json:"spec,omitempty"`
	Status APIGatewayAPIStatus `json:"status,omitempty"`
}

// APIGatewayAPISpec defines the desired state of APIGatewayAPI.
type APIGatewayAPISpec struct {
	// Optional. Display name.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// Optional. Resource labels to represent user-provided metadata. For more information, see the {{compute_name_short}} documentation: https://cloud.google.com/compute/docs/labeling-resources
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	// Optional. Immutable. The name of a Google Managed Service ( https://cloud.google.com/service-infrastructure/docs/glossary#managed). If not specified, a new Service will automatically be created in the same project as this API.
	ManagedService string `json:"managedService,omitempty" yaml:"managedService,omitempty"`
	// Optional. The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// The APIGatewayAPI name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// APIGatewayAPIStatus defines the observed state of APIGatewayAPI.
type APIGatewayAPIStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
