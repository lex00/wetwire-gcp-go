package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// BigQueryDataPolicy is the Schema for the BigQueryDataPolicy API
type BigQueryDataPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BigQueryDataPolicySpec   `json:"spec,omitempty"`
	Status BigQueryDataPolicyStatus `json:"status,omitempty"`
}

// BigQueryDataPolicySpec defines the desired state of BigQueryDataPolicy.
type BigQueryDataPolicySpec struct {
	// The data masking policy that specifies the data masking rule to use.
	DataMaskingPolicy map[string]interface{} `json:"dataMaskingPolicy,omitempty" yaml:"dataMaskingPolicy,omitempty"`
	// User-assigned (human readable) ID of the data policy that needs to be unique within a project. Used as {data_policy_id} in part of the resource name.
	DataPolicyID string `json:"dataPolicyID,omitempty" yaml:"dataPolicyID,omitempty"`
	// Type of data policy.
	DataPolicyType string `json:"dataPolicyType,omitempty" yaml:"dataPolicyType,omitempty"`
	// Required. The location of the application.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Reference to a Data Catalog Policy Tag resource.
	PolicyTagRef map[string]interface{} `json:"policyTagRef,omitempty" yaml:"policyTagRef,omitempty"`
	// Required. The host project of the application.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// The BigQueryDataPolicy name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// BigQueryDataPolicyStatus defines the observed state of BigQueryDataPolicy.
type BigQueryDataPolicyStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
