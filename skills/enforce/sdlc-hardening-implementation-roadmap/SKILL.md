---
name: sdlc-hardening-implementation-roadmap
description: >-
  Deploy defense-in-depth SDLC hardening across four phases: pre-commit hooks, CI/CD gates, runtime enforcement, and continuous audit evidence collection systems.
---

# SDLC Hardening Implementation Roadmap

## When to Use This Skill

This implementation roadmap provides a structured approach to hardening your Software Development Lifecycle (SDLC) across four critical phases:

1. **[Phase 1: Foundation](phase-1/index.md)** - Local enforcement and branch protection
2. **[Phase 2: Automation](phase-2/index.md)** - CI/CD gates and policy automation
3. **[Phase 3: Runtime](phase-3/index.md)** - Production policy enforcement
4. **[Phase 4: Advanced](phase-4/index.md)** - Audit evidence and compliance validation

Each phase builds on the previous one, creating defense-in-depth through multiple enforcement layers.

---


## Prerequisites

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


## Implementation

Every control in this roadmap is actionable and verifiable. No vague policies. No wishful thinking.

> **Audit Foundation**
>
> These controls are what auditors will verify. Skip items at your own risk. Each control must be fully deployed and evidenced before claiming compliance.
>

---

## Overview

This implementation roadmap provides a structured approach to hardening your Software Development Lifecycle (SDLC) across four critical phases:

1. **[Phase 1: Foundation](phase-1/index.md)** - Local enforcement and branch protection
2. **[Phase 2: Automation](phase-2/index.md)** - CI/CD gates and policy automation
3. **[Phase 3: Runtime](phase-3/index.md)** - Production policy enforcement
4. **[Phase 4: Advanced](phase-4/index.md)** - Audit evidence and compliance validation

Each phase builds on the previous one, creating defense-in-depth through multiple enforcement layers.

---

## Roadmap Phases

### Phase 1: Foundation (Weeks 1-4)

Establish local development controls and repository protection.

**Phase Components**:

- **[Pre-commit Hooks](phase-1/pre-commit-hooks.md)** - Secrets detection, linting, policy enforcement
- **[Branch Protection](phase-1/branch-protection.md)** - Required reviews, status checks, admin enforcement

**Why This Phase Matters**: If secrets enter git history, rotation doesn't help. If admins can bypass reviews, the policy is worthless. Foundation controls prevent bad code from ever entering the system.

**[View Phase 1 Overview →](phase-1/index.md)**

---

### Phase 2: Automation (Weeks 5-8)

Automate security and quality checks in CI/CD pipelines.

**Phase Components**:

- **[CI/CD Gates](phase-2/ci-gates.md)** - Required checks, SBOM generation, vulnerability scanning, SLSA provenance
- **[Evidence Collection](phase-2/evidence-collection.md)** - Automated archival and metrics tracking

**Why This Phase Matters**: Tests that fail, code with vulnerabilities, and builds without SBOMs never merge. CI becomes a gate, not a log. Supply chain security becomes automatic.

**[View Phase 2 Overview →](phase-2/index.md)**

---

### Phase 3: Runtime (Weeks 9-12)

Control what runs in production, not just what gets committed.

**Phase Components**:

- **[Policy Enforcement](phase-3/policy-enforcement.md)** - Core Kyverno policies, resource limits, image verification
- **[Advanced Policies](phase-3/advanced-policies.md)** - Namespace quotas, pod security, network policies
- **[Rollout Strategy](phase-3/rollout.md)** - Audit-first deployment and metrics

**Why This Phase Matters**: Pods without resource limits, images from untrusted registries, or missing security context cannot run. Policy is enforced before deployment, not after incidents.

**[View Phase 3 Overview →](phase-3/index.md)**

---

### Phase 4: Advanced (Month 4+)

Prove compliance through continuous evidence collection and validation.

**Phase Components**:

- **[Audit Evidence](phase-4/audit-evidence.md)** - Branch protection, PR reviews, signatures, SBOMs
- **[Compliance Validation](phase-4/compliance.md)** - OpenSSF Scorecard, SLSA verification, license checks
- **[Audit Simulation](phase-4/audit-simulation.md)** - Mock audit, gap analysis, remediation

**Why This Phase Matters**: Auditors will ask "prove branch protection was enabled on 2025-01-01". Archived config proves it. Evidence collection must be automatic and tamper-proof.

**[View Phase 4 Overview →](phase-4/index.md)**

