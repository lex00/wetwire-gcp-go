// Example: Define a GCS bucket with lifecycle rules
package gcp

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	storagev1beta1 "github.com/lex00/wetwire-gcp-go/resources/cnrm/storage/v1beta1"
)

// DataBucket is a GCS bucket for storing application data
var DataBucket = storagev1beta1.StorageBucket{
	ObjectMeta: metav1.ObjectMeta{
		Name:      "my-app-data",
		Namespace: "default",
	},
	Spec: storagev1beta1.StorageBucketSpec{
		Location:                 "US",
		StorageClass:             "STANDARD",
		UniformBucketLevelAccess: true,
		Versioning: map[string]interface{}{
			"enabled": true,
		},
		LifecycleRule: []map[string]interface{}{
			{
				"action": map[string]interface{}{
					"type":         "SetStorageClass",
					"storageClass": "NEARLINE",
				},
				"condition": map[string]interface{}{
					"age": 30,
				},
			},
			{
				"action": map[string]interface{}{
					"type": "Delete",
				},
				"condition": map[string]interface{}{
					"age": 365,
				},
			},
		},
	},
}

// BackupBucket is a GCS bucket for backups with COLDLINE storage
var BackupBucket = storagev1beta1.StorageBucket{
	ObjectMeta: metav1.ObjectMeta{
		Name:      "my-app-backups",
		Namespace: "default",
	},
	Spec: storagev1beta1.StorageBucketSpec{
		Location:                 "US",
		StorageClass:             "COLDLINE",
		UniformBucketLevelAccess: true,
	},
}
