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

See the full implementation guide in the [source documentation](https://adaptive-enforcement-lab.com/secure/github-actions-security/).


## Techniques


### What is OIDC Federation?

OpenID Connect (OIDC) allows GitHub Actions to authenticate to cloud providers without storing credentials as secrets.

**How It Works**:

1. GitHub Actions requests OIDC token via `id-token: write` permission
2. GitHub generates short-lived JWT with workflow claims (repo, branch, commit, etc.)
3. Workflow presents JWT to cloud provider's token exchange endpoint
4. Cloud provider validates claims against trust policy
5. Cloud provider issues temporary credentials (15 minutes to 1 hour)
6. Workflow uses temporary credentials to access cloud resources

**Key Benefits**:

- **No stored secrets**: Credentials never stored in GitHub
- **Short-lived tokens**: Expire in minutes, not years
- **Cryptographic binding**: Token tied to specific workflow context
- **Automatic rotation**: New token for every workflow run
- **Audit trail**: Cloud provider logs include workflow identity
- **Reduced attack surface**: Compromised workflow cannot exfiltrate long-lived credentials


### OIDC Token Claims

GitHub OIDC tokens include claims identifying the workflow context.

**Standard Claims**:

| Claim | Example | Description |
| ----- | ------- | ----------- |
| `sub` | `repo:org/repo:ref:refs/heads/main` | Subject identifier (most important for trust policies) |
| `aud` | `https://github.com/org` | Audience (usually organization or repo URL) |
| `iss` | `https://token.actions.githubusercontent.com` | Issuer (GitHub Actions) |
| `repository` | `org/repo` | Repository name |
| `repository_owner` | `org` | Organization or user |
| `ref` | `refs/heads/main` | Git ref that triggered workflow |
| `sha` | `abc123...` | Commit SHA |
| `workflow` | `CI` | Workflow name |
| `job_workflow_ref` | `org/repo/.github/workflows/ci.yml@refs/heads/main` | Workflow file reference |
| `environment` | `production` | Environment name (if used) |


### Subject Claim Patterns

The `sub` claim determines which workflows can assume cloud roles. Design subject patterns for least privilege.

### Repository-Level Trust

**Pattern**: Allow any workflow in specific repository

**Subject**: `repo:org/repo-name:*`

**Use Case**: All workflows in repository can access cloud resources

**Risk**: Any workflow file change can access credentials

**Example**:

```text
repo:adaptive-enforcement-lab/api-service:*
```

### Branch-Level Trust

**Pattern**: Allow workflows from specific branch only

**Subject**: `repo:org/repo-name:ref:refs/heads/main`

**Use Case**: Only main branch deployments

**Risk**: Lower risk, but all main workflows have access

**Example**:

```text
repo:adaptive-enforcement-lab/api-service:ref:refs/heads/main
```

### Environment-Level Trust (Recommended)

**Pattern**: Allow workflows targeting specific environment

**Subject**: `repo:org/repo-name:environment:production`

**Use Case**: Production deployments with approval gates

**Risk**: Lowest risk, combined with environment protection rules

**Example**:

```text
repo:adaptive-enforcement-lab/api-service:environment:production
```

### Pull Request Protection

**Pattern**: Block pull requests from assuming role

**Subject**: `repo:org/repo-name:ref:refs/heads/*` (excludes `refs/pull/*`)

**Use Case**: Prevent fork PRs from accessing production

**Risk**: Blocks legitimate PR workflows that need cloud access

**Example Subject Filter**:

```text
token.actions.githubusercontent.com:sub = "repo:adaptive-enforcement-lab/api-service:ref:refs/heads/*"
```

> **Cloud Provider OIDC Examples**
>
>
> For detailed cloud provider setup including GCP Workload Identity Federation and Azure Federated Credentials, see [Cloud Provider OIDC Patterns](./cloud-providers.md).


### GCP Workload Identity Federation

GCP uses Workload Identity Pools and Providers to validate GitHub tokens.

### Setup Process

#### Step 1: Create Workload Identity Pool

```bash
gcloud iam workload-identity-pools create github-pool \
  --location=global \
  --display-name="GitHub Actions Pool"
```

#### Step 2: Create Workload Identity Provider

```bash
gcloud iam workload-identity-pools providers create-oidc github-provider \
  --location=global \
  --workload-identity-pool=github-pool \
  --issuer-uri=https://token.actions.githubusercontent.com \
  --attribute-mapping="google.subject=assertion.sub,attribute.repository=assertion.repository,attribute.repository_owner=assertion.repository_owner,attribute.ref=assertion.ref" \
  --attribute-condition="assertion.repository_owner == 'adaptive-enforcement-lab'"
```

**Key Configuration**:

- `issuer-uri`: GitHub OIDC token issuer
- `attribute-mapping`: Maps GitHub claims to GCP attributes
- `attribute-condition`: Additional filtering (organization-level trust)

#### Step 3: Grant Service Account Access

```bash
gcloud iam service-accounts add-iam-policy-binding deploy@my-project.iam.gserviceaccount.com \
  --role=roles/iam.workloadIdentityUser \
  --member="principalSet://iam.googleapis.com/projects/123456789/locations/global/workloadIdentityPools/github-pool/attribute.repository/adaptive-enforcement-lab/api-service"
```

**Attribute Filtering** (environment-level):

```bash
--member="principalSet://iam.googleapis.com/projects/123456789/locations/global/workloadIdentityPools/github-pool/attribute.repository/adaptive-enforcement-lab/api-service/attribute.environment/production"
```

### Workflow Example

```yaml
name: Deploy to GCP
on:
  push:
    branches: [main]

permissions:
  id-token: write  # Required for OIDC token
  contents: read

jobs:
  deploy:
    runs-on: ubuntu-latest
    environment: production
    steps:


## Examples

See [examples.md](examples.md) for code examples.
## References

- [Source Documentation](https://adaptive-enforcement-lab.com/secure/github-actions-security/)
- [AEL Secure](https://adaptive-enforcement-lab.com/secure/)
