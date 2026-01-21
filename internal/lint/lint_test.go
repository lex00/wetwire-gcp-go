package lint

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLintFile_NoIssues(t *testing.T) {
	// Create a temporary file with valid code
	dir := t.TempDir()
	file := filepath.Join(dir, "valid.go")
	content := `package main

import "fmt"

func main() {
	fmt.Println("hello")
}
`
	if err := os.WriteFile(file, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	result, err := LintFile(file, Options{})
	if err != nil {
		t.Fatalf("LintFile error: %v", err)
	}

	if !result.Success {
		t.Errorf("expected success, got %d issues", len(result.Issues))
	}
}

func TestLintFile_PointerDeclaration(t *testing.T) {
	dir := t.TempDir()
	file := filepath.Join(dir, "pointer.go")
	content := `package gcp

import storagev1beta1 "example.com/storage/v1beta1"

var MyBucket = &storagev1beta1.StorageBucket{}
`
	if err := os.WriteFile(file, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	result, err := LintFile(file, Options{})
	if err != nil {
		t.Fatalf("LintFile error: %v", err)
	}

	if result.Success {
		t.Error("expected issues for pointer declaration")
	}

	found := false
	for _, issue := range result.Issues {
		if issue.Rule == "WGC002" {
			found = true
			break
		}
	}
	if !found {
		t.Error("expected WGC002 issue for pointer declaration")
	}
}

func TestLintFile_FunctionReturningCCType(t *testing.T) {
	dir := t.TempDir()
	file := filepath.Join(dir, "func.go")
	content := `package gcp

import computev1beta1 "example.com/compute/v1beta1"

func CreateInstance() computev1beta1.ComputeInstance {
	return computev1beta1.ComputeInstance{}
}
`
	if err := os.WriteFile(file, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	result, err := LintFile(file, Options{})
	if err != nil {
		t.Fatalf("LintFile error: %v", err)
	}

	if result.Success {
		t.Error("expected issues for function returning CC type")
	}

	found := false
	for _, issue := range result.Issues {
		if issue.Rule == "WGC001" {
			found = true
			break
		}
	}
	if !found {
		t.Error("expected WGC001 issue for function returning CC type")
	}
}

func TestLintFile_IAMMemberFormat(t *testing.T) {
	dir := t.TempDir()
	file := filepath.Join(dir, "iam.go")
	content := `package gcp

var binding = struct {
	Member string
}{
	Member: "someone@example.com",
}
`
	if err := os.WriteFile(file, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	result, err := LintFile(file, Options{})
	if err != nil {
		t.Fatalf("LintFile error: %v", err)
	}

	if result.Success {
		t.Error("expected issues for invalid IAM member format")
	}

	found := false
	for _, issue := range result.Issues {
		if issue.Rule == "WGC301" {
			found = true
			break
		}
	}
	if !found {
		t.Error("expected WGC301 issue for invalid IAM member format")
	}
}

func TestLintFile_ValidIAMMember(t *testing.T) {
	dir := t.TempDir()
	file := filepath.Join(dir, "iam_valid.go")
	content := `package gcp

var binding = struct {
	Member string
}{
	Member: "user:someone@example.com",
}
`
	if err := os.WriteFile(file, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	result, err := LintFile(file, Options{})
	if err != nil {
		t.Fatalf("LintFile error: %v", err)
	}

	// Should not have WGC301 issues
	for _, issue := range result.Issues {
		if issue.Rule == "WGC301" {
			t.Errorf("unexpected WGC301 issue for valid IAM member: %s", issue.Message)
		}
	}
}

func TestLintFile_PublicIAMWarning(t *testing.T) {
	dir := t.TempDir()
	file := filepath.Join(dir, "public.go")
	content := `package gcp

var binding = struct {
	Member string
}{
	Member: "allUsers",
}
`
	if err := os.WriteFile(file, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	result, err := LintFile(file, Options{})
	if err != nil {
		t.Fatalf("LintFile error: %v", err)
	}

	found := false
	for _, issue := range result.Issues {
		if issue.Rule == "WGC302" {
			found = true
			if issue.Severity != SeverityWarning {
				t.Errorf("expected warning severity, got %v", issue.Severity)
			}
			break
		}
	}
	if !found {
		t.Error("expected WGC302 warning for allUsers")
	}
}

func TestLintFile_DisabledRules(t *testing.T) {
	dir := t.TempDir()
	file := filepath.Join(dir, "disabled.go")
	content := `package gcp

import storagev1beta1 "example.com/storage/v1beta1"

var MyBucket = &storagev1beta1.StorageBucket{}
`
	if err := os.WriteFile(file, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	// Disable WGC002
	result, err := LintFile(file, Options{DisabledRules: []string{"WGC002"}})
	if err != nil {
		t.Fatalf("LintFile error: %v", err)
	}

	for _, issue := range result.Issues {
		if issue.Rule == "WGC002" {
			t.Error("WGC002 should be disabled")
		}
	}
}

func TestLintPackage_Recursive(t *testing.T) {
	dir := t.TempDir()

	// Create nested structure
	subdir := filepath.Join(dir, "sub")
	if err := os.MkdirAll(subdir, 0755); err != nil {
		t.Fatal(err)
	}

	// File in root
	file1 := filepath.Join(dir, "root.go")
	content1 := `package root

func main() {}
`
	if err := os.WriteFile(file1, []byte(content1), 0644); err != nil {
		t.Fatal(err)
	}

	// File in subdir
	file2 := filepath.Join(subdir, "sub.go")
	content2 := `package sub

func helper() {}
`
	if err := os.WriteFile(file2, []byte(content2), 0644); err != nil {
		t.Fatal(err)
	}

	// Use ... pattern
	result, err := LintPackage(dir+"/...", Options{})
	if err != nil {
		t.Fatalf("LintPackage error: %v", err)
	}

	// Both files should be linted without error
	if !result.Success {
		t.Logf("Issues: %v", result.Issues)
	}
}

func TestGetRules(t *testing.T) {
	rules := getRules(Options{})
	if len(rules) == 0 {
		t.Error("expected at least one rule")
	}

	// Check that all rules have ID and Description
	for _, r := range rules {
		if r.ID() == "" {
			t.Error("rule has empty ID")
		}
		if r.Description() == "" {
			t.Errorf("rule %s has empty description", r.ID())
		}
	}
}

func TestAllRules(t *testing.T) {
	rules := AllRules()

	expectedIDs := []string{"WGC001", "WGC002", "WGC101", "WGC201", "WGC301", "WGC302"}
	if len(rules) != len(expectedIDs) {
		t.Errorf("expected %d rules, got %d", len(expectedIDs), len(rules))
	}

	for i, r := range rules {
		if r.ID() != expectedIDs[i] {
			t.Errorf("expected rule %s at index %d, got %s", expectedIDs[i], i, r.ID())
		}
	}
}
