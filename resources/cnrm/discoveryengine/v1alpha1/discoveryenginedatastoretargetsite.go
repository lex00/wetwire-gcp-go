package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DiscoveryEngineDataStoreTargetSite is the Schema for the DiscoveryEngineDataStoreTargetSite API
type DiscoveryEngineDataStoreTargetSite struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DiscoveryEngineDataStoreTargetSiteSpec   `json:"spec,omitempty"`
	Status DiscoveryEngineDataStoreTargetSiteStatus `json:"status,omitempty"`
}

// DiscoveryEngineDataStoreTargetSiteSpec defines the desired state of DiscoveryEngineDataStoreTargetSite.
type DiscoveryEngineDataStoreTargetSiteSpec struct {
	// The DataStore this target site should be part of.
	DataStoreRef map[string]interface{} `json:"dataStoreRef,omitempty" yaml:"dataStoreRef,omitempty"`
	// Input only. If set to false, a uri_pattern is generated to include all pages whose address contains the provided_uri_pattern. If set to true, an uri_pattern is generated to try to be an exact match of the provided_uri_pattern or just the specific page if the provided_uri_pattern is a specific one. provided_uri_pattern is always normalized to generate the URI pattern to be used by the search engine.
	ExactMatch bool `json:"exactMatch,omitempty" yaml:"exactMatch,omitempty"`
	// Required. Input only. The user provided URI pattern from which the `generated_uri_pattern` is generated.
	ProvidedURIPattern string `json:"providedURIPattern,omitempty" yaml:"providedURIPattern,omitempty"`
	// The type of the target site, e.g., whether the site is to be included or excluded.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}

// DiscoveryEngineDataStoreTargetSiteStatus defines the observed state of DiscoveryEngineDataStoreTargetSite.
type DiscoveryEngineDataStoreTargetSiteStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
