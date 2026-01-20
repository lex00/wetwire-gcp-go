package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// AppEngineDomainMapping represents a appengine.cnrm.cloud.google.com AppEngineDomainMapping resource.
type AppEngineDomainMapping struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AppEngineDomainMappingSpec   `json:"spec,omitempty"`
	Status AppEngineDomainMappingStatus `json:"status,omitempty"`
}

// AppEngineDomainMappingSpec defines the desired state of AppEngineDomainMapping.
type AppEngineDomainMappingSpec struct {
	// Whether the domain creation should override any existing mappings for this domain.
	// By default, overrides are rejected. Default value: "STRICT" Possible values: ["STRICT", "OVERRIDE"].
	OverrideStrategy string `json:"overrideStrategy,omitempty" yaml:"overrideStrategy,omitempty"`
	// Immutable.
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
	// Immutable. Optional. The domainName of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// SSL configuration for this domain. If unconfigured, this domain will not serve with SSL.
	SslSettings map[string]interface{} `json:"sslSettings,omitempty" yaml:"sslSettings,omitempty"`
}

// AppEngineDomainMappingStatus defines the observed state of AppEngineDomainMapping.
type AppEngineDomainMappingStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
