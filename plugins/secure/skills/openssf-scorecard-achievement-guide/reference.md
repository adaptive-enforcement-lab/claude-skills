---
name: openssf-scorecard-achievement-guide - Reference
description: Complete reference for OpenSSF Scorecard Achievement Guide
---

# OpenSSF Scorecard Achievement Guide - Reference

This is the complete reference documentation extracted from the source.

# OpenSSF Scorecard Achievement Guide

Comprehensive guide for understanding, interpreting, and improving OpenSSF Scorecard results. Covers all 18 checks, false positive handling, controversial check guidance, and remediation playbooks.

> **Start Here, Not with Scorecard**
>
>
> Don't chase a score. Build secure practices first, then measure them. High Scorecard scores are a byproduct of good security engineering, not the goal.
>

---

## What is OpenSSF Scorecard?

OpenSSF Scorecard is an automated security tool that checks repositories for supply chain security best practices. It evaluates 18 different security checks and produces scores from 0 to 10 for each check.

**Why it matters:**

- **Compliance**: Required by some enterprise procurement processes and security questionnaires
- **Supply chain security**: Identifies real vulnerabilities in your development and release processes
- **Best practices enforcement**: Automated checks ensure security practices don't regress

**What it doesn't do:**

- Replace security audits or penetration testing
- Catch all vulnerabilities because heuristic-based checks have limitations
- Understand context because some failures may be intentional design choices

---

## Quick Reference: All 18 Checks

| Check | Weight | Difficulty | Category | Quick Fix |
| ----- | ------ | ---------- | -------- | --------- |
| Binary-Artifacts | High | Easy | Supply Chain | Remove binaries from git |
| Branch-Protection | High | Medium | Code Review | Enable GitHub settings |
| CI-Tests | Low | Easy | Quality | Add test workflow |
| CII-Best-Practices | Low | High | Certification | Complete questionnaire |
| Code-Review | High | Medium | Code Review | Require PR reviews |
| Contributors | Low | N/A | Community | Encourage contributions |
| Dangerous-Workflow | High | Medium | Supply Chain | Fix workflow patterns |
| Dependency-Update-Tool | High | Easy | Dependencies | Enable Renovate/Dependabot |
| Fuzzing | Medium | High | Security | Integrate fuzzing |
| License | Low | Easy | Legal | Add LICENSE file |
| Maintained | Low | N/A | Activity | Regular commits |
| Packaging | Medium | Medium | Distribution | Publish packages |
| Pinned-Dependencies | High | Medium | Supply Chain | Pin to SHA digests |
| SAST | Medium | Easy | Security | Add static analysis |
| Security-Policy | Medium | Easy | Documentation | Add SECURITY.md |
| Signed-Releases | High | High | Supply Chain | SLSA provenance |
| Token-Permissions | High | Easy | Security | Job-level permissions |
| Vulnerabilities | High | Varies | Security | Fix known CVEs |

**Weight definitions:**

- **High**: Critical for supply chain security
- **Medium**: Important but not critical
- **Low**: Nice to have, signals project health

**Difficulty estimates:**

- **Easy**: 1 to 2 hours to fix
- **Medium**: Half day to implement
- **High**: Full day or more of work
- **N/A**: Not directly controllable

For detailed descriptions, see the check category guides linked below.

---

## Check Categories

### Supply Chain Security (6 checks)

The highest impact security checks that prevent supply chain attacks:

- **Binary-Artifacts**: Detects checked-in binaries that could hide malware
- **Dangerous-Workflow**: Identifies workflows that could leak secrets or execute untrusted code
- **Dependency-Update-Tool**: Ensures dependencies stay current with security patches
- **Pinned-Dependencies**: Prevents unexpected behavior from dependency updates
- **Signed-Releases**: Cryptographic proof that releases are authentic
- **Token-Permissions**: Limits blast radius of compromised workflows

**Priority**: Fix these first. They protect against real supply chain attacks.

### Code Review & Quality (4 checks)

Checks that ensure code quality and review processes:

- **Branch-Protection**: Enforces review requirements and prevents force pushes
- **CI-Tests**: Verifies automated testing exists
- **Code-Review**: Ensures human review before merge
- **Contributors**: Measures community diversity

**Priority**: Medium. Important for code quality, less critical for security.

### Security Practices (5 checks)

Active security tooling and policies:

- **Fuzzing**: Tests for unexpected input handling
- **SAST**: Static analysis security testing
- **Security-Policy**: Documented vulnerability reporting process
- **Vulnerabilities**: Known CVEs in dependencies
- **CII-Best-Practices**: Comprehensive security certification

**Priority**: High for Vulnerabilities, SAST, and Security-Policy. Medium for others.

### Project Health (3 checks)

Signals about project maturity and maintenance:

- **License**: Legal clarity for users
- **Maintained**: Recent activity signals active maintenance
- **Packaging**: Distribution through package managers

**Priority**: Low for security, high for adoption.

---

## Common Score Ranges

### Score 7 to 8: Good Security Hygiene

**What you have:**

- Automated testing in CI
- Dependency scanning
- Basic branch protection
- Security policy documented

**What's missing:**

- SLSA provenance for releases
- Job-level token permissions
- Comprehensive dependency pinning

**Time to fix**: 4 to 8 hours focused work

**See**: [Stuck at 8: The Journey to 10/10](../../blog/posts/2025-12-18-scorecard-stuck-at-eight.md)

