package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// SecretManagerSecret is the Schema for the SecretManagerSecret API
type SecretManagerSecret struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SecretManagerSecretSpec   `json:"spec,omitempty"`
	Status SecretManagerSecretStatus `json:"status,omitempty"`
}

// SecretManagerSecretSpec defines the desired state of SecretManagerSecret.
type SecretManagerSecretSpec struct {
	// Optional. Custom metadata about the secret.
	// Annotations are distinct from various forms of labels.
	// Annotations exist to allow client tools to store their own state
	// information without requiring a database.
	// Annotation keys must be between 1 and 63 characters long, have a UTF-8
	// encoding of maximum 128 bytes, begin and end with an alphanumeric character
	// ([a-z0-9A-Z]), and may have dashes (-), underscores (_), dots (.), and
	// alphanumerics in between these symbols.
	// The total size of annotation keys and values must be less than 16KiB.
	Annotations map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	// Optional. Timestamp in UTC when the [Secret][google.cloud.secretmanager.v1.Secret] is scheduled to expire. This is always provided on output, regardless of what was sent on input.
	ExpireTime string `json:"expireTime,omitempty" yaml:"expireTime,omitempty"`
	// The labels assigned to this Secret.
	// Label keys must be between 1 and 63 characters long, have a UTF-8 encoding
	// of maximum 128 bytes, and must conform to the following PCRE regular
	// expression: `[\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}`
	// Label values must be between 0 and 63 characters long, have a UTF-8
	// encoding of maximum 128 bytes, and must conform to the following PCRE
	// regular expression: `[\p{Ll}\p{Lo}\p{N}_-]{0,63}`
	// No more than 64 labels can be assigned to a given resource.
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	// Optional. Immutable. The replication policy of the secret data attached to
	// the [Secret][google.cloud.secretmanager.v1.Secret].
	// The replication policy cannot be changed after the Secret has been created.
	Replication map[string]interface{} `json:"replication,omitempty" yaml:"replication,omitempty"`
	// The SecretManagerSecret name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Optional. Rotation policy attached to the [Secret][google.cloud.secretmanager.v1.Secret]. May be excluded if there is no rotation policy.
	Rotation map[string]interface{} `json:"rotation,omitempty" yaml:"rotation,omitempty"`
	// Optional. A list of up to 10 Pub/Sub topics to which messages are published when control plane operations are called on the secret or its versions.
	Topics []map[string]interface{} `json:"topics,omitempty" yaml:"topics,omitempty"`
	// Input only. A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.5s".
	Ttl string `json:"ttl,omitempty" yaml:"ttl,omitempty"`
	// Optional. Mapping from version alias to version name.
	// A version alias is a string with a maximum length of 63 characters and can
	// contain uppercase and lowercase letters, numerals, and the hyphen (`-`)
	// and underscore ('_') characters. An alias string must start with a
	// letter and cannot be the string 'latest' or 'NEW'.
	// No more than 50 aliases can be assigned to a given secret.
	// Version-Alias pairs will be viewable via GetSecret and modifiable via
	// UpdateSecret. Access by alias is only be supported on
	// GetSecretVersion and AccessSecretVersion.
	VersionAliases map[string]string `json:"versionAliases,omitempty" yaml:"versionAliases,omitempty"`
}

// SecretManagerSecretStatus defines the observed state of SecretManagerSecret.
type SecretManagerSecretStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
