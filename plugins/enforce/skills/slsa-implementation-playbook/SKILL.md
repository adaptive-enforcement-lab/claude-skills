---
name: slsa-implementation-playbook
description: >-
  Complete SLSA implementation playbook: clarify SLSA vs SBOM confusion, classify runner configurations, implement verification workflows, and adopt incrementally from Level 1 to Level 3.
---

# SLSA Implementation Playbook

## When to Use This Skill

Turn supply chain security from aspirational to operational.

> **What This Playbook Delivers**
>
>
> Clear implementation path from SLSA Level 1 to Level 3. Resolve SLSA vs SBOM confusion. Classify self-hosted runners correctly. Build verification workflows that actually work.


## Prerequisites

Before implementing SLSA:

- GitHub Actions (or equivalent CI/CD platform)
- Release artifact publishing (GitHub Releases, container registry)
- Basic understanding of cryptographic signing
- Decision on GitHub-hosted vs self-hosted runners

---


## When to Apply

### "We want OpenSSF Scorecard 10/10"

SLSA Level 3 provenance moves Signed-Releases from 8/10 to 10/10.

**What you need**:

1. Generate `.intoto.jsonl` attestation files
2. Upload to GitHub releases
3. Use GitHub-hosted runners (for Level 3)

**Implementation**: [Current SLSA Provenance Guide](slsa-provenance.md)

### "We use self-hosted runners for security"

Self-hosted runners don't automatically disqualify you, but they cap your SLSA level.

**Questions to answer**:

- Are builds isolated from each other?
- Can developers access runner infrastructure?
- Are runner images ephemeral or persistent?

**Detailed guidance coming in Runner Configuration guide.**

### "We already generate SBOMs"

Good. SBOM is complementary, not redundant.

- **SBOM**: Lists dependencies (inventory)
- **SLSA Provenance**: Proves build integrity (audit trail)

Both address supply chain risk from different angles.

**Detailed comparison coming in SLSA vs SBOM guide.**

---


## Implementation

Turn supply chain security from aspirational to operational.

> **What This Playbook Delivers**
>
>
> Clear implementation path from SLSA Level 1 to Level 3. Resolve SLSA vs SBOM confusion. Classify self-hosted runners correctly. Build verification workflows that actually work.
>

## The Problem

SLSA adoption stalls on three documented pain points:

1. **SLSA vs SBOM confusion** - Teams can't decide which to implement first or when to use each
2. **Self-hosted runner classification** - Unclear which SLSA level self-hosted runners qualify for
3. **Verification complexity** - Generating provenance is easy, verifying it is hard

This playbook addresses all three directly.

---

## What SLSA Actually Does

SLSA provenance **proves build integrity**. It answers:

- What source code produced this artifact?
- What build process ran?
- What environment executed the build?
- Was the build isolated from tampering?

SLSA is **not** a bill of materials. It's a cryptographic proof of the build process.

**Core value**: Detect compromised builds. Prevent supply chain attacks like SolarWinds.

---

## SLSA vs SBOM: The Confusion

The SLSA community acknowledges this confusion as a problem they "hope to address." We address it now.

**Quick answer**:

