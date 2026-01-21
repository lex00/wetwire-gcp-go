package gke_k8s

// Monitoring documentation for GKE K8s-native example.
//
// This example integrates with wetwire-observability-go for monitoring configuration.
// GKE provides managed Prometheus through Google Cloud Managed Service for Prometheus.
//
// Integration with wetwire-observability-go:
//
//	import (
//		"github.com/lex00/wetwire-observability-go/monitoring/gke"
//		"github.com/lex00/wetwire-observability-go/monitoring/k8s"
//	)
//
//	// Use GKE-specific monitoring configurations
//	var GKEMonitoringConfig = gke.DefaultConfig()
//
//	// Combine with generic K8s monitoring
//	var K8sMonitoringConfig = k8s.DefaultConfig()
//
// GKE Managed Prometheus:
//
// The cluster is configured with managed Prometheus enabled:
//
//	monitoringConfig:
//	  managedPrometheus:
//	    enabled: true
//
// This provides:
// - Automatic metric collection from GKE system components
// - Integration with Cloud Monitoring
// - PromQL query interface via Cloud Monitoring API
//
// Custom Metrics Collection:
//
// For application metrics, deploy PodMonitoring resources:
//
//	apiVersion: monitoring.googleapis.com/v1
//	kind: PodMonitoring
//	metadata:
//	  name: sample-app-monitoring
//	spec:
//	  selector:
//	    matchLabels:
//	      app: sample
//	  endpoints:
//	    - port: metrics
//	      interval: 30s
//
// Alerting:
//
// Use Google Cloud Alerting policies that query managed Prometheus:
//
//	import (
//		monitoringv1 "github.com/lex00/wetwire-gcp-go/resources/cnrm/monitoring/v1beta1"
//	)
//
//	var HighCPUAlert = monitoringv1.MonitoringAlertPolicy{
//		Spec: monitoringv1.MonitoringAlertPolicySpec{
//			DisplayName: "High CPU Usage",
//			Conditions: []monitoringv1.Condition{
//				{
//					DisplayName: "CPU > 80%",
//					// MQL or PromQL query
//				},
//			},
//		},
//	}
//
// Dashboards:
//
// Use Cloud Monitoring dashboards with Prometheus metrics:
//
//	var GKEDashboard = monitoringv1.MonitoringDashboard{
//		Spec: monitoringv1.MonitoringDashboardSpec{
//			DisplayName: "GKE Cluster Overview",
//			// Dashboard JSON configuration
//		},
//	}
