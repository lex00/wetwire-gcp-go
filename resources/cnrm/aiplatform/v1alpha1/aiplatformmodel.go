package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// AIPlatformModel is the Schema for the AIPlatformModel API
type AIPlatformModel struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AIPlatformModelSpec   `json:"spec,omitempty"`
	Status AIPlatformModelStatus `json:"status,omitempty"`
}

// AIPlatformModelSpec defines the desired state of AIPlatformModel.
type AIPlatformModelSpec struct {
	// Immutable. The path to the directory containing the Model artifact and any of its supporting files. Not required for AutoML Models.
	ArtifactURI string `json:"artifactURI,omitempty" yaml:"artifactURI,omitempty"`
	// Optional. User input field to specify the base model source. Currently it only supports specifying the Model Garden models and Genie models.
	BaseModelSource map[string]interface{} `json:"baseModelSource,omitempty" yaml:"baseModelSource,omitempty"`
	// Input only. The specification of the container that is to be used when deploying this Model. The specification is ingested upon [ModelService.UploadModel][google.cloud.aiplatform.v1.ModelService.UploadModel], and all binaries it contains are copied and stored internally by Vertex AI. Not required for AutoML Models.
	ContainerSpec map[string]interface{} `json:"containerSpec,omitempty" yaml:"containerSpec,omitempty"`
	// Stats of data used for training or evaluating the Model.
	// Only populated when the Model is trained by a TrainingPipeline with
	// [data_input_config][google.cloud.aiplatform.v1.TrainingPipeline.input_data_config].
	DataStats map[string]interface{} `json:"dataStats,omitempty" yaml:"dataStats,omitempty"`
	// The description of the Model.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Required. The display name of the Model. The name can be up to 128 characters long and can consist of any UTF-8 characters.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// Customer-managed encryption key spec for a Model. If set, this Model and all sub-resources of this Model will be secured by this key.
	EncryptionSpec map[string]interface{} `json:"encryptionSpec,omitempty" yaml:"encryptionSpec,omitempty"`
	// The default explanation specification for this Model.
	// The Model can be used for
	// [requesting
	// explanation][google.cloud.aiplatform.v1.PredictionService.Explain] after
	// being [deployed][google.cloud.aiplatform.v1.EndpointService.DeployModel] if
	// it is populated. The Model can be used for [batch
	// explanation][google.cloud.aiplatform.v1.BatchPredictionJob.generate_explanation]
	// if it is populated.
	// All fields of the explanation_spec can be overridden by
	// [explanation_spec][google.cloud.aiplatform.v1.DeployedModel.explanation_spec]
	// of
	// [DeployModelRequest.deployed_model][google.cloud.aiplatform.v1.DeployModelRequest.deployed_model],
	// or
	// [explanation_spec][google.cloud.aiplatform.v1.BatchPredictionJob.explanation_spec]
	// of [BatchPredictionJob][google.cloud.aiplatform.v1.BatchPredictionJob].
	// If the default explanation specification is not set for this Model, this
	// Model can still be used for
	// [requesting
	// explanation][google.cloud.aiplatform.v1.PredictionService.Explain] by
	// setting
	// [explanation_spec][google.cloud.aiplatform.v1.DeployedModel.explanation_spec]
	// of
	// [DeployModelRequest.deployed_model][google.cloud.aiplatform.v1.DeployModelRequest.deployed_model]
	// and for [batch
	// explanation][google.cloud.aiplatform.v1.BatchPredictionJob.generate_explanation]
	// by setting
	// [explanation_spec][google.cloud.aiplatform.v1.BatchPredictionJob.explanation_spec]
	// of [BatchPredictionJob][google.cloud.aiplatform.v1.BatchPredictionJob].
	ExplanationSpec map[string]interface{} `json:"explanationSpec,omitempty" yaml:"explanationSpec,omitempty"`
	// The labels with user-defined metadata to organize your Models.
	// Label keys and values can be no longer than 64 characters
	// (Unicode codepoints), can only contain lowercase letters, numeric
	// characters, underscores and dashes. International characters are allowed.
	// See https://goo.gl/xmQnxf for more information and examples of labels.
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	// Immutable. The location where the model should reside.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Immutable. An additional information about the Model; the schema of the metadata can be found in [metadata_schema][google.cloud.aiplatform.v1.Model.metadata_schema_uri]. Unset if the Model does not have any additional information.
	Metadata map[string]interface{} `json:"metadata,omitempty" yaml:"metadata,omitempty"`
	// Immutable. Points to a YAML file stored on Google Cloud Storage describing additional information about the Model, that is specific to it. Unset if the Model does not have any additional information. The schema is defined as an OpenAPI 3.0.2 [Schema Object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.2.md#schemaObject). AutoML Models always have this field populated by Vertex AI, if no additional metadata is needed, this field is set to an empty string. Note: The URI given on output will be immutable and probably different, including the URI scheme, than the one given on input. The output URI will point to a location where the user only has a read access.
	MetadataSchemaURI string `json:"metadataSchemaURI,omitempty" yaml:"metadataSchemaURI,omitempty"`
	// Optional. This field is populated if the model is produced by a pipeline job.
	PipelineJob string `json:"pipelineJob,omitempty" yaml:"pipelineJob,omitempty"`
	// The schemata that describe formats of the Model's predictions and explanations as given and returned via [PredictionService.Predict][google.cloud.aiplatform.v1.PredictionService.Predict] and [PredictionService.Explain][google.cloud.aiplatform.v1.PredictionService.Explain].
	PredictSchemata map[string]interface{} `json:"predictSchemata,omitempty" yaml:"predictSchemata,omitempty"`
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// The AIPlatformModel name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// User provided version aliases so that a model version can be referenced via alias (i.e. `projects/{project}/locations/{location}/models/{model_id}@{version_alias}` instead of auto-generated version id (i.e. `projects/{project}/locations/{location}/models/{model_id}@{version_id})`. The format is [a-z][a-zA-Z0-9-]{0,126}[a-z0-9] to distinguish from version_id. A default version alias will be created for the first version of the model, and there must be exactly one default version alias for a model.
	VersionAliases []string `json:"versionAliases,omitempty" yaml:"versionAliases,omitempty"`
	// The description of this version.
	VersionDescription string `json:"versionDescription,omitempty" yaml:"versionDescription,omitempty"`
}

// AIPlatformModelStatus defines the observed state of AIPlatformModel.
type AIPlatformModelStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
