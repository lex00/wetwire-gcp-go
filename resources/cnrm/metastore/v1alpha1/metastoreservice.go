package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MetastoreService is the Schema for the MetastoreService API
type MetastoreService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MetastoreServiceSpec   `json:"spec,omitempty"`
	Status MetastoreServiceStatus `json:"status,omitempty"`
}

// MetastoreServiceSpec defines the desired state of MetastoreService.
type MetastoreServiceSpec struct {
	// Immutable. The database type that the Metastore service stores its data.
	DatabaseType string `json:"databaseType,omitempty" yaml:"databaseType,omitempty"`
	// Immutable. Information used to configure the Dataproc Metastore service to encrypt customer data at rest. Cannot be updated.
	EncryptionConfig map[string]interface{} `json:"encryptionConfig,omitempty" yaml:"encryptionConfig,omitempty"`
	// Configuration information specific to running Hive metastore software as the metastore service.
	HiveMetastoreConfig map[string]interface{} `json:"hiveMetastoreConfig,omitempty" yaml:"hiveMetastoreConfig,omitempty"`
	// User-defined labels for the metastore service.
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// The one hour maintenance window of the metastore service. This specifies when the service can be restarted for maintenance purposes in UTC time. Maintenance window is not needed for services with the SPANNER database type.
	MaintenanceWindow map[string]interface{} `json:"maintenanceWindow,omitempty" yaml:"maintenanceWindow,omitempty"`
	// The configuration specifying the network settings for the Dataproc Metastore service.
	NetworkConfig map[string]interface{} `json:"networkConfig,omitempty" yaml:"networkConfig,omitempty"`
	// Immutable. The relative resource name of the VPC network on which the instance can be accessed.
	NetworkRef map[string]interface{} `json:"networkRef,omitempty" yaml:"networkRef,omitempty"`
	// The TCP port at which the metastore service is reached. Default: 9083.
	Port int32 `json:"port,omitempty" yaml:"port,omitempty"`
	// The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. The release channel of the service. If unspecified, defaults to `STABLE`.
	ReleaseChannel string `json:"releaseChannel,omitempty" yaml:"releaseChannel,omitempty"`
	// The MetastoreService name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Scaling configuration of the metastore service.
	ScalingConfig map[string]interface{} `json:"scalingConfig,omitempty" yaml:"scalingConfig,omitempty"`
	// The configuration specifying telemetry settings for the Dataproc Metastore service. If unspecified defaults to `JSON`.
	TelemetryConfig map[string]interface{} `json:"telemetryConfig,omitempty" yaml:"telemetryConfig,omitempty"`
	// The tier of the service.
	Tier string `json:"tier,omitempty" yaml:"tier,omitempty"`
}

// MetastoreServiceStatus defines the observed state of MetastoreService.
type MetastoreServiceStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
