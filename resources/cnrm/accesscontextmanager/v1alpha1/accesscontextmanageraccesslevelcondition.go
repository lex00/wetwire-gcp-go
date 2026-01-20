package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// AccessContextManagerAccessLevelCondition represents a accesscontextmanager.cnrm.cloud.google.com AccessContextManagerAccessLevelCondition resource.
type AccessContextManagerAccessLevelCondition struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AccessContextManagerAccessLevelConditionSpec   `json:"spec,omitempty"`
	Status AccessContextManagerAccessLevelConditionStatus `json:"status,omitempty"`
}

// AccessContextManagerAccessLevelConditionSpec defines the desired state of AccessContextManagerAccessLevelCondition.
type AccessContextManagerAccessLevelConditionSpec struct {
	AccessLevelRef map[string]interface{} `json:"accessLevelRef,omitempty" yaml:"accessLevelRef,omitempty"`
	// Immutable. Device specific restrictions, all restrictions must hold for
	// the Condition to be true. If not specified, all devices are
	// allowed.
	DevicePolicy map[string]interface{} `json:"devicePolicy,omitempty" yaml:"devicePolicy,omitempty"`
	// Immutable. A list of CIDR block IP subnetwork specification. May be IPv4
	// or IPv6.
	// Note that for a CIDR IP address block, the specified IP address
	// portion must be properly truncated (i.e. all the host bits must
	// be zero) or the input is considered malformed. For example,
	// "192.0.2.0/24" is accepted but "192.0.2.1/24" is not. Similarly,
	// for IPv6, "2001:db8::/32" is accepted whereas "2001:db8::1/32"
	// is not. The originating IP of a request must be in one of the
	// listed subnets in order for this Condition to be true.
	// If empty, all IP addresses are allowed.
	IPSubnetworks []string `json:"ipSubnetworks,omitempty" yaml:"ipSubnetworks,omitempty"`
	// Immutable. An allowed list of members (users, service accounts).
	// Using groups is not supported yet.
	// The signed-in user originating the request must be a part of one
	// of the provided members. If not specified, a request may come
	// from any user (logged in/not logged in, not present in any
	// groups, etc.).
	// Formats: 'user:{emailid}', 'serviceAccount:{emailid}'.
	Members []string `json:"members,omitempty" yaml:"members,omitempty"`
	// Immutable. Whether to negate the Condition. If true, the Condition becomes
	// a NAND over its non-empty fields, each field must be false for
	// the Condition overall to be satisfied. Defaults to false.
	Negate bool `json:"negate,omitempty" yaml:"negate,omitempty"`
	// Immutable. The request must originate from one of the provided
	// countries/regions.
	// Format: A valid ISO 3166-1 alpha-2 code.
	Regions []string `json:"regions,omitempty" yaml:"regions,omitempty"`
	// Immutable. A list of other access levels defined in the same Policy,
	// referenced by resource name. Referencing an AccessLevel which
	// does not exist is an error. All access levels listed must be
	// granted for the Condition to be true.
	// Format: accessPolicies/{policy_id}/accessLevels/{short_name}.
	RequiredAccessLevels []string `json:"requiredAccessLevels,omitempty" yaml:"requiredAccessLevels,omitempty"`
	// Immutable. Optional. The accessLevel of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// AccessContextManagerAccessLevelConditionStatus defines the observed state of AccessContextManagerAccessLevelCondition.
type AccessContextManagerAccessLevelConditionStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
