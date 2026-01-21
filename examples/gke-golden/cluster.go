package gke_golden

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	containerv1beta1 "github.com/lex00/wetwire-gcp-go/resources/cnrm/container/v1beta1"
)

// Cluster is the main GKE cluster resource.
var Cluster = containerv1beta1.ContainerCluster{
	TypeMeta: metav1.TypeMeta{
		APIVersion: "container.cnrm.cloud.google.com/v1beta1",
		Kind:       "ContainerCluster",
	},
	ObjectMeta: metav1.ObjectMeta{
		Name:      "gke-golden-cluster",
		Namespace: "config-connector",
		Labels: map[string]string{
			"environment": "production",
			"managed-by":  "wetwire",
		},
	},
	Spec: containerv1beta1.ContainerClusterSpec{
		Location:         region,
		InitialNodeCount: 1,
		// Note: Default node pool will be removed via terraform or manually
		// Config Connector doesn't have RemoveDefaultNodePool field
		NetworkRef: map[string]interface{}{
			"name": VPC.ObjectMeta.Name,
		},
		SubnetworkRef: map[string]interface{}{
			"name": GKESubnet.ObjectMeta.Name,
		},
		// VPC-native networking with alias IPs
		IPAllocationPolicy: map[string]interface{}{
			"clusterSecondaryRangeName":  "pods",
			"servicesSecondaryRangeName": "services",
		},
		// Workload Identity for pod authentication
		WorkloadIdentityConfig: map[string]interface{}{
			"workloadPool": "${PROJECT_ID}.svc.id.goog",
		},
		// Regular release channel for balanced stability/features
		ReleaseChannel: map[string]interface{}{
			"channel": "REGULAR",
		},
		// Enable essential add-ons
		AddonsConfig: map[string]interface{}{
			"httpLoadBalancing": map[string]interface{}{
				"disabled": false,
			},
			"horizontalPodAutoscaling": map[string]interface{}{
				"disabled": false,
			},
			"gcePersistentDiskCsiDriverConfig": map[string]interface{}{
				"enabled": true,
			},
			"gcsFuseCsiDriverConfig": map[string]interface{}{
				"enabled": true,
			},
			"configConnectorConfig": map[string]interface{}{
				"enabled": true,
			},
		},
		// Enable logging for system components and workloads
		LoggingConfig: map[string]interface{}{
			"enableComponents": []string{
				"SYSTEM_COMPONENTS",
				"WORKLOADS",
			},
		},
		// Enable monitoring with managed Prometheus
		MonitoringConfig: map[string]interface{}{
			"enableComponents": []string{
				"SYSTEM_COMPONENTS",
			},
			"managedPrometheus": map[string]interface{}{
				"enabled": true,
			},
		},
		// Private cluster configuration
		PrivateClusterConfig: map[string]interface{}{
			"enablePrivateNodes":    true,
			"enablePrivateEndpoint": false, // Allow public access to control plane
			"masterIpv4CidrBlock":   "172.16.0.0/28",
		},
		// Master authorized networks
		MasterAuthorizedNetworksConfig: map[string]interface{}{
			"cidrBlocks": []map[string]interface{}{
				{
					"displayName": "VPC",
					"cidrBlock":   "10.0.0.0/8",
				},
			},
		},
		// Enable Binary Authorization for container security
		BinaryAuthorization: map[string]interface{}{
			"evaluationMode": "PROJECT_SINGLETON_POLICY_ENFORCE",
		},
		// Enable Shielded Nodes
		NodeConfig: map[string]interface{}{
			"shieldedInstanceConfig": map[string]interface{}{
				"enableSecureBoot":          true,
				"enableIntegrityMonitoring": true,
			},
		},
		// Network policy for pod-to-pod traffic control
		NetworkPolicy: map[string]interface{}{
			"enabled":  true,
			"provider": "CALICO",
		},
		// Maintenance window - weekends 2-6 AM UTC
		MaintenancePolicy: map[string]interface{}{
			"recurringWindow": map[string]interface{}{
				"window": map[string]interface{}{
					"startTime": "2024-01-01T02:00:00Z",
					"endTime":   "2024-01-01T06:00:00Z",
				},
				"recurrence": "FREQ=WEEKLY;BYDAY=SA,SU",
			},
		},
	},
}

