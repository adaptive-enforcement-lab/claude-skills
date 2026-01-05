---
name: enforce - Reference
description: Complete reference for Enforce
---

# Enforce - Reference

This is the complete reference documentation extracted from the source.

# Enforce

Making security mandatory through automation.

> **Enforcement Over Education**
>
>
> **If you can't enforce it, it doesn't matter.** Documentation, training, and recommendations don't scale. Security controls that can be bypassed eventually will be bypassed.
>

## Overview

This section covers the **enforcement mechanisms** that make security policies mandatory, auditable, and impossible to ignore.

These controls pass SOC 2, ISO 27001, and PCI-DSS audits by shifting security left and making compliance automatic.

## Secure vs Enforce

Understanding the distinction:

- **Secure** ([see Secure](../secure/index.md)): Find and fix security issues
  - Vulnerability scanners that *identify* CVEs
  - SBOM generators that *document* dependencies
  - Security tools that *discover* weaknesses

- **Enforce** (this section): Make security mandatory through automation
  - Branch protection that *requires* reviews
  - Pre-commit hooks that *block* violations
  - Status checks that *prevent* merges
  - Policy-as-code that *rejects* non-compliant resources
  - SLSA provenance that *attests* build integrity

**Litmus test**: Can this be bypassed?

- If **yes** → Belongs in Enforce (make it mandatory)
- If **no** → Belongs in Secure (it's a finding/fix tool)

## What You'll Find Here

### Branch Protection

Require code reviews, passing status checks, commit signatures, and up-to-date branches before merging to protected branches.

**Why it matters**: Prevents direct commits to main, ensures peer review, and blocks broken code from reaching production.

**Key topics**:

- Required reviewers and review counts
- Required status checks (tests, security scans, linting)
- Commit signature verification
- Administrator bypass restrictions

### Pre-commit Hooks

Client-side and server-side hooks that block commits violating security policies, code standards, or compliance requirements.

**Why it matters**: Catch violations at commit time, before CI/CD ever runs. Fastest possible feedback loop.

**Key topics**:

- Secret detection (prevent credential leaks)
- Code formatting and linting enforcement
- Conventional commit enforcement
- Custom validation hooks

### Status Checks

GitHub status checks that gate pull request merges on passing tests, security scans, policy validation, and approval workflows.

**Why it matters**: Automated quality gates that prevent human error and enforce organizational standards.

**Key topics**:

- Required vs optional checks
- Check configuration patterns
- Failure handling and retries
- Matrix strategy checks

### Policy-as-Code

Runtime admission control using Kyverno and OPA to enforce security policies, compliance requirements, and operational standards in Kubernetes clusters.

**Why it matters**: Prevent misconfigured resources from ever being admitted to the cluster. Policy enforcement at the API server level cannot be bypassed.

**Key topics**:

- Kyverno policy patterns (validate, mutate, generate)
- OPA Gatekeeper constraints
- Local development validation
- CI integration (policy testing)
- Runtime deployment and monitoring
- Multi-source policy management

### SLSA Provenance

Generate cryptographically signed attestations proving the integrity of build processes, source code, and artifacts.

**Why it matters**: Supply chain attacks (SolarWinds, Log4Shell) exploit build process compromise. SLSA provenance proves your builds are tamper-proof.

**Key topics**:

- SLSA levels (1-4)
- Provenance generation with GitHub Actions
- Artifact signing and verification
- Rekor transparency log integration

### Testing Enforcement

Enforce minimum code coverage thresholds, require tests for new code, and block PRs that reduce coverage.

**Why it matters**: Code without tests is code that breaks in production. Enforce testing discipline at merge time.

**Key topics**:

- Coverage thresholds (80%+ recommended)
- Coverage enforcement in status checks
- Differential coverage (new code only)
- Test quality patterns

### Audit & Compliance

Automated collection of audit evidence, compliance documentation, and attestation generation for SOC 2, ISO 27001, and PCI-DSS audits.

**Why it matters**: Manual audit evidence collection is error-prone and time-consuming. Automate evidence generation to pass audits without scrambling.

**Key topics**:

- Evidence collection automation
- Audit log aggregation
- Compliance reporting
- Attestation workflows

## Common Workflows

### 1. Enforce Branch Protection

```bash
# Require 2 reviews, passing tests, and commit signatures
gh api repos/org/repo/branches/main/protection \
  --method PUT \
  --field required_pull_request_reviews[required_approving_review_count]=2 \
  --field required_status_checks[strict]=true \
  --field required_status_checks[contexts][]=test \
  --field required_status_checks[contexts][]=security-scan \
  --field required_signatures[enabled]=true
```

### 2. Pre-commit Hook for Secret Detection

```yaml
# .pre-commit-config.yaml
repos:
  - repo: https://github.com/trufflesecurity/trufflehog
    rev: v3.63.0
    hooks:
      - id: trufflehog
        name: TruffleHog
        entry: bash -c 'trufflehog git file://. --since-commit HEAD --only-verified --fail'
```

### 3. Kyverno Policy Enforcement

```yaml
# Enforce resource limits on all pods
apiVersion: kyverno.io/v1
kind: ClusterPolicy
metadata:
  name: require-resource-limits
spec:
  validationFailureAction: Enforce
  rules:
    - name: check-resource-limits
      match:
        resources:
          kinds:
            - Pod
      validate:
        message: "Resource limits are required"
        pattern:
          spec:
            containers:
              - resources:
                  limits:
                    memory: "?*"
                    cpu: "?*"
```

### 4. SLSA Provenance Generation

```yaml
# .github/workflows/release.yml
permissions:
  id-token: write  # Required for SLSA provenance
  contents: write

jobs:
  release:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - name: Build artifact
        run: make build
      - name: Generate SLSA provenance
        uses: slsa-framework/slsa-github-generator/.github/workflows/generator_generic_slsa3.yml@v1.9.0
        with:
          artifacts: dist/*
```

## Enforcement Hierarchy

Enforcement controls work in layers:

1. **Pre-commit hooks** (fastest feedback)
   - Catch violations before commit
   - Developer workstation enforcement
   - Can be bypassed with `--no-verify` (use server-side for critical policies)

2. **Status checks** (PR merge gates)
   - Automated quality gates
   - Enforce in CI/CD pipeline
   - Cannot be bypassed without admin override

3. **Branch protection** (repository controls)
   - Prevent direct commits
   - Require reviews and status checks
   - Restrict who can merge

4. **Policy-as-code** (runtime enforcement)
   - Admission control at API server
   - Cannot be bypassed by developers
   - Mutate or reject non-compliant resources

**Best practice**: Layer multiple enforcement mechanisms. Pre-commit hooks for fast feedback, status checks for automation, policy-as-code for runtime protection.

## Integration with Secure

Enforcement is only effective when paired with security tooling:

1. **Find vulnerabilities** ([Secure](../secure/index.md)) → **Block deployment** (Enforce)
2. **Generate SBOM** ([Secure](../secure/index.md)) → **Require SBOM in PR** (Enforce)
3. **Run Scorecard** ([Secure](../secure/index.md)) → **Enforce minimum score** (Enforce)
4. **Scan containers** ([Secure](../secure/index.md)) → **Block vulnerable images** (Enforce)

## Implementation Roadmap

See [Implementation Roadmap](implementation-roadmap/index.md) for phased rollout:

1. **Phase 1**: Branch protection (1 week)
2. **Phase 2**: Status checks (2 weeks)
3. **Phase 3**: Pre-commit hooks (1 week)
4. **Phase 4**: Policy-as-code (4 weeks)
5. **Phase 5**: SLSA provenance (2 weeks)

**Total timeline**: 10 weeks for complete enforcement stack.

## Getting Started

1. **Start with branch protection**: Require reviews and passing tests
2. **Add status checks**: Block PRs that fail security scans
3. **Deploy pre-commit hooks**: Catch secrets before they're committed
4. **Layer on policy-as-code**: Enforce runtime compliance
5. **Add SLSA provenance**: Prove build integrity

## Common Challenges

### "Enforcement slows down developers"

**Reality**: Finding and fixing issues in production is 10x slower than catching them in CI.

**Solution**: Layer enforcement to provide fast feedback (pre-commit hooks) before slow feedback (CI/CD).

### "Developers will just bypass the controls"

**Reality**: Some controls (like pre-commit hooks) can be bypassed. Others (like policy-as-code) cannot.

**Solution**: Use client-side enforcement for fast feedback, server-side enforcement for critical policies.

### "We need exceptions for emergencies"

**Reality**: Every organization needs break-glass procedures.

**Solution**: Document exception processes. Use temporary admin overrides with audit trails, not permanent bypasses.

## Related Content

- [Secure](../secure/index.md): Find and fix security issues
- [Build](../build/index.md): CI/CD pipelines and release automation
- [Patterns](../patterns/index.md): Reusable enforcement patterns

## Tags

Browse all content tagged with policy-enforcement, automation, compliance, and security on the [Tags](../tags.md) page.

