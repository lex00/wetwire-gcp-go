package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// LoggingLogExclusion represents a logging.cnrm.cloud.google.com LoggingLogExclusion resource.
type LoggingLogExclusion struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LoggingLogExclusionSpec   `json:"spec,omitempty"`
	Status LoggingLogExclusionStatus `json:"status,omitempty"`
}

// LoggingLogExclusionSpec defines the desired state of LoggingLogExclusion.
type LoggingLogExclusionSpec struct {
	// Immutable. The BillingAccount that this resource belongs to. Only one of [projectRef, folderRef, organizationRef, billingAccountRef] may be specified.
	BillingAccountRef map[string]interface{} `json:"billingAccountRef,omitempty" yaml:"billingAccountRef,omitempty"`
	// Optional. A description of this exclusion.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Optional. If set to True, then this exclusion is disabled and it does not exclude any log entries. You can update an exclusion to change the value of this field.
	Disabled bool `json:"disabled,omitempty" yaml:"disabled,omitempty"`
	// Required. An (https://cloud.google.com/logging/docs/view/advanced-queries#sample), you can exclude less than 100% of the matching log entries. For example, the following query matches 99% of low-severity log entries from Google Cloud Storage buckets: `"resource.type=gcs_bucket severity
	Filter string `json:"filter,omitempty" yaml:"filter,omitempty"`
	// Immutable. The Folder that this resource belongs to. Only one of [projectRef, folderRef, organizationRef, billingAccountRef] may be specified.
	FolderRef map[string]interface{} `json:"folderRef,omitempty" yaml:"folderRef,omitempty"`
	// Immutable. The Organization that this resource belongs to. Only one of [projectRef, folderRef, organizationRef, billingAccountRef] may be specified.
	OrganizationRef map[string]interface{} `json:"organizationRef,omitempty" yaml:"organizationRef,omitempty"`
	// Immutable. The Project that this resource belongs to. Only one of [projectRef, folderRef, organizationRef, billingAccountRef] may be specified.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// LoggingLogExclusionStatus defines the observed state of LoggingLogExclusion.
type LoggingLogExclusionStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
