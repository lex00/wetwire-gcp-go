package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// SQLInstance is the Schema for the sql API
type SQLInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SQLInstanceSpec   `json:"spec,omitempty"`
	Status SQLInstanceStatus `json:"status,omitempty"`
}

// SQLInstanceSpec defines the desired state of SQLInstance.
type SQLInstanceSpec struct {
	// Create this database as a clone of a source instance. Immutable.
	CloneSource map[string]interface{} `json:"cloneSource,omitempty" yaml:"cloneSource,omitempty"`
	// The MySQL, PostgreSQL or SQL Server (beta) version to use. Supported values include MYSQL_5_6, MYSQL_5_7, MYSQL_8_0, POSTGRES_9_6, POSTGRES_10, POSTGRES_11, POSTGRES_12, POSTGRES_13, POSTGRES_14, POSTGRES_15, SQLSERVER_2017_STANDARD, SQLSERVER_2017_ENTERPRISE, SQLSERVER_2017_EXPRESS, SQLSERVER_2017_WEB. Database Version Policies includes an up-to-date reference of supported versions.
	DatabaseVersion string `json:"databaseVersion,omitempty" yaml:"databaseVersion,omitempty"`
	EncryptionKMSCryptoKeyRef map[string]interface{} `json:"encryptionKMSCryptoKeyRef,omitempty" yaml:"encryptionKMSCryptoKeyRef,omitempty"`
	// The type of the instance. The valid values are:- 'SQL_INSTANCE_TYPE_UNSPECIFIED', 'CLOUD_SQL_INSTANCE', 'ON_PREMISES_INSTANCE' and 'READ_REPLICA_INSTANCE'.
	InstanceType string `json:"instanceType,omitempty" yaml:"instanceType,omitempty"`
	// Maintenance version.
	MaintenanceVersion string `json:"maintenanceVersion,omitempty" yaml:"maintenanceVersion,omitempty"`
	MasterInstanceRef map[string]interface{} `json:"masterInstanceRef,omitempty" yaml:"masterInstanceRef,omitempty"`
	// Immutable. The region the instance will sit in. Note, Cloud SQL is not available in all regions. A valid region must be provided to use this resource. If a region is not provided in the resource definition, the provider region will be used instead, but this will be an apply-time error for instances if the provider region is not supported with Cloud SQL. If you choose not to provide the region argument for this resource, make sure you understand this.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`
	// The configuration for replication.
	ReplicaConfiguration map[string]interface{} `json:"replicaConfiguration,omitempty" yaml:"replicaConfiguration,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Initial root password. Required for MS SQL Server.
	RootPassword map[string]interface{} `json:"rootPassword,omitempty" yaml:"rootPassword,omitempty"`
	// The settings to use for the database. The configuration is detailed below.
	Settings map[string]interface{} `json:"settings,omitempty" yaml:"settings,omitempty"`
}

// SQLInstanceStatus defines the observed state of SQLInstance.
type SQLInstanceStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
