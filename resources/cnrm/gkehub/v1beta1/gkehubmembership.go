package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GKEHubMembership represents a gkehub.cnrm.cloud.google.com GKEHubMembership resource.
type GKEHubMembership struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GKEHubMembershipSpec   `json:"spec,omitempty"`
	Status GKEHubMembershipStatus `json:"status,omitempty"`
}

// GKEHubMembershipSpec defines the desired state of GKEHubMembership.
type GKEHubMembershipSpec struct {
	// Optional. How to identify workloads from this Membership. See the documentation on Workload Identity for more details: https://cloud.google.com/kubernetes-engine/docs/how-to/workload-identity
	Authority map[string]interface{} `json:"authority,omitempty" yaml:"authority,omitempty"`
	// Description of this membership, limited to 63 characters. Must match the regex: `*` This field is present for legacy purposes.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Optional. Endpoint information to reach this member.
	Endpoint map[string]interface{} `json:"endpoint,omitempty" yaml:"endpoint,omitempty"`
	// Optional. An externally-generated and managed ID for this Membership. This ID may be modified after creation, but this is not recommended. The ID must match the regex: `*` If this Membership represents a Kubernetes cluster, this value should be set to the UID of the `kube-system` namespace object.
	ExternalID string `json:"externalId,omitempty" yaml:"externalId,omitempty"`
	// Optional. The infrastructure type this Membership is running on. Possible values: INFRASTRUCTURE_TYPE_UNSPECIFIED, ON_PREM, MULTI_CLOUD
	InfrastructureType string `json:"infrastructureType,omitempty" yaml:"infrastructureType,omitempty"`
	// Immutable. The location for the resource
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// GKEHubMembershipStatus defines the observed state of GKEHubMembership.
type GKEHubMembershipStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