- **SBOM** = Inventory (what's inside the artifact)
- **SLSA Provenance** = Proof (how the artifact was built)

You need both. They complement each other.

**Full clarification coming in subsequent playbook sections.**

---

## Playbook Structure

This playbook is organized into focused sections covering the complete SLSA implementation journey:

### Foundation & Conceptual Clarity

Start here to understand SLSA fundamentals:

- **SLSA vs SBOM** - When to use each, how they complement
- **SLSA Levels Explained** - Detailed requirements for Levels 1-4
- **[Current Implementation](slsa-provenance.md)** - GitHub Actions workflow patterns (available now)

### Decision Trees & Classification

Determine your target SLSA level:

- **Level Classification** - Decision trees for GitHub-hosted and self-hosted runners
- **Runner Configuration** - SLSA implications for different runner types

### Verification & Policy

Make SLSA provenance mandatory:

- **Verification Workflows** - Patterns for verifying provenance in CI/CD
- **Policy Templates** - Kyverno and OPA policies for enforcement

### Incremental Adoption

Implement SLSA progressively:

- **Adoption Roadmap** - Phased approach from SLSA 1 to SLSA 3

### Toolchain Integration

Language-specific implementation:

- **Toolchain Guides** - Go, Node.js, Python patterns
- **GitHub Actions Patterns** - Reusable workflow patterns

> **Playbook Content In Progress**
>
> Additional playbook sections are being developed. Start with the [current implementation guide](slsa-provenance.md) for immediate SLSA Level 3 provenance generation.
>

---

## Quick Start: Choose Your Path

### Path 1: I Need SLSA Level 3 Now

Use GitHub-hosted runners with `slsa-github-generator`:


*See [examples.md](examples.md) for detailed code examples.*

**Result**: SLSA Level 3 provenance in one workflow change.

**Learn more**: [Current SLSA Implementation](slsa-provenance.md)

### Path 2: I Have Self-Hosted Runners

Understand your SLSA level ceiling:

1. Review runner isolation requirements
2. Determine if self-hosted runners meet Level 2 isolation requirements
3. Consider migration to GitHub-hosted runners for Level 3

**Reality check**: Most self-hosted runners max out at SLSA Level 1 or 2.

**Detailed classification guidance coming in subsequent playbook sections.**

### Path 3: I Need to Understand SLSA First

Start with conceptual foundation:

1. Review SLSA vs SBOM distinction (detailed guide coming soon)
2. Understand SLSA level requirements (detailed guide coming soon)
3. Read the [current implementation guide](slsa-provenance.md) for practical patterns
4. Plan your incremental adoption (roadmap guide coming soon)

---

## SLSA Levels: Quick Reference

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

## Common Scenarios

### "We want OpenSSF Scorecard 10/10"

SLSA Level 3 provenance moves Signed-Releases from 8/10 to 10/10.

**What you need**:

1. Generate `.intoto.jsonl` attestation files
2. Upload to GitHub releases
3. Use GitHub-hosted runners (for Level 3)

**Implementation**: [Current SLSA Provenance Guide](slsa-provenance.md)

### "We use self-hosted runners for security"

Self-hosted runners don't automatically disqualify you, but they cap your SLSA level.

**Questions to answer**:

- Are builds isolated from each other?
- Can developers access runner infrastructure?
- Are runner images ephemeral or persistent?

**Detailed guidance coming in Runner Configuration guide.**

### "We already generate SBOMs"

Good. SBOM is complementary, not redundant.

- **SBOM**: Lists dependencies (inventory)
- **SLSA Provenance**: Proves build integrity (audit trail)

Both address supply chain risk from different angles.

**Detailed comparison coming in SLSA vs SBOM guide.**

---

## Integration with Existing Controls

SLSA provenance layers with other enforcement mechanisms:


*See [examples.md](examples.md) for detailed code examples.*

**Integration points**:

1. **Branch Protection** - Require provenance generation in required status checks
2. **Status Checks** - Block PRs without provenance
3. **Policy-as-Code** - Verify provenance before admission
4. **Audit Evidence** - Include provenance in compliance reports

---

## Why SLSA Matters

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

## Prerequisites

Before implementing SLSA:

- GitHub Actions (or equivalent CI/CD platform)
- Release artifact publishing (GitHub Releases, container registry)
- Basic understanding of cryptographic signing
- Decision on GitHub-hosted vs self-hosted runners

---

## Implementation Timeline

**Realistic estimates**:

- **SLSA Level 1**: 1-2 days
- **SLSA Level 2**: 1 week (automated provenance generation)
- **SLSA Level 3**: 2-3 weeks (isolated builds, verification workflows)

**Complexity drivers**: Verification workflows, policy enforcement, self-hosted runner migration.

**Phased approach guide coming soon.**

---

## Tools You'll Need

### Provenance Generation

- **[slsa-github-generator](https://github.com/slsa-framework/slsa-github-generator)** - GitHub Actions reusable workflows
- **[slsa-framework/provenance-action](https://github.com/slsa-framework/provenance-action)** - Alternative GitHub Action

### Verification

- **[slsa-verifier](https://github.com/slsa-framework/slsa-verifier)** - CLI tool for provenance verification
- **[cosign](https://github.com/sigstore/cosign)** - Container image signing and verification

### Policy Enforcement

- **[Kyverno](https://kyverno.io)** - Kubernetes admission control
- **[OPA Gatekeeper](https://open-policy-agent.github.io/gatekeeper/)** - Alternative policy engine

---

## Next Steps

1. **Start implementing**: Use the [current SLSA provenance guide](slsa-provenance.md) for immediate Level 3 provenance
2. **Understand SLSA vs SBOM**: Detailed comparison guide coming soon
3. **Learn SLSA levels**: Requirements guide coming soon
4. **Classify your setup**: Runner configuration and level classification guides coming soon
5. **Plan adoption**: Incremental adoption roadmap guide coming soon

---

## Related Content

- **[Current SLSA Implementation](slsa-provenance.md)** - Technical workflow details
- **[SBOM Generation](../../secure/sbom/sbom-generation.md)** - Complementary control
- **[OpenSSF Scorecard](../../secure/scorecard/scorecard-compliance.md)** - SLSA impact on scores
- **[Branch Protection](../branch-protection/branch-protection.md)** - Require provenance in status checks
- **[Policy-as-Code](../policy-as-code/index.md)** - Verify provenance at runtime

---

*SLSA provenance proves build integrity. Start with clarity, implement incrementally, verify everywhere.*

### The Problem

SLSA adoption stalls on three documented pain points:

1. **SLSA vs SBOM confusion** - Teams can't decide which to implement first or when to use each
2. **Self-hosted runner classification** - Unclear which SLSA level self-hosted runners qualify for
3. **Verification complexity** - Generating provenance is easy, verifying it is hard

This playbook addresses all three directly.

---

### What SLSA Actually Does

SLSA provenance **proves build integrity**. It answers:

- What source code produced this artifact?
- What build process ran?
- What environment executed the build?
- Was the build isolated from tampering?

SLSA is **not** a bill of materials. It's a cryptographic proof of the build process.

**Core value**: Detect compromised builds. Prevent supply chain attacks like SolarWinds.

---

### SLSA vs SBOM: The Confusion

The SLSA community acknowledges this confusion as a problem they "hope to address." We address it now.

**Quick answer**:

- **SBOM** = Inventory (what's inside the artifact)
- **SLSA Provenance** = Proof (how the artifact was built)

You need both. They complement each other.

**Full clarification coming in subsequent playbook sections.**

---

### Playbook Structure

This playbook is organized into focused sections covering the complete SLSA implementation journey:

### Foundation & Conceptual Clarity

Start here to understand SLSA fundamentals:

- **SLSA vs SBOM** - When to use each, how they complement
- **SLSA Levels Explained** - Detailed requirements for Levels 1-4
- **[Current Implementation](slsa-provenance.md)** - GitHub Actions workflow patterns (available now)

### Decision Trees & Classification

Determine your target SLSA level:

- **Level Classification** - Decision trees for GitHub-hosted and self-hosted runners
- **Runner Configuration** - SLSA implications for different runner types

### Verification & Policy

Make SLSA provenance mandatory:

- **Verification Workflows** - Patterns for verifying provenance in CI/CD
- **Policy Templates** - Kyverno and OPA policies for enforcement

### Incremental Adoption

Implement SLSA progressively:

- **Adoption Roadmap** - Phased approach from SLSA 1 to SLSA 3

### Toolchain Integration

Language-specific implementation:

- **Toolchain Guides** - Go, Node.js, Python patterns
- **GitHub Actions Patterns** - Reusable workflow patterns

> **Playbook Content In Progress**
>
> Additional playbook sections are being developed. Start with the [current implementation guide](slsa-provenance.md) for immediate SLSA Level 3 provenance generation.
>

---

### Quick Start: Choose Your Path

### Path 1: I Need SLSA Level 3 Now

Use GitHub-hosted runners with `slsa-github-generator`:


*See [examples.md](examples.md) for detailed code examples.*

**Result**: SLSA Level 3 provenance in one workflow change.

**Learn more**: [Current SLSA Implementation](slsa-provenance.md)

### Path 2: I Have Self-Hosted Runners

Understand your SLSA level ceiling:

1. Review runner isolation requirements
2. Determine if self-hosted runners meet Level 2 isolation requirements
3. Consider migration to GitHub-hosted runners for Level 3

**Reality check**: Most self-hosted runners max out at SLSA Level 1 or 2.

**Detailed classification guidance coming in subsequent playbook sections.**

### Path 3: I Need to Understand SLSA First

Start with conceptual foundation:

1. Review SLSA vs SBOM distinction (detailed guide coming soon)
2. Understand SLSA level requirements (detailed guide coming soon)
3. Read the [current implementation guide](slsa-provenance.md) for practical patterns
4. Plan your incremental adoption (roadmap guide coming soon)

---

### SLSA Levels: Quick Reference

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

### Common Scenarios

### "We want OpenSSF Scorecard 10/10"

SLSA Level 3 provenance moves Signed-Releases from 8/10 to 10/10.

**What you need**:

1. Generate `.intoto.jsonl` attestation files
2. Upload to GitHub releases
3. Use GitHub-hosted runners (for Level 3)

**Implementation**: [Current SLSA Provenance Guide](slsa-provenance.md)

### "We use self-hosted runners for security"

Self-hosted runners don't automatically disqualify you, but they cap your SLSA level.

**Questions to answer**:

- Are builds isolated from each other?
- Can developers access runner infrastructure?
- Are runner images ephemeral or persistent?

**Detailed guidance coming in Runner Configuration guide.**

### "We already generate SBOMs"

Good. SBOM is complementary, not redundant.

- **SBOM**: Lists dependencies (inventory)
- **SLSA Provenance**: Proves build integrity (audit trail)

Both address supply chain risk from different angles.

**Detailed comparison coming in SLSA vs SBOM guide.**

---

### Integration with Existing Controls

SLSA provenance layers with other enforcement mechanisms:


*See [examples.md](examples.md) for detailed code examples.*

**Integration points**:

1. **Branch Protection** - Require provenance generation in required status checks
2. **Status Checks** - Block PRs without provenance
3. **Policy-as-Code** - Verify provenance before admission
4. **Audit Evidence** - Include provenance in compliance reports

---

### Why SLSA Matters

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

### Prerequisites

Before implementing SLSA:

- GitHub Actions (or equivalent CI/CD platform)
- Release artifact publishing (GitHub Releases, container registry)
- Basic understanding of cryptographic signing
- Decision on GitHub-hosted vs self-hosted runners

---

### Implementation Timeline

**Realistic estimates**:

- **SLSA Level 1**: 1-2 days
- **SLSA Level 2**: 1 week (automated provenance generation)
- **SLSA Level 3**: 2-3 weeks (isolated builds, verification workflows)

**Complexity drivers**: Verification workflows, policy enforcement, self-hosted runner migration.

**Phased approach guide coming soon.**

---

### Tools You'll Need

### Provenance Generation

- **[slsa-github-generator](https://github.com/slsa-framework/slsa-github-generator)** - GitHub Actions reusable workflows
- **[slsa-framework/provenance-action](https://github.com/slsa-framework/provenance-action)** - Alternative GitHub Action

### Verification

- **[slsa-verifier](https://github.com/slsa-framework/slsa-verifier)** - CLI tool for provenance verification
- **[cosign](https://github.com/sigstore/cosign)** - Container image signing and verification

### Policy Enforcement

- **[Kyverno](https://kyverno.io)** - Kubernetes admission control
- **[OPA Gatekeeper](https://open-policy-agent.github.io/gatekeeper/)** - Alternative policy engine

---

### Next Steps

1. **Start implementing**: Use the [current SLSA provenance guide](slsa-provenance.md) for immediate Level 3 provenance
2. **Understand SLSA vs SBOM**: Detailed comparison guide coming soon
3. **Learn SLSA levels**: Requirements guide coming soon
4. **Classify your setup**: Runner configuration and level classification guides coming soon
5. **Plan adoption**: Incremental adoption roadmap guide coming soon

---

### Related Content

- **[Current SLSA Implementation](slsa-provenance.md)** - Technical workflow details
- **[SBOM Generation](../../secure/sbom/sbom-generation.md)** - Complementary control
- **[OpenSSF Scorecard](../../secure/scorecard/scorecard-compliance.md)** - SLSA impact on scores
- **[Branch Protection](../branch-protection/branch-protection.md)** - Require provenance in status checks
- **[Policy-as-Code](../policy-as-code/index.md)** - Verify provenance at runtime

---

*SLSA provenance proves build integrity. Start with clarity, implement incrementally, verify everywhere.*


## Comparison

The SLSA community acknowledges this confusion as a problem they "hope to address." We address it now.

**Quick answer**:

- **SBOM** = Inventory (what's inside the artifact)
- **SLSA Provenance** = Proof (how the artifact was built)

You need both. They complement each other.

**Full clarification coming in subsequent playbook sections.**

---


## Examples

See [examples.md](examples.md) for code examples.


## Full Reference

See [reference.md](reference.md) for complete documentation.


## Related Patterns

- Current SLSA Implementation
- SBOM Generation
- OpenSSF Scorecard
- Branch Protection
- Policy-as-Code

## References

- [Source Documentation](https://adaptive-enforcement-lab.com/enforce/slsa-provenance/)
- [AEL Enforce](https://adaptive-enforcement-lab.com/enforce/)
