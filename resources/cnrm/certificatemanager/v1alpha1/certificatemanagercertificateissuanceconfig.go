package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CertificateManagerCertificateIssuanceConfig is the Schema for the CertificateManagerCertificateIssuanceConfig API
type CertificateManagerCertificateIssuanceConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CertificateManagerCertificateIssuanceConfigSpec   `json:"spec,omitempty"`
	Status CertificateManagerCertificateIssuanceConfigStatus `json:"status,omitempty"`
}

// CertificateManagerCertificateIssuanceConfigSpec defines the desired state of CertificateManagerCertificateIssuanceConfig.
type CertificateManagerCertificateIssuanceConfigSpec struct {
	// Required. The CA that issues the workload certificate. It includes the CA address, type, authentication to CA service, etc.
	CertificateAuthorityConfigRef map[string]interface{} `json:"certificateAuthorityConfigRef,omitempty" yaml:"certificateAuthorityConfigRef,omitempty"`
	// One or more paragraphs of text description of a CertificateIssuanceConfig.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Required. The key algorithm to use when generating the private key.
	KeyAlgorithm string `json:"keyAlgorithm,omitempty" yaml:"keyAlgorithm,omitempty"`
	// Required. Workload certificate lifetime requested.
	Lifetime string `json:"lifetime,omitempty" yaml:"lifetime,omitempty"`
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// The CertificateManagerCertificateIssuanceConfig name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Required. Specifies the percentage of elapsed time of the certificate lifetime to wait before renewing the certificate. Must be a number between 1-99, inclusive.
	RotationWindowPercentage int32 `json:"rotationWindowPercentage,omitempty" yaml:"rotationWindowPercentage,omitempty"`
}

// CertificateManagerCertificateIssuanceConfigStatus defines the observed state of CertificateManagerCertificateIssuanceConfig.
type CertificateManagerCertificateIssuanceConfigStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
