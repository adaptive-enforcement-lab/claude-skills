---
name: enforce
description: >-
  Make security mandatory through automation. Branch protection, pre-commit hooks, status checks, policy-as-code, and SLSA provenance for SOC 2 compliance.
---

# Enforce

## When to Use This Skill

This section covers the **enforcement mechanisms** that make security policies mandatory, auditable, and impossible to ignore.

These controls pass SOC 2, ISO 27001, and PCI-DSS audits by shifting security left and making compliance automatic.


## Implementation

See [Implementation Roadmap](implementation-roadmap/index.md) for phased rollout:

1. **Phase 1**: Branch protection (1 week)
2. **Phase 2**: Status checks (2 weeks)
3. **Phase 3**: Pre-commit hooks (1 week)
4. **Phase 4**: Policy-as-code (4 weeks)
5. **Phase 5**: SLSA provenance (2 weeks)

**Total timeline**: 10 weeks for complete enforcement stack.


## Comparison

Understanding the distinction:

- **Secure** ([see Secure](../secure/index.md)): Find and fix security issues
  - Vulnerability scanners that *identify* CVEs
  - SBOM generators that *document* dependencies
  - Security tools that *discover* weaknesses

- **Enforce** (this section): Make security mandatory through automation
  - Branch protection that *requires* reviews
  - Pre-commit hooks that *block* violations
  - Status checks that *prevent* merges
  - Policy-as-code that *rejects* non-compliant resources
  - SLSA provenance that *attests* build integrity

**Litmus test**: Can this be bypassed?

- If **yes** → Belongs in Enforce (make it mandatory)
- If **no** → Belongs in Secure (it's a finding/fix tool)


## Examples

See [examples.md](examples.md) for code examples.


## Full Reference

See [reference.md](reference.md) for complete documentation.


## Related Patterns

- Secure
- Build
- Patterns

## References

- [Source Documentation](https://adaptive-enforcement-lab.com/enforce/index.md/)
- [AEL Enforce](https://adaptive-enforcement-lab.com/enforce/)
