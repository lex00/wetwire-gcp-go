package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// AppEngineServiceSplitTraffic represents a appengine.cnrm.cloud.google.com AppEngineServiceSplitTraffic resource.
type AppEngineServiceSplitTraffic struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AppEngineServiceSplitTrafficSpec   `json:"spec,omitempty"`
	Status AppEngineServiceSplitTrafficStatus `json:"status,omitempty"`
}

// AppEngineServiceSplitTrafficSpec defines the desired state of AppEngineServiceSplitTraffic.
type AppEngineServiceSplitTrafficSpec struct {
	// If set to true traffic will be migrated to this version.
	MigrateTraffic bool `json:"migrateTraffic,omitempty" yaml:"migrateTraffic,omitempty"`
	// Immutable.
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
	// Immutable. Optional. The service of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Mapping that defines fractional HTTP traffic diversion to different versions within the service.
	Split map[string]interface{} `json:"split,omitempty" yaml:"split,omitempty"`
}

// AppEngineServiceSplitTrafficStatus defines the observed state of AppEngineServiceSplitTraffic.
type AppEngineServiceSplitTrafficStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
