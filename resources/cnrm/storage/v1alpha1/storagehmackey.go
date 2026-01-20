package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// StorageHMACKey represents a storage.cnrm.cloud.google.com StorageHMACKey resource.
type StorageHMACKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   StorageHMACKeySpec   `json:"spec,omitempty"`
	Status StorageHMACKeyStatus `json:"status,omitempty"`
}

// StorageHMACKeySpec defines the desired state of StorageHMACKey.
type StorageHMACKeySpec struct {
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The service-generated accessId of the resource. Used for acquisition only. Leave unset to create a new resource.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Immutable. The email address of the key's associated service account.
	ServiceAccountEmail string `json:"serviceAccountEmail,omitempty" yaml:"serviceAccountEmail,omitempty"`
	// The state of the key. Can be set to one of ACTIVE, INACTIVE. Default value: "ACTIVE" Possible values: ["ACTIVE", "INACTIVE"].
	State string `json:"state,omitempty" yaml:"state,omitempty"`
}

// StorageHMACKeyStatus defines the observed state of StorageHMACKey.
type StorageHMACKeyStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
