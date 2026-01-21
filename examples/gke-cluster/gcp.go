// Example: Define a GKE cluster with node pool
package gcp

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	containerv1beta1 "github.com/lex00/wetwire-gcp-go/resources/cnrm/container/v1beta1"
)

// ProdCluster is a regional GKE cluster with workload identity
var ProdCluster = containerv1beta1.ContainerCluster{
	ObjectMeta: metav1.ObjectMeta{
		Name:      "prod-cluster",
		Namespace: "default",
	},
	Spec: containerv1beta1.ContainerClusterSpec{
		Location:         "us-central1",
		InitialNodeCount: 1,
		NetworkRef: map[string]interface{}{
			"name": "default",
		},
		IPAllocationPolicy: map[string]interface{}{
			"clusterSecondaryRangeName":  "pods",
			"servicesSecondaryRangeName": "services",
		},
		WorkloadIdentityConfig: map[string]interface{}{
			"workloadPool": "my-project.svc.id.goog",
		},
		ReleaseChannel: map[string]interface{}{
			"channel": "REGULAR",
		},
		AddonsConfig: map[string]interface{}{
			"httpLoadBalancing": map[string]interface{}{
				"disabled": false,
			},
			"horizontalPodAutoscaling": map[string]interface{}{
				"disabled": false,
			},
		},
		LoggingConfig: map[string]interface{}{
			"enableComponents": []string{"SYSTEM_COMPONENTS", "WORKLOADS"},
		},
		MonitoringConfig: map[string]interface{}{
			"enableComponents": []string{"SYSTEM_COMPONENTS"},
			"managedPrometheus": map[string]interface{}{
				"enabled": true,
			},
		},
	},
}

// ProdNodePool is the primary node pool for the cluster
var ProdNodePool = containerv1beta1.ContainerNodePool{
	ObjectMeta: metav1.ObjectMeta{
		Name:      "prod-nodes",
		Namespace: "default",
	},
	Spec: containerv1beta1.ContainerNodePoolSpec{
		Location: "us-central1",
		ClusterRef: map[string]interface{}{
			"name": ProdCluster.ObjectMeta.Name,
		},
		InitialNodeCount: 1,
		Autoscaling: map[string]interface{}{
			"minNodeCount": 1,
			"maxNodeCount": 10,
		},
		NodeConfig: map[string]interface{}{
			"machineType": "e2-standard-4",
			"diskSizeGb":  100,
			"diskType":    "pd-standard",
			"oauthScopes": []string{
				"https://www.googleapis.com/auth/cloud-platform",
			},
			"workloadMetadataConfig": map[string]interface{}{
				"mode": "GKE_METADATA",
			},
			"shieldedInstanceConfig": map[string]interface{}{
				"enableSecureBoot":          true,
				"enableIntegrityMonitoring": true,
			},
		},
		Management: map[string]interface{}{
			"autoRepair":  true,
			"autoUpgrade": true,
		},
		UpgradeSettings: map[string]interface{}{
			"maxSurge":       1,
			"maxUnavailable": 0,
		},
	},
}
