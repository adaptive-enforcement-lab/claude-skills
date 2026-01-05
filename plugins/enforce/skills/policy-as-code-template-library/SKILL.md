---
name: policy-as-code-template-library
description: >-
  Production-ready policy template library: 48 battle-tested policies (28 Kyverno + 20 OPA). Pod security, image validation, RBAC, resource governance, network security, and CI/CD integration.
---

# Policy-as-Code Template Library

## When to Use This Skill

**48 production-ready policies** for Kubernetes security and governance. Reduce the Rego learning curve. Copy, customize, deploy.

<!-- more -->

> **Template Library Overview**
>
> This library contains **28 Kyverno policies** and **20 OPA/Gatekeeper constraint templates** covering pod security, image validation, RBAC, resource governance, network security, mutation, and generation.
> Each template includes complete YAML/Rego, customization variables, validation commands, and real-world use cases.
>

---


## Implementation

> **Deploy in Audit Mode First**
>
> Always start with `audit` (Kyverno) or `dryrun` (OPA) mode. Monitor violations for 48 hours before switching to enforcement. Existing workloads may violate policies.
>

### Kyverno Quick Start (5 minutes)


*See [examples.md](examples.md) for detailed code examples.*

### OPA/Gatekeeper Quick Start (10 minutes)


*See [examples.md](examples.md) for detailed code examples.*

---


## Comparison

Choose the right policy engine for your team:

| Feature | Kyverno | OPA/Gatekeeper |
|---------|---------|----------------|
| **Policies** | 28 (validation, mutation, generation) | 20 (validation only) |
| **Language** | YAML + JMESPath | Rego (Go-like DSL) |
| **Learning Curve** | < 1 hour | 4-8 hours |
| **Best For** | Kubernetes-native teams, fast adoption | Multi-platform policies, complex logic |
| **Mutation** | ✅ Native support | ❌ Validation only |
| **Generation** | ✅ Auto-create resources | ❌ Validation only |

**See [Decision Guide →](decision-guide.md)** for detailed comparison and recommended starter paths.

---


## Examples

See [examples.md](examples.md) for code examples.


## Full Reference

See [reference.md](reference.md) for complete documentation.


## Related Patterns

- Kyverno Official Documentation
- OPA/Gatekeeper Documentation
- Kubernetes Pod Security Standards
- NIST SP 800-190
- CIS Kubernetes Benchmark

## References

- [Source Documentation](https://adaptive-enforcement-lab.com/enforce/policy-as-code/)
- [AEL Enforce](https://adaptive-enforcement-lab.com/enforce/)