---

## Implementation Timeline

| Phase | Timeline | Key Milestone | Validation Method |
|-------|----------|---------------|-------------------|
| **Phase 1: Foundation** | Weeks 1-4 | Branch protection on all repos | Test admin bypass attempt |
| **Phase 2: Automation** | Weeks 5-8 | CI gates block failing tests | Merge attempt with failing test |
| **Phase 3: Runtime** | Weeks 9-12 | Kyverno enforces pod policies | Deploy pod without limits |
| **Phase 4: Advanced** | Month 4+ | OpenSSF Scorecard 10/10 | Automated evidence retrieval |

---

## Critical Success Factors

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

## Validation Strategy

Each phase includes validation steps that prove controls are working:

**Foundation Phase**:

```bash
# Test pre-commit hook blocks secrets
echo "AKIAIOSFODNN7EXAMPLE" > .env && git add .env && git commit -m "test"
# Expected: Commit blocked by TruffleHog

# Test admin enforcement
gh api repos/org/repo/branches/main/protection | jq '.enforce_admins.enabled'
# Expected: true
```

**Automation Phase**:

```bash
# Test CI blocks failing tests
echo "func TestFail(t *testing.T) { t.Fatal() }" >> main_test.go
git push origin feature-branch
# Expected: Merge blocked by CI failure

# Verify SBOM generation
gsutil ls gs://audit-evidence/sbom/$(date +%Y-%m-%d)/
# Expected: SBOM files for today's builds
```

**Runtime Phase**:

```bash
# Test pod without resource limits is rejected
kubectl apply -f pod-no-limits.yaml
# Expected: Admission webhook denies request

# Test untrusted registry is blocked
kubectl apply -f pod-dockerhub.yaml
# Expected: Image source validation fails
```

**Advanced Phase**:

```bash
# Verify evidence archive exists
gsutil ls gs://audit-evidence/2025-01/branch-protection.json
# Expected: File exists with branch protection config

# Check OpenSSF Scorecard score
docker run gcr.io/openssf/scorecard-action:stable --repo=github.com/org/repo
# Expected: Score ≥ 8.0/10
```

---

## Prerequisites

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

## Common Pitfalls

### Pitfall 1: Deploying controls without validation

Don't assume a control works because it's deployed. Every control must be tested with an attack scenario.

**Solution**: Use validation commands from each phase to prove controls block violations.

### Pitfall 2: Admin bypass enabled "just in case"

Setting `enforce_admins: false` makes all other controls optional.

**Solution**: Keep admin enforcement enabled. If emergency bypass is needed, document it, use it, then re-enable immediately.

### Pitfall 3: Evidence collection runs but isn't verified

Automated evidence collection means nothing if the data is corrupt or incomplete.

**Solution**: Monthly spot-check of evidence archives. Verify JSON is valid, timestamps are correct, and files are complete.

### Pitfall 4: OpenSSF Scorecard score drops unnoticed

Your score can regress if controls are disabled or practices slip.

**Solution**: Run Scorecard monthly. Alert on score drops > 0.5 points. Investigate immediately.

---

## Next Steps

1. **Review** [Phase 1: Foundation](phase-1/index.md) and identify repositories for pilot deployment
2. **Prepare** cloud storage bucket for evidence collection
3. **Document** current state (how many repos have branch protection? How many require reviews?)
4. **Schedule** implementation kickoff with engineering teams
5. **Execute** Phase 1 controls on pilot repository
6. **Validate** controls work before scaling organization-wide

---

## Related Patterns

- **[Execution Guide](../execution.md)** - Progress tracking and rollback planning
- **[Branch Protection Enforcement](../../branch-protection/branch-protection.md)** - GitHub configuration
- **[Pre-commit Security Gates](../../pre-commit-hooks/pre-commit-hooks.md)** - Local enforcement
- **[SLSA Provenance](../../slsa-provenance/slsa-provenance.md)** - Build attestations
- **[Audit Evidence Collection](../../audit-compliance/audit-evidence.md)** - Long-term evidence storage
- **[Policy-as-Code with Kyverno](../../policy-as-code/kyverno/index.md)** - Runtime enforcement

---

*Foundation laid. Controls enforced. Evidence collected. Auditors get irrefutable proof. SDLC hardening is not a checklist item. It's operational reality.*

### Overview

This implementation roadmap provides a structured approach to hardening your Software Development Lifecycle (SDLC) across four critical phases:

