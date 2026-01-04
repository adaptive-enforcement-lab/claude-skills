---
name: installation-token-generation
description: >-
  Generate short-lived installation tokens from GitHub App credentials with actions/create-github-app-token. Organization-scoped and repository-scoped patterns for automated cross-repo workflows.
---

# Installation Token Generation

## When to Use This Skill

Installation tokens authenticate your GitHub App for specific repository operations. They enable:

- **Cross-repository automation** - Operate across multiple repositories
- **Organization-wide workflows** - Access all repositories in your organization
- **Automated processes** - No user interaction required
- **Scoped permissions** - Limit access to specific repositories
- **Short-lived credentials** - 1-hour expiration for security

> **Token Limitations**
>


    - 1-hour expiration (automatic refresh available)
    - Requires GitHub App installation on target repositories
    - Permissions limited to app's configured scope
    - Cannot perform user-attributed actions

##



## Implementation


See the full implementation guide in the source documentation.


## Examples

See [examples.md](examples.md) for code examples.






## References

- [Source Documentation](https://adaptive-enforcement-lab.com/patterns/github-actions/)
- [AEL Patterns](https://adaptive-enforcement-lab.com/patterns/)
