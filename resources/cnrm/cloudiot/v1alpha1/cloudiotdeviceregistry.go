package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CloudIOTDeviceRegistry represents a cloudiot.cnrm.cloud.google.com CloudIOTDeviceRegistry resource.
type CloudIOTDeviceRegistry struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CloudIOTDeviceRegistrySpec   `json:"spec,omitempty"`
	Status CloudIOTDeviceRegistryStatus `json:"status,omitempty"`
}

// CloudIOTDeviceRegistrySpec defines the desired state of CloudIOTDeviceRegistry.
type CloudIOTDeviceRegistrySpec struct {
	// List of public key certificates to authenticate devices.
	Credentials []map[string]interface{} `json:"credentials,omitempty" yaml:"credentials,omitempty"`
	// List of configurations for event notifications, such as PubSub topics
	// to publish device events to.
	EventNotificationConfigs []map[string]interface{} `json:"eventNotificationConfigs,omitempty" yaml:"eventNotificationConfigs,omitempty"`
	// Activate or deactivate HTTP.
	HTTPConfig map[string]interface{} `json:"httpConfig,omitempty" yaml:"httpConfig,omitempty"`
	// The default logging verbosity for activity from devices in this
	// registry. Specifies which events should be written to logs. For
	// example, if the LogLevel is ERROR, only events that terminate in
	// errors will be logged. LogLevel is inclusive; enabling INFO logging
	// will also enable ERROR logging. Default value: "NONE" Possible values: ["NONE", "ERROR", "INFO", "DEBUG"].
	LogLevel string `json:"logLevel,omitempty" yaml:"logLevel,omitempty"`
	// Activate or deactivate MQTT.
	MqttConfig map[string]interface{} `json:"mqttConfig,omitempty" yaml:"mqttConfig,omitempty"`
	// Immutable.
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
	// Immutable. The region in which the created registry should reside.
	// If it is not provided, the provider region is used.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// A PubSub topic to publish device state updates.
	StateNotificationConfig map[string]interface{} `json:"stateNotificationConfig,omitempty" yaml:"stateNotificationConfig,omitempty"`
}

// CloudIOTDeviceRegistryStatus defines the observed state of CloudIOTDeviceRegistry.
type CloudIOTDeviceRegistryStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
