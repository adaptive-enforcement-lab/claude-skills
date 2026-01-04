---
name: opa-rbac-templates
description: >-
  OPA RBAC policies preventing cluster-admin privilege escalation, restricting privileged verbs, and blocking wildcard permissions.
---

# OPA RBAC Templates

## When to Use This Skill

RBAC policies control who can perform which actions on which resources. These templates prevent privilege escalation through overly permissive roles.

> **Wildcards Grant Unrestricted Access**
>
> RBAC rules with `resources: ["*"]` or `verbs: ["*"]` grant access to all current and future resources or actions. Avoid wildcards except for break-glass admin roles.



## Implementation


See the full implementation guide in the source documentation.


## Examples

See [examples.md](examples.md) for code examples.






## References

- [Source Documentation](https://adaptive-enforcement-lab.com/enforce/policy-as-code/)
- [AEL Enforce](https://adaptive-enforcement-lab.com/enforce/)
