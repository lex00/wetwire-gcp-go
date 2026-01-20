package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ComputeInterconnectAttachment represents a compute.cnrm.cloud.google.com ComputeInterconnectAttachment resource.
type ComputeInterconnectAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeInterconnectAttachmentSpec   `json:"spec,omitempty"`
	Status ComputeInterconnectAttachmentStatus `json:"status,omitempty"`
}

// ComputeInterconnectAttachmentSpec defines the desired state of ComputeInterconnectAttachment.
type ComputeInterconnectAttachmentSpec struct {
	// Whether the VLAN attachment is enabled or disabled.  When using
	// PARTNER type this will Pre-Activate the interconnect attachment.
	AdminEnabled bool `json:"adminEnabled,omitempty" yaml:"adminEnabled,omitempty"`
	// Provisioned bandwidth capacity for the interconnect attachment.
	// For attachments of type DEDICATED, the user can set the bandwidth.
	// For attachments of type PARTNER, the Google Partner that is operating the interconnect must set the bandwidth.
	// Output only for PARTNER type, mutable for PARTNER_PROVIDER and DEDICATED,
	// Defaults to BPS_10G Possible values: ["BPS_50M", "BPS_100M", "BPS_200M", "BPS_300M", "BPS_400M", "BPS_500M", "BPS_1G", "BPS_2G", "BPS_5G", "BPS_10G", "BPS_20G", "BPS_50G"].
	Bandwidth string `json:"bandwidth,omitempty" yaml:"bandwidth,omitempty"`
	// Immutable. Up to 16 candidate prefixes that can be used to restrict the allocation
	// of cloudRouterIpAddress and customerRouterIpAddress for this attachment.
	// All prefixes must be within link-local address space (169.254.0.0/16)
	// and must be /29 or shorter (/28, /27, etc). Google will attempt to select
	// an unused /29 from the supplied candidate prefix(es). The request will
	// fail if all possible /29s are in use on Google's edge. If not supplied,
	// Google will randomly select an unused /29 from all of link-local space.
	CandidateSubnets []string `json:"candidateSubnets,omitempty" yaml:"candidateSubnets,omitempty"`
	// An optional description of this resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Immutable. Desired availability domain for the attachment. Only available for type
	// PARTNER, at creation time. For improved reliability, customers should
	// configure a pair of attachments with one per availability domain. The
	// selected availability domain will be provided to the Partner via the
	// pairing key so that the provisioned circuit will lie in the specified
	// domain. If not specified, the value will default to AVAILABILITY_DOMAIN_ANY.
	EdgeAvailabilityDomain string `json:"edgeAvailabilityDomain,omitempty" yaml:"edgeAvailabilityDomain,omitempty"`
	// Immutable. Indicates the user-supplied encryption option of this interconnect
	// attachment. Can only be specified at attachment creation for PARTNER or
	// DEDICATED attachments.
	// * NONE - This is the default value, which means that the VLAN attachment
	// carries unencrypted traffic. VMs are able to send traffic to, or receive
	// traffic from, such a VLAN attachment.
	// * IPSEC - The VLAN attachment carries only encrypted traffic that is
	// encrypted by an IPsec device, such as an HA VPN gateway or third-party
	// IPsec VPN. VMs cannot directly send traffic to, or receive traffic from,
	// such a VLAN attachment. To use HA VPN over Cloud Interconnect, the VLAN
	// attachment must be created with this option. Default value: "NONE" Possible values: ["NONE", "IPSEC"].
	Encryption string `json:"encryption,omitempty" yaml:"encryption,omitempty"`
	// Immutable. URL of the underlying Interconnect object that this attachment's
	// traffic will traverse through. Required if type is DEDICATED, must not
	// be set if type is PARTNER.
	Interconnect string `json:"interconnect,omitempty" yaml:"interconnect,omitempty"`
	IpsecInternalAddresses []map[string]interface{} `json:"ipsecInternalAddresses,omitempty" yaml:"ipsecInternalAddresses,omitempty"`
	// Maximum Transmission Unit (MTU), in bytes, of packets passing through
	// this interconnect attachment. Currently, only 1440 and 1500 are allowed. If not specified, the value will default to 1440.
	Mtu string `json:"mtu,omitempty" yaml:"mtu,omitempty"`
	// Region where the regional interconnect attachment resides.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// The Cloud Router to be used for dynamic routing. This router must
	// be in the same region as this ComputeInterconnectAttachment. The
	// ComputeInterconnectAttachment will automatically connect the
	// interconnect to the network & region within which the Cloud Router
	// is configured.
	RouterRef map[string]interface{} `json:"routerRef,omitempty" yaml:"routerRef,omitempty"`
	// Immutable. The type of InterconnectAttachment you wish to create. Defaults to
	// DEDICATED. Possible values: ["DEDICATED", "PARTNER", "PARTNER_PROVIDER"].
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
	// Immutable. The IEEE 802.1Q VLAN tag for this attachment, in the range 2-4094. When
	// using PARTNER type this will be managed upstream.
	VlanTag8021q int32 `json:"vlanTag8021q,omitempty" yaml:"vlanTag8021q,omitempty"`
}

// ComputeInterconnectAttachmentStatus defines the observed state of ComputeInterconnectAttachment.
type ComputeInterconnectAttachmentStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
