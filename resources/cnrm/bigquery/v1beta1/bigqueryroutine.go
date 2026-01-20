package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// BigQueryRoutine represents a bigquery.cnrm.cloud.google.com BigQueryRoutine resource.
type BigQueryRoutine struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BigQueryRoutineSpec   `json:"spec,omitempty"`
	Status BigQueryRoutineStatus `json:"status,omitempty"`
}

// BigQueryRoutineSpec defines the desired state of BigQueryRoutine.
type BigQueryRoutineSpec struct {
	// Input/output argument of a function or a stored procedure.
	Arguments []map[string]interface{} `json:"arguments,omitempty" yaml:"arguments,omitempty"`
	// The ID of the dataset containing this routine.
	DatasetRef map[string]interface{} `json:"datasetRef,omitempty" yaml:"datasetRef,omitempty"`
	// The body of the routine. For functions, this is the expression in the AS clause.
	// If language=SQL, it is the substring inside (but excluding) the parentheses.
	DefinitionBody string `json:"definitionBody,omitempty" yaml:"definitionBody,omitempty"`
	// The description of the routine if defined.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// The determinism level of the JavaScript UDF if defined. Possible values: ["DETERMINISM_LEVEL_UNSPECIFIED", "DETERMINISTIC", "NOT_DETERMINISTIC"].
	DeterminismLevel string `json:"determinismLevel,omitempty" yaml:"determinismLevel,omitempty"`
	// Optional. If language = "JAVASCRIPT", this field stores the path of the
	// imported JAVASCRIPT libraries.
	ImportedLibraries []string `json:"importedLibraries,omitempty" yaml:"importedLibraries,omitempty"`
	// The language of the routine. Possible values: ["SQL", "JAVASCRIPT"].
	Language string `json:"language,omitempty" yaml:"language,omitempty"`
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The routineId of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Optional. Can be set only if routineType = "TABLE_VALUED_FUNCTION".
	// If absent, the return table type is inferred from definitionBody at query time in each query
	// that references this routine. If present, then the columns in the evaluated table result will
	// be cast to match the column types specificed in return table type, at query time.
	ReturnTableType string `json:"returnTableType,omitempty" yaml:"returnTableType,omitempty"`
	// A JSON schema for the return type. Optional if language = "SQL"; required otherwise.
	// If absent, the return type is inferred from definitionBody at query time in each query
	// that references this routine. If present, then the evaluated result will be cast to
	// the specified returned type at query time. ~>**NOTE**: Because this field expects a JSON
	// string, any changes to the string will create a diff, even if the JSON itself hasn't
	// changed. If the API returns a different value for the same schema, e.g. it switche
	// d the order of values or replaced STRUCT field type with RECORD field type, we currently
	// cannot suppress the recurring diff this causes. As a workaround, we recommend using
	// the schema as returned by the API.
	ReturnType string `json:"returnType,omitempty" yaml:"returnType,omitempty"`
	// Immutable. The type of routine. Possible values: ["SCALAR_FUNCTION", "PROCEDURE", "TABLE_VALUED_FUNCTION"].
	RoutineType string `json:"routineType,omitempty" yaml:"routineType,omitempty"`
}

// BigQueryRoutineStatus defines the observed state of BigQueryRoutine.
type BigQueryRoutineStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
