package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// StorageDefaultObjectAccessControl represents a storage.cnrm.cloud.google.com StorageDefaultObjectAccessControl resource.
type StorageDefaultObjectAccessControl struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   StorageDefaultObjectAccessControlSpec   `json:"spec,omitempty"`
	Status StorageDefaultObjectAccessControlStatus `json:"status,omitempty"`
}

// StorageDefaultObjectAccessControlSpec defines the desired state of StorageDefaultObjectAccessControl.
type StorageDefaultObjectAccessControlSpec struct {
	// Reference to the bucket.
	BucketRef map[string]interface{} `json:"bucketRef,omitempty" yaml:"bucketRef,omitempty"`
	// The entity holding the permission, in one of the following forms:
	// * user-{{userId}}
	// * user-{{email}} (such as "user-liz@example.com")
	// * group-{{groupId}}
	// * group-{{email}} (such as "group-example@googlegroups.com")
	// * domain-{{domain}} (such as "domain-example.com")
	// * project-team-{{projectId}}
	// * allUsers
	// * allAuthenticatedUsers.
	Entity string `json:"entity,omitempty" yaml:"entity,omitempty"`
	// The name of the object, if applied to an object.
	Object string `json:"object,omitempty" yaml:"object,omitempty"`
	// The access permission for the entity. Possible values: ["OWNER", "READER"].
	Role string `json:"role,omitempty" yaml:"role,omitempty"`
}

// StorageDefaultObjectAccessControlStatus defines the observed state of StorageDefaultObjectAccessControl.
type StorageDefaultObjectAccessControlStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
