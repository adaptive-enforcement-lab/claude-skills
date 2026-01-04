---
name: kubernetes-integration - Examples
description: Code examples for Kubernetes Integration
---

# Kubernetes Integration - Examples


## Example 1: example-1.mermaid


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



## Example 2: example-2.go


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



