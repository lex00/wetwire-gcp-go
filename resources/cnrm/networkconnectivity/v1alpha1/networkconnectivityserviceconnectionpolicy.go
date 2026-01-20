package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NetworkConnectivityServiceConnectionPolicy is the Schema for the NetworkConnectivityServiceConnectionPolicy API
type NetworkConnectivityServiceConnectionPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NetworkConnectivityServiceConnectionPolicySpec   `json:"spec,omitempty"`
	Status NetworkConnectivityServiceConnectionPolicyStatus `json:"status,omitempty"`
}

// NetworkConnectivityServiceConnectionPolicySpec defines the desired state of NetworkConnectivityServiceConnectionPolicy.
type NetworkConnectivityServiceConnectionPolicySpec struct {
	// A description of this resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Immutable. Location of the resource.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// The resource path of the consumer network. Example: - projects/{projectNumOrId}/global/networks/{resourceId}.
	NetworkRef map[string]interface{} `json:"networkRef,omitempty" yaml:"networkRef,omitempty"`
	// Immutable. The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Configuration used for Private Service Connect connections. Used when Infrastructure is PSC.
	PscConfig map[string]interface{} `json:"pscConfig,omitempty" yaml:"pscConfig,omitempty"`
	// The NetworkConnectivityServiceConnectionPolicy name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// The service class identifier for which this ServiceConnectionPolicy is for. The service class identifier is a unique, symbolic representation of a ServiceClass. It is provided by the Service Producer. Google services have a prefix of gcp. For example, gcp-cloud-sql. 3rd party services do not. For example, test-service-a3dfcx.
	ServiceClass string `json:"serviceClass,omitempty" yaml:"serviceClass,omitempty"`
}

// NetworkConnectivityServiceConnectionPolicyStatus defines the observed state of NetworkConnectivityServiceConnectionPolicy.
type NetworkConnectivityServiceConnectionPolicyStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
