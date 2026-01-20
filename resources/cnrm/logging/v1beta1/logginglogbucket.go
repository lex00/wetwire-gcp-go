package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// LoggingLogBucket represents a logging.cnrm.cloud.google.com LoggingLogBucket resource.
type LoggingLogBucket struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LoggingLogBucketSpec   `json:"spec,omitempty"`
	Status LoggingLogBucketStatus `json:"status,omitempty"`
}

// LoggingLogBucketSpec defines the desired state of LoggingLogBucket.
type LoggingLogBucketSpec struct {
	// Immutable. The BillingAccount that this resource belongs to. Only one of [billingAccountRef, folderRef, organizationRef, projectRef] may be specified.
	BillingAccountRef map[string]interface{} `json:"billingAccountRef,omitempty" yaml:"billingAccountRef,omitempty"`
	// Describes this bucket.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Whether or not Log Analytics is enabled. Logs for buckets with Log Analytics enabled can be queried in the Log Analytics page using SQL queries. Cannot be disabled once enabled.
	EnableAnalytics bool `json:"enableAnalytics,omitempty" yaml:"enableAnalytics,omitempty"`
	// Immutable. The Folder that this resource belongs to. Only one of [billingAccountRef, folderRef, organizationRef, projectRef] may be specified.
	FolderRef map[string]interface{} `json:"folderRef,omitempty" yaml:"folderRef,omitempty"`
	// Immutable. The location of the resource. The supported locations are: global, us-central1, us-east1, us-west1, asia-east1, europe-west1.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Whether the bucket has been locked. The retention period on a locked bucket may not be changed. Locked buckets may only be deleted if they are empty.
	Locked bool `json:"locked,omitempty" yaml:"locked,omitempty"`
	// Immutable. The Organization that this resource belongs to. Only one of [billingAccountRef, folderRef, organizationRef, projectRef] may be specified.
	OrganizationRef map[string]interface{} `json:"organizationRef,omitempty" yaml:"organizationRef,omitempty"`
	// Immutable. The Project that this resource belongs to. Only one of [billingAccountRef, folderRef, organizationRef, projectRef] may be specified.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
	RetentionDays int64 `json:"retentionDays,omitempty" yaml:"retentionDays,omitempty"`
}

// LoggingLogBucketStatus defines the observed state of LoggingLogBucket.
type LoggingLogBucketStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
