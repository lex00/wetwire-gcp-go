<picture>
  <source media="(prefers-color-scheme: dark)" srcset="docs/wetwire-dark.svg">
  <img src="docs/wetwire-light.svg" width="100" height="67" align="right">
</picture>


# wetwire-gcp-go

[![CI](https://github.com/lex00/wetwire-gcp-go/actions/workflows/ci.yml/badge.svg)](https://github.com/lex00/wetwire-gcp-go/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/lex00/wetwire-gcp-go/branch/main/graph/badge.svg)](https://codecov.io/gh/lex00/wetwire-gcp-go)
[![Go](https://img.shields.io/badge/Go-1.25-blue?logo=go)](https://golang.org/)
[![Go Reference](https://pkg.go.dev/badge/github.com/lex00/wetwire-gcp-go.svg)](https://pkg.go.dev/github.com/lex00/wetwire-gcp-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/lex00/wetwire-gcp-go)](https://goreportcard.com/report/github.com/lex00/wetwire-gcp-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Go implementation of wetwire for Google Cloud Platform using Config Connector.

## Overview

wetwire-gcp-go enables defining GCP resources using Go code with full type safety and IDE support. It generates Config Connector manifests that can be applied to a Kubernetes cluster with Config Connector installed.

## Installation

```bash
go install github.com/lex00/wetwire-gcp-go/cmd/wetwire-gcp@latest
```

## Quick Start

```go
package main

import (
    computev1beta1 "github.com/lex00/wetwire-gcp-go/resources/cnrm/compute/v1beta1"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MyInstance defines a GCP Compute Engine instance.
// Resources are declared as package-level variables for AST discovery.
var MyInstance = computev1beta1.ComputeInstance{
    TypeMeta: metav1.TypeMeta{
        APIVersion: "compute.cnrm.cloud.google.com/v1beta1",
        Kind:       "ComputeInstance",
    },
    ObjectMeta: metav1.ObjectMeta{
        Name: "my-instance",
    },
    Spec: computev1beta1.ComputeInstanceSpec{
        MachineType: "n1-standard-1",
        Zone:        "us-central1-a",
    },
}

func main() {
    // wetwire-gcp build discovers resources via AST parsing
    // No runtime execution is needed
}
```

## Commands

- `wetwire-gcp build` - Generate Config Connector manifests from Go code
- `wetwire-gcp lint` - Lint GCP resource definitions
- `wetwire-gcp validate` - Validate against Config Connector schemas
- `wetwire-gcp init` - Initialize a new project
- `wetwire-gcp list` - List discovered resources
- `wetwire-gcp design` - AI-assisted infrastructure design
- `wetwire-gcp test` - Run synthesis tests with AI personas
- `wetwire-gcp mcp` - Start MCP server for Claude Code integration

## AI-Assisted Design

Use the `design` command for interactive, AI-assisted infrastructure creation:

```bash
wetwire-gcp design "Create a GKE cluster with 3 nodes"
```

## Config Connector

This package generates manifests for [Config Connector](https://cloud.google.com/config-connector/docs/overview), a Kubernetes add-on that manages GCP resources declaratively.

### Supported Resources

wetwire-gcp-go includes types for 450+ Config Connector CRDs across GCP services:

- Compute Engine (instances, disks, networks)
- Cloud Storage (buckets, IAM)
- Cloud SQL (instances, databases)
- GKE (clusters, node pools)
- IAM (service accounts, policies)
- And many more...

## Documentation

- [CLI Reference](docs/CLI.md)
- [FAQ](docs/FAQ.md)
- [Lint Rules](docs/LINT_RULES.md)
- [CLAUDE.md](CLAUDE.md) - AI assistant context

## Related Projects

- [wetwire-core-go](https://github.com/lex00/wetwire-core-go) - Core agent infrastructure
- [wetwire-k8s-go](https://github.com/lex00/wetwire-k8s-go) - Kubernetes manifests
- [wetwire](https://github.com/lex00/wetwire) - Central documentation

## License

Apache 2.0
