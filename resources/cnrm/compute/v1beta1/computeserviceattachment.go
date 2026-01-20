package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ComputeServiceAttachment represents a compute.cnrm.cloud.google.com ComputeServiceAttachment resource.
type ComputeServiceAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeServiceAttachmentSpec   `json:"spec,omitempty"`
	Status ComputeServiceAttachmentStatus `json:"status,omitempty"`
}

// ComputeServiceAttachmentSpec defines the desired state of ComputeServiceAttachment.
type ComputeServiceAttachmentSpec struct {
	// The connection preference of service attachment. The value can be set to `ACCEPT_AUTOMATIC`. An `ACCEPT_AUTOMATIC` service attachment is one that always accepts the connection from consumer forwarding rules. Possible values: CONNECTION_PREFERENCE_UNSPECIFIED, ACCEPT_AUTOMATIC, ACCEPT_MANUAL
	ConnectionPreference string `json:"connectionPreference,omitempty" yaml:"connectionPreference,omitempty"`
	// Projects that are allowed to connect to this service attachment.
	ConsumerAcceptLists []map[string]interface{} `json:"consumerAcceptLists,omitempty" yaml:"consumerAcceptLists,omitempty"`
	ConsumerRejectLists []map[string]interface{} `json:"consumerRejectLists,omitempty" yaml:"consumerRejectLists,omitempty"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Immutable. If true, enable the proxy protocol which is for supplying client TCP/IP address data in TCP connections that traverse proxies on their way to destination servers.
	EnableProxyProtocol bool `json:"enableProxyProtocol,omitempty" yaml:"enableProxyProtocol,omitempty"`
	// Immutable. The location for the resource
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	NatSubnets []map[string]interface{} `json:"natSubnets,omitempty" yaml:"natSubnets,omitempty"`
	// Immutable. The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Immutable.
	TargetServiceRef map[string]interface{} `json:"targetServiceRef,omitempty" yaml:"targetServiceRef,omitempty"`
}

// ComputeServiceAttachmentStatus defines the observed state of ComputeServiceAttachment.
type ComputeServiceAttachmentStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
