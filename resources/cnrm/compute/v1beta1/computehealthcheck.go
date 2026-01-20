package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ComputeHealthCheck represents a compute.cnrm.cloud.google.com ComputeHealthCheck resource.
type ComputeHealthCheck struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeHealthCheckSpec   `json:"spec,omitempty"`
	Status ComputeHealthCheckStatus `json:"status,omitempty"`
}

// ComputeHealthCheckSpec defines the desired state of ComputeHealthCheck.
type ComputeHealthCheckSpec struct {
	// How often (in seconds) to send a health check. The default value is 5
	// seconds.
	CheckIntervalSec int32 `json:"checkIntervalSec,omitempty" yaml:"checkIntervalSec,omitempty"`
	// An optional description of this resource. Provide this property when
	// you create the resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// A nested object resource.
	GrpcHealthCheck map[string]interface{} `json:"grpcHealthCheck,omitempty" yaml:"grpcHealthCheck,omitempty"`
	// A so-far unhealthy instance will be marked healthy after this many
	// consecutive successes. The default value is 2.
	HealthyThreshold int32 `json:"healthyThreshold,omitempty" yaml:"healthyThreshold,omitempty"`
	// A nested object resource.
	Http2HealthCheck map[string]interface{} `json:"http2HealthCheck,omitempty" yaml:"http2HealthCheck,omitempty"`
	// A nested object resource.
	HTTPHealthCheck map[string]interface{} `json:"httpHealthCheck,omitempty" yaml:"httpHealthCheck,omitempty"`
	// A nested object resource.
	HTTPSHealthCheck map[string]interface{} `json:"httpsHealthCheck,omitempty" yaml:"httpsHealthCheck,omitempty"`
	// Location represents the geographical location of the ComputeHealthCheck. Specify a region name or "global" for global resources. Reference: GCP definition of regions/zones (https://cloud.google.com/compute/docs/regions-zones/)
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Configure logging on this health check.
	LogConfig map[string]interface{} `json:"logConfig,omitempty" yaml:"logConfig,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// A nested object resource.
	SslHealthCheck map[string]interface{} `json:"sslHealthCheck,omitempty" yaml:"sslHealthCheck,omitempty"`
	// A nested object resource.
	TCPHealthCheck map[string]interface{} `json:"tcpHealthCheck,omitempty" yaml:"tcpHealthCheck,omitempty"`
	// How long (in seconds) to wait before claiming failure.
	// The default value is 5 seconds.  It is invalid for timeoutSec to have
	// greater value than checkIntervalSec.
	TimeoutSec int32 `json:"timeoutSec,omitempty" yaml:"timeoutSec,omitempty"`
	// A so-far healthy instance will be marked unhealthy after this many
	// consecutive failures. The default value is 2.
	UnhealthyThreshold int32 `json:"unhealthyThreshold,omitempty" yaml:"unhealthyThreshold,omitempty"`
}

// ComputeHealthCheckStatus defines the observed state of ComputeHealthCheck.
type ComputeHealthCheckStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
