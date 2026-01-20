package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// VMwareEnginePrivateCloud is the Schema for the VMwareEnginePrivateCloud API
type VMwareEnginePrivateCloud struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VMwareEnginePrivateCloudSpec   `json:"spec,omitempty"`
	Status VMwareEnginePrivateCloudStatus `json:"status,omitempty"`
}

// VMwareEnginePrivateCloudSpec defines the desired state of VMwareEnginePrivateCloud.
type VMwareEnginePrivateCloudSpec struct {
	// User-provided description for this private cloud.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Immutable.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Required. Input only. The management cluster for this private cloud.
	// This field is required during creation of the private cloud to provide
	// details for the default cluster.
	// The following fields can't be changed after private cloud creation:
	// `ManagementCluster.clusterId`, `ManagementCluster.nodeTypeId`.
	ManagementCluster map[string]interface{} `json:"managementCluster,omitempty" yaml:"managementCluster,omitempty"`
	// Required. Network configuration of the private cloud.
	NetworkConfig map[string]interface{} `json:"networkConfig,omitempty" yaml:"networkConfig,omitempty"`
	// The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// The VMwareEnginePrivateCloud name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Optional. Type of the private cloud. Defaults to STANDARD.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}

// VMwareEnginePrivateCloudStatus defines the observed state of VMwareEnginePrivateCloud.
type VMwareEnginePrivateCloudStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
