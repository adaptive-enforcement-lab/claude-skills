---
name: concurrency-control
description: >-
  Prevent workflow conflicts with mutex synchronization, semaphores for limited parallelism, and TTL strategies for automatic cleanup of completed workflows.
---

# Concurrency Control

## When to Use This Skill

When multiple workflows operate on shared resources, conflicts are inevitable. Two builds writing to the same output directory corrupt each other. Two deployments running simultaneously leave the system in an undefined state. Two cache rebuilds compete for the same ConfigMap.

Concurrency control prevents these conflicts. Argo Workflows provides several mechanisms: mutexes for exclusive access, semaphores for limited parallelism, and TTL strategies for cleanup.

---


## Implementation

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


## Techniques


### Patterns

| Pattern | Description |
| --------- | ------------- |
| [Mutex Synchronization](mutex.md) | Exclusive access to shared resources |
| [Semaphores](semaphores.md) | Limited concurrent access |
| [TTL Strategy](ttl.md) | Automatic cleanup of completed workflows |

---


## Anti-Patterns to Avoid

| Pattern | Description |
| --------- | ------------- |
| [Mutex Synchronization](mutex.md) | Exclusive access to shared resources |
| [Semaphores](semaphores.md) | Limited concurrent access |
| [TTL Strategy](ttl.md) | Automatic cleanup of completed workflows |

---


## Related Patterns

- Mutex Synchronization
- Semaphores
- TTL Strategy

## References

- [Source Documentation](https://adaptive-enforcement-lab.com/patterns/argo-workflows/)
- [AEL Patterns](https://adaptive-enforcement-lab.com/patterns/)
