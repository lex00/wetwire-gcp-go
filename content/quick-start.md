---
title: "Quick Start"
slug: quick-start
---

Get started with `wetwire-gcp` in 5 minutes.

## Installation

```bash
go install github.com/lex00/wetwire-gcp-go/cmd/wetwire-gcp@latest
```

## Quick Test (No Setup Required)

You can test `wetwire-gcp` without creating a Go module:

```bash
mkdir test && cd test
cat > main.go << 'EOF'
package infra

import (
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    storagev1beta1 "github.com/lex00/wetwire-gcp-go/resources/cnrm/storage/v1beta1"
)

var Bucket = storagev1beta1.StorageBucket{
    ObjectMeta: metav1.ObjectMeta{
        Name:      "my-bucket",
        Namespace: "default",
    },
    Spec: storagev1beta1.StorageBucketSpec{
        Location: "US",
    },
}
EOF

wetwire-gcp build .
```

This works because `wetwire-gcp` auto-generates a synthetic module when no `go.mod` is found.

---

## Create a Project

For a real project:

```bash
wetwire-gcp init myproject
cd myproject
```

This creates:
```
myproject/
├── go.mod
└── main.go
```

## Define Resources

Edit `main.go` to define your GCP resources:

```go
package infra

import (
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    computev1beta1 "github.com/lex00/wetwire-gcp-go/resources/cnrm/compute/v1beta1"
)

var Instance = computev1beta1.ComputeInstance{
    ObjectMeta: metav1.ObjectMeta{
        Name:      "my-instance",
        Namespace: "default",
    },
    Spec: computev1beta1.ComputeInstanceSpec{
        MachineType: "n1-standard-1",
        Zone:        "us-central1-a",
    },
}
```

## Build and Validate

```bash
# Generate Config Connector manifests
wetwire-gcp build .

# Lint your code
wetwire-gcp lint .

# List resources
wetwire-gcp list .
```

## Deploy

Apply the generated manifests to a cluster with Config Connector:

```bash
wetwire-gcp build . -o manifests.yaml
kubectl apply -f manifests.yaml
```

## Next Steps

- [CLI Reference](/cli/) - Full command documentation
- [Examples](/examples/) - Working examples
- [FAQ](/faq/) - Common questions
