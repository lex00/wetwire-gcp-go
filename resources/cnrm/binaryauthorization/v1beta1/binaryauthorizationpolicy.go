package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// BinaryAuthorizationPolicy represents a binaryauthorization.cnrm.cloud.google.com BinaryAuthorizationPolicy resource.
type BinaryAuthorizationPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BinaryAuthorizationPolicySpec   `json:"spec,omitempty"`
	Status BinaryAuthorizationPolicyStatus `json:"status,omitempty"`
}

// BinaryAuthorizationPolicySpec defines the desired state of BinaryAuthorizationPolicy.
type BinaryAuthorizationPolicySpec struct {
	// Optional. Admission policy allowlisting. A matching admission request will always be permitted. This feature is typically used to exclude Google or third-party infrastructure images from Binary Authorization policies.
	AdmissionWhitelistPatterns []map[string]interface{} `json:"admissionWhitelistPatterns,omitempty" yaml:"admissionWhitelistPatterns,omitempty"`
	// Optional. Per-cluster admission rules. Cluster spec format: location.clusterId. There can be at most one admission rule per cluster spec. A location is either a compute zone (e.g. us-central1-a) or a region (e.g. us-central1). For clusterId syntax restrictions see https://cloud.google.com/container-engine/reference/rest/v1/projects.zones.clusters.
	ClusterAdmissionRules map[string]map[string]interface{} `json:"clusterAdmissionRules,omitempty" yaml:"clusterAdmissionRules,omitempty"`
	// Required. Default admission rule for a cluster without a per-cluster, per-kubernetes-service-account, or per-istio-service-identity admission rule.
	DefaultAdmissionRule map[string]interface{} `json:"defaultAdmissionRule,omitempty" yaml:"defaultAdmissionRule,omitempty"`
	// Optional. A descriptive comment.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Optional. Controls the evaluation of a Google-maintained global admission policy for common system-level images. Images not covered by the global policy will be subject to the project admission policy. This setting has no effect when specified inside a global admission policy. Possible values: GLOBAL_POLICY_EVALUATION_MODE_UNSPECIFIED, ENABLE, DISABLE
	GlobalPolicyEvaluationMode string `json:"globalPolicyEvaluationMode,omitempty" yaml:"globalPolicyEvaluationMode,omitempty"`
	// Optional. Per-istio-service-identity admission rules. Istio service identity spec format: spiffe:///ns//sa/ or /ns//sa/ e.g. spiffe://example.com/ns/test-ns/sa/default
	IstioServiceIdentityAdmissionRules map[string]map[string]interface{} `json:"istioServiceIdentityAdmissionRules,omitempty" yaml:"istioServiceIdentityAdmissionRules,omitempty"`
	// Optional. Per-kubernetes-namespace admission rules. K8s namespace spec format: [a-z.-]+, e.g. 'some-namespace'
	KubernetesNamespaceAdmissionRules map[string]map[string]interface{} `json:"kubernetesNamespaceAdmissionRules,omitempty" yaml:"kubernetesNamespaceAdmissionRules,omitempty"`
	// Optional. Per-kubernetes-service-account admission rules. Service account spec format: namespace:serviceaccount. e.g. 'test-ns:default'
	KubernetesServiceAccountAdmissionRules map[string]map[string]interface{} `json:"kubernetesServiceAccountAdmissionRules,omitempty" yaml:"kubernetesServiceAccountAdmissionRules,omitempty"`
	// Immutable. The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
}

// BinaryAuthorizationPolicyStatus defines the observed state of BinaryAuthorizationPolicy.
type BinaryAuthorizationPolicyStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
