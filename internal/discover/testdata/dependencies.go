// Test fixture: Resources with dependencies
package testdata

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	containerv1beta1 "github.com/lex00/wetwire-gcp-go/resources/cnrm/container/v1beta1"
)

var Cluster = containerv1beta1.ContainerCluster{
	ObjectMeta: metav1.ObjectMeta{
		Name: "test-cluster",
	},
	Spec: containerv1beta1.ContainerClusterSpec{
		Location: "us-central1",
	},
}

var NodePool = containerv1beta1.ContainerNodePool{
	ObjectMeta: metav1.ObjectMeta{
		Name: "test-nodepool",
	},
	Spec: containerv1beta1.ContainerNodePoolSpec{
		Location: "us-central1",
		ClusterRef: map[string]interface{}{
			"name": Cluster.ObjectMeta.Name,
		},
	},
}
