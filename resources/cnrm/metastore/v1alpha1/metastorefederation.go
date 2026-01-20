package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MetastoreFederation is the Schema for the MetastoreFederation API
type MetastoreFederation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MetastoreFederationSpec   `json:"spec,omitempty"`
	Status MetastoreFederationStatus `json:"status,omitempty"`
}

// MetastoreFederationSpec defines the desired state of MetastoreFederation.
type MetastoreFederationSpec struct {
	// map type with key int32 as string and value as BackendMetastore
	BackendMetastores map[string]map[string]interface{} `json:"backendMetastores,omitempty" yaml:"backendMetastores,omitempty"`
	// User-defined labels for the metastore federation.
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// The MetastoreFederation name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Immutable. The Apache Hive metastore version of the federation. All backend metastore versions must be compatible with the federation version.
	Version string `json:"version,omitempty" yaml:"version,omitempty"`
}

// MetastoreFederationStatus defines the observed state of MetastoreFederation.
type MetastoreFederationStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
