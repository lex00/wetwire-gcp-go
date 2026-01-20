package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CloudFunctionsFunction represents a cloudfunctions.cnrm.cloud.google.com CloudFunctionsFunction resource.
type CloudFunctionsFunction struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CloudFunctionsFunctionSpec   `json:"spec,omitempty"`
	Status CloudFunctionsFunctionStatus `json:"status,omitempty"`
}

// CloudFunctionsFunctionSpec defines the desired state of CloudFunctionsFunction.
type CloudFunctionsFunctionSpec struct {
	// Memory (in MB), available to the function. Default value is 256MB. Allowed values are: 128MB, 256MB, 512MB, 1024MB, and 2048MB.
	AvailableMemoryMb int64 `json:"availableMemoryMb,omitempty" yaml:"availableMemoryMb,omitempty"`
	// User-provided description of a function.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Immutable. The name of the function (as defined in source code) that will be
	// executed. Defaults to the resource name suffix, if not specified. For
	// backward compatibility, if function with given name is not found, then the
	// system will try to use function named "function".
	// For Node.js this is name of a function exported by the module specified
	// in `source_location`.
	EntryPoint string `json:"entryPoint,omitempty" yaml:"entryPoint,omitempty"`
	// Environment variables that shall be available during function execution.
	EnvironmentVariables map[string]string `json:"environmentVariables,omitempty" yaml:"environmentVariables,omitempty"`
	// Immutable. A source that fires events in response to a condition in another service.
	EventTrigger map[string]interface{} `json:"eventTrigger,omitempty" yaml:"eventTrigger,omitempty"`
	// Immutable. An HTTPS endpoint type of source that can be triggered via URL.
	HTTPSTrigger map[string]interface{} `json:"httpsTrigger,omitempty" yaml:"httpsTrigger,omitempty"`
	// The ingress settings for the function, controlling what traffic can reach
	// it. Possible values: INGRESS_SETTINGS_UNSPECIFIED, ALLOW_ALL, ALLOW_INTERNAL_ONLY, ALLOW_INTERNAL_AND_GCLB
	IngressSettings string `json:"ingressSettings,omitempty" yaml:"ingressSettings,omitempty"`
	// The limit on the maximum number of function instances that may coexist at a
	// given time.
	MaxInstances int64 `json:"maxInstances,omitempty" yaml:"maxInstances,omitempty"`
	// Immutable. The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. The name of the Cloud Functions region of the function.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// The runtime in which to run the function. Required when deploying a new
	// function, optional when updating an existing function. For a complete
	// list of possible choices, see the
	// [`gcloud` command
	// reference](/sdk/gcloud/reference/functions/deploy#--runtime).
	Runtime string `json:"runtime,omitempty" yaml:"runtime,omitempty"`
	// Immutable.
	ServiceAccountRef map[string]interface{} `json:"serviceAccountRef,omitempty" yaml:"serviceAccountRef,omitempty"`
	// Immutable. The Google Cloud Storage URL, starting with gs://, pointing to the zip archive which contains the function.
	SourceArchiveURL string `json:"sourceArchiveUrl,omitempty" yaml:"sourceArchiveUrl,omitempty"`
	// Immutable. Represents parameters related to source repository where a function is hosted.
	SourceRepository map[string]interface{} `json:"sourceRepository,omitempty" yaml:"sourceRepository,omitempty"`
	// The function execution timeout. Execution is considered failed and
	// can be terminated if the function is not completed at the end of the
	// timeout period. Defaults to 60 seconds.
	Timeout string `json:"timeout,omitempty" yaml:"timeout,omitempty"`
	// The egress settings for the connector, controlling what traffic is diverted
	// through it. Possible values: VPC_CONNECTOR_EGRESS_SETTINGS_UNSPECIFIED, PRIVATE_RANGES_ONLY, ALL_TRAFFIC
	VpcConnectorEgressSettings string `json:"vpcConnectorEgressSettings,omitempty" yaml:"vpcConnectorEgressSettings,omitempty"`
	VpcConnectorRef map[string]interface{} `json:"vpcConnectorRef,omitempty" yaml:"vpcConnectorRef,omitempty"`
}

// CloudFunctionsFunctionStatus defines the observed state of CloudFunctionsFunction.
type CloudFunctionsFunctionStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
