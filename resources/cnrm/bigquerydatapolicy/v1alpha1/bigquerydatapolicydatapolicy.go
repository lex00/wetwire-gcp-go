package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// BigQueryDataPolicyDataPolicy represents a bigquerydatapolicy.cnrm.cloud.google.com BigQueryDataPolicyDataPolicy resource.
type BigQueryDataPolicyDataPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BigQueryDataPolicyDataPolicySpec   `json:"spec,omitempty"`
	Status BigQueryDataPolicyDataPolicyStatus `json:"status,omitempty"`
}

// BigQueryDataPolicyDataPolicySpec defines the desired state of BigQueryDataPolicyDataPolicy.
type BigQueryDataPolicyDataPolicySpec struct {
	// The data masking policy that specifies the data masking rule to use.
	DataMaskingPolicy map[string]interface{} `json:"dataMaskingPolicy,omitempty" yaml:"dataMaskingPolicy,omitempty"`
	// The enrollment level of the service. Possible values: ["COLUMN_LEVEL_SECURITY_POLICY", "DATA_MASKING_POLICY"].
	DataPolicyType string `json:"dataPolicyType,omitempty" yaml:"dataPolicyType,omitempty"`
	// Immutable. The name of the location of the data policy.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Policy tag resource name, in the format of projects/{project_number}/locations/{locationId}/taxonomies/{taxonomyId}/policyTags/{policyTag_id}.
	PolicyTag string `json:"policyTag,omitempty" yaml:"policyTag,omitempty"`
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The dataPolicyId of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// BigQueryDataPolicyDataPolicyStatus defines the observed state of BigQueryDataPolicyDataPolicy.
type BigQueryDataPolicyDataPolicyStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
