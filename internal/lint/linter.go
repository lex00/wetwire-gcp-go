// Package lint provides lint rules for wetwire-gcp-go code.
package lint

import (
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"

	corelint "github.com/lex00/wetwire-core-go/lint"
)

// Type aliases for compatibility with core lint package.
type (
	Issue    = corelint.Issue
	Severity = corelint.Severity
	Rule     = corelint.Rule
)

// Severity constants.
const (
	SeverityError   = corelint.SeverityError
	SeverityWarning = corelint.SeverityWarning
	SeverityInfo    = corelint.SeverityInfo
)

// Result contains the outcome of linting.
type Result struct {
	Success bool
	Issues  []Issue
}

// Options configures the linter.
type Options struct {
	// DisabledRules specifies rules to disable by ID (e.g., "WGC001").
	DisabledRules []string
	// Fix automatically fixes fixable issues (reserved for future use).
	Fix bool
}

// LintFile lints a single Go file.
func LintFile(path string, opts Options) (Result, error) {
	fset := token.NewFileSet()

	file, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		return Result{}, err
	}

	rules := getRules(opts)
	var issues []Issue

	for _, rule := range rules {
		ruleIssues := rule.Check(file, fset)
		issues = append(issues, ruleIssues...)
	}

	return Result{
		Success: len(issues) == 0,
		Issues:  issues,
	}, nil
}

// LintPackage lints all Go files in a package directory.
func LintPackage(pkgPath string, opts Options) (Result, error) {
	// Handle ... pattern
	if strings.HasSuffix(pkgPath, "/...") || strings.HasSuffix(pkgPath, "...") {
		root := strings.TrimSuffix(strings.TrimSuffix(pkgPath, "/..."), "...")
		return lintRecursive(root, opts)
	}

	fset := token.NewFileSet()

	pkgs, err := parser.ParseDir(fset, pkgPath, nil, parser.ParseComments)
	if err != nil {
		return Result{}, err
	}

	rules := getRules(opts)
	var allIssues []Issue

	for _, pkg := range pkgs {
		for _, file := range pkg.Files {
			for _, rule := range rules {
				issues := rule.Check(file, fset)
				allIssues = append(allIssues, issues...)
			}
		}
	}

	return Result{
		Success: len(allIssues) == 0,
		Issues:  allIssues,
	}, nil
}

// lintRecursive lints all Go packages recursively.
func lintRecursive(root string, opts Options) (Result, error) {
	var allIssues []Issue

	if root == "" || root == "." {
		root = "."
	}

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip vendor and hidden directories
		if info.IsDir() {
			if info.Name() == "vendor" || strings.HasPrefix(info.Name(), ".") {
				return filepath.SkipDir
			}
			// Skip resources/ directory (generated code)
			if info.Name() == "resources" {
				return filepath.SkipDir
			}
			return nil
		}

		// Only process Go files, skip tests
		if !strings.HasSuffix(path, ".go") || strings.HasSuffix(path, "_test.go") {
			return nil
		}

		result, err := LintFile(path, opts)
		if err != nil {
			return nil // Log but don't fail on parse errors
		}

		allIssues = append(allIssues, result.Issues...)
		return nil
	})

	if err != nil {
		return Result{}, err
	}

	return Result{
		Success: len(allIssues) == 0,
		Issues:  allIssues,
	}, nil
}

// getRules returns the rules to use based on options.
func getRules(opts Options) []Rule {
	all := AllRules()

	if len(opts.DisabledRules) == 0 {
		return all
	}

	disabled := make(map[string]bool)
	for _, id := range opts.DisabledRules {
		disabled[id] = true
	}

	var filtered []Rule
	for _, r := range all {
		if !disabled[r.ID()] {
			filtered = append(filtered, r)
		}
	}
	return filtered
}
