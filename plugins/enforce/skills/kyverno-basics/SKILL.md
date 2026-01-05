---
name: kyverno-basics
description: >-
  Install Kyverno, create validation policies, and understand audit vs enforce modes for Kubernetes admission control.
---

# Kyverno Basics

## When to Use This Skill

Kyverno runs as a dynamic admission controller in Kubernetes. It validates, mutates, and generates resources based on policies written in YAML.

---


## Implementation

Install Kyverno using Helm:


*See [examples.md](examples.md) for detailed code examples.*

Kyverno creates webhook configurations that intercept resource creation/updates before they reach etcd.

---


## Comparison

Roll out policies in audit mode first:

```yaml
spec:
  validationFailureAction: Audit  # Log violations, don't block
```

Check logs for violations:

```bash
kubectl get policyreport -A

NAMESPACE   NAME                          PASS   FAIL   WARN   ERROR   SKIP
default     polr-ns-default              12     3      0      0       0
production  polr-ns-production           45     1      0      0       0
```

Fix violations. Then switch to Enforce:

```yaml
spec:
  validationFailureAction: Enforce  # Block violations
```

### Gradual Rollout Strategy

1. Deploy policy in `Audit` mode
2. Monitor PolicyReports for 1 week
3. Remediate failures
4. Switch to `Enforce` mode
5. Handle exceptions with exclusions

Don't deploy straight to Enforce. Discover violations first.

---


## Examples

See [examples.md](examples.md) for code examples.


## Related Patterns

- Policy Patterns
- Testing and Exceptions
- CI/CD Integration

## References

- [Source Documentation](https://adaptive-enforcement-lab.com/enforce/policy-as-code/)
- [AEL Enforce](https://adaptive-enforcement-lab.com/enforce/)
