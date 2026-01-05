---
name: argo-workflows-patterns
description: >-
  Production Argo Workflows patterns: reusable templates, error handling, concurrency control, composition, and scheduled automation for Kubernetes operators.
---

# Argo Workflows Patterns

## When to Use This Skill

Production patterns for Argo Workflows: reusable templates, error handling, concurrency control, workflow composition, and scheduled automation.

---


## Implementation

1. **Define WorkflowTemplates** - Create reusable, tested building blocks
2. **Add Error Handling** - Configure retry strategies for transient failures
3. **Control Concurrency** - Use mutexes or semaphores for shared resources
4. **Compose Workflows** - Chain templates into complex pipelines
5. **Schedule Automation** - Run workflows on cron schedules

---


## Techniques


### Why Argo Workflows?

Kubernetes provides primitives (Pods, Jobs, CronJobs), but building complex automation from primitives is painful. You end up with shell scripts that check Pod status in loops, cleanup logic scattered across multiple places, and debugging that requires correlating logs from dozens of sources.

Argo Workflows provides higher-level abstractions designed for automation. Define workflows declaratively. Let the controller handle scheduling, retries, and cleanup. Visualize execution in a purpose-built UI. Focus on what the automation does, not how to orchestrate it.

---


### Pattern Categories

| Category | Description |
| ---------- | ------------- |
| [WorkflowTemplate Patterns](templates/index.md) | Reusable workflow definitions with error handling, volumes, and RBAC |
| [Concurrency Control](concurrency/index.md) | Mutex synchronization, semaphores, and TTL strategies |
| [Workflow Composition](composition/index.md) | Parent/child workflows, orchestration, and cross-workflow communication |
| [Scheduled Workflows](scheduled/index.md) | CronWorkflow patterns and GitHub integration |

---


### Quick Start

1. **Define WorkflowTemplates** - Create reusable, tested building blocks
2. **Add Error Handling** - Configure retry strategies for transient failures
3. **Control Concurrency** - Use mutexes or semaphores for shared resources
4. **Compose Workflows** - Chain templates into complex pipelines
5. **Schedule Automation** - Run workflows on cron schedules

---


### Troubleshooting

### Workflow Stuck in Pending

1. Check service account permissions: `kubectl describe rolebinding -n argo-workflows`
2. Verify resource quotas: `kubectl describe quota -n argo-workflows`
3. Check node resources: `kubectl top nodes`
4. Look for mutex waits: `kubectl get workflows -l workflows.argoproj.io/sync-id`

### Workflow Failed with RBAC Error

1. Verify ServiceAccount exists in workflow namespace
2. Check ClusterRoleBinding subjects match namespace
3. Use `kubectl auth can-i` to test permissions:

```bash
kubectl auth can-i patch deployments \
  --as=system:serviceaccount:argo-workflows:my-sa \
  -n target-namespace
```

### Mutex Deadlock

1. Find workflows waiting on mutex: `kubectl get workflows -l workflows.argoproj.io/sync-id`
2. Identify the workflow holding the lock
3. Check if the holding workflow is stuck or failed
4. Terminate stuck workflows to release mutex

---

> **Prerequisites**
>
> Argo Workflows must be installed in your cluster. See the [official installation guide](https://argo-workflows.readthedocs.io/en/latest/quick-start/) for setup instructions.
>

---


### Related

- [Argo Events Setup](../../patterns/argo-events/setup/index.md) - EventSource, EventBus, and Sensor configuration
- [ConfigMap as Cache](../efficiency/idempotency/caches.md) - Volume mounts for zero-API reads
- [Event-Driven Deployments](../../blog/posts/2025-12-14-event-driven-deployments-argo.md) - The journey to zero-latency automation


## Troubleshooting

See [troubleshooting.md](troubleshooting.md) for common issues and solutions.


## Related Patterns

- Argo Events Setup
- ConfigMap as Cache
- Event-Driven Deployments

## References

- [Source Documentation](https://adaptive-enforcement-lab.com/patterns/argo-workflows/)
- [AEL Patterns](https://adaptive-enforcement-lab.com/patterns/)
