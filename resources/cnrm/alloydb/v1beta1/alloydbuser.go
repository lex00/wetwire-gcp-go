package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// AlloyDBUser represents a alloydb.cnrm.cloud.google.com AlloyDBUser resource.
type AlloyDBUser struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AlloyDBUserSpec   `json:"spec,omitempty"`
	Status AlloyDBUserStatus `json:"status,omitempty"`
}

// AlloyDBUserSpec defines the desired state of AlloyDBUser.
type AlloyDBUserSpec struct {
	ClusterRef map[string]interface{} `json:"clusterRef,omitempty" yaml:"clusterRef,omitempty"`
	// List of database roles this database user has.
	DatabaseRoles []string `json:"databaseRoles,omitempty" yaml:"databaseRoles,omitempty"`
	// Password for this database user.
	Password map[string]interface{} `json:"password,omitempty" yaml:"password,omitempty"`
	// Immutable. Optional. The userId of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Immutable. The type of this user. Possible values: ["ALLOYDB_BUILT_IN", "ALLOYDB_IAM_USER"].
	UserType string `json:"userType,omitempty" yaml:"userType,omitempty"`
}

// AlloyDBUserStatus defines the observed state of AlloyDBUser.
type AlloyDBUserStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
