---
name: hardened-release-workflow
description: >-
  Production-ready release workflow examples with signed releases, SLSA provenance, artifact attestations, and minimal permissions.
---

# Hardened Release Workflow

## When to Use This Skill

Copy-paste ready release workflow templates with comprehensive security hardening. Each example demonstrates signed releases, SLSA provenance generation, artifact attestations, minimal permissions, and secure artifact distribution.

> **Complete Security Patterns**
>
>
> These workflows integrate all security patterns from the hub: SHA-pinned actions, minimal GITHUB_TOKEN permissions, SLSA provenance, artifact attestations, signature verification, and secure distribution. Use as production templates for secure software supply chain.


## Implementation

See the full implementation guide in the [source documentation](https://adaptive-enforcement-lab.com/secure/github-actions-security/).


## Key Principles

Every release workflow in this guide implements these controls:

1. **Action Pinning**: All third-party actions pinned to full SHA-256 commit hashes
2. **Minimal Permissions**: Only required permissions granted per job
3. **SLSA Provenance**: Build provenance attestations for supply chain transparency
4. **Artifact Attestations**: Cryptographic signatures for release artifacts
5. **Signature Verification**: Verifiable release authenticity
6. **Immutable Releases**: Tag protection and commit verification
7. **Approval Gates**: Environment protection for production releases


## Examples

See [examples.md](examples.md) for code examples.


## Full Reference

See [reference.md](reference.md) for complete documentation.
## References

- [Source Documentation](https://adaptive-enforcement-lab.com/secure/github-actions-security/)
- [AEL Secure](https://adaptive-enforcement-lab.com/secure/)
