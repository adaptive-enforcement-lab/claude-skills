---
name: opa-policy-templates
description: >-
  OPA Gatekeeper policy templates overview. 20 production-ready constraint templates for pod security, image validation, RBAC, and resource governance.
---

# OPA Policy Templates

## When to Use This Skill

> **Deploy in Audit Mode First**
>
> Use `enforcementAction: dryrun` initially. Existing resources may violate constraints. Monitor violations for 48 hours using `kubectl get constraints`, fix non-compliant resources, then switch to `deny`.
>

Production-ready OPA/Gatekeeper constraint templates for Kubernetes admission control. **20 policies** covering pod security, image
validation, RBAC, and resource governance. Each template includes complete Rego implementation, constraint examples, customization
options, validation commands, and real-world use cases.

---


## Implementation

Standard deployment workflow for all templates:


*See [examples.md](examples.md) for detailed code examples.*

---


## Comparison

Choosing between OPA/Gatekeeper and Kyverno depends on your team's expertise and requirements:

### Use OPA/Gatekeeper When

- You need **maximum flexibility** in policy logic (Rego is Turing-complete)
- Your team has **Rego expertise** or investment in OPA across multiple systems
- You require **cross-platform policy** (Kubernetes, Terraform, Envoy, etc.)
- Policies involve **complex conditional logic** or multi-resource validation
- You're building a **policy platform** for enterprise governance

### Use Kyverno When

- You want **Kubernetes-native YAML** policies (no DSL learning curve)
- You need **mutation and generation** features (OPA is validation-only)
- Your team prefers **JMESPath** over Rego for data extraction
- You want **faster time-to-value** with simpler policies
- You're **new to policy-as-code** and want quick adoption

**See [Decision Guide →](../decision-guide.md)** for detailed comparison and migration strategies.

---


## Examples

See [examples.md](examples.md) for code examples.


## Full Reference

See [reference.md](reference.md) for complete documentation.


## Related Patterns

- Kyverno Templates →
- Decision Guide →
- OPA/Kyverno Comparison →
- Migration Guide →
- Template Library Overview →

## References

- [Source Documentation](https://adaptive-enforcement-lab.com/enforce/policy-as-code/)
- [AEL Enforce](https://adaptive-enforcement-lab.com/enforce/)
