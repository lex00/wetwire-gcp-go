package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DeploymentManagerDeployment represents a deploymentmanager.cnrm.cloud.google.com DeploymentManagerDeployment resource.
type DeploymentManagerDeployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DeploymentManagerDeploymentSpec   `json:"spec,omitempty"`
	Status DeploymentManagerDeploymentStatus `json:"status,omitempty"`
}

// DeploymentManagerDeploymentSpec defines the desired state of DeploymentManagerDeployment.
type DeploymentManagerDeploymentSpec struct {
	// Immutable. Set the policy to use for creating new resources. Only used on
	// create and update. Valid values are 'CREATE_OR_ACQUIRE' (default) or
	// 'ACQUIRE'. If set to 'ACQUIRE' and resources do not already exist,
	// the deployment will fail. Note that updating this field does not
	// actually affect the deployment, just how it is updated. Default value: "CREATE_OR_ACQUIRE" Possible values: ["ACQUIRE", "CREATE_OR_ACQUIRE"].
	CreatePolicy string `json:"createPolicy,omitempty" yaml:"createPolicy,omitempty"`
	// Immutable. Set the policy to use for deleting new resources on update/delete.
	// Valid values are 'DELETE' (default) or 'ABANDON'. If 'DELETE',
	// resource is deleted after removal from Deployment Manager. If
	// 'ABANDON', the resource is only removed from Deployment Manager
	// and is not actually deleted. Note that updating this field does not
	// actually change the deployment, just how it is updated. Default value: "DELETE" Possible values: ["ABANDON", "DELETE"].
	DeletePolicy string `json:"deletePolicy,omitempty" yaml:"deletePolicy,omitempty"`
	// Optional user-provided description of deployment.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	Preview bool `json:"preview,omitempty" yaml:"preview,omitempty"`
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Parameters that define your deployment, including the deployment
	// configuration and relevant templates.
	Target map[string]interface{} `json:"target,omitempty" yaml:"target,omitempty"`
}

// DeploymentManagerDeploymentStatus defines the observed state of DeploymentManagerDeployment.
type DeploymentManagerDeploymentStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
