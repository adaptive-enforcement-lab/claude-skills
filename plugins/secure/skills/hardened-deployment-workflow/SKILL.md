---
name: hardened-deployment-workflow
description: >-
  Production-ready deployment workflow examples with OIDC authentication, environment protection, approval gates, and rollback patterns.
---

# Hardened Deployment Workflow

## When to Use This Skill

Copy-paste ready deployment workflow templates with comprehensive security hardening. Each example demonstrates OIDC authentication, environment protection, approval gates, zero-downtime deployments, and automated rollback patterns.

> **Complete Security Patterns**
>
>
> These workflows integrate all security patterns from the hub: OIDC federation (no stored secrets), environment protection with approval gates, SHA-pinned actions, minimal GITHUB_TOKEN permissions, deployment verification, and automated rollback. Use as production templates for secure deployments.


## Implementation

See the full implementation guide in the [source documentation](https://adaptive-enforcement-lab.com/secure/github-actions-security/).


## Key Principles

Every deployment workflow in this guide implements these controls:

1. **OIDC Authentication**: Secretless cloud authentication with short-lived tokens
2. **Environment Protection**: Required reviewers and wait timers for production
3. **Minimal Permissions**: `id-token: write` for OIDC, `contents: read` by default
4. **Approval Gates**: Human review before production deployment
5. **Deployment Verification**: Health checks after deployment
6. **Rollback Automation**: Automatic rollback on failure
7. **Audit Trail**: Deployment tracking and change logs


## Full Reference

See [reference.md](reference.md) for complete documentation.
## References

- [Source Documentation](https://adaptive-enforcement-lab.com/secure/github-actions-security/)
- [AEL Secure](https://adaptive-enforcement-lab.com/secure/)
