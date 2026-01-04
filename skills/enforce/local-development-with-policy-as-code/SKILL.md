---
name: local-development-with-policy-as-code
description: >-
  Run Kyverno policy validation locally with policy-platform containers. Zero local setup, same tools as CI, instant feedback before committing code changes.
---

# Local Development with Policy-as-Code

## When to Use This Skill

The policy-platform container includes all tools needed for local policy validation:

- **Kyverno CLI** - Policy validation and testing
- **Pluto** - Deprecated API detection
- **Helm** - Chart rendering and linting
- **Spectral** - OpenAPI/values schema validation
- **yq** - YAML processing

> **Zero Local Setup Required**
>
> One container contains all policies and tools. No local installations. Pull the container, run validations. Same environment as CI.
>

---



## Implementation


See the full implementation guide in the source documentation.


## Examples

See [examples.md](examples.md) for code examples.






## References

- [Source Documentation](https://adaptive-enforcement-lab.com/enforce/policy-as-code/)
- [AEL Enforce](https://adaptive-enforcement-lab.com/enforce/)
