package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// RecaptchaEnterpriseKey represents a recaptchaenterprise.cnrm.cloud.google.com RecaptchaEnterpriseKey resource.
type RecaptchaEnterpriseKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RecaptchaEnterpriseKeySpec   `json:"spec,omitempty"`
	Status RecaptchaEnterpriseKeyStatus `json:"status,omitempty"`
}

// RecaptchaEnterpriseKeySpec defines the desired state of RecaptchaEnterpriseKey.
type RecaptchaEnterpriseKeySpec struct {
	// Settings for keys that can be used by Android apps.
	AndroidSettings map[string]interface{} `json:"androidSettings,omitempty" yaml:"androidSettings,omitempty"`
	// Human-readable display name of this key. Modifiable by user.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// Settings for keys that can be used by iOS apps.
	IosSettings map[string]interface{} `json:"iosSettings,omitempty" yaml:"iosSettings,omitempty"`
	// Immutable. The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Immutable. Options for user acceptance testing.
	TestingOptions map[string]interface{} `json:"testingOptions,omitempty" yaml:"testingOptions,omitempty"`
	// Immutable. Settings specific to keys that can be used for WAF (Web Application Firewall).
	WafSettings map[string]interface{} `json:"wafSettings,omitempty" yaml:"wafSettings,omitempty"`
	// Settings for keys that can be used by websites.
	WebSettings map[string]interface{} `json:"webSettings,omitempty" yaml:"webSettings,omitempty"`
}

// RecaptchaEnterpriseKeyStatus defines the observed state of RecaptchaEnterpriseKey.
type RecaptchaEnterpriseKeyStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
