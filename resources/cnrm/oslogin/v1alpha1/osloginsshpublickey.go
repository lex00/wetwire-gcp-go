package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// OSLoginSSHPublicKey represents a oslogin.cnrm.cloud.google.com OSLoginSSHPublicKey resource.
type OSLoginSSHPublicKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OSLoginSSHPublicKeySpec   `json:"spec,omitempty"`
	Status OSLoginSSHPublicKeyStatus `json:"status,omitempty"`
}

// OSLoginSSHPublicKeySpec defines the desired state of OSLoginSSHPublicKey.
type OSLoginSSHPublicKeySpec struct {
	// An expiration time in microseconds since epoch.
	ExpirationTimeUsec string `json:"expirationTimeUsec,omitempty" yaml:"expirationTimeUsec,omitempty"`
	// Immutable. Public key text in SSH format, defined by RFC4253 section 6.6.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`
	// Immutable. The project ID of the Google Cloud Platform project.
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
	// Immutable. Optional. The service-generated fingerprint of the resource. Used for acquisition only. Leave unset to create a new resource.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Immutable. The user email.
	User string `json:"user,omitempty" yaml:"user,omitempty"`
}

// OSLoginSSHPublicKeyStatus defines the observed state of OSLoginSSHPublicKey.
type OSLoginSSHPublicKeyStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
