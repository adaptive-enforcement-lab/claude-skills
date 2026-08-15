# Secure

Source: https://adaptive-enforcement-lab.com/secure/

Finding and fixing security issues before they become incidents.

> **Security as a Process**
>
>
> Security isn't a one-time scan. It's a continuous process of finding vulnerabilities, generating evidence, and enabling secure development workflows.
>

## Overview

This section covers the **tools and practices** for discovering and remediating security issues in code, dependencies, containers, and supply chains.

## Secure vs Enforce

Understanding the distinction:

- **Secure** (this section): Find and fix security issues
  - Vulnerability scanners that *identify* CVEs
  - SBOM generators that *document* dependencies
  - Security tools that *discover* weaknesses
  - GitHub Apps that *provide* secure authentication

- **Enforce** ([see Enforce](../enforce/index.md)): Make security mandatory through automation
  - Branch protection that *requires* reviews
  - Pre-commit hooks that *block* violations
  - Status checks that *prevent* merges
  - Policy-as-code that *enforces* runtime compliance

**Litmus test**: Can this be bypassed?

- If **no** → It's a Secure tool (finding/fixing)
- If **yes** → It belongs in Enforce (making mandatory)

## What You'll Find Here

### GitHub Apps

Secure authentication for automated workflows. GitHub Apps offer granular permissions, auditable actions, and organization-level credential management.

**Why it matters**: Pass SOC 2 and ISO 27001 audits by replacing PATs with trackable, scoped authentication.

**Key topics**:

- Creating and configuring GitHub Apps
- Permission patterns for common workflows
- Credential storage and rotation
- Installation scopes and security

### Vulnerability Scanning

Find CVEs in dependencies, containers, and runtime environments before they reach production.

**Why it matters**: 84% of breaches exploit known vulnerabilities with available patches (Verizon DBIR 2024).

**Key topics**:

- Dependency scanning (npm, go mod, pip)
- Container image scanning (Trivy, Grype)
- Runtime vulnerability detection
- Remediation workflows

### SBOM (Software Bill of Materials)

Generate machine-readable inventories of all software components, dependencies, and transitive dependencies.

**Why it matters**: Executive Order 14028 and European Cyber Resilience Act require SBOMs for supply chain transparency.

**Key topics**:

- SBOM generation with Syft
- SPDX and CycloneDX formats
- Embedding SBOMs in container images
- Automated SBOM workflows

### Go Security Tooling

Leverage specialized security tooling for Go projects, including static analysis, vulnerability detection, and compliance checks.

**Why it matters**: Go's standard library security model requires specific tooling that understands Go's unique characteristics.

**Key topics**:

- `govulncheck` for vulnerability scanning
- `gosec` for static security analysis
- Go-specific SBOM generation
- CI/CD integration patterns

### Scorecard

OpenSSF Scorecard automated security checks for open-source best practices, SLSA compliance, and supply chain security.

**Why it matters**: Quantifiable security posture that passes compliance audits and satisfies customer security questionnaires.

**Key topics**:

- Scorecard check categories
- Achieving high scores (8+/10)
- Workflow examples
- Badge integration

## Common Workflows

### 1. Continuous Vulnerability Scanning

```yaml
# .github/workflows/security-scan.yml
name: Security Scan
on:
  push:
  schedule:
    - cron: '0 2 * * *'  # Daily at 2 AM

jobs:
  scan:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - name: Scan dependencies
        run: trivy fs --scanners vuln .
      - name: Scan containers
        run: trivy image myapp:latest
```

### 2. SBOM Generation on Release

```yaml
# Generate SBOM on every release
- name: Generate SBOM
  run: syft packages . -o spdx-json > sbom.spdx.json
- name: Attach to release
  run: gh release upload ${{ github.ref_name }} sbom.spdx.json
```

### 3. GitHub App Authentication

```yaml
# Secure authentication for cross-repo workflows
- uses: actions/create-github-app-token@v1
  id: app-token
  with:
    app-id: ${{ vars.APP_ID }}
    private-key: ${{ secrets.APP_PRIVATE_KEY }}
    owner: adaptive-enforcement-lab
- uses: actions/checkout@v4
  with:
    token: ${{ steps.app-token.outputs.token }}
```

## Integration with Enforce

Security findings are only valuable if they prevent insecure code from reaching production:

1. **Find vulnerabilities** (Secure) → **Block deployment** (Enforce)
2. **Generate SBOM** (Secure) → **Require SBOM in PR** (Enforce)
3. **Run Scorecard** (Secure) → **Enforce minimum score** (Enforce)
4. **Scan containers** (Secure) → **Block vulnerable images** (Enforce)

See [Enforce](../enforce/index.md) for enforcement mechanisms.

## Getting Started

1. **Start with GitHub Apps**: Replace PATs with secure, auditable authentication
2. **Add vulnerability scanning**: Catch known CVEs before they deploy
3. **Generate SBOMs**: Document your supply chain for compliance
4. **Run Scorecard**: Measure and improve security posture
5. **Layer on enforcement**: Make findings actionable with Enforce patterns

## Related Content

- [Enforce](../enforce/index.md): Make security mandatory through automation
- [Build](../build/index.md): CI/CD pipelines and release automation
- [Patterns](../patterns/index.md): Reusable security patterns

## Tags

Browse all content tagged with security, automation, supply-chain, and compliance on the [Tags](../tags.md) page.
