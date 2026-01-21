package lint

import (
	"go/ast"
	"go/token"
	"strings"
)

// AllRules returns all available lint rules.
func AllRules() []Rule {
	return []Rule{
		TopLevelDeclaration{},
		NoPointerDeclarations{},
		RequireObjectMeta{},
		ValidResourceRef{},
		IAMMemberFormat{},
		PublicIAMWarning{},
	}
}

// TopLevelDeclaration ensures Config Connector resources are top-level var declarations.
// Rule ID: WGC001
type TopLevelDeclaration struct{}

func (r TopLevelDeclaration) ID() string { return "WGC001" }
func (r TopLevelDeclaration) Description() string {
	return "Config Connector resources must be top-level var declarations"
}

func (r TopLevelDeclaration) Check(file *ast.File, fset *token.FileSet) []Issue {
	var issues []Issue

	// Check for function declarations that return Config Connector types
	for _, decl := range file.Decls {
		funcDecl, ok := decl.(*ast.FuncDecl)
		if !ok {
			continue
		}

		// Check if function returns a Config Connector type
		if funcDecl.Type.Results == nil {
			continue
		}

		for _, result := range funcDecl.Type.Results.List {
			if isConfigConnectorType(result.Type) {
				pos := fset.Position(funcDecl.Pos())
				issues = append(issues, Issue{
					Rule:     r.ID(),
					Message:  "Function returns Config Connector type - use top-level var declaration instead",
					File:     pos.Filename,
					Line:     pos.Line,
					Column:   pos.Column,
					Severity: SeverityError,
				})
			}
		}
	}

	return issues
}

// NoPointerDeclarations detects pointer assignments in resource declarations.
// Rule ID: WGC002
type NoPointerDeclarations struct{}

func (r NoPointerDeclarations) ID() string { return "WGC002" }
func (r NoPointerDeclarations) Description() string {
	return "Avoid pointer assignments (&Type{}) - use value types"
}

func (r NoPointerDeclarations) Check(file *ast.File, fset *token.FileSet) []Issue {
	var issues []Issue

	ast.Inspect(file, func(n ast.Node) bool {
		unary, ok := n.(*ast.UnaryExpr)
		if !ok || unary.Op != token.AND {
			return true
		}

		// Check if it's &Type{...}
		if _, ok := unary.X.(*ast.CompositeLit); ok {
			pos := fset.Position(unary.Pos())
			issues = append(issues, Issue{
				Rule:       r.ID(),
				Message:    "Avoid pointer assignment (&Type{}) - use value type",
				Suggestion: "Remove & and use value type directly",
				File:       pos.Filename,
				Line:       pos.Line,
				Column:     pos.Column,
				Severity:   SeverityError,
			})
		}

		return true
	})

	return issues
}

// RequireObjectMeta ensures resources have ObjectMeta with Name.
// Rule ID: WGC101
type RequireObjectMeta struct{}

func (r RequireObjectMeta) ID() string { return "WGC101" }
func (r RequireObjectMeta) Description() string {
	return "Config Connector resources must have ObjectMeta with Name"
}

func (r RequireObjectMeta) Check(file *ast.File, fset *token.FileSet) []Issue {
	var issues []Issue

	for _, decl := range file.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok || genDecl.Tok != token.VAR {
			continue
		}

		for _, spec := range genDecl.Specs {
			valueSpec, ok := spec.(*ast.ValueSpec)
			if !ok || len(valueSpec.Values) == 0 {
				continue
			}

			for _, value := range valueSpec.Values {
				comp, ok := value.(*ast.CompositeLit)
				if !ok {
					continue
				}

				// Check if this is a Config Connector resource type
				if !isConfigConnectorType(comp.Type) {
					continue
				}

				// Check for ObjectMeta field
				hasObjectMeta := false
				hasName := false

				for _, elt := range comp.Elts {
					kv, ok := elt.(*ast.KeyValueExpr)
					if !ok {
						continue
					}

					key, ok := kv.Key.(*ast.Ident)
					if !ok {
						continue
					}

					if key.Name == "ObjectMeta" {
						hasObjectMeta = true
						// Check if ObjectMeta has Name
						if metaComp, ok := kv.Value.(*ast.CompositeLit); ok {
							for _, metaElt := range metaComp.Elts {
								if metaKV, ok := metaElt.(*ast.KeyValueExpr); ok {
									if metaKey, ok := metaKV.Key.(*ast.Ident); ok {
										if metaKey.Name == "Name" {
											hasName = true
										}
									}
								}
							}
						}
					}
				}

				if !hasObjectMeta {
					pos := fset.Position(comp.Pos())
					issues = append(issues, Issue{
						Rule:       r.ID(),
						Message:    "Config Connector resource missing ObjectMeta",
						Suggestion: "Add ObjectMeta: metav1.ObjectMeta{Name: \"resource-name\"}",
						File:       pos.Filename,
						Line:       pos.Line,
						Column:     pos.Column,
						Severity:   SeverityError,
					})
				} else if !hasName {
					pos := fset.Position(comp.Pos())
					issues = append(issues, Issue{
						Rule:       r.ID(),
						Message:    "ObjectMeta missing Name field",
						Suggestion: "Add Name field to ObjectMeta",
						File:       pos.Filename,
						Line:       pos.Line,
						Column:     pos.Column,
						Severity:   SeverityError,
					})
				}
			}
		}
	}

	return issues
}

