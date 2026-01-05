---
name: github-actions-security-cheat-sheet - Reference
description: Complete reference for GitHub Actions Security Cheat Sheet
---

# GitHub Actions Security Cheat Sheet - Reference

This is the complete reference documentation extracted from the source.


# GitHub Actions Security Cheat Sheet

One-page security reference for hardening GitHub Actions workflows. Copy-paste ready patterns for production use.

> **Start Here**
>
>
> New to GitHub Actions security? Start with SHA pinning and minimal permissions. Both provide high impact with minimal workflow changes.
>

## Quick Security Checklist

Essential controls for every workflow:

- [ ] All actions pinned to full SHA-256 hashes with version comments
- [ ] Explicit minimal `permissions` block at workflow or job level
- [ ] OIDC federation for cloud access (no stored credentials)
- [ ] `pull_request` trigger for untrusted code (not `pull_request_target`)
- [ ] Input validation for any `github.event.*` values used in shell
- [ ] Secret scanning enabled with push protection
- [ ] Self-hosted runners use ephemeral patterns
- [ ] Environment protection for production deployments
- [ ] Dependabot enabled for automated action updates

## Action Pinning

Pin actions to immutable SHA-256 commits. Tags are mutable and vulnerable.

### SHA Pinning Pattern

```yaml
steps:
  # ✅ GOOD: SHA pinned with version comment
  - uses: actions/checkout@b4ffde65f46336ab88eb53be808477a3936bae11  # v4.1.1
  - uses: actions/setup-node@5e21ff4d9bc1a8cf6de233a3057d20ec6b3fb69d  # v3.8.1

  # ❌ BAD: Mutable tag reference
  # - uses: actions/checkout@v4
```

### Comment Formats

| Format | Example | Use Case |
| ------ | ------- | -------- |
| **Standard** | `# v4.1.1` | Most workflows |
| **Extended** | `# v4.1.1 (2023-11-15)` | Track update dates |
| **Date-based** | `# v4.1.1 @ 2023-11-15` | Compliance tracking |

### Common Actions Reference

| Action | Latest SHA (v4.1.1 / v3.8.1) | Trust Tier |
| ------ | ---------------------------- | ---------- |
| `actions/checkout` | `b4ffde65f46336ab88eb53be808477a3936bae11` | Tier 1 (GitHub) |
| `actions/setup-node` | `5e21ff4d9bc1a8cf6de233a3057d20ec6b3fb69d` | Tier 1 (GitHub) |
| `actions/cache` | `13aacd865c20de90d75de3b17ebe84f7a17d57d2` | Tier 1 (GitHub) |
| `actions/upload-artifact` | `26f96dfa697d77e81fd5907df203aa23a56210a8` | Tier 1 (GitHub) |
| `github/codeql-action/init` | `cdcdbb579706841c47f7063dda365e292e5cad7a` | Tier 1 (GitHub) |

[**See full pinning guide →**](../action-pinning/sha-pinning.md)

### Dependabot Auto-Updates

`.github/dependabot.yml`:

```yaml
version: 2
updates:
  - package-ecosystem: "github-actions"
    directory: "/"
    schedule:
      interval: "weekly"
    groups:
      # Group GitHub-maintained actions
      github-actions-core:
        patterns:
          - "actions/*"
          - "github/*"
```

[**See Dependabot guide →**](../action-pinning/dependabot.md)

## GITHUB_TOKEN Permissions

Minimize token scope. Default `write-all` is dangerous.

### Minimal Permissions Pattern

```yaml
name: Secure CI
on: [push, pull_request]

# Workflow-level: deny most access
permissions:
  contents: read

jobs:
  test:
    runs-on: ubuntu-latest
    # Job inherits workflow-level permissions
    steps:
      - uses: actions/checkout@b4ffde65f46336ab88eb53be808477a3936bae11  # v4.1.1
      - run: npm test

  publish:
    runs-on: ubuntu-latest
    # Job-level: override for specific needs
    permissions:
      contents: read
      packages: write  # Only this job can publish
    steps:
      - run: npm publish
```

