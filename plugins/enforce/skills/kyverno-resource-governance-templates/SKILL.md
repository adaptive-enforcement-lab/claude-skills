---
name: kyverno-resource-governance-templates
description: >-
  Kyverno resource governance policies enforcing CPU/memory limits, HPA requirements, and storage constraints for Kubernetes workloads.
---

# Kyverno Resource Governance Templates

## When to Use This Skill

Resource governance policies prevent overconsumption, enforce autoscaling requirements, and control storage allocation across your cluster.

> **Resource Limits Prevent Noisy Neighbors**
>
> Without resource limits, a single pod can consume all node capacity and starve other workloads. Enforce limits to guarantee fair resource allocation.


## When to Apply

### Scenario 1: Prevent Unbounded Resource Consumption

Require resource limits on all containers:

```yaml
# Enforced by: limits.yaml
# Result: All containers must define resources.limits.cpu and resources.limits.memory
# Impact: Prevents single pod from consuming entire node capacity
```

### Scenario 2: Mandate Autoscaling for Production

Require HPA for production Deployments:

```yaml
# Enforced by: hpa.yaml
# Result: Deployments in prod-* namespaces must have corresponding HPA
# Impact: Ensures production services scale automatically under load
```

### Scenario 3: Control Storage Costs

Restrict PVC size to prevent excessive allocations:

```yaml
# Enforced by: storage.yaml
# Result: PVCs cannot exceed 100Gi in dev namespaces
# Impact: Prevents accidental provisioning of expensive storage volumes
```


## Implementation

See the full implementation guide in the [source documentation](https://adaptive-enforcement-lab.com/enforce/policy-as-code/).


## Techniques


### Resource Management Patterns

### Resource Quotas vs Limits

Use both mechanisms for defense in depth:

- **ResourceQuota** - Namespace-level caps (total CPU/memory across all pods)
- **LimitRange** - Default and max values for individual pods
- **Kyverno Policies** - Validation and enforcement of resource configuration

Kyverno policies complement quotas by validating workload-level configuration before admission.

### Right-Sizing Workloads

Set appropriate resource values to balance cost and reliability:

- **Requests too low** → Pods scheduled on undersized nodes → OOMKilled
- **Requests too high** → Wasted capacity → Increased costs
- **Limits too low** → Pods throttled → Performance degradation
- **Limits too high** → Noisy neighbor problems → Node instability

Use Vertical Pod Autoscaler (VPA) recommendations to identify optimal values.

### Autoscaling Strategies

Choose the right autoscaling mechanism for your workload:

- **HPA (Horizontal)** - Scale replicas based on CPU/memory/custom metrics
- **VPA (Vertical)** - Adjust resource requests/limits automatically
- **Cluster Autoscaler** - Add/remove nodes based on pending pods

Kyverno policies enforce HPA presence and configuration validity.

*See [reference.md](reference.md) for additional techniques and detailed examples.*


## Examples

See [examples.md](examples.md) for code examples.


## Full Reference

See [reference.md](reference.md) for complete documentation.


## Related Patterns

- Kyverno Templates Overview
- Kyverno Pod Security
- OPA Resource Governance

## References

- [Source Documentation](https://adaptive-enforcement-lab.com/enforce/policy-as-code/)
- [AEL Enforce](https://adaptive-enforcement-lab.com/enforce/)
