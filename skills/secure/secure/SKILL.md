---
name: secure
description: >-
  Find and fix security issues before they become incidents. Vulnerability scanning, SBOM generation, supply chain security, and secure authentication workflows.
---

# Secure

## When to Use This Skill

This section covers the **tools and practices** for discovering and remediating security issues in code, dependencies, containers, and supply chains.


## Implementation

1. **Start with GitHub Apps**: Replace PATs with secure, auditable authentication
2. **Add vulnerability scanning**: Catch known CVEs before they deploy
3. **Generate SBOMs**: Document your supply chain for compliance
4. **Run Scorecard**: Measure and improve security posture
5. **Layer on enforcement**: Make findings actionable with Enforce patterns


## Comparison

Understanding the distinction:

- **Secure** (this section): Find and fix security issues
  - Vulnerability scanners that *identify* CVEs
  - SBOM generators that *document* dependencies
  - Security tools that *discover* weaknesses
  - GitHub Apps that *provide* secure authentication

- **Enforce** ([see Enforce](../enforce/index.md)): Make security mandatory through automation
  - Branch protection that *requires* reviews
  - Pre-commit hooks that *block* violations
  - Status checks that *prevent* merges
  - Policy-as-code that *enforces* runtime compliance

**Litmus test**: Can this be bypassed?

- If **no** → It's a Secure tool (finding/fixing)
- If **yes** → It belongs in Enforce (making mandatory)


## Examples

See [examples.md](examples.md) for code examples.


## Related Patterns

- Enforce
- Build
- Patterns

## References

- [Source Documentation](https://adaptive-enforcement-lab.com/secure/index.md/)
- [AEL Secure](https://adaptive-enforcement-lab.com/secure/)
