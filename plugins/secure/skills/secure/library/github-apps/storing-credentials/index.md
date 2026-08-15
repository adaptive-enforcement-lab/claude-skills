# Storing GitHub App Credentials

Source: https://adaptive-enforcement-lab.com/secure/github-apps/storing-credentials/

Comprehensive guide to securely storing GitHub App credentials across different environments and platforms.

> **Storage Environment Decision**
>
>
> - **GitHub Actions** - Native GitHub Secrets (recommended for GitHub-hosted workflows)
> - **External CI** - Platform-specific secret management (Jenkins, GitLab CI, CircleCI)
> - **Kubernetes** - External Secrets Operator or Sealed Secrets
> - **Local Development** - Environment variables or encrypted vaults (never in code)
>

## Storage Decision Tree

```mermaid
flowchart TD
    A["Where will credentials be used?"] --> B{"Execution<br/>environment?"}

    B -->|"GitHub Actions"| C["GitHub Secrets"]
    B -->|"External CI/CD"| D["Platform Secrets Manager"]
    B -->|"Kubernetes"| E["External Secrets Operator"]
    B -->|"Local Development"| F["Environment Variables<br/>+ Encrypted Vault"]

    C --> C1["Organization secrets<br/>(recommended)"]
    C --> C2["Repository secrets<br/>(single repo)"]
    C --> C3["Environment secrets<br/>(protected workflows)"]

    D --> D1["Jenkins Credentials"]
    D --> D2["GitLab CI Variables"]
    D --> D3["CircleCI Contexts"]

    E --> E1["AWS Secrets Manager"]
    E --> E2["HashiCorp Vault"]
    E --> E3["Google Secret Manager"]

    %% Ghostty Hardcore Theme
    style A fill:#515354,stroke:#ccccc7,stroke-width:2px,color:#ccccc7
    style B fill:#fd971e,stroke:#e6db74,stroke-width:2px,color:#1b1d1e
    style C fill:#a7e22e,stroke:#bded5f,stroke-width:2px,color:#1b1d1e
    style D fill:#65d9ef,stroke:#a3babf,stroke-width:2px,color:#1b1d1e
    style E fill:#9e6ffe,stroke:#9e6ffe,stroke-width:2px,color:#1b1d1e
    style F fill:#f92572,stroke:#ff669d,stroke-width:2px,color:#1b1d1e
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

## GitHub Actions Secrets

### Repository Secrets

For single-repository usage:

1. Navigate to repository **Settings**
2. Go to **Secrets and variables** → **Actions**
3. Click **New repository secret**
4. Add two secrets:
   - `CORE_APP_ID`: Numeric app ID
   - `CORE_APP_PRIVATE_KEY`: Complete `.pem` file contents

> **When to Use Repository Secrets**
>
>
> - Single repository workflows
> - Repository-specific GitHub Apps
> - Testing and development environments
> - Isolated security boundaries
>

### Organization Secrets

For organization-wide usage (recommended):

1. Navigate to **Organization Settings**
2. Go to **Secrets and variables** → **Actions**
3. Click **New organization secret**
4. Add secrets with same names as above
5. Configure **Repository access**:
   - **All repositories** - Available to all org repos
   - **Selected repositories** - Choose specific repos
   - **Private repositories only** - Additional security layer

> **Advantages of Organization Secrets**
>
>
> - **Single source of truth** - One credential set for all repositories
> - **Centralized rotation** - Update once, applies everywhere
> - **Consistent naming** - Same secret names across org
> - **Simplified auditing** - Track usage from one location
>

### Environment Secrets

For additional security with protected workflows:

1. Create an environment (e.g., `production`, `staging`)
2. Add secrets to the environment
3. Configure protection rules:
   - **Required reviewers** - Manual approval before access
   - **Wait timer** - Delay before secret access
   - **Deployment branches** - Restrict to specific branches
   - **Environment-specific values** - Different credentials per environment

```yaml
jobs:
  deploy:
    environment: production
    steps:
      - name: Generate token
        uses: actions/create-github-app-token@v1
        with:
          app-id: ${{ secrets.CORE_APP_ID }}
          private-key: ${{ secrets.CORE_APP_PRIVATE_KEY }}
```

> **Environment Protection Use Cases**
>
>
> - Production deployments requiring approval
> - Sensitive operations (database migrations, infrastructure changes)
> - Compliance requirements for change control
> - Multi-environment workflows (dev, staging, prod)
>

### Secret Naming Conventions

| Secret Name | Contents | Example | Scope |
| ------------- | ---------- | --------- | ------- |
| `CORE_APP_ID` | Numeric app ID | `123456` | All environments |
| `CORE_APP_PRIVATE_KEY` | Complete PEM file contents | `-----BEGIN RSA PRIVATE KEY-----...` | All environments |
| `PROD_APP_ID` | Production-specific app ID | `789012` | Environment-specific |
| `PROD_APP_PRIVATE_KEY` | Production private key | `-----BEGIN RSA PRIVATE KEY-----...` | Environment-specific |

> **Naming Best Practices**
>
>
> - Use `CORE_APP_*` prefix for organization-wide apps
> - Use environment prefixes (`PROD_`, `STAGING_`) for environment-specific credentials
> - Avoid ambiguous names like `APP_ID` or `PRIVATE_KEY`
> - Document naming conventions in your organization's security policy
>

### Repository vs Organization vs Environment Secrets

| Aspect | Repository Secrets | Organization Secrets | Environment Secrets |
| -------- | ------------------- | --------------------- | --------------------- |
| **Scope** | Single repository | Multiple repositories | Specific environment |
| **Management** | Per-repo updates | Centralized updates | Per-environment updates |
| **Rotation** | Update each repo | Update once | Update per environment |
| **Access Control** | Repository admins | Organization admins | Environment reviewers |
| **Protection** | Workflow-level | Repository selection | Approval + branch rules |
| **Use Case** | Isolated repos | Organization-wide automation | Production deployments |
| **Auditing** | Repository audit log | Organization audit log | Environment deployment log |

### Workflow Access Patterns

#### Basic Token Generation

```yaml
- name: Generate token
  id: app-token
  uses: actions/create-github-app-token@v2
  with:
    app-id: ${{ secrets.CORE_APP_ID }}
    private-key: ${{ secrets.CORE_APP_PRIVATE_KEY }}

- name: Use token
  env:
    GITHUB_TOKEN: ${{ steps.app-token.outputs.token }}
  run: |
    gh api /repos/${{ github.repository }}/issues
```

#### Organization-Scoped Token

```yaml
- name: Generate org-scoped token
  uses: actions/create-github-app-token@v2
  with:
    app-id: ${{ secrets.CORE_APP_ID }}
    private-key: ${{ secrets.CORE_APP_PRIVATE_KEY }}
    owner: my-organization
```

#### Multi-Repository Token

```yaml
- name: Generate multi-repo token
  uses: actions/create-github-app-token@v2
  with:
    app-id: ${{ secrets.CORE_APP_ID }}
    private-key: ${{ secrets.CORE_APP_PRIVATE_KEY }}
    repositories: |
      repo-one
      repo-two
      repo-three
```
