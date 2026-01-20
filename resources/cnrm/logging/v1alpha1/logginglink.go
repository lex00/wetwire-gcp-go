package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// LoggingLink is the Schema for the LoggingLink API
type LoggingLink struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LoggingLinkSpec   `json:"spec,omitempty"`
	Status LoggingLinkStatus `json:"status,omitempty"`
}

// LoggingLinkSpec defines the desired state of LoggingLink.
type LoggingLinkSpec struct {
	// Describes this link.
	// The maximum length of the description is 8000 characters.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Required. The LoggingLogBucket that this Link is associated with.
	LoggingLogBucketRef map[string]interface{} `json:"loggingLogBucketRef,omitempty" yaml:"loggingLogBucketRef,omitempty"`
	// Immutable. The LoggingLink name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// LoggingLinkStatus defines the observed state of LoggingLink.
type LoggingLinkStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
