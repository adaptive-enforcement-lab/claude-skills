---
name: workflowtemplate-patterns
description: >-
  WorkflowTemplate foundations: versioned, reusable automation building blocks with parameter contracts, error handling, volumes, and RBAC for production workflows.
---

# WorkflowTemplate Patterns

## When to Use This Skill

WorkflowTemplates are the foundation of reusable automation in Argo Workflows. Rather than defining workflows inline or copying YAML between projects, WorkflowTemplates let you create versioned, tested building blocks that can be invoked by events, schedules, or other workflows.

---


## Implementation

1. **Define the template** with clear parameter contracts
2. **Add error handling** with retry strategies for transient failures
3. **Configure volumes** for data persistence and secrets
4. **Set up RBAC** with minimal required permissions

---

> **Start Simple**
>
> Begin with basic structure and retry strategy. Add init containers and custom volumes only when the simpler approach proves insufficient.
>

---


## Techniques


### Why WorkflowTemplates Matter

The naive approach to workflow automation is embedding all logic directly in the triggering resource: a Sensor, CronWorkflow, or manual submission. This works for simple cases but quickly becomes unmaintainable.

Consider a documentation build pipeline. The first version might be a simple script triggered by a GitHub push. But then you need the same build for scheduled refreshes. And manual triggers for debugging. And a "full rebuild" variant that processes all repositories instead of just the changed one.

Without WorkflowTemplates, you end up with four copies of nearly identical YAML. When you fix a bug or add a feature, you update one copy and forget the others. Drift accumulates. Debugging becomes archaeology.

WorkflowTemplates solve this by extracting the workflow logic into a standalone resource. Triggers reference the template by name. Updates happen in one place. The template becomes a contract: "give me these parameters, and I'll do this work."

---


### Patterns

| Pattern | Description |
| --------- | ------------- |
| [Basic Structure](basic-structure.md) | Fundamental WorkflowTemplate anatomy and parameter handling |
| [Retry Strategy](retry-strategy.md) | Error handling with exponential backoff |
| [Init Containers](init-containers.md) | Multi-stage pipelines with sequential setup |
| [Volume Patterns](volume-patterns.md) | Persistent storage, secrets, and configuration |
| [RBAC Configuration](rbac.md) | Security and permission management |

---


### Quick Start

1. **Define the template** with clear parameter contracts
2. **Add error handling** with retry strategies for transient failures
3. **Configure volumes** for data persistence and secrets
4. **Set up RBAC** with minimal required permissions

---

> **Start Simple**
>
> Begin with basic structure and retry strategy. Add init containers and custom volumes only when the simpler approach proves insufficient.
>

---


### Related

- [Concurrency Control](../concurrency/index.md) - Mutex synchronization and TTL strategies
- [Workflow Composition](../composition/index.md) - Child workflows and orchestration
- [Scheduled Workflows](../scheduled/index.md) - CronWorkflow patterns


## Anti-Patterns to Avoid

| Pattern | Description |
| --------- | ------------- |
| [Basic Structure](basic-structure.md) | Fundamental WorkflowTemplate anatomy and parameter handling |
| [Retry Strategy](retry-strategy.md) | Error handling with exponential backoff |
| [Init Containers](init-containers.md) | Multi-stage pipelines with sequential setup |
| [Volume Patterns](volume-patterns.md) | Persistent storage, secrets, and configuration |
| [RBAC Configuration](rbac.md) | Security and permission management |

---


## Related Patterns

- Basic Structure
- Retry Strategy
- Init Containers
- Volume Patterns
- RBAC Configuration

## References

- [Source Documentation](https://adaptive-enforcement-lab.com/patterns/argo-workflows/)
- [AEL Patterns](https://adaptive-enforcement-lab.com/patterns/)
