package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// StorageBucket represents a storage.cnrm.cloud.google.com StorageBucket resource.
type StorageBucket struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   StorageBucketSpec   `json:"spec,omitempty"`
	Status StorageBucketStatus `json:"status,omitempty"`
}

// StorageBucketSpec defines the desired state of StorageBucket.
type StorageBucketSpec struct {
	// The bucket's autoclass configuration.
	Autoclass map[string]interface{} `json:"autoclass,omitempty" yaml:"autoclass,omitempty"`
	// DEPRECATED. Please use the `uniformBucketLevelAccess` field as this field has been renamed by Google. The `uniformBucketLevelAccess` field will supersede this field.
	// Enables Bucket PolicyOnly access to a bucket.
	BucketPolicyOnly bool `json:"bucketPolicyOnly,omitempty" yaml:"bucketPolicyOnly,omitempty"`
	// The bucket's Cross-Origin Resource Sharing (CORS) configuration.
	Cors []map[string]interface{} `json:"cors,omitempty" yaml:"cors,omitempty"`
	// The bucket's custom location configuration, which specifies the individual regions that comprise a dual-region bucket. If the bucket is designated a single or multi-region, the parameters are empty.
	CustomPlacementConfig map[string]interface{} `json:"customPlacementConfig,omitempty" yaml:"customPlacementConfig,omitempty"`
	// Whether or not to automatically apply an eventBasedHold to new objects added to the bucket.
	DefaultEventBasedHold bool `json:"defaultEventBasedHold,omitempty" yaml:"defaultEventBasedHold,omitempty"`
	// The bucket's encryption configuration.
	Encryption map[string]interface{} `json:"encryption,omitempty" yaml:"encryption,omitempty"`
	// The bucket's Lifecycle Rules configuration.
	LifecycleRule []map[string]interface{} `json:"lifecycleRule,omitempty" yaml:"lifecycleRule,omitempty"`
	// The Google Cloud Storage location.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// The bucket's Access & Storage Logs configuration.
	Logging map[string]interface{} `json:"logging,omitempty" yaml:"logging,omitempty"`
	// Prevents public access to a bucket.
	PublicAccessPrevention string `json:"publicAccessPrevention,omitempty" yaml:"publicAccessPrevention,omitempty"`
	// Enables Requester Pays on a storage bucket.
	RequesterPays bool `json:"requesterPays,omitempty" yaml:"requesterPays,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Configuration of the bucket's data retention policy for how long objects in the bucket should be retained.
	RetentionPolicy map[string]interface{} `json:"retentionPolicy,omitempty" yaml:"retentionPolicy,omitempty"`
	// The bucket's soft delete policy, which defines the period of time that soft-deleted objects will be retained, and cannot be permanently deleted. If it is not provided, by default Google Cloud Storage sets this to default soft delete policy.
	SoftDeletePolicy map[string]interface{} `json:"softDeletePolicy,omitempty" yaml:"softDeletePolicy,omitempty"`
	// The Storage Class of the new bucket. Supported values include: STANDARD, MULTI_REGIONAL, REGIONAL, NEARLINE, COLDLINE, ARCHIVE.
	StorageClass string `json:"storageClass,omitempty" yaml:"storageClass,omitempty"`
	// Enables uniform bucket-level access on a bucket.
	UniformBucketLevelAccess bool `json:"uniformBucketLevelAccess,omitempty" yaml:"uniformBucketLevelAccess,omitempty"`
	// The bucket's Versioning configuration.
	Versioning map[string]interface{} `json:"versioning,omitempty" yaml:"versioning,omitempty"`
	// Configuration if the bucket acts as a website.
	Website map[string]interface{} `json:"website,omitempty" yaml:"website,omitempty"`
}

// StorageBucketStatus defines the observed state of StorageBucket.
type StorageBucketStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
