package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ComputeNetworkAttachment is the Schema for the ComputeNetworkAttachment API
type ComputeNetworkAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeNetworkAttachmentSpec   `json:"spec,omitempty"`
	Status ComputeNetworkAttachmentStatus `json:"status,omitempty"`
}

// ComputeNetworkAttachmentSpec defines the desired state of ComputeNetworkAttachment.
type ComputeNetworkAttachmentSpec struct {
	// Check the ConnectionPreference enum for the list of possible values.
	ConnectionPreference string `json:"connectionPreference,omitempty" yaml:"connectionPreference,omitempty"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. An up-to-date fingerprint must be provided in order to patch.
	Fingerprint string `json:"fingerprint,omitempty" yaml:"fingerprint,omitempty"`
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Projects that are allowed to connect to this network attachment. The project can be specified using its id or number.
	ProducerAcceptLists []map[string]interface{} `json:"producerAcceptLists,omitempty" yaml:"producerAcceptLists,omitempty"`
	// Projects that are not allowed to connect to this network attachment. The project can be specified using its id or number.
	ProducerRejectLists []map[string]interface{} `json:"producerRejectLists,omitempty" yaml:"producerRejectLists,omitempty"`
	// The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// The ComputeNetworkAttachment name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// An array of URLs where each entry is the URL of a subnet provided by the service consumer to use for endpoints in the producers that connect to this network attachment.
	SubnetworkRefs []map[string]interface{} `json:"subnetworkRefs,omitempty" yaml:"subnetworkRefs,omitempty"`
}

// ComputeNetworkAttachmentStatus defines the observed state of ComputeNetworkAttachment.
type ComputeNetworkAttachmentStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
