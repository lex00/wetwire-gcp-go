package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DialogflowCXPage represents a dialogflowcx.cnrm.cloud.google.com DialogflowCXPage resource.
type DialogflowCXPage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DialogflowCXPageSpec   `json:"spec,omitempty"`
	Status DialogflowCXPageStatus `json:"status,omitempty"`
}

// DialogflowCXPageSpec defines the desired state of DialogflowCXPage.
type DialogflowCXPageSpec struct {
	// The human-readable name of the page, unique within the agent.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// The fulfillment to call when the session is entering the page.
	EntryFulfillment map[string]interface{} `json:"entryFulfillment,omitempty" yaml:"entryFulfillment,omitempty"`
	// Handlers associated with the page to handle events such as webhook errors, no match or no input.
	EventHandlers []map[string]interface{} `json:"eventHandlers,omitempty" yaml:"eventHandlers,omitempty"`
	// The form associated with the page, used for collecting parameters relevant to the page.
	Form map[string]interface{} `json:"form,omitempty" yaml:"form,omitempty"`
	// Immutable. The language of the following fields in page:
	// Page.entry_fulfillment.messages
	// Page.entry_fulfillment.conditional_cases
	// Page.event_handlers.trigger_fulfillment.messages
	// Page.event_handlers.trigger_fulfillment.conditional_cases
	// Page.form.parameters.fill_behavior.initial_prompt_fulfillment.messages
	// Page.form.parameters.fill_behavior.initial_prompt_fulfillment.conditional_cases
	// Page.form.parameters.fill_behavior.reprompt_event_handlers.messages
	// Page.form.parameters.fill_behavior.reprompt_event_handlers.conditional_cases
	// Page.transition_routes.trigger_fulfillment.messages
	// Page.transition_routes.trigger_fulfillment.conditional_cases
	// If not specified, the agent's default language is used. Many languages are supported. Note: languages must be enabled in the agent before they can be used.
	LanguageCode string `json:"languageCode,omitempty" yaml:"languageCode,omitempty"`
	// Immutable. The flow to create a page for.
	// Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>.
	Parent string `json:"parent,omitempty" yaml:"parent,omitempty"`
	// Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Ordered list of TransitionRouteGroups associated with the page. Transition route groups must be unique within a page.
	// If multiple transition routes within a page scope refer to the same intent, then the precedence order is: page's transition route -> page's transition route group -> flow's transition routes.
	// If multiple transition route groups within a page contain the same intent, then the first group in the ordered list takes precedence.
	// Format:projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>/transitionRouteGroups/<TransitionRouteGroup ID>.
	TransitionRouteGroups []string `json:"transitionRouteGroups,omitempty" yaml:"transitionRouteGroups,omitempty"`
	// A list of transitions for the transition rules of this page. They route the conversation to another page in the same flow, or another flow.
	// When we are in a certain page, the TransitionRoutes are evalauted in the following order:
	// TransitionRoutes defined in the page with intent specified.
	// TransitionRoutes defined in the transition route groups with intent specified.
	// TransitionRoutes defined in flow with intent specified.
	// TransitionRoutes defined in the transition route groups with intent specified.
	// TransitionRoutes defined in the page with only condition specified.
	// TransitionRoutes defined in the transition route groups with only condition specified.
	TransitionRoutes []map[string]interface{} `json:"transitionRoutes,omitempty" yaml:"transitionRoutes,omitempty"`
}

// DialogflowCXPageStatus defines the observed state of DialogflowCXPage.
type DialogflowCXPageStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
