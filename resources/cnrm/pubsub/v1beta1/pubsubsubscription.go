package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PubSubSubscription represents a pubsub.cnrm.cloud.google.com PubSubSubscription resource.
type PubSubSubscription struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PubSubSubscriptionSpec   `json:"spec,omitempty"`
	Status PubSubSubscriptionStatus `json:"status,omitempty"`
}

// PubSubSubscriptionSpec defines the desired state of PubSubSubscription.
type PubSubSubscriptionSpec struct {
	// This value is the maximum time after a subscriber receives a message
	// before the subscriber should acknowledge the message. After message
	// delivery but before the ack deadline expires and before the message is
	// acknowledged, it is an outstanding message and will not be delivered
	// again during that time (on a best-effort basis).
	// For pull subscriptions, this value is used as the initial value for
	// the ack deadline. To override this value for a given message, call
	// subscriptions.modifyAckDeadline with the corresponding ackId if using
	// pull. The minimum custom deadline you can specify is 10 seconds. The
	// maximum custom deadline you can specify is 600 seconds (10 minutes).
	// If this parameter is 0, a default value of 10 seconds is used.
	// For push delivery, this value is also used to set the request timeout
	// for the call to the push endpoint.
	// If the subscriber never acknowledges the message, the Pub/Sub system
	// will eventually redeliver the message.
	AckDeadlineSeconds int32 `json:"ackDeadlineSeconds,omitempty" yaml:"ackDeadlineSeconds,omitempty"`
	// If delivery to BigQuery is used with this subscription, this field is used to configure it.
	// Either pushConfig, bigQueryConfig or cloudStorageConfig can be set, but not combined.
	// If all three are empty, then the subscriber will pull and ack messages using API methods.
	BigqueryConfig map[string]interface{} `json:"bigqueryConfig,omitempty" yaml:"bigqueryConfig,omitempty"`
	// If delivery to Cloud Storage is used with this subscription, this field is used to configure it.
	// Either pushConfig, bigQueryConfig or cloudStorageConfig can be set, but not combined.
	// If all three are empty, then the subscriber will pull and ack messages using API methods.
	CloudStorageConfig map[string]interface{} `json:"cloudStorageConfig,omitempty" yaml:"cloudStorageConfig,omitempty"`
	// A policy that specifies the conditions for dead lettering messages in
	// this subscription. If dead_letter_policy is not set, dead lettering
	// is disabled.
	// The Cloud Pub/Sub service account associated with this subscription's
	// parent project (i.e.,
	// service-{project_number}@gcp-sa-pubsub.iam.gserviceaccount.com) must have
	// permission to Acknowledge() messages on this subscription.
	DeadLetterPolicy map[string]interface{} `json:"deadLetterPolicy,omitempty" yaml:"deadLetterPolicy,omitempty"`
	// If 'true', Pub/Sub provides the following guarantees for the delivery
	// of a message with a given value of messageId on this Subscriptions':
	// - The message sent to a subscriber is guaranteed not to be resent before the message's acknowledgement deadline expires.
	// - An acknowledged message will not be resent to a subscriber.
	// Note that subscribers may still receive multiple copies of a message when 'enable_exactly_once_delivery'
	// is true if the message was published multiple times by a publisher client. These copies are considered distinct by Pub/Sub and have distinct messageId values.
	EnableExactlyOnceDelivery bool `json:"enableExactlyOnceDelivery,omitempty" yaml:"enableExactlyOnceDelivery,omitempty"`
	// Immutable. If 'true', messages published with the same orderingKey in PubsubMessage will be delivered to
	// the subscribers in the order in which they are received by the Pub/Sub system. Otherwise, they
	// may be delivered in any order.
	EnableMessageOrdering bool `json:"enableMessageOrdering,omitempty" yaml:"enableMessageOrdering,omitempty"`
	// A policy that specifies the conditions for this subscription's expiration.
	// A subscription is considered active as long as any connected subscriber
	// is successfully consuming messages from the subscription or is issuing
	// operations on the subscription. If expirationPolicy is not set, a default
	// policy with ttl of 31 days will be used.  If it is set but ttl is "", the
	// resource never expires.  The minimum allowed value for expirationPolicy.ttl
	// is 1 day.
	ExpirationPolicy map[string]interface{} `json:"expirationPolicy,omitempty" yaml:"expirationPolicy,omitempty"`
	// Immutable. The subscription only delivers the messages that match the filter.
	// Pub/Sub automatically acknowledges the messages that don't match the filter. You can filter messages
	// by their attributes. The maximum length of a filter is 256 bytes. After creating the subscription,
	// you can't modify the filter.
	Filter string `json:"filter,omitempty" yaml:"filter,omitempty"`
	// How long to retain unacknowledged messages in the subscription's
	// backlog, from the moment a message is published. If
	// retain_acked_messages is true, then this also configures the retention
	// of acknowledged messages, and thus configures how far back in time a
	// subscriptions.seek can be done. Defaults to 7 days. Cannot be more
	// than 7 days ('"604800s"') or less than 10 minutes ('"600s"').
	// A duration in seconds with up to nine fractional digits, terminated
	// by 's'. Example: '"600.5s"'.
	MessageRetentionDuration string `json:"messageRetentionDuration,omitempty" yaml:"messageRetentionDuration,omitempty"`
	// If push delivery is used with this subscription, this field is used to
	// configure it. An empty pushConfig signifies that the subscriber will
	// pull and ack messages using API methods.
	PushConfig map[string]interface{} `json:"pushConfig,omitempty" yaml:"pushConfig,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Indicates whether to retain acknowledged messages. If 'true', then
	// messages are not expunged from the subscription's backlog, even if
	// they are acknowledged, until they fall out of the
	// messageRetentionDuration window.
	RetainAckedMessages bool `json:"retainAckedMessages,omitempty" yaml:"retainAckedMessages,omitempty"`
	// A policy that specifies how Pub/Sub retries message delivery for this subscription.
	// If not set, the default retry policy is applied. This generally implies that messages will be retried as soon as possible for healthy subscribers.
	// RetryPolicy will be triggered on NACKs or acknowledgement deadline exceeded events for a given message.
	RetryPolicy map[string]interface{} `json:"retryPolicy,omitempty" yaml:"retryPolicy,omitempty"`
	// Reference to a PubSubTopic.
	TopicRef map[string]interface{} `json:"topicRef,omitempty" yaml:"topicRef,omitempty"`
}

// PubSubSubscriptionStatus defines the observed state of PubSubSubscription.
type PubSubSubscriptionStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
