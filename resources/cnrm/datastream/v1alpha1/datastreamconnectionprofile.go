package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DatastreamConnectionProfile is the Schema for the DatastreamConnectionProfile API
type DatastreamConnectionProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DatastreamConnectionProfileSpec   `json:"spec,omitempty"`
	Status DatastreamConnectionProfileStatus `json:"status,omitempty"`
}

// DatastreamConnectionProfileSpec defines the desired state of DatastreamConnectionProfile.
type DatastreamConnectionProfileSpec struct {
	// BigQuery Connection Profile configuration.
	BigQueryProfile map[string]interface{} `json:"bigQueryProfile,omitempty" yaml:"bigQueryProfile,omitempty"`
	// Required. Display name.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// Forward SSH tunnel connectivity.
	ForwardSSHConnectivity map[string]interface{} `json:"forwardSSHConnectivity,omitempty" yaml:"forwardSSHConnectivity,omitempty"`
	// Cloud Storage ConnectionProfile configuration.
	GcsProfile map[string]interface{} `json:"gcsProfile,omitempty" yaml:"gcsProfile,omitempty"`
	// Labels.
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	// Immutable.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// MySQL ConnectionProfile configuration.
	MySQLProfile map[string]interface{} `json:"mySQLProfile,omitempty" yaml:"mySQLProfile,omitempty"`
	// Oracle ConnectionProfile configuration.
	OracleProfile map[string]interface{} `json:"oracleProfile,omitempty" yaml:"oracleProfile,omitempty"`
	// Private connectivity.
	PrivateConnectivity map[string]interface{} `json:"privateConnectivity,omitempty" yaml:"privateConnectivity,omitempty"`
	// The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// The DatastreamConnectionProfile name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// SQLServer Connection Profile configuration.
	SqlServerProfile map[string]interface{} `json:"sqlServerProfile,omitempty" yaml:"sqlServerProfile,omitempty"`
	// Static Service IP connectivity.
	StaticServiceIPConnectivity map[string]interface{} `json:"staticServiceIPConnectivity,omitempty" yaml:"staticServiceIPConnectivity,omitempty"`
}

// DatastreamConnectionProfileStatus defines the observed state of DatastreamConnectionProfile.
type DatastreamConnectionProfileStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
