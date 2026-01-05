---
name: opa-resource-governance-templates
description: >-
  OPA resource governance policies enforcing CPU/memory limits, ResourceQuota compliance, LimitRange validation, and storage constraints.
---

# OPA Resource Governance Templates

## When to Use This Skill

Resource governance policies prevent overconsumption, enforce quotas, and validate LimitRange compliance across your cluster.

> **ResourceQuota vs LimitRange vs OPA**
>
> ResourceQuota caps total namespace consumption. LimitRange sets defaults and bounds for individual pods. OPA validates configuration before admission. Use all three for comprehensive governance.


## When to Apply

### Scenario 1: Prevent Unbounded Resource Consumption

Require resource limits on all containers:

```yaml
# Enforced by: governance.yaml
# Result: All containers must define resources.limits.cpu and resources.limits.memory
# Impact: Prevents single pod from consuming entire node capacity

```

### Scenario 2: Enforce LimitRange Compliance

Block pods exceeding namespace LimitRange maximums:

```yaml
# Enforced by: limitrange.yaml
# Result: Pods cannot request more CPU/memory than LimitRange allows
# Impact: Ensures fair resource distribution across namespace

```

### Scenario 3: Control Storage Costs

Restrict PVC sizes based on environment:

```yaml
# Enforced by: storage.yaml
# Result: PVCs in dev-* namespaces cannot exceed 50Gi
# Impact: Prevents accidental provisioning of expensive storage volumes

```


## Implementation

See the full implementation guide in the [source documentation](https://adaptive-enforcement-lab.com/enforce/policy-as-code/).


## Techniques


### Resource Governance Patterns

### Three-Layer Resource Control

Implement overlapping controls for comprehensive governance:

1. **OPA Policies** - Validate resource configuration at admission time
2. **LimitRange** - Set namespace-level defaults and max values
3. **ResourceQuota** - Cap total namespace consumption

### Right-Sizing Workloads

Balance cost and reliability with appropriate resource values:

| Resource | Too Low | Too High | Sweet Spot |
|----------|---------|----------|------------|
| **Requests** | OOMKilled pods | Wasted capacity | Actual usage (P95) |
| **Limits** | Throttling | Noisy neighbors | 2x requests |
| **Storage** | Out of space | High costs | Actual data + 30% |

Use Vertical Pod Autoscaler (VPA) to identify optimal values.

### Quota Enforcement Strategy

Define quotas based on team size and workload type:

```yaml
# Development namespace quota
apiVersion: v1
kind: ResourceQuota
metadata:
  name: dev-quota
  namespace: dev-team
spec:
  hard:
    requests.cpu: "10"
    requests.memory: "20Gi"
    limits.cpu: "20"
    limits.memory: "40Gi"
    persistentvolumeclaims: "10"
    requests.storage: "100Gi"

```

OPA policies validate workloads against these quotas before admission.

*See [reference.md](reference.md) for additional techniques and detailed examples.*


## Examples

See [examples.md](examples.md) for code examples.


## Full Reference

See [reference.md](reference.md) for complete documentation.


## Related Patterns

- OPA Templates Overview
- OPA RBAC Policies
- Kyverno Resource Governance

## References

- [Source Documentation](https://adaptive-enforcement-lab.com/enforce/policy-as-code/)
- [AEL Enforce](https://adaptive-enforcement-lab.com/enforce/)
