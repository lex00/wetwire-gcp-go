package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CloudIdentityGroup is the Schema for the CloudIdentityGroup API
type CloudIdentityGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CloudIdentityGroupSpec   `json:"spec,omitempty"`
	Status CloudIdentityGroupStatus `json:"status,omitempty"`
}

// CloudIdentityGroupSpec defines the desired state of CloudIdentityGroup.
type CloudIdentityGroupSpec struct {
	// An extended description to help users determine the purpose of a `Group`. Must not be longer than 4,096 characters.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// The display name of the `Group`.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// Immutable. EntityKey of the Group.
	GroupKey map[string]interface{} `json:"groupKey,omitempty" yaml:"groupKey,omitempty"`
	// Immutable. The initial configuration options for creating a Group. See the [API reference](https://cloud.google.com/identity/docs/reference/rest/v1beta1/groups/create#initialgroupconfig) for possible values. Default value: "EMPTY" Possible values: ["INITIAL_GROUP_CONFIG_UNSPECIFIED", "WITH_INITIAL_OWNER", "EMPTY"].
	InitialGroupConfig string `json:"initialGroupConfig,omitempty" yaml:"initialGroupConfig,omitempty"`
	// One or more label entries that apply to the Group. Currently supported labels contain a key with an empty value. Google Groups are the default type of group and have a label with a key of cloudidentity.googleapis.com/groups.discussion_forum and an empty value. Existing Google Groups can have an additional label with a key of cloudidentity.googleapis.com/groups.security and an empty value added to them. This is an immutable change and the security label cannot be removed once added. Dynamic groups have a label with a key of cloudidentity.googleapis.com/groups.dynamic. Identity-mapped groups for Cloud Search have a label with a key of system/groups/external and an empty value.
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	// Immutable. The resource name of the entity under which this `Group` resides in the Cloud Identity resource hierarchy. Must be of the form `identitysources/{identity_source}` for external [identity-mapped groups](https://support.google.com/a/answer/9039510) or `customers/{customer_id}` for Google Groups. The `customer_id` must begin with "C" (for example, 'C046psxkn'). [Find your customer ID.] (https://support.google.com/cloudidentity/answer/10070793)
	Parent string `json:"parent,omitempty" yaml:"parent,omitempty"`
	// Immutable. Optional. The service-generated name of the resource. Format: groups/{groupID} or {groupID}. Used for acquisition only. Leave unset to create a new resource.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// CloudIdentityGroupStatus defines the observed state of CloudIdentityGroup.
type CloudIdentityGroupStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
