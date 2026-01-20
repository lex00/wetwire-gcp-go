package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// LoggingLogView represents a logging.cnrm.cloud.google.com LoggingLogView resource.
type LoggingLogView struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LoggingLogViewSpec   `json:"spec,omitempty"`
	Status LoggingLogViewStatus `json:"status,omitempty"`
}

// LoggingLogViewSpec defines the desired state of LoggingLogView.
type LoggingLogViewSpec struct {
	// Immutable. The BillingAccount that this resource belongs to. Only one of [billingAccountRef, folderRef, organizationRef, projectRef] may be specified.
	BillingAccountRef map[string]interface{} `json:"billingAccountRef,omitempty" yaml:"billingAccountRef,omitempty"`
	// Immutable.
	BucketRef map[string]interface{} `json:"bucketRef,omitempty" yaml:"bucketRef,omitempty"`
	// Describes this view.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Filter that restricts which log entries in a bucket are visible in this view. Filters are restricted to be a logical AND of ==/!= of any of the following: - originating project/folder/organization/billing account. - resource type - log id For example: SOURCE("projects/myproject") AND resource.type = "gce_instance" AND LOG_ID("stdout")
	Filter string `json:"filter,omitempty" yaml:"filter,omitempty"`
	// Immutable. The Folder that this resource belongs to. Only one of [billingAccountRef, folderRef, organizationRef, projectRef] may be specified.
	FolderRef map[string]interface{} `json:"folderRef,omitempty" yaml:"folderRef,omitempty"`
	// Immutable. The location of the resource. The supported locations are: global, us-central1, us-east1, us-west1, asia-east1, europe-west1.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Immutable. The Organization that this resource belongs to. Only one of [billingAccountRef, folderRef, organizationRef, projectRef] may be specified.
	OrganizationRef map[string]interface{} `json:"organizationRef,omitempty" yaml:"organizationRef,omitempty"`
	// Immutable. The Project that this resource belongs to. Only one of [billingAccountRef, folderRef, organizationRef, projectRef] may be specified.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// LoggingLogViewStatus defines the observed state of LoggingLogView.
type LoggingLogViewStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
