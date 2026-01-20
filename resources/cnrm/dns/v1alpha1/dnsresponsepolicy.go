package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DNSResponsePolicy represents a dns.cnrm.cloud.google.com DNSResponsePolicy resource.
type DNSResponsePolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DNSResponsePolicySpec   `json:"spec,omitempty"`
	Status DNSResponsePolicyStatus `json:"status,omitempty"`
}

// DNSResponsePolicySpec defines the desired state of DNSResponsePolicy.
type DNSResponsePolicySpec struct {
	// The description of the response policy, such as 'My new response policy'.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// The list of Google Kubernetes Engine clusters that can see this zone.
	GkeClusters []map[string]interface{} `json:"gkeClusters,omitempty" yaml:"gkeClusters,omitempty"`
	// The list of network names specifying networks to which this policy is applied.
	Networks []map[string]interface{} `json:"networks,omitempty" yaml:"networks,omitempty"`
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The responsePolicyName of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// DNSResponsePolicyStatus defines the observed state of DNSResponsePolicy.
type DNSResponsePolicyStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