### Permission Quick Reference

| Workflow Type | Required Permissions |
| ------------- | -------------------- |
| **CI/Test** | `contents: read` |
| **PR Comments** | `contents: read, pull-requests: write` |
| **Release** | `contents: write, packages: write` |
| **Deploy** | `id-token: write, contents: read` (OIDC) |
| **Security Scan** | `contents: read, security-events: write` |
| **GitHub Pages** | `contents: read, pages: write, id-token: write` |

### Common Permissions

| Permission | Read | Write |
| ---------- | ---- | ----- |
| `contents` | Clone repo | Push commits, tags |
| `pull-requests` | Read PRs | Create/update PRs, comments |
| `issues` | Read issues | Create/modify issues |
| `packages` | Download packages | Publish packages |
| `id-token` | - | Request OIDC JWT (cloud auth) |
| `security-events` | - | Upload SARIF to Security tab |

[**See permissions guide →**](../token-permissions/index.md)

## Secret Management

Eliminate long-lived credentials. Use OIDC federation.

### OIDC Federation (Recommended)

```yaml
jobs:
  deploy:
    runs-on: ubuntu-latest
    permissions:
      id-token: write  # Request OIDC token
      contents: read
    environment: production  # Restrict trust to environment
    steps:
      # ✅ GOOD: No stored credentials
      - uses: google-github-actions/auth@55bd3a7c6e2ae7cf1877fd1ccb9d54c0503c457c  # v2.1.2
        with:
          workload_identity_provider: 'projects/123/locations/global/workloadIdentityPools/github/providers/github'
          service_account: 'deploy@project.iam.gserviceaccount.com'

      # ❌ BAD: Stored service account key
      # - run: echo "${{ secrets.GCP_SA_KEY }}" | base64 -d > key.json
```

### OIDC Subject Claim Patterns

| Pattern | Subject Claim | Trust Level |
| ------- | ------------- | ----------- |
| **Environment** | `repo:org/repo:environment:prod` | **Recommended** |
| **Branch** | `repo:org/repo:ref:refs/heads/main` | Medium |
| **Repository** | `repo:org/repo` | Broad (use with caution) |

[**See OIDC guide →**](../secrets/oidc/index.md)

### Secret Rotation Schedule

| Credential Type | Rotation Frequency | Priority |
| --------------- | ------------------ | -------- |
| Production API keys | 30 days | Critical |
| CI/CD tokens | 60 days | High |
| Service account keys | 90 days (prefer OIDC) | High |
| Test environment | 180 days | Medium |
| Development tokens | 365 days | Low |

[**See rotation guide →**](../secrets/rotation/index.md)

### Secret Scanning

Enable push protection to block credential commits.

Configuration path: `Settings → Code security → Secret scanning → Push protection → Enable`

```yaml
# .github/workflows/secret-scan.yml
name: Secret Scanning
on: [push, pull_request]

jobs:
  scan:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@b4ffde65f46336ab88eb53be808477a3936bae11  # v4.1.1
        with:
          fetch-depth: 0  # Full history for scanning

      - uses: gitleaks/gitleaks-action@cb7149a9a1d86f1c2e3ab90ae2f43a75ed56e95a  # v2.3.2
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
```

[**See scanning guide →**](../secrets/scanning/index.md)

## Third-Party Actions

Evaluate before adopting. Not all actions are safe.

### Trust Tiers

| Tier | Publisher | Verification | Risk | Pinning Required |
| ---- | --------- | ------------ | ---- | ---------------- |
| **1** | GitHub (`actions/*`, `github/*`) | Official | Low | SHA recommended |
| **2** | Verified publishers (blue checkmark) | Verified org | Medium | SHA required |
| **3** | Community (active maintenance) | None | High | SHA + source review |
| **4** | Unknown/unmaintained | None | Very High | Avoid or fork |

### Action Evaluation Checklist

Before adding a third-party action:

