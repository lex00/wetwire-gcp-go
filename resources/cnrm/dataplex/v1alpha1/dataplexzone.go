package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DataplexZone is the Schema for the DataplexZone API
type DataplexZone struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DataplexZoneSpec   `json:"spec,omitempty"`
	Status DataplexZoneStatus `json:"status,omitempty"`
}

// DataplexZoneSpec defines the desired state of DataplexZone.
type DataplexZoneSpec struct {
	// Optional. Description of the zone.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Optional. Specification of the discovery feature applied to data in this zone.
	DiscoverySpec map[string]interface{} `json:"discoverySpec,omitempty" yaml:"discoverySpec,omitempty"`
	// Optional. User friendly display name.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// Reference to the parent DataplexLake that owns this Zone.
	LakeRef map[string]interface{} `json:"lakeRef,omitempty" yaml:"lakeRef,omitempty"`
	// The DataplexZone name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Required. Specification of the resources that are referenced by the assets within this zone.
	ResourceSpec map[string]interface{} `json:"resourceSpec,omitempty" yaml:"resourceSpec,omitempty"`
	// Required. Immutable. The type of the zone.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}

// DataplexZoneStatus defines the observed state of DataplexZone.
type DataplexZoneStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
