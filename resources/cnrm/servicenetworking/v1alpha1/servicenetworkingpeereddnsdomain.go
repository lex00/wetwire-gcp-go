package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ServiceNetworkingPeeredDNSDomain is the Schema for the ServiceNetworkingPeeredDNSDomain API
type ServiceNetworkingPeeredDNSDomain struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ServiceNetworkingPeeredDNSDomainSpec   `json:"spec,omitempty"`
	Status ServiceNetworkingPeeredDNSDomainStatus `json:"status,omitempty"`
}

// ServiceNetworkingPeeredDNSDomainSpec defines the desired state of ServiceNetworkingPeeredDNSDomain.
type ServiceNetworkingPeeredDNSDomainSpec struct {
	// The DNS domain name suffix e.g. `example.com.`. Cloud DNS requires that a DNS suffix ends with a trailing dot.
	DNSSuffix string `json:"dnsSuffix,omitempty" yaml:"dnsSuffix,omitempty"`
	// The network that this resource belongs to.
	NetworkRef map[string]interface{} `json:"networkRef,omitempty" yaml:"networkRef,omitempty"`
	// The ServiceNetworkingPeeredDNSDomain name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// ServiceNetworkingPeeredDNSDomainStatus defines the observed state of ServiceNetworkingPeeredDNSDomain.
type ServiceNetworkingPeeredDNSDomainStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