1. **[Phase 1: Foundation](phase-1/index.md)** - Local enforcement and branch protection
2. **[Phase 2: Automation](phase-2/index.md)** - CI/CD gates and policy automation
3. **[Phase 3: Runtime](phase-3/index.md)** - Production policy enforcement
4. **[Phase 4: Advanced](phase-4/index.md)** - Audit evidence and compliance validation

Each phase builds on the previous one, creating defense-in-depth through multiple enforcement layers.

---

### Roadmap Phases

### Phase 1: Foundation (Weeks 1-4)

Establish local development controls and repository protection.

**Phase Components**:

- **[Pre-commit Hooks](phase-1/pre-commit-hooks.md)** - Secrets detection, linting, policy enforcement
- **[Branch Protection](phase-1/branch-protection.md)** - Required reviews, status checks, admin enforcement

**Why This Phase Matters**: If secrets enter git history, rotation doesn't help. If admins can bypass reviews, the policy is worthless. Foundation controls prevent bad code from ever entering the system.

**[View Phase 1 Overview →](phase-1/index.md)**

---

### Phase 2: Automation (Weeks 5-8)

Automate security and quality checks in CI/CD pipelines.

**Phase Components**:

- **[CI/CD Gates](phase-2/ci-gates.md)** - Required checks, SBOM generation, vulnerability scanning, SLSA provenance
- **[Evidence Collection](phase-2/evidence-collection.md)** - Automated archival and metrics tracking

**Why This Phase Matters**: Tests that fail, code with vulnerabilities, and builds without SBOMs never merge. CI becomes a gate, not a log. Supply chain security becomes automatic.

**[View Phase 2 Overview →](phase-2/index.md)**

---

### Phase 3: Runtime (Weeks 9-12)

Control what runs in production, not just what gets committed.

**Phase Components**:

- **[Policy Enforcement](phase-3/policy-enforcement.md)** - Core Kyverno policies, resource limits, image verification
- **[Advanced Policies](phase-3/advanced-policies.md)** - Namespace quotas, pod security, network policies
- **[Rollout Strategy](phase-3/rollout.md)** - Audit-first deployment and metrics

**Why This Phase Matters**: Pods without resource limits, images from untrusted registries, or missing security context cannot run. Policy is enforced before deployment, not after incidents.

**[View Phase 3 Overview →](phase-3/index.md)**

---

### Phase 4: Advanced (Month 4+)

Prove compliance through continuous evidence collection and validation.

**Phase Components**:

- **[Audit Evidence](phase-4/audit-evidence.md)** - Branch protection, PR reviews, signatures, SBOMs
- **[Compliance Validation](phase-4/compliance.md)** - OpenSSF Scorecard, SLSA verification, license checks
- **[Audit Simulation](phase-4/audit-simulation.md)** - Mock audit, gap analysis, remediation

**Why This Phase Matters**: Auditors will ask "prove branch protection was enabled on 2025-01-01". Archived config proves it. Evidence collection must be automatic and tamper-proof.

**[View Phase 4 Overview →](phase-4/index.md)**

---

### Implementation Timeline

| Phase | Timeline | Key Milestone | Validation Method |
|-------|----------|---------------|-------------------|
| **Phase 1: Foundation** | Weeks 1-4 | Branch protection on all repos | Test admin bypass attempt |
| **Phase 2: Automation** | Weeks 5-8 | CI gates block failing tests | Merge attempt with failing test |
| **Phase 3: Runtime** | Weeks 9-12 | Kyverno enforces pod policies | Deploy pod without limits |
| **Phase 4: Advanced** | Month 4+ | OpenSSF Scorecard 10/10 | Automated evidence retrieval |

---

### Critical Success Factors

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

### Validation Strategy

Each phase includes validation steps that prove controls are working:

**Foundation Phase**:

```bash
# Test pre-commit hook blocks secrets
echo "AKIAIOSFODNN7EXAMPLE" > .env && git add .env && git commit -m "test"
# Expected: Commit blocked by TruffleHog

# Test admin enforcement
gh api repos/org/repo/branches/main/protection | jq '.enforce_admins.enabled'
# Expected: true
```

**Automation Phase**:

```bash
# Test CI blocks failing tests
echo "func TestFail(t *testing.T) { t.Fatal() }" >> main_test.go
git push origin feature-branch
# Expected: Merge blocked by CI failure

# Verify SBOM generation
gsutil ls gs://audit-evidence/sbom/$(date +%Y-%m-%d)/
# Expected: SBOM files for today's builds
```

