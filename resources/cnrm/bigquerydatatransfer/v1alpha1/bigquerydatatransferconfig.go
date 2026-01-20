package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// BigQueryDataTransferConfig is the Schema for the BigQueryDataTransferConfig API
type BigQueryDataTransferConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BigQueryDataTransferConfigSpec   `json:"spec,omitempty"`
	Status BigQueryDataTransferConfigStatus `json:"status,omitempty"`
}

// BigQueryDataTransferConfigSpec defines the desired state of BigQueryDataTransferConfig.
type BigQueryDataTransferConfigSpec struct {
	// The number of days to look back to automatically refresh the data. For example, if `data_refresh_window_days = 10`, then every day BigQuery reingests data for [today-10, today-1], rather than ingesting data for just [today-1]. Only valid if the data source supports the feature. Set the value to 0 to use the default value.
	DataRefreshWindowDays int32 `json:"dataRefreshWindowDays,omitempty" yaml:"dataRefreshWindowDays,omitempty"`
	// Immutable. Data source ID. This cannot be changed once data transfer is created. The full list of available data source IDs can be returned through an API call: https://cloud.google.com/bigquery-transfer/docs/reference/datatransfer/rest/v1/projects.locations.dataSources/list
	DataSourceID string `json:"dataSourceID,omitempty" yaml:"dataSourceID,omitempty"`
	// The BigQuery target dataset id.
	DatasetRef map[string]interface{} `json:"datasetRef,omitempty" yaml:"datasetRef,omitempty"`
	// Is this config disabled. When set to true, no runs will be scheduled for this transfer config.
	Disabled bool `json:"disabled,omitempty" yaml:"disabled,omitempty"`
	// User specified display name for the data transfer.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// Email notifications will be sent according to these preferences to the email address of the user who owns this transfer config.
	EmailPreferences map[string]interface{} `json:"emailPreferences,omitempty" yaml:"emailPreferences,omitempty"`
	// The encryption configuration part. Currently, it is only used for the optional KMS key name. The BigQuery service account of your project must be granted permissions to use the key. Read methods will return the key name applied in effect. Write methods will apply the key if it is present, or otherwise try to apply project default keys if it is absent.
	EncryptionConfiguration map[string]interface{} `json:"encryptionConfiguration,omitempty" yaml:"encryptionConfiguration,omitempty"`
	// Immutable.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Parameters specific to each data source. For more information see the bq tab in the 'Setting up a data transfer' section for each data source. For example the parameters for Cloud Storage transfers are listed here: https://cloud.google.com/bigquery-transfer/docs/cloud-storage-transfer#bq
	Params map[string]string `json:"params,omitempty" yaml:"params,omitempty"`
	// The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Pub/Sub topic where notifications will be sent after transfer runs associated with this transfer config finish.
	PubSubTopicRef map[string]interface{} `json:"pubSubTopicRef,omitempty" yaml:"pubSubTopicRef,omitempty"`
	// The BigQueryDataTransferConfig name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Data transfer schedule.
	// If the data source does not support a custom schedule, this should be
	// empty. If it is empty, the default value for the data source will be used.
	// The specified times are in UTC.
	// Examples of valid format:
	// `1st,3rd monday of month 15:30`,
	// `every wed,fri of jan,jun 13:15`, and
	// `first sunday of quarter 00:00`.
	// See more explanation about the format here:
	// https://cloud.google.com/appengine/docs/flexible/python/scheduling-jobs-with-cron-yaml#the_schedule_format
	// NOTE: The minimum interval time between recurring transfers depends on the
	// data source; refer to the documentation for your data source.
	Schedule string `json:"schedule,omitempty" yaml:"schedule,omitempty"`
	// Options customizing the data transfer schedule.
	ScheduleOptions map[string]interface{} `json:"scheduleOptions,omitempty" yaml:"scheduleOptions,omitempty"`
	// Service account email. If this field is set, the transfer config will be created with this service account's credentials. It requires that the requesting user calling this API has permissions to act as this service account. Note that not all data sources support service account credentials when creating a transfer config. For the latest list of data sources, please refer to https://cloud.google.com/bigquery/docs/use-service-accounts.
	ServiceAccountRef map[string]interface{} `json:"serviceAccountRef,omitempty" yaml:"serviceAccountRef,omitempty"`
}

// BigQueryDataTransferConfigStatus defines the observed state of BigQueryDataTransferConfig.
type BigQueryDataTransferConfigStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
