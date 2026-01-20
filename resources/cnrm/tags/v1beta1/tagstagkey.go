package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// TagsTagKey is the Schema for the TagsTagKey API
type TagsTagKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TagsTagKeySpec   `json:"spec,omitempty"`
	Status TagsTagKeyStatus `json:"status,omitempty"`
}

// TagsTagKeySpec defines the desired state of TagsTagKey.
type TagsTagKeySpec struct {
	// Optional. User-assigned description of the TagKey. Must not exceed 256
	// characters.
	// Read-write.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Immutable. The resource name of the TagKey's parent. A TagKey can be parented by an Organization or a Project. For a TagKey parented by an Organization, its parent must be in the form `organizations/{org_id}`. For a TagKey parented by a Project, its parent can be in the form `projects/{project_id}` or `projects/{project_number}`.
	Parent string `json:"parent,omitempty" yaml:"parent,omitempty"`
	// Optional. A purpose denotes that this Tag is intended for use in policies
	// of a specific policy engine, and will involve that policy engine in
	// management operations involving this Tag. A purpose does not grant a
	// policy engine exclusive rights to the Tag, and it may be referenced by
	// other policy engines.
	// A purpose cannot be changed once set.
	Purpose string `json:"purpose,omitempty" yaml:"purpose,omitempty"`
	// Optional. Purpose data corresponds to the policy system that the tag is
	// intended for. See documentation for `Purpose` for formatting of this field.
	// Purpose data cannot be changed once set.
	PurposeData map[string]string `json:"purposeData,omitempty" yaml:"purposeData,omitempty"`
	// The TagsTagKey name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Required. Immutable. The user friendly name for a TagKey. The short name
	// should be unique for TagKeys within the same tag namespace.
	// The short name must be 1-63 characters, beginning and ending with
	// an alphanumeric character ([a-z0-9A-Z]) with dashes (-), underscores (_),
	// dots (.), and alphanumerics between.
	ShortName string `json:"shortName,omitempty" yaml:"shortName,omitempty"`
}

// TagsTagKeyStatus defines the observed state of TagsTagKey.
type TagsTagKeyStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