### Score 8 to 9: Strong Security Posture

**What you have:**

- All checks from 7 to 8
- SLSA Level 3 provenance
- Job-level permissions
- SHA-pinned dependencies

**What's missing:**

- Perfect branch protection with 2+ reviewers and recent push approval
- Fuzzing integration
- CII Best Practices badge

**Time to fix**: 1 to 2 days

### Score 9 to 10: Exceptional Security

**What you have:**

- All previous checks passing
- Comprehensive security controls
- Advanced tooling including fuzzing and SLSA
- Community certification

**What's left:**

- Edge cases and false positives
- Documented exceptions for controversial checks
- Continuous monitoring and maintenance

**Time to fix**: Ongoing maintenance

---

## Detailed Guides

### Getting Started

Start with these guides for quick wins and foundational understanding:

- **[Scorecard Compliance](scorecard-compliance.md)** - Core patterns: job-level permissions, dependency pinning, source archive signing
- **[Workflow Examples](scorecard-workflow-examples.md)** - Production-ready workflows for 10/10 compliance

### Score Progression

Systematic approach to improving your score:

- **Score Progression Guide** *(Coming soon)* - Prioritized roadmap from 7 to 8 to 9 to 10

### Check-Specific Playbooks

Deep dives on check categories:

- **Supply Chain Checks** *(Coming soon)* - Pinned-Dependencies, Dangerous-Workflow, Binary-Artifacts, SAST
- **Code Review Checks** *(Coming soon)* - Code-Review, Contributors, Maintained, Branch-Protection
- **Security Practices Checks** *(Coming soon)* - Security-Policy, CII-Best-Practices, Vulnerabilities, Fuzzing, Token-Permissions
- **Release Security Checks** *(Coming soon)* - Signed-Releases, Packaging, License

### Advanced Topics

Navigate complexity and trade-offs:

- **False Positives Guide** *(Coming soon)* - Common false positive patterns and resolution approaches
- **Decision Framework** *(Coming soon)* - When to follow vs. deviate from Scorecard recommendations
- **CI/CD Integration** *(Coming soon)* - Automated Scorecard monitoring and regression prevention

---

## False Positives and Limitations

Scorecard uses heuristics, not perfect knowledge. Common false positive scenarios:

### Pinned-Dependencies Exceptions

**Issue**: Scorecard flags version tags for actions that **require** them.

**Examples**:

- `ossf/scorecard-action@v2.4.0` requires version tags for internal verification
- `slsa-framework/slsa-github-generator@v2.1.0` requires version tags for verifier validation

**Resolution**: Document the exception in Renovate config. These are legitimate deviations.

### Branch-Protection Admin Bypass

**Issue**: Scorecard penalizes allowing admins to bypass protections.

**Context**: Small teams may need admin bypass for emergency fixes.

**Resolution**: Decide based on team size and risk tolerance. Document the decision.

### Contributors Count

**Issue**: Solo-maintained projects can't increase contributor count.

**Context**: Single-maintainer projects are valid but flagged.

**Resolution**: Accept the score. Quality matters more than quantity.

---

## Controversial Recommendations

Not all Scorecard recommendations fit all contexts. Common debates:

### SHA Pinning vs. Semantic Versioning

**Scorecard position**: Pin everything to SHA digests.

**Counter-argument**: Version tags are more maintainable and Renovate handles updates.

**Our position**: SHA pin GitHub Actions for supply chain risk. Use version tags for dependencies with SemVer protection.

### Two Reviewers for Small Teams

**Scorecard position**: Require 2+ reviewers on all PRs.

**Counter-argument**: Solo maintainers or two-person teams can't meet this.

**Our position**: Enable for teams of 3+. Document exception for smaller teams.

### Fuzzing for All Projects

**Scorecard position**: All projects should have fuzzing.

**Counter-argument**: High implementation cost and low value for simple projects.

**Our position**: Prioritize for security-critical code such as parsers and crypto. Skip for CRUD apps.

---

## Related Content

### Blog Posts

Real-world Scorecard experiences:

- [OpenSSF Best Practices Badge in 2 Hours](../../blog/posts/2025-12-17-openssf-badge-two-hours.md) - Fast-track CII certification
- [Stuck at 8: The Journey to 10/10](../../blog/posts/2025-12-18-scorecard-stuck-at-eight.md) - SLSA provenance breakthrough

### Related Guides

- [SLSA Provenance](../../enforce/slsa-provenance/slsa-provenance.md) - Build attestations for Signed-Releases 10/10
- [SBOM Generation](../sbom/sbom-generation.md) - Complete attestation stack
- [GitHub Apps](../github-apps/index.md) - Secure authentication patterns

---

## Next Steps

1. **Run Scorecard**: Get your baseline score with `ossf/scorecard-action`
2. **Quick wins**: Fix Token-Permissions, add SECURITY.md, enable Dependabot
3. **Medium effort**: Implement SLSA provenance and pin dependencies to SHA
4. **High effort**: Add fuzzing, earn CII badge, perfect branch protection
5. **Maintenance**: Monitor score, prevent regressions, update as Scorecard evolves

**Remember**: Scorecard measures security practices. Don't game the score. Build secure systems.

---

*Scorecard is a tool, not a goal. Use it to find real security gaps, not to chase a number. The best score is the one that reflects actual security investment.*

