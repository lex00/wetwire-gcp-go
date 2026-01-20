package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// OSConfigPatchDeployment represents a osconfig.cnrm.cloud.google.com OSConfigPatchDeployment resource.
type OSConfigPatchDeployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OSConfigPatchDeploymentSpec   `json:"spec,omitempty"`
	Status OSConfigPatchDeploymentStatus `json:"status,omitempty"`
}

// OSConfigPatchDeploymentSpec defines the desired state of OSConfigPatchDeployment.
type OSConfigPatchDeploymentSpec struct {
	// Immutable. Description of the patch deployment. Length of the description is limited to 1024 characters.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Immutable. Duration of the patch. After the duration ends, the patch times out.
	// A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
	Duration string `json:"duration,omitempty" yaml:"duration,omitempty"`
	// Immutable. VM instances to patch.
	InstanceFilter map[string]interface{} `json:"instanceFilter,omitempty" yaml:"instanceFilter,omitempty"`
	// Immutable. Schedule a one-time execution.
	OneTimeSchedule map[string]interface{} `json:"oneTimeSchedule,omitempty" yaml:"oneTimeSchedule,omitempty"`
	// Immutable. Patch configuration that is applied.
	PatchConfig map[string]interface{} `json:"patchConfig,omitempty" yaml:"patchConfig,omitempty"`
	// Immutable. A name for the patch deployment in the project. When creating a name the following rules apply:
	// * Must contain only lowercase letters, numbers, and hyphens.
	// * Must start with a letter.
	// * Must be between 1-63 characters.
	// * Must end with a number or a letter.
	// * Must be unique within the project.
	PatchDeploymentID string `json:"patchDeploymentId,omitempty" yaml:"patchDeploymentId,omitempty"`
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Schedule recurring executions.
	RecurringSchedule map[string]interface{} `json:"recurringSchedule,omitempty" yaml:"recurringSchedule,omitempty"`
	// Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Immutable. Rollout strategy of the patch job.
	Rollout map[string]interface{} `json:"rollout,omitempty" yaml:"rollout,omitempty"`
}

// OSConfigPatchDeploymentStatus defines the observed state of OSConfigPatchDeployment.
type OSConfigPatchDeploymentStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
