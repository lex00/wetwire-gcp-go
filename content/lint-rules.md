---
title: "Lint Rules"
---

<picture>
  <source media="(prefers-color-scheme: dark)" srcset="./wetwire-dark.svg">
  <img src="./wetwire-light.svg" width="100" height="67">
</picture>

This document describes lint rules for wetwire-gcp-go.

## Overview

The wetwire-gcp linter enforces flat, declarative patterns optimized for AI generation and human readability. Rules check for structural patterns, Config Connector best practices, and GCP-specific constraints.

**Status:** Implemented. The linter is located in `internal/lint/` and implements the following rules.

## Rule Naming Convention

All rules follow the format: `WGCxxx`

- **W** - Wetwire
- **GC** - GCP Config connector
- **xxx** - Three-digit number

## Rule Categories

| Range | Category |
|-------|----------|
| WGC001-WGC099 | Structural patterns |
| WGC100-WGC199 | Config Connector requirements |
| WGC200-WGC299 | Resource references |
| WGC300-WGC399 | IAM and security |
| WGC400-WGC499 | Naming conventions |

---

## Implemented Rules

### WGC001: Top-level resource declarations

**Description:** Config Connector resources MUST be declared as top-level variables.

**Severity:** Error

**Why:** Resource discovery relies on finding top-level variable declarations via AST parsing.

**Bad:**

```go
func CreateBucket(name string) storagev1beta1.StorageBucket {
    return storagev1beta1.StorageBucket{...}
}
```

**Good:**

```go
var MyBucket = storagev1beta1.StorageBucket{
    ObjectMeta: metav1.ObjectMeta{Name: "my-bucket"},
}
```

---

### WGC002: No pointer declarations

**Description:** Avoid pointer assignments (&Type{}) - use value types.

**Severity:** Error

**Why:** Value types are simpler and match Config Connector's expected format.

**Bad:**

```go
var MyBucket = &storagev1beta1.StorageBucket{
    ObjectMeta: metav1.ObjectMeta{Name: "my-bucket"},
}
```

**Good:**

```go
var MyBucket = storagev1beta1.StorageBucket{
    ObjectMeta: metav1.ObjectMeta{Name: "my-bucket"},
}
```

---

### WGC101: Required ObjectMeta

**Description:** All Config Connector resources MUST have ObjectMeta with Name.

**Severity:** Error

**Bad:**

```go
var MyBucket = storagev1beta1.StorageBucket{
    Spec: storagev1beta1.StorageBucketSpec{...},
}
```

**Good:**

```go
var MyBucket = storagev1beta1.StorageBucket{
    ObjectMeta: metav1.ObjectMeta{
        Name: "my-bucket",
    },
    Spec: storagev1beta1.StorageBucketSpec{...},
}
```

---

### WGC102: Project annotation recommended

**Description:** Resources SHOULD have project annotation or rely on namespace default.

**Severity:** Warning

**Good:**

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

---

### WGC201: Valid resource references

**Description:** ResourceRef fields must reference existing resources.

**Severity:** Error

**Bad:**

```go
var MySubnet = computev1beta1.ComputeSubnetwork{
    Spec: computev1beta1.ComputeSubnetworkSpec{
        NetworkRef: computev1beta1.NetworkRef{
            Name: "nonexistent-network",  // No matching resource
        },
    },
}
```

**Good:**

```go
var MyNetwork = computev1beta1.ComputeNetwork{...}

var MySubnet = computev1beta1.ComputeSubnetwork{
    Spec: computev1beta1.ComputeSubnetworkSpec{
        NetworkRef: computev1beta1.NetworkRef{
            Name: MyNetwork.ObjectMeta.Name,  // Direct reference
        },
    },
}
```

---

### WGC301: IAM member format

**Description:** IAM member fields must use valid format (user:, serviceAccount:, group:).

**Severity:** Error

**Bad:**

```go
Member: "someone@example.com"
```

**Good:**

```go
Member: "user:someone@example.com"
Member: "serviceAccount:my-sa@my-project.iam.gserviceaccount.com"
```

---

### WGC302: Avoid allUsers/allAuthenticatedUsers

**Description:** Warn when using public IAM bindings.

**Severity:** Warning

**Triggers on:**

```go
Member: "allUsers"
Member: "allAuthenticatedUsers"
```

---

### WGC401: Resource naming conventions

**Description:** Resource names should follow GCP naming constraints.

**Severity:** Error

**Constraints:**
- Storage buckets: 3-63 characters, lowercase, numbers, hyphens
- Compute instances: 1-63 characters, lowercase, numbers, hyphens
- SQL instances: 1-98 characters, lowercase, numbers, hyphens

---

### WGC402: Location required

**Description:** Resources requiring a location MUST specify it.

**Severity:** Error

**Applies to:**
- StorageBucket (location)
- ComputeInstance (zone)
- SQLInstance (region)

---

## Inherited from wetwire-k8s-go

Since Config Connector resources are Kubernetes CRDs, the following WK8 rules also apply:

| Rule | Description |
|------|-------------|
| WK8001 | Top-level resource declarations |
| WK8002 | Avoid deeply nested structures |
| WK8003 | No duplicate resource names |
| WK8004 | Circular dependency detection |

See [wetwire-k8s-go LINT_RULES.md](https://github.com/lex00/wetwire-k8s-go/blob/main/docs/LINT_RULES.md) for details.

---

## Resources

- [Config Connector Resource Reference](https://cloud.google.com/config-connector/docs/reference/overview)
- [GCP Naming Conventions](https://cloud.google.com/compute/docs/naming-resources)
- [Wetwire Specification](https://github.com/lex00/wetwire/blob/main/docs/WETWIRE_SPEC.md)
