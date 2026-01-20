package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ComputeTargetSSLProxy represents a compute.cnrm.cloud.google.com ComputeTargetSSLProxy resource.
type ComputeTargetSSLProxy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeTargetSSLProxySpec   `json:"spec,omitempty"`
	Status ComputeTargetSSLProxyStatus `json:"status,omitempty"`
}

// ComputeTargetSSLProxySpec defines the desired state of ComputeTargetSSLProxy.
type ComputeTargetSSLProxySpec struct {
	// A reference to the ComputeBackendService resource.
	BackendServiceRef map[string]interface{} `json:"backendServiceRef,omitempty" yaml:"backendServiceRef,omitempty"`
	// A reference to the CertificateMap resource uri that identifies a
	// certificate map associated with the given target proxy. This
	// field can only be set for global target proxies. Accepted format is
	// '//certificatemanager.googleapis.com/projects/{project}/locations/{location}/certificateMaps/{resourceName}'.
	CertificateMapRef map[string]interface{} `json:"certificateMapRef,omitempty" yaml:"certificateMapRef,omitempty"`
	// Immutable. An optional description of this resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Specifies the type of proxy header to append before sending data to
	// the backend. Default value: "NONE" Possible values: ["NONE", "PROXY_V1"].
	ProxyHeader string `json:"proxyHeader,omitempty" yaml:"proxyHeader,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	SslCertificates []map[string]interface{} `json:"sslCertificates,omitempty" yaml:"sslCertificates,omitempty"`
	// A reference to the ComputeSSLPolicy resource that will be
	// associated with the TargetSslProxy resource. If not set, the
	// ComputeTargetSSLProxy resource will not have any SSL policy
	// configured.
	SslPolicyRef map[string]interface{} `json:"sslPolicyRef,omitempty" yaml:"sslPolicyRef,omitempty"`
}

// ComputeTargetSSLProxyStatus defines the observed state of ComputeTargetSSLProxy.
type ComputeTargetSSLProxyStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
