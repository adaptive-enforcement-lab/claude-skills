---
name: policy-packaging
description: >-
  Build policy-platform containers that aggregate Kyverno policies from multiple OCI sources. Multi-stage Docker builds for local dev and CI validation.
---

# Policy Packaging

## When to Use This Skill

The policy-platform container is a multi-stage Docker build that:

1. Pulls policy repositories as OCI containers
2. Installs policy validation tools (Kyverno, Pluto, Spectral, Helm)
3. Aggregates everything into a single distributable image

> **One Container, All Policies**
>
> The policy-platform image runs identically in local dev, CI pipelines, and reference environments. Zero configuration drift.
>

---



## Implementation


See the full implementation guide in the source documentation.


## Examples

See [examples.md](examples.md) for code examples.






## References

- [Source Documentation](https://adaptive-enforcement-lab.com/enforce/policy-as-code/)
- [AEL Enforce](https://adaptive-enforcement-lab.com/enforce/)
