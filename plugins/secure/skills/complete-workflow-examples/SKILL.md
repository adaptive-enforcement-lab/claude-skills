---
name: complete-workflow-examples
description: >-
  Copy-paste hardened CI/CD workflows with SHA-pinned actions, minimal GITHUB_TOKEN permissions, OIDC authentication, and comprehensive security scanning for GitHub Actions.
---

# Complete Workflow Examples

## When to Use This Skill

> **Ready-to-Deploy Templates**
>
>
> These examples integrate multiple security controls into production-ready workflows. Each template includes inline security comments, permission scoping, and cross-references to detailed pattern documentation.
>

Copy-paste ready workflows demonstrating all security patterns from this hub.

Each example integrates multiple security controls from across the hub: action pinning, minimal permissions, secret management, safe triggers, and more. All examples are complete and production-ready.


## Implementation

> **Ready-to-Deploy Templates**
>
>
> These examples integrate multiple security controls into production-ready workflows. Each template includes inline security comments, permission scoping, and cross-references to detailed pattern documentation.
>

Copy-paste ready workflows demonstrating all security patterns from this hub.

Each example integrates multiple security controls from across the hub: action pinning, minimal permissions, secret management, safe triggers, and more. All examples are complete and production-ready.

## Available Examples

### [Secure CI Workflow](ci-workflow/index.md)

Hardened continuous integration with comprehensive security controls.

**Key Patterns**:

- Fork PR security with two-stage workflows
- Language-specific testing (Node.js, Python, Go)
- Secret scanning prevention
- Minimal GITHUB_TOKEN permissions
- SHA-pinned actions with version comments

**Use Cases**: Test automation, PR validation, pre-merge quality gates

---

### [Release Workflow](release-workflow/index.md)

Signed releases with SLSA provenance and artifact attestations.

**Key Patterns**:

- Keyless signing with OIDC
- SLSA L2/L3 provenance generation
- Artifact attestations for GitHub releases, containers, and NPM packages
- Environment protection for release branches
- Minimal permissions with `id-token: write` and `attestations: write`

**Use Cases**: GitHub releases, container publishing, NPM publishing, signed artifacts

---

### [Deployment Workflow](deployment-workflow/index.md)

OIDC-based cloud deployment with environment protection and automated rollback.

**Key Patterns**:

- OIDC federation to GCP (no stored secrets)
- Environment protection with approval gates and wait timers
- Canary rollout with traffic migration (10% → 100%)
- Container scanning and signing
- Automated rollback on deployment failure

**Use Cases**: Cloud Run deployment, Kubernetes/Helm deployment, multi-environment pipelines, canary releases

---

### [Security Scanning](security-scanning/index.md)

Comprehensive SAST, dependency scanning, container scanning, and SARIF upload.

**Key Patterns**:

- CodeQL SAST with security-extended queries
- Dependency review with severity-based failure
- Container scanning (Trivy, SBOM generation with Syft/Grype)
- Language-specific scanning (Bandit, gosec, ESLint)
- SARIF aggregation and upload to Security tab
- Scheduled vulnerability scanning with issue creation

**Use Cases**: Security validation, compliance scanning, vulnerability detection, scheduled audits

---

## Common Security Controls

All examples use:

- **SHA-pinned actions** with version comments for supply chain security
- **Minimal GITHUB_TOKEN permissions** scoped to job requirements
- **Environment protection** where appropriate (deployments, releases)
- **OIDC federation** for cloud access (no stored secrets)
- **Input validation** and safe trigger patterns
- **Inline `# SECURITY:` comments** explaining security decisions

## Using These Examples

Each example includes:

- Complete workflow YAML ready to copy
- Inline `# SECURITY:` comments explaining security decisions
- Language-specific variants where applicable
- Security checklist for validation
- Common mistakes and how to avoid them
- Cross-references to relevant security patterns

## Integration Points

These examples reference patterns from:

- [Action Pinning](../action-pinning/index.md) - SHA pinning, Dependabot, automation
- [Token Permissions](../token-permissions/index.md) - Minimal scopes, job-level permissions
- [Secret Management](../secrets/secrets-management/index.md) - OIDC, rotation, scanning
- [Third-Party Actions](../third-party-actions/index.md) - Trust tiers, evaluation
- [Runner Security](../runners/index.md) - Hardening, ephemeral patterns
- [Workflow Patterns](../workflows/triggers/index.md) - Safe triggers, environments, reusable workflows

