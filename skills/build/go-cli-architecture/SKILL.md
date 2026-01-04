---
name: go-cli-architecture
description: >-
  Build Kubernetes-native CLIs in Go with type safety, testability, and complex orchestration logic for deployment tools and cluster automation.
---

# Go CLI Architecture

## When to Use This Skill

> **Go vs Shell Scripts**
>
>
> Start with shell scripts for prototyping. Graduate to Go when you need type safety, testability, or complex orchestration logic.
>

**Use Go when you need:**

- Direct Kubernetes API access with type-safe clients
- Complex orchestration logic across multiple resources
- Reusable tooling packaged as container images
- Performance-critical operations (milliseconds matter)
- Long-running controllers or operators

**Use shell scripts when you need:**

- Simple glue logic between existing tools
- Quick prototypes or one-off operations
- kubectl-based workflows without custom logic
- CI/CD steps that primarily call other CLIs

---



## Implementation


See the full implementation guide in the source documentation.


## Examples

See [examples.md](examples.md) for code examples.






## References

- [Source Documentation](https://adaptive-enforcement-lab.com/build/go-cli-architecture/)
- [AEL Build](https://adaptive-enforcement-lab.com/build/)
