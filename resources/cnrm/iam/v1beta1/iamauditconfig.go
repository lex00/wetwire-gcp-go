package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// IAMAuditConfig is the schema for the IAM audit logging API.
type IAMAuditConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IAMAuditConfigSpec   `json:"spec,omitempty"`
	Status IAMAuditConfigStatus `json:"status,omitempty"`
}

// IAMAuditConfigSpec defines the desired state of IAMAuditConfig.
type IAMAuditConfigSpec struct {
	// Required. The configuration for logging of each type of permission.
	AuditLogConfigs []map[string]interface{} `json:"auditLogConfigs,omitempty" yaml:"auditLogConfigs,omitempty"`
	// Immutable. Required. The GCP resource to set the IAMAuditConfig on (e.g. project).
	ResourceRef map[string]interface{} `json:"resourceRef,omitempty" yaml:"resourceRef,omitempty"`
	// Immutable. Required. The service for which to enable Data Access audit logs. The special value 'allServices' covers all services. Note that if there are audit configs covering both 'allServices' and a specific service, then the union of the two audit configs is used for that service: the 'logTypes' specified in each 'auditLogConfig' are enabled, and the 'exemptedMembers' in each 'auditLogConfig' are exempted.
	Service string `json:"service,omitempty" yaml:"service,omitempty"`
}

// IAMAuditConfigStatus defines the observed state of IAMAuditConfig.
type IAMAuditConfigStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
