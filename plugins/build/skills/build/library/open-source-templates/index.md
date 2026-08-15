# Open Source Project Templates

Source: https://adaptive-enforcement-lab.com/build/open-source-templates/

Copy-paste templates for open source project documentation based on real OpenSSF Best Practices Badge certification work. CONTRIBUTING.md, SECURITY.md, and GitHub issue forms with realistic SLAs and proven compliance.

> **Source Material**
>
> These templates come from the [readability project's OpenSSF certification](https://github.com/adaptive-enforcement-lab/readability) (PRs [#93](https://github.com/adaptive-enforcement-lab/readability/pull/93), [#94](https://github.com/adaptive-enforcement-lab/readability/pull/94), [#95](https://github.com/adaptive-enforcement-lab/readability/pull/95)).
>

---

## Why These Templates Matter

OpenSSF Best Practices Badge certification revealed the pattern: most projects have solid technical infrastructure but lack **documentation**.

The badge criteria require:

- ✅ CONTRIBUTING.md with development setup and testing requirements
- ✅ SECURITY.md with disclosure process and response timelines
- ✅ Issue templates for bug reports and feature requests
- ✅ Clear communication of contribution guidelines

These are the main gap for well-maintained projects. Templates fill the gap.

---

## Template Library

### CONTRIBUTING.md Template

The [CONTRIBUTING.md template](contributing-template.md) provides:

- Annotated template with placeholders for project-specific details
- Real example from production (readability project)
- Development setup, code style, testing requirements
- Pull request process and commit message conventions
- Code review requirements

[View CONTRIBUTING Template →](contributing-template.md)

---

### SECURITY.md Template

The [SECURITY.md template](security-template.md) provides:

- Template with realistic SLAs and private disclosure mechanisms
- Real example from production (readability project)
- Supported versions table
- Reporting process (GitHub Security Advisories + email)
- Response timelines by severity level
- Security measures documentation

[View SECURITY Template →](security-template.md)

---

### GitHub Issue Templates

The [GitHub issue templates](issue-templates.md) provide:

- Bug report YAML form template
- Feature request YAML form template
- Template configuration (config.yml)
- Structured fields with validation

[View Issue Templates →](issue-templates.md)

---

## OpenSSF Best Practices Alignment

How these templates satisfy OpenSSF Badge criteria:

| Criterion | Template | Compliance |
|-----------|----------|------------|
| **Documentation** | CONTRIBUTING.md | ✅ Explains how to contribute |
| **Bug Reporting** | Bug Report template | ✅ Structured process |
| **Enhancement Proposals** | Feature Request template | ✅ Clear submission path |
| **Security Process** | SECURITY.md | ✅ Disclosure mechanism |
| **Response Timelines** | SECURITY.md SLAs | ✅ Realistic commitments |
| **Testing Requirements** | CONTRIBUTING.md | ✅ Coverage thresholds |
| **Code Review** | CONTRIBUTING.md PR process | ✅ Approval requirements |

### Badge Checklist Mapping

✅ **Contributing file**: CONTRIBUTING.md with setup, testing, PR process

✅ **Bug reporting**: Issue templates with structured fields

✅ **Enhancement proposals**: Feature request template

✅ **Security disclosure**: SECURITY.md with private channel (Security Advisories)

✅ **Security response**: Documented SLAs (48hr initial, 7 day update, 90 day resolution)

---

## Common Gaps These Templates Fill

### 1. Missing Security Disclosure Process

**Problem**: No SECURITY.md, reporters open public issues with vulnerability details.

**Fix**: SECURITY.md with GitHub Security Advisories link + email fallback.

### 2. Unrealistic SLAs

**Problem**: SECURITY.md promises "immediate response" that never happens.

**Fix**: Realistic timelines (48 hours, not "immediately"). Better to under-promise and over-deliver.

### 3. Unstructured Bug Reports

**Problem**: Issues say "it doesn't work" with zero details.

**Fix**: YAML issue templates with required fields for reproduction steps, environment, logs.

### 4. No Testing Requirements

**Problem**: PRs without tests get merged, coverage drops.

**Fix**: CONTRIBUTING.md explicitly states coverage threshold and test commands.

### 5. No Development Setup

**Problem**: Contributors can't set up project locally.

**Fix**: CONTRIBUTING.md with prerequisite list and installation commands.

---

## Customization Checklist

When adapting these templates:

- [ ] Replace `[PROJECT_NAME]` with actual project name
- [ ] Replace `[ORG]` with GitHub organization
- [ ] Replace `[LANGUAGE]`, `[VERSION]` with tech stack
- [ ] Update install commands (`npm install`, `pip install`, `go mod download`)
- [ ] Update test commands (`npm test`, `pytest`, `go test`)
- [ ] Update linter/formatter names (`eslint`, `black`, `golangci-lint`)
- [ ] Set realistic SLA timelines based on team capacity
- [ ] Update supported versions table in SECURITY.md
- [ ] Configure issue template labels to match your project
- [ ] Update security measures list (SBOM, scanning tools, etc.)

---

## Related Patterns

- Blog: [OpenSSF Best Practices Badge in 2 Hours](../../blog/posts/2025-12-17-openssf-badge-two-hours.md) - How these templates enable fast certification
- [SLSA Provenance Implementation](../../enforce/slsa-provenance/slsa-provenance.md) - Security measures referenced in SECURITY.md
- [SBOM Generation](../../secure/sbom/sbom-generation.md) - Supply chain transparency

---

*The gap is never the code. It's the documentation. These templates close that gap. Copy, paste, customize, commit. OpenSSF Badge unlocked.*
