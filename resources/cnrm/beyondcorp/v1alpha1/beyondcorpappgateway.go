package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// BeyondCorpAppGateway represents a beyondcorp.cnrm.cloud.google.com BeyondCorpAppGateway resource.
type BeyondCorpAppGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BeyondCorpAppGatewaySpec   `json:"spec,omitempty"`
	Status BeyondCorpAppGatewayStatus `json:"status,omitempty"`
}

// BeyondCorpAppGatewaySpec defines the desired state of BeyondCorpAppGateway.
type BeyondCorpAppGatewaySpec struct {
	// Immutable. An arbitrary user-provided name for the AppGateway.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// Immutable. The type of hosting used by the AppGateway. Default value: "HOST_TYPE_UNSPECIFIED" Possible values: ["HOST_TYPE_UNSPECIFIED", "GCP_REGIONAL_MIG"].
	HostType string `json:"hostType,omitempty" yaml:"hostType,omitempty"`
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. The region of the AppGateway.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Immutable. The type of network connectivity used by the AppGateway. Default value: "TYPE_UNSPECIFIED" Possible values: ["TYPE_UNSPECIFIED", "TCP_PROXY"].
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}

// BeyondCorpAppGatewayStatus defines the observed state of BeyondCorpAppGateway.
type BeyondCorpAppGatewayStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
