---
title: "FAQ"
---

## General

<details>
<summary>What is wetwire-gcp-go?</summary>

wetwire-gcp-go generates Config Connector manifests from Go code. It enables type-safe GCP infrastructure definitions with IDE support and AI assistance.
</details>

<details>
<summary>How does it differ from Terraform?</summary>

| Aspect | wetwire-gcp-go | Terraform |
|--------|----------------|-----------|
| Language | Go | HCL |
| Type Safety | Full Go types | Limited |
| IDE Support | Full autocomplete | Plugin-based |
| Output | Config Connector YAML | Provider-specific |
| State | Kubernetes (via CC) | Terraform state |
</details>

<details>
<summary>What is Config Connector?</summary>

Config Connector is a Kubernetes add-on that manages GCP resources declaratively. You define resources as Kubernetes CRDs, and Config Connector creates/updates the actual GCP resources.
</details>

<details>
<summary>Do I need a Kubernetes cluster?</summary>

Yes, to apply the generated manifests. Config Connector runs inside a GKE cluster (or any cluster with Config Connector installed).
</details>

---

## Resources

<details>
<summary>What GCP services are supported?</summary>

wetwire-gcp-go supports 450+ Config Connector CRDs including:

- **Compute**: Instances, disks, networks, firewalls
- **Storage**: Buckets, IAM bindings
- **SQL**: Instances, databases, users
- **GKE**: Clusters, node pools
- **IAM**: Service accounts, roles, bindings
- **Pub/Sub**: Topics, subscriptions
- **BigQuery**: Datasets, tables
- **And many more...**
</details>

<details>
<summary>How do I reference other resources?</summary>

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
</details>

<details>
<summary>How do I set the GCP project?</summary>

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
</details>

<details>
<summary>How do I handle cross-project references?</summary>

For resources that need to reference resources in other projects, use external references with the project specified:

```go
var MySubnet = computev1beta1.ComputeSubnetwork{
    Spec: computev1beta1.ComputeSubnetworkSpec{
        NetworkRef: computev1beta1.NetworkRef{
            External: "projects/shared-vpc-project/global/networks/shared-network",
        },
    },
}
```

The `External` field allows referencing resources by their full GCP resource path, enabling cross-project dependencies.
</details>

---

## Syntax

<details>
<summary>Why package-level variables?</summary>

Package-level variables enable:
1. AST-based discovery without runtime execution
2. Direct references between resources
3. IDE autocomplete and type checking
</details>

<details>
<summary>Can I use functions?</summary>

Avoid functions for resource definitions. The AST parser discovers `var X = Type{...}` declarations. Functions would require runtime execution.

```go
// Good - discoverable
var MyInstance = computev1beta1.ComputeInstance{...}

// Bad - not discoverable
func createInstance() computev1beta1.ComputeInstance {
    return computev1beta1.ComputeInstance{...}
}
```
</details>

<details>
<summary>How do I handle optional fields?</summary>

Use Go's zero values. Empty/nil fields are omitted from output:

```go
var MyInstance = computev1beta1.ComputeInstance{
    Spec: computev1beta1.ComputeInstanceSpec{
        MachineType: "n1-standard-1",
        // Zone omitted - will use namespace default
    },
}
```
</details>

<details>
<summary>What's the recommended project structure?</summary>

Organize resources by service or environment:

```
my-gcp-infra/
├── go.mod
├── main.go           # Entry point
├── network/
│   ├── vpc.go        # VPC and subnets
│   └── firewall.go   # Firewall rules
├── compute/
│   ├── instances.go  # Compute instances
│   └── disks.go      # Persistent disks
├── storage/
│   └── buckets.go    # Storage buckets
└── iam/
    └── bindings.go   # IAM bindings
```

Group related resources together and use separate directories for different environments (dev, staging, prod) if needed.
</details>

---

## Building

<details>
<summary>How do I generate manifests?</summary>

```bash
wetwire-gcp build ./gcp -o manifests.yaml
```
</details>

<details>
<summary>Can I output JSON instead of YAML?</summary>

```bash
wetwire-gcp build ./gcp --format json
```
</details>

<details>
<summary>How do I apply the manifests?</summary>

```bash
wetwire-gcp build ./gcp | kubectl apply -f -
```
</details>

<details>
<summary>How do I integrate wetwire-gcp with my CI/CD pipeline?</summary>

Use wetwire-gcp in your pipeline to generate and validate manifests:

```yaml
# GitHub Actions example
steps:
  - name: Generate manifests
    run: wetwire-gcp build ./gcp -o manifests.yaml

  - name: Lint infrastructure code
    run: wetwire-gcp lint ./gcp

  - name: Validate against schemas
    run: wetwire-gcp validate ./gcp

  - name: Diff against current state
    run: |
      wetwire-gcp build ./gcp -o new.yaml
      wetwire-gcp diff current.yaml new.yaml
```

The `diff` command returns exit code 1 if changes are detected, enabling automated change detection.
</details>

<details>
<summary>Can I import existing Deployment Manager templates?</summary>

Direct import is not currently supported. However, you can:

1. Use the `design` command to describe your existing infrastructure
2. Manually translate templates to Go structs
3. Use Config Connector's documentation to map DM resources to CC CRDs

The Go type system will catch any misconfigurations during translation.
</details>

---

## Linting

<details>
<summary>What does the linter check?</summary>

See the [Lint Rules]({{< relref "/lint-rules" >}}) page for all rules. Key checks include:
- No pointers in declarations
- Valid Config Connector types
- Required fields present
- Naming conventions
</details>

<details>
<summary>How do I auto-fix issues?</summary>

```bash
wetwire-gcp lint ./gcp --fix
```
</details>

<details>
<summary>How does the linter help catch errors?</summary>

The linter provides multiple layers of protection:

1. **Structural checks**: Ensures resources are declared correctly as top-level variables
2. **Type validation**: Verifies Config Connector types and required fields
3. **Reference validation**: Checks that resource references point to existing resources
4. **Naming conventions**: Enforces GCP naming constraints (length, characters)
5. **Security warnings**: Alerts on public IAM bindings or missing project annotations

Run `wetwire-gcp lint --fix` to automatically correct fixable issues.
</details>

---

## AI Assistance

<details>
<summary>How do I use the design command?</summary>

```bash
wetwire-gcp design "Create a GKE cluster with autoscaling"
```

The AI generates Go code following wetwire patterns.
</details>

<details>
<summary>What AI providers are supported?</summary>

- **claude** (default) - Claude Code CLI
- **anthropic** - Anthropic API directly
- **kiro** - Kiro provider
</details>

---

## Troubleshooting

<details>
<summary>"No resources discovered"</summary>

Ensure resources are:
1. Package-level variables (not inside functions)
2. Using types from `resources/cnrm/`
3. In `.go` files (not `_test.go`)
</details>

<details>
<summary>"Unknown type"</summary>

The type may not be in the generated resources. Check:
1. Correct import path (`resources/cnrm/compute/v1beta1`)
2. Type name matches Config Connector CRD
</details>

<details>
<summary>Config Connector errors</summary>

Config Connector errors appear in the Kubernetes resource status. Check with:

```bash
kubectl describe <resource-type> <name>
```
</details>

---

## Resources

- [Config Connector Documentation](https://cloud.google.com/config-connector/docs/overview)
- [Config Connector CRD Reference](https://cloud.google.com/config-connector/docs/reference/overview)