- [ ] Check maintainer trustworthiness (organization, history, reputation)
- [ ] Review repository health (stars, forks, recent commits, open issues)
- [ ] Audit source code for suspicious patterns (secret exfiltration, network calls)
- [ ] Check security history (past vulnerabilities, incident response quality)
- [ ] Review permission requirements (does it need write access?)
- [ ] Verify maintenance activity (recent commits, responsive maintainers)
- [ ] Consider forking for critical workflows

[**See evaluation guide →**](../third-party-actions/evaluation.md)

### Organization Allowlisting

GitHub Enterprise: `Organization Settings → Actions → General → Policies`

```yaml
# Example policy: Tier 1 + Tier 2 only
Allowed actions and reusable workflows:
  - Allow actions created by GitHub: ✅
  - Allow actions by Marketplace verified creators: ✅
  - Allow specified actions and reusable workflows:
      - aquasecurity/trivy-action@*
      - google-github-actions/*@*
```

[**See allowlisting guide →**](../third-party-actions/allowlisting.md)

## Self-Hosted Runner Security

Never use persistent runners for untrusted code.

### Deployment Models

| Model | Security | Complexity | Use Case |
| ----- | -------- | ---------- | -------- |
| **GitHub-hosted** | High | None | Public repos, low trust requirement |
| **Ephemeral containers** | High | Medium | Private repos, moderate isolation |
| **Ephemeral VMs** | Very High | High | Production, compliance requirements |
| **Persistent runners** | Low | Low | **Avoid for public repos** |

### Ephemeral Runner Pattern

```bash
#!/bin/bash
# Podman ephemeral runner with strict isolation
podman run --rm \
  --security-opt=no-new-privileges:true \
  --cap-drop=ALL \
  --read-only \
  --tmpfs /tmp:rw,noexec,nosuid,size=1g \
  --network=slirp4netns:enable_ipv6=false \
  -e RUNNER_EPHEMERAL=true \
  -e GITHUB_TOKEN="${GITHUB_TOKEN}" \
  ghcr.io/myorg/runner:latest
```

[**See ephemeral patterns →**](../runners/ephemeral/index.md)

### Runner Hardening Checklist

- [ ] Use ephemeral mode (VMs or containers destroyed after each job)
- [ ] Deny-by-default firewall (UFW, iptables) with GitHub API allow-list
- [ ] Block cloud metadata endpoints (169.254.169.254)
- [ ] Dedicated unprivileged user (no sudo, restricted shell)
- [ ] No stored credentials (OIDC federation only)
- [ ] Restrict runner group to private repositories only
- [ ] Enable audit logging (auditd, centralized collection)
- [ ] Automatic security updates (unattended-upgrades, yum-cron)

[**See hardening guide →**](../runners/hardening/index.md)

### Runner Group Restrictions

Restrict sensitive runners to trusted repositories and workflows:

```bash
# Example: API-based runner group configuration
gh api --method PUT \
  /orgs/ORG/actions/runner-groups/GROUP_ID \
  -f name='production-runners' \
  -f visibility='selected' \
  -F selected_repository_ids='[123,456]' \
  -f allows_public_repositories=false \
  -f restricted_to_workflows=true \
  -F selected_workflows='[".github/workflows/deploy.yml@refs/heads/main"]'
```

[**See runner groups →**](../runners/groups/index.md)

## Workflow Triggers

Choose triggers carefully. `pull_request_target` is dangerous.

### Trigger Security Comparison

| Trigger | Execution Context | GITHUB_TOKEN | Secrets | Fork PR Safety |
| ------- | ----------------- | ------------ | ------- | -------------- |
| **`pull_request`** | Fork PR branch | Read-only | ❌ Not exposed | ✅ Safe |
| **`pull_request_target`** | Base branch | Write | ✅ Exposed | ❌ **Dangerous** |
| **`workflow_run`** | Base branch | Write | ✅ Exposed | ✅ Safe (with validation) |
| **`push`** | Pushed branch | Write | ✅ Exposed | N/A |

### Safe Fork PR Pattern

```yaml
# .github/workflows/ci.yml
name: CI
on:
  pull_request:  # Safe for untrusted code
    branches: [main]

permissions:
  contents: read  # Read-only access

jobs:

