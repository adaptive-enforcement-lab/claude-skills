---
name: installation-token-generation - Reference
description: Complete reference for Installation Token Generation
---

# Installation Token Generation - Reference

This is the complete reference documentation extracted from the source.


# Installation Token Generation

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

```mermaid
flowchart TD
    A["What repositories<br/>need access?"] --> B{"Access pattern?"}

    B -->|"All org repos<br/>(flexible scope)"| C["Organization-Scoped Token"]
    B -->|"Specific repos only<br/>(minimal scope)"| D["Repository-Scoped Token"]
    B -->|"Current repo only<br/>(workflow repo)"| E["Default Token"]

    C --> C1["Use owner parameter"]
    C --> C2["Access all installed repos"]
    C --> C3["Best for dynamic workflows"]

    D --> D1["Use repositories parameter"]
    D --> D2["Explicit allow list"]
    D --> D3["Best for security"]

    E --> E1["No parameters needed"]
    E --> E2["Single repo access"]
    E --> E3["Simplest pattern"]

    %% Ghostty Hardcore Theme
    style A fill:#515354,stroke:#ccccc7,stroke-width:2px,color:#ccccc7
    style B fill:#fd971e,stroke:#e6db74,stroke-width:2px,color:#1b1d1e
    style C fill:#a7e22e,stroke:#bded5f,stroke-width:2px,color:#1b1d1e
    style D fill:#9e6ffe,stroke:#9e6ffe,stroke-width:2px,color:#1b1d1e
    style E fill:#66d9ee,stroke:#a1efe4,stroke-width:2px,color:#1b1d1e
    style C1 fill:#515354,stroke:#ccccc7,stroke-width:1px,color:#ccccc7
    style C2 fill:#515354,stroke:#ccccc7,stroke-width:1px,color:#ccccc7
    style C3 fill:#515354,stroke:#ccccc7,stroke-width:1px,color:#ccccc7
    style D1 fill:#515354,stroke:#ccccc7,stroke-width:1px,color:#ccccc7
    style D2 fill:#515354,stroke:#ccccc7,stroke-width:1px,color:#ccccc7
    style D3 fill:#515354,stroke:#ccccc7,stroke-width:1px,color:#ccccc7
    style E1 fill:#515354,stroke:#ccccc7,stroke-width:1px,color:#ccccc7
    style E2 fill:#515354,stroke:#ccccc7,stroke-width:1px,color:#ccccc7
    style E3 fill:#515354,stroke:#ccccc7,stroke-width:1px,color:#ccccc7
```

## Basic Usage

### Single Repository Token

Generate a token scoped to the current repository.

```yaml
name: Single Repo Operation

on:
  workflow_dispatch:

jobs:
  example:
    runs-on: ubuntu-latest
    steps:
      - name: Generate repository token
        id: app_token
        uses: actions/create-github-app-token@v2
        with:
          app-id: ${{ secrets.CORE_APP_ID }}
          private-key: ${{ secrets.CORE_APP_PRIVATE_KEY }}

      - name: Use token
        env:
          GH_TOKEN: ${{ steps.app_token.outputs.token }}
        run: |
          # Token scoped to current repository only
          gh api repos/${{ github.repository }} --jq .full_name
```

**Output**: Token accessible via `${{ steps.app_token.outputs.token }}`

**Scope**: Current repository only (where workflow runs)

## Organization-Scoped Tokens

Generate tokens with access to all repositories where the app is installed.

```yaml
name: Organization-Wide Operation

on:
  workflow_dispatch:

jobs:
  org-scope:
    runs-on: ubuntu-latest
    steps:
      - name: Generate org-scoped token
        id: app_token
        uses: actions/create-github-app-token@v2
        with:
          app-id: ${{ secrets.CORE_APP_ID }}
          private-key: ${{ secrets.CORE_APP_PRIVATE_KEY }}
          owner: adaptive-enforcement-lab  # Organization name

      - name: List all org repositories
        env:
          GH_TOKEN: ${{ steps.app_token.outputs.token }}
        run: |
          echo "## Organization Repositories" >> $GITHUB_STEP_SUMMARY
          gh repo list adaptive-enforcement-lab \
            --limit 100 \
            --json name,description,visibility \
            --jq '.[] | "- **\(.name)** (\(.visibility)): \(.description)"' \
            >> $GITHUB_STEP_SUMMARY
```

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

```yaml
name: Multi-Repository Operation

on:
  workflow_dispatch:

jobs:
  repo-scope:
    runs-on: ubuntu-latest
    steps:
      - name: Generate repo-scoped token
        id: app_token
        uses: actions/create-github-app-token@v2
        with:
          app-id: ${{ secrets.CORE_APP_ID }}
          private-key: ${{ secrets.CORE_APP_PRIVATE_KEY }}
          repositories: |
            frontend-app
            backend-api
            infrastructure

      - name: Check repository status
        env:
          GH_TOKEN: ${{ steps.app_token.outputs.token }}
        run: |
          for repo in frontend-app backend-api infrastructure; do
            echo "Checking $repo..."
            gh api repos/adaptive-enforcement-lab/$repo \
              --jq '{name: .name, default_branch: .default_branch, private: .private}'
          done
```

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

