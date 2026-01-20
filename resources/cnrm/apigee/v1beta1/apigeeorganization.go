package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ApigeeOrganization is the Schema for the ApigeeOrganization API
type ApigeeOrganization struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ApigeeOrganizationSpec   `json:"spec,omitempty"`
	Status ApigeeOrganizationStatus `json:"status,omitempty"`
}

// ApigeeOrganizationSpec defines the desired state of ApigeeOrganization.
type ApigeeOrganizationSpec struct {
	// Addon configurations of the Apigee organization.
	AddonsConfig map[string]interface{} `json:"addonsConfig,omitempty" yaml:"addonsConfig,omitempty"`
	// Required. DEPRECATED: This field will eventually be deprecated and replaced with a differently-named field. Primary Google Cloud region for analytics data storage. For valid values, see [Create an Apigee organization](https://cloud.google.com/apigee/docs/api-platform/get-started/create-org).
	AnalyticsRegion string `json:"analyticsRegion,omitempty" yaml:"analyticsRegion,omitempty"`
	// Compute Engine network used for Service Networking to be peered with Apigee runtime instances. See [Getting started with the Service Networking API](https://cloud.google.com/service-infrastructure/docs/service-networking/getting-started). Valid only when [RuntimeType](#RuntimeType) is set to `CLOUD`. The value must be set before the creation of a runtime instance and can be updated only when there are no runtime instances. For example: `default`. Apigee also supports shared VPC (that is, the host network project is not the same as the one that is peering with Apigee). See [Shared VPC overview](https://cloud.google.com/vpc/docs/shared-vpc). To use a shared VPC network, use the following format: `projects/{host-project-id}/{region}/networks/{network-name}`. For example: `projects/my-sharedvpc-host/global/networks/mynetwork` **Note:** Not supported for Apigee hybrid.
	AuthorizedNetworkRef map[string]interface{} `json:"authorizedNetworkRef,omitempty" yaml:"authorizedNetworkRef,omitempty"`
	// Description of the Apigee organization.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Display name for the Apigee organization. Unused, but reserved for future use.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// Required. Name of the GCP project in which to associate the Apigee organization.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Properties defined in the Apigee organization profile.
	Properties map[string]string `json:"properties,omitempty" yaml:"properties,omitempty"`
	// The ApigeeOrganization name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Cloud KMS key name used for encrypting the data that is stored and replicated across runtime instances. Update is not allowed after the organization is created. If not specified or [RuntimeType](#RuntimeType) is `TRIAL`, a Google-Managed encryption key will be used. For example: "projects/foo/locations/us/keyRings/bar/cryptoKeys/baz". **Note:** Not supported for Apigee hybrid.
	RuntimeDatabaseEncryptionKeyRef map[string]interface{} `json:"runtimeDatabaseEncryptionKeyRef,omitempty" yaml:"runtimeDatabaseEncryptionKeyRef,omitempty"`
	// Required. Runtime type of the Apigee organization based on the Apigee subscription purchased.
	RuntimeType string `json:"runtimeType,omitempty" yaml:"runtimeType,omitempty"`
}

// ApigeeOrganizationStatus defines the observed state of ApigeeOrganization.
type ApigeeOrganizationStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
