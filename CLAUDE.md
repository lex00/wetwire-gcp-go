# CLAUDE.md

This file provides context for AI assistants working with wetwire-gcp-go.

## Overview

wetwire-gcp-go is a Go implementation of the wetwire pattern for Google Cloud Platform resources using Config Connector. It enables defining GCP resources using Go code with full type safety, IDE support, and AI-assisted development.

This package is part of the wetwire ecosystem and extends wetwire-k8s-go to support Config Connector CRDs.

## Repository Structure

```
wetwire-gcp-go/
├── cmd/
│   └── wetwire-gcp/      # Main CLI binary
├── domain/               # GCPDomain implementation
├── internal/             # Internal packages
├── resources/            # Generated Config Connector types (auto-generated)
│   └── cnrm/
│       ├── compute/v1beta1/
│       ├── storage/v1beta1/
│       └── ...
└── docs/                 # Documentation
```

## Key Concepts

### Config Connector

Config Connector is a Kubernetes add-on that allows managing GCP resources through Kubernetes CRDs. wetwire-gcp-go generates Config Connector manifests from Go code.

### Resource Types

GCP resources are defined as Go structs generated from Config Connector CRDs:

```go
import (
    computev1beta1 "github.com/lex00/wetwire-gcp-go/resources/cnrm/compute/v1beta1"
)

var MyInstance = computev1beta1.ComputeInstance{
    Metadata: metav1.ObjectMeta{
        Name: "my-instance",
    },
    Spec: computev1beta1.ComputeInstanceSpec{
        MachineType: "n1-standard-1",
        Zone:        "us-central1-a",
    },
}
```

### GCP Context

The GCPContext provides project, region, and zone resolution:

```go
var ctx = gcp.Context{
    Project: "my-project",
    Region:  "us-central1",
    Zone:    "us-central1-a",
}
```

## Common Tasks

### Building Manifests

```bash
# Build from current directory
wetwire-gcp build

# Build from specific path
wetwire-gcp build ./gcp

# Save to file
wetwire-gcp build -o manifests.yaml
```

### Linting

```bash
# Check for issues
wetwire-gcp lint

# Auto-fix issues
wetwire-gcp lint --fix
```

### Validation

```bash
# Validate against Config Connector schemas
wetwire-gcp validate
```

## Development Status

This package is in early development. Key features not yet implemented:

- [ ] Resource discovery using wetwire-k8s-go registry
- [ ] Config Connector CRD type generation
- [ ] GCP-specific lint rules
- [ ] Config Connector schema validation
- [ ] GCPContext for project/region/zone
- [ ] resourceRef pattern support

## Dependencies

- `wetwire-core-go` - Core agent infrastructure
- `wetwire-k8s-go` - CRD parsing and code generation
