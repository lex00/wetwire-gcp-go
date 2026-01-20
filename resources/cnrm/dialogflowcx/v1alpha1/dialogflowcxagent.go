package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DialogflowCXAgent represents a dialogflowcx.cnrm.cloud.google.com DialogflowCXAgent resource.
type DialogflowCXAgent struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DialogflowCXAgentSpec   `json:"spec,omitempty"`
	Status DialogflowCXAgentStatus `json:"status,omitempty"`
}

// DialogflowCXAgentSpec defines the desired state of DialogflowCXAgent.
type DialogflowCXAgentSpec struct {
	// The URI of the agent's avatar. Avatars are used throughout the Dialogflow console and in the self-hosted Web Demo integration.
	AvatarURI string `json:"avatarUri,omitempty" yaml:"avatarUri,omitempty"`
	// Immutable. The default language of the agent as a language tag. [See Language Support](https://cloud.google.com/dialogflow/cx/docs/reference/language)
	// for a list of the currently supported language codes. This field cannot be updated after creation.
	DefaultLanguageCode string `json:"defaultLanguageCode,omitempty" yaml:"defaultLanguageCode,omitempty"`
	// The description of this agent. The maximum length is 500 characters. If exceeded, the request is rejected.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// The human-readable name of the agent, unique within the location.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// Indicates if automatic spell correction is enabled in detect intent requests.
	EnableSpellCorrection bool `json:"enableSpellCorrection,omitempty" yaml:"enableSpellCorrection,omitempty"`
	// Determines whether this agent should log conversation queries.
	EnableStackdriverLogging bool `json:"enableStackdriverLogging,omitempty" yaml:"enableStackdriverLogging,omitempty"`
	// Immutable. The name of the location this agent is located in.
	// ~> **Note:** The first time you are deploying an Agent in your project you must configure location settings.
	// This is a one time step but at the moment you can only [configure location settings](https://cloud.google.com/dialogflow/cx/docs/concept/region#location-settings) via the Dialogflow CX console.
	// Another options is to use global location so you don't need to manually configure location settings.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Name of the SecuritySettings reference for the agent. Format: projects/<Project ID>/locations/<Location ID>/securitySettings/<Security Settings ID>.
	SecuritySettings string `json:"securitySettings,omitempty" yaml:"securitySettings,omitempty"`
	// Settings related to speech recognition.
	SpeechToTextSettings map[string]interface{} `json:"speechToTextSettings,omitempty" yaml:"speechToTextSettings,omitempty"`
	// The list of all languages supported by this agent (except for the default_language_code).
	SupportedLanguageCodes []string `json:"supportedLanguageCodes,omitempty" yaml:"supportedLanguageCodes,omitempty"`
	// The time zone of this agent from the [time zone database](https://www.iana.org/time-zones), e.g., America/New_York,
	// Europe/Paris.
	TimeZone string `json:"timeZone,omitempty" yaml:"timeZone,omitempty"`
}

// DialogflowCXAgentStatus defines the observed state of DialogflowCXAgent.
type DialogflowCXAgentStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
