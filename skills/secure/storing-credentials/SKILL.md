---
name: storing-credentials
description: >-
  Securely store GitHub App credentials across different environments. GitHub Actions secrets, external CI, Kubernetes, and automated rotation patterns.
---

# Storing Credentials

## When to Use This Skill

Comprehensive guide to securely storing GitHub App credentials across different environments and platforms.

> **Storage Environment Decision**
>
>
> - **GitHub Actions** - Native GitHub Secrets (recommended for GitHub-hosted workflows)
> - **External CI** - Platform-specific secret management (Jenkins, GitLab CI, CircleCI)
> - **Kubernetes** - External Secrets Operator or Sealed Secrets
> - **Local Development** - Environment variables or encrypted vaults (never in code)



## Implementation


See the full implementation guide in the source documentation.


## Examples

See [examples.md](examples.md) for code examples.






## References

- [Source Documentation](https://adaptive-enforcement-lab.com/secure/github-apps/)
- [AEL Secure](https://adaptive-enforcement-lab.com/secure/)
