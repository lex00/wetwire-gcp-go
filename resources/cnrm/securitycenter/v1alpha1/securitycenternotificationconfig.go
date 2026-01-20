package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// SecurityCenterNotificationConfig represents a securitycenter.cnrm.cloud.google.com SecurityCenterNotificationConfig resource.
type SecurityCenterNotificationConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SecurityCenterNotificationConfigSpec   `json:"spec,omitempty"`
	Status SecurityCenterNotificationConfigStatus `json:"status,omitempty"`
}

// SecurityCenterNotificationConfigSpec defines the desired state of SecurityCenterNotificationConfig.
type SecurityCenterNotificationConfigSpec struct {
	// Immutable. This must be unique within the organization.
	ConfigID string `json:"configId,omitempty" yaml:"configId,omitempty"`
	// The description of the notification config (max of 1024 characters).
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// The organization that this resource belongs to.
	OrganizationRef map[string]interface{} `json:"organizationRef,omitempty" yaml:"organizationRef,omitempty"`
	// The Pub/Sub topic to send notifications to. Its format is
	// "projects/[project_id]/topics/[topic]".
	PubsubTopic string `json:"pubsubTopic,omitempty" yaml:"pubsubTopic,omitempty"`
	// Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// The config for triggering streaming-based notifications.
	StreamingConfig map[string]interface{} `json:"streamingConfig,omitempty" yaml:"streamingConfig,omitempty"`
}

// SecurityCenterNotificationConfigStatus defines the observed state of SecurityCenterNotificationConfig.
type SecurityCenterNotificationConfigStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
