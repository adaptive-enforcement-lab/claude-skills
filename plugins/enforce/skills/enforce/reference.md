# Enforce — Full Reference

Generated from adaptive-enforcement-lab.com. For a scannable index with links to the live docs, see SKILL.md in this skill.

## Overview

Making security mandatory through automation.

> **Enforcement Over Education**
>
>
> **If you can't enforce it, it doesn't matter.** Documentation, training, and recommendations don't scale. Security controls that can be bypassed eventually will be bypassed.
>

### Overview

This section covers the **enforcement mechanisms** that make security policies mandatory, auditable, and impossible to ignore.

These controls pass SOC 2, ISO 27001, and PCI-DSS audits by shifting security left and making compliance automatic.

### Secure vs Enforce

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

### What You'll Find Here

#### Branch Protection

Enforce code reviews, status checks, commit signatures, and up-to-date branches on protected branches.

**Why it matters**: Prevents direct commits to main, ensures peer review, and blocks broken code from reaching production.

**Key topics**:

- Required reviewers and review counts
- Required status checks (tests, security scans, linting)
- Commit signature verification
- Administrator bypass restrictions

#### Pre-commit Hooks

Block commits violating security policies, code standards, or compliance, using client-side and server-side hooks.

**Why it matters**: Catch violations at commit time, before CI/CD ever runs. Fastest possible feedback loop.

**Key topics**:

- Secret detection (prevent credential leaks)
- Code formatting and linting enforcement
- Conventional commit enforcement
- Custom validation hooks

#### Status Checks

Gate pull request merges with GitHub status checks, requiring passing tests, security scans, policy validation, and approval.

**Why it matters**: Automated quality gates that prevent human error and enforce organizational standards.

**Key topics**:

- Required vs optional checks
- Check configuration patterns
- Failure handling and retries
- Matrix strategy checks

#### Policy-as-Code

Enforce security policies, compliance, and operational standards in Kubernetes clusters via runtime admission control with Kyverno and OPA.

**Why it matters**: Prevent misconfigured resources from ever being admitted to the cluster. Policy enforcement at the API server level cannot be bypassed.

**Key topics**:

- Kyverno policy patterns (validate, mutate, generate)
- OPA Gatekeeper constraints
- Local development validation
- CI integration (policy testing)
- Runtime deployment and monitoring
- Multi-source policy management

#### SLSA Provenance

Generate cryptographically signed attestations. Prove build process, source code, and artifact integrity.

**Why it matters**: Supply chain attacks (SolarWinds, Log4Shell) exploit build process compromise. SLSA provenance proves your builds are tamper-proof.

**Key topics**:

- SLSA levels (1-4)
- Provenance generation with GitHub Actions
- Artifact signing and verification
- Rekor transparency log integration

#### Testing Enforcement

Enforce minimum code coverage, require tests for new code, and block PRs that reduce coverage.

**Why it matters**: Code without tests is code that breaks in production. Enforce testing discipline at merge time.

**Key topics**:

- Coverage thresholds (80%+ recommended)
- Coverage enforcement in status checks
- Differential coverage (new code only)
- Test quality patterns

#### Audit & Compliance

Automate audit evidence collection, compliance documentation, and attestation generation for SOC 2, ISO 27001, and PCI-DSS audits.

**Why it matters**: Manual audit evidence collection is error-prone and time-consuming. Automate evidence generation to pass audits without scrambling.

**Key topics**:

- Evidence collection automation
- Audit log aggregation
- Compliance reporting
- Attestation workflows

### Common Workflows

#### 1. Enforce Branch Protection

```bash
## Require 2 reviews, passing tests, and commit signatures
gh api repos/org/repo/branches/main/protection \
  --method PUT \
  --field required_pull_request_reviews[required_approving_review_count]=2 \
  --field required_status_checks[strict]=true \
  --field required_status_checks[contexts][]=test \
  --field required_status_checks[contexts][]=security-scan \
  --field required_signatures[enabled]=true
```

#### 2. Pre-commit Hook for Secret Detection

```yaml
## .pre-commit-config.yaml
repos:
  - repo: https://github.com/trufflesecurity/trufflehog
    rev: v3.63.0
    hooks:
      - id: trufflehog
        name: TruffleHog
        entry: bash -c 'trufflehog git file://. --since-commit HEAD --only-verified --fail'
```

#### 3. Kyverno Policy Enforcement

```yaml
## Enforce resource limits on all pods
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

#### 4. SLSA Provenance Generation

```yaml
## .github/workflows/release.yml
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

### Enforcement Hierarchy

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

### Integration with Secure

Enforcement is only effective when paired with security tooling:

1. **Find vulnerabilities** ([Secure](../secure/index.md)) → **Block deployment** (Enforce)
2. **Generate SBOM** ([Secure](../secure/index.md)) → **Require SBOM in PR** (Enforce)
3. **Run Scorecard** ([Secure](../secure/index.md)) → **Enforce minimum score** (Enforce)
4. **Scan containers** ([Secure](../secure/index.md)) → **Block vulnerable images** (Enforce)

### Implementation Roadmap

See [Implementation Roadmap](implementation-roadmap/index.md) for phased rollout:

1. **Phase 1**: Branch protection (1 week)
2. **Phase 2**: Status checks (2 weeks)
3. **Phase 3**: Pre-commit hooks (1 week)
4. **Phase 4**: Policy-as-code (4 weeks)
5. **Phase 5**: SLSA provenance (2 weeks)

**Total timeline**: 10 weeks for complete enforcement stack.

### Getting Started

1. **Start with branch protection**: Require reviews and passing tests
2. **Add status checks**: Block PRs that fail security scans
3. **Deploy pre-commit hooks**: Catch secrets before they're committed
4. **Layer on policy-as-code**: Enforce runtime compliance
5. **Add SLSA provenance**: Prove build integrity

### Common Challenges

#### "Enforcement slows down developers"

**Reality**: Finding and fixing issues in production is 10x slower than catching them in CI.

**Solution**: Layer enforcement to provide fast feedback (pre-commit hooks) before slow feedback (CI/CD).

#### "Developers will just bypass the controls"

**Reality**: Some controls (like pre-commit hooks) can be bypassed. Others (like policy-as-code) cannot.

**Solution**: Use client-side enforcement for fast feedback, server-side enforcement for critical policies.

#### "We need exceptions for emergencies"

**Reality**: Every organization needs break-glass procedures.

**Solution**: Document exception processes. Use temporary admin overrides with audit trails, not permanent bypasses.

### Related Content

- [Secure](../secure/index.md): Find and fix security issues
- [Build](../build/index.md): CI/CD pipelines and release automation
- [Patterns](../patterns/index.md): Reusable enforcement patterns

### Tags

Browse all content tagged with policy-enforcement, automation, compliance, and security on the [Tags](../tags.md) page.

## Branch Protection Enforcement Patterns

Comprehensive branch protection configuration patterns with…

Branch protection rules transform security policies into automated enforcement. No manual oversight. No trust required.

> **Core Security Control**
>
> Branch protection is foundational. Without it, code reviews are optional, status checks are suggestions, and audit trails are worthless.
>

GitHub enforces the rules. Terraform maintains consistency at scale. GitHub Apps detect and remediate drift.

#### The Enforcement Gap

Most organizations have branch protection policies. Few enforce them consistently.

**The Problem**:

- New repositories inherit no protection
- Developers disable protection during incidents, forget to re-enable
- Configuration drift across 100+ repositories
- No automated detection when protection is weakened
- Exceptions bypass controls without audit trails

**The Solution**:

Automated enforcement with multiple defense layers:

1. **Security tier templates** - Standardized configurations for different risk levels
2. **Infrastructure as Code** - Terraform/OpenTofu modules for consistent deployment
3. **GitHub App enforcement** - Automated drift detection and remediation
4. **Audit reporting** - Compliance evidence collection
5. **Formalized bypass controls** - Time-boxed exceptions with approval workflows

---

#### Security Tiers

Different repositories require different protection levels.

| Tier | Use Case | Enforcement Level |
|------|----------|-------------------|
| **Standard** | Internal tools, documentation | Required reviews, basic status checks |
| **Enhanced** | Production services, customer-facing apps | Multi-reviewer, comprehensive checks, code owners |
| **Maximum** | Security-critical, compliance-regulated | Full enforcement, no admin bypass, mandatory signing |

> **Right-Sized Security**
>
> Not all repositories need maximum protection. Documentation repos can use Standard tier. Production infrastructure requires Maximum tier. Choose based on blast radius.
>

See **[Security Tiers](security-tiers.md)** for detailed configuration templates.

---

#### Architecture Overview

```mermaid
graph TD
    T[Terraform Module] -->|Applies| BP[Branch Protection Rules]
    GA[GitHub App] -->|Monitors| BP
    GA -->|Detects| DRIFT[Configuration Drift]
    DRIFT -->|Triggers| REM[Automated Remediation]
    REM -->|Restores| BP
    BP -->|Enforces| PR[Pull Requests]
    PR -->|Generates| AUDIT[Audit Evidence]

    %% Ghostty Hardcore Theme
    style T fill:#a7e22e,color:#1b1d1e
    style GA fill:#65d9ef,color:#1b1d1e
    style DRIFT fill:#f92572,color:#1b1d1e
    style BP fill:#fd971e,color:#1b1d1e

```

**Key Components**:

- **Terraform modules** - Declare protection rules as code
- **GitHub Apps** - Monitor and enforce compliance organization-wide
- **Drift detection** - Identify unauthorized changes
- **Automated remediation** - Restore protection without manual intervention
- **Audit collection** - Capture evidence for compliance reporting

---

#### What You'll Learn

This section covers comprehensive branch protection enforcement:

##### Configuration & Standards

- **[Security Tiers](security-tiers.md)** - Tiered protection templates for different risk levels
- **[Branch Protection Rules](branch-protection.md)** - Detailed configuration reference

##### Infrastructure as Code

- **[OpenTofu Modules](opentofu-modules.md)** - OpenTofu-specific patterns and considerations
- **[Multi-Repo Management](multi-repo-management.md)** - Patterns for 100+ repository enforcement

##### GitHub App Enforcement

- **[GitHub App Enforcement](github-app-enforcement.md)** - Centralized enforcement with GitHub Apps
- **[Enforcement Workflows](enforcement-workflows.md)** - Automated workflows for organization-wide compliance
- **[Drift Detection](drift-detection.md)** - Detecting and remediating configuration drift

##### Audit & Compliance

- **[Audit Evidence](audit-evidence.md)** - Collecting and storing compliance evidence
- **[Compliance Reporting](compliance-reporting.md)** - Automated reporting for SOC 2, ISO 27001, PCI-DSS
- **[Verification Scripts](verification-scripts.md)** - Audit preparation and continuous monitoring

##### Bypass Controls

- **[Bypass Controls](bypass-controls.md)** - Formalized bypass procedures with approval workflows
- **[Emergency Access](emergency-access.md)** - Break-glass procedures for production incidents
- **[Exception Management](exception-management.md)** - Managing permanent and temporary exceptions

##### Operations

- **[Troubleshooting](troubleshooting.md)** - Common issues and solutions

---

#### Quick Start

##### Step 1: Choose Your Security Tier

Start with **[Security Tiers](security-tiers.md)** to select the appropriate protection level for your repositories.

##### Step 2: Apply Protection

**Manual (single repository)**:

```bash
gh api --method PUT \
  repos/org/repo/branches/main/protection \
  --input protection-config.json
```

**Automated (organization-wide)**:

##### Step 3: Monitor Compliance

Deploy **[GitHub App Enforcement](github-app-enforcement.md)** to detect drift and maintain compliance.

##### Step 4: Collect Evidence

Implement **[Audit Evidence](audit-evidence.md)** patterns for compliance reporting.

---

#### Real-World Impact

**Before comprehensive enforcement**:

- 40% of repositories had no branch protection
- Admin bypass enabled on 80% of protected repositories
- Zero visibility into protection changes
- Manual verification before each audit (2 weeks of effort)

**After comprehensive enforcement**:

- 100% of repositories protected with tier-appropriate rules
- Automated drift detection and remediation within 5 minutes
- Complete audit trail of all protection changes
- Continuous compliance verification (zero manual effort)

**Key Metrics**:

- Configuration drift detected: **< 5 minutes**
- Automated remediation: **< 1 minute**
- Audit preparation time: **2 weeks → 15 minutes**

---

#### Prerequisites

- GitHub organization with admin access
- Terraform or OpenTofu (for IaC deployment)
- GitHub App with appropriate permissions (for automated enforcement)
- Basic understanding of Git workflow and branch protection concepts

---

#### Architecture Principles

##### 1. Defense in Depth

Multiple enforcement layers: local configuration, drift detection, audit verification.

##### 2. Automation Over Documentation

Don't document the policy. Enforce it automatically.

##### 3. Tier-Based Configuration

Standard, Enhanced, Maximum tiers prevent both under-protection and over-restriction.

##### 4. Immutable Audit Trail

GitHub API provides tamper-proof evidence of all enforcement actions.

##### 5. Formalized Exceptions

Bypass controls with approval workflows, time-boxing, and automatic re-enablement.

---

#### Related Patterns

- **[Required Status Checks](../status-checks/index.md)** - CI/CD as merge gates
- **[Commit Signing](../commit-signing/commit-signing.md)** - Cryptographic proof of authorship
- **[Audit & Compliance](../audit-compliance/audit-evidence.md)** - Evidence collection strategies
- **[GitHub Apps](../../secure/github-apps/index.md)** - Centralized authentication patterns

---

#### Next Steps

Start with **[Security Tiers](security-tiers.md)** to understand the tiered protection model, then review **[Branch Protection Rules](branch-protection.md)** for detailed configuration options.

---

*Protection became immutable. Drift was detected. Remediation was automatic. Auditors found zero gaps. The controls were real.*

## Implementation Roadmap

Phased rollout plan for SDLC hardening.

You can't harden everything at once. Prioritize controls by risk and audit value.

> **Phased Rollout**
>
> Follow the 12-week timeline to avoid disrupting existing workflows. Skip phases at your own risk.
>

Three-month plan from foundation to full enforcement.

---

#### Month 1: Foundation

Goal: Core enforcement in place. Evidence collection begins.

##### Week 1: Branch Protection

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

##### Week 2: CI/CD Status Checks

**Tasks**:

- Create `required-checks.yml` workflow (tests, lint)
- Configure branch protection to require checks
- Test on non-critical repository first

**Workflow**:

```yaml
name: Required Checks
on: [pull_request]
jobs:
  tests:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - run: make test
  lint:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - run: make lint
```

**Validation**: Open PR, verify checks block merge until passing.

---

##### Week 3: GitHub App Setup

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

##### Week 4: Evidence Archive

**Tasks**:

- Set up GCS bucket with lifecycle policy (3 year retention)
- Create monthly evidence collection workflow
- Archive first month's data (branch protection config, merged PRs)

**Workflow**:

```yaml
name: Monthly Evidence
on:
  schedule:
    - cron: '0 0 1 * *'
jobs:
  archive:
    runs-on: ubuntu-latest
    steps:
      - run: gh api repos/org/repo/branches/main/protection > config.json
      - run: gsutil cp *.json gs://audit-evidence/
```

**Validation**: Verify files appear in GCS bucket.

---

#### Month 2: Hardening

Goal: Add secrets detection, commit signing, and SBOM generation.

##### Week 5: Secrets Detection

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

##### Week 6: Signed Commits

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

##### Week 7: SBOM Generation

**Tasks**:

- Add Syft/Trivy to build pipelines
- Generate SBOM for each container build
- Upload SBOMs to artifact storage
- Verify license compliance (no GPL in proprietary code)

**Workflow**:

```yaml
- name: Generate SBOM
  uses: anchore/sbom-action@v0
  with:
    image: app:${{ github.sha }}
    format: cyclonedx-json
    output-file: sbom.json

- name: Upload SBOM
  uses: actions/upload-artifact@v4
  with:
    name: sbom
    path: sbom.json
```

**Validation**: Download artifact, verify SBOM contains expected dependencies.

See [SBOM Generation](../../secure/sbom/sbom-generation.md) for full implementation.

---

##### Week 8: Complete PAT Migration

**Tasks**:

- Audit all remaining PAT usages (`grep -r GITHUB_TOKEN .github/`)
- Create additional GitHub Apps for specific use cases
- Replace all PATs with app tokens
- Revoke old PATs

**Validation**: No PATs referenced in active workflows.

---

#### Month 3: Validation & Policy-as-Code

Goal: Simulate audit, fix gaps, add runtime enforcement.

##### Week 9: Vulnerability Scanning

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

##### Week 10: Policy-as-Code (Kyverno)

**Tasks**:

- Deploy Kyverno to Kubernetes clusters
- Install Policy Reporter for observability
- Implement core policies (resource limits, image sources, labels)
- Configure audit mode first, then enforcement mode

**Core policy**:

```yaml
apiVersion: kyverno.io/v1
kind: ClusterPolicy
metadata:
  name: require-resource-limits
spec:
  validationFailureAction: Enforce
  rules:
    - name: check-limits
      match:
        resources:
          kinds: [Pod]
      validate:
        message: "CPU and memory limits required"
        pattern:
          spec:
            containers:
              - resources:
                  limits:
                    memory: "?*"
                    cpu: "?*"
```

**Validation**: Deploy pod without limits, verify rejection.

See [Policy-as-Code with Kyverno](../../blog/posts/2025-12-13-policy-as-code-kyverno.md) for end-to-end implementation.

---

##### Week 11: Audit Simulation

**Tasks**:

- Pull evidence like an auditor would (API queries for March data)
- Generate summary report (PR reviews, check results, signed commits)
- Identify gaps in evidence or controls
- Document findings and remediation plan

**Simulation script**:

```bash
### Verify branch protection
gh api repos/org/repo/branches/main/protection

### Sample March PRs
gh api 'repos/org/repo/pulls?state=closed&base=main' \
  --jq '.[] | select(.merged_at | startswith("2025-03"))'

### Check signature coverage
./scripts/signature-coverage.sh 2025-03-01 2025-04-01
```

**Validation**: Evidence collection succeeds for sampled period.

---

##### Week 12: Remediation & Runbook

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

#### Next Steps

- **[Execution Guide](execution.md)** - Progress tracking, audit readiness criteria, rollback planning, cost estimation, success metrics

---

*Week 1: Protection enabled. Week 4: Evidence collected. Week 12: Audit simulation passed. Controls enforced. System hardened.*

### Phase 1: Foundation (Weeks 1-4)

Establish local development controls and repository protection. These controls prevent bad code from ever entering git history.

> **Real-World Impact**
>
> A fintech client deployed pre-commit hooks across 200 repositories in 2 weeks. Within 48 hours, the hooks blocked 14 attempted commits containing AWS keys, GCP service account tokens, and database credentials. None entered git history.
>

---

#### Phase Overview

Phase 1 establishes the foundation of SDLC security through two critical control layers:

1. **[Pre-commit Hooks](pre-commit-hooks.md)** - Block bad code locally before git commit
2. **[Branch Protection](branch-protection.md)** - Prevent unauthorized merges at repository level

These controls work together to create defense-in-depth at the source code level.

---

#### Phase Components

##### Pre-commit Hooks

Local enforcement that prevents secrets, policy violations, and code quality issues from entering git history.

**Key Controls**:

- Secrets detection with TruffleHog
- YAML/JSON validation
- Language-specific linting (Go, Python, etc.)
- Custom policy enforcement hooks
- Organization-wide distribution

**[View Pre-commit Hooks Details →](pre-commit-hooks.md)**

---

##### Branch Protection Rules

Repository-level enforcement that makes it impossible to merge without meeting security criteria.

**Key Controls**:

- Required pull request reviews
- Code owner approval requirements
- Required CI status checks
- Administrator enforcement (no bypasses)
- Commit signature requirements
- Force push and deletion prevention

**[View Branch Protection Details →](branch-protection.md)**

---

#### Phase 1 Validation Checklist

Before moving to Phase 2, verify all foundation controls work:

- [ ] Pre-commit hooks block secrets in test commit
- [ ] Pre-commit hooks block invalid YAML/JSON
- [ ] Pre-commit hooks enforce language-specific linting
- [ ] All repositories have `.pre-commit-config.yaml`
- [ ] Branch protection requires at least 1 review
- [ ] Branch protection dismisses stale reviews
- [ ] Branch protection requires code owner approval
- [ ] Branch protection enforces admins (`enforce_admins: true`)
- [ ] Branch protection requires signed commits
- [ ] Branch protection blocks force pushes and deletions
- [ ] Required CI status checks are configured (tests, lint, security)
- [ ] Organization-wide branch protection script runs successfully

---

#### Validation Commands

Test that controls are working:

```bash
### Test pre-commit secrets detection
echo "AWS_KEY=AKIAIOSFODNN7EXAMPLE" > .env
git add .env && git commit -m "test"
### Expected: Commit blocked by TruffleHog

### Verify branch protection admin enforcement
gh api repos/org/repo/branches/main/protection | jq '.enforce_admins.enabled'
### Expected: true

### Count repositories with protection
gh repo list org --limit 1000 --json name --jq '.[].name' | while read repo; do
  gh api repos/org/$repo/branches/main/protection >/dev/null 2>&1 && echo "✅"
done | grep "✅" | wc -l
```

---

#### Next Steps

With Phase 1 complete, you have:

- Pre-commit hooks blocking secrets and lint violations locally
- Branch protection preventing unauthorized merges
- Required reviews and signatures on all commits
- Automated distribution ensuring organization-wide coverage

**[Proceed to Phase 2: Automation →](../phase-2/index.md)**

Phase 2 builds on this foundation by adding CI/CD gates, SBOM generation, vulnerability scanning, and automated evidence collection.

---

#### Related Patterns

- **[Pre-commit Security Gates](../../../pre-commit-hooks/pre-commit-hooks.md)** - Detailed hook configuration
- **[Branch Protection Enforcement](../../../branch-protection/branch-protection.md)** - GitHub API automation
- **[Implementation Roadmap Overview](index.md)** - Complete roadmap
- **[Phase 2: Automation →](../phase-2/index.md)** - CI/CD gates

---

*Pre-commit hooks deployed. Secrets blocked at source. Branch protection enforced. Admins cannot bypass. Foundation is set.*

### Phase 2: Automation (Weeks 5-8)

Automate security, quality, and compliance checks in the pipeline. Tests that fail, code with vulnerabilities, and builds without SBOMs never merge.

> **Real-World Impact**
>
> An e-commerce platform implemented CI gates and SBOM generation in 3 weeks. Within the first month, gates blocked 23 merges with HIGH/CRITICAL vulnerabilities and generated SBOMs for 847 builds. When Log4Shell hit, they had complete dependency visibility across all services in under 2 hours.
>

---

#### Phase Overview

Phase 2 extends enforcement into the CI/CD pipeline through two critical areas:

1. **[CI/CD Gates](ci-gates.md)** - Required checks, SBOM generation, vulnerability scanning, SLSA provenance
2. **[Evidence Collection](evidence-collection.md)** - Automated archival and metrics tracking

These controls ensure failing builds never reach production and provide audit evidence.

---

#### Phase Components

##### CI/CD Gates

Pipeline enforcement that blocks merges with failing tests, vulnerabilities, or missing SBOMs.

**Key Controls**:

- Required status checks workflow
- SBOM generation for every build
- Vulnerability scanning with fail-fast
- SLSA provenance for releases
- Evidence storage integration

**[View CI/CD Gates Details →](ci-gates.md)**

---

##### Evidence Collection

Automated archival of branch protection configs, PR reviews, and build artifacts.

**Key Controls**:

- Branch protection config snapshots
- PR review metadata collection
- Workflow run log archival
- Integration with branch protection
- Metrics tracking

**[View Evidence Collection Details →](evidence-collection.md)**

---

#### Phase 2 Validation Checklist

Before moving to Phase 3, verify all automation controls work:

- [ ] CI workflow runs on every pull request
- [ ] Test failures block merge
- [ ] Lint failures block merge
- [ ] Security scan failures block merge
- [ ] SBOM is generated for every build
- [ ] SBOM is uploaded to evidence storage
- [ ] Vulnerability scanning fails on HIGH/CRITICAL
- [ ] SLSA provenance is generated for releases
- [ ] Provenance can be verified with `slsa-verifier`
- [ ] Monthly evidence collection runs successfully
- [ ] Evidence storage contains expected files

---

#### Validation Commands

Test that controls are working:

```bash
### Test CI blocks failing tests
echo "func TestFail(t *testing.T) { t.Fatal() }" >> main_test.go
git push origin feature-branch
### Expected: Merge blocked by CI failure

### Verify SBOM generation
gsutil ls gs://audit-evidence/sbom/$(date +%Y-%m-%d)/
### Expected: SBOM files for today's builds

### Verify SLSA provenance
gh release view vX.Y.Z --json assets | jq '.assets[].name' | grep intoto
### Expected: .intoto.jsonl file exists

### Verify evidence collection
gsutil ls gs://audit-evidence/2025-01/
### Expected: branch-protection.json, merged-prs.json
```

---

#### Next Steps

With Phase 2 complete, you have:

- CI gates blocking failing tests and security scans
- SBOM generation for every build
- Vulnerability scanning with fail-fast
- SLSA provenance for all releases
- Automated evidence collection

**[Proceed to Phase 3: Runtime →](../phase-3/index.md)**

Phase 3 extends enforcement to runtime by controlling what can actually deploy to production.

---

#### Related Patterns

- **[SLSA Provenance](../../../slsa-provenance/slsa-provenance.md)** - Build attestation details
- **[SBOM Generation](../../../../secure/sbom/sbom-generation.md)** - Software Bill of Materials
- **[Vulnerability Scanning](../../../../secure/vulnerability-scanning/vulnerability-scanning.md)** - Container image scanning
- **[Implementation Roadmap Overview](index.md)** - Complete roadmap
- **[Phase 1: Foundation](../phase-1/index.md)** - Pre-commit and branch protection
- **[Phase 3: Runtime →](../phase-3/index.md)** - Production policy enforcement

