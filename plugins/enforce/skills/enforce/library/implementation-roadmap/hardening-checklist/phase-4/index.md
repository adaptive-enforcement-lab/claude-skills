# Phase 4: Advanced (Month 4+)

Source: https://adaptive-enforcement-lab.com/enforce/implementation-roadmap/hardening-checklist/phase-4/

Prove compliance through continuous evidence collection and validation. Auditors get irrefutable proof that controls are real, not cosmetic.

> **Real-World Impact**
>
> A financial services company automated evidence collection for 6 months before their SOC 2 audit. Auditors requested 3 years of historical evidence. The team retrieved 36 months of branch protection configs, PR reviews, commit signatures, and SBOMs in under 10 minutes. Audit completed 2 weeks early.
>

---

## Phase Overview

Phase 4 completes the implementation with three critical areas:

1. **[Audit Evidence Collection](audit-evidence.md)** - Automated archival of branch protection, PR reviews, signatures, SBOMs
2. **[Compliance Validation](compliance.md)** - OpenSSF Scorecard, Best Practices Badge, SLSA verification, license checks
3. **[Audit Simulation](audit-simulation.md)** - Mock audit timeline, gap analysis, remediation

These controls provide continuous proof of compliance.

---

## Phase Components

### Audit Evidence Collection

Automated archival that ensures historical evidence exists when auditors ask.

**Key Controls**:

- Branch protection configuration archive
- Pull request review records
- Commit signature coverage tracking
- Workflow run logs
- SBOM and security scan results

**[View Audit Evidence Details →](audit-evidence.md)**

---

### Compliance Validation

Third-party verification and automated compliance reporting.

**Key Controls**:

- OpenSSF Scorecard monthly monitoring
- OpenSSF Best Practices Badge
- SLSA provenance verification
- Dependency license compliance
- Monthly compliance report generation

**[View Compliance Validation Details →](compliance.md)**

---

### Audit Simulation

Mock audit process to identify and fix gaps before real auditors arrive.

**Key Controls**:

- Document request simulation (Week 1)
- PR sampling and validation (Week 2)
- Gap analysis (Week 3)
- Remediation and verification (Week 4)

**[View Audit Simulation Details →](audit-simulation.md)**

---

## Phase 4 Validation Checklist

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

## Validation Commands

Test that controls are working:

```bash
# Verify evidence archive exists
gsutil ls gs://audit-evidence/2025-01/branch-protection.json
# Expected: File exists with branch protection config

# Check OpenSSF Scorecard score
docker run gcr.io/openssf/scorecard-action:stable --repo=github.com/org/repo
# Expected: Score ≥ 8.0/10

# Verify SLSA provenance
gh release view vX.Y.Z --json assets | jq '.assets[].name' | grep intoto
# Expected: .intoto.jsonl file exists

# Test evidence retrieval speed
time gsutil ls gs://audit-evidence/2024-*/branch-protection.json
# Expected: < 10 seconds
```

---

## Next Steps

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

## Related Patterns

- **[Audit Evidence Collection](../../../audit-compliance/audit-evidence.md)** - Evidence storage details
- **[OpenSSF Scorecard](../../../../secure/scorecard/scorecard-compliance.md)** - Scorecard configuration
- **[SLSA Provenance](../../../slsa-provenance/slsa-provenance.md)** - Build attestation
- **[Implementation Roadmap Overview](index.md)** - Complete roadmap
- **[Phase 3: Runtime](../phase-3/index.md)** - Runtime enforcement

---

*Evidence collected. Auditors satisfied. Scorecard 10/10. SLSA provenance verified. Compliance is automatic, not aspirational.*
