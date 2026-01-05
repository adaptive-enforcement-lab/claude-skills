---
name: actions-integration
description: >-
  Integrate GitHub Core Apps with Actions workflows for org-scoped automation. Generate tokens, access APIs, and implement cross-repository operations patterns.
---

# Actions Integration

## When to Use This Skill

This guide explains how to integrate your GitHub Core App with GitHub Actions
workflows for organization-level automation.

> **What You'll Learn**
>
> Generate short-lived tokens, use them with GitHub CLI and APIs, implement common workflow patterns, and handle errors gracefully.


## Prerequisites

Before integrating, ensure you have:

1. **Core App created and installed** - See [GitHub App Setup](../../../secure/github-apps/index.md)
2. **Secrets configured** - `CORE_APP_ID` and `CORE_APP_PRIVATE_KEY` stored in GitHub
3. **Required permissions** - App has permissions for your automation tasks


## Implementation

See the full implementation guide in the [source documentation](https://adaptive-enforcement-lab.com/patterns/github-actions/).


## Techniques


### Authentication Methods

GitHub Apps support three authentication methods, each serving different use cases:

| Method | Scope | Expiration | Primary Use Case |
|--------|-------|------------|------------------|
| **[JWT](jwt-authentication/index.md)** | App-level | 10 minutes | Installation discovery, app metadata, bootstrapping |
| **[Installation Tokens](token-generation/index.md)** | Repository/Org | 1 hour | Repository operations, API access, automation |
| **[OAuth](oauth-authentication/index.md)** | User context | Configurable | User-specific operations, web flows |

> **Which authentication method should I use?**
>
>
> - **Most workflows** → Installation Tokens (via `actions/create-github-app-token`)
> - **App management** → JWT (list installations, app configuration)
> - **User operations** → OAuth (actions on behalf of a user)
>
> See the [Authentication Decision Guide](../../../secure/github-apps/authentication-decision-guide.md) for detailed selection criteria.
## References

- [Source Documentation](https://adaptive-enforcement-lab.com/patterns/github-actions/)
- [AEL Patterns](https://adaptive-enforcement-lab.com/patterns/)