---

*CI gates deployed. SBOM generated. Vulnerabilities blocked. SLSA provenance signed. Evidence archived. Supply chain security is enforced, not suggested.*

### Phase 3: Runtime (Weeks 9-12)

Control what runs in production, not just what gets committed. Policy is enforced at admission time before pods can deploy.

> **Real-World Impact**
>
> A SaaS company deployed Kyverno policies to 5 Kubernetes clusters in 1 week. Within 72 hours, the policies blocked 34 pod deployments without resource limits, 18 images from untrusted registries, and 12 containers attempting to run as root. Zero manual intervention required.
>

---

#### Phase Overview

Phase 3 extends enforcement to runtime through three critical areas:

1. **[Policy Enforcement](policy-enforcement.md)** - Core Kyverno policies for resource limits, image verification, security context
2. **[Advanced Policies](advanced-policies.md)** - Namespace quotas, pod security standards, network policies
3. **[Rollout Strategy](rollout.md)** - Audit-first deployment approach and metrics

These controls ensure only compliant workloads run in production.

---

#### Phase Components

##### Policy Enforcement

Core admission control policies that block non-compliant pods.

**Key Controls**:

- Kyverno deployment and configuration
- Required resource limits (CPU/memory)
- Image source verification (approved registries)
- Security context requirements (non-root, read-only)
- Policy Reporter dashboard

**[View Policy Enforcement Details →](policy-enforcement.md)**

---

##### Advanced Policies

Extended runtime controls for comprehensive security.

**Key Controls**:

- Namespace resource quotas
- Pod security standards (baseline/restricted)
- Network policy requirements
- System namespace exclusions

**[View Advanced Policies Details →](advanced-policies.md)**

---

##### Rollout Strategy

Safe deployment approach with audit-first methodology.

**Key Controls**:

- Audit mode monitoring (Week 1)
- Violation remediation (Week 2)
- Enforce mode activation (Week 3)
- Metrics tracking and tuning (Week 4)

**[View Rollout Strategy Details →](rollout.md)**

---

#### Phase 3 Validation Checklist

Before moving to Phase 4, verify all runtime controls work:

- [ ] Kyverno is deployed and webhooks are running
- [ ] Policy Reporter UI is accessible and showing violations
- [ ] Resource limits policy blocks pods without limits
- [ ] Image source policy blocks untrusted registries
- [ ] Security context policy blocks root containers
- [ ] Pod security standards are enforced in production
- [ ] Network policies are required for namespaces
- [ ] Namespace quotas are enforced
- [ ] Policy violations are visible in dashboard
- [ ] All policies have been tested in Audit mode first

---

#### Validation Commands

Test that controls are working:

```bash
### Test pod without limits is rejected
kubectl apply -f - <<EOF
apiVersion: v1
kind: Pod
metadata:
  name: test
spec:
  containers:
    - name: app
      image: nginx
EOF
### Expected: Admission webhook denies request

### Test untrusted registry is blocked
kubectl apply -f - <<EOF
apiVersion: v1
kind: Pod
metadata:
  name: test
spec:
  containers:
    - name: app
      image: docker.io/nginx:latest
EOF
### Expected: Image source validation fails

### Check policy reports
kubectl get policyreport -A
### Expected: Shows pass/fail summary
```

---

#### Next Steps

With Phase 3 complete, you have:

- Kyverno enforcing policies at admission time
- Resource limits on all containers
- Image source verification
- Security context enforcement
- Policy observability dashboard

**[Proceed to Phase 4: Advanced →](../phase-4/index.md)**

Phase 4 completes the implementation with audit evidence collection, compliance validation, and OpenSSF Scorecard monitoring.

---

#### Related Patterns

- **[Policy-as-Code with Kyverno](../../../policy-as-code/kyverno/index.md)** - Detailed policy configuration
- **[Pod Security Standards](../../../../secure/cloud-native/gke-hardening/runtime-security/pod-security-standards.md)** - Security context requirements
- **[Runtime Security](../../../../secure/cloud-native/gke-hardening/runtime-security/index.md)** - Resource limits and runtime controls
- **[Implementation Roadmap Overview](index.md)** - Complete roadmap
- **[Phase 2: Automation](../phase-2/index.md)** - CI/CD gates
- **[Phase 4: Advanced →](../phase-4/index.md)** - Audit evidence and compliance

---

*Kyverno deployed. Policies enforced. Pods without limits blocked. Untrusted images rejected. Root containers denied. Runtime security is enforced, not hoped for.*

### Phase 4: Advanced (Month 4+)

Prove compliance through continuous evidence collection and validation. Auditors get irrefutable proof that controls are real, not cosmetic.

> **Real-World Impact**
>
> A financial services company automated evidence collection for 6 months before their SOC 2 audit. Auditors requested 3 years of historical evidence. The team retrieved 36 months of branch protection configs, PR reviews, commit signatures, and SBOMs in under 10 minutes. Audit completed 2 weeks early.
>

---

#### Phase Overview

Phase 4 completes the implementation with three critical areas:

1. **[Audit Evidence Collection](audit-evidence.md)** - Automated archival of branch protection, PR reviews, signatures, SBOMs
2. **[Compliance Validation](compliance.md)** - OpenSSF Scorecard, Best Practices Badge, SLSA verification, license checks
3. **[Audit Simulation](audit-simulation.md)** - Mock audit timeline, gap analysis, remediation

These controls provide continuous proof of compliance.

---

#### Phase Components

##### Audit Evidence Collection

Automated archival that ensures historical evidence exists when auditors ask.

**Key Controls**:

- Branch protection configuration archive
- Pull request review records
- Commit signature coverage tracking
- Workflow run logs
- SBOM and security scan results

**[View Audit Evidence Details →](audit-evidence.md)**

---

##### Compliance Validation

Third-party verification and automated compliance reporting.

**Key Controls**:

- OpenSSF Scorecard monthly monitoring
- OpenSSF Best Practices Badge
- SLSA provenance verification
- Dependency license compliance
- Monthly compliance report generation

**[View Compliance Validation Details →](compliance.md)**

---

##### Audit Simulation

Mock audit process to identify and fix gaps before real auditors arrive.

**Key Controls**:

- Document request simulation (Week 1)
- PR sampling and validation (Week 2)
- Gap analysis (Week 3)
- Remediation and verification (Week 4)

**[View Audit Simulation Details →](audit-simulation.md)**

---

#### Phase 4 Validation Checklist

Before declaring full implementation complete:

- [ ] Branch protection config archive runs monthly
- [ ] PR review records are collected and stored
- [ ] Commit signature coverage is tracked and reported
- [ ] Workflow run logs are archived monthly
- [ ] SBOM and scan results are archived for every release
- [ ] OpenSSF Scorecard runs monthly
- [ ] Scorecard score is ≥ 8.0/10
- [ ] OpenSSF Best Practices Badge obtained
- [ ] SLSA provenance verified on all releases
- [ ] Dependency license compliance checked
- [ ] Monthly compliance report generated automatically
- [ ] Evidence storage is tamper-proof and versioned

---

#### Validation Commands

Test that controls are working:

```bash
### Verify evidence archive exists
gsutil ls gs://audit-evidence/2025-01/branch-protection.json
### Expected: File exists with branch protection config

### Check OpenSSF Scorecard score
docker run gcr.io/openssf/scorecard-action:stable --repo=github.com/org/repo
### Expected: Score ≥ 8.0/10

### Verify SLSA provenance
gh release view vX.Y.Z --json assets | jq '.assets[].name' | grep intoto
### Expected: .intoto.jsonl file exists

### Test evidence retrieval speed
time gsutil ls gs://audit-evidence/2024-*/branch-protection.json
### Expected: < 10 seconds
```

---

#### Next Steps

With Phase 4 complete, you have:

- Automated monthly evidence collection
- Branch protection and PR review archives
- Commit signature tracking
- SBOM and scan result storage
- OpenSSF Scorecard monitoring
- Compliance report generation
- Audit-ready evidence repository

**Your SDLC is fully hardened and compliance-ready.**

---

#### Related Patterns

- **[Audit Evidence Collection](../../../audit-compliance/audit-evidence.md)** - Evidence storage details
- **[OpenSSF Scorecard](../../../../secure/scorecard/scorecard-compliance.md)** - Scorecard configuration
- **[SLSA Provenance](../../../slsa-provenance/slsa-provenance.md)** - Build attestation
- **[Implementation Roadmap Overview](index.md)** - Complete roadmap
- **[Phase 3: Runtime](../phase-3/index.md)** - Runtime enforcement

---

*Evidence collected. Auditors satisfied. Scorecard 10/10. SLSA provenance verified. Compliance is automatic, not aspirational.*

### SDLC Hardening Implementation Roadmap

Every control in this roadmap is actionable and verifiable. No vague policies. No wishful thinking.

> **Audit Foundation**
>
> These controls are what auditors will verify. Skip items at your own risk. Each control must be fully deployed and evidenced before claiming compliance.
>

---

#### Overview

This implementation roadmap provides a structured approach to hardening your Software Development Lifecycle (SDLC) across four critical phases:

1. **[Phase 1: Foundation](phase-1/index.md)** - Local enforcement and branch protection
2. **[Phase 2: Automation](phase-2/index.md)** - CI/CD gates and policy automation
3. **[Phase 3: Runtime](phase-3/index.md)** - Production policy enforcement
4. **[Phase 4: Advanced](phase-4/index.md)** - Audit evidence and compliance validation

Each phase builds on the previous one, creating defense-in-depth through multiple enforcement layers.

---

#### Roadmap Phases

##### Phase 1: Foundation (Weeks 1-4)

Establish local development controls and repository protection.

**Phase Components**:

- **[Pre-commit Hooks](phase-1/pre-commit-hooks.md)** - Secrets detection, linting, policy enforcement
- **[Branch Protection](phase-1/branch-protection.md)** - Required reviews, status checks, admin enforcement

**Why This Phase Matters**: If secrets enter git history, rotation doesn't help. If admins can bypass reviews, the policy is worthless. Foundation controls prevent bad code from ever entering the system.

**[View Phase 1 Overview →](phase-1/index.md)**

---

##### Phase 2: Automation (Weeks 5-8)

Automate security and quality checks in CI/CD pipelines.

**Phase Components**:

- **[CI/CD Gates](phase-2/ci-gates.md)** - Required checks, SBOM generation, vulnerability scanning, SLSA provenance
- **[Evidence Collection](phase-2/evidence-collection.md)** - Automated archival and metrics tracking

**Why This Phase Matters**: Tests that fail, code with vulnerabilities, and builds without SBOMs never merge. CI becomes a gate, not a log. Supply chain security becomes automatic.

**[View Phase 2 Overview →](phase-2/index.md)**

---

##### Phase 3: Runtime (Weeks 9-12)

Control what runs in production, not just what gets committed.

**Phase Components**:

- **[Policy Enforcement](phase-3/policy-enforcement.md)** - Core Kyverno policies, resource limits, image verification
- **[Advanced Policies](phase-3/advanced-policies.md)** - Namespace quotas, pod security, network policies
- **[Rollout Strategy](phase-3/rollout.md)** - Audit-first deployment and metrics

**Why This Phase Matters**: Pods without resource limits, images from untrusted registries, or missing security context cannot run. Policy is enforced before deployment, not after incidents.

**[View Phase 3 Overview →](phase-3/index.md)**

---

##### Phase 4: Advanced (Month 4+)

Prove compliance through continuous evidence collection and validation.

**Phase Components**:

- **[Audit Evidence](phase-4/audit-evidence.md)** - Branch protection, PR reviews, signatures, SBOMs
- **[Compliance Validation](phase-4/compliance.md)** - OpenSSF Scorecard, SLSA verification, license checks
- **[Audit Simulation](phase-4/audit-simulation.md)** - Mock audit, gap analysis, remediation

**Why This Phase Matters**: Auditors will ask "prove branch protection was enabled on 2025-01-01". Archived config proves it. Evidence collection must be automatic and tamper-proof.

**[View Phase 4 Overview →](phase-4/index.md)**

---

#### Implementation Timeline

| Phase | Timeline | Key Milestone | Validation Method |
|-------|----------|---------------|-------------------|
| **Phase 1: Foundation** | Weeks 1-4 | Branch protection on all repos | Test admin bypass attempt |
| **Phase 2: Automation** | Weeks 5-8 | CI gates block failing tests | Merge attempt with failing test |
| **Phase 3: Runtime** | Weeks 9-12 | Kyverno enforces pod policies | Deploy pod without limits |
| **Phase 4: Advanced** | Month 4+ | OpenSSF Scorecard 10/10 | Automated evidence retrieval |

---

#### Critical Success Factors

> **These are non-negotiable**
>
> - **Branch protection on every release branch** (main, production, release-*)
> - **`enforce_admins: true`** (no admin bypasses)
> - **100% signature coverage** on all repositories
> - **SLSA provenance** on every release
> - **OpenSSF Scorecard 10/10** for critical repositories
> - **Monthly evidence collection** with tamper-proof storage
> - **Audit simulation** before real auditors arrive
>

---

#### Validation Strategy

Each phase includes validation steps that prove controls are working:

**Foundation Phase**:

```bash
### Test pre-commit hook blocks secrets
echo "AKIAIOSFODNN7EXAMPLE" > .env && git add .env && git commit -m "test"
### Expected: Commit blocked by TruffleHog

### Test admin enforcement
gh api repos/org/repo/branches/main/protection | jq '.enforce_admins.enabled'
### Expected: true
```

**Automation Phase**:

```bash
### Test CI blocks failing tests
echo "func TestFail(t *testing.T) { t.Fatal() }" >> main_test.go
git push origin feature-branch
### Expected: Merge blocked by CI failure

### Verify SBOM generation
gsutil ls gs://audit-evidence/sbom/$(date +%Y-%m-%d)/
### Expected: SBOM files for today's builds
```

**Runtime Phase**:

```bash
### Test pod without resource limits is rejected
kubectl apply -f pod-no-limits.yaml
### Expected: Admission webhook denies request

### Test untrusted registry is blocked
kubectl apply -f pod-dockerhub.yaml
### Expected: Image source validation fails
```

**Advanced Phase**:

```bash
### Verify evidence archive exists
gsutil ls gs://audit-evidence/2025-01/branch-protection.json
### Expected: File exists with branch protection config

### Check OpenSSF Scorecard score
docker run gcr.io/openssf/scorecard-action:stable --repo=github.com/org/repo
### Expected: Score ≥ 8.0/10
```

---

#### Prerequisites

Before starting Phase 1, ensure you have:

- [ ] Access to GitHub organization settings
- [ ] Cloud storage bucket for evidence (GCS, S3, Azure Blob)
- [ ] GitHub App or token with appropriate permissions
- [ ] Kubernetes cluster for runtime policies (Phase 3)
- [ ] Team buy-in on enforcement approach

> **Start Small**
>
> Begin with a single repository or team. Validate controls work before scaling organization-wide. Use pilot repository as reference implementation for others.
>

---

#### Common Pitfalls

##### Pitfall 1: Deploying controls without validation

Don't assume a control works because it's deployed. Every control must be tested with an attack scenario.

**Solution**: Use validation commands from each phase to prove controls block violations.

##### Pitfall 2: Admin bypass enabled "just in case"

Setting `enforce_admins: false` makes all other controls optional.

**Solution**: Keep admin enforcement enabled. If emergency bypass is needed, document it, use it, then re-enable immediately.

##### Pitfall 3: Evidence collection runs but isn't verified

Automated evidence collection means nothing if the data is corrupt or incomplete.

**Solution**: Monthly spot-check of evidence archives. Verify JSON is valid, timestamps are correct, and files are complete.

##### Pitfall 4: OpenSSF Scorecard score drops unnoticed

Your score can regress if controls are disabled or practices slip.

**Solution**: Run Scorecard monthly. Alert on score drops > 0.5 points. Investigate immediately.

---

#### Next Steps

1. **Review** [Phase 1: Foundation](phase-1/index.md) and identify repositories for pilot deployment
2. **Prepare** cloud storage bucket for evidence collection
3. **Document** current state (how many repos have branch protection? How many require reviews?)
4. **Schedule** implementation kickoff with engineering teams
5. **Execute** Phase 1 controls on pilot repository
6. **Validate** controls work before scaling organization-wide

---

#### Related Patterns

- **[Execution Guide](../execution.md)** - Progress tracking and rollback planning
- **[Branch Protection Enforcement](../../branch-protection/branch-protection.md)** - GitHub configuration
- **[Pre-commit Security Gates](../../pre-commit-hooks/pre-commit-hooks.md)** - Local enforcement
- **[SLSA Provenance](../../slsa-provenance/slsa-provenance.md)** - Build attestations
- **[Audit Evidence Collection](../../audit-compliance/audit-evidence.md)** - Long-term evidence storage
- **[Policy-as-Code with Kyverno](../../policy-as-code/kyverno/index.md)** - Runtime enforcement

---

*Foundation laid. Controls enforced. Evidence collected. Auditors get irrefutable proof. SDLC hardening is not a checklist item. It's operational reality.*

## Incident Readiness

### Incident Response Playbook Templates

Operational runbooks for Kubernetes security incidents. Each playbook combines decision trees, step-by-step procedures, and validation criteria to enable rapid, confident response to common incident patterns.

This library is designed for teams operating Kubernetes infrastructure at scale, where incident response speed and consistency directly impact security posture and business continuity.

---

#### How to Use This Library

##### Before an Incident

1. **Review** each playbook relevant to your environment and threat model
2. **Customize** commands and thresholds for your cluster configuration
3. **Test** playbook steps in non-production environments
4. **Train** on-call engineers on decision trees and escalation paths
5. **Integrate** with monitoring and alerting systems

##### During an Incident

1. **Identify** which playbook applies using decision trees
2. **Follow** procedures in sequence without skipping steps
3. **Document** actions and timestamps as you proceed
4. **Validate** success criteria before moving to next phase
5. **Escalate** if playbook doesn't resolve issue or if conditions change

##### After an Incident

1. **Collect** evidence using post-incident procedures
2. **Complete** RCA templates to identify root causes
3. **Track** improvements in incident tracking system
4. **Update** playbooks based on lessons learned

---

#### Playbook Categories

> **Practice Exercises**
>
> Tabletop exercises and simulation scenarios. Use these to train teams and validate playbook effectiveness before real incidents.
>
> *Coming soon: Tabletop exercise templates*
>

---

#### Alert Classification Decision Tree

```mermaid
graph TD
    A["Incident Alert Triggered"] --> B{"Severity Level?"}
    B -->|Critical - Cluster unavailable| C["CRITICAL: Execute <br/>Active Threat Response"]
    B -->|High - Service degradation| D["HIGH: Execute <br/>Active Container Threat"]
    B -->|Medium - Anomalous behavior| E["MEDIUM: Execute <br/>Suspicious Activity Assessment"]
    B -->|Low - Policy violation| F["LOW: Execute <br/>Compliance Audit"]

    C --> C1["Page on-call engineer<br/>Declare SEV-1 incident<br/>Start war room"]
    D --> D1["Alert primary on-call<br/>Declare SEV-2 incident<br/>Notify team"]
    E --> E1["Create incident ticket<br/>Assign to engineer<br/>Set 1-hour response SLO"]
    F --> F1["Log for audit<br/>Schedule review<br/>No immediate response"]

    %% Ghostty Hardcore Theme
    style C fill:#dc2626
    style D fill:#ea580c
    style E fill:#eab308
    style F fill:#22c55e

```

---

#### Quick Reference: Incident Severity Levels

| Level | Criteria | Response | Playbook |
|---|---|---|---|
| **SEV-1 (Critical)** | Cluster unavailable, widespread pod failures, data loss risk | Page all on-call, declare war room, 15-min SLO | Detection → Containment → Remediation (parallel) |
| **SEV-2 (High)** | Service degradation, one pod compromised, customer impact | Page primary on-call, 1-hour SLO | Detection → Containment → Remediation (sequential) |
| **SEV-3 (Medium)** | Anomalous behavior, no customer impact, security alert | Create ticket, assign to engineer, 4-hour SLO | Detection → Investigation (no immediate action) |
| **SEV-4 (Low)** | Policy violation, compliance finding, no immediate threat | Log for audit, schedule review, no SLO | Audit only, no immediate action |

---

#### Continuous Improvement

##### Playbook Review Schedule

- **Monthly:** Review alerts that triggered playbooks for false positives
- **Quarterly:** Update playbooks based on lessons learned from incidents
- **Semi-annually:** Review against new threats and attack patterns
- **Annually:** Comprehensive review and rewrite of all playbooks

##### Metrics to Track

- **Time to Detect:** Goal: < 5 minutes from incident start
- **Time to Contain:** Goal: < 15 minutes from detection
- **Time to Resolve:** Goal: < 1 hour from detection
- **Accuracy:** % of playbook steps that applied without modification
- **False Positives:** % of alerts that weren't actual incidents

##### Feedback Loop

After each incident:

1. RCA identifies gaps in playbook
2. Update playbook with lessons learned
3. Add new alert rules for faster detection
4. Update runbook links in monitoring
5. Run training on updated playbook
6. Update metrics and SLOs based on performance

---

#### Additional Resources

- **Kubernetes Security Documentation:** <https://kubernetes.io/docs/concepts/security/>
- **Network Policies:** <https://kubernetes.io/docs/concepts/services-networking/network-policies/>
- **RBAC Authorization:** <https://kubernetes.io/docs/reference/access-authn-authz/rbac/>
- **Audit Logging:** <https://kubernetes.io/docs/tasks/debug-application-cluster/audit/>
- **Pod Security Standards:** <https://kubernetes.io/docs/concepts/security/pod-security-standards/>

---

#### Version History

| Version | Date | Changes |
|---|---|---|
| 1.0 | 2026-01-02 | Initial version with Detection, Containment, Remediation, and Post-Incident playbooks |

---

#### Contact and Support

For playbook updates, questions, or incident support:

- **On-Call:** Page the primary on-call through your alerting system
- **Playbook Issues:** File GitHub issue in the incident-response repository
- **Training:** Contact your security/SRE team lead
- **Escalation:** Use the escalation phone tree in your incident response plan

## Policy-as-Code: End-to-End Enforcement

Enforce security and compliance policies across…

Enforce security and operational policies across the entire SDLC: local development, CI pipelines, and runtime admission control.

#### Overview

Policy-as-Code ensures compliance through automated enforcement at three critical checkpoints:

```mermaid
graph LR
    A[Local Development] -->|Push Code| B[CI Pipeline]
    B -->|Deploy| C[Runtime Admission]

    A -->|policy-platform container| A1[Kyverno CLI]
    B -->|policy-platform container| B1[Kyverno CLI]
    C -->|Kyverno controller| C1[Admission Webhook]

    %% Ghostty Hardcore Theme
    style A fill:#a7e22e,color:#1b1d1e
    style B fill:#a7e22e,color:#1b1d1e
    style C fill:#65d9ef,color:#1b1d1e

```

**The Core Principle**: Same policies, three enforcement points. Zero gaps.

---

#### The Problem with Scattered Enforcement

Traditional approaches fail in predictable ways:

| Approach                      | Problem                   | Result                  |
| -------------------------------- | ---------------------------- | -------------------------- |
| Documentation only            | Nobody reads it           | Violations in production |
| CI-only checks | Local testing incomplete | Broken pipelines |
| Runtime-only admission control | Issues caught too late | Failed deployments |

> **The Gap Problem**
>
> CI checks resource limits, but Kyverno policy doesn't match. Developer tests locally, CI passes, runtime rejects deployment. This gap causes production failures.
>

**Solution**: One policy source, distributed everywhere.

---

#### Architecture

##### Policy Sources

Policies originate from version-controlled repositories:

```mermaid
graph TD
    SP[security-policy repo] -->|OCI container| PP[policy-platform]
    DP[devops-policy repo] -->|OCI container| PP
    KC[Kyverno curated policies] -->|OCI container| PP

    PP -->|Container image| L[Local Dev]
    PP -->|Container image| CI[CI Pipeline]
    PP -->|Helm chart| RT[Runtime]

    %% Ghostty Hardcore Theme
    style SP fill:#f92572,color:#1b1d1e
    style DP fill:#fd971e,color:#1b1d1e
    style KC fill:#9e6ffe,color:#1b1d1e
    style PP fill:#65d9ef,color:#1b1d1e

```

##### Policy Aggregation

The policy-platform container aggregates policies from multiple sources:

**Dockerfile (multi-stage build)**:

```dockerfile
FROM security-policy-repo:main AS security_policy_repo
FROM devops-policy-repo:main AS devops_policy_repo

FROM alpine:3.24
RUN apk add helm kyverno pluto spectral

COPY --from=security_policy_repo /repos/security-policy/ /repos/security-policy/
COPY --from=devops_policy_repo /repos/devops-policy/ /repos/devops-policy/
```

**Result**: Single container with all policies, ready to run anywhere.

---

#### Three-Layer Enforcement

##### Layer 1: Local Development

Developer runs policy checks before commit:

```bash
docker run policy-platform:latest \
  kyverno apply /repos/security-policy/policies.yaml \
  --resource my-deployment.yaml
```

**Benefits**:

- Instant feedback
- No CI wait time
- Same validation as CI

##### Layer 2: CI Pipeline

Automated validation in every pull request:

```yaml
steps:
  - name: Validate Security Policy
    image: policy-platform:latest
    script:
      - kyverno apply security-policy.yaml --resource app.yaml
      - kyverno apply devops-policy.yaml --resource app.yaml
```

