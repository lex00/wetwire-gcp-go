package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ContainerAnalysisNote represents a containeranalysis.cnrm.cloud.google.com ContainerAnalysisNote resource.
type ContainerAnalysisNote struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ContainerAnalysisNoteSpec   `json:"spec,omitempty"`
	Status ContainerAnalysisNoteStatus `json:"status,omitempty"`
}

// ContainerAnalysisNoteSpec defines the desired state of ContainerAnalysisNote.
type ContainerAnalysisNoteSpec struct {
	// A note describing an attestation role.
	Attestation map[string]interface{} `json:"attestation,omitempty" yaml:"attestation,omitempty"`
	// A note describing build provenance for a verifiable build.
	Build map[string]interface{} `json:"build,omitempty" yaml:"build,omitempty"`
	// A note describing something that can be deployed.
	Deployment map[string]interface{} `json:"deployment,omitempty" yaml:"deployment,omitempty"`
	// A note describing the initial analysis of a resource.
	Discovery map[string]interface{} `json:"discovery,omitempty" yaml:"discovery,omitempty"`
	// Time of expiration for this note. Empty if note does not expire.
	ExpirationTime string `json:"expirationTime,omitempty" yaml:"expirationTime,omitempty"`
	// A note describing a base image.
	Image map[string]interface{} `json:"image,omitempty" yaml:"image,omitempty"`
	// A detailed description of this note.
	LongDescription string `json:"longDescription,omitempty" yaml:"longDescription,omitempty"`
	// Required for non-Windows OS. The package this Upgrade is for.
	Package map[string]interface{} `json:"package,omitempty" yaml:"package,omitempty"`
	RelatedNoteNames []map[string]interface{} `json:"relatedNoteNames,omitempty" yaml:"relatedNoteNames,omitempty"`
	// URLs associated with this note.
	RelatedURL []map[string]interface{} `json:"relatedUrl,omitempty" yaml:"relatedUrl,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// A one sentence description of this note.
	ShortDescription string `json:"shortDescription,omitempty" yaml:"shortDescription,omitempty"`
	// A note describing a package vulnerability.
	Vulnerability map[string]interface{} `json:"vulnerability,omitempty" yaml:"vulnerability,omitempty"`
}

// ContainerAnalysisNoteStatus defines the observed state of ContainerAnalysisNote.
type ContainerAnalysisNoteStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
