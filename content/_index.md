---
title: "Wetwire GCP"
---

[![Go Reference](https://pkg.go.dev/badge/github.com/lex00/wetwire-gcp-go.svg)](https://pkg.go.dev/github.com/lex00/wetwire-gcp-go)
[![CI](https://github.com/lex00/wetwire-gcp-go/actions/workflows/ci.yml/badge.svg)](https://github.com/lex00/wetwire-gcp-go/actions/workflows/ci.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Generate Google Cloud Deployment Manager templates from Go structs with AI-assisted design.

## Philosophy

Wetwire uses typed constraints to reduce the model capability required for accurate code generation.

**Core hypothesis:** Typed input + smaller model ≈ Semantic input + larger model

The type system and lint rules act as a force multiplier — cheaper models produce quality output when guided by schema-generated types and iterative lint feedback.

## Documentation

| Document | Description |
|----------|-------------|
| [CLI Reference]({{< relref "/cli" >}}) | Command-line interface |
| [FAQ]({{< relref "/faq" >}}) | Frequently asked questions |
| [Lint Rules]({{< relref "/lint-rules" >}}) | Linting reference |

## Installation

```bash
go install github.com/lex00/wetwire-gcp-go@latest
```

## Quick Example

```go
var MyBucket = storage.Bucket{
    Name:     "my-data-bucket",
    Location: "us-central1",
}
```
