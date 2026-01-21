package gke_golden

// This file demonstrates how to integrate wetwire-observability-go monitoring
// modules with the GKE golden example.
//
// The monitoring configuration imports from wetwire-observability-go/monitoring/gke
// which provides GKE-specific scrape configs, alerts, recording rules, and dashboards.

/*
Usage with wetwire-observability-go:

import (
	"github.com/lex00/wetwire-observability-go/monitoring/gke"
	"github.com/lex00/wetwire-observability-go/prometheus"
)

// PrometheusConfig creates a complete Prometheus configuration for GKE monitoring.
var PrometheusConfig = prometheus.Config{
	Global: prometheus.GlobalConfig{
		ScrapeInterval:     15 * prometheus.Second,
		EvaluationInterval: 15 * prometheus.Second,
		ExternalLabels: map[string]string{
			"cluster": "gke-golden-cluster",
			"region":  "us-central1",
		},
	},
	// Use all GKE-specific scrape configs
	ScrapeConfigs: gke.AllScrapeConfigs(),
	// Use GKE remote write helper for managed Prometheus
	RemoteWrite: []*prometheus.RemoteWriteConfig{
		gke.GCPManagedPrometheusRemoteWrite("my-project"),
	},
}

// AlertingRules combines base K8s alerts with GKE-specific alerts.
var AlertingRules = gke.AllAlerts()

// RecordingRules provides pre-computed metrics for efficient querying.
var RecordingRules = gke.AllRecordingRules()

// ClusterDashboard provides a comprehensive GKE cluster overview.
var ClusterDashboard = gke.GKEClusterDashboard

// GCLBDashboard provides Google Cloud Load Balancer monitoring.
var GCLBDashboard = gke.GKEGCLBDashboard

// The monitoring/gke package provides:
//
// Scrape Configs:
// - Base K8s scrapes (nodes, pods, services, kube-state-metrics, cadvisor)
// - GCP Managed Prometheus endpoint
// - Stackdriver exporter
// - GKE Ingress Controller metrics
// - Config Connector metrics
// - Workload Identity metrics
//
// Alerts:
// - Base K8s alerts (node, pod, cluster health)
// - GKE node pool scaling alerts
// - GCLB health and performance alerts
// - Config Connector reconciliation alerts
// - Workload Identity credential errors
// - GKE Autopilot-specific alerts (if using Autopilot)
// - Persistent Disk CSI alerts
//
// PromQL Expressions:
// - Node pool utilization metrics
// - GCLB request rates and latencies
// - Config Connector operation counts
// - Workload Identity token refresh rates
//
// Dashboards:
// - GKE Cluster Overview
// - GCLB Performance
// - GKE Autopilot (if applicable)
// - Config Connector Operations
*/
