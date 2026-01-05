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
>
> - 1-hour expiration (automatic refresh available)
> - Requires GitHub App installation on target repositories
> - Permissions limited to app's configured scope
> - Cannot perform user-attributed actions


## Implementation

Installation tokens provide automated, secure access to repositories where your GitHub App is installed. Use installation tokens for GitHub Actions workflows, CI/CD automation, and cross-repository operations.

> **When to Use Installation Tokens**
>
>
> Installation tokens are for **automated repository operations**. Use JWT for app-level operations and OAuth for user-attributed actions.
>

## Overview

Installation tokens authenticate your GitHub App for specific repository operations. They enable:

- **Cross-repository automation** - Operate across multiple repositories
- **Organization-wide workflows** - Access all repositories in your organization
- **Automated processes** - No user interaction required
- **Scoped permissions** - Limit access to specific repositories
- **Short-lived credentials** - 1-hour expiration for security

> **Token Limitations**
>
>
> - 1-hour expiration (automatic refresh available)
> - Requires GitHub App installation on target repositories
> - Permissions limited to app's configured scope
> - Cannot perform user-attributed actions
>

## Token Scoping Decision


*See [examples.md](examples.md) for detailed code examples.*

## Basic Usage

### Single Repository Token

Generate a token scoped to the current repository.


*See [examples.md](examples.md) for detailed code examples.*

**Output**: Token accessible via `${{ steps.app_token.outputs.token }}`

**Scope**: Current repository only (where workflow runs)

## Organization-Scoped Tokens

Generate tokens with access to all repositories where the app is installed.


*See [examples.md](examples.md) for detailed code examples.*

> **Owner Parameter is Critical**
>
>
> - **With `owner`**: Access all repositories in the organization
> - **Without `owner`**: Access only the current repository
> - Must match your GitHub organization name exactly
>

**Use cases**:

- Discovery workflows (list all repositories)
- Cross-repository automation
- Organization-wide policy enforcement
- Dynamic repository targeting

## Repository-Scoped Tokens

Limit token access to specific repositories for enhanced security.


*See [examples.md](examples.md) for detailed code examples.*

> **Security Best Practice**
>
>
> Use repository-scoped tokens when you know exactly which repositories need access. This follows the principle of least privilege.
>

**Benefits**:

- Explicit allow list of repositories
- Reduces blast radius if token is compromised
- Clear audit trail of intended access
- Enforces access boundaries

## When NOT to Use Installation Tokens

> **Don't Use Installation Tokens For**
>
>
> - **User-attributed actions** - Use OAuth instead
> - **App-level operations** - Use JWT (list installations, get app manifest)
> - **Public repository read-only access** - Use `GITHUB_TOKEN` if simpler
> - **Personal repository access** - Use OAuth for user's private repos
> - **Operations requiring user identity** - Actions appear as "bot" with installation tokens
>

## Next Steps

- [Workflow Patterns](workflow-patterns.md) - Cross-repository automation patterns
- [Use Cases](use-cases.md) - Real-world implementation examples
- [Lifecycle and Security](lifecycle-security.md) - Token management and security best practices

### Overview

Installation tokens authenticate your GitHub App for specific repository operations. They enable:

- **Cross-repository automation** - Operate across multiple repositories
- **Organization-wide workflows** - Access all repositories in your organization
- **Automated processes** - No user interaction required
- **Scoped permissions** - Limit access to specific repositories
- **Short-lived credentials** - 1-hour expiration for security

> **Token Limitations**
>
>
> - 1-hour expiration (automatic refresh available)
> - Requires GitHub App installation on target repositories
> - Permissions limited to app's configured scope
> - Cannot perform user-attributed actions

### Token Scoping Decision


*See [examples.md](examples.md) for detailed code examples.*

### Basic Usage

### Single Repository Token

Generate a token scoped to the current repository.


*See [examples.md](examples.md) for detailed code examples.*

**Output**: Token accessible via `${{ steps.app_token.outputs.token }}`

**Scope**: Current repository only (where workflow runs)

### Organization-Scoped Tokens

Generate tokens with access to all repositories where the app is installed.


*See [examples.md](examples.md) for detailed code examples.*

> **Owner Parameter is Critical**
>
>
> - **With `owner`**: Access all repositories in the organization
> - **Without `owner`**: Access only the current repository
> - Must match your GitHub organization name exactly
>

**Use cases**:

- Discovery workflows (list all repositories)
- Cross-repository automation
- Organization-wide policy enforcement
- Dynamic repository targeting

### Repository-Scoped Tokens

Limit token access to specific repositories for enhanced security.


*See [examples.md](examples.md) for detailed code examples.*

> **Security Best Practice**
>
>
> Use repository-scoped tokens when you know exactly which repositories need access. This follows the principle of least privilege.
>

**Benefits**:

- Explicit allow list of repositories
- Reduces blast radius if token is compromised
- Clear audit trail of intended access
- Enforces access boundaries

### When NOT to Use Installation Tokens

> **Don't Use Installation Tokens For**
>
>
> - **User-attributed actions** - Use OAuth instead
> - **App-level operations** - Use JWT (list installations, get app manifest)
> - **Public repository read-only access** - Use `GITHUB_TOKEN` if simpler
> - **Personal repository access** - Use OAuth for user's private repos
> - **Operations requiring user identity** - Actions appear as "bot" with installation tokens

### Next Steps

- [Workflow Patterns](workflow-patterns.md) - Cross-repository automation patterns
- [Use Cases](use-cases.md) - Real-world implementation examples
- [Lifecycle and Security](lifecycle-security.md) - Token management and security best practices


## Examples

See [examples.md](examples.md) for code examples.


## Full Reference

See [reference.md](reference.md) for complete documentation.
## References

- [Source Documentation](https://adaptive-enforcement-lab.com/patterns/github-actions/)
- [AEL Patterns](https://adaptive-enforcement-lab.com/patterns/)