// SystemNodePool runs critical system workloads like CoreDNS.
var SystemNodePool = containerv1beta1.ContainerNodePool{
	TypeMeta: metav1.TypeMeta{
		APIVersion: "container.cnrm.cloud.google.com/v1beta1",
		Kind:       "ContainerNodePool",
	},
	ObjectMeta: metav1.ObjectMeta{
		Name:      "gke-golden-system-pool",
		Namespace: "config-connector",
		Labels: map[string]string{
			"environment": "production",
			"managed-by":  "wetwire",
			"pool-type":   "system",
		},
	},
	Spec: containerv1beta1.ContainerNodePoolSpec{
		Location: region,
		ClusterRef: map[string]interface{}{
			"name": Cluster.ObjectMeta.Name,
		},
		InitialNodeCount: 1,
		Autoscaling: map[string]interface{}{
			"minNodeCount": 2,
			"maxNodeCount": 4,
		},
		NodeConfig: map[string]interface{}{
			"machineType": "e2-standard-4",
			"diskSizeGb":  100,
			"diskType":    "pd-ssd",
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
			"labels": map[string]string{
				"pool-type": "system",
			},
			"taints": []map[string]interface{}{
				{
					"key":    "CriticalAddonsOnly",
					"value":  "true",
					"effect": "NO_SCHEDULE",
				},
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
		NodeLocations: []string{
			region + "-a",
			region + "-b",
			region + "-c",
		},
	},
}

// AppNodePool runs general application workloads.
var AppNodePool = containerv1beta1.ContainerNodePool{
	TypeMeta: metav1.TypeMeta{
		APIVersion: "container.cnrm.cloud.google.com/v1beta1",
		Kind:       "ContainerNodePool",
	},
	ObjectMeta: metav1.ObjectMeta{
		Name:      "gke-golden-app-pool",
		Namespace: "config-connector",
		Labels: map[string]string{
			"environment": "production",
			"managed-by":  "wetwire",
			"pool-type":   "application",
		},
	},
	Spec: containerv1beta1.ContainerNodePoolSpec{
		Location: region,
		ClusterRef: map[string]interface{}{
			"name": Cluster.ObjectMeta.Name,
		},
		InitialNodeCount: 1,
		Autoscaling: map[string]interface{}{
			"minNodeCount": 2,
			"maxNodeCount": 10,
		},
		NodeConfig: map[string]interface{}{
			"machineType": "e2-standard-8",
			"diskSizeGb":  200,
			"diskType":    "pd-ssd",
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
			"labels": map[string]string{
				"pool-type": "application",
			},
		},
		Management: map[string]interface{}{
			"autoRepair":  true,
			"autoUpgrade": true,
		},
		UpgradeSettings: map[string]interface{}{
			"maxSurge":       2,
			"maxUnavailable": 1,
		},
		NodeLocations: []string{
			region + "-a",
			region + "-b",
			region + "-c",
		},
	},
}

// SpotNodePool runs fault-tolerant workloads on preemptible/spot instances.
var SpotNodePool = containerv1beta1.ContainerNodePool{
	TypeMeta: metav1.TypeMeta{
		APIVersion: "container.cnrm.cloud.google.com/v1beta1",
		Kind:       "ContainerNodePool",
	},
	ObjectMeta: metav1.ObjectMeta{
		Name:      "gke-golden-spot-pool",
		Namespace: "config-connector",
		Labels: map[string]string{
			"environment": "production",
			"managed-by":  "wetwire",
			"pool-type":   "spot",
		},
	},
	Spec: containerv1beta1.ContainerNodePoolSpec{
		Location: region,
		ClusterRef: map[string]interface{}{
			"name": Cluster.ObjectMeta.Name,
		},
		InitialNodeCount: 0,
		Autoscaling: map[string]interface{}{
			"minNodeCount": 0,
			"maxNodeCount": 20,
		},
		NodeConfig: map[string]interface{}{
			"machineType": "e2-standard-8",
			"diskSizeGb":  100,
			"diskType":    "pd-standard",
			"spot":        true,
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
			"labels": map[string]string{
				"pool-type": "spot",
				"lifecycle": "spot",
			},
			"taints": []map[string]interface{}{
				{
					"key":    "cloud.google.com/gke-spot",
					"value":  "true",
					"effect": "NO_SCHEDULE",
				},
			},
		},
		Management: map[string]interface{}{
			"autoRepair":  true,
			"autoUpgrade": true,
		},
		UpgradeSettings: map[string]interface{}{
			"maxSurge":       5,
			"maxUnavailable": 0,
		},
		NodeLocations: []string{
			region + "-a",
			region + "-b",
			region + "-c",
		},
	},
}
