---
name: implementation-roadmap
description: >-
  Phased rollout plan for SDLC hardening. Foundation to runtime enforcement in 90 days. Prioritized by risk and audit importance.
---

# Implementation Roadmap

## When to Use This Skill

You can't harden everything at once. Prioritize controls by risk and audit value.

> **Phased Rollout**
>
> Follow the 12-week timeline to avoid disrupting existing workflows. Skip phases at your own risk.
>

Three-month plan from foundation to full enforcement.

---


## Implementation

You can't harden everything at once. Prioritize controls by risk and audit value.

> **Phased Rollout**
>
> Follow the 12-week timeline to avoid disrupting existing workflows. Skip phases at your own risk.
>

Three-month plan from foundation to full enforcement.

---

## Month 1: Foundation

Goal: Core enforcement in place. Evidence collection begins.

### Week 1: Branch Protection

**Tasks**:

- Enable branch protection on `main` and production branches
- Require 1+ approving reviews
- Enable `enforce_admins`
- Require linear history

**Validation**:

```bash
gh api repos/org/repo/branches/main/protection \
  | jq '{reviews: .required_pull_request_reviews, admins: .enforce_admins}'
```

**Documentation**: Update CONTRIBUTING.md with review requirements.

---

### Week 2: CI/CD Status Checks

**Tasks**:

- Create `required-checks.yml` workflow (tests, lint)
- Configure branch protection to require checks
- Test on non-critical repository first

**Workflow**:


*See [examples.md](examples.md) for detailed code examples.*

**Validation**: Open PR, verify checks block merge until passing.

---

### Week 3: GitHub App Setup

**Tasks**:

- Create GitHub App for automation (see [Setup Guide](../../secure/github-apps/index.md))
- Configure permissions (releases, PRs, contents)
- Generate and store private key in secrets
- Replace first PAT usage in workflows

**Validation**:

```yaml
- name: Test app token
  uses: actions/create-github-app-token@v2
  with:
    app-id: ${{ secrets.APP_ID }}
    private-key: ${{ secrets.PRIVATE_KEY }}
```

**Migration tracking**: Document remaining PAT usages for month 2.

---

### Week 4: Evidence Archive

**Tasks**:

- Set up GCS bucket with lifecycle policy (3 year retention)
- Create monthly evidence collection workflow
- Archive first month's data (branch protection config, merged PRs)

**Workflow**:


*See [examples.md](examples.md) for detailed code examples.*

**Validation**: Verify files appear in GCS bucket.

---

## Month 2: Hardening

Goal: Add secrets detection, commit signing, and SBOM generation.

### Week 5: Secrets Detection

**Tasks**:

- Add TruffleHog to `.pre-commit-config.yaml`
- Deploy pre-commit config to all repositories
- Add secrets scan to CI workflow
- Document bypass procedure (`--no-verify` tracking)

**Pre-commit hook**:

```yaml
repos:
  - repo: https://github.com/trufflesecurity/trufflehog
    rev: v3.63.0
    hooks:
      - id: trufflehog
        entry: trufflehog filesystem --fail --no-update
```

**Validation**: Attempt to commit AWS key, verify block.

See [Pre-commit Security Gates](../../blog/posts/2025-12-04-pre-commit-security-gates.md) for full implementation.

---

### Week 6: Signed Commits

**Tasks**:

- Generate GPG keys for core team
- Add public keys to GitHub
- Configure Git to sign commits automatically
- Enable `required_signatures` on protected branches

**Configuration**:

```bash
git config --global user.signingkey YOUR_GPG_KEY_ID
git config --global commit.gpgsign true
```

**Validation**:

```bash
git log --show-signature | grep "Good signature"
```

See [Commit Signing](../commit-signing/commit-signing.md) for setup guide.

---

### Week 7: SBOM Generation

**Tasks**:

- Add Syft/Trivy to build pipelines
- Generate SBOM for each container build
- Upload SBOMs to artifact storage
- Verify license compliance (no GPL in proprietary code)

**Workflow**:


*See [examples.md](examples.md) for detailed code examples.*

**Validation**: Download artifact, verify SBOM contains expected dependencies.

See [SBOM Generation](../../secure/sbom/sbom-generation.md) for full implementation.

---

### Week 8: Complete PAT Migration

**Tasks**:

- Audit all remaining PAT usages (`grep -r GITHUB_TOKEN .github/`)
- Create additional GitHub Apps for specific use cases
- Replace all PATs with app tokens
- Revoke old PATs

**Validation**: No PATs referenced in active workflows.

---

## Month 3: Validation & Policy-as-Code

Goal: Simulate audit, fix gaps, add runtime enforcement.

