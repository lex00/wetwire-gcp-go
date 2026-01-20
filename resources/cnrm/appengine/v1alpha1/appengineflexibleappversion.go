package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// AppEngineFlexibleAppVersion represents a appengine.cnrm.cloud.google.com AppEngineFlexibleAppVersion resource.
type AppEngineFlexibleAppVersion struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AppEngineFlexibleAppVersionSpec   `json:"spec,omitempty"`
	Status AppEngineFlexibleAppVersionStatus `json:"status,omitempty"`
}

// AppEngineFlexibleAppVersionSpec defines the desired state of AppEngineFlexibleAppVersion.
type AppEngineFlexibleAppVersionSpec struct {
	// Serving configuration for Google Cloud Endpoints.
	APIConfig map[string]interface{} `json:"apiConfig,omitempty" yaml:"apiConfig,omitempty"`
	// Automatic scaling is based on request rate, response latencies, and other application metrics.
	AutomaticScaling map[string]interface{} `json:"automaticScaling,omitempty" yaml:"automaticScaling,omitempty"`
	// Metadata settings that are supplied to this version to enable beta runtime features.
	BetaSettings map[string]string `json:"betaSettings,omitempty" yaml:"betaSettings,omitempty"`
	// Duration that static files should be cached by web proxies and browsers.
	// Only applicable if the corresponding StaticFilesHandler does not specify its own expiration time.
	DefaultExpiration string `json:"defaultExpiration,omitempty" yaml:"defaultExpiration,omitempty"`
	// If set to 'true', the service will be deleted if it is the last version.
	DeleteServiceOnDestroy bool `json:"deleteServiceOnDestroy,omitempty" yaml:"deleteServiceOnDestroy,omitempty"`
	// Code and application artifacts that make up this version.
	Deployment map[string]interface{} `json:"deployment,omitempty" yaml:"deployment,omitempty"`
	// Code and application artifacts that make up this version.
	EndpointsAPIService map[string]interface{} `json:"endpointsApiService,omitempty" yaml:"endpointsApiService,omitempty"`
	// The entrypoint for the application.
	Entrypoint map[string]interface{} `json:"entrypoint,omitempty" yaml:"entrypoint,omitempty"`
	EnvVariables map[string]string `json:"envVariables,omitempty" yaml:"envVariables,omitempty"`
	// An ordered list of URL-matching patterns that should be applied to incoming requests.
	// The first matching URL handles the request and other request handlers are not attempted.
	Handlers []map[string]interface{} `json:"handlers,omitempty" yaml:"handlers,omitempty"`
	// A list of the types of messages that this application is able to receive. Possible values: ["INBOUND_SERVICE_MAIL", "INBOUND_SERVICE_MAIL_BOUNCE", "INBOUND_SERVICE_XMPP_ERROR", "INBOUND_SERVICE_XMPP_MESSAGE", "INBOUND_SERVICE_XMPP_SUBSCRIBE", "INBOUND_SERVICE_XMPP_PRESENCE", "INBOUND_SERVICE_CHANNEL_PRESENCE", "INBOUND_SERVICE_WARMUP"].
	InboundServices []string `json:"inboundServices,omitempty" yaml:"inboundServices,omitempty"`
	// Instance class that is used to run this version. Valid values are
	// AutomaticScaling: F1, F2, F4, F4_1G
	// ManualScaling: B1, B2, B4, B8, B4_1G
	// Defaults to F1 for AutomaticScaling and B1 for ManualScaling.
	InstanceClass string `json:"instanceClass,omitempty" yaml:"instanceClass,omitempty"`
	// Health checking configuration for VM instances. Unhealthy instances are killed and replaced with new instances.
	LivenessCheck map[string]interface{} `json:"livenessCheck,omitempty" yaml:"livenessCheck,omitempty"`
	// A service with manual scaling runs continuously, allowing you to perform complex initialization and rely on the state of its memory over time.
	ManualScaling map[string]interface{} `json:"manualScaling,omitempty" yaml:"manualScaling,omitempty"`
	// Extra network settings.
	Network map[string]interface{} `json:"network,omitempty" yaml:"network,omitempty"`
	// Files that match this pattern will not be built into this version. Only applicable for Go runtimes.
	NobuildFilesRegex string `json:"nobuildFilesRegex,omitempty" yaml:"nobuildFilesRegex,omitempty"`
	// If set to 'true', the application version will not be deleted.
	NoopOnDestroy bool `json:"noopOnDestroy,omitempty" yaml:"noopOnDestroy,omitempty"`
	// Immutable.
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
	// Configures readiness health checking for instances. Unhealthy instances are not put into the backend traffic rotation.
	ReadinessCheck map[string]interface{} `json:"readinessCheck,omitempty" yaml:"readinessCheck,omitempty"`
	// Immutable. Optional. The versionId of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Machine resources for a version.
	Resources map[string]interface{} `json:"resources,omitempty" yaml:"resources,omitempty"`
	// Desired runtime. Example python27.
	Runtime string `json:"runtime,omitempty" yaml:"runtime,omitempty"`
	// The version of the API in the given runtime environment.
	// Please see the app.yaml reference for valid values at 'https://cloud.google.com/appengine/docs/standard/<language>/config/appref'\
	// Substitute '<language>' with 'python', 'java', 'php', 'ruby', 'go' or 'nodejs'.
	RuntimeAPIVersion string `json:"runtimeApiVersion,omitempty" yaml:"runtimeApiVersion,omitempty"`
	// The channel of the runtime to use. Only available for some runtimes.
	RuntimeChannel string `json:"runtimeChannel,omitempty" yaml:"runtimeChannel,omitempty"`
	// The path or name of the app's main executable.
	RuntimeMainExecutablePath string `json:"runtimeMainExecutablePath,omitempty" yaml:"runtimeMainExecutablePath,omitempty"`
	// The identity that the deployed version will run as. Admin API will use the App Engine Appspot service account as
	// default if this field is neither provided in app.yaml file nor through CLI flag.
	ServiceAccount string `json:"serviceAccount,omitempty" yaml:"serviceAccount,omitempty"`
	ServiceRef map[string]interface{} `json:"serviceRef,omitempty" yaml:"serviceRef,omitempty"`
	// Current serving status of this version. Only the versions with a SERVING status create instances and can be billed. Default value: "SERVING" Possible values: ["SERVING", "STOPPED"].
	ServingStatus string `json:"servingStatus,omitempty" yaml:"servingStatus,omitempty"`
	// Enables VPC connectivity for standard apps.
	VpcAccessConnector map[string]interface{} `json:"vpcAccessConnector,omitempty" yaml:"vpcAccessConnector,omitempty"`
}

// AppEngineFlexibleAppVersionStatus defines the observed state of AppEngineFlexibleAppVersion.
type AppEngineFlexibleAppVersionStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
