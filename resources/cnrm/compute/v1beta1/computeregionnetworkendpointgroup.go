package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ComputeRegionNetworkEndpointGroup represents a compute.cnrm.cloud.google.com ComputeRegionNetworkEndpointGroup resource.
type ComputeRegionNetworkEndpointGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeRegionNetworkEndpointGroupSpec   `json:"spec,omitempty"`
	Status ComputeRegionNetworkEndpointGroupStatus `json:"status,omitempty"`
}

// ComputeRegionNetworkEndpointGroupSpec defines the desired state of ComputeRegionNetworkEndpointGroup.
type ComputeRegionNetworkEndpointGroupSpec struct {
	// Immutable. Only valid when networkEndpointType is "SERVERLESS".
	// Only one of cloud_run, app_engine, cloud_function or serverless_deployment may be set.
	CloudFunction map[string]interface{} `json:"cloudFunction,omitempty" yaml:"cloudFunction,omitempty"`
	// Immutable. Only valid when networkEndpointType is "SERVERLESS".
	// Only one of cloud_run, app_engine, cloud_function or serverless_deployment may be set.
	CloudRun map[string]interface{} `json:"cloudRun,omitempty" yaml:"cloudRun,omitempty"`
	// Immutable. An optional description of this resource. Provide this property when
	// you create the resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Immutable. Type of network endpoints in this network endpoint group. Defaults to SERVERLESS Default value: "SERVERLESS" Possible values: ["SERVERLESS", "PRIVATE_SERVICE_CONNECT"].
	NetworkEndpointType string `json:"networkEndpointType,omitempty" yaml:"networkEndpointType,omitempty"`
	// Immutable. This field is only used for PSC.
	// The URL of the network to which all network endpoints in the NEG belong. Uses
	// "default" project network if unspecified.
	NetworkRef map[string]interface{} `json:"networkRef,omitempty" yaml:"networkRef,omitempty"`
	// Immutable. The target service url used to set up private service connection to
	// a Google API or a PSC Producer Service Attachment.
	PscTargetService string `json:"pscTargetService,omitempty" yaml:"pscTargetService,omitempty"`
	// Immutable. A reference to the region where the Serverless NEGs Reside.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Immutable. This field is only used for PSC.
	// Optional URL of the subnetwork to which all network endpoints in the NEG belong.
	SubnetworkRef map[string]interface{} `json:"subnetworkRef,omitempty" yaml:"subnetworkRef,omitempty"`
}

// ComputeRegionNetworkEndpointGroupStatus defines the observed state of ComputeRegionNetworkEndpointGroup.
type ComputeRegionNetworkEndpointGroupStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
