# Phase 2: Automation (Weeks 5-8)

Source: https://adaptive-enforcement-lab.com/enforce/implementation-roadmap/hardening-checklist/phase-2/

Automate security, quality, and compliance checks in the pipeline. Tests that fail, code with vulnerabilities, and builds without SBOMs never merge.

> **Real-World Impact**
>
> An e-commerce platform implemented CI gates and SBOM generation in 3 weeks. Within the first month, gates blocked 23 merges with HIGH/CRITICAL vulnerabilities and generated SBOMs for 847 builds. When Log4Shell hit, they had complete dependency visibility across all services in under 2 hours.
>

---

## Phase Overview

Phase 2 extends enforcement into the CI/CD pipeline through two critical areas:

1. **[CI/CD Gates](ci-gates.md)** - Required checks, SBOM generation, vulnerability scanning, SLSA provenance
2. **[Evidence Collection](evidence-collection.md)** - Automated archival and metrics tracking

These controls ensure failing builds never reach production and provide audit evidence.

---

## Phase Components

### CI/CD Gates

Pipeline enforcement that blocks merges with failing tests, vulnerabilities, or missing SBOMs.

**Key Controls**:

- Required status checks workflow
- SBOM generation for every build
- Vulnerability scanning with fail-fast
- SLSA provenance for releases
- Evidence storage integration

**[View CI/CD Gates Details →](ci-gates.md)**

---

### Evidence Collection

Automated archival of branch protection configs, PR reviews, and build artifacts.

**Key Controls**:

- Branch protection config snapshots
- PR review metadata collection
- Workflow run log archival
- Integration with branch protection
- Metrics tracking

**[View Evidence Collection Details →](evidence-collection.md)**

---

## Phase 2 Validation Checklist

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

## Validation Commands

Test that controls are working:

```bash
# Test CI blocks failing tests
echo "func TestFail(t *testing.T) { t.Fatal() }" >> main_test.go
git push origin feature-branch
# Expected: Merge blocked by CI failure

# Verify SBOM generation
gsutil ls gs://audit-evidence/sbom/$(date +%Y-%m-%d)/
# Expected: SBOM files for today's builds

# Verify SLSA provenance
gh release view vX.Y.Z --json assets | jq '.assets[].name' | grep intoto
# Expected: .intoto.jsonl file exists

# Verify evidence collection
gsutil ls gs://audit-evidence/2025-01/
# Expected: branch-protection.json, merged-prs.json
```

---

## Next Steps

With Phase 2 complete, you have:

- CI gates blocking failing tests and security scans
- SBOM generation for every build
- Vulnerability scanning with fail-fast
- SLSA provenance for all releases
- Automated evidence collection

**[Proceed to Phase 3: Runtime →](../phase-3/index.md)**

Phase 3 extends enforcement to runtime by controlling what can actually deploy to production.

---

## Related Patterns

- **[SLSA Provenance](../../../slsa-provenance/slsa-provenance.md)** - Build attestation details
- **[SBOM Generation](../../../../secure/sbom/sbom-generation.md)** - Software Bill of Materials
- **[Vulnerability Scanning](../../../../secure/vulnerability-scanning/vulnerability-scanning.md)** - Container image scanning
- **[Implementation Roadmap Overview](index.md)** - Complete roadmap
- **[Phase 1: Foundation](../phase-1/index.md)** - Pre-commit and branch protection
- **[Phase 3: Runtime →](../phase-3/index.md)** - Production policy enforcement

---

*CI gates deployed. SBOM generated. Vulnerabilities blocked. SLSA provenance signed. Evidence archived. Supply chain security is enforced, not suggested.*
