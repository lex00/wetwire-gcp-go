package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DialogflowCXWebhook represents a dialogflowcx.cnrm.cloud.google.com DialogflowCXWebhook resource.
type DialogflowCXWebhook struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DialogflowCXWebhookSpec   `json:"spec,omitempty"`
	Status DialogflowCXWebhookStatus `json:"status,omitempty"`
}

// DialogflowCXWebhookSpec defines the desired state of DialogflowCXWebhook.
type DialogflowCXWebhookSpec struct {
	// Indicates whether the webhook is disabled.
	Disabled bool `json:"disabled,omitempty" yaml:"disabled,omitempty"`
	// The human-readable name of the webhook, unique within the agent.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// Indicates if automatic spell correction is enabled in detect intent requests.
	EnableSpellCorrection bool `json:"enableSpellCorrection,omitempty" yaml:"enableSpellCorrection,omitempty"`
	// Determines whether this agent should log conversation queries.
	EnableStackdriverLogging bool `json:"enableStackdriverLogging,omitempty" yaml:"enableStackdriverLogging,omitempty"`
	// Configuration for a generic web service.
	GenericWebService map[string]interface{} `json:"genericWebService,omitempty" yaml:"genericWebService,omitempty"`
	// Immutable. The agent to create a webhook for.
	// Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>.
	Parent string `json:"parent,omitempty" yaml:"parent,omitempty"`
	// Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Name of the SecuritySettings reference for the agent. Format: projects/<Project ID>/locations/<Location ID>/securitySettings/<Security Settings ID>.
	SecuritySettings string `json:"securitySettings,omitempty" yaml:"securitySettings,omitempty"`
	// Configuration for a Service Directory service.
	ServiceDirectory map[string]interface{} `json:"serviceDirectory,omitempty" yaml:"serviceDirectory,omitempty"`
	// Webhook execution timeout.
	Timeout string `json:"timeout,omitempty" yaml:"timeout,omitempty"`
}

// DialogflowCXWebhookStatus defines the observed state of DialogflowCXWebhook.
type DialogflowCXWebhookStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
