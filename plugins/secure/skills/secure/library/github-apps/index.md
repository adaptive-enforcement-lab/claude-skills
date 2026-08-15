# GitHub Core App for Organizations

Source: https://adaptive-enforcement-lab.com/secure/github-apps/

This guide describes the concept, setup, and configuration of a GitHub Core App for organization-level automation.

## What is a GitHub Core App?

> **Definition**
>
>
> A **GitHub Core App** is an organization-level GitHub App that provides centralized, secure authentication for GitHub Actions workflows operating across multiple repositories. It serves as the foundational authentication mechanism for org-wide automation.
>

### Why Use a Core App?

> **Core App Approach**
>
>
> - Organization-owned identity independent of individuals
> - Survives personnel changes
> - Complete audit trail of all actions
> - Fine-grained, repository-scoped permissions
> - Higher rate limits (5000 requests/hour per installation)
> - Team-based repository access control
>

### Use Cases

A GitHub Core App enables:

- **Cross-repository operations** - Synchronize files across multiple repositories
- **Team-scoped automation** - Query and operate on team repositories
- **Centralized CI/CD** - Single authentication source for all workflows
- **Compliance automation** - Enforce policies across organization
- **Repository management** - Create, configure, and manage repositories programmatically

## Core App vs. Standard GitHub Apps

| Aspect | Core App | Standard App |
| -------- | ---------- | -------------- |
| **Scope** | Organization-wide | Single repository or selected repos |
| **Purpose** | Infrastructure automation | Feature-specific functionality |
| **Permissions** | Broad, covers common operations | Narrow, task-specific |
| **Installation** | All repositories | Selective repositories |
| **Ownership** | Organization-level admin | Project or team |
| **Lifespan** | Permanent infrastructure | Project lifecycle |

## Prerequisites

### Required Access

> **Required Access**
>
>
> To create a Core App, you need:
>
> - **Organization owner** role
> - Access to organization settings: `https://github.com/organizations/{ORG}/settings/apps`
>

### Planning Considerations

> **Planning Considerations**
>
>
> Before creating the app, determine:
>
> 1. **Permission scope** - Which repository and organization permissions are needed
> 2. **Installation scope** - All repositories or specific teams
> 3. **Token management** - Where secrets will be stored (repository or organization level)
> 4. **Naming convention** - Standard naming (e.g., "CORE App", "Automation Core")
>

## Guide Sections

- [Authentication Decision Guide](authentication-decision-guide.md) - Choose between JWT, installation tokens, and OAuth
- [Creating the Core App](creating-the-app.md) - Step-by-step app creation and configuration
- [Storing Credentials](storing-credentials/index.md) - Managing secrets in GitHub
- [Permission Patterns](permission-patterns.md) - Common permission configurations
- [Security Best Practices](security-best-practices.md) - Securing your Core App
- [Installation Scopes](installation-scopes.md) - Choosing the right installation scope
- [Common Permissions](common-permissions.md) - Permission requirements by use case
- [Troubleshooting](troubleshooting.md) - Common issues and solutions
- [Maintenance](maintenance.md) - Ongoing care and key rotation

## Next Steps

After setting up your Core App:

1. **[GitHub Actions Integration](../../patterns/github-actions/actions-integration/index.md)** - Learn how to use the app in workflows
2. **[Distribution Workflows](../../patterns/github-actions/use-cases/file-distribution/index.md)** - Example use case patterns

## References

- [GitHub Apps Documentation](https://docs.github.com/en/apps)
- [GitHub Apps Permissions](https://docs.github.com/en/rest/overview/permissions-required-for-github-apps)
- [GitHub GraphQL API](https://docs.github.com/en/graphql)
- [Organization Security Best Practices](https://docs.github.com/en/organizations/keeping-your-organization-secure)
