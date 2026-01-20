package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CloudBuildTrigger represents a cloudbuild.cnrm.cloud.google.com CloudBuildTrigger resource.
type CloudBuildTrigger struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CloudBuildTriggerSpec   `json:"spec,omitempty"`
	Status CloudBuildTriggerStatus `json:"status,omitempty"`
}

// CloudBuildTriggerSpec defines the desired state of CloudBuildTrigger.
type CloudBuildTriggerSpec struct {
	// Configuration for manual approval to start a build invocation of this BuildTrigger.
	// Builds created by this trigger will require approval before they execute.
	// Any user with a Cloud Build Approver role for the project can approve a build.
	ApprovalConfig map[string]interface{} `json:"approvalConfig,omitempty" yaml:"approvalConfig,omitempty"`
	// BitbucketServerTriggerConfig describes the configuration of a trigger that creates a build whenever a Bitbucket Server event is received.
	BitbucketServerTriggerConfig map[string]interface{} `json:"bitbucketServerTriggerConfig,omitempty" yaml:"bitbucketServerTriggerConfig,omitempty"`
	// Contents of the build template. Either a filename or build template must be provided.
	Build map[string]interface{} `json:"build,omitempty" yaml:"build,omitempty"`
	// Human-readable description of the trigger.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Whether the trigger is disabled or not. If true, the trigger will never result in a build.
	Disabled bool `json:"disabled,omitempty" yaml:"disabled,omitempty"`
	// Path, from the source root, to a file whose contents is used for the template.
	// Either a filename or build template must be provided. Set this only when using trigger_template or github.
	// When using Pub/Sub, Webhook or Manual set the file name using git_file_source instead.
	Filename string `json:"filename,omitempty" yaml:"filename,omitempty"`
	// A Common Expression Language string. Used only with Pub/Sub and Webhook.
	Filter string `json:"filter,omitempty" yaml:"filter,omitempty"`
	// The file source describing the local or remote Build template.
	GitFileSource map[string]interface{} `json:"gitFileSource,omitempty" yaml:"gitFileSource,omitempty"`
	// Describes the configuration of a trigger that creates a build whenever a GitHub event is received.
	// One of 'trigger_template', 'github', 'pubsub_config' or 'webhook_config' must be provided.
	Github map[string]interface{} `json:"github,omitempty" yaml:"github,omitempty"`
	// ignoredFiles and includedFiles are file glob matches using https://golang.org/pkg/path/filepath/#Match
	// extended with support for '**'.
	// If ignoredFiles and changed files are both empty, then they are not
	// used to determine whether or not to trigger a build.
	// If ignoredFiles is not empty, then we ignore any files that match any
	// of the ignored_file globs. If the change has no files that are outside
	// of the ignoredFiles globs, then we do not trigger a build.
	IgnoredFiles []string `json:"ignoredFiles,omitempty" yaml:"ignoredFiles,omitempty"`
	// Build logs will be sent back to GitHub as part of the checkrun
	// result.  Values can be INCLUDE_BUILD_LOGS_UNSPECIFIED or
	// INCLUDE_BUILD_LOGS_WITH_STATUS Possible values: ["INCLUDE_BUILD_LOGS_UNSPECIFIED", "INCLUDE_BUILD_LOGS_WITH_STATUS"].
	IncludeBuildLogs string `json:"includeBuildLogs,omitempty" yaml:"includeBuildLogs,omitempty"`
	// ignoredFiles and includedFiles are file glob matches using https://golang.org/pkg/path/filepath/#Match
	// extended with support for '**'.
	// If any of the files altered in the commit pass the ignoredFiles filter
	// and includedFiles is empty, then as far as this filter is concerned, we
	// should trigger the build.
	// If any of the files altered in the commit pass the ignoredFiles filter
	// and includedFiles is not empty, then we make sure that at least one of
	// those files matches a includedFiles glob. If not, then we do not trigger
	// a build.
	IncludedFiles []string `json:"includedFiles,omitempty" yaml:"includedFiles,omitempty"`
	// Immutable. The location of the Cloud Build trigger. If not specified, "global" is used. More info: cloud.google.com/build/docs/locations.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// PubsubConfig describes the configuration of a trigger that creates
	// a build whenever a Pub/Sub message is published.
	// One of 'trigger_template', 'github', 'pubsub_config' 'webhook_config' or 'source_to_build' must be provided.
	PubsubConfig map[string]interface{} `json:"pubsubConfig,omitempty" yaml:"pubsubConfig,omitempty"`
	// The configuration of a trigger that creates a build whenever an event from Repo API is received.
	RepositoryEventConfig map[string]interface{} `json:"repositoryEventConfig,omitempty" yaml:"repositoryEventConfig,omitempty"`
	// The service account used for all user-controlled operations including
	// triggers.patch, triggers.run, builds.create, and builds.cancel.
	// If no service account is set, then the standard Cloud Build service account
	// ([PROJECT_NUM]@system.gserviceaccount.com) will be used instead.
	// When populating via the external field, the following format is supported:
	// projects/{PROJECT_ID}/serviceAccounts/{SERVICE_ACCOUNT_EMAIL}
	ServiceAccountRef map[string]interface{} `json:"serviceAccountRef,omitempty" yaml:"serviceAccountRef,omitempty"`
	// The repo and ref of the repository from which to build.
	// This field is used only for those triggers that do not respond to SCM events.
	// Triggers that respond to such events build source at whatever commit caused the event.
	// This field is currently only used by Webhook, Pub/Sub, Manual, and Cron triggers.
	// One of 'trigger_template', 'github', 'pubsub_config' 'webhook_config' or 'source_to_build' must be provided.
	SourceToBuild map[string]interface{} `json:"sourceToBuild,omitempty" yaml:"sourceToBuild,omitempty"`
	// Substitutions data for Build resource.
	Substitutions map[string]string `json:"substitutions,omitempty" yaml:"substitutions,omitempty"`
	// Tags for annotation of a BuildTrigger.
	Tags []string `json:"tags,omitempty" yaml:"tags,omitempty"`
	// Template describing the types of source changes to trigger a build.
	// Branch and tag names in trigger templates are interpreted as regular
	// expressions. Any branch or tag change that matches that regular
	// expression will trigger a build.
	// One of 'trigger_template', 'github', 'pubsub_config', 'webhook_config' or 'source_to_build' must be provided.
	TriggerTemplate map[string]interface{} `json:"triggerTemplate,omitempty" yaml:"triggerTemplate,omitempty"`
	// WebhookConfig describes the configuration of a trigger that creates
	// a build whenever a webhook is sent to a trigger's webhook URL.
	// One of 'trigger_template', 'github', 'pubsub_config' 'webhook_config' or 'source_to_build' must be provided.
	WebhookConfig map[string]interface{} `json:"webhookConfig,omitempty" yaml:"webhookConfig,omitempty"`
}

// CloudBuildTriggerStatus defines the observed state of CloudBuildTrigger.
type CloudBuildTriggerStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
