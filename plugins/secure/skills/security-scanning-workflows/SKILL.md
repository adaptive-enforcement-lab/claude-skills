---
name: security-scanning-workflows
description: >-
  Comprehensive security scanning examples with SAST (CodeQL), dependency scanning, container vulnerability detection (Trivy), and SARIF upload to GitHub Security tab.
---

# Security Scanning Workflows

## When to Use This Skill

Copy-paste ready security scanning workflow templates with comprehensive coverage. Each example demonstrates SAST with CodeQL, dependency vulnerability detection, container image scanning with Trivy, and SARIF upload to GitHub Security tab for centralized visibility.

> **Complete Security Patterns**
>
>
> These workflows integrate all security scanning patterns: SHA-pinned actions, minimal GITHUB_TOKEN permissions (`security-events: write` for SARIF upload), automated scanning on every PR and push, SARIF result aggregation in GitHub Security tab, and security gates that block merges on critical findings.


## Implementation

See the full implementation guide in the [source documentation](https://adaptive-enforcement-lab.com/secure/github-actions-security/).


## Key Principles

Every security scanning workflow in this guide implements these controls:

1. **SAST Integration**: Static analysis with CodeQL to detect code-level vulnerabilities
2. **Dependency Scanning**: Automated vulnerability detection in dependencies with severity-based gates
3. **Container Scanning**: Image vulnerability scanning with Trivy before deployment
4. **SARIF Upload**: Centralized findings in GitHub Security tab for audit and tracking
5. **Security Gates**: Block merges on critical/high severity findings
6. **Minimal Permissions**: `security-events: write` scoped to scanning jobs only
7. **Scan All Changes**: Automated scanning on every PR and main branch push


## Full Reference

See [reference.md](reference.md) for complete documentation.
## References

- [Source Documentation](https://adaptive-enforcement-lab.com/secure/github-actions-security/)
- [AEL Secure](https://adaptive-enforcement-lab.com/secure/)
