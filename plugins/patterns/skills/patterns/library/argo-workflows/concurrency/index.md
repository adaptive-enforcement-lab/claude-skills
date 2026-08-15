# Concurrency Control

Source: https://adaptive-enforcement-lab.com/patterns/argo-workflows/concurrency/

When multiple workflows operate on shared resources, conflicts are inevitable. Two builds writing to the same output directory corrupt each other. Two deployments running simultaneously leave the system in an undefined state. Two cache rebuilds compete for the same ConfigMap.

Concurrency control prevents these conflicts. Argo Workflows provides several mechanisms: mutexes for exclusive access, semaphores for limited parallelism, and TTL strategies for cleanup.

---

## Why Concurrency Control Matters

Without concurrency control, the system behavior depends on timing. Sometimes workflows complete successfully. Sometimes they fail mysteriously. Sometimes they produce incorrect results that aren't detected until much later.

Consider a documentation build pipeline triggered by Git pushes. Developer A pushes a change and the build starts. Developer B pushes another change before A's build completes. Now two builds run simultaneously, both reading from and writing to the same directories.

Mutex synchronization ensures only one build runs at a time. B's build waits for A's to complete. The output is always consistent. The tradeoff is latency. B waits instead of starting immediately. But consistent results are worth more than fast chaos.

---

## Patterns

| Pattern | Description |
| --------- | ------------- |
| [Mutex Synchronization](mutex.md) | Exclusive access to shared resources |
| [Semaphores](semaphores.md) | Limited concurrent access |
| [TTL Strategy](ttl.md) | Automatic cleanup of completed workflows |

---

## Quick Start

1. **Identify shared resources** - What can only be accessed by one workflow at a time?
2. **Choose the right pattern** - Mutex for exclusive access, semaphore for limited parallelism
3. **Configure TTL** - Prevent unbounded growth of completed workflows
4. **Test under load** - Verify behavior when multiple workflows trigger simultaneously

---

> **Start with Mutex**
>
> When in doubt, start with a mutex. It's simpler to configure and debug. Only switch to semaphores when you need controlled parallelism.
>

---

## Related

- [WorkflowTemplate Patterns](../templates/index.md) - Basic workflow structure
- [Workflow Composition](../composition/index.md) - Parent/child workflow coordination
- [Scheduled Workflows](../scheduled/index.md) - CronWorkflow concurrency policies
