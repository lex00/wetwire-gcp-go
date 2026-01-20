package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// KMSKeyHandle is the Schema for the KMSKeyHandle API
type KMSKeyHandle struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KMSKeyHandleSpec   `json:"spec,omitempty"`
	Status KMSKeyHandleStatus `json:"status,omitempty"`
}

// KMSKeyHandleSpec defines the desired state of KMSKeyHandle.
type KMSKeyHandleSpec struct {
	// Location name to create KeyHandle
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Project hosting KMSKeyHandle
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// The KMS Key Handle ID used for resource creation or acquisition. For creation: If specified, this value is used as the key handle ID. If not provided, a UUID will be generated and assigned as the key handle ID. For acquisition: This field must be provided to identify the key handle resource to acquire.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Indicates the resource type that the resulting [CryptoKey][] is meant to protect, e.g. `{SERVICE}.googleapis.com/{TYPE}`. See documentation for supported resource types https://cloud.google.com/kms/docs/autokey-overview#compatible-services.
	ResourceTypeSelector string `json:"resourceTypeSelector,omitempty" yaml:"resourceTypeSelector,omitempty"`
}

// KMSKeyHandleStatus defines the observed state of KMSKeyHandle.
type KMSKeyHandleStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
