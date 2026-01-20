package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// AlloyDBCluster is the Schema for the AlloyDBCluster API
type AlloyDBCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AlloyDBClusterSpec   `json:"spec,omitempty"`
	Status AlloyDBClusterStatus `json:"status,omitempty"`
}

// AlloyDBClusterSpec defines the desired state of AlloyDBCluster.
type AlloyDBClusterSpec struct {
	// The automated backup policy for this cluster.
	// If no policy is provided then the default policy will be used. If backups
	// are supported for the cluster, the default policy takes one backup a day,
	// has a backup window of 1 hour, and retains backups for 14 days.
	// For more information on the defaults, consult the
	// documentation for the message type.
	AutomatedBackupPolicy map[string]interface{} `json:"automatedBackupPolicy,omitempty" yaml:"automatedBackupPolicy,omitempty"`
	// The type of cluster. If not set, defaults to PRIMARY. Default value: "PRIMARY" Possible values: ["PRIMARY", "SECONDARY"].
	ClusterType string `json:"clusterType,omitempty" yaml:"clusterType,omitempty"`
	// Optional. Continuous backup configuration for this cluster.
	ContinuousBackupConfig map[string]interface{} `json:"continuousBackupConfig,omitempty" yaml:"continuousBackupConfig,omitempty"`
	// Optional. The database engine major version. This is an optional field and it is populated at the Cluster creation time. If a database version is not supplied at cluster creation time, then a default database version will be used.
	DatabaseVersion string `json:"databaseVersion,omitempty" yaml:"databaseVersion,omitempty"`
	// Policy to determine if the cluster should be deleted forcefully. Deleting a cluster forcefully, deletes the cluster and all its associated instances within the cluster. Deleting a Secondary cluster with a secondary instance REQUIRES setting deletion_policy = "FORCE" otherwise an error is returned. This is needed as there is no support to delete just the secondary instance, and the only way to delete secondary instance is to delete the associated secondary cluster forcefully which also deletes the secondary instance.
	DeletionPolicy string `json:"deletionPolicy,omitempty" yaml:"deletionPolicy,omitempty"`
	// User-settable and human-readable display name for the Cluster.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// Optional. The encryption config can be specified to encrypt the data disks and other persistent data resources of a cluster with a customer-managed encryption key (CMEK). When this field is not specified, the cluster will then use default encryption scheme to protect the user data.
	EncryptionConfig map[string]interface{} `json:"encryptionConfig,omitempty" yaml:"encryptionConfig,omitempty"`
	// Input only. Initial user to setup during cluster creation. Required. If used in `RestoreCluster` this is ignored.
	InitialUser map[string]interface{} `json:"initialUser,omitempty" yaml:"initialUser,omitempty"`
	// Immutable. The location where the alloydb cluster should reside.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Optional. The maintenance update policy determines when to allow or deny updates.
	MaintenanceUpdatePolicy map[string]interface{} `json:"maintenanceUpdatePolicy,omitempty" yaml:"maintenanceUpdatePolicy,omitempty"`
	NetworkConfig map[string]interface{} `json:"networkConfig,omitempty" yaml:"networkConfig,omitempty"`
	// The resource link for the VPC network in which cluster resources are created and from which they are accessible via Private IP. The network must belong to the same project as the cluster. It is specified in the form: `projects/{project}/global/networks/{network_id}`. This is required to create a cluster. Deprecated, use network_config.network instead.
	NetworkRef map[string]interface{} `json:"networkRef,omitempty" yaml:"networkRef,omitempty"`
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// The AlloyDBCluster name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Immutable. The source when restoring from a backup. Conflicts with 'restoreContinuousBackupSource', both can't be set together.
	RestoreBackupSource map[string]interface{} `json:"restoreBackupSource,omitempty" yaml:"restoreBackupSource,omitempty"`
	// Immutable. The source when restoring via point in time recovery (PITR). Conflicts with 'restoreBackupSource', both can't be set together.
	RestoreContinuousBackupSource map[string]interface{} `json:"restoreContinuousBackupSource,omitempty" yaml:"restoreContinuousBackupSource,omitempty"`
	// Cross Region replication config specific to SECONDARY cluster.
	SecondaryConfig map[string]interface{} `json:"secondaryConfig,omitempty" yaml:"secondaryConfig,omitempty"`
}

// AlloyDBClusterStatus defines the observed state of AlloyDBCluster.
type AlloyDBClusterStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
