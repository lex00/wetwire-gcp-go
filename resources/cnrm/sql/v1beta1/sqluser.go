package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// SQLUser represents a sql.cnrm.cloud.google.com SQLUser resource.
type SQLUser struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SQLUserSpec   `json:"spec,omitempty"`
	Status SQLUserStatus `json:"status,omitempty"`
}

// SQLUserSpec defines the desired state of SQLUser.
type SQLUserSpec struct {
	// Immutable. The host the user can connect from. This is only supported for MySQL instances. Don't set this field for PostgreSQL instances. Can be an IP address. Changing this forces a new resource to be created.
	Host string `json:"host,omitempty" yaml:"host,omitempty"`
	InstanceRef map[string]interface{} `json:"instanceRef,omitempty" yaml:"instanceRef,omitempty"`
	// The password for the user. Can be updated. For Postgres instances this is a Required field, unless type is set to
	// either CLOUD_IAM_USER or CLOUD_IAM_SERVICE_ACCOUNT.
	Password map[string]interface{} `json:"password,omitempty" yaml:"password,omitempty"`
	PasswordPolicy map[string]interface{} `json:"passwordPolicy,omitempty" yaml:"passwordPolicy,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Immutable. The user type. It determines the method to authenticate the user during login.
	// The default is the database's built-in user type. Flags include "BUILT_IN", "CLOUD_IAM_USER", or "CLOUD_IAM_SERVICE_ACCOUNT".
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}

// SQLUserStatus defines the observed state of SQLUser.
type SQLUserStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
