---
name: jmespath-for-kyverno
description: >-
  Master JMESPath for Kyverno policies. Query nested resources, build complex conditions, and validate Kubernetes workloads with production-tested patterns.
---

# JMESPath for Kyverno

## When to Use This Skill

**Use JMESPath when:**

- Pattern matching can't express your logic
- You need conditionals or transformations
- Validation depends on multiple fields
- You're filtering or comparing arrays

**Skip JMESPath when:**

- Simple pattern matching works (`pattern`, `anyPattern`)
- You're only checking field existence
- No cross-field validation needed

> **Test Before Deploying**
>
> Always test JMESPath expressions with `kyverno jp` before adding them to policies. Syntax errors fail silently in audit mode and block resources in enforce mode.
>

---



## Implementation


See the full implementation guide in the source documentation.


## Examples

See [examples.md](examples.md) for code examples.






## References

- [Source Documentation](https://adaptive-enforcement-lab.com/enforce/policy-as-code/)
- [AEL Enforce](https://adaptive-enforcement-lab.com/enforce/)
