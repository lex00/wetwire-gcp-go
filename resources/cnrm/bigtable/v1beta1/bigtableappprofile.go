package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// BigtableAppProfile is the Schema for the BigtableAppProfile API
type BigtableAppProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BigtableAppProfileSpec   `json:"spec,omitempty"`
	Status BigtableAppProfileStatus `json:"status,omitempty"`
}

// BigtableAppProfileSpec defines the desired state of BigtableAppProfile.
type BigtableAppProfileSpec struct {
	// Specifies that this app profile is intended for read-only usage via the Data Boost feature. Please opt-in to this feature by setting the `alpha.cnrm.cloud.google.com/reconciler: direct` annotation.
	DataBoostIsolationReadOnly map[string]interface{} `json:"dataBoostIsolationReadOnly,omitempty" yaml:"dataBoostIsolationReadOnly,omitempty"`
	// Long form description of the use case for this AppProfile.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// InstanceRef defines the resource reference to BigtableInstance, which "External" field holds the GCP identifier for the KRM object.
	InstanceRef map[string]interface{} `json:"instanceRef,omitempty" yaml:"instanceRef,omitempty"`
	// The set of clusters to route to, if using multi cluster routing. The order is ignored; clusters will be tried in order of distance. If left empty, all clusters are eligible.
	MultiClusterRoutingClusterIds []string `json:"multiClusterRoutingClusterIds,omitempty" yaml:"multiClusterRoutingClusterIds,omitempty"`
	// Use a multi-cluster routing policy.
	MultiClusterRoutingUseAny bool `json:"multiClusterRoutingUseAny,omitempty" yaml:"multiClusterRoutingUseAny,omitempty"`
	// The BigtableAppProfile name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Use a single-cluster routing policy.
	SingleClusterRouting map[string]interface{} `json:"singleClusterRouting,omitempty" yaml:"singleClusterRouting,omitempty"`
	// The standard options used for isolating this app profile's traffic from other use cases.
	StandardIsolation map[string]interface{} `json:"standardIsolation,omitempty" yaml:"standardIsolation,omitempty"`
}

// BigtableAppProfileStatus defines the observed state of BigtableAppProfile.
type BigtableAppProfileStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
