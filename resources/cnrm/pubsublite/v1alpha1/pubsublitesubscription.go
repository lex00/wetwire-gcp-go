package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PubSubLiteSubscription represents a pubsublite.cnrm.cloud.google.com PubSubLiteSubscription resource.
type PubSubLiteSubscription struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PubSubLiteSubscriptionSpec   `json:"spec,omitempty"`
	Status PubSubLiteSubscriptionStatus `json:"status,omitempty"`
}

// PubSubLiteSubscriptionSpec defines the desired state of PubSubLiteSubscription.
type PubSubLiteSubscriptionSpec struct {
	// The settings for this subscription's message delivery.
	DeliveryConfig map[string]interface{} `json:"deliveryConfig,omitempty" yaml:"deliveryConfig,omitempty"`
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// The region of the pubsub lite topic.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Immutable. A reference to a Topic resource.
	Topic string `json:"topic,omitempty" yaml:"topic,omitempty"`
	// The zone of the pubsub lite topic.
	Zone string `json:"zone,omitempty" yaml:"zone,omitempty"`
}

// PubSubLiteSubscriptionStatus defines the observed state of PubSubLiteSubscription.
type PubSubLiteSubscriptionStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
