package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ComputeBackendBucket represents a compute.cnrm.cloud.google.com ComputeBackendBucket resource.
type ComputeBackendBucket struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeBackendBucketSpec   `json:"spec,omitempty"`
	Status ComputeBackendBucketStatus `json:"status,omitempty"`
}

// ComputeBackendBucketSpec defines the desired state of ComputeBackendBucket.
type ComputeBackendBucketSpec struct {
	// Reference to the bucket.
	BucketRef map[string]interface{} `json:"bucketRef,omitempty" yaml:"bucketRef,omitempty"`
	// Cloud CDN configuration for this Backend Bucket.
	CdnPolicy map[string]interface{} `json:"cdnPolicy,omitempty" yaml:"cdnPolicy,omitempty"`
	// Compress text responses using Brotli or gzip compression, based on the client's Accept-Encoding header. Possible values: ["AUTOMATIC", "DISABLED"].
	CompressionMode string `json:"compressionMode,omitempty" yaml:"compressionMode,omitempty"`
	// Headers that the HTTP/S load balancer should add to proxied responses.
	CustomResponseHeaders []string `json:"customResponseHeaders,omitempty" yaml:"customResponseHeaders,omitempty"`
	// An optional textual description of the resource; provided by the
	// client when the resource is created.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// The security policy associated with this backend bucket.
	EdgeSecurityPolicy string `json:"edgeSecurityPolicy,omitempty" yaml:"edgeSecurityPolicy,omitempty"`
	// If true, enable Cloud CDN for this BackendBucket.
	EnableCdn bool `json:"enableCdn,omitempty" yaml:"enableCdn,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// ComputeBackendBucketStatus defines the observed state of ComputeBackendBucket.
type ComputeBackendBucketStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
