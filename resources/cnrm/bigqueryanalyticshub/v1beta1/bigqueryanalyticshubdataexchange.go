package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// BigQueryAnalyticsHubDataExchange is the Schema for the BigQueryAnalyticsHubDataExchange API
type BigQueryAnalyticsHubDataExchange struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BigQueryAnalyticsHubDataExchangeSpec   `json:"spec,omitempty"`
	Status BigQueryAnalyticsHubDataExchangeStatus `json:"status,omitempty"`
}

// BigQueryAnalyticsHubDataExchangeSpec defines the desired state of BigQueryAnalyticsHubDataExchange.
type BigQueryAnalyticsHubDataExchangeSpec struct {
	// Optional. Description of the data exchange. The description must not contain Unicode non-characters as well as C0 and C1 control codes except tabs (HT), new lines (LF), carriage returns (CR), and page breaks (FF). Default value is an empty string. Max length: 2000 bytes.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Optional. Type of discovery on the discovery page for all the listings under this exchange. Updating this field also updates (overwrites) the discovery_type field for all the listings under this exchange.
	DiscoveryType string `json:"discoveryType,omitempty" yaml:"discoveryType,omitempty"`
	// Required. Human-readable display name of the data exchange. The display name must contain only Unicode letters, numbers (0-9), underscores (_), dashes (-), spaces ( ), ampersands (&) and must not start or end with spaces. Default value is an empty string. Max length: 63 bytes.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// Optional. Documentation describing the data exchange.
	Documentation string `json:"documentation,omitempty" yaml:"documentation,omitempty"`
	// Immutable. The name of the location this data exchange.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Optional. Email or URL of the primary point of contact of the data exchange. Max Length: 1000 bytes.
	PrimaryContact string `json:"primaryContact,omitempty" yaml:"primaryContact,omitempty"`
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. The BigQueryAnalyticsHubDataExchange name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// BigQueryAnalyticsHubDataExchangeStatus defines the observed state of BigQueryAnalyticsHubDataExchange.
type BigQueryAnalyticsHubDataExchangeStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
