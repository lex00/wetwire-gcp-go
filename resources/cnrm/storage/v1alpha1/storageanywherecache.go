package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// StorageAnywhereCache is the Schema for the StorageAnywhereCache API
type StorageAnywhereCache struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   StorageAnywhereCacheSpec   `json:"spec,omitempty"`
	Status StorageAnywhereCacheStatus `json:"status,omitempty"`
}

// StorageAnywhereCacheSpec defines the desired state of StorageAnywhereCache.
type StorageAnywhereCacheSpec struct {
	// Cache admission policy. Valid values include: `admit-on-first-miss` and `admit-on-second-miss`. Defaults to `admit-on-first-miss`.
	AdmissionPolicy string `json:"admissionPolicy,omitempty" yaml:"admissionPolicy,omitempty"`
	// Immutable. The reference to bucket where cache needs to be created.
	BucketRef map[string]interface{} `json:"bucketRef,omitempty" yaml:"bucketRef,omitempty"`
	// The desired state of the cache. Possible values include "running", "disabled", and "paused". If not specified, the default value is "running". This field controls the runtime behavior of the cache. Please note that changes to the `desiredState` are prioritized over any other updates. For instance, if both the `desiredState` and `ttl` are updated simultaneously, the state would be updated first, followed by `ttl`.
	DesiredState string `json:"desiredState,omitempty" yaml:"desiredState,omitempty"`
	// The AnywhereCacheID generated via backend. It can be used by users to manage an existing cache.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Cache entry TTL (ranges between 1h to 7d). This is a cache-level config that defines how long a cache entry can live. Defaults to "86400s". TTL must be in whole seconds.
	Ttl string `json:"ttl,omitempty" yaml:"ttl,omitempty"`
	// Immutable. The zone in which the cache instance needs to be created. For example, us-central1-a.
	Zone string `json:"zone,omitempty" yaml:"zone,omitempty"`
}

// StorageAnywhereCacheStatus defines the observed state of StorageAnywhereCache.
type StorageAnywhereCacheStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