## Quick Start

1. **Choose** the example that matches your use case
2. **Copy** the workflow YAML to `.github/workflows/`
3. **Customize** the language/framework-specific steps
4. **Review** the security checklist
5. **Test** with `act` or a draft PR
6. **Deploy** to production

For additional guidance, see the [Quick Reference Cheat Sheet](../cheat-sheet/index.md).

### Available Examples

### [Secure CI Workflow](ci-workflow/index.md)

Hardened continuous integration with comprehensive security controls.

**Key Patterns**:

- Fork PR security with two-stage workflows
- Language-specific testing (Node.js, Python, Go)
- Secret scanning prevention
- Minimal GITHUB_TOKEN permissions
- SHA-pinned actions with version comments

**Use Cases**: Test automation, PR validation, pre-merge quality gates

---

### [Release Workflow](release-workflow/index.md)

Signed releases with SLSA provenance and artifact attestations.

**Key Patterns**:

- Keyless signing with OIDC
- SLSA L2/L3 provenance generation
- Artifact attestations for GitHub releases, containers, and NPM packages
- Environment protection for release branches
- Minimal permissions with `id-token: write` and `attestations: write`

**Use Cases**: GitHub releases, container publishing, NPM publishing, signed artifacts

---

### [Deployment Workflow](deployment-workflow/index.md)

OIDC-based cloud deployment with environment protection and automated rollback.

**Key Patterns**:

- OIDC federation to GCP (no stored secrets)
- Environment protection with approval gates and wait timers
- Canary rollout with traffic migration (10% → 100%)
- Container scanning and signing
- Automated rollback on deployment failure

**Use Cases**: Cloud Run deployment, Kubernetes/Helm deployment, multi-environment pipelines, canary releases

---

### [Security Scanning](security-scanning/index.md)

Comprehensive SAST, dependency scanning, container scanning, and SARIF upload.

**Key Patterns**:

- CodeQL SAST with security-extended queries
- Dependency review with severity-based failure
- Container scanning (Trivy, SBOM generation with Syft/Grype)
- Language-specific scanning (Bandit, gosec, ESLint)
- SARIF aggregation and upload to Security tab
- Scheduled vulnerability scanning with issue creation

**Use Cases**: Security validation, compliance scanning, vulnerability detection, scheduled audits

---

### Common Security Controls

All examples use:

- **SHA-pinned actions** with version comments for supply chain security
- **Minimal GITHUB_TOKEN permissions** scoped to job requirements
- **Environment protection** where appropriate (deployments, releases)
- **OIDC federation** for cloud access (no stored secrets)
- **Input validation** and safe trigger patterns
- **Inline `# SECURITY:` comments** explaining security decisions

### Using These Examples

Each example includes:

- Complete workflow YAML ready to copy
- Inline `# SECURITY:` comments explaining security decisions
- Language-specific variants where applicable
- Security checklist for validation
- Common mistakes and how to avoid them
- Cross-references to relevant security patterns

### Integration Points

These examples reference patterns from:

- [Action Pinning](../action-pinning/index.md) - SHA pinning, Dependabot, automation
- [Token Permissions](../token-permissions/index.md) - Minimal scopes, job-level permissions
- [Secret Management](../secrets/secrets-management/index.md) - OIDC, rotation, scanning
- [Third-Party Actions](../third-party-actions/index.md) - Trust tiers, evaluation
- [Runner Security](../runners/index.md) - Hardening, ephemeral patterns
- [Workflow Patterns](../workflows/triggers/index.md) - Safe triggers, environments, reusable workflows

### Quick Start

1. **Choose** the example that matches your use case
2. **Copy** the workflow YAML to `.github/workflows/`
3. **Customize** the language/framework-specific steps
4. **Review** the security checklist
5. **Test** with `act` or a draft PR
6. **Deploy** to production

For additional guidance, see the [Quick Reference Cheat Sheet](../cheat-sheet/index.md).
## References

- [Source Documentation](https://adaptive-enforcement-lab.com/secure/github-actions-security/)
- [AEL Secure](https://adaptive-enforcement-lab.com/secure/)
