package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DataflowFlexTemplateJob is the Schema for the DataflowFlexTemplateJob API
type DataflowFlexTemplateJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DataflowFlexTemplateJobSpec   `json:"spec,omitempty"`
	Status DataflowFlexTemplateJobStatus `json:"status,omitempty"`
}

// DataflowFlexTemplateJobSpec defines the desired state of DataflowFlexTemplateJob.
type DataflowFlexTemplateJobSpec struct {
	// Additional experiment flags for the job.
	AdditionalExperiments []string `json:"additionalExperiments,omitempty" yaml:"additionalExperiments,omitempty"`
	// The algorithm to use for autoscaling
	AutoscalingAlgorithm string `json:"autoscalingAlgorithm,omitempty" yaml:"autoscalingAlgorithm,omitempty"`
	// Cloud Storage path to a file with json serialized ContainerSpec as content.
	ContainerSpecGcsPath string `json:"containerSpecGcsPath,omitempty" yaml:"containerSpecGcsPath,omitempty"`
	// Whether to enable Streaming Engine for the job.
	EnableStreamingEngine bool `json:"enableStreamingEngine,omitempty" yaml:"enableStreamingEngine,omitempty"`
	// Configuration for VM IPs.
	IPConfiguration string `json:"ipConfiguration,omitempty" yaml:"ipConfiguration,omitempty"`
	// The Cloud KMS key for the job.
	KmsKeyNameRef map[string]interface{} `json:"kmsKeyNameRef,omitempty" yaml:"kmsKeyNameRef,omitempty"`
	// The machine type to use for launching the job. The default is n1-standard-1.
	LauncherMachineType string `json:"launcherMachineType,omitempty" yaml:"launcherMachineType,omitempty"`
	// The machine type to use for the job. Defaults to the value from the template if not specified.
	MachineType string `json:"machineType,omitempty" yaml:"machineType,omitempty"`
	// The maximum number of Google Compute Engine instances to be made available to your pipeline during execution, from 1 to 1000.
	MaxWorkers int32 `json:"maxWorkers,omitempty" yaml:"maxWorkers,omitempty"`
	// Network to which VMs will be assigned.  If empty or unspecified, the service will use the network "default".
	NetworkRef map[string]interface{} `json:"networkRef,omitempty" yaml:"networkRef,omitempty"`
	// The initial number of Google Compute Engine instances for the job.
	NumWorkers int32 `json:"numWorkers,omitempty" yaml:"numWorkers,omitempty"`
	// The parameters for FlexTemplate. Ex. {"num_workers":"5"}
	Parameters map[string]interface{} `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	// Immutable. The region in which the created job should run.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`
	// Docker registry location of container image to use for the 'worker harness. Default is the container for the version of the SDK. Note this field is only valid for portable pipelines.
	SdkContainerImage string `json:"sdkContainerImage,omitempty" yaml:"sdkContainerImage,omitempty"`
	// The email address of the service account to run the job as.
	ServiceAccountEmailRef map[string]interface{} `json:"serviceAccountEmailRef,omitempty" yaml:"serviceAccountEmailRef,omitempty"`
	// The Cloud Storage path for staging local files. Must be a valid Cloud Storage URL, beginning with `gs://`.
	StagingLocation string `json:"stagingLocation,omitempty" yaml:"stagingLocation,omitempty"`
	// Subnetwork to which VMs will be assigned, if desired. You can specify a subnetwork using either a complete URL or an abbreviated path. Expected to be of the form "https://www.googleapis.com/compute/v1/projects/HOST_PROJECT_ID/regions/REGION/subnetworks/SUBNETWORK" or "regions/REGION/subnetworks/SUBNETWORK". If the subnetwork is located in a Shared VPC network, you must use the complete URL.
	SubnetworkRef map[string]interface{} `json:"subnetworkRef,omitempty" yaml:"subnetworkRef,omitempty"`
	// The Cloud Storage path to use for temporary files. Must be a valid Cloud Storage URL, beginning with `gs://`.
	TempLocation string `json:"tempLocation,omitempty" yaml:"tempLocation,omitempty"`
	// Map of transform name prefixes of the job to be replaced with the corresponding name prefixes of the new job. Only applicable when updating a pipeline.
	TransformNameMapping map[string]interface{} `json:"transformNameMapping,omitempty" yaml:"transformNameMapping,omitempty"`
}

// DataflowFlexTemplateJobStatus defines the observed state of DataflowFlexTemplateJob.
type DataflowFlexTemplateJobStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
