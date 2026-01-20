package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CloudAssetFolderFeed represents a cloudasset.cnrm.cloud.google.com CloudAssetFolderFeed resource.
type CloudAssetFolderFeed struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CloudAssetFolderFeedSpec   `json:"spec,omitempty"`
	Status CloudAssetFolderFeedStatus `json:"status,omitempty"`
}

// CloudAssetFolderFeedSpec defines the desired state of CloudAssetFolderFeed.
type CloudAssetFolderFeedSpec struct {
	// A list of the full names of the assets to receive updates. You must specify either or both of
	// assetNames and assetTypes. Only asset updates matching specified assetNames and assetTypes are
	// exported to the feed. For example: //compute.googleapis.com/projects/my_project_123/zones/zone1/instances/instance1.
	// See https://cloud.google.com/apis/design/resourceNames#fullResourceName for more info.
	AssetNames []string `json:"assetNames,omitempty" yaml:"assetNames,omitempty"`
	// A list of types of the assets to receive updates. You must specify either or both of assetNames
	// and assetTypes. Only asset updates matching specified assetNames and assetTypes are exported to
	// the feed. For example: "compute.googleapis.com/Disk"
	// See https://cloud.google.com/asset-inventory/docs/supported-asset-types for a list of all
	// supported asset types.
	AssetTypes []string `json:"assetTypes,omitempty" yaml:"assetTypes,omitempty"`
	// Immutable. The project whose identity will be used when sending messages to the
	// destination pubsub topic. It also specifies the project for API
	// enablement check, quota, and billing.
	BillingProject string `json:"billingProject,omitempty" yaml:"billingProject,omitempty"`
	// A condition which determines whether an asset update should be published. If specified, an asset
	// will be returned only when the expression evaluates to true. When set, expression field
	// must be a valid CEL expression on a TemporalAsset with name temporal_asset. Example: a Feed with
	// expression "temporal_asset.deleted == true" will only publish Asset deletions. Other fields of
	// condition are optional.
	Condition map[string]interface{} `json:"condition,omitempty" yaml:"condition,omitempty"`
	// Asset content type. If not specified, no content but the asset name and type will be returned. Possible values: ["CONTENT_TYPE_UNSPECIFIED", "RESOURCE", "IAM_POLICY", "ORG_POLICY", "OS_INVENTORY", "ACCESS_POLICY"].
	ContentType string `json:"contentType,omitempty" yaml:"contentType,omitempty"`
	// Immutable. This is the client-assigned asset feed identifier and it needs to be unique under a specific parent.
	FeedID string `json:"feedId,omitempty" yaml:"feedId,omitempty"`
	// Output configuration for asset feed destination.
	FeedOutputConfig map[string]interface{} `json:"feedOutputConfig,omitempty" yaml:"feedOutputConfig,omitempty"`
	// Immutable. The folder this feed should be created in.
	Folder string `json:"folder,omitempty" yaml:"folder,omitempty"`
	// The folder that this resource belongs to.
	FolderRef map[string]interface{} `json:"folderRef,omitempty" yaml:"folderRef,omitempty"`
	// Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// CloudAssetFolderFeedStatus defines the observed state of CloudAssetFolderFeed.
type CloudAssetFolderFeedStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
