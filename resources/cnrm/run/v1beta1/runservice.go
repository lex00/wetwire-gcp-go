package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// RunService represents a run.cnrm.cloud.google.com RunService resource.
type RunService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RunServiceSpec   `json:"spec,omitempty"`
	Status RunServiceStatus `json:"status,omitempty"`
}

// RunServiceSpec defines the desired state of RunService.
type RunServiceSpec struct {
	// Unstructured key value map that may be set by external tools to store and arbitrary metadata. They are not queryable and should be preserved when modifying objects.
	// Cloud Run API v2 does not support annotations with 'run.googleapis.com', 'cloud.googleapis.com', 'serving.knative.dev', or 'autoscaling.knative.dev' namespaces, and they will be rejected in new resources.
	// All system annotations in v1 now have a corresponding field in v2 Service.
	// This field follows Kubernetes annotations' namespacing, limits, and rules.
	Annotations map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	// Settings for the Binary Authorization feature.
	BinaryAuthorization map[string]interface{} `json:"binaryAuthorization,omitempty" yaml:"binaryAuthorization,omitempty"`
	// Arbitrary identifier for the API client.
	Client string `json:"client,omitempty" yaml:"client,omitempty"`
	// Arbitrary version identifier for the API client.
	ClientVersion string `json:"clientVersion,omitempty" yaml:"clientVersion,omitempty"`
	// One or more custom audiences that you want this service to support. Specify each custom audience as the full URL in a string. The custom audiences are encoded in the token and used to authenticate requests.
	// For more information, see https://cloud.google.com/run/docs/configuring/custom-audiences.
	CustomAudiences []string `json:"customAudiences,omitempty" yaml:"customAudiences,omitempty"`
	// User-provided description of the Service. This field currently has a 512-character limit.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Provides the ingress settings for this Service. On output, returns the currently observed ingress settings, or INGRESS_TRAFFIC_UNSPECIFIED if no revision is active. Possible values: ["INGRESS_TRAFFIC_ALL", "INGRESS_TRAFFIC_INTERNAL_ONLY", "INGRESS_TRAFFIC_INTERNAL_LOAD_BALANCER"].
	Ingress string `json:"ingress,omitempty" yaml:"ingress,omitempty"`
	// The launch stage as defined by [Google Cloud Platform Launch Stages](https://cloud.google.com/products#product-launch-stages). Cloud Run supports ALPHA, BETA, and GA.
	// If no value is specified, GA is assumed. Set the launch stage to a preview stage on input to allow use of preview features in that stage. On read (or output), describes whether the resource uses preview features.
	// For example, if ALPHA is provided as input, but only BETA and GA-level features are used, this field will be BETA on output. Possible values: ["UNIMPLEMENTED", "PRELAUNCH", "EARLY_ACCESS", "ALPHA", "BETA", "GA", "DEPRECATED"].
	LaunchStage string `json:"launchStage,omitempty" yaml:"launchStage,omitempty"`
	// Immutable. The location of the cloud run service.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// The template used to create revisions for this Service.
	Template map[string]interface{} `json:"template,omitempty" yaml:"template,omitempty"`
	// Specifies how to distribute traffic over a collection of Revisions belonging to the Service. If traffic is empty or not provided, defaults to 100% traffic to the latest Ready Revision.
	Traffic []map[string]interface{} `json:"traffic,omitempty" yaml:"traffic,omitempty"`
}

// RunServiceStatus defines the observed state of RunService.
type RunServiceStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
