package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// APIGatewayGateway represents a apigateway.cnrm.cloud.google.com APIGatewayGateway resource.
type APIGatewayGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   APIGatewayGatewaySpec   `json:"spec,omitempty"`
	Status APIGatewayGatewayStatus `json:"status,omitempty"`
}

// APIGatewayGatewaySpec defines the desired state of APIGatewayGateway.
type APIGatewayGatewaySpec struct {
	// Resource name of the API Config for this Gateway. Format: projects/{project}/locations/global/apis/{api}/configs/{apiConfig}.
	// When changing api configs please ensure the new config is a new resource and the lifecycle rule 'create_before_destroy' is set.
	APIConfig string `json:"apiConfig,omitempty" yaml:"apiConfig,omitempty"`
	// A user-visible name for the API.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. The region of the gateway for the API.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`
	// Immutable. Optional. The gatewayId of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// APIGatewayGatewayStatus defines the observed state of APIGatewayGateway.
type APIGatewayGatewayStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