// ValidResourceRef checks that ResourceRef fields reference valid resources.
// Rule ID: WGC201
type ValidResourceRef struct{}

func (r ValidResourceRef) ID() string { return "WGC201" }
func (r ValidResourceRef) Description() string {
	return "ResourceRef fields should reference existing resources"
}

func (r ValidResourceRef) Check(file *ast.File, fset *token.FileSet) []Issue {
	// This is a placeholder - full implementation would require package-level context
	// to track all defined resources across files
	return nil
}

// IAMMemberFormat validates IAM member field format.
// Rule ID: WGC301
type IAMMemberFormat struct{}

func (r IAMMemberFormat) ID() string { return "WGC301" }
func (r IAMMemberFormat) Description() string {
	return "IAM member fields must use valid format (user:, serviceAccount:, group:)"
}

var validIAMPrefixes = []string{
	"user:",
	"serviceAccount:",
	"group:",
	"domain:",
	"allUsers",
	"allAuthenticatedUsers",
}

func (r IAMMemberFormat) Check(file *ast.File, fset *token.FileSet) []Issue {
	var issues []Issue

	ast.Inspect(file, func(n ast.Node) bool {
		kv, ok := n.(*ast.KeyValueExpr)
		if !ok {
			return true
		}

		// Check if this is a Member field
		key, ok := kv.Key.(*ast.Ident)
		if !ok || key.Name != "Member" {
			return true
		}

		// Check the value
		lit, ok := kv.Value.(*ast.BasicLit)
		if !ok || lit.Kind != token.STRING {
			return true
		}

		value := strings.Trim(lit.Value, `"`)

		// Check if it has a valid prefix
		hasValidPrefix := false
		for _, prefix := range validIAMPrefixes {
			if strings.HasPrefix(value, prefix) || value == prefix {
				hasValidPrefix = true
				break
			}
		}

		if !hasValidPrefix && value != "" {
			pos := fset.Position(lit.Pos())
			issues = append(issues, Issue{
				Rule:       r.ID(),
				Message:    "IAM member must have valid prefix (user:, serviceAccount:, group:, domain:)",
				Suggestion: "Use format like 'user:email@example.com' or 'serviceAccount:sa@project.iam.gserviceaccount.com'",
				File:       pos.Filename,
				Line:       pos.Line,
				Column:     pos.Column,
				Severity:   SeverityError,
			})
		}

		return true
	})

	return issues
}

// PublicIAMWarning warns about public IAM bindings.
// Rule ID: WGC302
type PublicIAMWarning struct{}

func (r PublicIAMWarning) ID() string { return "WGC302" }
func (r PublicIAMWarning) Description() string {
	return "Warn when using public IAM bindings (allUsers, allAuthenticatedUsers)"
}

func (r PublicIAMWarning) Check(file *ast.File, fset *token.FileSet) []Issue {
	var issues []Issue

	ast.Inspect(file, func(n ast.Node) bool {
		lit, ok := n.(*ast.BasicLit)
		if !ok || lit.Kind != token.STRING {
			return true
		}

		value := strings.Trim(lit.Value, `"`)

		if value == "allUsers" || value == "allAuthenticatedUsers" {
			pos := fset.Position(lit.Pos())
			issues = append(issues, Issue{
				Rule:     r.ID(),
				Message:  "Public IAM binding detected - this grants access to everyone",
				File:     pos.Filename,
				Line:     pos.Line,
				Column:   pos.Column,
				Severity: SeverityWarning,
			})
		}

		return true
	})

	return issues
}

// isConfigConnectorType checks if a type expression is a Config Connector resource.
func isConfigConnectorType(expr ast.Expr) bool {
	switch t := expr.(type) {
	case *ast.SelectorExpr:
		// Check for patterns like computev1beta1.ComputeInstance
		if x, ok := t.X.(*ast.Ident); ok {
			// Check if the package name contains version indicator
			name := x.Name
			if strings.Contains(name, "v1") || strings.Contains(name, "v1beta1") || strings.Contains(name, "v1alpha1") {
				return true
			}
		}
	case *ast.Ident:
		// Direct type reference - check if it starts with common CC prefixes
		name := t.Name
		ccPrefixes := []string{
			"Compute", "Storage", "SQL", "Container", "IAM",
			"Pubsub", "BigQuery", "Spanner", "DNS", "Network",
		}
		for _, prefix := range ccPrefixes {
			if strings.HasPrefix(name, prefix) {
				return true
			}
		}
	}
	return false
}
