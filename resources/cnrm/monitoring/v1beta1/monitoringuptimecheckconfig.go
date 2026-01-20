package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MonitoringUptimeCheckConfig represents a monitoring.cnrm.cloud.google.com MonitoringUptimeCheckConfig resource.
type MonitoringUptimeCheckConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MonitoringUptimeCheckConfigSpec   `json:"spec,omitempty"`
	Status MonitoringUptimeCheckConfigStatus `json:"status,omitempty"`
}

// MonitoringUptimeCheckConfigSpec defines the desired state of MonitoringUptimeCheckConfig.
type MonitoringUptimeCheckConfigSpec struct {
	// The content that is expected to appear in the data returned by the target server against which the check is run.  Currently, only the first entry in the `content_matchers` list is supported, and additional entries will be ignored. This field is optional and should only be specified if a content match is required as part of the/ Uptime check.
	ContentMatchers []map[string]interface{} `json:"contentMatchers,omitempty" yaml:"contentMatchers,omitempty"`
	// A human-friendly name for the Uptime check configuration. The display name should be unique within a Stackdriver Workspace in order to make it easier to identify; however, uniqueness is not enforced. Required.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// Contains information needed to make an HTTP or HTTPS check.
	HTTPCheck map[string]interface{} `json:"httpCheck,omitempty" yaml:"httpCheck,omitempty"`
	// Immutable. The [monitored resource](https://cloud.google.com/monitoring/api/resources) associated with the configuration. The following monitored resource types are supported for Uptime checks:   `uptime_url`,   `gce_instance`,   `gae_app`,   `aws_ec2_instance`,   `aws_elb_load_balancer`
	MonitoredResource map[string]interface{} `json:"monitoredResource,omitempty" yaml:"monitoredResource,omitempty"`
	// How often, in seconds, the Uptime check is performed. Currently, the only supported values are `60s` (1 minute), `300s` (5 minutes), `600s` (10 minutes), and `900s` (15 minutes). Optional, defaults to `60s`.
	Period string `json:"period,omitempty" yaml:"period,omitempty"`
	// Immutable. The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. The group resource associated with the configuration.
	ResourceGroup map[string]interface{} `json:"resourceGroup,omitempty" yaml:"resourceGroup,omitempty"`
	// Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// The list of regions from which the check will be run. Some regions contain one location, and others contain more than one. If this field is specified, enough regions must be provided to include a minimum of 3 locations.  Not specifying this field will result in Uptime checks running from all available regions.
	SelectedRegions []string `json:"selectedRegions,omitempty" yaml:"selectedRegions,omitempty"`
	// Contains information needed to make a TCP check.
	TCPCheck map[string]interface{} `json:"tcpCheck,omitempty" yaml:"tcpCheck,omitempty"`
	// The maximum amount of time to wait for the request to complete (must be between 1 and 60 seconds). Required.
	Timeout string `json:"timeout,omitempty" yaml:"timeout,omitempty"`
}

// MonitoringUptimeCheckConfigStatus defines the observed state of MonitoringUptimeCheckConfig.
type MonitoringUptimeCheckConfigStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
