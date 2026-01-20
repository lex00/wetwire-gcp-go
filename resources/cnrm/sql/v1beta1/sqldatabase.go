package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// SQLDatabase represents a sql.cnrm.cloud.google.com SQLDatabase resource.
type SQLDatabase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SQLDatabaseSpec   `json:"spec,omitempty"`
	Status SQLDatabaseStatus `json:"status,omitempty"`
}

// SQLDatabaseSpec defines the desired state of SQLDatabase.
type SQLDatabaseSpec struct {
	// The charset value. See MySQL's
	// [Supported Character Sets and Collations](https://dev.mysql.com/doc/refman/5.7/en/charset-charsets.html)
	// and Postgres' [Character Set Support](https://www.postgresql.org/docs/9.6/static/multibyte.html)
	// for more details and supported values. Postgres databases only support
	// a value of 'UTF8' at creation time.
	Charset string `json:"charset,omitempty" yaml:"charset,omitempty"`
	// The collation value. See MySQL's
	// [Supported Character Sets and Collations](https://dev.mysql.com/doc/refman/5.7/en/charset-charsets.html)
	// and Postgres' [Collation Support](https://www.postgresql.org/docs/9.6/static/collation.html)
	// for more details and supported values. Postgres databases only support
	// a value of 'en_US.UTF8' at creation time.
	Collation string `json:"collation,omitempty" yaml:"collation,omitempty"`
	// The deletion policy for the database. Setting ABANDON allows the resource
	// to be abandoned rather than deleted. This is useful for Postgres, where databases cannot be
	// deleted from the API if there are users other than cloudsqlsuperuser with access. Possible
	// values are: "ABANDON", "DELETE". Defaults to "DELETE".
	DeletionPolicy string `json:"deletionPolicy,omitempty" yaml:"deletionPolicy,omitempty"`
	// The Cloud SQL instance.
	InstanceRef map[string]interface{} `json:"instanceRef,omitempty" yaml:"instanceRef,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// SQLDatabaseStatus defines the observed state of SQLDatabase.
type SQLDatabaseStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
