package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DialogflowCXFlow represents a dialogflowcx.cnrm.cloud.google.com DialogflowCXFlow resource.
type DialogflowCXFlow struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DialogflowCXFlowSpec   `json:"spec,omitempty"`
	Status DialogflowCXFlowStatus `json:"status,omitempty"`
}

// DialogflowCXFlowSpec defines the desired state of DialogflowCXFlow.
type DialogflowCXFlowSpec struct {
	// The description of the flow. The maximum length is 500 characters. If exceeded, the request is rejected.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// The human-readable name of the flow.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// A flow's event handlers serve two purposes:
	// They are responsible for handling events (e.g. no match, webhook errors) in the flow.
	// They are inherited by every page's [event handlers][Page.event_handlers], which can be used to handle common events regardless of the current page. Event handlers defined in the page have higher priority than those defined in the flow.
	// Unlike transitionRoutes, these handlers are evaluated on a first-match basis. The first one that matches the event get executed, with the rest being ignored.
	EventHandlers []map[string]interface{} `json:"eventHandlers,omitempty" yaml:"eventHandlers,omitempty"`
	// Immutable. The language of the following fields in flow:
	// Flow.event_handlers.trigger_fulfillment.messages
	// Flow.event_handlers.trigger_fulfillment.conditional_cases
	// Flow.transition_routes.trigger_fulfillment.messages
	// Flow.transition_routes.trigger_fulfillment.conditional_cases
	// If not specified, the agent's default language is used. Many languages are supported. Note: languages must be enabled in the agent before they can be used.
	LanguageCode string `json:"languageCode,omitempty" yaml:"languageCode,omitempty"`
	// NLU related settings of the flow.
	NluSettings map[string]interface{} `json:"nluSettings,omitempty" yaml:"nluSettings,omitempty"`
	// Immutable. The agent to create a flow for.
	// Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>.
	Parent string `json:"parent,omitempty" yaml:"parent,omitempty"`
	// Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// A flow's transition route group serve two purposes:
	// They are responsible for matching the user's first utterances in the flow.
	// They are inherited by every page's [transition route groups][Page.transition_route_groups]. Transition route groups defined in the page have higher priority than those defined in the flow.
	// Format:projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>/transitionRouteGroups/<TransitionRouteGroup ID>.
	TransitionRouteGroups []string `json:"transitionRouteGroups,omitempty" yaml:"transitionRouteGroups,omitempty"`
	// A flow's transition routes serve two purposes:
	// They are responsible for matching the user's first utterances in the flow.
	// They are inherited by every page's [transition routes][Page.transition_routes] and can support use cases such as the user saying "help" or "can I talk to a human?", which can be handled in a common way regardless of the current page. Transition routes defined in the page have higher priority than those defined in the flow.
	// TransitionRoutes are evalauted in the following order:
	// TransitionRoutes with intent specified.
	// TransitionRoutes with only condition specified.
	// TransitionRoutes with intent specified are inherited by pages in the flow.
	TransitionRoutes []map[string]interface{} `json:"transitionRoutes,omitempty" yaml:"transitionRoutes,omitempty"`
}

// DialogflowCXFlowStatus defines the observed state of DialogflowCXFlow.
type DialogflowCXFlowStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
