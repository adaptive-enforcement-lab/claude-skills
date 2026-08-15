# Kubernetes Integration

Source: https://adaptive-enforcement-lab.com/build/go-cli-architecture/kubernetes-integration/

Integrate your Go CLI with Kubernetes using client-go.

> **Universal Client**
>
> Build clients that work everywhere: developer laptops, CI runners, and cluster pods. Automatic config detection handles the differences.
>

---

## Overview

A well-designed Kubernetes CLI works seamlessly both on developer laptops and inside cluster pods. This section covers:

- **[Client Configuration](client-configuration.md)** - Automatic config detection for all environments
- **[RBAC Setup](rbac-setup.md)** - Service accounts and permissions
- **[Common Operations](common-operations/index.md)** - List, patch, and restart resources

---

## Configuration Flow

```mermaid
graph TB
    Start[Client Request] --> ExplicitKC{Explicit<br/>kubeconfig?}
    ExplicitKC -->|Yes| UseExplicit[Use Specified Path]
    ExplicitKC -->|No| InCluster{In-Cluster<br/>Token Exists?}

    InCluster -->|Yes| UseInCluster[Use In-Cluster Config]
    InCluster -->|No| EnvKC{KUBECONFIG<br/>Env Set?}

    EnvKC -->|Yes| UseEnv[Use KUBECONFIG Path]
    EnvKC -->|No| UseHome[Use ~/.kube/config]

    UseExplicit --> CreateClient[Create Clientset]
    UseInCluster --> CreateClient
    UseEnv --> CreateClient
    UseHome --> CreateClient

    CreateClient --> Ready[Client Ready]

    %% Start node - cyan

    %% Ghostty Hardcore Theme
    style Start fill:#65d9ef,color:#1b1d1e

    %% Decision nodes - orange
    style ExplicitKC fill:#fd971e,color:#1b1d1e
    style InCluster fill:#fd971e,color:#1b1d1e
    style EnvKC fill:#fd971e,color:#1b1d1e

    %% Config resolution nodes - purple
    style UseExplicit fill:#9e6ffe,color:#1b1d1e
    style UseInCluster fill:#9e6ffe,color:#1b1d1e
    style UseEnv fill:#9e6ffe,color:#1b1d1e
    style UseHome fill:#9e6ffe,color:#1b1d1e

    %% Processing and success nodes
    style CreateClient fill:#65d9ef,color:#1b1d1e
    style Ready fill:#a7e22e,color:#1b1d1e

```

---

## Quick Start

```go
import "k8s.io/client-go/kubernetes"

// Create a client that works everywhere
client, err := k8s.NewClient(kubeconfig, namespace)
if err != nil {
    return fmt.Errorf("failed to create client: %w", err)
}

// Use the client
deployments, err := client.ListDeployments(ctx)
```

---

## Best Practices

| Practice | Description |
| ---------- | ------------- |
| **Use contexts everywhere** | Pass `context.Context` to all Kubernetes operations |
| **Handle cancellation** | Respect context cancellation for clean shutdowns |
| **Wrap errors with context** | Include resource type and name in error messages |
| **Default to current namespace** | Match kubectl behavior for namespace resolution |
| **Support both configs** | Always handle in-cluster and out-of-cluster scenarios |
| **Minimal RBAC** | Request only the permissions your CLI needs |

---

*Build clients that work everywhere: laptop, CI runner, or pod.*
