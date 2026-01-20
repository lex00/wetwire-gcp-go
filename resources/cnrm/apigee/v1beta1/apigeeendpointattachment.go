package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ApigeeEndpointAttachment is the Schema for the ApigeeEndpointAttachment API
type ApigeeEndpointAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ApigeeEndpointAttachmentSpec   `json:"spec,omitempty"`
	Status ApigeeEndpointAttachmentStatus `json:"status,omitempty"`
}

// ApigeeEndpointAttachmentSpec defines the desired state of ApigeeEndpointAttachment.
type ApigeeEndpointAttachmentSpec struct {
	// Required. Location of the endpoint attachment.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Reference to parent Apigee Organization.
	OrganizationRef map[string]interface{} `json:"organizationRef,omitempty" yaml:"organizationRef,omitempty"`
	// The ApigeeEndpointAttachment name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Reference to the ServiceAttachment for the EndpointAttachment.
	ServiceAttachmentRef map[string]interface{} `json:"serviceAttachmentRef,omitempty" yaml:"serviceAttachmentRef,omitempty"`
}

// ApigeeEndpointAttachmentStatus defines the observed state of ApigeeEndpointAttachment.
type ApigeeEndpointAttachmentStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
