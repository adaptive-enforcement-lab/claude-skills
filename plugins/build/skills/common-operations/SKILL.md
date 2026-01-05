---
name: common-operations
description: >-
  Implement idiomatic Kubernetes operations with label selectors, strategic merge patches, and proper error handling for production-grade CLI tooling.
---

# Common Operations

## When to Use This Skill

A well-designed Kubernetes CLI provides idiomatic operations that work consistently across resource types. This section covers:

- **[List Resources](list-resources.md)** - Query resources with label selectors
- **[Rollout Restart](rollout-restart.md)** - Trigger rolling restarts without downtime
- **[ConfigMap Operations](configmap-operations.md)** - Store and retrieve configuration data
- **[Watch Resources](watch-resources.md)** - React to real-time resource changes

---


## Implementation

See the full implementation guide in the [source documentation](https://adaptive-enforcement-lab.com/build/go-cli-architecture/).


## Key Principles

| Practice | Description |
| ---------- | ------------- |
| **Use label selectors** | Filter resources server-side, not client-side |
| **Prefer patches over updates** | Patches are safer for concurrent modifications |
| **Use strategic merge patches** | Kubernetes-native patch format for resources |
| **Handle not found errors** | Check `apierrors.IsNotFound(err)` before creating |
| **Respect resource versions** | Use optimistic concurrency for updates |

---

*Use the Kubernetes API idiomatically: label selectors, patches, and proper error handling.*


## Techniques


### Operation Patterns

```mermaid
graph TB
    CLI[CLI Command] --> List[List Resources]
    CLI --> Mutate[Mutate Resources]
    CLI --> Watch[Watch Changes]

    List --> Filter[Label Selectors]
    Mutate --> Patch[Strategic Merge Patch]
    Mutate --> Create[Get-or-Create]
    Watch --> Events[Event Stream]

    %% Ghostty Hardcore Theme
    style CLI fill:#65d9ef,color:#1b1d1e
    style List fill:#a7e22e,color:#1b1d1e
    style Mutate fill:#fd971e,color:#1b1d1e
    style Watch fill:#9e6ffe,color:#1b1d1e
    style Filter fill:#5e7175,color:#f8f8f3
    style Patch fill:#5e7175,color:#f8f8f3
    style Create fill:#5e7175,color:#f8f8f3
    style Events fill:#5e7175,color:#f8f8f3

```

---
## References

- [Source Documentation](https://adaptive-enforcement-lab.com/build/go-cli-architecture/)
- [AEL Build](https://adaptive-enforcement-lab.com/build/)
