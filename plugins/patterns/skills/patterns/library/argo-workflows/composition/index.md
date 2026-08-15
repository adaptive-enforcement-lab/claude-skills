# Workflow Composition

Source: https://adaptive-enforcement-lab.com/patterns/argo-workflows/composition/

As automation pipelines grow, a single monolithic workflow becomes unmaintainable. Composition patterns let you build complex pipelines from smaller, reusable pieces. A parent workflow can spawn children, wait for their completion, and orchestrate the overall flow.

---

## Why Compose Workflows?

The obvious approach to building a multi-stage pipeline is putting everything in one WorkflowTemplate. Clone the repository, run tests, build the artifact, deploy to staging, run integration tests, and promote to production. All of this runs as a single workflow with sequential steps.

This works until it doesn't.

The problems emerge gradually. First, you need the same build step in a different pipeline. So you copy it. Now you have two copies. Then someone fixes a bug in one copy but forgets the other.

Then you need to run just the build step for debugging. But you can't, because it's entangled with everything else.

Composition solves this by treating workflows as functions. Each workflow does one thing well. Parent workflows orchestrate the pieces. When you need the build step elsewhere, you call it. When you need to test the build step in isolation, you run it directly.

The tradeoff is complexity. Composed workflows have more moving parts, more YAML files, more potential failure points. Use composition when the benefits of reusability outweigh the costs of coordination.

---

## Patterns

| Pattern | Description |
| --------- | ------------- |
| [Spawning Child Workflows](spawning-children.md) | Create and wait for child workflow completion |
| [Parallel Execution](parallel.md) | Run multiple workflows simultaneously |
| [DAG Orchestration](dag.md) | Dependency-based execution ordering |
| [Cross-Workflow Communication](communication.md) | Passing data and triggering decoupled workflows |

---

## Quick Start

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

## Related

- [WorkflowTemplate Patterns](../templates/index.md) - Building the components to compose
- [Concurrency Control](../concurrency/index.md) - Preventing conflicts between composed workflows
- [Scheduled Workflows](../scheduled/index.md) - Time-based orchestration
