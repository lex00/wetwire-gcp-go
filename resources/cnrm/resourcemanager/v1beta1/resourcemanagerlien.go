package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ResourceManagerLien represents a resourcemanager.cnrm.cloud.google.com ResourceManagerLien resource.
type ResourceManagerLien struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ResourceManagerLienSpec   `json:"spec,omitempty"`
	Status ResourceManagerLienStatus `json:"status,omitempty"`
}

// ResourceManagerLienSpec defines the desired state of ResourceManagerLien.
type ResourceManagerLienSpec struct {
	// Immutable. A stable, user-visible/meaningful string identifying the origin
	// of the Lien, intended to be inspected programmatically. Maximum length of
	// 200 characters.
	Origin string `json:"origin,omitempty" yaml:"origin,omitempty"`
	Parent map[string]interface{} `json:"parent,omitempty" yaml:"parent,omitempty"`
	// Immutable. Concise user-visible strings indicating why an action cannot be performed
	// on a resource. Maximum length of 200 characters.
	Reason string `json:"reason,omitempty" yaml:"reason,omitempty"`
	// Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Immutable. The types of operations which should be blocked as a result of this Lien.
	// Each value should correspond to an IAM permission. The server will validate
	// the permissions against those for which Liens are supported.  An empty
	// list is meaningless and will be rejected.
	// e.g. ['resourcemanager.projects.delete'].
	Restrictions []string `json:"restrictions,omitempty" yaml:"restrictions,omitempty"`
}

// ResourceManagerLienStatus defines the observed state of ResourceManagerLien.
type ResourceManagerLienStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
