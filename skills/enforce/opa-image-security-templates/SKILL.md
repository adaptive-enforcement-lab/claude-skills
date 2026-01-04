---
name: opa-image-security-templates
description: >-
  OPA image security policies for container registry allowlisting, digest enforcement, and signature verification in Kubernetes.
---

# OPA Image Security Templates

## When to Use This Skill

Image security policies control which container images can run in your cluster. These templates enforce registry allowlists, require immutable digests, and validate cryptographic signatures.

> **Image Tags Are Mutable**
>
> Tags like `latest` or `v1.2.3` can be overwritten by attackers who compromise registries. Use digest-based references (`sha256:...`) for immutable deployments.



## Implementation


See the full implementation guide in the source documentation.


## Examples

See [examples.md](examples.md) for code examples.






## References

- [Source Documentation](https://adaptive-enforcement-lab.com/enforce/policy-as-code/)
- [AEL Enforce](https://adaptive-enforcement-lab.com/enforce/)
