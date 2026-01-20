package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DialogflowIntent represents a dialogflow.cnrm.cloud.google.com DialogflowIntent resource.
type DialogflowIntent struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DialogflowIntentSpec   `json:"spec,omitempty"`
	Status DialogflowIntentStatus `json:"status,omitempty"`
}

// DialogflowIntentSpec defines the desired state of DialogflowIntent.
type DialogflowIntentSpec struct {
	// The name of the action associated with the intent.
	// Note: The action name must not contain whitespaces.
	Action string `json:"action,omitempty" yaml:"action,omitempty"`
	// The list of platforms for which the first responses will be copied from the messages in PLATFORM_UNSPECIFIED
	// (i.e. default platform). Possible values: ["FACEBOOK", "SLACK", "TELEGRAM", "KIK", "SKYPE", "LINE", "VIBER", "ACTIONS_ON_GOOGLE", "GOOGLE_HANGOUTS"].
	DefaultResponsePlatforms []string `json:"defaultResponsePlatforms,omitempty" yaml:"defaultResponsePlatforms,omitempty"`
	// The name of this intent to be displayed on the console.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// The collection of event names that trigger the intent. If the collection of input contexts is not empty, all of
	// the contexts must be present in the active user session for an event to trigger this intent. See the
	// [events reference](https://cloud.google.com/dialogflow/docs/events-overview) for more details.
	Events []string `json:"events,omitempty" yaml:"events,omitempty"`
	// The list of context names required for this intent to be triggered.
	// Format: projects/<Project ID>/agent/sessions/-/contexts/<Context ID>.
	InputContextNames []string `json:"inputContextNames,omitempty" yaml:"inputContextNames,omitempty"`
	// Indicates whether this is a fallback intent.
	IsFallback bool `json:"isFallback,omitempty" yaml:"isFallback,omitempty"`
	// Indicates whether Machine Learning is disabled for the intent.
	// Note: If mlDisabled setting is set to true, then this intent is not taken into account during inference in ML
	// ONLY match mode. Also, auto-markup in the UI is turned off.
	MlDisabled bool `json:"mlDisabled,omitempty" yaml:"mlDisabled,omitempty"`
	// Immutable. The unique identifier of the parent intent in the chain of followup intents.
	// Format: projects/<Project ID>/agent/intents/<Intent ID>.
	ParentFollowupIntentName string `json:"parentFollowupIntentName,omitempty" yaml:"parentFollowupIntentName,omitempty"`
	// The priority of this intent. Higher numbers represent higher priorities.
	// - If the supplied value is unspecified or 0, the service translates the value to 500,000, which corresponds
	// to the Normal priority in the console.
	// - If the supplied value is negative, the intent is ignored in runtime detect intent requests.
	Priority int32 `json:"priority,omitempty" yaml:"priority,omitempty"`
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Indicates whether to delete all contexts in the current session when this intent is matched.
	ResetContexts bool `json:"resetContexts,omitempty" yaml:"resetContexts,omitempty"`
	// Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Indicates whether webhooks are enabled for the intent.
	// * WEBHOOK_STATE_ENABLED: Webhook is enabled in the agent and in the intent.
	// * WEBHOOK_STATE_ENABLED_FOR_SLOT_FILLING: Webhook is enabled in the agent and in the intent. Also, each slot
	// filling prompt is forwarded to the webhook. Possible values: ["WEBHOOK_STATE_ENABLED", "WEBHOOK_STATE_ENABLED_FOR_SLOT_FILLING"].
	WebhookState string `json:"webhookState,omitempty" yaml:"webhookState,omitempty"`
}

// DialogflowIntentStatus defines the observed state of DialogflowIntent.
type DialogflowIntentStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
