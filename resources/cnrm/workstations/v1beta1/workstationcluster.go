package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// WorkstationCluster is the Schema for the WorkstationCluster API
type WorkstationCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WorkstationClusterSpec   `json:"spec,omitempty"`
	Status WorkstationClusterStatus `json:"status,omitempty"`
}

// WorkstationClusterSpec defines the desired state of WorkstationCluster.
type WorkstationClusterSpec struct {
	// Optional. Client-specified annotations.
	Annotations []map[string]interface{} `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	// Optional. Human-readable name for this workstation cluster.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// Optional. [Labels](https://cloud.google.com/workstations/docs/label-resources) that are applied to the workstation cluster and that are also propagated to the underlying Compute Engine resources.
	Labels []map[string]interface{} `json:"labels,omitempty" yaml:"labels,omitempty"`
	// The location of the cluster.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Immutable. Reference to the Compute Engine network in which instances associated with this workstation cluster will be created.
	NetworkRef map[string]interface{} `json:"networkRef,omitempty" yaml:"networkRef,omitempty"`
	// Optional. Configuration for private workstation cluster.
	PrivateClusterConfig map[string]interface{} `json:"privateClusterConfig,omitempty" yaml:"privateClusterConfig,omitempty"`
	// Immutable. The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// The WorkstationCluster name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Immutable. Reference to the Compute Engine subnetwork in which instances associated with this workstation cluster will be created. Must be part of the subnetwork specified for this workstation cluster.
	SubnetworkRef map[string]interface{} `json:"subnetworkRef,omitempty" yaml:"subnetworkRef,omitempty"`
}

// WorkstationClusterStatus defines the observed state of WorkstationCluster.
type WorkstationClusterStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
