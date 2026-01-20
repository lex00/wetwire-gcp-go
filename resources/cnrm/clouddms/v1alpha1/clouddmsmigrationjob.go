package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CloudDMSMigrationJob is the Schema for the CloudDMSMigrationJob API
type CloudDMSMigrationJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CloudDMSMigrationJobSpec   `json:"spec,omitempty"`
	Status CloudDMSMigrationJobStatus `json:"status,omitempty"`
}

// CloudDMSMigrationJobSpec defines the desired state of CloudDMSMigrationJob.
type CloudDMSMigrationJobSpec struct {
	// The CMEK (customer-managed encryption key) fully qualified key name used for the migration job. This field supports all migration jobs types except for: * Mysql to Mysql (use the cmek field in the cloudsql connection profile instead). * PostrgeSQL to PostgreSQL (use the cmek field in the cloudsql connection profile instead). * PostgreSQL to AlloyDB (use the kms_key_name field in the alloydb connection profile instead). Each Cloud CMEK key has the following format: projects/[PROJECT]/locations/[REGION]/keyRings/[RING]/cryptoKeys/[KEY_NAME]
	CmekKeyNameRef map[string]interface{} `json:"cmekKeyNameRef,omitempty" yaml:"cmekKeyNameRef,omitempty"`
	// The conversion workspace used by the migration.
	ConversionWorkspace map[string]interface{} `json:"conversionWorkspace,omitempty" yaml:"conversionWorkspace,omitempty"`
	// The database engine type and provider of the destination.
	DestinationDatabase map[string]interface{} `json:"destinationDatabase,omitempty" yaml:"destinationDatabase,omitempty"`
	// Required. The Connection Profile of the destination connection profile.
	DestinationRef map[string]interface{} `json:"destinationRef,omitempty" yaml:"destinationRef,omitempty"`
	// The migration job display name.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// The initial dump flags. This field and the "dump_path" field are mutually exclusive.
	DumpFlags map[string]interface{} `json:"dumpFlags,omitempty" yaml:"dumpFlags,omitempty"`
	// The path to the dump file in Google Cloud Storage, in the format: (gs://[BUCKET_NAME]/[OBJECT_NAME]). This field and the "dump_flags" field are mutually exclusive.
	DumpPath string `json:"dumpPath,omitempty" yaml:"dumpPath,omitempty"`
	// This field can be used to select the entities to migrate as part of the migration job. It uses AIP-160 notation to select a subset of the entities configured on the associated conversion-workspace. This field should not be set on migration-jobs that are not associated with a conversion workspace.
	Filter string `json:"filter,omitempty" yaml:"filter,omitempty"`
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Optional. Data dump parallelism settings used by the migration. Currently applicable only for MySQL to Cloud SQL for MySQL migrations only.
	PerformanceConfig map[string]interface{} `json:"performanceConfig,omitempty" yaml:"performanceConfig,omitempty"`
	// The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// The CloudDMSMigrationJob name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// The details needed to communicate to the source over Reverse SSH tunnel connectivity.
	ReverseSSHConnectivity map[string]interface{} `json:"reverseSSHConnectivity,omitempty" yaml:"reverseSSHConnectivity,omitempty"`
	// The database engine type and provider of the source.
	SourceDatabase map[string]interface{} `json:"sourceDatabase,omitempty" yaml:"sourceDatabase,omitempty"`
	// Required. The Connection Profile resource of the source connection profile.
	SourceRef map[string]interface{} `json:"sourceRef,omitempty" yaml:"sourceRef,omitempty"`
	// static ip connectivity data (default, no additional details needed).
	StaticIPConnectivity map[string]interface{} `json:"staticIPConnectivity,omitempty" yaml:"staticIPConnectivity,omitempty"`
	// Required. The migration job type.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
	// The details of the VPC network that the source database is located in.
	VpcPeeringConnectivity map[string]interface{} `json:"vpcPeeringConnectivity,omitempty" yaml:"vpcPeeringConnectivity,omitempty"`
}

// CloudDMSMigrationJobStatus defines the observed state of CloudDMSMigrationJob.
type CloudDMSMigrationJobStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
