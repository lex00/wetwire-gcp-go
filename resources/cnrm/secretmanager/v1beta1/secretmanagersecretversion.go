package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// SecretManagerSecretVersion is the Schema for the SecretManagerSecretVersion API
type SecretManagerSecretVersion struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SecretManagerSecretVersionSpec   `json:"spec,omitempty"`
	Status SecretManagerSecretVersionStatus `json:"status,omitempty"`
}

// SecretManagerSecretVersionSpec defines the desired state of SecretManagerSecretVersion.
type SecretManagerSecretVersionSpec struct {
	// DEPRECATED. You do not need to set this field in direct reconciler mode. Use delete-policy annotation instead. https://cloud.google.com/config-connector/docs/how-to/managing-deleting-resources#keeping_resources_after_deletion The deletion policy for the secret version. Setting 'ABANDON' allows the resource to be abandoned rather than deleted. Setting 'DISABLE' allows the resource to be disabled rather than deleted. Default is 'DELETE'. Possible values are: * DELETE * DISABLE * ABANDON.
	DeletionPolicy string `json:"deletionPolicy,omitempty" yaml:"deletionPolicy,omitempty"`
	// Should enable or disable the current SecretVersion. - Enabled version can be accessed and described. - Disabled version cannot be accessed, but the secret's contents still exist
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	// DEPRECATED. You do not need to set this field in direct reconciler mode.
	IsSecretDataBase64 bool `json:"isSecretDataBase64,omitempty" yaml:"isSecretDataBase64,omitempty"`
	// The SecretVersion number. If given, Config Connector acquires the resource from the Secret Manager service. If not given, Config Connector adds a new secret version to the GCP service, and you can find out the version number from `status.observedState.version`
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// The actual secret data. Config Connector supports secret data stored in Kubernetes secret or plain data (base64)
	SecretData map[string]interface{} `json:"secretData,omitempty" yaml:"secretData,omitempty"`
	// The resource name of the [Secret][google.cloud.secretmanager.v1.Secret] to create a [SecretVersion][google.cloud.secretmanager.v1.SecretVersion] for.
	SecretRef map[string]interface{} `json:"secretRef,omitempty" yaml:"secretRef,omitempty"`
}

// SecretManagerSecretVersionStatus defines the observed state of SecretManagerSecretVersion.
type SecretManagerSecretVersionStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
