---
title: "FAQ"
---

<picture>
  <source media="(prefers-color-scheme: dark)" srcset="./wetwire-dark.svg">
  <img src="./wetwire-light.svg" width="100" height="67">
</picture>

This FAQ covers questions specific to wetwire-gcp-go. For general wetwire questions, see the [central FAQ](https://github.com/lex00/wetwire/blob/main/docs/FAQ.md).

---

## General

### What is wetwire-gcp-go?

wetwire-gcp-go generates Config Connector manifests from Go code. It enables type-safe GCP infrastructure definitions with IDE support and AI assistance.

### How does it differ from Terraform?

| Aspect | wetwire-gcp-go | Terraform |
|--------|----------------|-----------|
| Language | Go | HCL |
| Type Safety | Full Go types | Limited |
| IDE Support | Full autocomplete | Plugin-based |
| Output | Config Connector YAML | Provider-specific |
| State | Kubernetes (via CC) | Terraform state |

### What is Config Connector?

Config Connector is a Kubernetes add-on that manages GCP resources declaratively. You define resources as Kubernetes CRDs, and Config Connector creates/updates the actual GCP resources.

### Do I need a Kubernetes cluster?

Yes, to apply the generated manifests. Config Connector runs inside a GKE cluster (or any cluster with Config Connector installed).

---

## Resources

### What GCP services are supported?

wetwire-gcp-go supports 450+ Config Connector CRDs including:

- **Compute**: Instances, disks, networks, firewalls
- **Storage**: Buckets, IAM bindings
- **SQL**: Instances, databases, users
- **GKE**: Clusters, node pools
- **IAM**: Service accounts, roles, bindings
- **Pub/Sub**: Topics, subscriptions
- **BigQuery**: Datasets, tables
- **And many more...**

### How do I reference other resources?

Use Go variable references:

```go
var MyNetwork = computev1beta1.ComputeNetwork{
    ObjectMeta: metav1.ObjectMeta{
        Name: "my-network",
    },
}

var MySubnet = computev1beta1.ComputeSubnetwork{
    Spec: computev1beta1.ComputeSubnetworkSpec{
        NetworkRef: computev1beta1.NetworkRef{
            Name: MyNetwork.ObjectMeta.Name,  // Direct reference
        },
    },
}
```

### How do I set the GCP project?

Use annotations on resources:

```go
var MyBucket = storagev1beta1.StorageBucket{
    ObjectMeta: metav1.ObjectMeta{
        Name: "my-bucket",
        Annotations: map[string]string{
            "cnrm.cloud.google.com/project-id": "my-project",
        },
    },
}
```

Or set a namespace-level default in Config Connector.

---

## Syntax

### Why package-level variables?

Package-level variables enable:
1. AST-based discovery without runtime execution
2. Direct references between resources
3. IDE autocomplete and type checking

### Can I use functions?

Avoid functions for resource definitions. The AST parser discovers `var X = Type{...}` declarations. Functions would require runtime execution.

```go
// Good - discoverable
var MyInstance = computev1beta1.ComputeInstance{...}

// Bad - not discoverable
func createInstance() computev1beta1.ComputeInstance {
    return computev1beta1.ComputeInstance{...}
}
```

### How do I handle optional fields?

Use Go's zero values. Empty/nil fields are omitted from output:

```go
var MyInstance = computev1beta1.ComputeInstance{
    Spec: computev1beta1.ComputeInstanceSpec{
        MachineType: "n1-standard-1",
        // Zone omitted - will use namespace default
    },
}
```

---

## Building

### How do I generate manifests?

```bash
wetwire-gcp build ./gcp -o manifests.yaml
```

### Can I output JSON instead of YAML?

```bash
wetwire-gcp build ./gcp --format json
```

### How do I apply the manifests?

```bash
wetwire-gcp build ./gcp | kubectl apply -f -
```

---

## Linting

### What does the linter check?

See [LINT_RULES.md](LINT_RULES.md) for all rules. Key checks include:
- No pointers in declarations
- Valid Config Connector types
- Required fields present
- Naming conventions

### How do I auto-fix issues?

```bash
wetwire-gcp lint ./gcp --fix
```

---

## AI Assistance

### How do I use the design command?

```bash
wetwire-gcp design "Create a GKE cluster with autoscaling"
```

The AI generates Go code following wetwire patterns.

### What AI providers are supported?

- **claude** (default) - Claude Code CLI
- **anthropic** - Anthropic API directly
- **kiro** - Kiro provider

---

## Troubleshooting

### "No resources discovered"

Ensure resources are:
1. Package-level variables (not inside functions)
2. Using types from `resources/cnrm/`
3. In `.go` files (not `_test.go`)

### "Unknown type"

The type may not be in the generated resources. Check:
1. Correct import path (`resources/cnrm/compute/v1beta1`)
2. Type name matches Config Connector CRD

### Config Connector errors

Config Connector errors appear in the Kubernetes resource status. Check with:

```bash
kubectl describe <resource-type> <name>
```

---

## Resources

- [Config Connector Documentation](https://cloud.google.com/config-connector/docs/overview)
- [Config Connector CRD Reference](https://cloud.google.com/config-connector/docs/reference/overview)
- [Wetwire Specification](https://github.com/lex00/wetwire/blob/main/docs/WETWIRE_SPEC.md)
