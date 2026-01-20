package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// BigQueryDatasetAccess represents a bigquery.cnrm.cloud.google.com BigQueryDatasetAccess resource.
type BigQueryDatasetAccess struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BigQueryDatasetAccessSpec   `json:"spec,omitempty"`
	Status BigQueryDatasetAccessStatus `json:"status,omitempty"`
}

// BigQueryDatasetAccessSpec defines the desired state of BigQueryDatasetAccess.
type BigQueryDatasetAccessSpec struct {
	// Immutable. Grants all resources of particular types in a particular dataset read access to the current dataset.
	Dataset map[string]interface{} `json:"dataset,omitempty" yaml:"dataset,omitempty"`
	// Immutable. A unique ID for this dataset, without the project name. The ID
	// must contain only letters (a-z, A-Z), numbers (0-9), or
	// underscores (_). The maximum length is 1,024 characters.
	DatasetID string `json:"datasetId,omitempty" yaml:"datasetId,omitempty"`
	// Immutable. A domain to grant access to. Any users signed in with the
	// domain specified will be granted the specified access.
	Domain string `json:"domain,omitempty" yaml:"domain,omitempty"`
	// Immutable. An email address of a Google Group to grant access to.
	GroupByEmail string `json:"groupByEmail,omitempty" yaml:"groupByEmail,omitempty"`
	// Immutable. Some other type of member that appears in the IAM Policy but isn't a user,
	// group, domain, or special group. For example: 'allUsers'.
	IamMember string `json:"iamMember,omitempty" yaml:"iamMember,omitempty"`
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The routine of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Immutable. Describes the rights granted to the user specified by the other
	// member of the access object. Basic, predefined, and custom roles are
	// supported. Predefined roles that have equivalent basic roles are
	// swapped by the API to their basic counterparts, and will show a diff
	// post-create. See
	// [official docs](https://cloud.google.com/bigquery/docs/access-control).
	Role string `json:"role,omitempty" yaml:"role,omitempty"`
	// Immutable. A special group to grant access to. Possible values include:
	// * 'projectOwners': Owners of the enclosing project.
	// * 'projectReaders': Readers of the enclosing project.
	// * 'projectWriters': Writers of the enclosing project.
	// * 'allAuthenticatedUsers': All authenticated BigQuery users.
	SpecialGroup string `json:"specialGroup,omitempty" yaml:"specialGroup,omitempty"`
	// Immutable. An email address of a user to grant access to. For example:
	// fred@example.com.
	UserByEmail string `json:"userByEmail,omitempty" yaml:"userByEmail,omitempty"`
	// Immutable. A view from a different dataset to grant access to. Queries
	// executed against that view will have read access to tables in
	// this dataset. The role field is not required when this field is
	// set. If that view is updated by any user, access to the view
	// needs to be granted again via an update operation.
	View map[string]interface{} `json:"view,omitempty" yaml:"view,omitempty"`
}

// BigQueryDatasetAccessStatus defines the observed state of BigQueryDatasetAccess.
type BigQueryDatasetAccessStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
