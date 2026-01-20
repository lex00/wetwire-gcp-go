package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// OSConfigGuestPolicy represents a osconfig.cnrm.cloud.google.com OSConfigGuestPolicy resource.
type OSConfigGuestPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OSConfigGuestPolicySpec   `json:"spec,omitempty"`
	Status OSConfigGuestPolicyStatus `json:"status,omitempty"`
}

// OSConfigGuestPolicySpec defines the desired state of OSConfigGuestPolicy.
type OSConfigGuestPolicySpec struct {
	// Specifies the VMs that are assigned this policy. This allows you to target sets or groups of VMs by different parameters such as labels, names, OS, or zones. Empty assignments will target ALL VMs underneath this policy. Conflict Management Policies that exist higher up in the resource hierarchy (closer to the Org) will override those lower down if there is a conflict. At the same level in the resource hierarchy (ie. within a project), the service will prevent the creation of multiple policies that conflict with each other. If there are multiple policies that specify the same config (eg. package, software recipe, repository, etc.), the service will ensure that no VM could potentially receive instructions from both policies. To create multiple policies that specify different versions of a package or different configs for different Operating Systems, each policy must be mutually exclusive in their targeting according to labels, OS, or other criteria. Different configs are identified for conflicts in different ways. Packages are identified by their name and the package manager(s) they target. Package repositories are identified by their unique id where applicable. Some package managers don't have a unique identifier for repositories and where that's the case, no uniqueness is validated by the service. Note that if OS Inventory is disabled, a VM will not be assigned a policy that targets by OS because the service will see this VM's OS as unknown.
	Assignment map[string]interface{} `json:"assignment,omitempty" yaml:"assignment,omitempty"`
	// Description of the GuestPolicy. Length of the description is limited to 1024 characters.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// List of package repository configurations assigned to the VM instance.
	PackageRepositories []map[string]interface{} `json:"packageRepositories,omitempty" yaml:"packageRepositories,omitempty"`
	// List of package configurations assigned to the VM instance.
	Packages []map[string]interface{} `json:"packages,omitempty" yaml:"packages,omitempty"`
	// Optional. A list of Recipes to install on the VM.
	Recipes []map[string]interface{} `json:"recipes,omitempty" yaml:"recipes,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// OSConfigGuestPolicyStatus defines the observed state of OSConfigGuestPolicy.
type OSConfigGuestPolicyStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
