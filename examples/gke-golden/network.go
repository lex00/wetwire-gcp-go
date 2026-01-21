// Package gke_golden provides a production-ready GKE cluster example.
//
// This example demonstrates best practices for deploying a GKE cluster
// with proper networking, security, and observability configurations.
package gke_golden

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	computev1beta1 "github.com/lex00/wetwire-gcp-go/resources/cnrm/compute/v1beta1"
)

// Helper variables for configuration
var (
	region       = "us-central1"
	networkName  = "gke-golden-vpc"
	gkeSubnet    = "gke-subnet"
	podsCIDR     = "10.4.0.0/14"
	servicesCIDR = "10.8.0.0/20"
)

// VPC is the virtual network for the GKE cluster.
var VPC = computev1beta1.ComputeNetwork{
	TypeMeta: metav1.TypeMeta{
		APIVersion: "compute.cnrm.cloud.google.com/v1beta1",
		Kind:       "ComputeNetwork",
	},
	ObjectMeta: metav1.ObjectMeta{
		Name:      networkName,
		Namespace: "config-connector",
		Labels: map[string]string{
			"environment": "production",
			"managed-by":  "wetwire",
		},
	},
	Spec: computev1beta1.ComputeNetworkSpec{
		AutoCreateSubnetworks: false, // Custom subnet mode for GKE
		RoutingMode:           "REGIONAL",
		Description:           "VPC for production GKE cluster",
	},
}

// GKESubnet is the subnet for GKE nodes.
// Using a /20 CIDR provides ~4000 IPs for nodes.
// Secondary ranges are used for pods and services with VPC-native networking.
var GKESubnet = computev1beta1.ComputeSubnetwork{
	TypeMeta: metav1.TypeMeta{
		APIVersion: "compute.cnrm.cloud.google.com/v1beta1",
		Kind:       "ComputeSubnetwork",
	},
	ObjectMeta: metav1.ObjectMeta{
		Name:      gkeSubnet,
		Namespace: "config-connector",
		Labels: map[string]string{
			"environment": "production",
			"managed-by":  "wetwire",
		},
	},
	Spec: computev1beta1.ComputeSubnetworkSpec{
		Region: region,
		NetworkRef: map[string]interface{}{
			"name": VPC.ObjectMeta.Name,
		},
		IPCidrRange:           "10.0.0.0/20",
		PrivateIPGoogleAccess: true,
		SecondaryIPRange: []map[string]interface{}{
			{
				"rangeName":   "pods",
				"ipCidrRange": podsCIDR,
			},
			{
				"rangeName":   "services",
				"ipCidrRange": servicesCIDR,
			},
		},
		LogConfig: map[string]interface{}{
			"aggregationInterval": "INTERVAL_5_SEC",
			"flowSampling":        0.5,
			"metadata":            "INCLUDE_ALL_METADATA",
		},
	},
}

// CloudRouter is the Cloud Router for NAT gateway.
var CloudRouter = computev1beta1.ComputeRouter{
	TypeMeta: metav1.TypeMeta{
		APIVersion: "compute.cnrm.cloud.google.com/v1beta1",
		Kind:       "ComputeRouter",
	},
	ObjectMeta: metav1.ObjectMeta{
		Name:      "gke-golden-router",
		Namespace: "config-connector",
		Labels: map[string]string{
			"environment": "production",
			"managed-by":  "wetwire",
		},
	},
	Spec: computev1beta1.ComputeRouterSpec{
		Region: region,
		NetworkRef: map[string]interface{}{
			"name": VPC.ObjectMeta.Name,
		},
		Description: "Cloud Router for GKE NAT gateway",
	},
}

// CloudNAT provides egress for private GKE nodes.
var CloudNAT = computev1beta1.ComputeRouterNAT{
	TypeMeta: metav1.TypeMeta{
		APIVersion: "compute.cnrm.cloud.google.com/v1beta1",
		Kind:       "ComputeRouterNAT",
	},
	ObjectMeta: metav1.ObjectMeta{
		Name:      "gke-golden-nat",
		Namespace: "config-connector",
		Labels: map[string]string{
			"environment": "production",
			"managed-by":  "wetwire",
		},
	},
	Spec: computev1beta1.ComputeRouterNATSpec{
		Region: region,
		RouterRef: map[string]interface{}{
			"name": CloudRouter.ObjectMeta.Name,
		},
		NatIPAllocateOption:            "AUTO_ONLY",
		SourceSubnetworkIPRangesToNat: "ALL_SUBNETWORKS_ALL_IP_RANGES",
		MinPortsPerVm:                   64,
		EnableEndpointIndependentMapping: false,
		LogConfig: map[string]interface{}{
			"enable": true,
			"filter": "ERRORS_ONLY",
		},
	},
}

// AllowInternalFirewall allows all internal VPC traffic.
var AllowInternalFirewall = computev1beta1.ComputeFirewall{
	TypeMeta: metav1.TypeMeta{
		APIVersion: "compute.cnrm.cloud.google.com/v1beta1",
		Kind:       "ComputeFirewall",
	},
	ObjectMeta: metav1.ObjectMeta{
		Name:      "gke-golden-allow-internal",
		Namespace: "config-connector",
		Labels: map[string]string{
			"environment": "production",
			"managed-by":  "wetwire",
		},
	},
	Spec: computev1beta1.ComputeFirewallSpec{
		NetworkRef: map[string]interface{}{
			"name": VPC.ObjectMeta.Name,
		},
		Direction:   "INGRESS",
		Priority:    1000,
		Description: "Allow all internal traffic within VPC",
		SourceRanges: []string{
			"10.0.0.0/8",
		},
		Allow: []map[string]interface{}{
			{
				"protocol": "tcp",
			},
			{
				"protocol": "udp",
			},
			{
				"protocol": "icmp",
			},
		},
	},
}

// AllowHealthChecksFirewall allows GCP health check traffic.
var AllowHealthChecksFirewall = computev1beta1.ComputeFirewall{
	TypeMeta: metav1.TypeMeta{
		APIVersion: "compute.cnrm.cloud.google.com/v1beta1",
		Kind:       "ComputeFirewall",
	},
	ObjectMeta: metav1.ObjectMeta{
		Name:      "gke-golden-allow-health-checks",
		Namespace: "config-connector",
		Labels: map[string]string{
			"environment": "production",
			"managed-by":  "wetwire",
		},
	},
	Spec: computev1beta1.ComputeFirewallSpec{
		NetworkRef: map[string]interface{}{
			"name": VPC.ObjectMeta.Name,
		},
		Direction:   "INGRESS",
		Priority:    1000,
		Description: "Allow GCP health check ranges",
		SourceRanges: []string{
			"35.191.0.0/16",  // GCP health check range
			"130.211.0.0/22", // GCP health check range
		},
		Allow: []map[string]interface{}{
			{
				"protocol": "tcp",
			},
		},
	},
}
