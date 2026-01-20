package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// AssetFeed is the Schema for the AssetFeed API
type AssetFeed struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AssetFeedSpec   `json:"spec,omitempty"`
	Status AssetFeedStatus `json:"status,omitempty"`
}

// AssetFeedSpec defines the desired state of AssetFeed.
type AssetFeedSpec struct {
	// A list of the full names of the assets to receive updates. You must specify either or both of asset_names and asset_types. Only asset updates matching specified asset_names or asset_types are exported to the feed. Example: `//compute.googleapis.com/projects/my_project_123/zones/zone1/instances/instance1`. For a list of the full names for supported asset types, see [Resource name format](/asset-inventory/docs/resource-name-format).
	AssetNames []string `json:"assetNames,omitempty" yaml:"assetNames,omitempty"`
	// A list of types of the assets to receive updates. You must specify either
	// or both of asset_names and asset_types. Only asset updates matching
	// specified asset_names or asset_types are exported to the feed.
	// Example: `"compute.googleapis.com/Disk"`
	// For a list of all supported asset types, see
	// [Supported asset types](/asset-inventory/docs/supported-asset-types).
	AssetTypes []string `json:"assetTypes,omitempty" yaml:"assetTypes,omitempty"`
	// A condition which determines whether an asset update should be published.
	// If specified, an asset will be returned only when the expression evaluates
	// to true.
	// When set, `expression` field in the `Expr` must be a valid [CEL expression]
	// (https://github.com/google/cel-spec) on a TemporalAsset with name
	// `temporal_asset`. Example: a Feed with expression ("temporal_asset.deleted
	// == true") will only publish Asset deletions. Other fields of `Expr` are
	// optional.
	// See our [user
	// guide](https://cloud.google.com/asset-inventory/docs/monitoring-asset-changes-with-condition)
	// for detailed instructions.
	Condition map[string]interface{} `json:"condition,omitempty" yaml:"condition,omitempty"`
	// Asset content type. If not specified, no content but the asset name and type will be returned.
	ContentType string `json:"contentType,omitempty" yaml:"contentType,omitempty"`
	// Required. Feed output configuration defining where the asset updates are published to.
	FeedOutputConfig map[string]interface{} `json:"feedOutputConfig,omitempty" yaml:"feedOutputConfig,omitempty"`
	// FolderRef represents the Folder that this resource belongs to.
	FolderRef map[string]interface{} `json:"folderRef,omitempty" yaml:"folderRef,omitempty"`
	// OrganizationRef represents the Organization that this resource belongs to.
	OrganizationRef map[string]interface{} `json:"organizationRef,omitempty" yaml:"organizationRef,omitempty"`
	// The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// A list of relationship types to output, for example: `INSTANCE_TO_INSTANCEGROUP`. This field should only be specified if content_type=RELATIONSHIP. * If specified: it outputs specified relationship updates on the [asset_names] or the [asset_types]. It returns an error if any of the [relationship_types] doesn't belong to the supported relationship types of the [asset_names] or [asset_types], or any of the [asset_names] or the [asset_types] doesn't belong to the source types of the [relationship_types]. * Otherwise: it outputs the supported relationships of the types of [asset_names] and [asset_types] or returns an error if any of the [asset_names] or the [asset_types] has no replationship support. See [Introduction to Cloud Asset Inventory](https://cloud.google.com/asset-inventory/docs/overview) for all supported asset types and relationship types.
	RelationshipTypes []string `json:"relationshipTypes,omitempty" yaml:"relationshipTypes,omitempty"`
	// The AssetFeed name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// AssetFeedStatus defines the observed state of AssetFeed.
type AssetFeedStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
