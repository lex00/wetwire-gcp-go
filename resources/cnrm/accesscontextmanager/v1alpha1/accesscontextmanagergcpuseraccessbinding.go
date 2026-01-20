package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// AccessContextManagerGCPUserAccessBinding represents a accesscontextmanager.cnrm.cloud.google.com AccessContextManagerGCPUserAccessBinding resource.
type AccessContextManagerGCPUserAccessBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AccessContextManagerGCPUserAccessBindingSpec   `json:"spec,omitempty"`
	Status AccessContextManagerGCPUserAccessBindingStatus `json:"status,omitempty"`
}

// AccessContextManagerGCPUserAccessBindingSpec defines the desired state of AccessContextManagerGCPUserAccessBinding.
type AccessContextManagerGCPUserAccessBindingSpec struct {
	// Required. Access level that a user must have to be granted access. Only one access level is supported, not multiple. This repeated field must have exactly one element. Example: "accessPolicies/9522/accessLevels/device_trusted".
	AccessLevels []string `json:"accessLevels,omitempty" yaml:"accessLevels,omitempty"`
	// Immutable. Required. Immutable. Google Group id whose members are subject to this binding's restrictions. See "id" in the G Suite Directory API's Groups resource. If a group's email address/alias is changed, this resource will continue to point at the changed group. This field does not accept group email addresses or aliases. Example: "01d520gv4vjcrht".
	GroupKey string `json:"groupKey,omitempty" yaml:"groupKey,omitempty"`
	// The organization that this resource belongs to.
	OrganizationRef map[string]interface{} `json:"organizationRef,omitempty" yaml:"organizationRef,omitempty"`
	// Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// AccessContextManagerGCPUserAccessBindingStatus defines the observed state of AccessContextManagerGCPUserAccessBinding.
type AccessContextManagerGCPUserAccessBindingStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
