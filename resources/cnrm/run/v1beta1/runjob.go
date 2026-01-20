package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// RunJob is the Schema for the RunJob API
type RunJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RunJobSpec   `json:"spec,omitempty"`
	Status RunJobStatus `json:"status,omitempty"`
}

// RunJobSpec defines the desired state of RunJob.
type RunJobSpec struct {
	// Optional. User-provided annotations, which are stored in GCP.
	Annotations map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	// Optional. Settings for Binary Authorization feature.
	BinaryAuthorization map[string]interface{} `json:"binaryAuthorization,omitempty" yaml:"binaryAuthorization,omitempty"`
	// Optional. Arbitrary identifier for the API client.
	Client string `json:"client,omitempty" yaml:"client,omitempty"`
	// Optional. Arbitrary version identifier for the API client.
	ClientVersion string `json:"clientVersion,omitempty" yaml:"clientVersion,omitempty"`
	// Optional. The launch stage of the job. Possible values are `LAUNCH_STAGE_UNSPECIFIED`, `UNIMPLEMENTED`, `PRELAUNCH`, `EARLY_ACCESS`, `ALPHA`, `BETA`, `GA`, `DEPRECATED`.
	LaunchStage string `json:"launchStage,omitempty" yaml:"launchStage,omitempty"`
	// The location of the cloud run job
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// The RunJob name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Required. The template used to create executions for this Job.
	Template map[string]interface{} `json:"template,omitempty" yaml:"template,omitempty"`
}

// RunJobStatus defines the observed state of RunJob.
type RunJobStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
