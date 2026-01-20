package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// AccessContextManagerServicePerimeterResource represents a accesscontextmanager.cnrm.cloud.google.com AccessContextManagerServicePerimeterResource resource.
type AccessContextManagerServicePerimeterResource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AccessContextManagerServicePerimeterResourceSpec   `json:"spec,omitempty"`
	Status AccessContextManagerServicePerimeterResourceStatus `json:"status,omitempty"`
}

// AccessContextManagerServicePerimeterResourceSpec defines the desired state of AccessContextManagerServicePerimeterResource.
type AccessContextManagerServicePerimeterResourceSpec struct {
	// Only the `external` field is supported to configure the reference.
	// The name of the Service Perimeter to add this resource to.
	// Referencing a resource name leads to recursive reference and Config Connector does not support the feature for now.
	PerimeterNameRef map[string]interface{} `json:"perimeterNameRef,omitempty" yaml:"perimeterNameRef,omitempty"`
	// A GCP resource that is inside of the service perimeter.
	ResourceRef map[string]interface{} `json:"resourceRef,omitempty" yaml:"resourceRef,omitempty"`
}

// AccessContextManagerServicePerimeterResourceStatus defines the observed state of AccessContextManagerServicePerimeterResource.
type AccessContextManagerServicePerimeterResourceStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
