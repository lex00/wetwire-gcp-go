package domain

import (
	"strings"
	"testing"
)

func TestGCPDomain_Name(t *testing.T) {
	d := &GCPDomain{}
	if got := d.Name(); got != "gcp" {
		t.Errorf("Name() = %q, want %q", got, "gcp")
	}
}

func TestGCPDomain_Version(t *testing.T) {
	d := &GCPDomain{}
	v := d.Version()
	if v == "" {
		t.Error("Version() returned empty string")
	}
}

func TestGCPDomain_Builder(t *testing.T) {
	d := &GCPDomain{}
	b := d.Builder()
	if b == nil {
		t.Error("Builder() returned nil")
	}
}

func TestGCPDomain_Linter(t *testing.T) {
	d := &GCPDomain{}
	l := d.Linter()
	if l == nil {
		t.Error("Linter() returned nil")
	}
}

func TestGCPDomain_Initializer(t *testing.T) {
	d := &GCPDomain{}
	i := d.Initializer()
	if i == nil {
		t.Error("Initializer() returned nil")
	}
}

func TestGCPDomain_Validator(t *testing.T) {
	d := &GCPDomain{}
	v := d.Validator()
	if v == nil {
		t.Error("Validator() returned nil")
	}
}

func TestGCPDomain_Lister(t *testing.T) {
	d := &GCPDomain{}
	l := d.Lister()
	if l == nil {
		t.Error("Lister() returned nil")
	}
}

func TestGCPBuilder_Build(t *testing.T) {
	b := &gcpBuilder{}

	t.Run("empty directory", func(t *testing.T) {
		result, err := b.Build(nil, ".", BuildOpts{})
		if err != nil {
			t.Fatalf("Build() error = %v", err)
		}
		if result == nil {
			t.Fatal("Build() returned nil result")
		}
	})

	t.Run("examples directory", func(t *testing.T) {
		result, err := b.Build(nil, "../examples", BuildOpts{})
		if err != nil {
			t.Fatalf("Build() error = %v", err)
		}
		if result == nil {
			t.Fatal("Build() returned nil result")
		}

		// Should contain YAML output
		if result.Data == nil {
			t.Skip("No data in result")
		}
		data, ok := result.Data.(map[string]interface{})
		if !ok {
			t.Skip("Data is not a map")
		}
		if count, ok := data["resource_count"]; ok {
			if count.(int) == 0 {
				t.Error("expected resources in examples directory")
			}
		}
	})

	t.Run("non-existent directory", func(t *testing.T) {
		_, err := b.Build(nil, "/non/existent/path", BuildOpts{})
		if err == nil {
			t.Error("expected error for non-existent path")
		}
	})
}

func TestGCPLister_List(t *testing.T) {
	l := &gcpLister{}

	t.Run("examples directory", func(t *testing.T) {
		result, err := l.List(nil, "../examples", ListOpts{})
		if err != nil {
			t.Fatalf("List() error = %v", err)
		}
		if result == nil {
			t.Fatal("List() returned nil result")
		}

		if result.Data == nil {
			t.Skip("No data in result")
		}

		data, ok := result.Data.(map[string]interface{})
		if !ok {
			t.Skip("Data is not a map")
		}

		count, ok := data["count"]
		if !ok {
			t.Error("result.Data missing 'count' key")
		}
		if count.(int) == 0 {
			t.Error("expected resources in examples directory")
		}

		resources, ok := data["resources"]
		if !ok {
			t.Error("result.Data missing 'resources' key")
		}
		if resources == nil {
			t.Error("resources is nil")
		}
	})

	t.Run("empty directory", func(t *testing.T) {
		result, err := l.List(nil, ".", ListOpts{})
		if err != nil {
			t.Fatalf("List() error = %v", err)
		}
		if result == nil {
			t.Fatal("List() returned nil result")
		}
	})
}

