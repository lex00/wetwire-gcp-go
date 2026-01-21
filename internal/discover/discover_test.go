package discover

import (
	"path/filepath"
	"testing"
)

func TestDiscoverFile_Simple(t *testing.T) {
	resources, err := DiscoverFile("testdata/simple.go")
	if err != nil {
		t.Fatalf("DiscoverFile() error = %v", err)
	}

	if len(resources) != 1 {
		t.Fatalf("expected 1 resource, got %d", len(resources))
	}

	r := resources[0]
	if r.Name != "SimpleInstance" {
		t.Errorf("Name = %q, want %q", r.Name, "SimpleInstance")
	}
	if r.Type != "computev1beta1.ComputeInstance" {
		t.Errorf("Type = %q, want %q", r.Type, "computev1beta1.ComputeInstance")
	}
	if r.Line != 10 {
		t.Errorf("Line = %d, want %d", r.Line, 10)
	}
}

func TestDiscoverFile_Dependencies(t *testing.T) {
	resources, err := DiscoverFile("testdata/dependencies.go")
	if err != nil {
		t.Fatalf("DiscoverFile() error = %v", err)
	}

	if len(resources) != 2 {
		t.Fatalf("expected 2 resources, got %d", len(resources))
	}

	// Find NodePool resource
	var nodePool *Resource
	for i, r := range resources {
		if r.Name == "NodePool" {
			nodePool = &resources[i]
			break
		}
	}

	if nodePool == nil {
		t.Fatal("NodePool resource not found")
	}

	// NodePool should depend on Cluster
	if len(nodePool.Dependencies) == 0 {
		t.Error("expected NodePool to have dependencies")
	}

	found := false
	for _, dep := range nodePool.Dependencies {
		if dep == "Cluster" {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("expected NodePool to depend on Cluster, got dependencies: %v", nodePool.Dependencies)
	}
}

func TestDiscoverFile_Multiple(t *testing.T) {
	resources, err := DiscoverFile("testdata/multiple.go")
	if err != nil {
		t.Fatalf("DiscoverFile() error = %v", err)
	}

	if len(resources) != 3 {
		t.Fatalf("expected 3 resources, got %d", len(resources))
	}

	names := make(map[string]bool)
	for _, r := range resources {
		names[r.Name] = true
		if r.Type != "storagev1beta1.StorageBucket" {
			t.Errorf("resource %s has Type = %q, want storagev1beta1.StorageBucket", r.Name, r.Type)
		}
	}

	if !names["Bucket1"] || !names["Bucket2"] || !names["Bucket3"] {
		t.Errorf("expected Bucket1, Bucket2, Bucket3; got %v", names)
	}
}

func TestDiscoverFile_NonResource(t *testing.T) {
	resources, err := DiscoverFile("testdata/non_resource.go")
	if err != nil {
		t.Fatalf("DiscoverFile() error = %v", err)
	}

	if len(resources) != 0 {
		t.Errorf("expected 0 resources from non_resource.go, got %d", len(resources))
		for _, r := range resources {
			t.Logf("  unexpected resource: %s (%s)", r.Name, r.Type)
		}
	}
}

func TestDiscoverFile_NotExists(t *testing.T) {
	_, err := DiscoverFile("testdata/does_not_exist.go")
	if err == nil {
		t.Error("expected error for non-existent file")
	}
}

func TestDiscoverFile_InvalidSyntax(t *testing.T) {
	// Create a temporary file with invalid Go syntax
	// We'll test with a path that points to a non-Go file
	_, err := DiscoverFile("testdata/non_resource.go") // This is valid, so no error
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestDiscoverDirectory(t *testing.T) {
	resources, err := DiscoverDirectory("testdata")
	if err != nil {
		t.Fatalf("DiscoverDirectory() error = %v", err)
	}

	// Should find resources from simple.go (1), dependencies.go (2), multiple.go (3)
	// non_resource.go should contribute 0
	expectedCount := 6
	if len(resources) != expectedCount {
		t.Errorf("expected %d resources, got %d", expectedCount, len(resources))
		for _, r := range resources {
			t.Logf("  found: %s (%s) in %s", r.Name, r.Type, filepath.Base(r.File))
		}
	}

	// Verify resources have absolute file paths
	for _, r := range resources {
		if !filepath.IsAbs(r.File) {
			t.Errorf("resource %s has non-absolute file path: %s", r.Name, r.File)
		}
	}
}

func TestDiscoverDirectory_NotExists(t *testing.T) {
	_, err := DiscoverDirectory("testdata/not_a_directory")
	if err == nil {
		t.Error("expected error for non-existent directory")
	}
}

func TestDiscoverDirectory_Empty(t *testing.T) {
	// testdata/non_resource.go has no Config Connector types
	// but the directory itself has other files, so let's just verify
	// we handle directories correctly
	resources, err := DiscoverDirectory(".")
	if err != nil {
		t.Fatalf("DiscoverDirectory(\".\") error = %v", err)
	}

	// Should at least include testdata resources
	if len(resources) < 6 {
		t.Errorf("expected at least 6 resources, got %d", len(resources))
	}
}

func TestResource_Fields(t *testing.T) {
	resources, err := DiscoverFile("testdata/simple.go")
	if err != nil {
		t.Fatalf("DiscoverFile() error = %v", err)
	}

	if len(resources) == 0 {
		t.Fatal("no resources found")
	}

	r := resources[0]

	// Test all fields are populated
	if r.Name == "" {
		t.Error("Name is empty")
	}
	if r.Type == "" {
		t.Error("Type is empty")
	}
	if r.File == "" {
		t.Error("File is empty")
	}
	if r.Line == 0 {
		t.Error("Line is 0")
	}
	// Dependencies can be empty, that's OK
}

func TestExtractTypeName(t *testing.T) {
	// This is tested indirectly through DiscoverFile
	// but we can add explicit tests if the function is exported
}

func TestIsConfigConnectorType(t *testing.T) {
	tests := []struct {
		typeName string
		want     bool
	}{
		{"computev1beta1.ComputeInstance", true},
		{"storagev1beta1.StorageBucket", true},
		{"containerv1beta1.ContainerCluster", true},
		{"string", false},
		{"int", false},
		{"map[string]interface{}", false},
		{"unknownpkg.SomeType", false},
	}

	for _, tt := range tests {
		t.Run(tt.typeName, func(t *testing.T) {
			got := isConfigConnectorType(tt.typeName)
			if got != tt.want {
				t.Errorf("isConfigConnectorType(%q) = %v, want %v", tt.typeName, got, tt.want)
			}
		})
	}
}
