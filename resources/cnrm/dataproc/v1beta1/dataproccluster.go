package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DataprocCluster represents a dataproc.cnrm.cloud.google.com DataprocCluster resource.
type DataprocCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DataprocClusterSpec   `json:"spec,omitempty"`
	Status DataprocClusterStatus `json:"status,omitempty"`
}

// DataprocClusterSpec defines the desired state of DataprocCluster.
type DataprocClusterSpec struct {
	// Immutable. The cluster config. Note that Dataproc may set default values, and values may change when clusters are updated.
	Config map[string]interface{} `json:"config,omitempty" yaml:"config,omitempty"`
	// Immutable. The location for the resource, usually a GCP region.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Immutable. The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Immutable. Optional. The virtual cluster config is used when creating a Dataproc cluster that does not directly control the underlying compute resources, for example, when creating a [Dataproc-on-GKE cluster](https://cloud.google.com/dataproc/docs/guides/dpgke/dataproc-gke). Dataproc may set default values, and values may change when clusters are updated. Exactly one of config or virtual_cluster_config must be specified.
	VirtualClusterConfig map[string]interface{} `json:"virtualClusterConfig,omitempty" yaml:"virtualClusterConfig,omitempty"`
}

// DataprocClusterStatus defines the observed state of DataprocCluster.
type DataprocClusterStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
