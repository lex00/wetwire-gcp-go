package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ContainerAttachedCluster is the Schema for the ContainerAttachedCluster API
type ContainerAttachedCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ContainerAttachedClusterSpec   `json:"spec,omitempty"`
	Status ContainerAttachedClusterStatus `json:"status,omitempty"`
}

// ContainerAttachedClusterSpec defines the desired state of ContainerAttachedCluster.
type ContainerAttachedClusterSpec struct {
	// Optional. Annotations on the cluster.
	// This field has the same restrictions as Kubernetes annotations.
	// The total size of all keys and values combined is limited to 256k.
	// Key can have 2 segments: prefix (optional) and name (required),
	// separated by a slash (/).
	// Prefix must be a DNS subdomain.
	// Name must be 63 characters or less, begin and end with alphanumerics,
	// with dashes (-), underscores (_), dots (.), and alphanumerics between.
	Annotations map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	// Optional. Configuration related to the cluster RBAC settings.
	Authorization map[string]interface{} `json:"authorization,omitempty" yaml:"authorization,omitempty"`
	// Optional. Binary Authorization configuration for this cluster.
	BinaryAuthorization map[string]interface{} `json:"binaryAuthorization,omitempty" yaml:"binaryAuthorization,omitempty"`
	// Optional. Policy to determine what flags to send on delete.
	DeletionPolicy string `json:"deletionPolicy,omitempty" yaml:"deletionPolicy,omitempty"`
	// Optional. A human readable description of this Attached cluster. Cannot be longer than 255 UTF-8 encoded bytes.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Immutable. The Kubernetes distribution of the underlying attached cluster.
	// Supported values: ["eks", "aks", "generic"].
	Distribution string `json:"distribution,omitempty" yaml:"distribution,omitempty"`
	// Required. Fleet configuration.
	Fleet map[string]interface{} `json:"fleet,omitempty" yaml:"fleet,omitempty"`
	// Immutable. The location for the resource.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Optional. Logging configuration for this cluster.
	LoggingConfig map[string]interface{} `json:"loggingConfig,omitempty" yaml:"loggingConfig,omitempty"`
	// Optional. Monitoring configuration for this cluster.
	MonitoringConfig map[string]interface{} `json:"monitoringConfig,omitempty" yaml:"monitoringConfig,omitempty"`
	// Required. OpenID Connect (OIDC) discovery information of the target cluster.
	// Kubernetes Service Account (KSA) tokens are JWT tokens signed by the cluster
	// API server. This field indicates how GCP services	validate KSA tokens in order
	// to allow system workloads (such as GKE Connect and telemetry agents) to
	// authenticate back to GCP.
	// Both clusters with public and private issuer URLs are supported.
	// Clusters with public issuers only need to specify the 'issuerUrl' field
	// while clusters with private issuers need to provide both 'issuerUrl' and 'jwks'.
	OidcConfig map[string]interface{} `json:"oidcConfig,omitempty" yaml:"oidcConfig,omitempty"`
	// Required. The platform version for the cluster (e.g. `1.30.0-gke.1`).
	PlatformVersion string `json:"platformVersion,omitempty" yaml:"platformVersion,omitempty"`
	// The ID of the project in which the resource belongs.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Optional. The ContainerAttachedCluster name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// ContainerAttachedClusterStatus defines the observed state of ContainerAttachedCluster.
type ContainerAttachedClusterStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
