package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// WorkstationConfig is the Schema for the WorkstationConfig API
type WorkstationConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WorkstationConfigSpec   `json:"spec,omitempty"`
	Status WorkstationConfigStatus `json:"status,omitempty"`
}

// WorkstationConfigSpec defines the desired state of WorkstationConfig.
type WorkstationConfigSpec struct {
	// Optional. Client-specified annotations.
	Annotations []map[string]interface{} `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	// Optional. Container that runs upon startup for each workstation using this workstation configuration.
	Container map[string]interface{} `json:"container,omitempty" yaml:"container,omitempty"`
	// Optional. Human-readable name for this workstation configuration.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// Immutable. Encrypts resources of this workstation configuration using a
	// customer-managed encryption key (CMEK).
	// If specified, the boot disk of the Compute Engine instance and the
	// persistent disk are encrypted using this encryption key. If
	// this field is not set, the disks are encrypted using a generated
	// key. Customer-managed encryption keys do not protect disk metadata.
	// If the customer-managed encryption key is rotated, when the workstation
	// instance is stopped, the system attempts to recreate the
	// persistent disk with the new version of the key. Be sure to keep
	// older versions of the key until the persistent disk is recreated.
	// Otherwise, data on the persistent disk might be lost.
	// If the encryption key is revoked, the workstation session automatically
	// stops within 7 hours.
	// Immutable after the workstation configuration is created.
	EncryptionKey map[string]interface{} `json:"encryptionKey,omitempty" yaml:"encryptionKey,omitempty"`
	// Optional. Runtime host for the workstation.
	Host map[string]interface{} `json:"host,omitempty" yaml:"host,omitempty"`
	// Optional. Number of seconds to wait before automatically stopping a
	// workstation after it last received user traffic.
	// A value of `"0s"` indicates that Cloud Workstations VMs created with this
	// configuration should never time out due to idleness.
	// Provide
	// [duration](https://developers.google.com/protocol-buffers/docs/reference/google.protobuf#duration)
	// terminated by `s` for seconds—for example, `"7200s"` (2 hours).
	// The default is `"1200s"` (20 minutes).
	IdleTimeout string `json:"idleTimeout,omitempty" yaml:"idleTimeout,omitempty"`
	// Optional. [Labels](https://cloud.google.com/workstations/docs/label-resources) that are applied to the workstation configuration and that are also propagated to the underlying Compute Engine resources.
	Labels []map[string]interface{} `json:"labels,omitempty" yaml:"labels,omitempty"`
	// Parent is a reference to the parent WorkstationCluster for this WorkstationConfig.
	ParentRef map[string]interface{} `json:"parentRef,omitempty" yaml:"parentRef,omitempty"`
	// Optional. Directories to persist across workstation sessions.
	PersistentDirectories []map[string]interface{} `json:"persistentDirectories,omitempty" yaml:"persistentDirectories,omitempty"`
	// Optional. Readiness checks to perform when starting a workstation using this workstation configuration. Mark a workstation as running only after all specified readiness checks return 200 status codes.
	ReadinessChecks []map[string]interface{} `json:"readinessChecks,omitempty" yaml:"readinessChecks,omitempty"`
	// Optional. Immutable. Specifies the zones used to replicate the VM and disk
	// resources within the region. If set, exactly two zones within the
	// workstation cluster's region must be specified—for example,
	// `['us-central1-a', 'us-central1-f']`. If this field is empty, two default
	// zones within the region are used.
	// Immutable after the workstation configuration is created.
	ReplicaZones []string `json:"replicaZones,omitempty" yaml:"replicaZones,omitempty"`
	// The WorkstationConfig name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Optional. Number of seconds that a workstation can run until it is
	// automatically shut down. We recommend that workstations be shut down daily
	// to reduce costs and so that security updates can be applied upon restart.
	// The
	// [idle_timeout][google.cloud.workstations.v1.WorkstationConfig.idle_timeout]
	// and
	// [running_timeout][google.cloud.workstations.v1.WorkstationConfig.running_timeout]
	// fields are independent of each other. Note that the
	// [running_timeout][google.cloud.workstations.v1.WorkstationConfig.running_timeout]
	// field shuts down VMs after the specified time, regardless of whether or not
	// the VMs are idle.
	// Provide duration terminated by `s` for seconds—for example, `"54000s"`
	// (15 hours). Defaults to `"43200s"` (12 hours). A value of `"0s"` indicates
	// that workstations using this configuration should never time out. If
	// [encryption_key][google.cloud.workstations.v1.WorkstationConfig.encryption_key]
	// is set, it must be greater than `"0s"` and less than
	// `"86400s"` (24 hours).
	// Warning: A value of `"0s"` indicates that Cloud Workstations VMs created
	// with this configuration have no maximum running time. This is strongly
	// discouraged because you incur costs and will not pick up security updates.
	RunningTimeout string `json:"runningTimeout,omitempty" yaml:"runningTimeout,omitempty"`
}

// WorkstationConfigStatus defines the observed state of WorkstationConfig.
type WorkstationConfigStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
