package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// StorageBucketAccessControl represents a storage.cnrm.cloud.google.com StorageBucketAccessControl resource.
type StorageBucketAccessControl struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   StorageBucketAccessControlSpec   `json:"spec,omitempty"`
	Status StorageBucketAccessControlStatus `json:"status,omitempty"`
}

// StorageBucketAccessControlSpec defines the desired state of StorageBucketAccessControl.
type StorageBucketAccessControlSpec struct {
	// Reference to the bucket.
	BucketRef map[string]interface{} `json:"bucketRef,omitempty" yaml:"bucketRef,omitempty"`
	// Immutable. The entity holding the permission, in one of the following forms:
	// user-userId
	// user-email
	// group-groupId
	// group-email
	// domain-domain
	// project-team-projectId
	// allUsers
	// allAuthenticatedUsers
	// Examples:
	// The user liz@example.com would be user-liz@example.com.
	// The group example@googlegroups.com would be
	// group-example@googlegroups.com.
	// To refer to all members of the Google Apps for Business domain
	// example.com, the entity would be domain-example.com.
	Entity string `json:"entity,omitempty" yaml:"entity,omitempty"`
	// The access permission for the entity. Possible values: ["OWNER", "READER", "WRITER"].
	Role string `json:"role,omitempty" yaml:"role,omitempty"`
}

// StorageBucketAccessControlStatus defines the observed state of StorageBucketAccessControl.
type StorageBucketAccessControlStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
