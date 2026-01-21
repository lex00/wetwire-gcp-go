// Test fixture: Multiple resources in one file
package testdata

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	storagev1beta1 "github.com/lex00/wetwire-gcp-go/resources/cnrm/storage/v1beta1"
)

var Bucket1 = storagev1beta1.StorageBucket{
	ObjectMeta: metav1.ObjectMeta{
		Name: "bucket-one",
	},
	Spec: storagev1beta1.StorageBucketSpec{
		Location: "US",
	},
}

var Bucket2 = storagev1beta1.StorageBucket{
	ObjectMeta: metav1.ObjectMeta{
		Name: "bucket-two",
	},
	Spec: storagev1beta1.StorageBucketSpec{
		Location: "EU",
	},
}

var Bucket3 = storagev1beta1.StorageBucket{
	ObjectMeta: metav1.ObjectMeta{
		Name: "bucket-three",
	},
	Spec: storagev1beta1.StorageBucketSpec{
		Location: "ASIA",
	},
}
