package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DiscoveryEngineDataStore is the Schema for the DiscoveryEngineDataStore API
type DiscoveryEngineDataStore struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DiscoveryEngineDataStoreSpec   `json:"spec,omitempty"`
	Status DiscoveryEngineDataStoreStatus `json:"status,omitempty"`
}

// DiscoveryEngineDataStoreSpec defines the desired state of DiscoveryEngineDataStore.
type DiscoveryEngineDataStoreSpec struct {
	// Immutable. The collection for the DataStore.
	Collection string `json:"collection,omitempty" yaml:"collection,omitempty"`
	// Immutable. The content config of the data store. If this field is unset, the server behavior defaults to [ContentConfig.NO_CONTENT][google.cloud.discoveryengine.v1.DataStore.ContentConfig.NO_CONTENT].
	ContentConfig string `json:"contentConfig,omitempty" yaml:"contentConfig,omitempty"`
	// Required. The data store display name.
	// This field must be a UTF-8 encoded string with a length limit of 128
	// characters. Otherwise, an INVALID_ARGUMENT error is returned.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// Immutable. The industry vertical that the data store registers.
	IndustryVertical string `json:"industryVertical,omitempty" yaml:"industryVertical,omitempty"`
	// Immutable. The location for the resource.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// The ID of the project in which the resource belongs.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// The DiscoveryEngineDataStore name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// The solutions that the data store enrolls. Available solutions for each
	// [industry_vertical][google.cloud.discoveryengine.v1.DataStore.industry_vertical]:
	// * `MEDIA`: `SOLUTION_TYPE_RECOMMENDATION` and `SOLUTION_TYPE_SEARCH`.
	// * `SITE_SEARCH`: `SOLUTION_TYPE_SEARCH` is automatically enrolled. Other
	// solutions cannot be enrolled.
	SolutionTypes []string `json:"solutionTypes,omitempty" yaml:"solutionTypes,omitempty"`
	// Config to store data store type configuration for workspace data. This must be set when [DataStore.content_config][google.cloud.discoveryengine.v1.DataStore.content_config] is set as [DataStore.ContentConfig.GOOGLE_WORKSPACE][google.cloud.discoveryengine.v1.DataStore.ContentConfig.GOOGLE_WORKSPACE].
	WorkspaceConfig map[string]interface{} `json:"workspaceConfig,omitempty" yaml:"workspaceConfig,omitempty"`
}

// DiscoveryEngineDataStoreStatus defines the observed state of DiscoveryEngineDataStore.
type DiscoveryEngineDataStoreStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
