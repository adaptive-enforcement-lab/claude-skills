---
name: oidc-federation-patterns
description: >-
  Secretless authentication to cloud providers using OpenID Connect federation. GCP, Azure, and cloud-agnostic examples with subject claim patterns and trust policies.
---

# OIDC Federation Patterns

## When to Use This Skill

Eliminate stored credentials entirely. OIDC federation replaces long-lived secrets with short-lived tokens tied to workflow context.

> **The Win**
>
>
> OIDC federation means zero stored secrets for cloud authentication. No rotation burden, no credential sprawl, no leaked keys in logs. Tokens expire in minutes and are cryptographically bound to your repository, branch, and commit.



## Implementation


See the full implementation guide in the source documentation.


## Examples

See [examples.md](examples.md) for code examples.






## References

- [Source Documentation](https://adaptive-enforcement-lab.com/secure/github-actions-security/)
- [AEL Secure](https://adaptive-enforcement-lab.com/secure/)
