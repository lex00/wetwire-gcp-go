package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// FilestoreInstance represents a filestore.cnrm.cloud.google.com FilestoreInstance resource.
type FilestoreInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FilestoreInstanceSpec   `json:"spec,omitempty"`
	Status FilestoreInstanceStatus `json:"status,omitempty"`
}

// FilestoreInstanceSpec defines the desired state of FilestoreInstance.
type FilestoreInstanceSpec struct {
	// The description of the instance (2048 characters or less).
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// File system shares on the instance. For this version, only a single file share is supported.
	FileShares []map[string]interface{} `json:"fileShares,omitempty" yaml:"fileShares,omitempty"`
	// Immutable. The location for the resource
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Immutable. VPC networks to which the instance is connected. For this version, only a single network is supported.
	Networks []map[string]interface{} `json:"networks,omitempty" yaml:"networks,omitempty"`
	// Immutable. The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Immutable. The service tier of the instance. Possible values: TIER_UNSPECIFIED, STANDARD, PREMIUM, BASIC_HDD, BASIC_SSD, HIGH_SCALE_SSD, ENTERPRISE
	Tier string `json:"tier,omitempty" yaml:"tier,omitempty"`
}

// FilestoreInstanceStatus defines the observed state of FilestoreInstance.
type FilestoreInstanceStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
