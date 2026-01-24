---
title: "Examples"
---

The `examples/` directory contains example wetwire-gcp-go projects demonstrating common GCP infrastructure patterns using Config Connector resources.

## Available Examples

### [compute-instance](https://github.com/lex00/wetwire-gcp-go/tree/main/examples/compute-instance)

A minimal example showing how to define a GCP Compute Engine instance.

**Concepts demonstrated:**
- Basic resource definition
- Machine type configuration
- Boot disk setup
- Network interface

```go
var MyInstance = computev1beta1.ComputeInstance{
    ObjectMeta: metav1.ObjectMeta{
        Name:      "my-instance",
        Namespace: "default",
    },
    Spec: computev1beta1.ComputeInstanceSpec{
        MachineType: "n1-standard-1",
        Zone:        "us-central1-a",
        BootDisk: map[string]interface{}{
            "initializeParams": map[string]interface{}{
                "image": "debian-cloud/debian-11",
            },
        },
    },
}
```

### [storage-bucket](https://github.com/lex00/wetwire-gcp-go/tree/main/examples/storage-bucket)

Cloud Storage bucket configuration.

**Concepts demonstrated:**
- Storage bucket setup
- Location configuration
- Access control

### [gke-cluster](https://github.com/lex00/wetwire-gcp-go/tree/main/examples/gke-cluster)

Google Kubernetes Engine cluster definition.

**Concepts demonstrated:**
- GKE cluster configuration
- Node pool setup
- Networking options

### [gke-golden](https://github.com/lex00/wetwire-gcp-go/tree/main/examples/gke-golden)

Production-ready GKE cluster configuration with best practices.

### [gke-k8s](https://github.com/lex00/wetwire-gcp-go/tree/main/examples/gke-k8s)

GKE cluster with Kubernetes resource integration.

### [sql-instance](https://github.com/lex00/wetwire-gcp-go/tree/main/examples/sql-instance)

Cloud SQL database instance configuration.

### [web-app](https://github.com/lex00/wetwire-gcp-go/tree/main/examples/web-app)

Complete web application infrastructure pattern.

## Running Examples

```bash
# Build Config Connector manifests
cd examples/compute-instance
wetwire-gcp build .

# Build and save to file
wetwire-gcp build . -o manifests.yaml

# Lint the configuration
wetwire-gcp lint .

# View dependency graph
wetwire-gcp graph . --format mermaid
```

## Deploying to GCP

After generating the manifests:

```bash
# Apply to cluster with Config Connector installed
kubectl apply -f manifests.yaml

# Check resource status
kubectl get computeinstance my-instance -o yaml
```

## Example Structure

Each example follows this structure:

```
example-name/
├── main.go         # Resource definitions
└── go.mod          # Go module file
```

## Creating Your Own

1. Initialize a new project:
   ```bash
   wetwire-gcp init myproject
   ```

2. Define resources in `main.go`

3. Build and validate:
   ```bash
   wetwire-gcp build .
   wetwire-gcp lint .
   ```

## See Also

- [CLI](/cli/) - Command-line interface reference
- [FAQ](/faq/) - Frequently asked questions
