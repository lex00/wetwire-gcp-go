package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DialogflowCXIntent represents a dialogflowcx.cnrm.cloud.google.com DialogflowCXIntent resource.
type DialogflowCXIntent struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DialogflowCXIntentSpec   `json:"spec,omitempty"`
	Status DialogflowCXIntentStatus `json:"status,omitempty"`
}

// DialogflowCXIntentSpec defines the desired state of DialogflowCXIntent.
type DialogflowCXIntentSpec struct {
	// Human readable description for better understanding an intent like its scope, content, result etc. Maximum character limit: 140 characters.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// The human-readable name of the intent, unique within the agent.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// Indicates whether this is a fallback intent. Currently only default fallback intent is allowed in the agent, which is added upon agent creation.
	// Adding training phrases to fallback intent is useful in the case of requests that are mistakenly matched, since training phrases assigned to fallback intents act as negative examples that triggers no-match event.
	IsFallback bool `json:"isFallback,omitempty" yaml:"isFallback,omitempty"`
	// Immutable. The language of the following fields in intent:
	// Intent.training_phrases.parts.text
	// If not specified, the agent's default language is used. Many languages are supported. Note: languages must be enabled in the agent before they can be used.
	LanguageCode string `json:"languageCode,omitempty" yaml:"languageCode,omitempty"`
	// The collection of parameters associated with the intent.
	Parameters []map[string]interface{} `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	// Immutable. The agent to create an intent for.
	// Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>.
	Parent string `json:"parent,omitempty" yaml:"parent,omitempty"`
	// The priority of this intent. Higher numbers represent higher priorities.
	// If the supplied value is unspecified or 0, the service translates the value to 500,000, which corresponds to the Normal priority in the console.
	// If the supplied value is negative, the intent is ignored in runtime detect intent requests.
	Priority int32 `json:"priority,omitempty" yaml:"priority,omitempty"`
	// Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// The collection of training phrases the agent is trained on to identify the intent.
	TrainingPhrases []map[string]interface{} `json:"trainingPhrases,omitempty" yaml:"trainingPhrases,omitempty"`
}

// DialogflowCXIntentStatus defines the observed state of DialogflowCXIntent.
type DialogflowCXIntentStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
