package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ApigeeInstance is the Schema for the ApigeeInstance API
type ApigeeInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ApigeeInstanceSpec   `json:"spec,omitempty"`
	Status ApigeeInstanceStatus `json:"status,omitempty"`
}

// ApigeeInstanceSpec defines the desired state of ApigeeInstance.
type ApigeeInstanceSpec struct {
	// Optional. Access logging configuration enables the access logging feature at the instance. Apigee customers can enable access logging to ship the access logs to their own project's cloud logging.
	AccessLoggingConfig map[string]interface{} `json:"accessLoggingConfig,omitempty" yaml:"accessLoggingConfig,omitempty"`
	// Optional. Customer accept list represents the list of projects (id/number) on customer side that can privately connect to the service attachment. It is an optional field which the customers can provide during the instance creation. By default, the customer project associated with the Apigee organization will be included to the list.
	ConsumerAcceptList []string `json:"consumerAcceptList,omitempty" yaml:"consumerAcceptList,omitempty"`
	// Optional. Description of the instance.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Customer Managed Encryption Key (CMEK) used for disk and volume encryption. If not specified, a Google-Managed encryption key will be used.
	DiskEncryptionKMSCryptoKeyRef map[string]interface{} `json:"diskEncryptionKMSCryptoKeyRef,omitempty" yaml:"diskEncryptionKMSCryptoKeyRef,omitempty"`
	// Optional. Display name for the instance.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// Optional. Comma-separated list of CIDR blocks of length 22 and/or 28 used to create the Apigee instance. Providing CIDR ranges is optional. You can provide just /22 or /28 or both (or neither). Ranges you provide should be freely available as part of a larger named range you have allocated to the Service Networking peering. If this parameter is not provided, Apigee automatically requests an available /22 and /28 CIDR block from Service Networking. Use the /22 CIDR block for configuring your firewall needs to allow traffic from Apigee. Input formats: `a.b.c.d/22` or `e.f.g.h/28` or `a.b.c.d/22,e.f.g.h/28`
	IPRange string `json:"ipRange,omitempty" yaml:"ipRange,omitempty"`
	// Required. Compute Engine location where the instance resides.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Reference to parent Apigee Organization.
	OrganizationRef map[string]interface{} `json:"organizationRef,omitempty" yaml:"organizationRef,omitempty"`
	// Optional. Size of the CIDR block range that will be reserved by the instance. PAID organizations support `SLASH_16` to `SLASH_20` and defaults to `SLASH_16`. Evaluation organizations support only `SLASH_23`.
	PeeringCIDRRange string `json:"peeringCIDRRange,omitempty" yaml:"peeringCIDRRange,omitempty"`
	// The ApigeeInstance name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// ApigeeInstanceStatus defines the observed state of ApigeeInstance.
type ApigeeInstanceStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