**Runtime Phase**:

```bash
# Test pod without resource limits is rejected
kubectl apply -f pod-no-limits.yaml
# Expected: Admission webhook denies request

# Test untrusted registry is blocked
kubectl apply -f pod-dockerhub.yaml
# Expected: Image source validation fails
```

**Advanced Phase**:

```bash
# Verify evidence archive exists
gsutil ls gs://audit-evidence/2025-01/branch-protection.json
# Expected: File exists with branch protection config

# Check OpenSSF Scorecard score
docker run gcr.io/openssf/scorecard-action:stable --repo=github.com/org/repo
# Expected: Score ≥ 8.0/10
```

---

### Prerequisites

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

### Common Pitfalls

### Pitfall 1: Deploying controls without validation

Don't assume a control works because it's deployed. Every control must be tested with an attack scenario.

**Solution**: Use validation commands from each phase to prove controls block violations.

### Pitfall 2: Admin bypass enabled "just in case"

Setting `enforce_admins: false` makes all other controls optional.

**Solution**: Keep admin enforcement enabled. If emergency bypass is needed, document it, use it, then re-enable immediately.

### Pitfall 3: Evidence collection runs but isn't verified

Automated evidence collection means nothing if the data is corrupt or incomplete.

**Solution**: Monthly spot-check of evidence archives. Verify JSON is valid, timestamps are correct, and files are complete.

### Pitfall 4: OpenSSF Scorecard score drops unnoticed

Your score can regress if controls are disabled or practices slip.

**Solution**: Run Scorecard monthly. Alert on score drops > 0.5 points. Investigate immediately.

---

### Next Steps

1. **Review** [Phase 1: Foundation](phase-1/index.md) and identify repositories for pilot deployment
2. **Prepare** cloud storage bucket for evidence collection
3. **Document** current state (how many repos have branch protection? How many require reviews?)
4. **Schedule** implementation kickoff with engineering teams
5. **Execute** Phase 1 controls on pilot repository
6. **Validate** controls work before scaling organization-wide

---

### Related Patterns

- **[Execution Guide](../execution.md)** - Progress tracking and rollback planning
- **[Branch Protection Enforcement](../../branch-protection/branch-protection.md)** - GitHub configuration
- **[Pre-commit Security Gates](../../pre-commit-hooks/pre-commit-hooks.md)** - Local enforcement
- **[SLSA Provenance](../../slsa-provenance/slsa-provenance.md)** - Build attestations
- **[Audit Evidence Collection](../../audit-compliance/audit-evidence.md)** - Long-term evidence storage
- **[Policy-as-Code with Kyverno](../../policy-as-code/kyverno/index.md)** - Runtime enforcement

---

*Foundation laid. Controls enforced. Evidence collected. Auditors get irrefutable proof. SDLC hardening is not a checklist item. It's operational reality.*


## Anti-Patterns to Avoid

### Pitfall 1: Deploying controls without validation

Don't assume a control works because it's deployed. Every control must be tested with an attack scenario.

**Solution**: Use validation commands from each phase to prove controls block violations.

### Pitfall 2: Admin bypass enabled "just in case"

Setting `enforce_admins: false` makes all other controls optional.

**Solution**: Keep admin enforcement enabled. If emergency bypass is needed, document it, use it, then re-enable immediately.

### Pitfall 3: Evidence collection runs but isn't verified

Automated evidence collection means nothing if the data is corrupt or incomplete.

**Solution**: Monthly spot-check of evidence archives. Verify JSON is valid, timestamps are correct, and files are complete.

### Pitfall 4: OpenSSF Scorecard score drops unnoticed

Your score can regress if controls are disabled or practices slip.

**Solution**: Run Scorecard monthly. Alert on score drops > 0.5 points. Investigate immediately.

---


## Examples

See [examples.md](examples.md) for code examples.


## Full Reference

See [reference.md](reference.md) for complete documentation.


## Related Patterns

- Execution Guide
- Branch Protection Enforcement
- Pre-commit Security Gates
- SLSA Provenance
- Audit Evidence Collection
- Policy-as-Code with Kyverno

## References

- [Source Documentation](https://adaptive-enforcement-lab.com/enforce/implementation-roadmap/)
- [AEL Enforce](https://adaptive-enforcement-lab.com/enforce/)
