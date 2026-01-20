package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DataFusionInstance represents a datafusion.cnrm.cloud.google.com DataFusionInstance resource.
type DataFusionInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DataFusionInstanceSpec   `json:"spec,omitempty"`
	Status DataFusionInstanceStatus `json:"status,omitempty"`
}

// DataFusionInstanceSpec defines the desired state of DataFusionInstance.
type DataFusionInstanceSpec struct {
	DataprocServiceAccountRef map[string]interface{} `json:"dataprocServiceAccountRef,omitempty" yaml:"dataprocServiceAccountRef,omitempty"`
	// Immutable. A description of this instance.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Immutable. Display name for an instance.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// Option to enable Stackdriver Logging.
	EnableStackdriverLogging bool `json:"enableStackdriverLogging,omitempty" yaml:"enableStackdriverLogging,omitempty"`
	// Option to enable Stackdriver Monitoring.
	EnableStackdriverMonitoring bool `json:"enableStackdriverMonitoring,omitempty" yaml:"enableStackdriverMonitoring,omitempty"`
	// Immutable. The location for the resource
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Immutable. Network configuration options. These are required when a private Data Fusion instance is to be created.
	NetworkConfig map[string]interface{} `json:"networkConfig,omitempty" yaml:"networkConfig,omitempty"`
	// Immutable. Map of additional options used to configure the behavior of Data Fusion instance.
	Options map[string]string `json:"options,omitempty" yaml:"options,omitempty"`
	// Immutable. Specifies whether the Data Fusion instance should be private. If set to true, all Data Fusion nodes will have private IP addresses and will not be able to access the public internet.
	PrivateInstance bool `json:"privateInstance,omitempty" yaml:"privateInstance,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Immutable. Required. Instance type. Possible values: TYPE_UNSPECIFIED, BASIC, ENTERPRISE, DEVELOPER
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
	// Current version of the Data Fusion.
	Version string `json:"version,omitempty" yaml:"version,omitempty"`
	// Immutable. Name of the zone in which the Data Fusion instance will be created. Only DEVELOPER instances use this field.
	Zone string `json:"zone,omitempty" yaml:"zone,omitempty"`
}

// DataFusionInstanceStatus defines the observed state of DataFusionInstance.
type DataFusionInstanceStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
