package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ContainerCluster represents a container.cnrm.cloud.google.com ContainerCluster resource.
type ContainerCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ContainerClusterSpec   `json:"spec,omitempty"`
	Status ContainerClusterStatus `json:"status,omitempty"`
}

// ContainerClusterSpec defines the desired state of ContainerCluster.
type ContainerClusterSpec struct {
	// The configuration for addons supported by GKE.
	AddonsConfig map[string]interface{} `json:"addonsConfig,omitempty" yaml:"addonsConfig,omitempty"`
	// Enable NET_ADMIN for this cluster.
	AllowNetAdmin bool `json:"allowNetAdmin,omitempty" yaml:"allowNetAdmin,omitempty"`
	// Configuration for the Google Groups for GKE feature.
	AuthenticatorGroupsConfig map[string]interface{} `json:"authenticatorGroupsConfig,omitempty" yaml:"authenticatorGroupsConfig,omitempty"`
	// Configuration options for the Binary Authorization feature.
	BinaryAuthorization map[string]interface{} `json:"binaryAuthorization,omitempty" yaml:"binaryAuthorization,omitempty"`
	// Per-cluster configuration of Node Auto-Provisioning with Cluster Autoscaler to automatically adjust the size of the cluster and create/delete node pools based on the current needs of the cluster's workload. See the guide to using Node Auto-Provisioning for more details.
	ClusterAutoscaling map[string]interface{} `json:"clusterAutoscaling,omitempty" yaml:"clusterAutoscaling,omitempty"`
	// Immutable. The IP address range of the Kubernetes pods in this cluster in CIDR notation (e.g. 10.96.0.0/14). Leave blank to have one automatically chosen or specify a /14 block in 10.0.0.0/8. This field will only work for routes-based clusters, where ip_allocation_policy is not defined.
	ClusterIpv4Cidr string `json:"clusterIpv4Cidr,omitempty" yaml:"clusterIpv4Cidr,omitempty"`
	// Telemetry integration for the cluster.
	ClusterTelemetry map[string]interface{} `json:"clusterTelemetry,omitempty" yaml:"clusterTelemetry,omitempty"`
	// Immutable. Configuration for the confidential nodes feature, which makes nodes run on confidential VMs. Warning: This configuration can't be changed (or added/removed) after cluster creation without deleting and recreating the entire cluster.
	ConfidentialNodes map[string]interface{} `json:"confidentialNodes,omitempty" yaml:"confidentialNodes,omitempty"`
	// Configuration for all of the cluster's control plane endpoints. Currently supports only DNS endpoint configuration and disable IP endpoint. Other IP endpoint configurations are available in private_cluster_config.
	ControlPlaneEndpointsConfig map[string]interface{} `json:"controlPlaneEndpointsConfig,omitempty" yaml:"controlPlaneEndpointsConfig,omitempty"`
	// Cost management configuration for the cluster.
	CostManagementConfig map[string]interface{} `json:"costManagementConfig,omitempty" yaml:"costManagementConfig,omitempty"`
	// Application-layer Secrets Encryption settings. The object format is {state = string, key_name = string}. Valid values of state are: "ENCRYPTED"; "DECRYPTED". key_name is the name of a CloudKMS key.
	DatabaseEncryption map[string]interface{} `json:"databaseEncryption,omitempty" yaml:"databaseEncryption,omitempty"`
	// Immutable. The desired datapath provider for this cluster. By default, uses the IPTables-based kube-proxy implementation.
	DatapathProvider string `json:"datapathProvider,omitempty" yaml:"datapathProvider,omitempty"`
	// Immutable. The default maximum number of pods per node in this cluster. This doesn't work on "routes-based" clusters, clusters that don't have IP Aliasing enabled.
	DefaultMaxPodsPerNode int32 `json:"defaultMaxPodsPerNode,omitempty" yaml:"defaultMaxPodsPerNode,omitempty"`
	// Whether the cluster disables default in-node sNAT rules. In-node sNAT rules will be disabled when defaultSnatStatus is disabled.
	DefaultSnatStatus map[string]interface{} `json:"defaultSnatStatus,omitempty" yaml:"defaultSnatStatus,omitempty"`
	// Immutable.  Description of the cluster.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Immutable. Configuration for Cloud DNS for Kubernetes Engine.
	DNSConfig map[string]interface{} `json:"dnsConfig,omitempty" yaml:"dnsConfig,omitempty"`
	// Immutable. Enable Autopilot for this cluster.
	EnableAutopilot bool `json:"enableAutopilot,omitempty" yaml:"enableAutopilot,omitempty"`
	// DEPRECATED. Deprecated in favor of binary_authorization. Enable Binary Authorization for this cluster. If enabled, all container images will be validated by Google Binary Authorization.
	EnableBinaryAuthorization bool `json:"enableBinaryAuthorization,omitempty" yaml:"enableBinaryAuthorization,omitempty"`
	// Whether Cilium cluster-wide network policy is enabled on this cluster.
	EnableCiliumClusterwideNetworkPolicy bool `json:"enableCiliumClusterwideNetworkPolicy,omitempty" yaml:"enableCiliumClusterwideNetworkPolicy,omitempty"`
	// Whether FQDN Network Policy is enabled on this cluster.
	EnableFqdnNetworkPolicy bool `json:"enableFqdnNetworkPolicy,omitempty" yaml:"enableFqdnNetworkPolicy,omitempty"`
	// Whether Intra-node visibility is enabled for this cluster. This makes same node pod to pod traffic visible for VPC network.
	EnableIntranodeVisibility bool `json:"enableIntranodeVisibility,omitempty" yaml:"enableIntranodeVisibility,omitempty"`
	// Configuration for Kubernetes Beta APIs.
	EnableK8sBetaApis map[string]interface{} `json:"enableK8sBetaApis,omitempty" yaml:"enableK8sBetaApis,omitempty"`
	// Immutable. Whether to enable Kubernetes Alpha features for this cluster. Note that when this option is enabled, the cluster cannot be upgraded and will be automatically deleted after 30 days.
	EnableKubernetesAlpha bool `json:"enableKubernetesAlpha,omitempty" yaml:"enableKubernetesAlpha,omitempty"`
	// Whether L4ILB Subsetting is enabled for this cluster.
	EnableL4IlbSubsetting bool `json:"enableL4IlbSubsetting,omitempty" yaml:"enableL4IlbSubsetting,omitempty"`
	// Whether the ABAC authorizer is enabled for this cluster. When enabled, identities in the system, including service accounts, nodes, and controllers, will have statically granted permissions beyond those provided by the RBAC configuration or IAM. Defaults to false.
	EnableLegacyAbac bool `json:"enableLegacyAbac,omitempty" yaml:"enableLegacyAbac,omitempty"`
	// Immutable. Whether multi-networking is enabled for this cluster.
	EnableMultiNetworking bool `json:"enableMultiNetworking,omitempty" yaml:"enableMultiNetworking,omitempty"`
	// Enable Shielded Nodes features on all nodes in this cluster. Defaults to true.
	EnableShieldedNodes bool `json:"enableShieldedNodes,omitempty" yaml:"enableShieldedNodes,omitempty"`
	// Immutable. Whether to enable Cloud TPU resources in this cluster.
	EnableTpu bool `json:"enableTpu,omitempty" yaml:"enableTpu,omitempty"`
	// Configuration for GKE Gateway API controller.
	GatewayAPIConfig map[string]interface{} `json:"gatewayApiConfig,omitempty" yaml:"gatewayApiConfig,omitempty"`
	// Configuration for Identity Service which allows customers to use external identity providers with the K8S API.
	IdentityServiceConfig map[string]interface{} `json:"identityServiceConfig,omitempty" yaml:"identityServiceConfig,omitempty"`
	// Immutable. The number of nodes to create in this cluster's default node pool. In regional or multi-zonal clusters, this is the number of nodes per zone. Must be set if node_pool is not set. If you're using google_container_node_pool objects with no default node pool, you'll need to set this to a value of at least 1, alongside setting remove_default_node_pool to true.
	InitialNodeCount int32 `json:"initialNodeCount,omitempty" yaml:"initialNodeCount,omitempty"`
	// Immutable. Configuration of cluster IP allocation for VPC-native clusters. Adding this block enables IP aliasing, making the cluster VPC-native instead of routes-based.
	IPAllocationPolicy map[string]interface{} `json:"ipAllocationPolicy,omitempty" yaml:"ipAllocationPolicy,omitempty"`
	// Immutable. The location (region or zone) in which the cluster master will be created, as well as the default node location. If you specify a zone (such as us-central1-a), the cluster will be a zonal cluster with a single cluster master. If you specify a region (such as us-west1), the cluster will be a regional cluster with multiple masters spread across zones in the region, and with default node locations in those zones as well.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Logging configuration for the cluster.
	LoggingConfig map[string]interface{} `json:"loggingConfig,omitempty" yaml:"loggingConfig,omitempty"`
	// The logging service that the cluster should write logs to. Available options include logging.googleapis.com(Legacy Stackdriver), logging.googleapis.com/kubernetes(Stackdriver Kubernetes Engine Logging), and none. Defaults to logging.googleapis.com/kubernetes.
	LoggingService string `json:"loggingService,omitempty" yaml:"loggingService,omitempty"`
	// The maintenance policy to use for the cluster.
	MaintenancePolicy map[string]interface{} `json:"maintenancePolicy,omitempty" yaml:"maintenancePolicy,omitempty"`
	// DEPRECATED. Basic authentication was removed for GKE cluster versions >= 1.19. The authentication information for accessing the Kubernetes master. Some values in this block are only returned by the API if your service account has permission to get credentials for your GKE cluster. If you see an unexpected diff unsetting your client cert, ensure you have the container.clusters.getCredentials permission.
	MasterAuth map[string]interface{} `json:"masterAuth,omitempty" yaml:"masterAuth,omitempty"`
	// The desired configuration options for master authorized networks. Omit the nested cidr_blocks attribute to disallow external access (except the cluster node IPs, which GKE automatically whitelists).
	MasterAuthorizedNetworksConfig map[string]interface{} `json:"masterAuthorizedNetworksConfig,omitempty" yaml:"masterAuthorizedNetworksConfig,omitempty"`
	// If set, and enable_certificates=true, the GKE Workload Identity Certificates controller and node agent will be deployed in the cluster.
	MeshCertificates map[string]interface{} `json:"meshCertificates,omitempty" yaml:"meshCertificates,omitempty"`
	// The minimum version of the master. GKE will auto-update the master to new versions, so this does not guarantee the current master version--use the read-only master_version field to obtain that. If unset, the cluster's version will be set by GKE to the version of the most recent official release (which is not necessarily the latest version).
	MinMasterVersion string `json:"minMasterVersion,omitempty" yaml:"minMasterVersion,omitempty"`
	// Monitoring configuration for the cluster.
	MonitoringConfig map[string]interface{} `json:"monitoringConfig,omitempty" yaml:"monitoringConfig,omitempty"`
	// The monitoring service that the cluster should write metrics to. Automatically send metrics from pods in the cluster to the Google Cloud Monitoring API. VM metrics will be collected by Google Compute Engine regardless of this setting Available options include monitoring.googleapis.com(Legacy Stackdriver), monitoring.googleapis.com/kubernetes(Stackdriver Kubernetes Engine Monitoring), and none. Defaults to monitoring.googleapis.com/kubernetes.
	MonitoringService string `json:"monitoringService,omitempty" yaml:"monitoringService,omitempty"`
	// Configuration options for the NetworkPolicy feature.
	NetworkPolicy map[string]interface{} `json:"networkPolicy,omitempty" yaml:"networkPolicy,omitempty"`
	NetworkRef map[string]interface{} `json:"networkRef,omitempty" yaml:"networkRef,omitempty"`
	// Immutable. Determines whether alias IPs or routes will be used for pod IPs in the cluster.
	NetworkingMode string `json:"networkingMode,omitempty" yaml:"networkingMode,omitempty"`
	// Immutable. The configuration of the nodepool.
	NodeConfig map[string]interface{} `json:"nodeConfig,omitempty" yaml:"nodeConfig,omitempty"`
	// The list of zones in which the cluster's nodes are located. Nodes must be in the region of their regional cluster or in the same region as their cluster's zone for zonal clusters. If this is specified for a zonal cluster, omit the cluster's zone.
	NodeLocations []string `json:"nodeLocations,omitempty" yaml:"nodeLocations,omitempty"`
	// Node pool configs that apply to all auto-provisioned node pools in autopilot clusters and node auto-provisioning enabled clusters.
	NodePoolAutoConfig map[string]interface{} `json:"nodePoolAutoConfig,omitempty" yaml:"nodePoolAutoConfig,omitempty"`
	// The default nodel pool settings for the entire cluster.
	NodePoolDefaults map[string]interface{} `json:"nodePoolDefaults,omitempty" yaml:"nodePoolDefaults,omitempty"`
	NodeVersion string `json:"nodeVersion,omitempty" yaml:"nodeVersion,omitempty"`
	// The notification config for sending cluster upgrade notifications.
	NotificationConfig map[string]interface{} `json:"notificationConfig,omitempty" yaml:"notificationConfig,omitempty"`
	// Configuration for the PodSecurityPolicy feature.
	PodSecurityPolicyConfig map[string]interface{} `json:"podSecurityPolicyConfig,omitempty" yaml:"podSecurityPolicyConfig,omitempty"`
	// Configuration for private clusters, clusters with private nodes.
	PrivateClusterConfig map[string]interface{} `json:"privateClusterConfig,omitempty" yaml:"privateClusterConfig,omitempty"`
	// The desired state of IPv6 connectivity to Google Services. By default, no private IPv6 access to or from Google Services (all access will be via IPv4).
	PrivateIpv6GoogleAccess string `json:"privateIpv6GoogleAccess,omitempty" yaml:"privateIpv6GoogleAccess,omitempty"`
	// Enable/Disable Protect API features for the cluster.
	ProtectConfig map[string]interface{} `json:"protectConfig,omitempty" yaml:"protectConfig,omitempty"`
	// Configuration options for the Release channel feature, which provide more control over automatic upgrades of your GKE clusters. Note that removing this field from your config will not unenroll it. Instead, use the "UNSPECIFIED" channel.
	ReleaseChannel map[string]interface{} `json:"releaseChannel,omitempty" yaml:"releaseChannel,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Configuration for the ResourceUsageExportConfig feature.
	ResourceUsageExportConfig map[string]interface{} `json:"resourceUsageExportConfig,omitempty" yaml:"resourceUsageExportConfig,omitempty"`
	// Defines the config needed to enable/disable features for the Security Posture API.
	SecurityPostureConfig map[string]interface{} `json:"securityPostureConfig,omitempty" yaml:"securityPostureConfig,omitempty"`
	// If set, and enabled=true, services with external ips field will not be blocked.
	ServiceExternalIpsConfig map[string]interface{} `json:"serviceExternalIpsConfig,omitempty" yaml:"serviceExternalIpsConfig,omitempty"`
	SubnetworkRef map[string]interface{} `json:"subnetworkRef,omitempty" yaml:"subnetworkRef,omitempty"`
	// Vertical Pod Autoscaling automatically adjusts the resources of pods controlled by it.
	VerticalPodAutoscaling map[string]interface{} `json:"verticalPodAutoscaling,omitempty" yaml:"verticalPodAutoscaling,omitempty"`
	// Configuration for the use of Kubernetes Service Accounts in GCP IAM policies.
	WorkloadIdentityConfig map[string]interface{} `json:"workloadIdentityConfig,omitempty" yaml:"workloadIdentityConfig,omitempty"`
}

// ContainerClusterStatus defines the observed state of ContainerCluster.
type ContainerClusterStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
