---
name: opa-pod-security-templates
description: >-
  OPA pod security policies preventing privileged containers, restricting Linux capabilities, and enforcing security contexts in Kubernetes.
---

# OPA Pod Security Templates

## When to Use This Skill

Pod security policies written in Rego prevent privilege escalation and enforce security boundaries for containerized workloads.

> **Capabilities Bypass Security Boundaries**
>
> Linux capabilities grant fine-grained privileges. A container with `CAP_SYS_ADMIN` can bypass most kernel security mechanisms. Drop all capabilities by default.



## Implementation


See the full implementation guide in the source documentation.


## Examples

See [examples.md](examples.md) for code examples.






## References

- [Source Documentation](https://adaptive-enforcement-lab.com/enforce/policy-as-code/)
- [AEL Enforce](https://adaptive-enforcement-lab.com/enforce/)
