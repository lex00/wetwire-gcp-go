package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CloudIOTDevice represents a cloudiot.cnrm.cloud.google.com CloudIOTDevice resource.
type CloudIOTDevice struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CloudIOTDeviceSpec   `json:"spec,omitempty"`
	Status CloudIOTDeviceStatus `json:"status,omitempty"`
}

// CloudIOTDeviceSpec defines the desired state of CloudIOTDevice.
type CloudIOTDeviceSpec struct {
	// If a device is blocked, connections or requests from this device will fail.
	Blocked bool `json:"blocked,omitempty" yaml:"blocked,omitempty"`
	// The credentials used to authenticate this device.
	Credentials []map[string]interface{} `json:"credentials,omitempty" yaml:"credentials,omitempty"`
	// Gateway-related configuration and state.
	GatewayConfig map[string]interface{} `json:"gatewayConfig,omitempty" yaml:"gatewayConfig,omitempty"`
	// The logging verbosity for device activity. Possible values: ["NONE", "ERROR", "INFO", "DEBUG"].
	LogLevel string `json:"logLevel,omitempty" yaml:"logLevel,omitempty"`
	// The metadata key-value pairs assigned to the device.
	Metadata map[string]string `json:"metadata,omitempty" yaml:"metadata,omitempty"`
	// Immutable. The name of the device registry where this device should be created.
	Registry string `json:"registry,omitempty" yaml:"registry,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// CloudIOTDeviceStatus defines the observed state of CloudIOTDevice.
type CloudIOTDeviceStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