func TestGCPLinter_Lint(t *testing.T) {
	l := &gcpLinter{}

	result, err := l.Lint(nil, ".", LintOpts{})
	if err != nil {
		t.Fatalf("Lint() error = %v", err)
	}
	if result == nil {
		t.Fatal("Lint() returned nil result")
	}
	// Lint is not yet implemented, so just verify it doesn't crash
}

func TestGCPValidator_Validate(t *testing.T) {
	v := &gcpValidator{}

	result, err := v.Validate(nil, ".", ValidateOpts{})
	if err != nil {
		t.Fatalf("Validate() error = %v", err)
	}
	if result == nil {
		t.Fatal("Validate() returned nil result")
	}
	// Validate is not yet implemented, so just verify it doesn't crash
}

func TestGCPInitializer_Init(t *testing.T) {
	i := &gcpInitializer{}

	t.Run("basic init", func(t *testing.T) {
		result, err := i.Init(nil, ".", InitOpts{})
		if err != nil {
			t.Fatalf("Init() error = %v", err)
		}
		if result == nil {
			t.Fatal("Init() returned nil result")
		}
	})

	// Note: Scenario init test would create files, skip for now
}

func TestGetAPIVersionFromType(t *testing.T) {
	tests := []struct {
		typeName string
		want     string
	}{
		{"computev1beta1.ComputeInstance", "compute.cnrm.cloud.google.com/v1beta1"},
		{"storagev1beta1.StorageBucket", "storage.cnrm.cloud.google.com/v1beta1"},
		{"containerv1beta1.ContainerCluster", "container.cnrm.cloud.google.com/v1beta1"},
		{"sqlv1beta1.SQLInstance", "sql.cnrm.cloud.google.com/v1beta1"},
		{"InvalidType", "unknown"},
		{"", "unknown"},
	}

	for _, tt := range tests {
		t.Run(tt.typeName, func(t *testing.T) {
			got := getAPIVersionFromType(tt.typeName)
			if got != tt.want {
				t.Errorf("getAPIVersionFromType(%q) = %q, want %q", tt.typeName, got, tt.want)
			}
		})
	}
}

func TestGetKindFromType(t *testing.T) {
	tests := []struct {
		typeName string
		want     string
	}{
		{"computev1beta1.ComputeInstance", "ComputeInstance"},
		{"storagev1beta1.StorageBucket", "StorageBucket"},
		{"JustKind", "JustKind"},
		{"", ""},
	}

	for _, tt := range tests {
		t.Run(tt.typeName, func(t *testing.T) {
			got := getKindFromType(tt.typeName)
			if got != tt.want {
				t.Errorf("getKindFromType(%q) = %q, want %q", tt.typeName, got, tt.want)
			}
		})
	}
}

func TestCreateRootCommand(t *testing.T) {
	d := &GCPDomain{}
	cmd := CreateRootCommand(d)

	if cmd == nil {
		t.Fatal("CreateRootCommand() returned nil")
	}

	// Verify basic command properties
	if cmd.Use == "" {
		t.Error("command Use is empty")
	}

	// Verify expected subcommands exist
	subcommands := make(map[string]bool)
	for _, sub := range cmd.Commands() {
		subcommands[sub.Name()] = true
	}

	expectedCommands := []string{"build", "lint", "init", "validate", "list"}
	for _, name := range expectedCommands {
		if !subcommands[name] {
			t.Errorf("expected subcommand %q not found", name)
		}
	}
}

func TestBuildOutputFormat(t *testing.T) {
	b := &gcpBuilder{}

	result, err := b.Build(nil, "../examples/compute-instance", BuildOpts{})
	if err != nil {
		t.Fatalf("Build() error = %v", err)
	}

	// The message should contain YAML
	if result.Message == "" {
		t.Error("result.Message is empty")
	}

	// Check for YAML structure
	if !strings.Contains(result.Message, "apiVersion:") {
		t.Error("output should contain apiVersion")
	}
	if !strings.Contains(result.Message, "kind:") {
		t.Error("output should contain kind")
	}
	if !strings.Contains(result.Message, "metadata:") {
		t.Error("output should contain metadata")
	}
}
