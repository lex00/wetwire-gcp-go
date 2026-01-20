package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// FirebaseHostingChannel represents a firebasehosting.cnrm.cloud.google.com FirebaseHostingChannel resource.
type FirebaseHostingChannel struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FirebaseHostingChannelSpec   `json:"spec,omitempty"`
	Status FirebaseHostingChannelStatus `json:"status,omitempty"`
}

// FirebaseHostingChannelSpec defines the desired state of FirebaseHostingChannel.
type FirebaseHostingChannelSpec struct {
	// The time at which the channel will be automatically deleted. If null, the channel
	// will not be automatically deleted. This field is present in the output whether it's
	// set directly or via the 'ttl' field.
	ExpireTime string `json:"expireTime,omitempty" yaml:"expireTime,omitempty"`
	// Immutable. Optional. The channelId of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// The number of previous releases to retain on the channel for rollback or other
	// purposes. Must be a number between 1-100. Defaults to 10 for new channels.
	RetainedReleaseCount int32 `json:"retainedReleaseCount,omitempty" yaml:"retainedReleaseCount,omitempty"`
	// Immutable. Required. The ID of the site in which to create this channel.
	SiteID string `json:"siteId,omitempty" yaml:"siteId,omitempty"`
	// Immutable. Input only. A time-to-live for this channel. Sets 'expire_time' to the provided
	// duration past the time of the request. A duration in seconds with up to nine fractional
	// digits, terminated by 's'. Example: "86400s" (one day).
	Ttl string `json:"ttl,omitempty" yaml:"ttl,omitempty"`
}

// FirebaseHostingChannelStatus defines the observed state of FirebaseHostingChannel.
type FirebaseHostingChannelStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