**Benefits**:

- Blocks non-compliant merges
- Generates policy reports
- Environment-specific validation

##### Layer 3: Runtime Admission

Kyverno admission controller in Kubernetes:

```yaml
apiVersion: kyverno.io/v1
kind: ClusterPolicy
metadata:
  name: require-resource-limits
spec:
  validationFailureAction: Enforce
```

**Benefits**:

- Final safety net
- Prevents misconfigured deployments
- Continuous compliance monitoring

---

#### Enforcement Guarantees

| Stage  | Enforcement         | Bypassable?           | Purpose                       |
| --------- | ---------------------- | ------------------------ | -------------------------------- |
| Local | Developer-initiated | Yes (developer choice) | Fast feedback, early detection |
| CI     | Automated on PR     | No (blocks merge)     | Gate for code review          |
| Runtime | Admission webhook   | No (rejects pod)      | Production safety             |

**Key Insight**: Local and CI use **same container**, runtime uses **same policies**.

---

#### What You'll Learn

This section covers complete policy-as-code implementation:

1. **[Local Development](local-development/index.md)** - Running policies in containers locally
2. **[CI Integration](ci-integration/index.md)** - Automated validation in pipelines
3. **[Runtime Deployment](runtime-deployment/index.md)** - Kyverno admission control
4. **[Multi-Source Policies](multi-source-policies/index.md)** - Aggregating policy repositories
5. **[Policy Packaging](policy-packaging/index.md)** - Building the policy-platform container
6. **[Operations](operations/index.md)** - Day-to-day policy management

---

#### Prerequisites

- Kubernetes cluster (for runtime deployment)
- Container runtime (Docker/Podman for local dev)
- CI platform (GitHub Actions, Bitbucket Pipelines, GitLab CI)
- Basic Kyverno knowledge (see [Kyverno guide](kyverno/index.md))

---

#### Quick Start

> **Start Local, Scale Up**
>
> Test policies locally first. Fix violations in seconds, not hours. Only after local validation works should you move to CI integration and runtime deployment.
>

**Step 1**: Run policies locally

```bash
docker run policy-platform:latest \
  kyverno apply /repos/security-policy/ \
  --resource deployment.yaml
```

**Step 2**: Add to CI pipeline

```yaml
- name: Policy Check
  image: policy-platform:latest
  script:
    - kyverno apply /repos/security-policy/ --resource app.yaml
```

**Step 3**: Deploy Kyverno to cluster

```bash
helm install kyverno kyverno/kyverno -f kyverno-values.yaml
helm install policy-reporter policy-reporter/policy-reporter
```

---

#### Real-World Impact

**Before Policy-as-Code**:

- Pods deployed without resource limits → OOMKilled nodes
- Secrets in ConfigMaps → Security incidents
- Deprecated APIs → Failed upgrades

**After Policy-as-Code**:

- 100% of deployments have resource limits
- Zero secrets in clear text
- Deprecated API usage blocked before merge

**Key Metric**: Issues caught in **local dev** (5 min fix) vs **production** (incident response).

---

#### Architecture Principles

##### 1. Single Source of Truth

Policies live in Git repositories. Everything derives from there.

##### 2. Container-Based Distribution

One container runs everywhere. No "works on my machine."

##### 3. Progressive Enforcement

Local (warn) → CI (fail) → Runtime (block).

##### 4. Separation of Concerns

- **Policy repos**: Define rules
- **Policy-platform**: Package and distribute
- **Kyverno**: Enforce at runtime

---

#### Related Patterns

- **[SDLC Hardening](../index.md)** - Broader enforcement strategies
- **[Kyverno Implementation](kyverno/index.md)** - Runtime policy details
- **[Pre-commit Hooks](../pre-commit-hooks/pre-commit-hooks.md)** - Complementary local checks
- **[CI/CD Patterns](../../patterns/architecture/index.md)** - Pipeline architecture

---

#### Next Steps

Start with **[Local Development](local-development/index.md)** to run policies on your machine, then progress to CI and runtime deployment.

### CI Integration: Automated Policy Enforcement

Block non-compliant code at merge time. Same container as local dev, zero configuration drift.

#### Overview

CI integration enforces policies automatically in every pull request using the **same policy-platform container** developers run locally.

```mermaid
graph LR
    PR[Pull Request] --> ENV[Detect Environment]
    ENV --> LINT[Lint Values]
    LINT --> BUILD[Build Manifests]
    BUILD --> VAL[Validate Policies]
    VAL --> MERGE{All Pass?}
    MERGE -->|Yes| ALLOW[Allow Merge]
    MERGE -->|No| BLOCK[Block Merge]

    %% Ghostty Hardcore Theme
    style ALLOW fill:#a7e22e,color:#1b1d1e
    style BLOCK fill:#f92572,color:#1b1d1e

```

**Key Principle**: CI uses identical validation to local development. No surprises.

---

#### Pipeline Architecture

##### Environment Detection

Policy validation is environment-specific. Detect target environment from branch:

```yaml
- step:
    name: Detect Environment
    script:
      - |
        if [ -n "$BITBUCKET_PR_ID" ]; then
##          # Pull Request - check destination branch
          case $BITBUCKET_PR_DESTINATION_BRANCH in
            "development") ENVIRONMENT="dev" ;;
            "qac")         ENVIRONMENT="qac" ;;
            "staging")     ENVIRONMENT="stg" ;;
            "production")  ENVIRONMENT="prd" ;;
            *)
              echo "Unknown destination branch"
              exit 0
              ;;
          esac
        else
##          # Direct push - check current branch
          case $BITBUCKET_BRANCH in
            "development") ENVIRONMENT="dev" ;;
            "qac")         ENVIRONMENT="qac" ;;
            "staging")     ENVIRONMENT="stg" ;;
            "production")  ENVIRONMENT="prd" ;;
          esac
        fi
        echo "export ENVIRONMENT=${ENVIRONMENT}" > environment.sh
    artifacts:
      - environment.sh
```

> **Environment Detection is Critical**
>
> Production policies are stricter than dev. Wrong environment detection means applying dev policies to production code. This creates security gaps.
>

---

#### Pipeline Stages

##### Stage 1: Schema Validation

Validate Helm values against schemas **before** rendering manifests:

```yaml
- step:
    name: Lint Values Against Schema
    image: policy-platform:latest
    script:
      - source environment.sh
      - |
##        # Merge base values + environment values
        yq eval-all 'select(fileIndex == 0) * select(fileIndex == 1)' \
          /repos/backend-applications/charts/backend-app/values.yaml \
          ./cd/values.${ENVIRONMENT}.yaml \
        > combined_values.yaml

      - |
##        # Validate merged values against schema
        spectral lint \
          -r /repos/backend-applications/.spectral.yaml \
          combined_values.yaml
```

**Catches**:

- Missing required fields
- Type mismatches
- Invalid enum values

**Fails fast**: no point rendering manifests if values are invalid.

---

##### Stage 2: Manifest Rendering

Render environment-specific manifests:

```yaml
- step:
    name: Build Environment Manifests
    image: policy-platform:latest
    script:
      - source environment.sh
      - |
##        # Define chart paths
        sec_pol_chart=/repos/security-policy/charts/security-policy
        dev_pol_chart=/repos/devops-policy/charts/devops-policy
        be_apps_chart=/repos/backend-applications/charts/backend-applications

      - |
##        # Render DevOps policies
        helm template devops-policy ${dev_pol_chart} \
          -f ${dev_pol_chart}/values.yaml \
          -f /repos/devops-policy/cd/values.yaml \
          -f /repos/devops-policy/cd/${ENVIRONMENT}/values.yaml \
        > devops-policy.yaml

      - |
##        # Render Security policies
        helm template security-policy ${sec_pol_chart} \
          -f ${sec_pol_chart}/values.yaml \
          -f /repos/security-policy/cd/values.yaml \
          -f /repos/security-policy/cd/${ENVIRONMENT}/values.yaml \
        > security-policy.yaml

      - |
##        # Render application manifests
        helm template backend-app ${be_apps_chart} \
          -f ${be_apps_chart}/values.yaml \
          -f ./cd/values.${ENVIRONMENT}.yaml \
        > backend-app.yaml

    artifacts:
      - devops-policy.yaml
      - security-policy.yaml
      - backend-app.yaml
```

**Three artifact types**:

1. **security-policy.yaml** for Security rules (resource limits, image policies, etc.)
2. **devops-policy.yaml** for Operational rules (labels, annotations, naming)
3. **backend-app.yaml** for Application manifests to validate

---

##### Stage 3: Policy Validation (Parallel)

Validate against DevOps and Security policies **in parallel**:

```yaml
- parallel:
    steps:
      - step:
          name: Validate DevOps Policy
          image: policy-platform:latest
          script:
            - |
##              # Generate policy report
              kyverno apply devops-policy.yaml \
                --resource backend-app.yaml \
                --output mutated-resources \
                --policy-report \
                --audit-warn \
              > tmp-policy-report.yaml

            - |
##              # Extract YAML report for download
              sed -n '/^POLICY REPORT:/,$p' tmp-policy-report.yaml \
                | tail -n +3 \
                | { echo '---'; cat; } \
              > policy-report.yaml

            - |
##              # Display summary
              kyverno apply devops-policy.yaml \
                --resource backend-app.yaml \
                --output mutated-resources \
                --remove-color

            - |
##              # Display detailed results table
              kyverno apply devops-policy.yaml \
                --resource backend-app.yaml \
                --output mutated-resources \
                -t --detailed-results \
                --remove-color
          artifacts:
            - policy-report.yaml

      - step:
          name: Validate Security Policy
          image: policy-platform:latest
          script:
            - |
              kyverno apply security-policy.yaml \
                --resource backend-app.yaml \
                --output mutated-resources \
                --policy-report \
                --audit-warn \
              > tmp-policy-report.yaml

            - sed -n '/^POLICY REPORT:/,$p' tmp-policy-report.yaml \
                | tail -n +3 \
                | { echo '---'; cat; } \
              > policy-report.yaml

            - kyverno apply security-policy.yaml \
                --resource backend-app.yaml \
                --output mutated-resources \
                --remove-color

            - kyverno apply security-policy.yaml \
                --resource backend-app.yaml \
                --output mutated-resources \
                -t --detailed-results \
                --remove-color
          artifacts:
            - policy-report.yaml
```

> **Parallel Validation Saves Time**
>
> DevOps and Security policies are independent. Running them in parallel cuts pipeline time in half. Both must pass for merge approval.
>

**Artifacts**: Policy reports downloadable for detailed review.

---

#### Complete Bitbucket Pipeline

Full pipeline showing all stages:

```yaml
image:
  name: policy-platform:main
  username: _json_key
  password: "$GCLOUD_API_KEYFILE"

pipelines:
  pull-requests:
    '**':
##      # Stage 1: Detect environment from PR destination
      - step:
          name: Detect Environment
          script:
            - |
              case $BITBUCKET_PR_DESTINATION_BRANCH in
                "development") ENVIRONMENT="dev" ;;
                "qac")         ENVIRONMENT="qac" ;;
                "staging")     ENVIRONMENT="stg" ;;
                "production")  ENVIRONMENT="prd" ;;
                *)
                  echo "Unknown branch. Skipping."
                  exit 0
                  ;;
              esac
              echo "export ENVIRONMENT=${ENVIRONMENT}" > environment.sh
          artifacts:
            - environment.sh

##      # Stage 2: Validate Helm values schema
      - step:
          name: Lint Values Schema
          script:
            - source environment.sh
            - yq eval-all 'select(fileIndex == 0) * select(fileIndex == 1)' \
                /repos/backend-applications/charts/backend-app/values.yaml \
                ./cd/values.${ENVIRONMENT}.yaml \
              > combined_values.yaml
            - spectral lint -r /repos/backend-applications/.spectral.yaml \
                combined_values.yaml

##      # Stage 3: Render manifests
      - step:
          name: Build Manifests
          script:
            - source environment.sh
            - helm template devops-policy \
                /repos/devops-policy/charts/devops-policy \
                -f /repos/devops-policy/charts/devops-policy/values.yaml \
                -f /repos/devops-policy/cd/${ENVIRONMENT}/values.yaml \
              > devops-policy.yaml
            - helm template security-policy \
                /repos/security-policy/charts/security-policy \
                -f /repos/security-policy/charts/security-policy/values.yaml \
                -f /repos/security-policy/cd/${ENVIRONMENT}/values.yaml \
              > security-policy.yaml
            - helm template backend-app \
                /repos/backend-applications/charts/backend-app \
                -f /repos/backend-applications/charts/backend-app/values.yaml \
                -f ./cd/values.${ENVIRONMENT}.yaml \
              > backend-app.yaml
          artifacts:
            - devops-policy.yaml
            - security-policy.yaml
            - backend-app.yaml

##      # Stage 4: Validate policies (parallel)
      - parallel:
          steps:
            - step:
                name: DevOps Policy
                script:
                  - kyverno apply devops-policy.yaml \
                      --resource backend-app.yaml \
                      --policy-report --audit-warn \
                    > tmp-report.yaml
                  - kyverno apply devops-policy.yaml \
                      --resource backend-app.yaml \
                      --remove-color
                  - kyverno apply devops-policy.yaml \
                      --resource backend-app.yaml \
                      -t --detailed-results --remove-color
                artifacts:
                  - tmp-report.yaml

            - step:
                name: Security Policy
                script:
                  - kyverno apply security-policy.yaml \
                      --resource backend-app.yaml \
                      --policy-report --audit-warn \
                    > tmp-report.yaml
                  - kyverno apply security-policy.yaml \
                      --resource backend-app.yaml \
                      --remove-color
                  - kyverno apply security-policy.yaml \
                      --resource backend-app.yaml \
                      -t --detailed-results --remove-color
                artifacts:
                  - tmp-report.yaml
```

> **Policy Report Artifacts**
>
> Each validation step generates a `policy-report.yaml` artifact. Download these for detailed offline review and compliance tracking.
>

---

#### Next Steps

- **[GitHub Actions Integration](github-actions.md)** for GitHub Actions workflow examples
- **[Runtime Deployment](../runtime-deployment/index.md)** for Deploy Kyverno admission control
- **[Multi-Source Policies](../multi-source-policies/index.md)** for Aggregate multiple policy repos
- **[Operations](../operations/index.md)** for Day-to-day policy management

### JMESPath for Kyverno

Master JMESPath to unlock advanced Kyverno policy capabilities. Query nested JSON structures, build complex validation logic, and enforce standards that simple pattern matching cannot express.

> **What You'll Learn**
>
> JMESPath extends Kyverno beyond basic pattern matching. Use it for cross-field validation, dynamic conditions, array filtering, and string transformations. Essential for enterprise-grade policy enforcement.
>

---

#### Why JMESPath Matters

**Simple pattern matching fails when you need to:**

- Compare multiple fields (requests vs limits)
- Validate conditionally (if label exists, require annotation)
- Parse and transform data (extract image tags)
- Filter arrays dynamically (containers with specific images)
- Build complex boolean logic across nested structures

**JMESPath solves all of this.** It's a query language for JSON, purpose-built for navigating Kubernetes resources and extracting validation data.

---

#### Documentation Structure

##### Getting Started

**[JMESPath Patterns](patterns.md)**
Core patterns for common Kyverno use cases. Start here if you're new to JMESPath in policies.

- Projection and filtering
- Cross-field validation
- Array operations
- String transformations
- Boolean logic

##### Advanced Techniques

**[Advanced Patterns](advanced.md)**
Sophisticated validation logic for complex scenarios.

- Label and annotation transformations
- Multi-level array operations
- Dynamic naming validation
- Conditional enforcement based on metadata

**[Enterprise Supply Chain](enterprise-supply-chain.md)**
Supply chain security patterns for production workloads.

- Image signature validation
- Provenance verification
- SBOM enforcement
- Attestation checks

**[Enterprise Patterns](enterprise.md)**
Production-grade policies for enterprise Kubernetes.

- Multi-cluster enforcement
- Compliance validation
- Resource governance
- Security hardening

##### Reference Material

**[Function Reference](reference.md)**
Complete JMESPath function library for Kyverno.

- String functions (split, join, contains)
- Array operations (map, filter, sort)
- Comparison operators
- Logical expressions

**[Testing Guide](testing.md)**
Test JMESPath expressions before deploying policies.

- `kyverno jp` CLI usage
- Test case development
- Debugging patterns
- Common mistakes and fixes

---

#### Quick Start

**Install Kyverno CLI for testing:**

```bash
### Install kyverno CLI
brew install kyverno/kyverno/kyverno

### Test JMESPath expression
kyverno jp query -i manifest.yaml 'spec.template.spec.containers[*].name'
```

**Simple validation example:**

```yaml
apiVersion: kyverno.io/v1
kind: ClusterPolicy
metadata:
  name: require-resource-limits
spec:
  validationFailureAction: Enforce
  rules:
  - name: validate-limits
    match:
      any:
      - resources:
          kinds:
          - Pod
    validate:
      message: "All containers must define resource limits"
      deny:
        conditions:
          any:
          - key: "{{ request.object.spec.containers[?!resources.limits.memory].name | length(@) }}"
            operator: GreaterThan
            value: 0
```

**What this does:**

- Filters containers without memory limits: `containers[?!resources.limits.memory]`
- Extracts their names: `.name`
- Counts them: `| length(@)`
- Denies if count > 0

---

#### When to Use JMESPath

**Use JMESPath when:**

- Pattern matching can't express your logic
- You need conditionals or transformations
- Validation depends on multiple fields
- You're filtering or comparing arrays

**Skip JMESPath when:**

- Simple pattern matching works (`pattern`, `anyPattern`)
- You're only checking field existence
- No cross-field validation needed

> **Test Before Deploying**
>
> Always test JMESPath expressions with `kyverno jp` before adding them to policies. Syntax errors fail silently in audit mode and block resources in enforce mode.
>

---

#### Learning Path

**Beginner:**

1. Read [JMESPath Patterns](patterns.md) - core techniques
2. Use [Testing Guide](testing.md) - validate your expressions
3. Reference [Function Reference](reference.md) - lookup syntax

**Intermediate:**

1. Study [Advanced Patterns](advanced.md) - complex scenarios
2. Apply [Enterprise Patterns](enterprise.md) - production use cases

**Advanced:**

1. Implement [Enterprise Supply Chain](enterprise-supply-chain.md) - security hardening
2. Build custom patterns for your environment

---

#### External Resources

