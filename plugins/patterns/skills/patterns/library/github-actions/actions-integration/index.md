# GitHub Actions Integration with Core App

Source: https://adaptive-enforcement-lab.com/patterns/github-actions/actions-integration/

This guide explains how to integrate your GitHub Core App with GitHub Actions
workflows for organization-level automation.

> **What You'll Learn**
>
> Generate short-lived tokens, use them with GitHub CLI and APIs, implement common workflow patterns, and handle errors gracefully.
>

## Prerequisites

Before integrating, ensure you have:

1. **Core App created and installed** - See [GitHub App Setup](../../../secure/github-apps/index.md)
2. **Secrets configured** - `CORE_APP_ID` and `CORE_APP_PRIVATE_KEY` stored in GitHub
3. **Required permissions** - App has permissions for your automation tasks

## Authentication Methods

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
>

## What's Covered

This section walks through the complete integration lifecycle:

**Authentication Methods:**

- [JWT Authentication](jwt-authentication/index.md) - App-level authentication for installation discovery and management
- [Installation Tokens](token-generation/index.md) - Generate short-lived tokens from Core App credentials
- [OAuth Authentication](oauth-authentication/index.md) - User-context authentication for web and device flows
- [Token Lifecycle](token-lifecycle/index.md) - Token expiration, refresh strategies, and caching patterns

**Integration Patterns:**

- [Using Tokens](using-tokens.md) - Integrate tokens with GitHub CLI, Git, and APIs
- [Workflow Patterns](token-generation/workflow-patterns.md) - Common automation patterns
- [Token Validation](token-validation.md) - Verify token scope and permissions
- [Workflow Permissions](workflow-permissions.md) - Configure workflow-level permissions

**Operations:**

- [Error Handling](error-handling/index.md) - Handle authentication errors and token expiration
- [Security Best Practices](security-best-practices.md) - Keep tokens secure
- [Troubleshooting](troubleshooting.md) - Debug common issues
- [Performance Optimization](performance-optimization.md) - Optimize for speed

## References

- [actions/create-github-app-token](https://github.com/actions/create-github-app-token)
- [GitHub Actions Permissions](https://docs.github.com/en/actions/security-guides/automatic-token-authentication)
- [GitHub CLI Manual](https://cli.github.com/manual/)
- [GitHub GraphQL API](https://docs.github.com/en/graphql)
- [GitHub Core App Setup](../../../secure/github-apps/index.md)
