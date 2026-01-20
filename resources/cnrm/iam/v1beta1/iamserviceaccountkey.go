package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// IAMServiceAccountKey represents a iam.cnrm.cloud.google.com IAMServiceAccountKey resource.
type IAMServiceAccountKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IAMServiceAccountKeySpec   `json:"spec,omitempty"`
	Status IAMServiceAccountKeyStatus `json:"status,omitempty"`
}

// IAMServiceAccountKeySpec defines the desired state of IAMServiceAccountKey.
type IAMServiceAccountKeySpec struct {
	// Immutable. The algorithm used to generate the key, used only on create. KEY_ALG_RSA_2048 is the default algorithm. Valid values are: "KEY_ALG_RSA_1024", "KEY_ALG_RSA_2048".
	KeyAlgorithm string `json:"keyAlgorithm,omitempty" yaml:"keyAlgorithm,omitempty"`
	// Immutable.
	PrivateKeyType string `json:"privateKeyType,omitempty" yaml:"privateKeyType,omitempty"`
	// Immutable. A field that allows clients to upload their own public key. If set, use this public key data to create a service account key for given service account. Please note, the expected format for this field is a base64 encoded X509_PEM.
	PublicKeyData string `json:"publicKeyData,omitempty" yaml:"publicKeyData,omitempty"`
	// Immutable.
	PublicKeyType string `json:"publicKeyType,omitempty" yaml:"publicKeyType,omitempty"`
	ServiceAccountRef map[string]interface{} `json:"serviceAccountRef,omitempty" yaml:"serviceAccountRef,omitempty"`
}

// IAMServiceAccountKeyStatus defines the observed state of IAMServiceAccountKey.
type IAMServiceAccountKeyStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
