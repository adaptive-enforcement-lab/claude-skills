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



## Implementation


See the full implementation guide in the source documentation.


## Examples

See [examples.md](examples.md) for code examples.






## References

- [Source Documentation](https://adaptive-enforcement-lab.com/enforce/policy-as-code/)
- [AEL Enforce](https://adaptive-enforcement-lab.com/enforce/)
