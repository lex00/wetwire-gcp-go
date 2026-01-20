package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DNSManagedZone represents a dns.cnrm.cloud.google.com DNSManagedZone resource.
type DNSManagedZone struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DNSManagedZoneSpec   `json:"spec,omitempty"`
	Status DNSManagedZoneStatus `json:"status,omitempty"`
}

// DNSManagedZoneSpec defines the desired state of DNSManagedZone.
type DNSManagedZoneSpec struct {
	// Cloud logging configuration.
	CloudLoggingConfig map[string]interface{} `json:"cloudLoggingConfig,omitempty" yaml:"cloudLoggingConfig,omitempty"`
	// A textual description field. Defaults to 'Managed by Config Connector'.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Immutable. The DNS name of this managed zone, for instance "example.com.".
	DNSName string `json:"dnsName,omitempty" yaml:"dnsName,omitempty"`
	// DNSSEC configuration.
	DnssecConfig map[string]interface{} `json:"dnssecConfig,omitempty" yaml:"dnssecConfig,omitempty"`
	// The presence for this field indicates that outbound forwarding is enabled
	// for this zone. The value of this field contains the set of destinations
	// to forward to.
	ForwardingConfig map[string]interface{} `json:"forwardingConfig,omitempty" yaml:"forwardingConfig,omitempty"`
	// The presence of this field indicates that DNS Peering is enabled for this
	// zone. The value of this field contains the network to peer with.
	PeeringConfig map[string]interface{} `json:"peeringConfig,omitempty" yaml:"peeringConfig,omitempty"`
	// For privately visible zones, the set of Virtual Private Cloud
	// resources that the zone is visible from. At least one of 'gke_clusters' or 'networks' must be specified.
	PrivateVisibilityConfig map[string]interface{} `json:"privateVisibilityConfig,omitempty" yaml:"privateVisibilityConfig,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Immutable. Specifies if this is a managed reverse lookup zone. If true, Cloud DNS will resolve reverse
	// lookup queries using automatically configured records for VPC resources. This only applies
	// to networks listed under 'private_visibility_config'.
	ReverseLookup bool `json:"reverseLookup,omitempty" yaml:"reverseLookup,omitempty"`
	// Immutable. The presence of this field indicates that this zone is backed by Service Directory. The value of this field contains information related to the namespace associated with the zone.
	ServiceDirectoryConfig map[string]interface{} `json:"serviceDirectoryConfig,omitempty" yaml:"serviceDirectoryConfig,omitempty"`
	// Immutable. The zone's visibility: public zones are exposed to the Internet,
	// while private zones are visible only to Virtual Private Cloud resources. Default value: "public" Possible values: ["private", "public"].
	Visibility string `json:"visibility,omitempty" yaml:"visibility,omitempty"`
}

// DNSManagedZoneStatus defines the observed state of DNSManagedZone.
type DNSManagedZoneStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
