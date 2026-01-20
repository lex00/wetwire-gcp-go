package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DNSRecordSet represents a dns.cnrm.cloud.google.com DNSRecordSet resource.
type DNSRecordSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DNSRecordSetSpec   `json:"spec,omitempty"`
	Status DNSRecordSetStatus `json:"status,omitempty"`
}

// DNSRecordSetSpec defines the desired state of DNSRecordSet.
type DNSRecordSetSpec struct {
	ManagedZoneRef map[string]interface{} `json:"managedZoneRef,omitempty" yaml:"managedZoneRef,omitempty"`
	// Immutable. The DNS name this record set will apply to.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	// The configuration for steering traffic based on query. You can specify either Weighted Round Robin(WRR) type or Geolocation(GEO) type.
	RoutingPolicy map[string]interface{} `json:"routingPolicy,omitempty" yaml:"routingPolicy,omitempty"`
	// DEPRECATED. Although this field is still available, there is limited support. We recommend that you use `spec.rrdatasRefs` instead.
	Rrdatas []string `json:"rrdatas,omitempty" yaml:"rrdatas,omitempty"`
	RrdatasRefs []map[string]interface{} `json:"rrdatasRefs,omitempty" yaml:"rrdatasRefs,omitempty"`
	// The time-to-live of this record set (seconds).
	Ttl int32 `json:"ttl,omitempty" yaml:"ttl,omitempty"`
	// The DNS record set type.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}

// DNSRecordSetStatus defines the observed state of DNSRecordSet.
type DNSRecordSetStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
