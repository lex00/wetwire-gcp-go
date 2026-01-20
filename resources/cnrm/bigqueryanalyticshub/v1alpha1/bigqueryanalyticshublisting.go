package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// BigQueryAnalyticsHubListing is the Schema for the BigQueryAnalyticsHubListing API
type BigQueryAnalyticsHubListing struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BigQueryAnalyticsHubListingSpec   `json:"spec,omitempty"`
	Status BigQueryAnalyticsHubListingStatus `json:"status,omitempty"`
}

// BigQueryAnalyticsHubListingSpec defines the desired state of BigQueryAnalyticsHubListing.
type BigQueryAnalyticsHubListingSpec struct {
	// Optional. Categories of the listing. Up to two categories are allowed.
	Categories []string `json:"categories,omitempty" yaml:"categories,omitempty"`
	// BigQueryAnalyticsHubDataExchangeRef defines the resource reference to BigQueryAnalyticsHubDataExchange, which "External" field holds the GCP identifier for the KRM object.
	DataExchangeRef map[string]interface{} `json:"dataExchangeRef,omitempty" yaml:"dataExchangeRef,omitempty"`
	// Optional. Details of the data provider who owns the source data.
	DataProvider map[string]interface{} `json:"dataProvider,omitempty" yaml:"dataProvider,omitempty"`
	// Optional. Short description of the listing. The description must contain only Unicode characters or tabs  (HT), new lines (LF), carriage returns (CR), and page breaks (FF). Default value is an empty string. Max length: 2000 bytes.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Optional. Type of discovery of the listing on the discovery page.
	DiscoveryType string `json:"discoveryType,omitempty" yaml:"discoveryType,omitempty"`
	// Required. Human-readable display name of the listing. The display name must contain only Unicode letters, numbers (0-9), underscores (_), dashes (-), spaces ( ), ampersands (&) and can't start or end with spaces. Default value is an empty string. Max length: 63 bytes.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// Optional. Documentation describing the listing.
	Documentation string `json:"documentation,omitempty" yaml:"documentation,omitempty"`
	// Immutable. The name of the location this data exchange.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Optional. Email or URL of the primary point of contact of the listing. Max Length: 1000 bytes.
	PrimaryContact string `json:"primaryContact,omitempty" yaml:"primaryContact,omitempty"`
	// The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Optional. Details of the publisher who owns the listing and who can share the source data.
	Publisher map[string]interface{} `json:"publisher,omitempty" yaml:"publisher,omitempty"`
	// Optional. Email or URL of the request access of the listing. Subscribers can use this reference to request access. Max Length: 1000 bytes.
	RequestAccess string `json:"requestAccess,omitempty" yaml:"requestAccess,omitempty"`
	// Immutable. The BigQueryAnalyticsHubDataExchangeListing name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	Source map[string]interface{} `json:"source,omitempty" yaml:"source,omitempty"`
}

// BigQueryAnalyticsHubListingStatus defines the observed state of BigQueryAnalyticsHubListing.
type BigQueryAnalyticsHubListingStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
