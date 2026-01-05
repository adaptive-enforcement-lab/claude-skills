---
name: workflow-composition
description: >-
  Build complex pipelines from reusable workflow components. Compose parent-child workflows, orchestrate multi-stage automation, and eliminate copy-paste YAML.
---

# Workflow Composition

## When to Use This Skill

As automation pipelines grow, a single monolithic workflow becomes unmaintainable. Composition patterns let you build complex pipelines from smaller, reusable pieces. A parent workflow can spawn children, wait for their completion, and orchestrate the overall flow.

---


## Implementation

1. **Extract reusable logic** into separate WorkflowTemplates
2. **Create a parent workflow** that spawns children
3. **Define success/failure conditions** for proper status propagation
4. **Test each child independently** before composing

---

> **Test Children First**
>
> Always test child workflows independently before composing them into a parent. Debugging failures in composed workflows is much harder than debugging standalone workflows.
>

---


## Techniques


### Patterns

| Pattern | Description |
| --------- | ------------- |
| [Spawning Child Workflows](spawning-children.md) | Create and wait for child workflow completion |
| [Parallel Execution](parallel.md) | Run multiple workflows simultaneously |
| [DAG Orchestration](dag.md) | Dependency-based execution ordering |
| [Cross-Workflow Communication](communication.md) | Passing data and triggering decoupled workflows |

---


## Anti-Patterns to Avoid

| Pattern | Description |
| --------- | ------------- |
| [Spawning Child Workflows](spawning-children.md) | Create and wait for child workflow completion |
| [Parallel Execution](parallel.md) | Run multiple workflows simultaneously |
| [DAG Orchestration](dag.md) | Dependency-based execution ordering |
| [Cross-Workflow Communication](communication.md) | Passing data and triggering decoupled workflows |

---


## Related Patterns

- Spawning Child Workflows
- Parallel Execution
- DAG Orchestration
- Cross-Workflow Communication

## References

- [Source Documentation](https://adaptive-enforcement-lab.com/patterns/argo-workflows/)
- [AEL Patterns](https://adaptive-enforcement-lab.com/patterns/)
