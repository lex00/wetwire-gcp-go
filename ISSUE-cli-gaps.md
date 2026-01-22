# Issue: Implement GCP CLI Commands Consistent with Other Domains

## Summary

wetwire-gcp-go has stub implementations for several CLI commands that are fully implemented in other domain packages (aws, azure, k8s, observability). These need to be implemented for feature parity.

## Current State

| Command | AWS | Azure | K8s | Obs | GCP |
|---------|-----|-------|-----|-----|-----|
| build | ✓ | ✓ | ✓ | ✓ | ✓ |
| lint | ✓ | ✓ | ✓ | ✓ | ✓ |
| list | ✓ | ✓ | ✓ | ✓ | ✓ |
| mcp | ✓ | ✓ | ✓ | ✓ | ✓ |
| **design** | ✓ full | ✓ full | ✓ full | ✓ full | **stub** |
| **test** | ✓ full | ✓ full | ✓ full | ✓ full | **stub** |
| **diff** | ✓ full | stub | ✓ full | stub | **stub** |
| **watch** | ✓ full | stub | ✓ full | stub | **stub** |

## Commands to Implement

### 1. `design` - AI-Assisted Infrastructure Design

**Purpose:** Interactive session where a human works with an AI to design and generate infrastructure code.

**What it does:**
- Accepts natural language prompt describing desired infrastructure
- AI asks clarifying questions about requirements
- Generates Go code using wetwire-gcp patterns
- Runs linter and fixes issues iteratively
- Builds final Config Connector manifests

**Providers to support:**
- `anthropic` (default) - Direct Anthropic API
- `kiro` - Kiro CLI integration

**Reference implementation:** `wetwire-aws-go/cmd/wetwire-aws/design.go` (203 lines)

**Key dependencies:**
- `wetwire-core-go/agent/agents.NewRunnerAgent`
- `wetwire-core-go/agent/orchestrator.NewHumanDeveloper`
- `wetwire-core-go/providers/anthropic`
- `internal/kiro` package (needs creation)

### 2. `test` - Persona-Based Synthesis Testing

**Purpose:** Automated testing with AI personas to evaluate code generation quality.

**What it does:**
- Runs AI agents with different personas (beginner, intermediate, expert)
- Each persona responds to the runner agent differently
- Validates that generated code passes lint and build
- Reports success/failure per persona

**Flags:**
- `--persona` - Which persona to test with
- `--all-personas` - Run all personas sequentially
- `--provider` - AI provider (anthropic, kiro)
- `--scenario` - Scenario name for tracking

**Reference implementation:** `wetwire-aws-go/cmd/wetwire-aws/test.go` (327 lines)

**Key dependencies:**
- `wetwire-core-go/agent/personas`
- `wetwire-core-go/agent/results`
- `internal/kiro.TestRunner` (needs creation)

### 3. `diff` - Semantic Manifest Comparison

**Purpose:** Compare two Config Connector manifest files and show semantic differences.

**What it does:**
- Parses two YAML/JSON manifest files
- Compares resources semantically (not textual diff)
- Reports added, removed, and modified resources
- Shows property-level changes for modifications

**Flags:**
- `--format` - Output format (text, json)
- `--ignore-order` - Ignore array element ordering

**Reference implementation:** `wetwire-aws-go/cmd/wetwire-aws/diff.go` (121 lines) + `internal/differ/`

**Key dependencies:**
- `internal/differ` package (needs creation)
- YAML parsing for Config Connector manifests

### 4. `watch` - File Watcher for Auto-Rebuild

**Purpose:** Monitor source files and automatically rebuild on changes.

**What it does:**
- Watches `.go` files in specified packages
- Debounces rapid changes
- Runs lint on each change
- Rebuilds if lint passes
- Reports results in terminal

**Flags:**
- `--lint-only` - Only run lint, skip build
- `--debounce` - Debounce duration (default 500ms)
- `--format` - Output format for build
- `--output` - Output file

**Reference implementation:** `wetwire-aws-go/cmd/wetwire-aws/watch.go` (371 lines)

**Key dependencies:**
- `github.com/fsnotify/fsnotify` (already in go.mod)
- `internal/discover`
- `internal/lint`

## Implementation Order

1. **watch** - Simplest, no AI dependencies, useful for development
2. **diff** - No AI dependencies, useful for CI/CD
3. **design** - Requires kiro package, most user-facing value
4. **test** - Requires design infrastructure, useful for quality assurance

## Additional Work Required

### Create `internal/kiro` Package

Port from wetwire-aws-go/internal/kiro:
- `config.go` - Kiro agent configuration
- `launch.go` - Launch Kiro CLI sessions
- `testrunner.go` - Test runner for persona testing

### Update GCPDomain for Agent Integration

Ensure `domain.GCPDomain` implements all methods needed by `agents.RunnerAgent`:
- `Build()`
- `Lint()`
- Domain-specific system prompts

## Acceptance Criteria

- [ ] `wetwire-gcp watch ./examples/gke-golden` works
- [ ] `wetwire-gcp diff manifest1.yaml manifest2.yaml` works
- [ ] `wetwire-gcp design "Create a GKE cluster"` starts interactive session
- [ ] `wetwire-gcp test --persona beginner "Create a storage bucket"` runs persona test
- [ ] All commands have tests matching other domain coverage
- [ ] CLI help text matches other domains' style

## Notes

- GCP uses Config Connector (K8s CRDs), so output format is YAML manifests, not CloudFormation/ARM
- The `mcp` command is already fully implemented and working
- Reference the AWS implementation as the most complete example
