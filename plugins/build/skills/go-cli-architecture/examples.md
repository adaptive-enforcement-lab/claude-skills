---
name: go-cli-architecture - Examples
description: Code examples for Go CLI Architecture
---

# Go CLI Architecture - Examples


## Example 1: example-1.mermaid


```mermaid
graph TB
    subgraph CLI["CLI Layer"]
        Root[root.go<br/>Cobra Root Command]
        Orch[orchestrate.go<br/>Workflow Entry Point]
        Sub1[check.go<br/>Subcommand]
        Sub2[rebuild.go<br/>Subcommand]
        Sub3[select.go<br/>Subcommand]
    end

    subgraph Pkg["Business Logic Layer"]
        Cache[pkg/cache<br/>Cache Management]
        K8s[pkg/k8s<br/>Client Wrapper]
        Selector[pkg/selector<br/>Business Logic]
        Restarter[pkg/restarter<br/>Deployment Logic]
    end

    subgraph External["External Systems"]
        API[Kubernetes API]
        Argo[Argo Workflows]
        Store[Cache Store]
    end

    Root --> Orch
    Orch --> Sub1
    Orch --> Sub2
    Orch --> Sub3
    Sub1 --> Cache
    Sub2 --> Cache
    Sub3 --> Selector
    Selector --> Restarter
    Cache --> K8s
    Selector --> K8s
    Restarter --> K8s
    K8s --> API
    K8s --> Argo
    Cache --> Store

    %% CLI Layer nodes - cyan

    %% Ghostty Hardcore Theme
    style Root fill:#65d9ef,color:#1b1d1e
    style Orch fill:#65d9ef,color:#1b1d1e
    style Sub1 fill:#65d9ef,color:#1b1d1e
    style Sub2 fill:#65d9ef,color:#1b1d1e
    style Sub3 fill:#65d9ef,color:#1b1d1e

    %% Business Logic Layer nodes - green
    style Cache fill:#a7e22e,color:#1b1d1e
    style K8s fill:#a7e22e,color:#1b1d1e
    style Selector fill:#a7e22e,color:#1b1d1e
    style Restarter fill:#a7e22e,color:#1b1d1e

    %% External Systems nodes - purple
    style API fill:#9e6ffe,color:#1b1d1e
    style Argo fill:#9e6ffe,color:#1b1d1e
    style Store fill:#9e6ffe,color:#1b1d1e
```



## Example 2: example-2.text


```text
myctl/
├── cmd/
│   ├── root.go           # Cobra root command, global flags
│   ├── orchestrate.go    # Main workflow orchestrator
│   ├── check.go          # Cache check command
│   ├── rebuild.go        # Cache rebuild command
│   └── select.go         # Deployment selector
├── pkg/
│   ├── cache/            # Cache management logic
│   │   ├── cache.go
│   │   └── cache_test.go
│   ├── k8s/              # Kubernetes client wrapper
│   │   ├── client.go
│   │   └── client_test.go
│   ├── selector/         # Business logic
│   │   ├── selector.go
│   │   └── selector_test.go
│   └── restarter/        # Deployment restart logic
│       ├── restarter.go
│       └── restarter_test.go
├── Dockerfile
├── go.mod
├── go.sum
└── main.go               # Entry point
```



