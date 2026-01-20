package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ComputeHTTPSHealthCheck represents a compute.cnrm.cloud.google.com ComputeHTTPSHealthCheck resource.
type ComputeHTTPSHealthCheck struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeHTTPSHealthCheckSpec   `json:"spec,omitempty"`
	Status ComputeHTTPSHealthCheckStatus `json:"status,omitempty"`
}

// ComputeHTTPSHealthCheckSpec defines the desired state of ComputeHTTPSHealthCheck.
type ComputeHTTPSHealthCheckSpec struct {
	// How often (in seconds) to send a health check. The default value is 5
	// seconds.
	CheckIntervalSec int32 `json:"checkIntervalSec,omitempty" yaml:"checkIntervalSec,omitempty"`
	// An optional description of this resource. Provide this property when
	// you create the resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// A so-far unhealthy instance will be marked healthy after this many
	// consecutive successes. The default value is 2.
	HealthyThreshold int32 `json:"healthyThreshold,omitempty" yaml:"healthyThreshold,omitempty"`
	// The value of the host header in the HTTPS health check request. If
	// left empty (default value), the public IP on behalf of which this
	// health check is performed will be used.
	Host string `json:"host,omitempty" yaml:"host,omitempty"`
	// The TCP port number for the HTTPS health check request.
	// The default value is 443.
	Port int32 `json:"port,omitempty" yaml:"port,omitempty"`
	// The request path of the HTTPS health check request.
	// The default value is /.
	RequestPath string `json:"requestPath,omitempty" yaml:"requestPath,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// How long (in seconds) to wait before claiming failure.
	// The default value is 5 seconds.  It is invalid for timeoutSec to have
	// greater value than checkIntervalSec.
	TimeoutSec int32 `json:"timeoutSec,omitempty" yaml:"timeoutSec,omitempty"`
	// A so-far healthy instance will be marked unhealthy after this many
	// consecutive failures. The default value is 2.
	UnhealthyThreshold int32 `json:"unhealthyThreshold,omitempty" yaml:"unhealthyThreshold,omitempty"`
}

// ComputeHTTPSHealthCheckStatus defines the observed state of ComputeHTTPSHealthCheck.
type ComputeHTTPSHealthCheckStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
