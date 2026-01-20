package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ContainerAnalysisOccurrence represents a containeranalysis.cnrm.cloud.google.com ContainerAnalysisOccurrence resource.
type ContainerAnalysisOccurrence struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ContainerAnalysisOccurrenceSpec   `json:"spec,omitempty"`
	Status ContainerAnalysisOccurrenceStatus `json:"status,omitempty"`
}

// ContainerAnalysisOccurrenceSpec defines the desired state of ContainerAnalysisOccurrence.
type ContainerAnalysisOccurrenceSpec struct {
	// Occurrence that represents a single "attestation". The authenticity
	// of an attestation can be verified using the attached signature.
	// If the verifier trusts the public key of the signer, then verifying
	// the signature is sufficient to establish trust. In this circumstance,
	// the authority to which this attestation is attached is primarily
	// useful for lookup (how to find this attestation if you already
	// know the authority and artifact to be verified) and intent (for
	// which authority this attestation was intended to sign.
	Attestation map[string]interface{} `json:"attestation,omitempty" yaml:"attestation,omitempty"`
	// Immutable. The analysis note associated with this occurrence, in the form of
	// projects/[PROJECT]/notes/[NOTE_ID]. This field can be used as a
	// filter in list requests.
	NoteName string `json:"noteName,omitempty" yaml:"noteName,omitempty"`
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// A description of actions that can be taken to remedy the note.
	Remediation string `json:"remediation,omitempty" yaml:"remediation,omitempty"`
	// Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Immutable. Required. Immutable. A URI that represents the resource for which
	// the occurrence applies. For example,
	// https://gcr.io/project/image@sha256:123abc for a Docker image.
	ResourceURI string `json:"resourceUri,omitempty" yaml:"resourceUri,omitempty"`
}

// ContainerAnalysisOccurrenceStatus defines the observed state of ContainerAnalysisOccurrence.
type ContainerAnalysisOccurrenceStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