- [JMESPath Official Tutorial](https://jmespath.org/tutorial.html) - language fundamentals
- [Kyverno JMESPath Documentation](https://kyverno.io/docs/writing-policies/jmespath/) - policy-specific usage
- [JMESPath Playground](https://jmespath.org/) - interactive testing

---

**Next:** Start with [JMESPath Patterns](patterns.md) to learn core techniques.

### Kyverno Basics

Kyverno runs as a dynamic admission controller in Kubernetes. It validates, mutates, and generates resources based on policies written in YAML.

---

#### Installation

Install Kyverno using Helm:

```bash
### Add Kyverno Helm repository
helm repo add kyverno https://kyverno.github.io/kyverno/
helm repo update

### Install Kyverno
helm install kyverno kyverno/kyverno \
  --namespace kyverno \
  --create-namespace \
  --set replicaCount=3

### Verify installation
kubectl get pods -n kyverno
```

Kyverno creates webhook configurations that intercept resource creation/updates before they reach etcd.

---

#### Basic Kyverno Policy

> **Quick Start**
>
> This guide is part of a modular documentation set. Refer to related guides in the navigation for complete context.
>

Require resource limits on all deployments:

```yaml
apiVersion: kyverno.io/v1
kind: ClusterPolicy
metadata:
  name: require-resource-limits
spec:
  validationFailureAction: Enforce
  background: true
  rules:
    - name: check-resource-limits
      match:
        any:
          - resources:
              kinds:
                - Deployment
      validate:
        message: "Resource limits are required for all containers"
        pattern:
          spec:
            template:
              spec:
                containers:
                  - resources:
                      limits:
                        memory: "?*"
                        cpu: "?*"
```

Try to deploy without limits:

```bash
$ kubectl apply -f deployment.yaml
Error from server: admission webhook "validate.kyverno.svc-fail" denied the request:

policy Deployment/default/api for resource violation:

require-resource-limits:
  check-resource-limits: validation error: Resource limits are required for all containers
```

Deployment blocked. Policy enforced.

---

#### Audit Mode vs Enforce Mode

Roll out policies in audit mode first:

```yaml
spec:
  validationFailureAction: Audit  # Log violations, don't block
```

Check logs for violations:

```bash
kubectl get policyreport -A

NAMESPACE   NAME                          PASS   FAIL   WARN   ERROR   SKIP
default     polr-ns-default              12     3      0      0       0
production  polr-ns-production           45     1      0      0       0
```

Fix violations. Then switch to Enforce:

```yaml
spec:
  validationFailureAction: Enforce  # Block violations
```

##### Gradual Rollout Strategy

1. Deploy policy in `Audit` mode
2. Monitor PolicyReports for 1 week
3. Remediate failures
4. Switch to `Enforce` mode
5. Handle exceptions with exclusions

Don't deploy straight to Enforce. Discover violations first.

---

#### Policy Structure

All Kyverno policies follow this structure:

```yaml
apiVersion: kyverno.io/v1
kind: ClusterPolicy  # or Policy for namespaced
metadata:
  name: policy-name
spec:
  validationFailureAction: Enforce | Audit
  background: true | false  # Apply to existing resources
  rules:
    - name: rule-name
      match:  # What resources to check
        any:
          - resources:
              kinds: [Deployment, StatefulSet]
              namespaces: [production, staging]
      exclude:  # What to skip
        any:
          - resources:
              namespaces: [kube-system]
      validate | mutate | generate:  # What to do
##        # Policy logic here
```

---

#### Related Guides

- **[Policy Patterns](policy-patterns.md)** - Common validation and mutation patterns
- **[Testing and Exceptions](testing-approaches.md)** - Test policies before production
- **[CI/CD Integration](ci-cd-integration.md)** - Automate policy validation

---

*Policy deployed in audit mode. Violations logged. Teams notified. Fixes deployed. Policy switched to enforce. Zero production impact.*

### Kyverno Generation Templates

Kyverno generation policies automatically create supporting resources when specific conditions are met. Instead of relying on manual configuration or documentation, generation policies enforce security and resilience by default.

> **Manual Configuration Creates Gaps**
>
> Relying on documentation to ensure every namespace has ResourceQuotas or every HA workload has a PodDisruptionBudget creates inconsistencies. Generation policies enforce these requirements automatically at creation time.
>

#### What is Generation?

Generation policies create new Kubernetes resources in response to triggers. When a namespace is created, a generation policy can automatically add ResourceQuotas, NetworkPolicies, and LimitRanges. When a high-availability Deployment appears, a generation policy can create a PodDisruptionBudget.

This eliminates the gap between "we should have this" and "we do have this."

---

#### Key Concepts

##### Triggers

Generation policies react to resource creation or changes:

- **Namespace creation** → Generate default-deny NetworkPolicies and ResourceQuotas
- **Deployment with 2+ replicas** → Generate PodDisruptionBudget
- **Production workload** → Generate stricter quotas and policies
- **Label changes** → Synchronize generated resources with new requirements

##### Synchronization

When `synchronize: true` is set, Kyverno keeps generated resources in sync with the source. If you change a namespace label, the generated ResourceQuota updates to match. If you scale a Deployment to 1 replica, the PodDisruptionBudget is removed.

This enforces consistency without manual intervention.

##### Exclusions and Preconditions

Generation policies use exclusions to skip system namespaces and preconditions to enforce requirements:

```yaml
exclude:
  resources:
    names:
      - kube-system
      - kube-public
      - kube-node-lease

preconditions:
  all:
    - key: "{{ request.object.spec.replicas }}"
      operator: GreaterThanOrEquals
      value: 2
```

Only workloads meeting specific criteria get generated resources.

---

#### Template Categories

##### Namespace Resources

Automatically create security and resource governance for new namespaces:

- **[ResourceQuotas and NetworkPolicies](namespace.md)** - Default quotas and default-deny networking for every new namespace

**Use cases:**

- Prevent resource exhaustion from unconstrained namespaces
- Enforce zero-trust networking by default
- Automatically apply environment-specific quotas (dev vs production)
- Ensure DNS egress is allowed while blocking other traffic

---

##### Workload Resources

Automatically create resilience and availability controls for high-availability workloads:

- **[PodDisruptionBudgets](workload.md)** - Automatic PDBs for Deployments and StatefulSets with multiple replicas

**Use cases:**

- Prevent downtime during cluster upgrades and node maintenance
- Enforce SLA compliance for critical services
- Protect against mass pod evictions during autoscaling
- Maintain service availability during partial failures

---

#### When to Use Generation

**Use generation when:**

- You want to enforce security-by-default for all new resources
- Manual configuration creates gaps and inconsistencies
- You need automatic synchronization with changing requirements
- Supporting resources should follow workload lifecycle (create, update, delete)

**Do not use generation when:**

- Resources require unique, per-workload customization
- Generated resources would conflict with existing resources
- You need human approval before resource creation
- The triggering resource does not contain enough context to generate correctly

---

#### Validation Strategy

After deploying generation policies, validate that resources are created correctly:

```bash
### Check that new namespaces get ResourceQuotas
kubectl create namespace test-gen
kubectl get resourcequotas -n test-gen

### Check that multi-replica Deployments get PDBs
kubectl create deployment nginx --image=nginx --replicas=3 -n test-gen
kubectl label deployment nginx app=nginx -n test-gen
kubectl get pdb -n test-gen

### Audit resources without expected generated objects
kubectl get namespaces -o json | jq -r '.items[] | select(.metadata.name != "kube-system") | .metadata.name' | while read ns; do
  quota_count=$(kubectl get resourcequotas -n $ns --no-headers 2>/dev/null | wc -l)
  if [ $quota_count -eq 0 ]; then
    echo "WARNING: Namespace $ns has no ResourceQuota"
  fi
done
```

---

#### Related Resources

- **[Kyverno Templates Overview](../index.md)** - Back to Kyverno templates
- **[Template Library Overview](../index.md)** - Back to main template library

### Kyverno Image Validation Templates

Enforce container image security controls before deployment. These policies validate image sources, require cryptographic signatures, enforce digest-based references, and block images with critical vulnerabilities.

---

#### Purpose

Image validation is your first line of defense against supply chain attacks. These templates ensure:

- Only images from approved registries can deploy
- Images use immutable digest references, not mutable tags
- Images are cryptographically signed by trusted build pipelines
- Images without recent vulnerability scans are blocked

---

#### Template Categories

##### [Image Validation](validation.md)

Registry allowlists, digest requirements, and tag validation. Block untrusted registries and prohibit mutable `latest` tags. Enforce SHA256 digest references for immutable deployments and supply chain transparency.

Use this for:

- Preventing public Docker Hub images in production
- Requiring digest-based image references
- Blocking images without explicit version tags
- Enforcing approved internal registries

---

##### [Image Signing](signing.md)

Cosign signature verification for trusted images. Verify container images are cryptographically signed before deployment. Support both key-based and keyless (OIDC) signing with SLSA provenance attestations.

Use this for:

- Verifying images come from approved build pipelines
- SLSA compliance and provenance attestations
- Preventing unauthorized image pushes
- Enforcing GitHub Actions workflow signatures

---

##### [Base Image Security](security.md)

Approved base image enforcement. Restrict workloads to approved, maintained base images. Block deprecated distributions and enforce minimal base images for high-security workloads.

Use this for:

- Centralizing base image management
- Blocking vulnerable or EOL distributions
- Enforcing distroless or minimal images
- Standardizing across teams

---

##### [CVE Scanning Gates](cve-scanning.md)

Vulnerability scan attestations and CVE thresholds. Require Trivy vulnerability scan attestations before deployment. Block images with critical or high severity CVEs based on environment risk tolerance.

Use this for:

- Zero-day vulnerability protection
- PCI-DSS and SOC2 compliance
- Shift-left security in CI/CD
- Different CVE thresholds per environment

---

#### Implementation Strategy

> **Phased Rollout Recommended**
>
> Start with registry controls and digest requirements before adding signature verification and CVE scanning. This minimizes disruption while building security layers progressively.
>

##### 1. Start with Registry Allowlists

Block untrusted registries before enforcing signatures or scans.

```bash
kubectl apply -f registry-allowlist-policy.yaml  # Registry controls first
kubectl get clusterpolicy -w   # Watch for Ready status
```

##### 2. Add Digest Requirements

Enforce immutable image references.

```bash
kubectl apply -f digest-enforcement-policy.yaml  # Digest enforcement
kubectl get clusterpolicy -w
```

##### 3. Implement Image Signing

Verify images come from trusted sources.

```bash
kubectl apply -f signature-verification-policy.yaml  # Signature verification
kubectl get clusterpolicy -w
```

##### 4. Enforce CVE Scanning

Block vulnerable images based on scan attestations.

```bash
kubectl apply -f cve-scanning-policy.yaml  # CVE gates
kubectl get clusterpolicy -w
```

##### 5. Centralize Base Images

Standardize on approved, maintained base images.

```bash
kubectl apply -f base-image-policy.yaml  # Base image enforcement
kubectl get clusterpolicy -w
```

---

#### Related Resources

- **[Kyverno Labels →](../labels.md)** - Mandatory metadata enforcement
- **[Kyverno Pod Security →](../pod-security/standards.md)** - Security contexts and capabilities
- **[Kyverno Resource Limits →](../resource/limits.md)** - CPU and memory enforcement
- **[Template Library Overview →](../index.md)** - Back to main page

---

#### External Documentation

- **[Kyverno Image Verification](https://kyverno.io/docs/writing-policies/verify-images/)** - Official Kyverno image verification guide
- **[Sigstore Cosign](https://docs.sigstore.dev/cosign/overview/)** - Container image signing and verification
- **[SLSA Framework](https://slsa.dev/)** - Supply chain security levels
- **[Trivy Scanner](https://aquasecurity.github.io/trivy/)** - Vulnerability scanning tool

### Kyverno Mutation Templates

Mutation policies modify resources at admission time, before they're persisted to etcd. This approach enforces standards without blocking deployments or requiring manual manifest updates.

#### Why Mutation Over Validation

Validation blocks non-compliant resources. Mutation fixes them automatically.

> **Fix vs Block**
>
> Mutations reduce friction by auto-correcting resources at admission time. This approach enforces standards without breaking deployments or requiring manual updates.
>

**Use mutations when:**

- Adding required labels or annotations to all workloads
- Injecting sidecars for logging, monitoring, or security
- Setting default resource limits or security contexts
- Enforcing organizational standards that shouldn't block deployments

**Use validation when:**

- Security boundaries must never be crossed (privileged containers, host paths)
- Resource constraints are non-negotiable (quotas, limits)
- Audit requirements demand explicit opt-in (PII handling, compliance tags)

#### Available Templates

##### [Label Injection](labels.md)

Auto-inject required labels and annotations into workloads:

- Default organizational labels (team, environment, cost center)
- Conditional label injection based on namespace or existing labels
- Label propagation from namespaces to pods

**Apply a policy:**

```bash
kubectl apply -f labels.yaml
```

##### [Sidecar Injection](sidecar.md)

Auto-inject sidecar containers for observability and security:

- Logging sidecars (Fluent Bit, Fluentd)
- Monitoring agents (Prometheus exporters, metrics collectors)
- Security sidecars (secret managers, policy enforcers)

**Apply a policy:**

```bash
kubectl apply -f sidecar.yaml
```

#### Mutation Execution Order

Kyverno processes mutations in this order:

1. **Mutating Admission Webhooks** - Kyverno intercepts CREATE/UPDATE requests
2. **Policy Evaluation** - Matches resource against mutation rules
3. **Mutation Application** - Modifies resource in memory
4. **Validation** - Resource proceeds to validation policies (if any)
5. **Persistence** - Modified resource written to etcd

**Critical:** Mutations only apply to CREATE/UPDATE operations. Set `background: false` to prevent mutations from applying to existing resources during policy sync.

#### Testing Mutations

Use `kubectl apply --dry-run=server` to test mutations without creating resources:

```bash
### Test label mutation
kubectl apply --dry-run=server -f test-deployment.yaml -o yaml | grep -A5 labels

### Test sidecar injection
kubectl apply --dry-run=server -f test-pod.yaml -o yaml | grep -A10 containers
```

#### Common Patterns

##### Conditional Mutations

Only mutate resources that match specific criteria:

- Namespace-scoped mutations (dev vs prod)
- Label-based mutations (inject monitoring only for `app.kubernetes.io/monitored=true`)
- Resource type mutations (different rules for Deployments vs StatefulSets)

##### Mutation Conflicts

When multiple policies mutate the same field:

- **Last-write-wins** - Policies execute in alphabetical order by name
- **Merge strategies** - Use `patchStrategicMerge` or `patchesJson6902` for predictable merging
- **Exclusions** - Use `exclude` blocks to prevent conflicting mutations

##### Security Boundaries

Never mutate security-critical fields:

- Security contexts (runAsUser, capabilities, privileged)
- Resource limits (mutations can escalate privileges)
- Host paths or volumes (mutations can grant filesystem access)

Use validation policies for security boundaries. Use mutations for operational standards.

#### Related Resources

- [Kyverno Templates Overview](../index.md)
- [Kyverno Generation Templates](../generation/index.md)
- [Kyverno Image Security](../image/index.md)

### Kyverno Network Security Templates

Network policies control traffic between pods, namespaces, and external endpoints. These templates enforce network segmentation and prevent unauthorized communication.

> **Network Policies Require CNI Support**
>
> NetworkPolicy resources only function when your CNI plugin supports them. Verify your cluster's CNI (Calico, Cilium, Weave Net) before deploying network policies.
>

#### Why Network Policy Enforcement Matters

Default Kubernetes behavior allows all pod-to-pod communication. This creates a flat network where any compromised pod can reach any other pod.

**Network policies provide:**

- Namespace isolation (prevent cross-namespace traffic)
- Service-level segmentation (databases only accessible to specific apps)
- Egress controls (block unauthorized external connections)
- Zero-trust networking (explicit allow-lists only)

#### Available Templates

##### [NetworkPolicy Requirements](security.md)

Require NetworkPolicy resources in every namespace:

- Mandate default-deny policies for new namespaces
- Enforce ingress and egress rules for production workloads
- Block namespaces without network segmentation

**Apply a policy:**

```bash
kubectl apply -f security.yaml
```

##### [Ingress Class Validation](ingress-class.md)

Enforce approved IngressClass usage:

- Restrict Ingress resources to approved IngressClass values
- Prevent direct exposure through unapproved ingress controllers
- Validate ingress annotations for security requirements

**Apply a policy:**

```bash
kubectl apply -f ingress-class.yaml
```

##### [Ingress TLS Requirements](ingress-tls.md)

Mandate TLS termination for all Ingress resources:

- Require TLS configuration on all Ingress objects
- Validate TLS secret references exist
- Enforce HTTPS-only traffic for external services

**Apply a policy:**

```bash
kubectl apply -f ingress-tls.yaml
```

##### [Service Type Restrictions](services.md)

Control Service exposure and external access:

- Restrict LoadBalancer and NodePort Service types
- Require annotations for external-facing services
- Validate service selectors and port configurations

**Apply a policy:**

```bash
kubectl apply -f services.yaml
```

#### Network Security Patterns

##### Defense in Depth

Layer network controls across multiple boundaries:

1. **Namespace NetworkPolicies** - Default deny all traffic
2. **Service Restrictions** - Limit LoadBalancer/NodePort usage
3. **Ingress Controls** - Require TLS and approved ingress classes
4. **Egress Filtering** - Block unauthorized external connections

##### Zero-Trust Networking

Never assume trust based on network location:

- Require explicit NetworkPolicy allow rules (no implicit trust)
- Mandate mTLS for service-to-service communication (use service mesh if needed)
- Validate identity at every network boundary (authentication, not IP allowlisting)

##### Production vs Non-Production

Use different enforcement levels based on environment:

- **Production** - Strict NetworkPolicy requirements, TLS mandatory, LoadBalancer restricted
- **Development** - Relaxed policies, allow broader access for testing
- **Staging** - Production-like policies to catch configuration issues early

#### Common Enforcement Scenarios

##### Scenario 1: Prevent Unapproved External Exposure

Block LoadBalancer services except for approved namespaces:

```yaml
### Enforced by: services.yaml
### Result: Only ingress-nginx namespace can create LoadBalancer services
### Impact: Prevents accidental exposure of internal services to the internet
```

##### Scenario 2: Mandate TLS for Public Services

Require TLS configuration on all Ingress resources:

```yaml
### Enforced by: ingress-tls.yaml
### Result: All Ingress objects must define spec.tls with valid secrets
### Impact: Eliminates plaintext HTTP exposure for external services
```

##### Scenario 3: Enforce Namespace Isolation

Require NetworkPolicy in every namespace before pod creation:

```yaml
### Enforced by: security.yaml
### Result: Namespaces must have NetworkPolicy resources before accepting workloads
### Impact: Prevents pods from communicating across namespace boundaries by default
```

#### Testing Network Policies

Validate NetworkPolicy enforcement without disrupting traffic:

```bash
### Test NetworkPolicy requirement (should fail without policy)
kubectl create namespace test-ns
kubectl run test-pod --image=nginx -n test-ns
### Expected: Blocked by policy requiring NetworkPolicy in namespace

### Test Ingress TLS requirement (should fail without TLS)
cat <<EOF | kubectl apply -f -
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: test-ingress
  namespace: test-ns
spec:
  rules:
    - host: test.example.com
      http:
        paths:
          - path: /
            pathType: Prefix
            backend:
              service:
                name: test-service
                port:
                  number: 80
EOF
### Expected: Blocked by policy requiring spec.tls

### Test Service type restriction (should fail for LoadBalancer)
kubectl expose deployment test-app --type=LoadBalancer --port=80 -n test-ns
### Expected: Blocked by policy restricting LoadBalancer type
```

#### Related Resources

- [Kyverno Templates Overview](../index.md)
- [Kyverno Pod Security](../pod-security/index.md)
- [Kyverno Resource Governance](../resource/index.md)

### Kyverno Pod Security Templates

Pod security policies prevent privilege escalation, restrict dangerous capabilities, and enforce security boundaries for containerized workloads.

> **Pod Security Standards Replace PSP**
>
> PodSecurityPolicy was deprecated in Kubernetes 1.21 and removed in 1.25. Use Pod Security Standards (PSS) via admission controllers or Kyverno policies instead.
>

#### Why Pod Security Matters

Containers inherit privileges from their configuration. Without enforcement, workloads can:

- Run as root with unrestricted filesystem access
- Mount host paths to access node secrets or modify system files
- Escalate privileges using dangerous Linux capabilities
- Break out of container isolation through privileged mode

#### Available Templates

##### [Pod Security Standards](standards.md)

Enforce Kubernetes Pod Security Standards (Baseline, Restricted):

- Block privileged containers and hostPath volumes
- Require non-root execution and read-only root filesystems
- Enforce seccomp, AppArmor, and SELinux profiles

**Apply a policy:**

```bash
kubectl apply -f standards.yaml

```

##### [Privilege Restrictions](privileges.md)

Prevent privilege escalation and dangerous execution modes:

- Block `privileged: true` containers
- Prevent `allowPrivilegeEscalation: true`
- Restrict host namespaces (PID, IPC, Network)
- Block host port bindings

**Apply a policy:**

```bash
kubectl apply -f privileges.yaml

```

##### [Security Profiles](profiles.md)

Enforce security profiles and runtime restrictions:

- Require seccomp profiles (RuntimeDefault or custom)
- Mandate AppArmor annotations for workloads
- Enforce SELinux contexts for pod isolation
- Block containers running as UID 0 (root)

**Apply a policy:**

```bash
kubectl apply -f profiles.yaml

```

#### Pod Security Standards Levels

Kubernetes defines three PSS levels. Choose based on risk tolerance.

##### Privileged (Unrestricted)

No restrictions. Only use for trusted system components.

- **Use cases:** CNI plugins, storage drivers, monitoring agents with node access
- **Risk:** Full cluster compromise if container is exploited
- **Recommendation:** Avoid. Use Restricted where possible.

##### Baseline (Minimize Known Privilege Escalations)

Prevents known privilege escalation vectors:

- No privileged containers
- No host namespace sharing
- No host path mounts
- Limited capabilities (drops `ALL`, allows safe subset)

**Use for:** Most production workloads without special requirements.

##### Restricted (Hardened for High-Security Environments)

Enforces current security best practices:

- Non-root execution (`runAsNonRoot: true`)
- Read-only root filesystem
- Seccomp profile required
- Drops all capabilities
- No privilege escalation

**Use for:** Internet-facing services, multi-tenant clusters, compliance requirements.

#### Common Enforcement Scenarios

##### Scenario 1: Block All Privileged Containers

Prevent privileged mode across the cluster:

```yaml
### Enforced by: privileges.yaml
### Result: No containers can run with privileged: true
### Impact: Eliminates most container breakout vectors

```

##### Scenario 2: Require Non-Root Execution

Force all containers to run as non-root users:

```yaml
### Enforced by: profiles.yaml
### Result: Containers must define runAsNonRoot: true
### Impact: Prevents root-level filesystem access and privilege escalation

```

##### Scenario 3: Enforce Seccomp Profiles

Mandate seccomp profiles for syscall filtering:

```yaml
### Enforced by: standards.yaml
### Result: Pods must define securityContext.seccompProfile
### Impact: Reduces kernel attack surface by blocking dangerous syscalls

```

#### Testing Pod Security Policies

Validate enforcement without disrupting production:

```bash
### Test privileged container block (should fail)
kubectl run privileged-test --image=nginx --privileged=true
### Expected: Blocked by privilege restriction policy

### Test root user block (should fail)
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Pod
metadata:
  name: root-test
spec:
  containers:
    - name: nginx
      image: nginx
      securityContext:
        runAsUser: 0
EOF
### Expected: Blocked by non-root requirement policy

### Test hostPath mount block (should fail)
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Pod
metadata:
  name: hostpath-test
spec:
  containers:
    - name: nginx
      image: nginx
      volumeMounts:
        - name: host
          mountPath: /host
  volumes:
    - name: host
      hostPath:
        path: /
EOF
### Expected: Blocked by Pod Security Standards policy

### Test compliant pod (should succeed)
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Pod
metadata:
  name: compliant-test
spec:
  securityContext:
    runAsNonRoot: true
    runAsUser: 1000
    seccompProfile:
      type: RuntimeDefault
  containers:
    - name: nginx
      image: nginx
      securityContext:
        allowPrivilegeEscalation: false
        readOnlyRootFilesystem: true
        capabilities:
          drop:
            - ALL
EOF
### Expected: Allowed by all policies

```

#### Migration from PodSecurityPolicy

Replace deprecated PSPs with Kyverno policies:

1. **Audit current PSP usage:**

   ```bash
   kubectl get psp
   kubectl get pods --all-namespaces -o jsonpath='{range .items[*]}{.metadata.name}{"\t"}{.metadata.annotations.kubernetes\.io/psp}{"\n"}{end}'

   ```

2. **Map PSP rules to Kyverno policies:**
   - `privileged: false` → Use `privileges.yaml`
   - `allowPrivilegeEscalation: false` → Use `privileges.yaml`
   - `runAsUser` rules → Use `profiles.yaml`
   - `volumes` restrictions → Use `standards.yaml`

3. **Deploy Kyverno policies in audit mode:**

   ```bash
   kubectl apply -f standards.yaml
   kubectl apply -f privileges.yaml
   kubectl apply -f profiles.yaml

   ```

4. **Review policy reports for violations:**

   ```bash
   kubectl get polr -A  # Policy Reports
   kubectl describe polr <report-name> -n <namespace>

   ```

5. **Switch to enforce mode after validation:**
   Update `validationFailureAction: Enforce` in policies.

#### Related Resources

- [Kyverno Templates Overview](../index.md)
- [Kyverno Network Security](../network/index.md)
- [OPA Pod Security Templates](../pod-security/index.md)

### Kyverno Policy Templates

> **Start with Audit Mode**
>
> Deploy in `audit` mode first. Existing workloads may violate these policies. Monitor violations for 48 hours, fix non-compliant resources, then switch to `enforce`.
>

Production-ready Kyverno policies for Kubernetes admission control. **28 policies** covering validation, mutation, and generation patterns. Each template includes complete configuration, customization options, validation commands, and real-world use cases.

---

#### Available Templates

##### Pod Security (5 Policies)

Enforce pod security standards, prevent privileged containers, control host namespaces, and enforce security profiles.

**Files:**

- **[Pod Security Standards →](pod-security/standards.md)** (2 policies)
- **[Privilege Escalation Prevention →](pod-security/privileges.md)** (1 policy)
- **[Security Profiles →](pod-security/profiles.md)** (2 policies)

Key policies:

- Pod Security Standards Enforcement (Baseline/Restricted)
- Host Namespace Restrictions (hostNetwork, hostPID, hostIPC, hostPort)
- Privilege Escalation Prevention (allowPrivilegeEscalation, privileged containers)
- Seccomp Profile Enforcement (RuntimeDefault, Localhost, unconfined blocking)
- AppArmor Profile Requirements (runtime/default, custom profiles)

---

##### Image Validation (5 Policies)

Control container images with digest requirements, registry allowlists, signature verification, base image enforcement, and CVE scanning gates.

**Files:**

- **[Image Digest & Registry Validation →](image/validation.md)** (2 policies)
- **[Image Signing Verification →](image/signing.md)** (1 policy)
- **[Base Image Enforcement →](image/security.md)** (1 policy)
- **[CVE Scanning Gates →](image/cve-scanning.md)** (1 policy)

Key policies:

- Image Digest Requirements (SHA256 enforcement)
- Registry Allowlist and Tag Validation (block `latest`, untrusted registries)
- Cosign Image Signature Verification (keyless and key-based)
- Base Image Enforcement (approved base images, deprecated blocklist)
- CVE Scanning Gates (Trivy attestations, severity thresholds)

---

##### Resource Management (5 Policies)

Ensure resource requests and limits, enforce CPU/memory ratios, control ephemeral storage, constrain PVC sizes, and require HPA configuration.

**Files:**

- **[Resource Limits & Ratios →](resource/limits.md)** (2 policies)
- **[Storage Limits →](resource/storage.md)** (2 policies)
- **[HPA Requirements →](resource/hpa.md)** (1 policy)

Key policies:

- Resource Limits and Requests Enforcement (CPU, memory, QoS classes)
- CPU and Memory Ratio Enforcement (prevent over-provisioning)
- Ephemeral Storage Limits (ephemeral storage, emptyDir controls)
- PVC Size Constraints (min/max sizes, storage class governance)
- HPA Configuration Requirements (replica bounds, metrics validation)

---

##### Network Security (5 Policies)

Enforce network policies, restrict egress traffic, require ingress class validation, mandate TLS encryption, and control service types.

**Files:**

- **[Network Policies & Egress →](network/security.md)** (2 policies)
- **[Ingress Class Requirements →](network/ingress-class.md)** (1 policy)
- **[Ingress TLS Requirements →](network/ingress-tls.md)** (1 policy)
- **[Service Type Restrictions →](network/services.md)** (1 policy)

Key policies:

- Require Network Policies (namespace coverage, default-deny enforcement)
- Egress Restrictions (destination controls, external IP blocking)
- Ingress Class Requirements (approved controllers, deprecated annotation blocking)
- Ingress TLS Requirements (encryption enforcement, cert-manager integration)
- Service Type Restrictions (LoadBalancer approval, NodePort controls)

---

##### Mutation & Generation (7 Policies)

Automatically inject labels, add sidecars, generate resource quotas, create network policies, and ensure pod disruption budgets.

**Files:**

- **[Label Mutation →](mutation/labels.md)** (2 policies)
- **[Sidecar Injection →](mutation/sidecar.md)** (2 policies)
- **[Namespace Resource Generation →](generation/namespace.md)** (2 policies)
- **[Workload Resource Generation →](generation/workload.md)** (1 policy)

Key policies:

- Default Label Injection (team, environment, version, cost-center)
- Namespace Label Propagation (inherit team, compliance, SLA labels)
- Logging Sidecar Injection (Fluent Bit with Elasticsearch/Loki)
- Monitoring Sidecar Injection (Nginx exporter, JMX exporter, Prometheus)
- Automatic ResourceQuota Generation (default quotas, production quotas)
- Default-Deny NetworkPolicy Generation (default-deny ingress, strict egress)
- Automatic PodDisruptionBudget Generation (2+ replicas, critical workloads)

---

##### Mandatory Labels (1 Policy)

Enforce required metadata for observability, cost tracking, and compliance auditing.

**Files:**

- **[Mandatory Labels →](labels.md)** (1 policy)

Key policy:

- Mandatory Labels and Annotations (require app, team, version, environment labels)

---

#### Policy Types

Kyverno supports three policy types:

##### Validation Policies

**Block** resources that violate security rules.

Examples: Pod security restrictions, image allowlists, resource limits, network security

##### Mutation Policies

**Modify** resources before admission to enforce standards.

Examples: Add labels, inject sidecars, set default resource limits

##### Generation Policies

**Create** new resources when triggers match.

Examples: Generate ResourceQuotas for new namespaces, create default-deny NetworkPolicies

---

#### Quick Start

All templates follow the same deployment pattern:

```bash
### Apply policy in audit mode first
kubectl apply -f policy.yaml

### Monitor policy violations
kubectl logs -f -n kyverno deployment/kyverno

### Check policy reports
kubectl get polr -A  # PolicyReports
kubectl get cpolr    # ClusterPolicyReports

### Switch to enforce mode after validation
kubectl patch clusterpolicy <policy-name> \
  --type merge \
  -p '{"spec":{"validationFailureAction":"enforce"}}'
```

#### Policy Customization

Every template includes a customization table:

| Variable | Default | Purpose |
|----------|---------|---------|
| `validationFailureAction` | `audit` | Use `audit` for testing, `enforce` for production |
| `background` | `true` | Scan existing resources (not just new admission requests) |
| Resource selectors | Varies | Target specific namespaces, kinds, or labels |

#### Related Resources

- **[JMESPath Patterns →](../jmespath/patterns.md)** - Advanced Kyverno pattern examples
- **[OPA Templates →](../opa/index.md)** - Gatekeeper constraint templates
- **[Decision Guide →](../decision-guide.md)** - OPA vs Kyverno selection guide
- **[Template Library Overview →](index.md)** - Back to main page

### Kyverno Resource Governance Templates

Resource governance policies prevent overconsumption, enforce autoscaling requirements, and control storage allocation across your cluster.

> **Resource Limits Prevent Noisy Neighbors**
>
> Without resource limits, a single pod can consume all node capacity and starve other workloads. Enforce limits to guarantee fair resource allocation.
>

#### Why Resource Governance Matters

Kubernetes does not enforce resource limits by default. This creates operational risks:

- **Node Exhaustion** - Pods without limits can consume all CPU/memory
- **OOMKilled Pods** - Memory requests too low cause evictions
- **Autoscaling Failures** - HPA requires resource metrics from limits/requests
- **Cost Overruns** - Uncontrolled storage provisioning inflates cloud bills

#### Available Templates

##### [Resource Limits and Requests](limits.md)

Enforce CPU and memory limits/requests on all workloads:

- Require resource requests for scheduling decisions
- Mandate resource limits to prevent node exhaustion
- Validate requests ≤ limits for all containers
- Block workloads with excessive resource claims

**Apply a policy:**

```bash
kubectl apply -f limits.yaml
```

##### [Horizontal Pod Autoscaler Requirements](hpa.md)

Mandate HPA for production workloads:

- Require HPA for Deployments in production namespaces
- Validate HPA min/max replica bounds
- Ensure HPA targets exist and are valid
- Block HPAs without resource-based metrics

**Apply a policy:**

```bash
kubectl apply -f hpa.yaml
```

##### [Storage Constraints](storage.md)

Control PersistentVolume and PersistentVolumeClaim allocation:

- Restrict PVC sizes to prevent excessive storage claims
- Require specific StorageClasses for production data
- Block dynamic provisioning in restricted namespaces
- Validate volume access modes and reclaim policies

**Apply a policy:**

```bash
kubectl apply -f storage.yaml
```

#### Resource Management Patterns

##### Resource Quotas vs Limits

Use both mechanisms for defense in depth:

- **ResourceQuota** - Namespace-level caps (total CPU/memory across all pods)
- **LimitRange** - Default and max values for individual pods
- **Kyverno Policies** - Validation and enforcement of resource configuration

Kyverno policies complement quotas by validating workload-level configuration before admission.

##### Right-Sizing Workloads

Set appropriate resource values to balance cost and reliability:

- **Requests too low** → Pods scheduled on undersized nodes → OOMKilled
- **Requests too high** → Wasted capacity → Increased costs
- **Limits too low** → Pods throttled → Performance degradation
- **Limits too high** → Noisy neighbor problems → Node instability

Use Vertical Pod Autoscaler (VPA) recommendations to identify optimal values.

##### Autoscaling Strategies

Choose the right autoscaling mechanism for your workload:

- **HPA (Horizontal)** - Scale replicas based on CPU/memory/custom metrics
- **VPA (Vertical)** - Adjust resource requests/limits automatically
- **Cluster Autoscaler** - Add/remove nodes based on pending pods

Kyverno policies enforce HPA presence and configuration validity.

#### Common Enforcement Scenarios

##### Scenario 1: Prevent Unbounded Resource Consumption

Require resource limits on all containers:

```yaml
### Enforced by: limits.yaml
### Result: All containers must define resources.limits.cpu and resources.limits.memory
### Impact: Prevents single pod from consuming entire node capacity
```

##### Scenario 2: Mandate Autoscaling for Production

Require HPA for production Deployments:

```yaml
### Enforced by: hpa.yaml
### Result: Deployments in prod-* namespaces must have corresponding HPA
### Impact: Ensures production services scale automatically under load
```

##### Scenario 3: Control Storage Costs

Restrict PVC size to prevent excessive allocations:

```yaml
### Enforced by: storage.yaml
### Result: PVCs cannot exceed 100Gi in dev namespaces
### Impact: Prevents accidental provisioning of expensive storage volumes
```

#### Testing Resource Policies

Validate enforcement without disrupting workloads:

```bash
### Test resource limit requirement (should fail without limits)
kubectl run no-limits --image=nginx
### Expected: Blocked by policy requiring resource limits

### Test excessive resource request (should fail if beyond policy limits)
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Pod
metadata:
  name: excessive-request
spec:
  containers:
    - name: nginx
      image: nginx
      resources:
        requests:
          cpu: "100"
          memory: "1000Gi"
EOF
### Expected: Blocked by policy restricting maximum requests

### Test HPA requirement (should fail without HPA)
kubectl create deployment test-app --image=nginx --replicas=3 -n production
### Expected: Blocked by policy requiring HPA for production Deployments

### Test storage size restriction (should fail for excessive PVC)
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: large-pvc
  namespace: dev
spec:
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 500Gi
EOF
### Expected: Blocked by policy restricting dev namespace PVC sizes

### Test compliant workload (should succeed)
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Pod
metadata:
  name: compliant-pod
spec:
  containers:
    - name: nginx
      image: nginx
      resources:
        requests:
          cpu: "100m"
          memory: "128Mi"
        limits:
          cpu: "500m"
          memory: "512Mi"
EOF
### Expected: Allowed by resource limit policies
```

#### Resource Budgeting

Plan cluster capacity using policy-enforced constraints:

1. **Calculate namespace budgets:**

   ```bash
   kubectl get resourcequota -n production -o yaml
##   # Review quota limits vs current usage
   ```

2. **Identify outliers:**

   ```bash
   kubectl top pods --all-namespaces --sort-by=memory
   kubectl top pods --all-namespaces --sort-by=cpu
   ```

3. **Validate policy alignment:**
   - Do LimitRanges match policy-enforced limits?
   - Are ResourceQuotas enforced at namespace level?
   - Do HPA min/max bounds align with capacity planning?

4. **Monitor policy violations:**

   ```bash
   kubectl get polr -A  # Policy Reports
##   # Review workloads blocked by resource policies
   ```

#### Cost Optimization

Use resource policies to reduce cloud infrastructure costs:

- **Right-size workloads** - Block oversized resource requests
- **Prevent storage sprawl** - Restrict PVC sizes in non-production
- **Enforce autoscaling** - Scale down during off-peak hours with HPA
- **Use cheaper storage classes** - Require specific StorageClasses for dev/test

#### Related Resources

- [Kyverno Templates Overview](../index.md)
- [Kyverno Pod Security](../pod-security/index.md)
- [OPA Resource Governance](../resource/index.md)

### Local Development with Policy-as-Code

Run the same policy validation locally that runs in CI. Catch issues in seconds, not hours.

#### Overview

The policy-platform container includes all tools needed for local policy validation:

- **Kyverno CLI** - Policy validation and testing
- **Pluto** - Deprecated API detection
- **Helm** - Chart rendering and linting
- **Spectral** - OpenAPI/values schema validation
- **yq** - YAML processing

> **Zero Local Setup Required**
>
> One container contains all policies and tools. No local installations. Pull the container, run validations. Same environment as CI.
>

---

#### The Local Development Container

##### What's Inside

The policy-platform container is a multi-stage build that aggregates policies from multiple repositories:

```dockerfile
### Pull policy repositories as OCI containers
FROM security-policy-repo:main AS security_policy_repo
FROM devops-policy-repo:main AS devops_policy_repo

### Alpine base with all tools
FROM alpine:3.24
RUN apk add curl bash ca-certificates git helm yq

### Install Kyverno CLI
RUN curl -sSL https://github.com/kyverno/kyverno/releases/download/v1.13.2/kyverno-cli_v1.13.2_linux_x86_64.tar.gz \
  | tar -xz -C /usr/local/bin

### Install Pluto
RUN curl -sSL https://github.com/FairwindsOps/pluto/releases/download/v5.21.1/pluto_5.21.1_linux_amd64.tar.gz \
  | tar xz -C /usr/local/bin

### Copy policies from dependent containers
COPY --from=security_policy_repo /repos/security-policy/ /repos/security-policy/
COPY --from=devops_policy_repo /repos/devops-policy/ /repos/devops-policy/

WORKDIR /repos
```

**Key Insight**: Policy repos are also OCI containers. Multi-stage build pulls them automatically.

---

#### Basic Usage

##### Running Policy Validation

Validate a Kubernetes manifest against all policies:

```bash
docker run --rm \
  -v $(pwd):/workspace \
  policy-platform:latest \
  kyverno apply /repos/security-policy/ \
  --resource /workspace/deployment.yaml
```

**Output**:

```text
Applying 1 policy to 1 resource...

pass: 12/12
fail: 0/12
warn: 0/12
error: 0/12
skip: 0/12

All resources passed policy validation!
```

##### Validation Failure Example

```bash
$ docker run --rm -v $(pwd):/workspace policy-platform:latest \
  kyverno apply /repos/security-policy/ \
  --resource /workspace/bad-deployment.yaml

fail: 2/12
  require-resource-limits:
    Deployment/default/nginx: CPU and memory limits required
  disallow-latest-tag:
    Deployment/default/nginx: Container uses :latest tag
```

> **Fix Before Committing**
>
> Local validation catches issues before CI. Fix violations now - no 20-minute CI feedback loop.
>

---

#### Helm Chart Validation

##### Rendering Charts with Environment Values

Real-world Helm charts need environment-specific values:

```bash
### Render chart for staging environment
docker run --rm \
  -v $(pwd):/workspace \
  policy-platform:latest \
  helm template my-app /workspace/charts/my-app \
    -f /workspace/charts/my-app/values.yaml \
    -f /workspace/cd/staging/values.yaml \
  > staging-manifests.yaml

### Validate rendered manifests
docker run --rm \
  -v $(pwd):/workspace \
  policy-platform:latest \
  kyverno apply /repos/security-policy/ \
    --resource /workspace/staging-manifests.yaml
```

**This matches exactly what CI does.**

##### Multi-Environment Validation

Validate across all environments before pushing:

```bash
for env in dev qac stg prd; do
  echo "Validating ${env} environment..."

##  # Render manifests
  docker run --rm -v $(pwd):/workspace policy-platform:latest \
    helm template app /workspace/charts/app \
      -f /workspace/charts/app/values.yaml \
      -f /workspace/cd/${env}/values.yaml \
    > ${env}-manifests.yaml

##  # Validate policies
  docker run --rm -v $(pwd):/workspace policy-platform:latest \
    kyverno apply /repos/security-policy/ \
      --resource /workspace/${env}-manifests.yaml \
      --audit-warn
done
```

**Catch environment-specific issues locally.**

---

#### Next Steps

- **[Advanced Validation](advanced-validation.md)** - Policy reports, deprecated API detection, schema validation
- **[Workflow Integration](workflow-integration.md)** - Pre-commit hooks, Make targets, troubleshooting
- **[CI Integration](../ci-integration/index.md)** - Automate policy checks in pipelines

### Multi-Source Policy Aggregation

Combine policies from multiple repositories into a single enforcement container.

#### Overview

Real-world policy management requires aggregating policies from different teams and sources:

```mermaid
graph TD
    SP[security-policy repo] -->|OCI container| MS[Multi-stage Build]
    DP[devops-policy repo] -->|OCI container| MS
    KC[Kyverno curated] -->|OCI container| MS
    APP[Application repo] -->|OCI container| MS

    MS -->|Single image| PP[policy-platform:latest]

    PP -->|Contains all policies| OUT[Unified Enforcement]

    %% Ghostty Hardcore Theme
    style SP fill:#f92572,color:#1b1d1e
    style DP fill:#fd971e,color:#1b1d1e
    style KC fill:#9e6ffe,color:#1b1d1e
    style APP fill:#a7e22e,color:#1b1d1e
    style PP fill:#65d9ef,color:#1b1d1e

```

> **Policy Repos as OCI Containers**
>
> Each policy repository is **also** an OCI container. Multi-stage Docker builds pull them all automatically. No manual copying or Git submodules.
>

---

#### Policy Repository Structure

##### Repository as OCI Container

Each policy repository builds its own container:

**security-policy repo Dockerfile**:

```dockerfile
FROM scratch
COPY . /repos/security-policy/
```

**Build and push**:

```bash
docker build -t security-policy-repo:main .
docker push europe-west6-docker.pkg.dev/ops/charts/security-policy-repo:main
```

**Result**: Policies packaged as artifact, versioned, and distributed via registry.

---

#### Multi-Stage Build Pattern

##### Aggregating Policy Sources

The policy-platform Dockerfile uses multi-stage builds to pull all policy sources:

```dockerfile
### Stage 1: Pull DevOps policies
FROM europe-west6-docker.pkg.dev/ops/charts/devops-policy-repo:main AS devops_policy_repo

### Stage 2: Pull Security policies
FROM europe-west6-docker.pkg.dev/ops/charts/security-policy-repo:main AS security_policy_repo

### Stage 3: Pull Application-specific policies
FROM europe-west6-docker.pkg.dev/ops/charts/backend-applications-repo:main AS backend_applications_repo

### Final stage: Combine everything
FROM alpine:3.24.1

### Install policy tools
RUN apk add --no-cache \
    curl bash ca-certificates git tar gzip \
    libc6-compat helm yq

### Install Kyverno CLI
RUN mkdir -p ./bin/kyverno && \
    curl -sSL https://github.com/kyverno/kyverno/releases/download/v1.13.2/kyverno-cli_v1.13.2_linux_x86_64.tar.gz \
      | tar -xz -C ./bin/kyverno && \
    chmod +x ./bin/kyverno/kyverno && \
    cp ./bin/kyverno/kyverno /usr/local/bin/

### Install Pluto
RUN mkdir -p ./bin/pluto && \
    curl -sSL https://github.com/FairwindsOps/pluto/releases/download/v5.21.1/pluto_5.21.1_linux_amd64.tar.gz \
      | tar xz -C ./bin/pluto && \
    chmod +x ./bin/pluto && \
    cp ./bin/pluto/pluto /usr/local/bin/

### Install Spectral
RUN mkdir -p ./bin/spectral && \
    latest_spectral=$(curl -sSL https://api.github.com/repos/stoplightio/spectral/releases/latest | grep 'tag_name' | cut -d\" -f4) && \
    curl -sSL https://github.com/stoplightio/spectral/releases/download/${latest_spectral}/spectral-alpine-x64 \
      -o ./bin/spectral/spectral && \
    chmod +x ./bin/spectral/spectral && \
    cp ./bin/spectral/spectral /usr/local/bin/

### Copy all policy repositories
COPY --from=devops_policy_repo /repos/devops-policy/ /repos/devops-policy/
COPY --from=security_policy_repo /repos/security-policy/ /repos/security-policy/
COPY --from=backend_applications_repo /repos/backend-applications/ /repos/backend-applications/

WORKDIR /repos
```

**Architecture**: Each `COPY --from=` pulls a policy repo. Final image contains all.

---

#### Policy Repository Organization

##### Security Policy Repository

**Directory structure**:

```text
security-policy/
├── charts/
│   └── security-policy/
│       ├── Chart.yaml
│       ├── values.yaml
│       └── templates/
│           ├── require-resource-limits.yaml
│           ├── disallow-latest-tag.yaml
│           ├── require-probes.yaml
│           ├── disallow-privileged.yaml
│           └── ...
├── cd/
│   ├── values.yaml          # Common values
│   ├── dev/values.yaml      # Dev overrides
│   ├── qac/values.yaml
│   ├── stg/values.yaml
│   └── prd/values.yaml
└── Dockerfile               # Build repo as OCI container
```

**Example policy** (`require-resource-limits.yaml`):

```yaml
apiVersion: kyverno.io/v1
kind: ClusterPolicy
metadata:
  name: require-resource-limits
  annotations:
    policies.kyverno.io/title: Require Resource Limits
    policies.kyverno.io/category: Best Practices
    policies.kyverno.io/severity: medium
spec:
  validationFailureAction: {{ .Values.policies.resourceLimits.validationFailureAction }}
  background: true
  rules:
    - name: check-cpu-memory
      match:
        resources:
          kinds:
            - Deployment
            - StatefulSet
            - DaemonSet
      validate:
        message: "CPU and memory limits required"
        pattern:
          spec:
            template:
              spec:
                containers:
                  - resources:
                      limits:
                        memory: "?*"
                        cpu: "?*"
```

##### DevOps Policy Repository

**Focus**: Operational best practices

```text
devops-policy/
├── charts/
│   └── devops-policy/
│       └── templates/
│           ├── require-labels.yaml         # Standard labels
│           ├── require-annotations.yaml    # Ownership metadata
│           ├── naming-conventions.yaml     # Resource naming
│           └── ...
├── cd/
│   ├── dev/values.yaml
│   └── prd/values.yaml
└── Dockerfile
```

**Example policy** (`require-labels.yaml`):

```yaml
apiVersion: kyverno.io/v1
kind: ClusterPolicy
metadata:
  name: require-standard-labels
spec:
  validationFailureAction: {{ .Values.policies.labels.validationFailureAction }}
  background: true
  rules:
    - name: check-labels
      match:
        resources:
          kinds:
            - Deployment
            - StatefulSet
      validate:
        message: "Standard labels required: app, environment, team"
        pattern:
          metadata:
            labels:
              app: "?*"
              environment: "?*"
              team: "?*"
```

> **Separation of Concerns**
>
> Security team owns security-policy repo. DevOps team owns devops-policy repo. Each team maintains their domain independently.
>

##### Application-Specific Policies

**Focus**: Application Helm chart schemas and configurations

```text
backend-applications/
├── charts/
│   └── backend-applications/
│       ├── values.yaml
│       └── ...
├── .spectral.yaml           # Values schema validation
└── Dockerfile
```

**.spectral.yaml** (schema validation):

```yaml
extends: spectral:oas
rules:
  required-fields:
    given: $
    severity: error
    then:
      - field: name
        function: truthy
      - field: namespace
        function: truthy
      - field: replicas
        function: truthy
  replica-count:
    given: $.replicas
    severity: warning
    then:
      function: schema
      functionOptions:
        schema:
          type: integer
          minimum: 1
          maximum: 10
```

---

#### Next Steps

- **[Policy Management](policy-management.md)** - Rendering, versioning, and discovery
- **[Testing and Operations](testing-and-operations.md)** - Testing multi-source policies
- **[Policy Packaging](../policy-packaging/index.md)** - Build the policy-platform container

### OPA Image Security Templates

Image security policies control which container images can run in your cluster. These templates enforce registry allowlists, require immutable digests, and validate cryptographic signatures.

> **Image Tags Are Mutable**
>
> Tags like `latest` or `v1.2.3` can be overwritten by attackers who compromise registries. Use digest-based references (`sha256:...`) for immutable deployments.
>

#### Why Image Security Matters

Container images are the primary attack vector for supply chain compromises:

- **Registry Poisoning** - Attackers push malicious images to public registries
- **Tag Mutation** - `latest` tags updated with backdoored code
- **Typosquatting** - Misspelled image names redirect to attacker-controlled registries
- **Unsigned Images** - No cryptographic proof of provenance

#### Available Templates

##### [Registry Allowlist](base.md)

Restrict container images to approved registries:

- Block public registries (Docker Hub, Quay, GCR)
- Allow only corporate registry domains
- Enforce namespace-specific registry restrictions
- Prevent deployment of untrusted images

**Apply a policy:**

```bash
kubectl apply -f base.yaml

```

##### [Digest Enforcement](digest.md)

Require immutable digest references instead of mutable tags:

- Block tag-based image references (`nginx:latest`)
- Mandate digest-based references (`nginx@sha256:...`)
- Validate digest format and checksum integrity
- Prevent image tag mutation attacks

**Apply a policy:**

```bash
kubectl apply -f digest.yaml

```

##### [Image Scanning Requirements](security.md)

Enforce vulnerability scanning and security assessments:

- Require scan annotations on all images
- Block images with high/critical CVEs
- Validate scan freshness (no stale scans)
- Enforce minimum scan score thresholds

**Apply a policy:**

```bash
kubectl apply -f security.yaml

```

##### [Signature Verification](verification.md)

Validate cryptographic signatures using Cosign or Notary:

- Require valid signatures from trusted keys
- Block unsigned or invalidly signed images
- Enforce signature transparency logs (Rekor)
- Validate attestations for build provenance

**Apply a policy:**

```bash
kubectl apply -f verification.yaml

```

#### Image Security Defense Layers

Implement multiple controls for defense in depth:

1. **Registry Allowlist** - Only approved registries (base.yaml)
2. **Digest Enforcement** - Immutable image references (digest.yaml)
3. **Vulnerability Scanning** - No high/critical CVEs (security.yaml)
4. **Signature Verification** - Cryptographic provenance (verification.yaml)

Each layer addresses different attack vectors. Use all four for production environments.

#### Common Enforcement Scenarios

##### Scenario 1: Block Public Registries

Prevent deployment of images from untrusted sources:

```yaml
### Enforced by: base.yaml
### Result: Only images from registry.company.com allowed
### Impact: Eliminates supply chain attacks via public registries

```

##### Scenario 2: Prevent Tag Mutation

Require digest-based image references:

```yaml
### Enforced by: digest.yaml
### Result: Image references must use @sha256:... format
### Impact: Guarantees deployed image matches approved version

```

##### Scenario 3: Block Vulnerable Images

Reject images with known CVEs:

```yaml
### Enforced by: security.yaml
### Result: Images must have scan results with no high/critical vulnerabilities
### Impact: Prevents deployment of exploitable container images

```

##### Scenario 4: Verify Build Provenance

Validate cryptographic signatures on all images:

```yaml
### Enforced by: verification.yaml
### Result: Images must be signed by trusted key in KMS
### Impact: Ensures images originated from approved CI/CD pipelines

```

#### Testing Image Security Policies

Validate enforcement without blocking legitimate workloads:

```bash
### Test registry allowlist (should fail for Docker Hub)
kubectl run docker-hub-test --image=nginx:latest
### Expected: Blocked by registry allowlist policy

### Test digest requirement (should fail for tag-based reference)
kubectl run tag-test --image=registry.company.com/nginx:v1.21
### Expected: Blocked by digest enforcement policy

### Test digest-based reference (should succeed if from approved registry)
kubectl run digest-test --image=registry.company.com/nginx@sha256:abcdef123456...
### Expected: Allowed by registry and digest policies

### Test unsigned image (should fail without valid signature)
kubectl run unsigned-test --image=registry.company.com/app@sha256:123456...
### Expected: Blocked by signature verification policy

### Test signed image (should succeed with valid Cosign signature)
### First, sign the image:
### cosign sign --key cosign.key registry.company.com/app@sha256:123456...
kubectl run signed-test --image=registry.company.com/app@sha256:123456...
### Expected: Allowed after signature verification

```

#### Signature Verification with Cosign

Deploy Cosign-based signature verification:

1. **Generate signing keys:**

   ```bash
   cosign generate-key-pair
##   # Creates cosign.key (private) and cosign.pub (public)

   ```

2. **Sign container images in CI/CD:**

   ```bash
   cosign sign --key cosign.key registry.company.com/app:v1.2.3

   ```

3. **Store public key in Kubernetes Secret:**

   ```bash
   kubectl create secret generic cosign-pub \
     --from-file=cosign.pub=./cosign.pub \
     -n opa-system

   ```

4. **Configure OPA policy to verify signatures:**
   Policy references `cosign-pub` Secret for signature validation.

5. **Validate signature verification:**

   ```bash
   kubectl run test-app --image=registry.company.com/app@sha256:...
##   # OPA validates signature before admission

   ```

#### Image Scanning Integration

Integrate vulnerability scanning into admission control:

##### Scan in CI/CD

Scan images during build and push scan results as annotations:

```bash
### Trivy example
trivy image --format json registry.company.com/app:v1.2.3 > scan.json

### Push image with scan annotation
crane mutate registry.company.com/app:v1.2.3 \
  --annotation trivy.scan.result="$(cat scan.json)"

```

##### Enforce Scan Results

OPA policies read scan annotations and block vulnerable images:

```rego
### Pseudo-code: Full implementation in security.yaml
deny[msg] {
  scan_result := input.metadata.annotations["trivy.scan.result"]
  criticals := count(scan_result.Results[_].Vulnerabilities[_] | _.Severity == "CRITICAL")
  criticals > 0
  msg := sprintf("Image has %d critical vulnerabilities", [criticals])
}

```

#### Migrating from Tag-Based to Digest-Based Deployments

Transition existing workloads to use digests:

1. **Audit current image references:**

   ```bash
   kubectl get pods --all-namespaces -o jsonpath='{range .items[*]}{.spec.containers[*].image}{"\n"}{end}' | sort -u

   ```

2. **Convert tags to digests:**

   ```bash
##   # Get digest for tagged image
   crane digest registry.company.com/nginx:v1.21
##   # Output: sha256:abcdef123456...

##   # Update deployment to use digest
   kubectl set image deployment/nginx nginx=registry.company.com/nginx@sha256:abcdef123456...

   ```

3. **Deploy digest enforcement policy in audit mode:**
   Set `enforcementAction: warn` to identify non-compliant workloads.

4. **Fix violations and enable enforcement:**
   After all workloads use digests, set `enforcementAction: deny`.

#### Related Resources

- [OPA Templates Overview](../index.md)
- [OPA Pod Security](../pod-security/index.md)
- [Kyverno Image Validation](../image/index.md)

### OPA Pod Security Templates

Pod security policies written in Rego prevent privilege escalation and enforce security boundaries for containerized workloads.

> **Capabilities Bypass Security Boundaries**
>
> Linux capabilities grant fine-grained privileges. A container with `CAP_SYS_ADMIN` can bypass most kernel security mechanisms. Drop all capabilities by default.
>

#### Why Pod Security Matters

Container isolation relies on kernel namespaces, cgroups, and capabilities. Misconfigurations break this isolation:

- **Privileged Mode** - Disables all security boundaries
- **Host Namespaces** - Shares PID/IPC/Network with node
- **Dangerous Capabilities** - `CAP_SYS_ADMIN`, `CAP_NET_RAW`, `CAP_SYS_PTRACE`
- **Root Execution** - Unnecessary privileges for most workloads

#### Available Templates

##### [Privileged Container Prevention](overview.md)

Block privileged containers and host access:

- Prevent `privileged: true` in container security contexts
- Block hostPath, hostPID, hostIPC, hostNetwork usage
- Restrict host port bindings
- Prevent sharing node resources with containers

**Apply a policy:**

```bash
kubectl apply -f overview.yaml

```

##### [Capability Restrictions](capabilities.md)

Control Linux capabilities granted to containers:

- Require dropping ALL capabilities by default
- Block dangerous capabilities (SYS_ADMIN, NET_RAW, SYS_PTRACE)
- Allow safe capabilities only (NET_BIND_SERVICE, CHOWN)
- Validate capability drops in security contexts

**Apply a policy:**

```bash
kubectl apply -f capabilities.yaml

```

##### [Security Context Enforcement](contexts.md)

Mandate security context configuration:

- Require non-root user execution
- Enforce read-only root filesystems
- Require runAsNonRoot: true
- Block privilege escalation via allowPrivilegeEscalation

**Apply a policy:**

```bash
kubectl apply -f contexts.yaml

```

##### [Privilege Escalation Prevention](escalation.md)

Block privilege escalation mechanisms:

- Prevent `allowPrivilegeEscalation: true`
- Block setuid/setgid binaries in containers
- Enforce seccomp profiles to restrict syscalls
- Validate AppArmor/SELinux profiles

**Apply a policy:**

```bash
kubectl apply -f escalation.yaml

```

#### Pod Security Defense Layers

Implement overlapping controls for defense in depth:

1. **Privileged Prevention** - Block privileged mode and host access (overview.yaml)
2. **Capability Restrictions** - Drop dangerous Linux capabilities (capabilities.yaml)
3. **Non-Root Execution** - Require runAsNonRoot (contexts.yaml)
4. **Escalation Prevention** - Block privilege escalation paths (escalation.yaml)

Each layer addresses different attack vectors. Production workloads should pass all four.

#### Common Enforcement Scenarios

##### Scenario 1: Block Privileged Containers

Prevent unrestricted container execution:

```yaml
### Enforced by: overview.yaml
### Result: No containers can run with privileged: true
### Impact: Eliminates most container breakout vectors

```

##### Scenario 2: Drop Dangerous Capabilities

Remove capabilities that grant excessive privileges:

```yaml
### Enforced by: capabilities.yaml
### Result: All containers must drop CAP_SYS_ADMIN, CAP_NET_RAW
### Impact: Prevents kernel manipulation and network sniffing

```

##### Scenario 3: Enforce Non-Root Execution

Require all containers to run as non-root users:

```yaml
### Enforced by: contexts.yaml
### Result: Containers must define runAsNonRoot: true and runAsUser > 0
### Impact: Prevents root-level filesystem access and privilege escalation

```

##### Scenario 4: Block Privilege Escalation

Prevent containers from gaining privileges after start:

```yaml
### Enforced by: escalation.yaml
### Result: Containers must set allowPrivilegeEscalation: false
### Impact: Blocks setuid binaries and capability inheritance

```

#### Testing Pod Security Policies

Validate enforcement without disrupting workloads:

```bash
### Test privileged container block (should fail)
kubectl run privileged-test --image=nginx --privileged=true
### Expected: Admission denied by overview.yaml

### Test capability violation (should fail with CAP_SYS_ADMIN)
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Pod
metadata:
  name: cap-test
spec:
  containers:
    - name: nginx
      image: nginx
      securityContext:
        capabilities:
          add:
            - SYS_ADMIN
EOF
### Expected: Admission denied by capabilities.yaml

### Test root execution (should fail with runAsUser: 0)
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Pod
metadata:
  name: root-test
spec:
  containers:
    - name: nginx
      image: nginx
      securityContext:
        runAsUser: 0
EOF
### Expected: Admission denied by contexts.yaml

### Test privilege escalation (should fail with allowPrivilegeEscalation: true)
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Pod
metadata:
  name: escalation-test
spec:
  containers:
    - name: nginx
      image: nginx
      securityContext:
        allowPrivilegeEscalation: true
EOF
### Expected: Admission denied by escalation.yaml

### Test compliant pod (should succeed)
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Pod
metadata:
  name: compliant-test
spec:
  securityContext:
    runAsNonRoot: true
    runAsUser: 1000
    seccompProfile:
      type: RuntimeDefault
  containers:
    - name: nginx
      image: nginx
      securityContext:
        allowPrivilegeEscalation: false
        readOnlyRootFilesystem: true
        capabilities:
          drop:
            - ALL
EOF
### Expected: Admission allowed by all policies

```

#### Understanding Linux Capabilities

Capabilities split root privileges into fine-grained permissions:

##### Safe Capabilities (Usually Allowed)

- `NET_BIND_SERVICE` - Bind to ports < 1024
- `CHOWN` - Change file ownership
- `DAC_OVERRIDE` - Bypass file permission checks
- `SETGID` / `SETUID` - Change user/group IDs

##### Dangerous Capabilities (Always Blocked)

- `SYS_ADMIN` - Virtually unlimited kernel access
- `NET_RAW` - Create raw sockets (packet sniffing)
- `SYS_PTRACE` - Debug arbitrary processes (credential theft)
- `SYS_MODULE` - Load kernel modules
- `CAP_SYS_BOOT` - Reboot system

**Best practice:** Drop ALL capabilities, then add only required safe capabilities.

#### Security Context Configuration

Every pod should define security contexts at both pod and container levels:

##### Pod-Level Security Context

```yaml
spec:
  securityContext:
    runAsNonRoot: true
    runAsUser: 1000
    runAsGroup: 1000
    fsGroup: 1000
    seccompProfile:
      type: RuntimeDefault

```

##### Container-Level Security Context

```yaml
spec:
  containers:
    - name: app
      securityContext:
        allowPrivilegeEscalation: false
        readOnlyRootFilesystem: true
        capabilities:
          drop:
            - ALL
          add:
            - NET_BIND_SERVICE  # Only if binding to port 80/443

```

#### Migration from Permissive to Restrictive Policies

Transition existing workloads to secure configurations:

1. **Audit current security contexts:**

   ```bash
   kubectl get pods --all-namespaces -o json | \
     jq '.items[] | select(.spec.containers[].securityContext.privileged == true)'

   ```

2. **Deploy policies in audit mode:**
   Use OPA audit mode to log violations without blocking:

   ```bash
   kubectl apply -f overview.yaml  # Set enforcementAction: warn

   ```

3. **Review violations:**

   ```bash
   kubectl get constrainttemplates
   kubectl get <constraint-name> -o yaml
##   # Check status.violations for non-compliant pods

   ```

4. **Fix workload security contexts:**
   Update Deployments/StatefulSets to add security context fields.

5. **Enable enforcement:**
   Change `enforcementAction: deny` after validation period.

#### Related Resources

- [OPA Templates Overview](../index.md)
- [OPA RBAC Policies](../rbac/index.md)
- [Kyverno Pod Security Templates](../pod-security/index.md)

### OPA Policy Templates

> **Deploy in Audit Mode First**
>
> Use `enforcementAction: dryrun` initially. Existing resources may violate constraints. Monitor violations for 48 hours using `kubectl get constraints`, fix non-compliant resources, then switch to `deny`.
>

Production-ready OPA/Gatekeeper constraint templates for Kubernetes admission control. **20 policies** covering pod security, image
validation, RBAC, and resource governance. Each template includes complete Rego implementation, constraint examples, customization
options, validation commands, and real-world use cases.

---

#### Available Templates

##### Pod Security (5 Policies)

Prevent privileged containers, block host namespace access, enforce capability drops, require secure contexts, and prevent privilege escalation.

**Files:**

- **[Privileged & Host Namespaces →](pod-security/overview.md)** (2 policies)
- **[Capabilities Drop →](pod-security/capabilities.md)** (1 policy)
- **[Security Contexts →](pod-security/contexts.md)** (1 policy)
- **[Privilege Escalation Prevention →](pod-security/escalation.md)** (1 policy)

Key policies:

- Privileged Container Prevention (block `privileged: true`)
- Host Namespace Restrictions (block hostNetwork, hostPID, hostIPC, hostPort)
- Required Capabilities Drop (enforce `drop: ["ALL"]`, restrict dangerous capabilities)
- Security Context Requirements (runAsNonRoot, readOnlyRootFilesystem, UID restrictions)
- Privilege Escalation Prevention (block `allowPrivilegeEscalation: true`)

---

##### Image Security (5 Policies)

Control container images with registry allowlists, tag requirements, digest enforcement, signature verification annotations, and base image governance.

**Files:**

- **[Registry & Tag Validation →](image/security.md)** (2 policies)
- **[Digest Enforcement →](image/digest.md)** (1 policy)
- **[Signature Verification →](image/verification.md)** (1 policy)
- **[Base Image Enforcement →](image/base.md)** (1 policy)

Key policies:

- Registry Allowlist (enforce approved registries, block public Docker Hub)
- Tag Requirements (block `latest` tags, require specific tag patterns)
- Digest Enforcement (require SHA256 digest references, block tag-only images)
- Image Signature Verification Annotations (require proof of cosign verification in CI/CD)
- Base Image Enforcement (require approved base images via annotations, block deprecated)

---

##### RBAC (5 Policies)

Restrict service accounts, prevent cross-namespace role bindings, block cluster-admin assignments, restrict privileged verbs, and prevent wildcard permissions.

**Files:**

- **[Service Accounts & Role Bindings →](rbac/overview.md)** (2 policies)
- **[Cluster-Admin Prevention →](rbac/cluster-admin.md)** (1 policy)
- **[Privileged Verbs Restrictions →](rbac/privileged-verbs.md)** (1 policy)
- **[Wildcard Prevention →](rbac/wildcards.md)** (1 policy)

Key policies:

- Service Account Restrictions (block default SA usage, require dedicated SAs, prevent auto-mount tokens)
- Role Binding Namespace Enforcement (prevent cross-namespace subjects in RoleBindings)
- Cluster-Admin Prevention (block cluster-admin and system:masters role assignments)
- Privileged Verbs Restrictions (block escalate, impersonate, bind verbs)
- Wildcard Resource Prevention (block `*` in resources, apiGroups, verbs)

---

##### Resource Governance (5 Policies)

Enforce resource limits and requests, require namespace quotas, mandate LimitRanges, control ephemeral storage, and restrict storage classes and PVC sizes.

**Files:**

- **[Resource Limits & Quotas →](resource/governance.md)** (2 policies)
- **[LimitRange & Ephemeral Storage →](resource/limitrange.md)** (2 policies)
- **[Storage Class & PVC Constraints →](resource/storage.md)** (2 policies)

Key policies:

- Resource Limits and Requests Enforcement (CPU, memory limits required, max limits enforcement)
- Resource Quota Requirements (namespace quotas required, prevent unbounded consumption)
- LimitRange Requirements (default limits in namespaces, prevent extreme requests)
- Ephemeral Storage Limits (ephemeral-storage limits required, prevent disk exhaustion)
- Storage Class Restrictions (allowlist/blocklist, cost control, migration enforcement)
- PVC Size Constraints (min/max sizes, approval workflow for large volumes)

---

#### OPA vs Kyverno

Choosing between OPA/Gatekeeper and Kyverno depends on your team's expertise and requirements:

##### Use OPA/Gatekeeper When

- You need **maximum flexibility** in policy logic (Rego is Turing-complete)
- Your team has **Rego expertise** or investment in OPA across multiple systems
- You require **cross-platform policy** (Kubernetes, Terraform, Envoy, etc.)
- Policies involve **complex conditional logic** or multi-resource validation
- You're building a **policy platform** for enterprise governance

##### Use Kyverno When

- You want **Kubernetes-native YAML** policies (no DSL learning curve)
- You need **mutation and generation** features (OPA is validation-only)
- Your team prefers **JMESPath** over Rego for data extraction
- You want **faster time-to-value** with simpler policies
- You're **new to policy-as-code** and want quick adoption

**See [Decision Guide →](../decision-guide.md)** for detailed comparison and migration strategies.

---

#### Policy Deployment

All OPA constraint templates follow the same two-step deployment pattern:

##### Step 1: Deploy ConstraintTemplate

The `ConstraintTemplate` defines the policy logic in Rego:

```yaml
apiVersion: templates.gatekeeper.sh/v1beta1
kind: ConstraintTemplate
metadata:
  name: k8sblockprivileged
spec:
  crd:
    spec:
      names:
        kind: K8sBlockPrivileged
  targets:
    - target: admission.k8s.gatekeeper.sh
      rego: |
        package k8sblockprivileged
        violation[{"msg": msg}] {
##          # Rego policy logic here
        }
```

##### Step 2: Deploy Constraint

The `Constraint` activates the template with specific parameters:

```yaml
apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sBlockPrivileged
metadata:
  name: block-privileged-containers
spec:
  enforcementAction: dryrun  # Use 'deny' for enforcement
  match:
    kinds:
      - apiGroups: [""]
        kinds: ["Pod"]
    namespaces:
      - "production"
      - "staging"
```

---

#### Quick Start

Standard deployment workflow for all templates:

```bash
### 1. Deploy constraint template (defines policy logic)
kubectl apply -f constraint-template.yaml

### 2. Deploy constraint in audit mode (dryrun)
kubectl apply -f constraint.yaml

### 3. Monitor violations
kubectl get constraints
kubectl get <constraint-kind> <constraint-name> -o yaml

### 4. Check audit results
kubectl get constraints -o json | jq '.items[].status.violations'

### 5. Fix non-compliant resources
kubectl get pods -n production --show-labels

### 6. Switch to enforcement mode after validation
kubectl patch <constraint-kind> <constraint-name> \
  --type merge \
  -p '{"spec":{"enforcementAction":"deny"}}'
```

---

#### Policy Customization

Every template includes a customization table with these common parameters:

| Parameter | Default | Purpose |
|-----------|---------|---------|
| `enforcementAction` | `dryrun` | Use `dryrun` for testing, `deny` for enforcement |
| `match.kinds` | Varies | Target specific Kubernetes resource types |
| `match.namespaces` | `[]` | Target specific namespaces (empty = all) |
| `match.excludedNamespaces` | `["kube-system"]` | Exempt system namespaces |
| `match.labelSelector` | None | Target resources with specific labels |

Template-specific parameters (e.g., `exemptImages`, `allowedRegistries`, `maxCPU`) are documented in each policy's customization table.

---

#### Constraint Status

Monitor policy violations and audit results:

```bash
### List all constraints
kubectl get constraints

### Get detailed status for a specific constraint
kubectl get k8sblockprivileged block-privileged-containers -o yaml

### Extract violations from constraint status
kubectl get k8sblockprivileged block-privileged-containers \
  -o jsonpath='{.status.violations[*].message}' | jq

### Count total violations across all constraints
kubectl get constraints -o json | \
  jq '[.items[].status.totalViolations] | add'
```

---

#### Rego Testing

All templates include unit test examples for Rego policies:

```bash
### Install OPA CLI
brew install opa  # macOS
### or download from https://www.openpolicyagent.org/docs/latest/#running-opa

### Test Rego policy locally
opa test constraint-template.yaml test-cases.yaml -v

### Example test case
### test-cases.yaml
package k8sblockprivileged

test_privileged_container_blocked {
  violation[{"msg": msg}] with input as {
    "review": {
      "object": {
        "spec": {
          "containers": [{
            "name": "test",
            "securityContext": {"privileged": true}
          }]
        }
      }
    }
  }
}
```

See **[Privilege Escalation Prevention →](pod-security/escalation.md#rego_unit_testing)** for complete testing guide.

---

#### Related Resources

- **[Kyverno Templates →](../kyverno/index.md)** - 28 Kyverno policies for comparison
- **[Decision Guide →](../decision-guide.md)** - OPA vs Kyverno selection criteria
- **[OPA/Kyverno Comparison →](../opa-kyverno-comparison.md)** - Detailed feature comparison
- **[Migration Guide →](../opa-kyverno-migration.md)** - Switching between OPA and Kyverno
- **[Template Library Overview →](index.md)** - Back to main page

### OPA RBAC Templates

RBAC policies control who can perform which actions on which resources. These templates prevent privilege escalation through overly permissive roles.

> **Wildcards Grant Unrestricted Access**
>
> RBAC rules with `resources: ["*"]` or `verbs: ["*"]` grant access to all current and future resources or actions. Avoid wildcards except for break-glass admin roles.
>

#### Why RBAC Governance Matters

Kubernetes RBAC is additive (permissions accumulate across bindings). Without enforcement:

- **Privilege Escalation** - Users create ClusterRoleBindings granting themselves cluster-admin
- **Wildcard Abuse** - Roles grant `*` permissions instead of least privilege
- **Dangerous Verbs** - `escalate`, `impersonate`, `bind` allow RBAC manipulation
- **Permanent Admin Access** - ClusterRole `cluster-admin` granted to service accounts

#### Available Templates

##### [Cluster-Admin Prevention](cluster-admin.md)

Block cluster-admin role bindings except for approved subjects:

- Prevent creation of ClusterRoleBindings to `cluster-admin`
- Allow only break-glass admin accounts or system components
- Validate subject identity before granting cluster-admin
- Audit cluster-admin grants for compliance

**Apply a policy:**

```bash
kubectl apply -f cluster-admin.yaml
```

##### [Privileged Verb Restrictions](privileged-verbs.md)

Block dangerous RBAC verbs that enable privilege escalation:

- Prevent `escalate` verb (bypass RBAC validation)
- Block `bind` verb (assign ClusterRoles to arbitrary subjects)
- Restrict `impersonate` verb (act as other users/service accounts)
- Limit `*` verb grants to approved roles

**Apply a policy:**

```bash
kubectl apply -f privileged-verbs.yaml
```

##### [Wildcard Prevention](wildcards.md)

Restrict wildcard usage in RBAC rules:

- Block `resources: ["*"]` in Role/ClusterRole rules
- Prevent `verbs: ["*"]` except for read-only access
- Require explicit resource and verb lists
- Allow wildcards only for monitoring/observability roles

**Apply a policy:**

```bash
kubectl apply -f wildcards.yaml
```

##### [RBAC Policy Overview](overview.md)

General RBAC governance and least privilege principles:

- Namespace-scoped roles preferred over ClusterRoles
- Service account permissions limited to pod requirements
- Time-bounded RoleBindings with expiration annotations
- Regular RBAC audits and privilege reviews

**Apply a policy:**

```bash
kubectl apply -f overview.yaml
```

#### RBAC Security Patterns

##### Least Privilege Principle

Grant minimum permissions required for each workload:

1. **Start with no permissions** - Service accounts have no default permissions
2. **Add specific resources** - `pods`, `configmaps`, not `*`
3. **Add specific verbs** - `get`, `list`, not `*`
4. **Scope to namespace** - Use Role instead of ClusterRole when possible

##### Defense Against Privilege Escalation

Block RBAC manipulation verbs:

- **`escalate`** - Allows creating roles with more permissions than creator has
- **`bind`** - Allows granting roles to arbitrary subjects
- **`impersonate`** - Allows acting as other users without authentication

Only cluster admins should have these verbs.

##### Time-Bounded Permissions

Use annotations to enforce temporary access:

```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: temporary-debug-access
  annotations:
    rbac.expires: "2025-01-05T00:00:00Z"
subjects:
  - kind: User
    name: engineer@company.com
roleRef:
  kind: ClusterRole
  name: debug-read-only
```

OPA policies can validate expiration and block expired bindings.

#### Common Enforcement Scenarios

##### Scenario 1: Prevent Unauthorized Cluster-Admin

Block cluster-admin except for approved break-glass accounts:

```yaml
### Enforced by: cluster-admin.yaml
### Result: Only subjects in approved list can receive cluster-admin binding
### Impact: Prevents privilege escalation to cluster-admin
```

##### Scenario 2: Block Dangerous RBAC Verbs

Prevent use of `escalate`, `bind`, `impersonate`:

```yaml
### Enforced by: privileged-verbs.yaml
### Result: Roles cannot include escalate/bind/impersonate verbs
### Impact: Prevents users from granting themselves additional permissions
```

##### Scenario 3: Eliminate Wildcard Permissions

Require explicit resource and verb lists:

```yaml
### Enforced by: wildcards.yaml
### Result: Roles must specify resources: ["pods"], not resources: ["*"]
### Impact: Reduces blast radius of compromised service accounts
```

#### Testing RBAC Policies

Validate enforcement without disrupting operations:

```bash
### Test cluster-admin prevention (should fail for unapproved subject)
kubectl create clusterrolebinding test-admin \
  --clusterrole=cluster-admin \
  --user=attacker@example.com
### Expected: Admission denied by cluster-admin.yaml

### Test privileged verb block (should fail with escalate verb)
cat <<EOF | kubectl apply -f -
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: escalate-test
rules:
  - apiGroups: ["rbac.authorization.k8s.io"]
    resources: ["clusterroles"]
    verbs: ["escalate"]
EOF
### Expected: Admission denied by privileged-verbs.yaml

### Test wildcard prevention (should fail with resources: ["*"])
cat <<EOF | kubectl apply -f -
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: wildcard-test
  namespace: default
rules:
  - apiGroups: ["*"]
    resources: ["*"]
    verbs: ["get", "list"]
EOF
### Expected: Admission denied by wildcards.yaml

### Test compliant role (should succeed)
cat <<EOF | kubectl apply -f -
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: compliant-role
  namespace: default
rules:
  - apiGroups: [""]
    resources: ["pods", "configmaps"]
    verbs: ["get", "list", "watch"]
EOF
### Expected: Admission allowed by all policies
```

#### RBAC Audit and Review

Regularly audit RBAC configuration for compliance:

##### Identify Privileged Bindings

```bash
### List all ClusterRoleBindings to cluster-admin
kubectl get clusterrolebindings -o json | \
  jq '.items[] | select(.roleRef.name == "cluster-admin") | .metadata.name'

### List roles with wildcard permissions
kubectl get roles,clusterroles --all-namespaces -o json | \
  jq '.items[] | select(.rules[].resources[] == "*") | .metadata.name'
```

##### Find Service Accounts with Excessive Permissions

```bash
### List service accounts with cluster-admin
kubectl get clusterrolebindings -o json | \
  jq '.items[] | select(.roleRef.name == "cluster-admin" and .subjects[].kind == "ServiceAccount")'

### Find service accounts with escalate/bind/impersonate verbs
kubectl get roles,clusterroles --all-namespaces -o json | \
  jq '.items[] | select(.rules[].verbs[] | IN("escalate", "bind", "impersonate"))'
```

##### Validate RoleBinding Subjects

```bash
### List all human users with cluster-level access
kubectl get clusterrolebindings -o json | \
  jq '.items[] | select(.subjects[].kind == "User") | {binding: .metadata.name, user: .subjects[].name, role: .roleRef.name}'
```

#### Break-Glass Admin Access

Maintain emergency access while enforcing policies:

##### Approved Admin Accounts

Define break-glass accounts in OPA policy:

```rego
### Pseudo-code: Full implementation in cluster-admin.yaml
approved_admins := {
  "break-glass-admin@company.com",
  "oncall-sre@company.com",
  "system:masters",  # For kubeadm bootstrap
}

deny[msg] {
  input.kind == "ClusterRoleBinding"
  input.roleRef.name == "cluster-admin"
  not approved_admins[input.subjects[_].name]
  msg := "cluster-admin can only be granted to approved break-glass accounts"
}
```

##### Temporary Elevation

Use short-lived credentials instead of permanent cluster-admin:

```bash
### Generate temporary kubeconfig with cluster-admin (expires in 1 hour)
kubectl create token break-glass-admin --duration=1h

### Use temporary token for emergency operations
kubectl --token=$(kubectl create token break-glass-admin --duration=1h) get nodes
```

#### Related Resources

- [OPA Templates Overview](../index.md)
- [OPA Pod Security](../pod-security/index.md)
- [OPA Resource Governance](../resource/index.md)

### OPA Resource Governance Templates

Resource governance policies prevent overconsumption, enforce quotas, and validate LimitRange compliance across your cluster.

> **ResourceQuota vs LimitRange vs OPA**
>
> ResourceQuota caps total namespace consumption. LimitRange sets defaults and bounds for individual pods. OPA validates configuration before admission. Use all three for comprehensive governance.
>

#### Why Resource Governance Matters

Kubernetes does not enforce resource limits by default. This creates operational and financial risks:

- **Node Exhaustion** - Unbounded pods consume all CPU/memory
- **Quota Violations** - Deployments exceed namespace ResourceQuota
- **Cost Overruns** - Excessive storage provisioning or oversized VMs
- **Cluster Instability** - OOMKilled pods cascade across nodes

#### Available Templates

##### [Resource Governance](governance.md)

Enforce resource limits, requests, and quota compliance:

- Require resource requests for all containers
- Mandate resource limits to prevent node exhaustion
- Validate requests ≤ limits for CPU and memory
- Block workloads exceeding namespace quotas

**Apply a policy:**

```bash
kubectl apply -f governance.yaml

```

##### [LimitRange Validation](limitrange.md)

Enforce LimitRange compliance for pods and containers:

- Validate pod resource requests against LimitRange defaults
- Block pods exceeding LimitRange max values
- Require LimitRange in all non-system namespaces
- Enforce container-level and pod-level limits

**Apply a policy:**

```bash
kubectl apply -f limitrange.yaml

```

##### [Storage Constraints](storage.md)

Control PersistentVolume and PersistentVolumeClaim allocation:

- Restrict PVC sizes based on namespace or StorageClass
- Require approved StorageClasses for production data
- Block dynamic provisioning in restricted namespaces
- Validate volume access modes and reclaim policies

**Apply a policy:**

```bash
kubectl apply -f storage.yaml

```

#### Resource Governance Patterns

##### Three-Layer Resource Control

Implement overlapping controls for comprehensive governance:

1. **OPA Policies** - Validate resource configuration at admission time
2. **LimitRange** - Set namespace-level defaults and max values
3. **ResourceQuota** - Cap total namespace consumption

##### Right-Sizing Workloads

Balance cost and reliability with appropriate resource values:

| Resource | Too Low | Too High | Sweet Spot |
|----------|---------|----------|------------|
| **Requests** | OOMKilled pods | Wasted capacity | Actual usage (P95) |
| **Limits** | Throttling | Noisy neighbors | 2x requests |
| **Storage** | Out of space | High costs | Actual data + 30% |

Use Vertical Pod Autoscaler (VPA) to identify optimal values.

##### Quota Enforcement Strategy

Define quotas based on team size and workload type:

```yaml
### Development namespace quota
apiVersion: v1
kind: ResourceQuota
metadata:
  name: dev-quota
  namespace: dev-team
spec:
  hard:
    requests.cpu: "10"
    requests.memory: "20Gi"
    limits.cpu: "20"
    limits.memory: "40Gi"
    persistentvolumeclaims: "10"
    requests.storage: "100Gi"

```

OPA policies validate workloads against these quotas before admission.

#### Common Enforcement Scenarios

##### Scenario 1: Prevent Unbounded Resource Consumption

Require resource limits on all containers:

```yaml
### Enforced by: governance.yaml
### Result: All containers must define resources.limits.cpu and resources.limits.memory
### Impact: Prevents single pod from consuming entire node capacity

```

##### Scenario 2: Enforce LimitRange Compliance

Block pods exceeding namespace LimitRange maximums:

```yaml
### Enforced by: limitrange.yaml
### Result: Pods cannot request more CPU/memory than LimitRange allows
### Impact: Ensures fair resource distribution across namespace

```

##### Scenario 3: Control Storage Costs

Restrict PVC sizes based on environment:

```yaml
### Enforced by: storage.yaml
### Result: PVCs in dev-* namespaces cannot exceed 50Gi
### Impact: Prevents accidental provisioning of expensive storage volumes

```

#### Testing Resource Governance Policies

Validate enforcement without blocking legitimate workloads:

```bash
### Test resource limit requirement (should fail without limits)
kubectl run no-limits --image=nginx
### Expected: Admission denied by governance.yaml

### Test excessive resource request (should fail above quota)
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Pod
metadata:
  name: excessive-request
  namespace: dev-team
spec:
  containers:
    - name: nginx
      image: nginx
      resources:
        requests:
          cpu: "50"  # Exceeds namespace quota of 10 CPU
          memory: "100Gi"
EOF
### Expected: Admission denied by governance.yaml (quota violation)

### Test LimitRange violation (should fail above max)
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Pod
metadata:
  name: limitrange-violation
  namespace: dev-team
spec:
  containers:
    - name: nginx
      image: nginx
      resources:
        requests:
          cpu: "100"  # Exceeds LimitRange max of 2 CPU
          memory: "200Gi"
EOF
### Expected: Admission denied by limitrange.yaml

### Test storage size restriction (should fail for excessive PVC)
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: large-pvc
  namespace: dev-team
spec:
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 500Gi  # Exceeds policy max of 50Gi for dev namespaces
EOF
### Expected: Admission denied by storage.yaml

### Test compliant workload (should succeed)
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Pod
metadata:
  name: compliant-pod
  namespace: dev-team
spec:
  containers:
    - name: nginx
      image: nginx
      resources:
        requests:
          cpu: "100m"
          memory: "128Mi"
        limits:
          cpu: "500m"
          memory: "512Mi"
EOF
### Expected: Admission allowed by all policies

```

#### ResourceQuota and LimitRange Integration

OPA policies complement native Kubernetes resource controls:

##### ResourceQuota Example

```yaml
apiVersion: v1
kind: ResourceQuota
metadata:
  name: production-quota
  namespace: production
spec:
  hard:
    requests.cpu: "100"
    requests.memory: "200Gi"
    limits.cpu: "200"
    limits.memory: "400Gi"
    persistentvolumeclaims: "50"
    requests.storage: "1Ti"

```

##### LimitRange Example

```yaml
apiVersion: v1
kind: LimitRange
metadata:
  name: production-limits
  namespace: production
spec:
  limits:
    - max:
        cpu: "4"
        memory: "8Gi"
      min:
        cpu: "10m"
        memory: "64Mi"
      default:
        cpu: "500m"
        memory: "512Mi"
      defaultRequest:
        cpu: "100m"
        memory: "128Mi"
      type: Container

```

##### OPA Validation

OPA policies validate that:

- Pod resource requests/limits respect LimitRange bounds
- Total namespace consumption stays within ResourceQuota
- Requests ≤ Limits for all containers
- Required fields are present (not relying on LimitRange defaults)

#### Capacity Planning with Resource Policies

Use policy enforcement data for cluster sizing:

1. **Collect resource requests across namespaces:**

   ```bash
   kubectl get pods --all-namespaces -o json | \
     jq '[.items[].spec.containers[].resources.requests] | add'

   ```

2. **Compare against node capacity:**

   ```bash
   kubectl get nodes -o json | \
     jq '[.items[].status.allocatable] | add'

   ```

3. **Identify overcommitment:**

   ```bash
##   # If total requests > allocatable capacity, cluster is overcommitted
##   # Add nodes or reduce resource requests

   ```

4. **Adjust policies based on usage:**
   - Increase LimitRange maximums if legitimate workloads are blocked
   - Decrease quotas if namespace consumption is consistently low
   - Update OPA policies to reflect new capacity constraints

#### Cost Optimization with Resource Governance

Reduce cloud costs through policy enforcement:

##### CPU/Memory Optimization

- **Block oversized requests** - Prevent requesting 16 CPU for 100m usage
- **Enforce limits** - Prevent burst usage that triggers autoscaler
- **Right-size VPA recommendations** - Use VPA to identify bloated requests

##### Storage Optimization

- **Restrict PVC sizes** - Cap dev/test at 50Gi, prod at 500Gi
- **Require cheaper StorageClasses** - Use `standard` for non-critical data
- **Block dynamic provisioning** - Require pre-provisioned PVs for large databases

##### Autoscaling Optimization

- **Require HPA** - Scale replicas instead of oversizing pods
- **Set reasonable bounds** - Limit HPA max replicas to prevent cost spikes
- **Use cluster autoscaler** - Add nodes only when pending pods exist

#### Related Resources

- [OPA Templates Overview](../index.md)
- [OPA RBAC Policies](../rbac/index.md)
- [Kyverno Resource Governance](../resource/index.md)

### Policy Packaging

Build the policy-platform container that aggregates all policy sources and tools.

#### Overview

The policy-platform container is a multi-stage Docker build that:

1. Pulls policy repositories as OCI containers
2. Installs policy validation tools (Kyverno, Pluto, Spectral, Helm)
3. Aggregates everything into a single distributable image

> **One Container, All Policies**
>
> The policy-platform image runs identically in local dev, CI pipelines, and reference environments. Zero configuration drift.
>

---

#### Dockerfile Architecture

##### Multi-Stage Build Pattern

```dockerfile
### Stage 1-3: Pull policy repos as OCI containers
FROM policy-repo-1:tag AS policy_repo_1
FROM policy-repo-2:tag AS policy_repo_2
FROM policy-repo-3:tag AS policy_repo_3

### Final stage: Aggregate and install tools
FROM alpine:3.24.1

### Install tools
RUN apk add curl bash helm yq

### Install Kyverno CLI
RUN curl -sSL ...kyverno.tar.gz | tar -xz

### Copy all policy repos
COPY --from=policy_repo_1 /repos/repo1/ /repos/repo1/
COPY --from=policy_repo_2 /repos/repo2/ /repos/repo2/
COPY --from=policy_repo_3 /repos/repo3/ /repos/repo3/
```

**Key Benefits**:

- Policy repos versioned independently
- Single container aggregates all
- Tools bundled with policies

---

#### Complete Dockerfile

See [Multi-Source Policies](../multi-source-policies/index.md) for full Dockerfile example.

**Core components**:

1. **Base image**: Alpine Linux (small, secure)
2. **Tools**: Kyverno CLI v1.13.2, Pluto v5.21.1, Spectral, Helm, yq
3. **Policy repos**: Copied from dependent OCI containers

---

#### Tool Selection

##### Kyverno CLI

**Purpose**: Policy validation and testing

**Installation**:

```dockerfile
RUN curl -sSL https://github.com/kyverno/kyverno/releases/download/v1.13.2/kyverno-cli_v1.13.2_linux_x86_64.tar.gz \
  | tar -xz -C /usr/local/bin
```

##### Pluto

**Purpose**: Deprecated API detection

**Installation**:

```dockerfile
RUN curl -sSL https://github.com/FairwindsOps/pluto/releases/download/v5.21.1/pluto_5.21.1_linux_amd64.tar.gz \
  | tar xz -C /usr/local/bin
```

##### Spectral

**Purpose**: Schema validation

**Installation**:

```dockerfile
RUN latest_spectral=$(curl -sSL https://api.github.com/repos/stoplightio/spectral/releases/latest | grep 'tag_name' | cut -d\" -f4) && \
    curl -sSL https://github.com/stoplightio/spectral/releases/download/${latest_spectral}/spectral-alpine-x64 \
      -o /usr/local/bin/spectral && \
    chmod +x /usr/local/bin/spectral
```

> **Pin Tool Versions**
>
> Always pin specific tool versions in Dockerfile. Dynamic `latest` tags cause non-reproducible builds.
>

---

#### Build Process

##### Local Build

```bash
docker build -t policy-platform:latest -f ci/Dockerfile .
```

##### CI Build

```yaml
### Bitbucket Pipelines
- step:
    name: Build Policy Platform
    services:
      - docker
    script:
      - docker build -t policy-platform:${BITBUCKET_BUILD_NUMBER} -f ci/Dockerfile .
      - docker tag policy-platform:${BITBUCKET_BUILD_NUMBER} policy-platform:latest
      - docker push policy-platform:${BITBUCKET_BUILD_NUMBER}
      - docker push policy-platform:latest
```

##### GitHub Actions

```yaml
- name: Build and Push
  uses: docker/build-push-action@v5
  with:
    context: .
    file: ./ci/Dockerfile
    push: true
    tags: |
      policy-platform:${{ github.sha }}
      policy-platform:latest
```

---

#### Next Steps

- **[Distribution](distribution.md)** - Versioning, testing, optimization
- **[Maintenance](maintenance.md)** - Troubleshooting and best practices
- **[Multi-Source Policies](../multi-source-policies/index.md)** - Policy aggregation patterns

### Policy-as-Code Operations

Day-to-day management, updates, and monitoring of policy enforcement.

#### Overview

Operating a policy-as-code platform requires:

1. **Policy updates** - Rolling out new policies safely
2. **Monitoring** - Tracking compliance and violations
3. **Exception management** - Handling edge cases
4. **Troubleshooting** - Resolving policy issues
5. **Auditing** - Proving compliance

> **Operations at Scale**
>
> Policy-as-code operations follow GitOps principles. All changes go through Git. All deployments are tracked. All violations are logged.
>

---

#### Policy Lifecycle

##### Adding New Policies

**Step 1**: Create policy in policy repo

```yaml
### security-policy/charts/security-policy/templates/require-network-policy.yaml
apiVersion: kyverno.io/v1
kind: ClusterPolicy
metadata:
  name: require-network-policy
  annotations:
    policies.kyverno.io/title: Require NetworkPolicy
    policies.kyverno.io/category: Security
    policies.kyverno.io/severity: high
spec:
  validationFailureAction: {{ .Values.policies.networkPolicy.validationFailureAction }}
  background: true
  rules:
    - name: check-network-policy-exists
      match:
        resources:
          kinds:
            - Namespace
      validate:
        message: "Namespace must have a NetworkPolicy"
        deny:
          conditions:
            - key: "{{ request.object.metadata.name }}"
              operator: AnyNotIn
              value: [" kube-system", "kube-public", "kube-node-lease"]
```

**Step 2**: Add to values.yaml

```yaml
### security-policy/charts/security-policy/values.yaml
policies:
  networkPolicy:
    enabled: true
    validationFailureAction: Audit  # Start with Audit
```

**Step 3**: Test locally

```bash
docker run --rm -v $(pwd):/workspace policy-platform:latest bash -c '\
  helm template security /repos/security-policy/charts/security-policy \
    -f /repos/security-policy/charts/security-policy/values.yaml \
  > /tmp/policies.yaml &&\
  kyverno apply /tmp/policies.yaml --resource /workspace/test-namespace.yaml\
'
```

**Step 4**: Deploy to dev

```bash
### Update policy-platform container (rebuild with new policy)
docker build -t policy-platform:v1.0.3 -f ci/Dockerfile .
docker push policy-platform:v1.0.3

### Deploy to dev cluster
helm upgrade security-policy /repos/security-policy/charts/security-policy \
  --namespace kyverno \
  --values /repos/security-policy/cd/dev/values.yaml
```

**Step 5**: Monitor PolicyReports

```bash
kubectl get policyreport -A

### Check for violations
kubectl get policyreport polr-ns-default -o yaml
```

**Step 6**: Switch to Enforce after validation

```yaml
### security-policy/cd/prd/values.yaml
policies:
  networkPolicy:
    validationFailureAction: Enforce  # Now block violations
```

> **Always Start with Audit**
>
> New policies must start in Audit mode. Monitor violations for at least one week before switching to Enforce. This prevents breaking existing workloads.
>

---

#### Updating Existing Policies

##### Policy Refinement

Refine policy based on violations:

**Original policy** (too strict):

```yaml
validate:
  pattern:
    spec:
      containers:
        - resources:
            limits:
              memory: "?*"
              cpu: "?*"
            requests:              # Requires both limits AND requests
              memory: "?*"
              cpu: "?*"
```

**Updated policy** (allow limits-only):

```yaml
validate:
  pattern:
    spec:
      containers:
        - resources:
            limits:
              memory: "?*"
              cpu: "?*"
##          # Requests optional
```

**Deployment**:

1. Update policy in repo
2. Increment version (`v2.1.2` → `v2.1.3`)
3. Rebuild policy-platform container
4. Deploy to clusters

---

#### Next Steps

- **[Monitoring](monitoring.md)** - Compliance dashboards, metrics, and alerting
- **[Workflows](workflows.md)** - Updates, backup, performance tuning
- **[Runtime Deployment](../runtime-deployment/index.md)** - Kyverno deployment guide

### Policy-as-Code Template Library

**48 production-ready policies** for Kubernetes security and governance. Reduce the Rego learning curve. Copy, customize, deploy.

<!-- more -->

> **Template Library Overview**
>
> This library contains **28 Kyverno policies** and **20 OPA/Gatekeeper constraint templates** covering pod security, image validation, RBAC, resource governance, network security, mutation, and generation.
> Each template includes complete YAML/Rego, customization variables, validation commands, and real-world use cases.
>

---

#### What You Get

This library provides ready-to-use policies for common security scenarios:

- **48 Total Policies**: 28 Kyverno + 20 OPA/Gatekeeper
- **Complete Implementation**: Full YAML/Rego with production-ready configuration
- **Customization Tables**: Variables, defaults, and purpose for each parameter
- **Validation Commands**: Test policies before enforcement
- **Real-World Use Cases**: 4-6 production scenarios per policy
- **Testing Guidance**: Audit mode, policy reports, troubleshooting

---

#### Template Categories

##### [Decision Guide →](decision-guide.md)

**Choose between OPA and Kyverno** based on team expertise, policy complexity, and operational requirements.

- Quick decision matrix (expertise, scope, complexity)
- Recommended starter paths
- [Detailed comparison →](opa-kyverno-comparison.md)
- [Migration strategies →](opa-kyverno-migration.md)

---

##### [Kyverno Templates →](kyverno/index.md)

**28 production-ready Kyverno policies** for Kubernetes admission control, mutation, and resource generation.

###### [Pod Security →](kyverno/pod-security/index.md) (5 Policies)

- Pod Security Standards Enforcement
- Host Namespace Restrictions
- Privilege Escalation Prevention
- Seccomp Profile Enforcement
- AppArmor Profile Requirements

###### [Image Validation →](kyverno/image/index.md) (5 Policies)

- Image Digest Requirements
- Registry Allowlist and Tag Validation
- Cosign Image Signature Verification
- Base Image Enforcement
- CVE Scanning Gates

###### [Resource Management →](kyverno/resource/index.md) (5 Policies)

- Resource Limits and Requests Enforcement
- CPU and Memory Ratio Enforcement
- Ephemeral Storage Limits
- PVC Size Constraints
- HPA Configuration Requirements

###### [Network Security →](kyverno/network/index.md) (5 Policies)

- Require Network Policies
- Egress Restrictions
- Ingress Class Requirements
- Ingress TLS Requirements
- Service Type Restrictions

###### Mutation & Generation (7 Policies)

- [Mutation Policies →](kyverno/mutation/index.md) - Default Label Injection, Namespace Label Propagation, Logging Sidecar Injection, Monitoring Sidecar Injection
- [Generation Policies →](kyverno/generation/index.md) - Automatic ResourceQuota Generation, Default-Deny NetworkPolicy Generation, Automatic PodDisruptionBudget Generation

###### [Labels & Metadata →](kyverno/labels.md) (1 Policy)

- Mandatory Labels and Annotations

---

##### [OPA/Gatekeeper Templates →](opa/index.md)

**20 production-ready OPA constraint templates** with complete Rego implementation for advanced policy enforcement.

###### [Pod Security →](opa/pod-security/index.md) (5 Policies)

- Privileged Container Prevention
- Host Namespace Restrictions
- Required Capabilities Drop
- Security Context Requirements
- Privilege Escalation Prevention

###### [Image Security →](opa/image/index.md) (5 Policies)

- Registry Allowlist
- Tag Requirements
- Digest Enforcement
- Image Signature Verification Annotations
- Base Image Enforcement

###### [RBAC →](opa/rbac/index.md) (5 Policies)

- Service Account Restrictions
- Role Binding Namespace Enforcement
- Cluster-Admin Prevention
- Privileged Verbs Restrictions
- Wildcard Resource Prevention

###### [Resource Governance →](opa/resource/index.md) (5 Policies)

- Resource Limits and Requests Enforcement
- Resource Quota Requirements
- LimitRange Requirements
- Ephemeral Storage Limits
- Storage Class Restrictions

---

##### [JMESPath Patterns →](jmespath/index.md)

**Advanced Kyverno pattern library** for complex validation logic using JMESPath.

- Pattern fundamentals (projection, filtering, multi-select)
- Cross-field validation (requests vs limits, label dependencies)
- Complex conditions (nested logic, transformations)
- [Advanced patterns →](jmespath/advanced.md) (aggregation, arithmetic, string manipulation)
- [Enterprise examples →](jmespath/enterprise.md) (registry policies, cost controls, HA requirements)
- [Testing guide →](jmespath/testing.md) (kyverno jp CLI, debugging, validation)

---

##### [CI/CD Integration →](ci-cd-integration.md)

Automated policy validation in development pipelines:

- GitHub Actions pre-flight validation
- ArgoCD policy gating
- Pre-commit hooks

---

##### [Usage Guide →](usage-guide.md)

Template customization workflow, validation best practices, and quick start guides:

- Customization workflow
- Validation best practices
- Quick start guides
- Troubleshooting

---

#### Policy Engine Comparison

Choose the right policy engine for your team:

| Feature | Kyverno | OPA/Gatekeeper |
|---------|---------|----------------|
| **Policies** | 28 (validation, mutation, generation) | 20 (validation only) |
| **Language** | YAML + JMESPath | Rego (Go-like DSL) |
| **Learning Curve** | < 1 hour | 4-8 hours |
| **Best For** | Kubernetes-native teams, fast adoption | Multi-platform policies, complex logic |
| **Mutation** | ✅ Native support | ❌ Validation only |
| **Generation** | ✅ Auto-create resources | ❌ Validation only |

**See [Decision Guide →](decision-guide.md)** for detailed comparison and recommended starter paths.

---

#### Quick Start

> **Deploy in Audit Mode First**
>
> Always start with `audit` (Kyverno) or `dryrun` (OPA) mode. Monitor violations for 48 hours before switching to enforcement. Existing workloads may violate policies.
>

##### Kyverno Quick Start (5 minutes)

```bash
### 1. Install Kyverno
helm repo add kyverno https://kyverno.github.io/kyverno/
helm install kyverno kyverno/kyverno --namespace kyverno --create-namespace

### 2. Apply a policy (starts in audit mode)
kubectl apply -f https://raw.githubusercontent.com/adaptive-enforcement-lab/docs/main/kyverno-pod-security.yaml

### 3. Monitor violations
kubectl get polr -A  # PolicyReports
kubectl get cpolr    # ClusterPolicyReports

### 4. Switch to enforcement after validation
kubectl patch clusterpolicy require-pod-security \
  --type merge \
  -p '{"spec":{"validationFailureAction":"enforce"}}'
```

##### OPA/Gatekeeper Quick Start (10 minutes)

```bash
### 1. Install Gatekeeper
kubectl apply -f https://raw.githubusercontent.com/open-policy-agent/gatekeeper/master/deploy/gatekeeper.yaml

### 2. Deploy constraint template (policy logic)
kubectl apply -f https://raw.githubusercontent.com/adaptive-enforcement-lab/docs/main/opa-pod-security.yaml

### 3. Deploy constraint (starts in dryrun mode)
kubectl apply -f constraint.yaml

### 4. Monitor violations
kubectl get constraints
kubectl get k8sblockprivileged -o yaml

### 5. Switch to enforcement after validation
kubectl patch k8sblockprivileged block-privileged \
  --type merge \
  -p '{"spec":{"enforcementAction":"deny"}}'
```

---

#### Related Resources

- **[Kyverno Official Documentation](https://kyverno.io/docs/)** - Kyverno guides and API reference
- **[OPA/Gatekeeper Documentation](https://open-policy-agent.org/docs/latest/kubernetes-admission-control/)** - Gatekeeper deployment and Rego reference
- **[Kubernetes Pod Security Standards](https://kubernetes.io/docs/concepts/security/pod-security-standards/)** - Baseline and Restricted profiles
- **[NIST SP 800-190](https://csrc.nist.gov/publications/detail/sp/800-190/final)** - Application Container Security Guide
- **[CIS Kubernetes Benchmark](https://www.cisecurity.org/benchmark/kubernetes)** - Security configuration standards

### Runtime Deployment: Admission Control with Kyverno

Deploy policy enforcement directly in Kubernetes clusters. Final safety net before production.

#### Overview

Runtime admission control enforces policies at the cluster level using Kyverno admission webhooks:

```mermaid
graph TD
    K[kubectl apply] --> API[Kubernetes API Server]
    API --> ADM[Admission Controller]
    ADM --> KYV[Kyverno Webhook]
    KYV --> POL{Policy Check}
    POL -->|Pass| ETCD[(etcd)]
    POL -->|Fail| REJECT[Reject Request]

    %% Ghostty Hardcore Theme
    style ETCD fill:#a7e22e,color:#1b1d1e
    style REJECT fill:#f92572,color:#1b1d1e

```

> **Runtime is the Final Safety Net**
>
> Local dev and CI checks can be bypassed. Runtime admission control is the last line of defense. If it fails, non-compliant resources never reach production.
>

---

#### Architecture Components

##### 1. Kyverno Admission Controller

Intercepts API requests before they reach etcd:

- **Admission Webhooks**: Validate, mutate, or generate resources
- **Background Scans**: Continuous compliance checking
- **Policy Reports**: Violation tracking

##### 2. Policy Reporter

Visualization and notification layer:

- **Dashboard**: Policy compliance overview
- **Metrics**: Prometheus integration
- **Alerts**: Slack, Teams, email notifications

##### 3. Policy Sources

Policies deployed as Kubernetes resources:

```yaml
apiVersion: kyverno.io/v1
kind: ClusterPolicy
metadata:
  name: require-resource-limits
spec:
  validationFailureAction: Enforce
  background: true
  rules:
    - name: check-cpu-memory
      match:
        resources:
          kinds:
            - Deployment
      validate:
        message: "CPU and memory limits required"
        pattern:
          spec:
            template:
              spec:
                containers:
                  - resources:
                      limits:
                        memory: "?*"
                        cpu: "?*"
```

---

#### Installation

##### Kyverno Deployment

Deploy Kyverno using Helm:

```bash
helm repo add kyverno https://kyverno.github.io/kyverno/
helm repo update

helm install kyverno kyverno/kyverno \
  --namespace kyverno \
  --create-namespace \
  --values kyverno-values.yaml
```

**kyverno-values.yaml**:

```yaml
features:
  logging:
    logLevel: -2  # Info level

  backgroundScan:
    backgroundScanInterval: 6h

### Clean up old reports
policyReportsCleanup:
  enabled: true

cleanupJobs:
  admissionReports:
    enabled: true
    schedule: "0 0 * * *"  # Daily

  clusterAdmissionReports:
    enabled: true
    schedule: "0 0 * * SUN"  # Weekly

  policyReports:
    enabled: true
    schedule: "0 0 * * *"  # Daily

  clusterPolicyReports:
    enabled: true
    schedule: "0 0 1 * *"  # Monthly

  resources:
    limits:
      memory: 128Mi
    requests:
      cpu: 50m
      memory: 64Mi

### Exclude system namespaces
resourceFilters:
  resourceFiltersExcludeNamespaces:
    - kube-system
    - gmp-system
    - cnrm-system

### Logging configuration
admissionController:
  container:
    extraArgs:
      "loggingFormat": "json"
      "v": "1"

backgroundController:
  enabled: true
  rbac:
    clusterRole:
      extraResources:
        - apiGroups: ["apps"]
          resources: ["deployments", "statefulsets", "daemonsets"]
          verbs: ["get", "list", "watch"]
```

> **Background Scan Interval**
>
> Set `backgroundScanInterval` to 6h for most clusters. Reduce to 1h for high-compliance environments. Increase to 12h for large clusters (1000+ nodes).
>

##### Policy Reporter Deployment

```bash
helm repo add policy-reporter https://kyverno.github.io/policy-reporter
helm repo update

helm install policy-reporter policy-reporter/policy-reporter \
  --namespace policy-reporter \
  --create-namespace \
  --values policy-reporter-values.yaml
```

**policy-reporter-values.yaml**:

```yaml
metrics:
  enabled: true

logging:
  encoding: json
  logLevel: -2
  development: false

api:
  logging: false

ui:
  enabled: true
  displayMode: dark

kyvernoPlugin:
  enabled: true
  metrics:
    enabled: true
```

> **Policy Reporter UI**
>
> Access the dashboard with `kubectl port-forward -n policy-reporter svc/policy-reporter-ui 8080:8080`. Navigate to [http://localhost:8080](http://localhost:8080).
>

---

#### Verification

##### Verify Kyverno Installation

```bash
### Check Kyverno pods
kubectl get pods -n kyverno

### Expected output:
### kyverno-admission-controller-xxx   Running
### kyverno-background-controller-xxx  Running
### kyverno-cleanup-controller-xxx     Running
### kyverno-reports-controller-xxx     Running
```

##### Verify Webhook Registration

```bash
### Check ValidatingWebhookConfiguration
kubectl get validatingwebhookconfiguration | grep kyverno

### Check MutatingWebhookConfiguration
kubectl get mutatingwebhookconfiguration | grep kyverno
```

##### Test Policy Enforcement

```bash
### Try deploying without resource limits
kubectl run test --image=nginx

### Expected: Denied by admission webhook
```

---

#### Next Steps

- **[Policy Enforcement](policy-enforcement.md)** - Deploy policies, configure modes, enable scanning
- **[Monitoring](monitoring.md)** - Dashboards, alerts, troubleshooting
- **[Operations](../operations/index.md)** - Day-to-day policy management

## Required Status Checks

CI/CD pipelines as merge gates.

Branch protection can require CI checks to pass before merge.

This turns "you should test" into "untested code cannot merge."

---

#### The Contract

GitHub won't allow merge until all required checks report success.

```yaml
### .github/workflows/required-checks.yml
name: Required Checks

on:
  pull_request:
    branches: [main]

jobs:
  tests:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - name: Run tests
        run: make test

  security-scan:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - name: Container scan
        run: |
          trivy image --severity HIGH,CRITICAL --exit-code 1 \
            gcr.io/project/app:${{ github.sha }}

  lint:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - name: Lint
        run: golangci-lint run
```

Configure branch protection to require these checks:

```yaml
required_status_checks:
  strict: true
  contexts:
    - "tests"
    - "security-scan"
    - "lint"
```

Code that fails tests, has HIGH CVEs, or doesn't pass linting cannot merge.

---

#### Automatic Audit Trail

> **Quick Start**
>
> This guide is part of a modular documentation set. Refer to related guides in the navigation for complete context.
>

GitHub stores check results permanently:

- Workflow run logs
- Exit codes and failure reasons
- Timestamps proving continuous enforcement
- Which commits triggered which checks

Auditors can query historical check results:

```bash
### Get workflow runs for a PR
gh api repos/org/repo/actions/runs \
  --jq '.workflow_runs[] | select(.head_branch=="feature-branch") |
    {name: .name, conclusion, created_at}'
```

---

#### Required Check Types

##### Unit and Integration Tests

```yaml
tests:
  runs-on: ubuntu-latest
  steps:
    - uses: actions/checkout@v4
    - name: Run tests
      run: |
        go test ./... -v -cover
        go test ./... -race
```

Test coverage isn't a metric. It's a gate.

##### Security Scanning

```yaml
security-scan:
  runs-on: ubuntu-latest
  steps:
    - uses: actions/checkout@v4
    - name: Build container
      run: docker build -t app:${{ github.sha }} .
    - name: Scan
      uses: aquasecurity/trivy-action@master
      with:
        image-ref: app:${{ github.sha }}
        severity: HIGH,CRITICAL
        exit-code: 1  # Fail if vulnerabilities found
```

See [Zero-Vulnerability Pipelines](../../blog/posts/2025-12-15-zero-vulnerability-pipelines.md) for full implementation.

##### Linting

```yaml
lint:
  runs-on: ubuntu-latest
  steps:
    - uses: actions/checkout@v4
    - uses: golangci/golangci-lint-action@v4
      with:
        args: --timeout=5m
```

Prevent code style bikeshedding in reviews. Linter enforces standards.

##### SBOM Generation

```yaml
sbom:
  runs-on: ubuntu-latest
  steps:
    - uses: actions/checkout@v4
    - uses: anchore/sbom-action@v0
      with:
        image: app:${{ github.sha }}
        format: cyclonedx-json
        output-file: sbom.json
```

Supply chain visibility becomes required evidence.

##### Forbidden Technology Check

```yaml
forbidden-tech:
  runs-on: ubuntu-latest
  steps:
    - uses: actions/checkout@v4
    - name: Check for forbidden technologies
      run: |
        ./scripts/check-forbidden-tech.sh
```

See [Pre-commit Security Gates](../../blog/posts/2025-12-04-pre-commit-security-gates.md) for forbidden technology enforcement.

---

#### Strict Mode

```yaml
required_status_checks:
  strict: true
```

**Strict mode** requires branch to be up-to-date with base before merge.

Without strict mode:

1. PR created from commit `abc123`
2. New commit `def456` merged to `main`
3. PR merges without integrating `def456`
4. Integration issues appear in `main`

With strict mode:

1. PR created from commit `abc123`
2. New commit `def456` merged to `main`
3. PR cannot merge until rebased on `def456`
4. All checks re-run against integration

Prevents "works in PR, breaks in main" scenarios.

---

#### Matrix Testing

Run checks across multiple configurations:

```yaml
tests:
  strategy:
    matrix:
      go: ['1.21', '1.22']
      os: [ubuntu-latest, macos-latest]
  runs-on: ${{ matrix.os }}
  steps:
    - uses: actions/checkout@v4
    - uses: actions/setup-go@v5
      with:
        go-version: ${{ matrix.go }}
    - name: Test
      run: go test ./...
```

All matrix combinations must pass.

---

#### Next Steps

- **[Configuration Patterns](configuration.md)** - Required vs optional checks, flaky tests, timing strategies
- **[Operations Guide](operations.md)** - Debugging, audit evidence, cost optimization

**Branch Protection Integration**:

- **[Branch Protection Rules](../branch-protection/branch-protection.md)** - Enforcement framework
- **[Security Tiers](../branch-protection/security-tiers.md)** - Required checks by tier
- **[Enforcement Workflows](../branch-protection/enforcement-workflows.md)** - Automated enforcement

**Related Controls**:

- **[Pre-commit Hooks](../pre-commit-hooks/pre-commit-hooks.md)** - Earlier validation
- **[Commit Signing](../commit-signing/commit-signing.md)** - Cryptographic proof of authorship

---

*Required checks blocked the PR. Tests failed. Vulnerabilities found. The code didn't merge. The pipeline worked.*

## SLSA Implementation Playbook

Complete SLSA implementation playbook: clarify SLSA…

Turn supply chain security from aspirational to operational.

> **What This Playbook Delivers**
>
>
> Clear implementation path from SLSA Level 1 to Level 3. Resolve SLSA vs SBOM confusion. Classify self-hosted runners correctly. Build verification workflows that actually work.
>

#### The Problem

SLSA adoption stalls on three documented pain points:

1. **SLSA vs SBOM confusion** - Teams can't decide which to implement first or when to use each
2. **Self-hosted runner classification** - Unclear which SLSA level self-hosted runners qualify for
3. **Verification complexity** - Generating provenance is easy, verifying it is hard

This playbook addresses all three directly.

---

#### What SLSA Actually Does

SLSA provenance **proves build integrity**. It answers:

- What source code produced this artifact?
- What build process ran?
- What environment executed the build?
- Was the build isolated from tampering?

SLSA is **not** a bill of materials. It's a cryptographic proof of the build process.

**Core value**: Detect compromised builds. Prevent supply chain attacks like SolarWinds.

---

#### SLSA vs SBOM: The Confusion

The SLSA community acknowledges this confusion as a problem they "hope to address." We address it now.

**Quick answer**:

- **SBOM** = Inventory (what's inside the artifact)
- **SLSA Provenance** = Proof (how the artifact was built)

You need both. They complement each other.

**Full clarification coming in subsequent playbook sections.**

---

#### Playbook Structure

This playbook is organized into focused sections covering the complete SLSA implementation journey:

##### Foundation & Conceptual Clarity

Start here to understand SLSA fundamentals:

- **SLSA vs SBOM** - When to use each, how they complement
- **SLSA Levels Explained** - Detailed requirements for Levels 1-4
- **[Current Implementation](slsa-provenance.md)** - GitHub Actions workflow patterns (available now)

##### Decision Trees & Classification

Determine your target SLSA level:

- **Level Classification** - Decision trees for GitHub-hosted and self-hosted runners
- **Runner Configuration** - SLSA implications for different runner types

##### Verification & Policy

Make SLSA provenance mandatory:

- **Verification Workflows** - Patterns for verifying provenance in CI/CD
- **Policy Templates** - Kyverno and OPA policies for enforcement

##### Incremental Adoption

Implement SLSA progressively:

- **Adoption Roadmap** - Phased approach from SLSA 1 to SLSA 3

##### Toolchain Integration

Language-specific implementation:

- **Toolchain Guides** - Go, Node.js, Python patterns
- **GitHub Actions Patterns** - Reusable workflow patterns

> **Playbook Content In Progress**
>
> Additional playbook sections are being developed. Start with the [current implementation guide](slsa-provenance.md) for immediate SLSA Level 3 provenance generation.
>

---

#### Quick Start: Choose Your Path

##### Path 1: I Need SLSA Level 3 Now

Use GitHub-hosted runners with `slsa-github-generator`:

```yaml
jobs:
  provenance:
    permissions:
      actions: read
      id-token: write
      contents: write
    uses: slsa-framework/slsa-github-generator/.github/workflows/generator_generic_slsa3.yml@v2.1.0
    with:
      base64-subjects: "${{ needs.build.outputs.hashes }}"
      upload-assets: true
```

**Result**: SLSA Level 3 provenance in one workflow change.

**Learn more**: [Current SLSA Implementation](slsa-provenance.md)

##### Path 2: I Have Self-Hosted Runners

Understand your SLSA level ceiling:

1. Review runner isolation requirements
2. Determine if self-hosted runners meet Level 2 isolation requirements
3. Consider migration to GitHub-hosted runners for Level 3

**Reality check**: Most self-hosted runners max out at SLSA Level 1 or 2.

**Detailed classification guidance coming in subsequent playbook sections.**

##### Path 3: I Need to Understand SLSA First

Start with conceptual foundation:

1. Review SLSA vs SBOM distinction (detailed guide coming soon)
2. Understand SLSA level requirements (detailed guide coming soon)
3. Read the [current implementation guide](slsa-provenance.md) for practical patterns
4. Plan your incremental adoption (roadmap guide coming soon)

---

#### SLSA Levels: Quick Reference

| Level | Build Script | Source Provenance | Build Isolation | Provenance Signing |
|-------|--------------|-------------------|-----------------|-------------------|
| **0** | None | None | None | None |
| **1** | Manual | Recorded | None | Manual |
| **2** | Automated | Versioned | None | Service-generated |
| **3** | Automated | Versioned | **Isolated** | **Non-falsifiable** |
| **4** | Automated | Versioned | **Hermetic** | **Non-falsifiable** + 2-party review |

**Critical distinction**: Level 3 requires **isolated build environments**. This is where self-hosted runners struggle.

**Full details coming in SLSA Levels guide.**

---

#### Common Scenarios

##### "We want OpenSSF Scorecard 10/10"

SLSA Level 3 provenance moves Signed-Releases from 8/10 to 10/10.

**What you need**:

1. Generate `.intoto.jsonl` attestation files
2. Upload to GitHub releases
3. Use GitHub-hosted runners (for Level 3)

**Implementation**: [Current SLSA Provenance Guide](slsa-provenance.md)

##### "We use self-hosted runners for security"

Self-hosted runners don't automatically disqualify you, but they cap your SLSA level.

**Questions to answer**:

- Are builds isolated from each other?
- Can developers access runner infrastructure?
- Are runner images ephemeral or persistent?

**Detailed guidance coming in Runner Configuration guide.**

##### "We already generate SBOMs"

Good. SBOM is complementary, not redundant.

- **SBOM**: Lists dependencies (inventory)
- **SLSA Provenance**: Proves build integrity (audit trail)

Both address supply chain risk from different angles.

**Detailed comparison coming in SLSA vs SBOM guide.**

---

#### Integration with Existing Controls

SLSA provenance layers with other enforcement mechanisms:

```mermaid
graph TD
    A[Source Code] -->|Branch Protection| B[Protected Branch]
    B -->|Status Checks| C[CI Pipeline]
    C -->|SLSA Provenance| D[Signed Artifact]
    D -->|Verification| E[Runtime Deployment]
    E -->|Policy-as-Code| F[Production]

    %% Ghostty Hardcore Theme
    style A fill:#a7e22e,color:#1b1d1e
    style D fill:#65d9ef,color:#1b1d1e
    style F fill:#f92572,color:#1b1d1e

```

**Integration points**:

1. **Branch Protection** - Require provenance generation in required status checks
2. **Status Checks** - Block PRs without provenance
3. **Policy-as-Code** - Verify provenance before admission
4. **Audit Evidence** - Include provenance in compliance reports

---

#### Why SLSA Matters

Supply chain attacks exploit build process compromise:

- **SolarWinds (2020)**: Attackers injected malware during build
- **Codecov (2021)**: Compromised Bash uploader script
- **ua-parser-js (2021)**: Hijacked npm package with backdoor

SLSA provenance detects these attacks by proving:

1. Artifact came from known source commit
2. Build ran in isolated environment
3. Build process matches expected workflow

**The gap SLSA fills**: It's not enough to sign artifacts. You must prove the build process itself is trustworthy.

---

#### Prerequisites

Before implementing SLSA:

- GitHub Actions (or equivalent CI/CD platform)
- Release artifact publishing (GitHub Releases, container registry)
- Basic understanding of cryptographic signing
- Decision on GitHub-hosted vs self-hosted runners

---

#### Implementation Timeline

**Realistic estimates**:

- **SLSA Level 1**: 1-2 days
- **SLSA Level 2**: 1 week (automated provenance generation)
- **SLSA Level 3**: 2-3 weeks (isolated builds, verification workflows)

**Complexity drivers**: Verification workflows, policy enforcement, self-hosted runner migration.

**Phased approach guide coming soon.**

---

#### Tools You'll Need

##### Provenance Generation

- **[slsa-github-generator](https://github.com/slsa-framework/slsa-github-generator)** - GitHub Actions reusable workflows
- **[slsa-framework/provenance-action](https://github.com/slsa-framework/slsa-github-generator)** - Alternative GitHub Action

##### Verification

- **[slsa-verifier](https://github.com/slsa-framework/slsa-verifier)** - CLI tool for provenance verification
- **[cosign](https://github.com/sigstore/cosign)** - Container image signing and verification

##### Policy Enforcement

- **[Kyverno](https://kyverno.io)** - Kubernetes admission control
- **[OPA Gatekeeper](https://open-policy-agent.github.io/gatekeeper/)** - Alternative policy engine

---

#### Next Steps

1. **Start implementing**: Use the [current SLSA provenance guide](slsa-provenance.md) for immediate Level 3 provenance
2. **Understand SLSA vs SBOM**: Detailed comparison guide coming soon
3. **Learn SLSA levels**: Requirements guide coming soon
4. **Classify your setup**: Runner configuration and level classification guides coming soon
5. **Plan adoption**: Incremental adoption roadmap guide coming soon

---

#### Related Content

- **[Current SLSA Implementation](slsa-provenance.md)** - Technical workflow details
- **[SBOM Generation](../../secure/sbom/sbom-generation.md)** - Complementary control
- **[OpenSSF Scorecard](../../secure/scorecard/scorecard-compliance.md)** - SLSA impact on scores
- **[Branch Protection](../branch-protection/branch-protection.md)** - Require provenance in status checks
- **[Policy-as-Code](../policy-as-code/index.md)** - Verify provenance at runtime

---

*SLSA provenance proves build integrity. Start with clarity, implement incrementally, verify everywhere.*

### SLSA Provenance: Toolchain Integration

Turn language-specific builds into provable pipelines.

<!-- more -->

> **Toolchain Integration Overview**
>
> This section covers SLSA Level 3 provenance generation for Go, Node.js, and Python projects. Each guide includes binary builds, package publishing, container image patterns, and dependency verification workflows.
>

#### Overview

Language-specific toolchains have unique SLSA integration points:

- **Build artifacts**: Binaries, packages, wheels, container images
- **Package registries**: npm, PyPI, Go modules, GitHub Packages
- **Dependency management**: go.sum, package-lock.json, poetry.lock
- **Build tools**: GoReleaser, npm scripts, setuptools, build isolation

Each toolchain guide covers:

- SLSA Level 3 provenance generation patterns
- Multi-platform and cross-compilation workflows
- Package registry integration
- Dependency lockfile verification
- Container image attestation
- Verification workflows
- Common gotchas and troubleshooting

---

#### Toolchain Guides

##### [Go Integration →](go-integration.md)

SLSA provenance for Go binary builds, multi-platform releases, and GoReleaser integration:

- **Binary builds**: Single and multi-platform cross-compilation
- **GoReleaser**: Automated release workflows with provenance
- **Go modules**: Dependency verification with go.sum
- **Container images**: Distroless patterns ([advanced guide](go-advanced.md))

**Key pattern**: Go's reproducible builds + SLSA provenance = non-falsifiable build integrity

##### [Node.js Integration →](node-integration.md)

SLSA provenance for npm packages, application artifacts, and container images:

- **npm packages**: Publishing with `npm publish --provenance`
- **Application artifacts**: Bundled JavaScript and TypeScript builds
- **Container images**: Multi-stage builds with Node runtime
- **Dependency lockfiles**: package-lock.json, yarn.lock, pnpm-lock.yaml
- **Registry verification**: npm audit signatures ([advanced guide](node-advanced.md))

**Key pattern**: Lockfile integrity + SLSA provenance = verified supply chain

##### [Python Integration →](python-integration.md)

SLSA provenance for PyPI packages, wheels, and container images:

- **PyPI packages**: Publishing wheels and source distributions
- **Application artifacts**: Wheels (.whl), source distributions (.tar.gz)
- **Container images**: Python runtime with application code
- **Dependency lockfiles**: requirements.txt, Pipfile.lock, poetry.lock

**Key pattern**: pip lockfiles + SLSA provenance = provable package publishing

---

#### Quick Reference

##### Toolchain Comparison

| Toolchain | Primary Artifact | Package Registry | Lockfile | SLSA Tool |
|-----------|------------------|------------------|----------|-----------|
| **Go** | Binary | Go modules | `go.sum` | `slsa-github-generator` |
| **Node.js** | npm package | npm, GitHub Packages | `package-lock.json` | `slsa-github-generator` |
| **Python** | Wheel (.whl) | PyPI, private registries | `poetry.lock` | `slsa-github-generator` |

##### Common Patterns

All toolchain guides follow these patterns:

1. Single artifact provenance
2. Multi-platform/multi-artifact builds
3. Package registry integration
4. Container image attestation
5. Verification with slsa-verifier
6. Dependency lockfile verification

##### Quick Start Commands

=== "Go"

    ```bash
##    # Generate provenance for Go binary
    go build -trimpath -ldflags="-buildid=" -o myapp

##    # Verify Go module checksums
    go mod verify

##    # Verify SLSA provenance
    slsa-verifier verify-artifact myapp \
      --provenance-path myapp.intoto.jsonl \
      --source-uri github.com/org/repo
    ```

=== "Node.js"

    ```bash
##    # Publish npm package with provenance
    npm publish --provenance

##    # Verify lockfile integrity
    npm ci --audit

##    # Verify SLSA provenance
    slsa-verifier verify-artifact artifact.tgz \
      --provenance-path artifact.tgz.intoto.jsonl \
      --source-uri github.com/org/repo
    ```

=== "Python"

    ```bash
##    # Build Python wheel
    python -m build

##    # Verify dependency hashes
    pip install --require-hashes -r requirements.txt

##    # Verify SLSA provenance
    slsa-verifier verify-artifact dist/mypackage-1.0.0-py3-none-any.whl \
      --provenance-path provenance.intoto.jsonl \
      --source-uri github.com/org/repo
    ```

---

#### Common Integration Patterns

##### Pattern: Multi-Artifact Provenance

All toolchains support generating provenance for multiple artifacts in a single build:

```yaml
jobs:
  build:
    outputs:
      hashes: ${{ steps.hash.outputs.hashes }}
    steps:
      - name: Build artifacts
        run: |
##          # Toolchain-specific build commands

      - name: Generate hashes
        id: hash
        run: |
          sha256sum artifacts/* | base64 -w0 > hashes.txt
          echo "hashes=$(cat hashes.txt)" >> "$GITHUB_OUTPUT"

  provenance:
    needs: [build]
    permissions:
      actions: read
      id-token: write
      contents: write
    uses: slsa-framework/slsa-github-generator/.github/workflows/generator_generic_slsa3.yml@v2.1.0
    with:
      base64-subjects: "${{ needs.build.outputs.hashes }}"
      upload-assets: true
```

This pattern works for:

- Go: Multiple binaries, multi-platform builds
- Node.js: Multiple npm packages, container images
- Python: Multiple wheels, source distributions

##### Pattern: Container Image Provenance

All toolchains support container image attestation:

```yaml
jobs:
  build-image:
    outputs:
      digest: ${{ steps.build.outputs.digest }}
    steps:
      - name: Build container image
        id: build
        run: |
##          # Toolchain-specific container build
          podman build -t myapp:latest .
          DIGEST=$(podman inspect myapp:latest --format='{{.Id}}')
          echo "digest=${DIGEST}" >> "$GITHUB_OUTPUT"

  provenance:
    needs: [build-image]
    permissions:
      actions: read
      id-token: write
      packages: write
    uses: slsa-framework/slsa-github-generator/.github/workflows/generator_container_slsa3.yml@v2.1.0
    with:
      image: ghcr.io/org/myapp
      digest: "${{ needs.build-image.outputs.digest }}"
```

See toolchain-specific guides for:

- Go: Distroless base images ([go-advanced.md](go-advanced.md))
- Node.js: Multi-stage builds ([node-advanced.md](node-advanced.md))
- Python: Python slim images ([python-integration.md](python-integration.md))

##### Pattern: Dependency Lockfile Verification

All toolchains support dependency verification:

=== "Go"

    ```yaml
    - name: Verify Go modules
      run: |
        go mod verify
        go mod download -json | jq -r '.Error' | grep -q '^null$'
    ```

=== "Node.js"

    ```yaml
    - name: Verify npm dependencies
      run: |
        npm ci --audit
        npm audit signatures
    ```

=== "Python"

    ```yaml
    - name: Verify Python dependencies
      run: |
        pip install --require-hashes -r requirements.txt
        pip check
    ```

---

#### Choosing Your Toolchain Guide

Pick the guide matching your project's main language:

- **Go projects** → [Go Integration Guide](go-integration.md)
- **Node.js/TypeScript projects** → [Node.js Integration Guide](node-integration.md)
- **Python projects** → [Python Integration Guide](python-integration.md)

For multi-language projects, start with your primary build artifact's guide, then use multi-artifact provenance pattern to cover all outputs.

---

#### Integration Checklist

Use this checklist when integrating SLSA provenance into your toolchain:

- [ ] **Choose toolchain guide** based on primary language
- [ ] **Review build patterns** for your artifact type (binary, package, container)
- [ ] **Implement provenance generation** using `slsa-github-generator`
- [ ] **Test provenance verification** with `slsa-verifier`
- [ ] **Verify dependency lockfiles** in CI/CD
- [ ] **Add deployment gates** requiring provenance verification
- [ ] **Document workflow** for team onboarding
- [ ] **Monitor OpenSSF Scorecard** for Signed-Releases improvement

---

#### Advanced Patterns

For advanced integration scenarios, see:

- **[Go Advanced Patterns](go-advanced.md)** - Container images, verification workflows, best practices
- **[Node.js Advanced Patterns](node-advanced.md)** - Registry verification, npm audit signatures, deployment gates
- **[Verification Workflows](../verification-workflows.md)** - Cross-toolchain verification patterns
- **[Policy Templates](../policy-templates.md)** - Kyverno and OPA enforcement for all toolchains

---

#### Common Questions

##### Do I need separate provenance for each language?

No. Use multi-artifact provenance to cover all build outputs in a single attestation.

##### Can I use the same verification workflow for all languages?

Yes. `slsa-verifier` works identically across toolchains regardless of artifact type.

##### Should I verify lockfiles or SLSA provenance?

**Both**. Lockfiles verify dependency inputs, SLSA provenance proves build environment integrity.

---

#### Next Steps

1. **Choose your toolchain guide**: [Go](go-integration.md), [Node.js](node-integration.md), or [Python](python-integration.md)
2. **Implement provenance generation**: Follow Pattern 1 in your toolchain guide
3. **Add verification workflow**: See toolchain-specific verification sections
4. **Enforce with policy**: Review [Policy Templates](../policy-templates.md)
5. **Scale adoption**: Follow [Adoption Roadmap](../adoption-roadmap.md)

---

#### Related Content

- **[SLSA Implementation Playbook](../index.md)** - Complete SLSA adoption guide
- **[SLSA Levels Explained](../slsa-levels.md)** - Understand Level 1-4 requirements
- **[Verification Workflows](../verification-workflows.md)** - CI/CD verification patterns
- **[Runner Configuration](../runner-configuration.md)** - GitHub-hosted vs self-hosted implications
- **[Adoption Roadmap](../adoption-roadmap.md)** - Phased implementation strategy

---

*Language-specific builds become provable pipelines. Choose your toolchain, implement provenance, verify everywhere.*
