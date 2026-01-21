// Package domain provides the GCPDomain implementation for wetwire-core-go.
package domain

import (
	"fmt"
	"path/filepath"
	"sort"
	"strings"

	coredomain "github.com/lex00/wetwire-core-go/domain"
	"github.com/lex00/wetwire-gcp-go/internal/discover"
	"github.com/lex00/wetwire-gcp-go/internal/lint"
	_ "github.com/lex00/wetwire-gcp-go/internal/registry" // Register Config Connector types
	"github.com/lex00/wetwire-k8s-go/differ"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// Version is set at build time
var Version = "dev"

// Re-export core types for convenience
type (
	Context      = coredomain.Context
	BuildOpts    = coredomain.BuildOpts
	LintOpts     = coredomain.LintOpts
	InitOpts     = coredomain.InitOpts
	ValidateOpts = coredomain.ValidateOpts
	ListOpts     = coredomain.ListOpts
	GraphOpts    = coredomain.GraphOpts
	DiffOpts     = coredomain.DiffOpts
	DiffResult   = coredomain.DiffResult
	DiffEntry    = coredomain.DiffEntry
	DiffSummary  = coredomain.DiffSummary
	Result       = coredomain.Result
	Error        = coredomain.Error
)

var (
	NewResult              = coredomain.NewResult
	NewResultWithData      = coredomain.NewResultWithData
	NewErrorResult         = coredomain.NewErrorResult
	NewErrorResultMultiple = coredomain.NewErrorResultMultiple
)

// GCPDomain implements the Domain interface for GCP Config Connector manifest generation.
type GCPDomain struct{}

// Compile-time interface verification
var (
	_ coredomain.Domain       = (*GCPDomain)(nil)
	_ coredomain.ListerDomain = (*GCPDomain)(nil)
	_ coredomain.DifferDomain = (*GCPDomain)(nil)
)

// Name returns "gcp"
func (d *GCPDomain) Name() string {
	return "gcp"
}

// Version returns the current version
func (d *GCPDomain) Version() string {
	return Version
}

// Builder returns the GCP builder implementation
func (d *GCPDomain) Builder() coredomain.Builder {
	return &gcpBuilder{}
}

// Linter returns the GCP linter implementation
func (d *GCPDomain) Linter() coredomain.Linter {
	return &gcpLinter{}
}

// Initializer returns the GCP initializer implementation
func (d *GCPDomain) Initializer() coredomain.Initializer {
	return &gcpInitializer{}
}

// Validator returns the GCP validator implementation
func (d *GCPDomain) Validator() coredomain.Validator {
	return &gcpValidator{}
}

// Lister returns the GCP lister implementation
func (d *GCPDomain) Lister() coredomain.Lister {
	return &gcpLister{}
}

// Differ returns the GCP differ implementation (uses shared K8s differ)
func (d *GCPDomain) Differ() coredomain.Differ {
	return differ.New()
}

// CreateRootCommand creates the root command using the domain interface.
func CreateRootCommand(d coredomain.Domain) *cobra.Command {
	return coredomain.Run(d)
}

// gcpBuilder implements domain.Builder
type gcpBuilder struct{}

func (b *gcpBuilder) Build(ctx *Context, path string, opts BuildOpts) (*Result, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, fmt.Errorf("resolve path: %w", err)
	}

	// Discover all Config Connector resources
	resources, err := discover.DiscoverDirectory(absPath)
	if err != nil {
		return nil, fmt.Errorf("discover resources: %w", err)
	}

	if len(resources) == 0 {
		return NewResult("No Config Connector resources found"), nil
	}

	// Build YAML output
	var yamlDocs []string
	for _, r := range resources {
		// For now, output a placeholder YAML document
		// Full implementation would evaluate Go code and serialize the actual values
		doc := map[string]interface{}{
			"apiVersion": getAPIVersionFromType(r.Type),
			"kind":       getKindFromType(r.Type),
			"metadata": map[string]interface{}{
				"name": strings.ToLower(r.Name),
			},
			"spec": map[string]interface{}{
				"# TODO": "Implement full Go evaluation",
			},
		}

		yamlBytes, err := yaml.Marshal(doc)
		if err != nil {
			return nil, fmt.Errorf("marshal %s: %w", r.Name, err)
		}
		yamlDocs = append(yamlDocs, string(yamlBytes))
	}

	output := strings.Join(yamlDocs, "---\n")

	return NewResultWithData(output, map[string]interface{}{
		"resource_count": len(resources),
	}), nil
}

