package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// AppEngineStandardAppVersion represents a appengine.cnrm.cloud.google.com AppEngineStandardAppVersion resource.
type AppEngineStandardAppVersion struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AppEngineStandardAppVersionSpec   `json:"spec,omitempty"`
	Status AppEngineStandardAppVersionStatus `json:"status,omitempty"`
}

// AppEngineStandardAppVersionSpec defines the desired state of AppEngineStandardAppVersion.
type AppEngineStandardAppVersionSpec struct {
	// Allows App Engine second generation runtimes to access the legacy bundled services.
	AppEngineApis bool `json:"appEngineApis,omitempty" yaml:"appEngineApis,omitempty"`
	// Automatic scaling is based on request rate, response latencies, and other application metrics.
	AutomaticScaling map[string]interface{} `json:"automaticScaling,omitempty" yaml:"automaticScaling,omitempty"`
	// Basic scaling creates instances when your application receives requests. Each instance will be shut down when the application becomes idle. Basic scaling is ideal for work that is intermittent or driven by user activity.
	BasicScaling map[string]interface{} `json:"basicScaling,omitempty" yaml:"basicScaling,omitempty"`
	// If set to 'true', the service will be deleted if it is the last version.
	DeleteServiceOnDestroy bool `json:"deleteServiceOnDestroy,omitempty" yaml:"deleteServiceOnDestroy,omitempty"`
	// Code and application artifacts that make up this version.
	Deployment map[string]interface{} `json:"deployment,omitempty" yaml:"deployment,omitempty"`
	// The entrypoint for the application.
	Entrypoint map[string]interface{} `json:"entrypoint,omitempty" yaml:"entrypoint,omitempty"`
	// Environment variables available to the application.
	EnvVariables map[string]string `json:"envVariables,omitempty" yaml:"envVariables,omitempty"`
	// An ordered list of URL-matching patterns that should be applied to incoming requests.
	// The first matching URL handles the request and other request handlers are not attempted.
	Handlers []map[string]interface{} `json:"handlers,omitempty" yaml:"handlers,omitempty"`
	// A list of the types of messages that this application is able to receive. Possible values: ["INBOUND_SERVICE_MAIL", "INBOUND_SERVICE_MAIL_BOUNCE", "INBOUND_SERVICE_XMPP_ERROR", "INBOUND_SERVICE_XMPP_MESSAGE", "INBOUND_SERVICE_XMPP_SUBSCRIBE", "INBOUND_SERVICE_XMPP_PRESENCE", "INBOUND_SERVICE_CHANNEL_PRESENCE", "INBOUND_SERVICE_WARMUP"].
	InboundServices []string `json:"inboundServices,omitempty" yaml:"inboundServices,omitempty"`
	// Instance class that is used to run this version. Valid values are
	// AutomaticScaling: F1, F2, F4, F4_1G
	// BasicScaling or ManualScaling: B1, B2, B4, B4_1G, B8
	// Defaults to F1 for AutomaticScaling and B2 for ManualScaling and BasicScaling. If no scaling is specified, AutomaticScaling is chosen.
	InstanceClass string `json:"instanceClass,omitempty" yaml:"instanceClass,omitempty"`
	// Configuration for third-party Python runtime libraries that are required by the application.
	Libraries []map[string]interface{} `json:"libraries,omitempty" yaml:"libraries,omitempty"`
	// A service with manual scaling runs continuously, allowing you to perform complex initialization and rely on the state of its memory over time.
	ManualScaling map[string]interface{} `json:"manualScaling,omitempty" yaml:"manualScaling,omitempty"`
	// If set to 'true', the application version will not be deleted.
	NoopOnDestroy bool `json:"noopOnDestroy,omitempty" yaml:"noopOnDestroy,omitempty"`
	// Immutable.
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
	// Immutable. Optional. The versionId of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Desired runtime. Example python27.
	Runtime string `json:"runtime,omitempty" yaml:"runtime,omitempty"`
	// The version of the API in the given runtime environment.
	// Please see the app.yaml reference for valid values at 'https://cloud.google.com/appengine/docs/standard/<language>/config/appref'\
	// Substitute '<language>' with 'python', 'java', 'php', 'ruby', 'go' or 'nodejs'.
	RuntimeAPIVersion string `json:"runtimeApiVersion,omitempty" yaml:"runtimeApiVersion,omitempty"`
	// The identity that the deployed version will run as. Admin API will use the App Engine Appspot service account as default if this field is neither provided in app.yaml file nor through CLI flag.
	ServiceAccount string `json:"serviceAccount,omitempty" yaml:"serviceAccount,omitempty"`
	ServiceRef map[string]interface{} `json:"serviceRef,omitempty" yaml:"serviceRef,omitempty"`
	// Whether multiple requests can be dispatched to this version at once.
	Threadsafe bool `json:"threadsafe,omitempty" yaml:"threadsafe,omitempty"`
	// Enables VPC connectivity for standard apps.
	VpcAccessConnector map[string]interface{} `json:"vpcAccessConnector,omitempty" yaml:"vpcAccessConnector,omitempty"`
}

// AppEngineStandardAppVersionStatus defines the observed state of AppEngineStandardAppVersion.
type AppEngineStandardAppVersionStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