### Week 9: Vulnerability Scanning

**Tasks**:

- Add Trivy/Grype container scanning to CI
- Set severity threshold (HIGH/CRITICAL block merge)
- Configure vulnerability database auto-update

**Workflow**:

```yaml
- name: Scan container
  run: |
    trivy image --severity HIGH,CRITICAL --exit-code 1 \
      gcr.io/project/app:${{ github.sha }}
```

**Validation**: Introduce test vulnerability, verify build fails.

See [Zero-Vulnerability Pipelines](../../blog/posts/2025-12-15-zero-vulnerability-pipelines.md).

---

### Week 10: Policy-as-Code (Kyverno)

**Tasks**:

- Deploy Kyverno to Kubernetes clusters
- Install Policy Reporter for observability
- Implement core policies (resource limits, image sources, labels)
- Configure audit mode first, then enforcement mode

**Core policy**:


*See [examples.md](examples.md) for detailed code examples.*

**Validation**: Deploy pod without limits, verify rejection.

See [Policy-as-Code with Kyverno](../../blog/posts/2025-12-13-policy-as-code-kyverno.md) for end-to-end implementation.

---

### Week 11: Audit Simulation

**Tasks**:

- Pull evidence like an auditor would (API queries for March data)
- Generate summary report (PR reviews, check results, signed commits)
- Identify gaps in evidence or controls
- Document findings and remediation plan

**Simulation script**:

```bash
# Verify branch protection
gh api repos/org/repo/branches/main/protection

# Sample March PRs
gh api 'repos/org/repo/pulls?state=closed&base=main' \
  --jq '.[] | select(.merged_at | startswith("2025-03"))'

# Check signature coverage
./scripts/signature-coverage.sh 2025-03-01 2025-04-01
```

**Validation**: Evidence collection succeeds for sampled period.

---

### Week 12: Remediation & Runbook

**Tasks**:

- Fix gaps identified in simulation
- Create runbook for responding to audit requests
- Train team on SDLC controls (why they exist, how to use them)
- Document exception processes (emergency bypass, post-review)

**Runbook sections**:

- How to retrieve branch protection evidence
- How to query PR review history
- How to generate compliance reports
- Exception request template
- Bypass logging procedure

**Validation**: Team can retrieve evidence without assistance.

---

## Next Steps

- **[Execution Guide](execution.md)** - Progress tracking, audit readiness criteria, rollback planning, cost estimation, success metrics

---

*Week 1: Protection enabled. Week 4: Evidence collected. Week 12: Audit simulation passed. Controls enforced. System hardened.*

### Month 1: Foundation

Goal: Core enforcement in place. Evidence collection begins.

### Week 1: Branch Protection

**Tasks**:

- Enable branch protection on `main` and production branches
- Require 1+ approving reviews
- Enable `enforce_admins`
- Require linear history

**Validation**:

```bash
gh api repos/org/repo/branches/main/protection \
  | jq '{reviews: .required_pull_request_reviews, admins: .enforce_admins}'
```

**Documentation**: Update CONTRIBUTING.md with review requirements.

---

### Week 2: CI/CD Status Checks

**Tasks**:

- Create `required-checks.yml` workflow (tests, lint)
- Configure branch protection to require checks
- Test on non-critical repository first

**Workflow**:


*See [examples.md](examples.md) for detailed code examples.*

**Validation**: Open PR, verify checks block merge until passing.

---

### Week 3: GitHub App Setup

**Tasks**:

- Create GitHub App for automation (see [Setup Guide](../../secure/github-apps/index.md))
- Configure permissions (releases, PRs, contents)
- Generate and store private key in secrets
- Replace first PAT usage in workflows

**Validation**:

```yaml
- name: Test app token
  uses: actions/create-github-app-token@v2
  with:
    app-id: ${{ secrets.APP_ID }}
    private-key: ${{ secrets.PRIVATE_KEY }}
```

**Migration tracking**: Document remaining PAT usages for month 2.

---

### Week 4: Evidence Archive

**Tasks**:

- Set up GCS bucket with lifecycle policy (3 year retention)
- Create monthly evidence collection workflow
- Archive first month's data (branch protection config, merged PRs)

**Workflow**:


*See [examples.md](examples.md) for detailed code examples.*

**Validation**: Verify files appear in GCS bucket.

---

### Month 2: Hardening

Goal: Add secrets detection, commit signing, and SBOM generation.

### Week 5: Secrets Detection

**Tasks**:

- Add TruffleHog to `.pre-commit-config.yaml`
- Deploy pre-commit config to all repositories
- Add secrets scan to CI workflow
- Document bypass procedure (`--no-verify` tracking)

**Pre-commit hook**:

