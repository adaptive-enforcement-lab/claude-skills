---
name: kyverno-generation-templates
description: >-
  Kyverno generation policy templates that auto-create supporting resources like NetworkPolicies, ResourceQuotas, and PodDisruptionBudgets for new workloads and namespaces.
---

# Kyverno Generation Templates

## When to Use This Skill

**Use generation when:**

- You want to enforce security-by-default for all new resources
- Manual configuration creates gaps and inconsistencies
- You need automatic synchronization with changing requirements
- Supporting resources should follow workload lifecycle (create, update, delete)

**Do not use generation when:**

- Resources require unique, per-workload customization
- Generated resources would conflict with existing resources
- You need human approval before resource creation
- The triggering resource does not contain enough context to generate correctly

---


## Implementation

See the full implementation guide in the [source documentation](https://adaptive-enforcement-lab.com/enforce/policy-as-code/).


## Examples

See [examples.md](examples.md) for code examples.


## Related Patterns

- Kyverno Templates Overview
- Template Library Overview

## References

- [Source Documentation](https://adaptive-enforcement-lab.com/enforce/policy-as-code/)
- [AEL Enforce](https://adaptive-enforcement-lab.com/enforce/)
