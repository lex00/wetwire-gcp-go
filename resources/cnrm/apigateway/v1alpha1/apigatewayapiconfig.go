package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// APIGatewayAPIConfig represents a apigateway.cnrm.cloud.google.com APIGatewayAPIConfig resource.
type APIGatewayAPIConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   APIGatewayAPIConfigSpec   `json:"spec,omitempty"`
	Status APIGatewayAPIConfigStatus `json:"status,omitempty"`
}

// APIGatewayAPIConfigSpec defines the desired state of APIGatewayAPIConfig.
type APIGatewayAPIConfigSpec struct {
	// Immutable. The API to attach the config to.
	API string `json:"api,omitempty" yaml:"api,omitempty"`
	// Immutable. Creates a unique name beginning with the specified prefix. If this and api_config_id are unspecified, a random value is chosen for the name.
	APIConfigIDPrefix string `json:"apiConfigIdPrefix,omitempty" yaml:"apiConfigIdPrefix,omitempty"`
	// A user-visible name for the API.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// Immutable. Immutable. Gateway specific configuration.
	// If not specified, backend authentication will be set to use OIDC authentication using the default compute service account.
	GatewayConfig map[string]interface{} `json:"gatewayConfig,omitempty" yaml:"gatewayConfig,omitempty"`
	// gRPC service definition files. If specified, openapiDocuments must not be included.
	GrpcServices []map[string]interface{} `json:"grpcServices,omitempty" yaml:"grpcServices,omitempty"`
	// Optional. Service Configuration files. At least one must be included when using gRPC service definitions. See https://cloud.google.com/endpoints/docs/grpc/grpc-service-config#service_configuration_overview for the expected file contents.
	// If multiple files are specified, the files are merged with the following rules: * All singular scalar fields are merged using "last one wins" semantics in the order of the files uploaded. * Repeated fields are concatenated. * Singular embedded messages are merged using these rules for nested fields.
	ManagedServiceConfigs []map[string]interface{} `json:"managedServiceConfigs,omitempty" yaml:"managedServiceConfigs,omitempty"`
	// OpenAPI specification documents. If specified, grpcServices and managedServiceConfigs must not be included.
	OpenapiDocuments []map[string]interface{} `json:"openapiDocuments,omitempty" yaml:"openapiDocuments,omitempty"`
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The apiConfigId of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// APIGatewayAPIConfigStatus defines the observed state of APIGatewayAPIConfig.
type APIGatewayAPIConfigStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