```yaml
repos:
  - repo: https://github.com/trufflesecurity/trufflehog
    rev: v3.63.0
    hooks:
      - id: trufflehog
        entry: trufflehog filesystem --fail --no-update
```

**Validation**: Attempt to commit AWS key, verify block.

See [Pre-commit Security Gates](../../blog/posts/2025-12-04-pre-commit-security-gates.md) for full implementation.

---

### Week 6: Signed Commits

**Tasks**:

- Generate GPG keys for core team
- Add public keys to GitHub
- Configure Git to sign commits automatically
- Enable `required_signatures` on protected branches

**Configuration**:

```bash
git config --global user.signingkey YOUR_GPG_KEY_ID
git config --global commit.gpgsign true
```

**Validation**:

```bash
git log --show-signature | grep "Good signature"
```

See [Commit Signing](../commit-signing/commit-signing.md) for setup guide.

---

### Week 7: SBOM Generation

**Tasks**:

- Add Syft/Trivy to build pipelines
- Generate SBOM for each container build
- Upload SBOMs to artifact storage
- Verify license compliance (no GPL in proprietary code)

**Workflow**:


*See [examples.md](examples.md) for detailed code examples.*

**Validation**: Download artifact, verify SBOM contains expected dependencies.

See [SBOM Generation](../../secure/sbom/sbom-generation.md) for full implementation.

---

### Week 8: Complete PAT Migration

**Tasks**:

- Audit all remaining PAT usages (`grep -r GITHUB_TOKEN .github/`)
- Create additional GitHub Apps for specific use cases
- Replace all PATs with app tokens
- Revoke old PATs

**Validation**: No PATs referenced in active workflows.

---

### Month 3: Validation & Policy-as-Code

Goal: Simulate audit, fix gaps, add runtime enforcement.

### Week 9: Vulnerability Scanning

**Tasks**:

- Add Trivy/Grype container scanning to CI
- Set severity threshold (HIGH/CRITICAL block merge)
- Configure vulnerability database auto-update

**Workflow**:

```yaml
- name: Scan container
  run: |
    trivy image --severity HIGH,CRITICAL --exit-code 1 \
      gcr.io/project/app:${{ github.sha }}
```

**Validation**: Introduce test vulnerability, verify build fails.

See [Zero-Vulnerability Pipelines](../../blog/posts/2025-12-15-zero-vulnerability-pipelines.md).

---

### Week 10: Policy-as-Code (Kyverno)

**Tasks**:

- Deploy Kyverno to Kubernetes clusters
- Install Policy Reporter for observability
- Implement core policies (resource limits, image sources, labels)
- Configure audit mode first, then enforcement mode

**Core policy**:


*See [examples.md](examples.md) for detailed code examples.*

**Validation**: Deploy pod without limits, verify rejection.

See [Policy-as-Code with Kyverno](../../blog/posts/2025-12-13-policy-as-code-kyverno.md) for end-to-end implementation.

---

### Week 11: Audit Simulation

**Tasks**:

- Pull evidence like an auditor would (API queries for March data)
- Generate summary report (PR reviews, check results, signed commits)
- Identify gaps in evidence or controls
- Document findings and remediation plan

**Simulation script**:

```bash
# Verify branch protection
gh api repos/org/repo/branches/main/protection

# Sample March PRs
gh api 'repos/org/repo/pulls?state=closed&base=main' \
  --jq '.[] | select(.merged_at | startswith("2025-03"))'

# Check signature coverage
./scripts/signature-coverage.sh 2025-03-01 2025-04-01
```

**Validation**: Evidence collection succeeds for sampled period.

---

### Week 12: Remediation & Runbook

**Tasks**:

- Fix gaps identified in simulation
- Create runbook for responding to audit requests
- Train team on SDLC controls (why they exist, how to use them)
- Document exception processes (emergency bypass, post-review)

**Runbook sections**:

- How to retrieve branch protection evidence
- How to query PR review history
- How to generate compliance reports
- Exception request template
- Bypass logging procedure

**Validation**: Team can retrieve evidence without assistance.

---

### Next Steps

- **[Execution Guide](execution.md)** - Progress tracking, audit readiness criteria, rollback planning, cost estimation, success metrics

---

*Week 1: Protection enabled. Week 4: Evidence collected. Week 12: Audit simulation passed. Controls enforced. System hardened.*


## Examples

See [examples.md](examples.md) for code examples.


## Full Reference

See [reference.md](reference.md) for complete documentation.
## References

- [Source Documentation](https://adaptive-enforcement-lab.com/enforce/implementation-roadmap/)
- [AEL Enforce](https://adaptive-enforcement-lab.com/enforce/)
