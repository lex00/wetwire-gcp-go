package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// BillingBudgetsBudget represents a billingbudgets.cnrm.cloud.google.com BillingBudgetsBudget resource.
type BillingBudgetsBudget struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BillingBudgetsBudgetSpec   `json:"spec,omitempty"`
	Status BillingBudgetsBudgetStatus `json:"status,omitempty"`
}

// BillingBudgetsBudgetSpec defines the desired state of BillingBudgetsBudget.
type BillingBudgetsBudgetSpec struct {
	// Optional. Rules to apply to notifications sent based on budget spend and thresholds.
	AllUpdatesRule map[string]interface{} `json:"allUpdatesRule,omitempty" yaml:"allUpdatesRule,omitempty"`
	// Required. Budgeted amount.
	Amount map[string]interface{} `json:"amount,omitempty" yaml:"amount,omitempty"`
	// Immutable.
	BillingAccountRef map[string]interface{} `json:"billingAccountRef,omitempty" yaml:"billingAccountRef,omitempty"`
	// Optional. Filters that define which resources are used to compute the actual spend against the budget amount, such as projects, services, and the budget's time period, as well as other filters.
	BudgetFilter map[string]interface{} `json:"budgetFilter,omitempty" yaml:"budgetFilter,omitempty"`
	// User data for display name in UI. The name must be less than or equal to 60 characters.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Optional. Rules that trigger alerts (notifications of thresholds being crossed) when spend exceeds the specified percentages of the budget.
	ThresholdRules []map[string]interface{} `json:"thresholdRules,omitempty" yaml:"thresholdRules,omitempty"`
}

// BillingBudgetsBudgetStatus defines the observed state of BillingBudgetsBudget.
type BillingBudgetsBudgetStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
