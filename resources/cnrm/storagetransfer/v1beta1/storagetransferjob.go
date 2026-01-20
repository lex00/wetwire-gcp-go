package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// StorageTransferJob represents a storagetransfer.cnrm.cloud.google.com StorageTransferJob resource.
type StorageTransferJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   StorageTransferJobSpec   `json:"spec,omitempty"`
	Status StorageTransferJobStatus `json:"status,omitempty"`
}

// StorageTransferJobSpec defines the desired state of StorageTransferJob.
type StorageTransferJobSpec struct {
	// Unique description to identify the Transfer Job.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Notification configuration.
	NotificationConfig map[string]interface{} `json:"notificationConfig,omitempty" yaml:"notificationConfig,omitempty"`
	// Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Schedule specification defining when the Transfer Job should be scheduled to start, end and what time to run.
	Schedule map[string]interface{} `json:"schedule,omitempty" yaml:"schedule,omitempty"`
	// Status of the job. Default: ENABLED. NOTE: The effect of the new job status takes place during a subsequent job run. For example, if you change the job status from ENABLED to DISABLED, and an operation spawned by the transfer is running, the status change would not affect the current operation.
	Status string `json:"status,omitempty" yaml:"status,omitempty"`
	// Transfer specification.
	TransferSpec map[string]interface{} `json:"transferSpec,omitempty" yaml:"transferSpec,omitempty"`
}

// StorageTransferJobStatus defines the observed state of StorageTransferJob.
type StorageTransferJobStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
