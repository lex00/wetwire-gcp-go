package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EssentialContactsContact is the Schema for the EssentialContactsContact API
type EssentialContactsContact struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EssentialContactsContactSpec   `json:"spec,omitempty"`
	Status EssentialContactsContactStatus `json:"status,omitempty"`
}

// EssentialContactsContactSpec defines the desired state of EssentialContactsContact.
type EssentialContactsContactSpec struct {
	// Required. The email address to send notifications to. The email address does not need to be a Google Account.
	Email string `json:"email,omitempty" yaml:"email,omitempty"`
	// FolderRef represents the Folder that this resource belongs to.
	FolderRef map[string]interface{} `json:"folderRef,omitempty" yaml:"folderRef,omitempty"`
	// Required. The preferred language for notifications, as a ISO 639-1 language code. See [Supported languages](https://cloud.google.com/resource-manager/docs/managing-notification-contacts#supported-languages) for a list of supported languages.
	LanguageTag string `json:"languageTag,omitempty" yaml:"languageTag,omitempty"`
	// Required. The categories of notifications that the contact will receive communications for.
	NotificationCategorySubscriptions []string `json:"notificationCategorySubscriptions,omitempty" yaml:"notificationCategorySubscriptions,omitempty"`
	// OrganizationRef represents the Organization that this resource belongs to.
	OrganizationRef map[string]interface{} `json:"organizationRef,omitempty" yaml:"organizationRef,omitempty"`
	// The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// The EssentialContactsContact name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// EssentialContactsContactStatus defines the observed state of EssentialContactsContact.
type EssentialContactsContactStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
