---
name: kyverno-mutation-templates
description: >-
  Kyverno mutation policies that auto-inject labels, sidecars, and configuration into Kubernetes workloads at admission time.
---

# Kyverno Mutation Templates

## When to Use This Skill

Mutation policies modify resources at admission time, before they're persisted to etcd. This approach enforces standards without blocking deployments or requiring manual manifest updates.


## Implementation

See the full implementation guide in the [source documentation](https://adaptive-enforcement-lab.com/enforce/policy-as-code/).


## Techniques


### Common Patterns

### Conditional Mutations

Only mutate resources that match specific criteria:

- Namespace-scoped mutations (dev vs prod)
- Label-based mutations (inject monitoring only for `app.kubernetes.io/monitored=true`)
- Resource type mutations (different rules for Deployments vs StatefulSets)

### Mutation Conflicts

When multiple policies mutate the same field:

- **Last-write-wins** - Policies execute in alphabetical order by name
- **Merge strategies** - Use `patchStrategicMerge` or `patchesJson6902` for predictable merging
- **Exclusions** - Use `exclude` blocks to prevent conflicting mutations

### Security Boundaries

Never mutate security-critical fields:

- Security contexts (runAsUser, capabilities, privileged)
- Resource limits (mutations can escalate privileges)
- Host paths or volumes (mutations can grant filesystem access)

Use validation policies for security boundaries. Use mutations for operational standards.


## Examples

See [examples.md](examples.md) for code examples.


## Related Patterns

- Kyverno Templates Overview
- Kyverno Generation Templates
- Kyverno Image Security

## References

- [Source Documentation](https://adaptive-enforcement-lab.com/enforce/policy-as-code/)
- [AEL Enforce](https://adaptive-enforcement-lab.com/enforce/)
