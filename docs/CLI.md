<picture>
  <source media="(prefers-color-scheme: dark)" srcset="./wetwire-dark.svg">
  <img src="./wetwire-light.svg" width="100" height="67">
</picture>

The `wetwire-gcp` command provides tools for generating and validating Config Connector manifests from Go code.

## Quick Reference

| Command | Description |
|---------|-------------|
| `wetwire-gcp build` | Generate Config Connector manifests from Go source |
| `wetwire-gcp lint` | Lint code for issues |
| `wetwire-gcp init` | Initialize a new project |
| `wetwire-gcp validate` | Validate resources against schemas |
| `wetwire-gcp list` | List discovered resources |
| `wetwire-gcp design` | AI-assisted infrastructure design |
| `wetwire-gcp test` | Run automated persona-based testing |
| `wetwire-gcp mcp` | Start MCP server for Claude Code |

```bash
wetwire-gcp --help     # Show help
```

---

## build

Generate Config Connector manifests from Go source files.

```bash
# Generate YAML to stdout
wetwire-gcp build ./gcp

# Save to file
wetwire-gcp build ./gcp -o manifests.yaml

# Generate JSON format
wetwire-gcp build ./gcp --format json
```

### Options

| Option | Description |
|--------|-------------|
| `PATH` | Directory containing Go source files |
| `--format, -f {yaml,json}` | Output format (default: yaml) |
| `--output, -o FILE` | Output file (default: stdout) |

### How It Works

1. Parses Go source files using `go/ast`
2. Discovers `var X = Type{...}` resource declarations
3. Resolves Config Connector types from `resources/cnrm/`
4. Extracts resource dependencies from references
5. Outputs manifests in dependency order

---

## lint

Check code for issues and optionally auto-fix them.

```bash
# Check for issues
wetwire-gcp lint ./gcp

# Auto-fix issues
wetwire-gcp lint ./gcp --fix

# JSON output
wetwire-gcp lint ./gcp --format json
```

### Options

| Option | Description |
|--------|-------------|
| `PATH` | Directory to lint |
| `--fix` | Auto-fix issues where possible |
| `--format {text,json}` | Output format |

See [LINT_RULES.md](LINT_RULES.md) for all lint rules.

---

## init

Initialize a new wetwire-gcp project.

```bash
# Create new project
wetwire-gcp init my-gcp-infra

# Initialize in current directory
wetwire-gcp init .
```

### Generated Files

| File | Purpose |
|------|---------|
| `go.mod` | Go module definition |
| `main.go` | Entry point with example resource |
| `network.go` | Example VPC network |
| `compute.go` | Example compute instance |

---

## validate

Validate resources against Config Connector schemas.

```bash
# Validate all resources
wetwire-gcp validate ./gcp

# Validate specific file
wetwire-gcp validate ./gcp/compute.go
```

### Options

| Option | Description |
|--------|-------------|
| `PATH` | Directory or file to validate |
| `--strict` | Fail on warnings |

---

## list

List all discovered resources.

```bash
# List resources
wetwire-gcp list ./gcp

# Show details
wetwire-gcp list ./gcp --verbose
```

### Output

```
compute.cnrm.cloud.google.com/v1beta1/ComputeInstance
  MyInstance (compute.go:15)
  WebServer (compute.go:30)

storage.cnrm.cloud.google.com/v1beta1/StorageBucket
  DataBucket (storage.go:10)
```

---

## design

AI-assisted infrastructure design using Claude.

```bash
# Interactive design session
wetwire-gcp design "Create a GKE cluster with 3 nodes"

# With specific provider
wetwire-gcp design --provider anthropic "Create a Cloud SQL instance"
```

### Options

| Option | Description |
|--------|-------------|
| `PROMPT` | Description of desired infrastructure |
| `--provider {anthropic,claude,kiro}` | AI provider (default: claude) |
| `--output, -o DIR` | Output directory |

---

## test

Run automated persona-based testing.

```bash
# Run all personas
wetwire-gcp test ./scenarios/gke-cluster

# Run specific persona
wetwire-gcp test ./scenarios/gke-cluster --persona beginner

# With verbose output
wetwire-gcp test ./scenarios/gke-cluster --verbose
```

### Options

| Option | Description |
|--------|-------------|
| `SCENARIO` | Path to scenario directory |
| `--persona {beginner,intermediate,expert}` | Specific persona |
| `--all` | Run all personas |
| `--verbose` | Show detailed output |

---

## mcp

Start the MCP (Model Context Protocol) server for Claude Code integration.

```bash
wetwire-gcp mcp
```

### Claude Code Configuration

Add to your Claude Code MCP settings:

```json
{
  "mcpServers": {
    "wetwire-gcp": {
      "command": "wetwire-gcp",
      "args": ["mcp"]
    }
  }
}
```

### Available Tools

| Tool | Description |
|------|-------------|
| `wetwire_init` | Initialize new project |
| `wetwire_build` | Generate manifests |
| `wetwire_lint` | Lint code |
| `wetwire_validate` | Validate schemas |

---

## Environment Variables

| Variable | Description |
|----------|-------------|
| `ANTHROPIC_API_KEY` | API key for Anthropic provider |
| `GOOGLE_PROJECT` | Default GCP project |
| `GOOGLE_REGION` | Default GCP region |
| `GOOGLE_ZONE` | Default GCP zone |

---

## Resources

- [Config Connector Documentation](https://cloud.google.com/config-connector/docs/overview)
- [Wetwire Specification](https://github.com/lex00/wetwire/blob/main/docs/WETWIRE_SPEC.md)
