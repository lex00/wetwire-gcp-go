package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DiscoveryEngineEngine is the Schema for the DiscoveryEngineEngine API
type DiscoveryEngineEngine struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DiscoveryEngineEngineSpec   `json:"spec,omitempty"`
	Status DiscoveryEngineEngineStatus `json:"status,omitempty"`
}

// DiscoveryEngineEngineSpec defines the desired state of DiscoveryEngineEngine.
type DiscoveryEngineEngineSpec struct {
	// Configurations for the Chat Engine. Only applicable if solution_type is SOLUTION_TYPE_CHAT.
	ChatEngineConfig map[string]interface{} `json:"chatEngineConfig,omitempty" yaml:"chatEngineConfig,omitempty"`
	// Immutable. The collection for the Engine.
	Collection string `json:"collection,omitempty" yaml:"collection,omitempty"`
	// Common config spec that specifies the metadata of the engine.
	CommonConfig map[string]interface{} `json:"commonConfig,omitempty" yaml:"commonConfig,omitempty"`
	// The data stores associated with this engine. For SOLUTION_TYPE_SEARCH and SOLUTION_TYPE_RECOMMENDATION type of engines, they can only associate with at most one data store. If solution_type is SOLUTION_TYPE_CHAT, multiple DataStores in the same Collection can be associated here. Note that when used in CreateEngineRequest, one DataStore must be provided as the system will use it for necessary initializations.
	DataStoreRefs []map[string]interface{} `json:"dataStoreRefs,omitempty" yaml:"dataStoreRefs,omitempty"`
	// Optional. Whether to disable analytics for searches performed on this engine.
	DisableAnalytics bool `json:"disableAnalytics,omitempty" yaml:"disableAnalytics,omitempty"`
	// Required. The display name of the engine. Should be human readable. UTF-8 encoded string with limit of 1024 characters.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// The industry vertical that the engine registers. The restriction of the Engine industry vertical is based on DataStore: If unspecified, default to `GENERIC`. Vertical on Engine has to match vertical of the DataStore linked to the engine.
	IndustryVertical string `json:"industryVertical,omitempty" yaml:"industryVertical,omitempty"`
	// Immutable. Location of the resource.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Immutable. The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. The DiscoveryEngineChatEngine name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Configurations for the Search Engine. Only applicable if solution_type is SOLUTION_TYPE_SEARCH.
	SearchEngineConfig map[string]interface{} `json:"searchEngineConfig,omitempty" yaml:"searchEngineConfig,omitempty"`
	// Required. The solutions of the engine.
	SolutionType string `json:"solutionType,omitempty" yaml:"solutionType,omitempty"`
}

// DiscoveryEngineEngineStatus defines the observed state of DiscoveryEngineEngine.
type DiscoveryEngineEngineStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
