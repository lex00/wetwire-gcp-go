package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// AssuredWorkloadsWorkload is the Schema for the AssuredWorkloadsWorkload API
type AssuredWorkloadsWorkload struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AssuredWorkloadsWorkloadSpec   `json:"spec,omitempty"`
	Status AssuredWorkloadsWorkloadStatus `json:"status,omitempty"`
}

// AssuredWorkloadsWorkloadSpec defines the desired state of AssuredWorkloadsWorkload.
type AssuredWorkloadsWorkloadSpec struct {
	// Optional. The billing account used for the resources which are direct children of workload. This billing account is initially associated with the resources created as part of Workload creation. After the initial creation of these resources, the customer can change the assigned billing account. The resource name has the form `billingAccounts/{billing_account_id}`. For example, `billingAccounts/012345-567890-ABCDEF`.
	BillingAccountRef map[string]interface{} `json:"billingAccountRef,omitempty" yaml:"billingAccountRef,omitempty"`
	// Required. Immutable. Compliance Regime associated with this workload.
	ComplianceRegime string `json:"complianceRegime,omitempty" yaml:"complianceRegime,omitempty"`
	// Required. The user-assigned display name of the Workload.
	// When present it must be between 4 to 30 characters.
	// Allowed characters are: lowercase and uppercase letters, numbers,
	// hyphen, and spaces.
	// Example: My Workload
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// Optional. Indicates the sovereignty status of the given workload. Currently meant to be used by Europe/Canada customers.
	EnableSovereignControls bool `json:"enableSovereignControls,omitempty" yaml:"enableSovereignControls,omitempty"`
	// Optional. Compliance Regime associated with this workload.
	Partner string `json:"partner,omitempty" yaml:"partner,omitempty"`
	// The AssuredWorkloadsWorkload name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Input only. Resource properties that are used to customize workload resources. These properties (such as custom project id) will be used to create workload resources if possible. This field is optional.
	ResourceSettings []map[string]interface{} `json:"resourceSettings,omitempty" yaml:"resourceSettings,omitempty"`
}

// AssuredWorkloadsWorkloadStatus defines the observed state of AssuredWorkloadsWorkload.
type AssuredWorkloadsWorkloadStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
