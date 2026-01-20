package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// APIQuotaPreference is the Schema for the APIQuotaPreference API
type APIQuotaPreference struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   APIQuotaPreferenceSpec   `json:"spec,omitempty"`
	Status APIQuotaPreferenceStatus `json:"status,omitempty"`
}

// APIQuotaPreferenceSpec defines the desired state of APIQuotaPreference.
type APIQuotaPreferenceSpec struct {
	// Input only. An email address that can be used to contact the the user, in
	// case Google Cloud needs more information to make a decision before
	// additional quota can be granted.
	// When requesting a quota increase, the email address is required.
	// When requesting a quota decrease, the email address is optional.
	// For example, the email address is optional when the
	// `QuotaConfig.preferred_value` is smaller than the
	// `QuotaDetails.reset_value`.
	ContactEmail string `json:"contactEmail,omitempty" yaml:"contactEmail,omitempty"`
	// Immutable. The dimensions that this quota preference applies to. The key of
	// the map entry is the name of a dimension, such as "region", "zone",
	// "network_id", and the value of the map entry is the dimension value.
	// If a dimension is missing from the map of dimensions, the quota preference
	// applies to all the dimension values except for those that have other quota
	// preferences configured for the specific value.
	// NOTE: QuotaPreferences can only be applied across all values of "user" and
	// "resource" dimension. Do not set values for "user" or "resource" in the
	// dimension map.
	// Example: {"provider", "Foo Inc"} where "provider" is a service specific
	// dimension.
	Dimensions map[string]string `json:"dimensions,omitempty" yaml:"dimensions,omitempty"`
	// FolderRef represents the Folder that this resource belongs to.
	FolderRef map[string]interface{} `json:"folderRef,omitempty" yaml:"folderRef,omitempty"`
	// The reason / justification for this quota preference.
	Justification string `json:"justification,omitempty" yaml:"justification,omitempty"`
	// OrganizationRef represents the Organization that this resource belongs to.
	OrganizationRef map[string]interface{} `json:"organizationRef,omitempty" yaml:"organizationRef,omitempty"`
	// The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Required. Preferred quota configuration.
	QuotaConfig map[string]interface{} `json:"quotaConfig,omitempty" yaml:"quotaConfig,omitempty"`
	// Required. The id of the quota to which the quota preference is applied. A quota name is unique in the service. Example: `CpusPerProjectPerRegion`
	QuotaID string `json:"quotaID,omitempty" yaml:"quotaID,omitempty"`
	// The APIQuotaPreference name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Required. The name of the service to which the quota preference is applied.
	Service string `json:"service,omitempty" yaml:"service,omitempty"`
}

// APIQuotaPreferenceStatus defines the observed state of APIQuotaPreference.
type APIQuotaPreferenceStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
