package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// BillingAccount is the Schema for the BillingAccount API
type BillingAccount struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BillingAccountSpec   `json:"spec,omitempty"`
	Status BillingAccountStatus `json:"status,omitempty"`
}

// BillingAccountSpec defines the desired state of BillingAccount.
type BillingAccountSpec struct {
	// Optional. The currency in which the billing account is billed and charged,
	// represented as an ISO 4217 code such as `USD`.
	// Billing account currency is determined at the time of billing account
	// creation and cannot be updated subsequently, so this field should not be
	// set on update requests. In addition, a subaccount always matches the
	// currency of its parent billing account, so this field should not be set on
	// subaccounts. Clients can read this field to determine the
	// currency of an existing billing account.
	CurrencyCode string `json:"currencyCode,omitempty" yaml:"currencyCode,omitempty"`
	// The display name given to the billing account, such as `My Billing Account`. This name is displayed in the Google Cloud Console.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// Optional. The billing account's parent resource.
	ParentRef map[string]interface{} `json:"parentRef,omitempty" yaml:"parentRef,omitempty"`
	// The BillingAccount name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// BillingAccountStatus defines the observed state of BillingAccount.
type BillingAccountStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