// getAPIVersionFromType extracts apiVersion from a type name like "computev1beta1.ComputeInstance"
func getAPIVersionFromType(typeName string) string {
	parts := strings.Split(typeName, ".")
	if len(parts) != 2 {
		return "unknown"
	}
	pkg := parts[0]
	// Extract service and version from package name
	// e.g., "computev1beta1" -> "compute.cnrm.cloud.google.com/v1beta1"
	for i := len(pkg) - 1; i >= 0; i-- {
		if pkg[i] == 'v' {
			service := pkg[:i]
			version := pkg[i:]
			return fmt.Sprintf("%s.cnrm.cloud.google.com/%s", service, version)
		}
	}
	return "unknown"
}

// getKindFromType extracts the Kind from a type name
func getKindFromType(typeName string) string {
	parts := strings.Split(typeName, ".")
	if len(parts) == 2 {
		return parts[1]
	}
	return typeName
}

// gcpLinter implements domain.Linter
type gcpLinter struct{}

func (l *gcpLinter) Lint(ctx *Context, path string, opts LintOpts) (*Result, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, fmt.Errorf("resolve path: %w", err)
	}

	lintOpts := lint.Options{
		DisabledRules: opts.Disable,
		Fix:           opts.Fix,
	}

	result, err := lint.LintPackage(absPath, lintOpts)
	if err != nil {
		return nil, fmt.Errorf("lint: %w", err)
	}

	if result.Success {
		return NewResult("No lint issues found"), nil
	}

	// Convert lint issues to domain errors
	var errors []Error
	for _, issue := range result.Issues {
		errors = append(errors, Error{
			Message:  issue.Message,
			Path:     issue.File,
			Line:     issue.Line,
			Column:   issue.Column,
			Severity: issue.Severity.String(),
			Code:     issue.Rule,
		})
	}

	return NewErrorResultMultiple(fmt.Sprintf("Found %d lint issues", len(errors)), errors), nil
}

// gcpInitializer implements domain.Initializer
type gcpInitializer struct{}

func (i *gcpInitializer) Init(ctx *Context, path string, opts InitOpts) (*Result, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, fmt.Errorf("resolve path: %w", err)
	}

	// Use opts.Path if provided, otherwise fall back to path argument
	targetPath := opts.Path
	if targetPath == "" || targetPath == "." {
		targetPath = absPath
	}

	// Handle scenario initialization
	if opts.Scenario {
		return i.initScenario(ctx, targetPath, opts)
	}

	// Basic project initialization
	return i.initProject(ctx, targetPath, opts)
}

// initScenario creates a full scenario structure with prompts and expected outputs
func (i *gcpInitializer) initScenario(ctx *Context, path string, opts InitOpts) (*Result, error) {
	name := opts.Name
	if name == "" {
		name = filepath.Base(path)
	}

	description := opts.Description
	if description == "" {
		description = "GCP Config Connector scenario"
	}

	// Use core's scenario scaffolding
	scenario := coredomain.ScaffoldScenario(name, description, "gcp")
	created, err := coredomain.WriteScenario(path, scenario)
	if err != nil {
		return nil, fmt.Errorf("write scenario: %w", err)
	}

	return NewResultWithData(fmt.Sprintf("Created scenario at %s", path), map[string]interface{}{
		"files_created": created,
	}), nil
}

// initProject creates a basic GCP project structure
func (i *gcpInitializer) initProject(ctx *Context, path string, opts InitOpts) (*Result, error) {
	// TODO: Implement GCP project initialization
	return NewResult(fmt.Sprintf("GCP init for %s (not yet implemented)", path)), nil
}

// gcpValidator implements domain.Validator
type gcpValidator struct{}

func (v *gcpValidator) Validate(ctx *Context, path string, opts ValidateOpts) (*Result, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, fmt.Errorf("resolve path: %w", err)
	}

	// TODO: Implement Config Connector schema validation
	return NewResult(fmt.Sprintf("GCP validate for %s (not yet implemented)", absPath)), nil
}

// gcpLister implements domain.Lister
type gcpLister struct{}

func (l *gcpLister) List(ctx *Context, path string, opts ListOpts) (*Result, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, fmt.Errorf("resolve path: %w", err)
	}

	resources, err := discover.DiscoverDirectory(absPath)
	if err != nil {
		return nil, fmt.Errorf("discover resources: %w", err)
	}

	if len(resources) == 0 {
		return NewResult("No Config Connector resources found"), nil
	}

	// Format output
	var lines []string
	for _, r := range resources {
		lines = append(lines, fmt.Sprintf("%s (%s) - %s:%d", r.Name, r.Type, filepath.Base(r.File), r.Line))
	}
	sort.Strings(lines)

	return NewResultWithData(
		fmt.Sprintf("Found %d Config Connector resources", len(resources)),
		map[string]interface{}{
			"resources": resources,
			"count":     len(resources),
		},
	), nil
}
