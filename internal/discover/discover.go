// Package discover provides GCP Config Connector resource discovery from Go source files.
package discover

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"

	"github.com/lex00/wetwire-gcp-go/internal/registry"
)

// Resource represents a discovered GCP Config Connector resource.
type Resource struct {
	Name         string   // Variable name in Go code
	Type         string   // Full type name (e.g., "computev1beta1.ComputeInstance")
	File         string   // Source file path
	Line         int      // Line number in source
	Dependencies []string // Names of other resources this one references
}

// DiscoverFile discovers Config Connector resources in a single Go source file.
func DiscoverFile(filePath string) ([]Resource, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, filePath, nil, 0)
	if err != nil {
		return nil, fmt.Errorf("parse file %s: %w", filePath, err)
	}

	absPath, err := filepath.Abs(filePath)
	if err != nil {
		absPath = filePath
	}

	var resources []Resource

	for _, decl := range file.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok || genDecl.Tok != token.VAR {
			continue
		}

		for _, spec := range genDecl.Specs {
			valueSpec, ok := spec.(*ast.ValueSpec)
			if !ok {
				continue
			}

			for i, name := range valueSpec.Names {
				if name.Name == "_" {
					continue
				}

				var resourceType string
				if valueSpec.Type != nil {
					resourceType = getResourceType(valueSpec.Type)
				} else if i < len(valueSpec.Values) {
					resourceType = getResourceTypeFromExpr(valueSpec.Values[i])
				}

				if resourceType == "" {
					continue
				}

				var deps []string
				if i < len(valueSpec.Values) {
					deps = findDependencies(valueSpec.Values[i], file)
				}

				resource := Resource{
					Name:         name.Name,
					Type:         resourceType,
					File:         absPath,
					Line:         fset.Position(name.Pos()).Line,
					Dependencies: deps,
				}
				resources = append(resources, resource)
			}
		}
	}

	return resources, nil
}

// DiscoverDirectory discovers Config Connector resources in all Go files within a directory.
func DiscoverDirectory(dir string) ([]Resource, error) {
	var allResources []Resource

	absDir, err := filepath.Abs(dir)
	if err != nil {
		return nil, fmt.Errorf("resolve path: %w", err)
	}

	err = filepath.Walk(absDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories
		if info.IsDir() {
			// Skip vendor, .git, and hidden directories
			base := filepath.Base(path)
			if base == "vendor" || base == ".git" || strings.HasPrefix(base, ".") {
				return filepath.SkipDir
			}
			return nil
		}

		// Skip non-Go files and test files
		if !strings.HasSuffix(path, ".go") || strings.HasSuffix(path, "_test.go") {
			return nil
		}

		resources, err := DiscoverFile(path)
		if err != nil {
			// Log but continue
			return nil
		}

		allResources = append(allResources, resources...)
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("walk directory %s: %w", dir, err)
	}

	return allResources, nil
}

// getResourceType extracts the Config Connector resource type from an AST type expression.
func getResourceType(typeExpr ast.Expr) string {
	typeName, pkgName := extractTypeName(typeExpr)
	if typeName == "" {
		return ""
	}

	fullTypeName := typeName
	if pkgName != "" {
		fullTypeName = fmt.Sprintf("%s.%s", pkgName, typeName)
	}

	if isConfigConnectorType(fullTypeName) {
		return fullTypeName
	}
	return ""
}

// getResourceTypeFromExpr extracts the type from a value expression.
func getResourceTypeFromExpr(expr ast.Expr) string {
	typeName, pkgName := inferTypeFromValue(expr)
	if typeName == "" {
		return ""
	}

	fullTypeName := typeName
	if pkgName != "" {
		fullTypeName = fmt.Sprintf("%s.%s", pkgName, typeName)
	}

	if isConfigConnectorType(fullTypeName) {
		return fullTypeName
	}
	return ""
}

// extractTypeName extracts type and package name from a type expression.
func extractTypeName(expr ast.Expr) (typeName, pkgName string) {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name, ""
	case *ast.SelectorExpr:
		if ident, ok := t.X.(*ast.Ident); ok {
			return t.Sel.Name, ident.Name
		}
	case *ast.StarExpr:
		return extractTypeName(t.X)
	}
	return "", ""
}

// inferTypeFromValue infers type from a composite literal.
func inferTypeFromValue(expr ast.Expr) (typeName, pkgName string) {
	switch e := expr.(type) {
	case *ast.CompositeLit:
		return extractTypeName(e.Type)
	case *ast.UnaryExpr:
		if e.Op == token.AND {
			return inferTypeFromValue(e.X)
		}
	}
	return "", ""
}

// isConfigConnectorType checks if a type name is a known Config Connector resource type.
func isConfigConnectorType(typeName string) bool {
	if registry.DefaultRegistry.IsKnownType(typeName) {
		return true
	}

	parts := strings.Split(typeName, ".")
	if len(parts) == 2 {
		if registry.DefaultRegistry.IsKnownPackage(parts[0]) {
			return true
		}
	}

	return false
}

// findDependencies finds references to other top-level variables in an expression.
func findDependencies(expr ast.Expr, file *ast.File) []string {
	deps := make(map[string]bool)

	ast.Inspect(expr, func(n ast.Node) bool {
		switch node := n.(type) {
		case *ast.Ident:
			if isTopLevelVar(node.Name, file) {
				deps[node.Name] = true
			}
		case *ast.SelectorExpr:
			if ident, ok := node.X.(*ast.Ident); ok {
				if isTopLevelVar(ident.Name, file) {
					deps[ident.Name] = true
				}
			}
		}
		return true
	})

	var result []string
	for dep := range deps {
		result = append(result, dep)
	}
	return result
}

// isTopLevelVar checks if a name is a top-level variable in the file.
func isTopLevelVar(name string, file *ast.File) bool {
	for _, decl := range file.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok || genDecl.Tok != token.VAR {
			continue
		}

		for _, spec := range genDecl.Specs {
			valueSpec, ok := spec.(*ast.ValueSpec)
			if !ok {
				continue
			}

			for _, varName := range valueSpec.Names {
				if varName.Name == name {
					return true
				}
			}
		}
	}
	return false
}
