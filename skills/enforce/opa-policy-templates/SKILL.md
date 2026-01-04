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


See the full implementation guide in the source documentation.


## Examples

See [examples.md](examples.md) for code examples.






## References

- [Source Documentation](https://adaptive-enforcement-lab.com/enforce/policy-as-code/)
- [AEL Enforce](https://adaptive-enforcement-lab.com/enforce/)
