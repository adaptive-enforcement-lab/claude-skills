# Kyverno Mutation Templates

Source: https://adaptive-enforcement-lab.com/enforce/policy-as-code/template-library/kyverno/mutation/

Mutation policies modify resources at admission time, before they're persisted to etcd. This approach enforces standards without blocking deployments or requiring manual manifest updates.

## Why Mutation Over Validation

Validation blocks non-compliant resources. Mutation fixes them automatically.

> **Fix vs Block**
>
> Mutations reduce friction by auto-correcting resources at admission time. This approach enforces standards without breaking deployments or requiring manual updates.
>

**Use mutations when:**

- Adding required labels or annotations to all workloads
- Injecting sidecars for logging, monitoring, or security
- Setting default resource limits or security contexts
- Enforcing organizational standards that shouldn't block deployments

**Use validation when:**

- Security boundaries must never be crossed (privileged containers, host paths)
- Resource constraints are non-negotiable (quotas, limits)
- Audit requirements demand explicit opt-in (PII handling, compliance tags)

## Available Templates

### [Label Injection](labels.md)

Auto-inject required labels and annotations into workloads:

- Default organizational labels (team, environment, cost center)
- Conditional label injection based on namespace or existing labels
- Label propagation from namespaces to pods

**Apply a policy:**

```bash
kubectl apply -f labels.yaml
```

### [Sidecar Injection](sidecar.md)

Auto-inject sidecar containers for observability and security:

- Logging sidecars (Fluent Bit, Fluentd)
- Monitoring agents (Prometheus exporters, metrics collectors)
- Security sidecars (secret managers, policy enforcers)

**Apply a policy:**

```bash
kubectl apply -f sidecar.yaml
```

## Mutation Execution Order

Kyverno processes mutations in this order:

1. **Mutating Admission Webhooks** - Kyverno intercepts CREATE/UPDATE requests
2. **Policy Evaluation** - Matches resource against mutation rules
3. **Mutation Application** - Modifies resource in memory
4. **Validation** - Resource proceeds to validation policies (if any)
5. **Persistence** - Modified resource written to etcd

**Critical:** Mutations only apply to CREATE/UPDATE operations. Set `background: false` to prevent mutations from applying to existing resources during policy sync.

## Testing Mutations

Use `kubectl apply --dry-run=server` to test mutations without creating resources:

```bash
# Test label mutation
kubectl apply --dry-run=server -f test-deployment.yaml -o yaml | grep -A5 labels

# Test sidecar injection
kubectl apply --dry-run=server -f test-pod.yaml -o yaml | grep -A10 containers
```

## Common Patterns

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

## Related Resources

- [Kyverno Templates Overview](../index.md)
- [Kyverno Generation Templates](../generation/index.md)
- [Kyverno Image Security](../image/index.md)
