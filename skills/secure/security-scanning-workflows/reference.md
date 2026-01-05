---
name: security-scanning-workflows - Reference
description: Complete reference for Security Scanning Workflows
---

# Security Scanning Workflows - Reference

This is the complete reference documentation extracted from the source.


# Security Scanning Workflows

Copy-paste ready security scanning workflow templates with comprehensive coverage. Each example demonstrates SAST with CodeQL, dependency vulnerability detection, container image scanning with Trivy, and SARIF upload to GitHub Security tab for centralized visibility.

> **Complete Security Patterns**
>
>
> These workflows integrate all security scanning patterns: SHA-pinned actions, minimal GITHUB_TOKEN permissions (`security-events: write` for SARIF upload), automated scanning on every PR and push, SARIF result aggregation in GitHub Security tab, and security gates that block merges on critical findings.
>

## Security Scanning Principles

Every security scanning workflow in this guide implements these controls:

1. **SAST Integration**: Static analysis with CodeQL to detect code-level vulnerabilities
2. **Dependency Scanning**: Automated vulnerability detection in dependencies with severity-based gates
3. **Container Scanning**: Image vulnerability scanning with Trivy before deployment
4. **SARIF Upload**: Centralized findings in GitHub Security tab for audit and tracking
5. **Security Gates**: Block merges on critical/high severity findings
6. **Minimal Permissions**: `security-events: write` scoped to scanning jobs only
7. **Scan All Changes**: Automated scanning on every PR and main branch push

## Universal Security Scanning Workflow

Comprehensive scanning workflow covering SAST, dependencies, and containers in one pipeline.

### Multi-Scanner Security Pipeline

Complete security scanning with CodeQL, dependency review, and Trivy.

```yaml
name: Security Scanning
on:
  push:
    branches: [main, develop]
  pull_request:
    branches: [main, develop]
  schedule:
    # SECURITY: Weekly scheduled scan catches newly-disclosed vulnerabilities
    # Run every Monday at 08:00 UTC
    - cron: '0 8 * * 1'

# SECURITY: Minimal permissions by default
permissions:
  contents: read

jobs:
  # Job 1: SAST with CodeQL
  codeql-analysis:
    name: CodeQL SAST Analysis
    runs-on: ubuntu-latest
    permissions:
      contents: read        # Read repository code
      security-events: write  # Upload SARIF to Security tab
      actions: read         # Read workflow metadata
    strategy:
      fail-fast: false
      matrix:
        # SECURITY: Scan all languages in monorepo
        language: ['javascript', 'python']
    steps:
      # SECURITY: All actions pinned to full SHA-256 commit hashes
      - name: Checkout code
        uses: actions/checkout@b4ffde65f46336ab88eb53be808477a3936bae11  # v4.1.1
        with:
          persist-credentials: false

      # SECURITY: Initialize CodeQL for static analysis
      - name: Initialize CodeQL
        uses: github/codeql-action/init@cdcdbb579706841c47f7063dda365e292e5cad7a  # v2.13.4
        with:
          languages: ${{ matrix.language }}
          # SECURITY: security-extended includes additional checks beyond default suite
          # Use security-and-quality for maximum coverage (slower)
          queries: security-extended
          # SECURITY: Threat modeling configuration
          # Identifies sources (user input) and sinks (sensitive operations)
          config-file: ./.github/codeql/codeql-config.yml

      # SECURITY: Autobuild for compiled languages (Java, C++, C#, Go)
      # For interpreted languages (JavaScript, Python), this is a no-op
      - name: Autobuild
        uses: github/codeql-action/autobuild@cdcdbb579706841c47f7063dda365e292e5cad7a  # v2.13.4

      # SECURITY: Perform CodeQL analysis and upload results
      - name: Perform CodeQL Analysis
        uses: github/codeql-action/analyze@cdcdbb579706841c47f7063dda365e292e5cad7a  # v2.13.4
        with:
          # SECURITY: Category allows multiple analyses per repository
          # Use language as category for monorepo scanning
          category: "/language:${{ matrix.language }}"
          # SECURITY: Upload SARIF to Security tab (requires security-events: write)
          upload: true
          # SECURITY: Fail workflow on high/critical findings
          # Comment out for informational-only scanning
          # fail-on: high

  # Job 2: Dependency vulnerability scanning
  dependency-scan:
    name: Dependency Vulnerability Scan
    runs-on: ubuntu-latest
    permissions:
      contents: read
      pull-requests: write  # Post review comments on PRs
    steps:
      - name: Checkout code
        uses: actions/checkout@b4ffde65f46336ab88eb53be808477a3936bae11  # v4.1.1
        with:
          persist-credentials: false

      # SECURITY: Dependency review detects vulnerable and malicious packages in PRs
      # Only runs on pull_request events (not push)
      - name: Dependency Review
        if: github.event_name == 'pull_request'
        uses: actions/dependency-review-action@c74b580d73376b7750d3d2a50bfb8adc2c937507  # v3.1.0
        with:
          # SECURITY: Fail on critical/high vulnerabilities
          fail-on-severity: high
          # SECURITY: Deny licenses incompatible with your policy
          deny-licenses: AGPL-3.0, GPL-3.0
          # SECURITY: Warn on moderate/low vulnerabilities
          warn-on-severity: moderate
          # SECURITY: Comment threshold reduces PR noise
          comment-summary-in-pr: true
          # SECURITY: Allow specific packages if needed (use sparingly)
          # allow-dependencies-licenses: MIT, Apache-2.0

  # Job 3: Container image vulnerability scanning
  container-scan:
    name: Container Image Scan
    runs-on: ubuntu-latest
    # SECURITY: Only scan containers on main branch and PRs from same repo
    # Prevents fork PRs from triggering container builds
    if: github.event_name == 'push' || github.event.pull_request.head.repo.full_name == github.repository
    permissions:
      contents: read
      security-events: write  # Upload SARIF to Security tab
    steps:
      - name: Checkout code
        uses: actions/checkout@b4ffde65f46336ab88eb53be808477a3936bae11  # v4.1.1
        with:
          persist-credentials: false

      # SECURITY: Build container image for scanning
      # In production, scan images from registry instead of building
      - name: Build container image
        run: |
          podman build -t myapp:${{ github.sha }} .

      # SECURITY: Scan container image with Trivy
      - name: Run Trivy vulnerability scanner
        uses: aquasecurity/trivy-action@d43c1f16c00cfd3978dde6c07f4bbcf9eb6993ca  # 0.16.1
        with:
          # SECURITY: Scan filesystem and dependencies in container
          scan-type: 'image'
          image-ref: 'myapp:${{ github.sha }}'
          # SECURITY: SARIF format for GitHub Security tab upload
          format: 'sarif'
          output: 'trivy-results.sarif'
          # SECURITY: Fail on critical/high vulnerabilities
          severity: 'CRITICAL,HIGH'
          # SECURITY: Exit code 1 if vulnerabilities found (blocks merge)
          exit-code: '1'
          # SECURITY: Ignore unfixed vulnerabilities (optional)
          # ignore-unfixed: true

      # SECURITY: Upload Trivy results to Security tab
      - name: Upload Trivy SARIF results
        if: always()  # Upload even if scan fails
        uses: github/codeql-action/upload-sarif@cdcdbb579706841c47f7063dda365e292e5cad7a  # v2.13.4
        with:
          sarif_file: 'trivy-results.sarif'
          # SECURITY: Category allows multiple scan results
          category: 'trivy-container'

      # SECURITY: Also generate human-readable report
      - name: Generate Trivy report
        if: always()
        uses: aquasecurity/trivy-action@d43c1f16c00cfd3978dde6c07f4bbcf9eb6993ca  # 0.16.1
        with:
          scan-type: 'image'
          image-ref: 'myapp:${{ github.sha }}'
          format: 'table'
          output: 'trivy-report.txt'

      - name: Upload Trivy report
        if: always()
        uses: actions/upload-artifact@c7d193f32edcb7bfad88892161225aeda64e9392  # v4.0.0
        with:
          name: trivy-report
          path: trivy-report.txt
          retention-days: 30

  # Job 4: Secret scanning verification
  secret-scan:
    name: Secret Scanning Verification
    runs-on: ubuntu-latest
    permissions:
      contents: read
    steps:
      - name: Checkout code
        uses: actions/checkout@b4ffde65f46336ab88eb53be808477a3936bae11  # v4.1.1
        with:
          # SECURITY: Fetch full history to scan all commits in PR
          fetch-depth: 0
          persist-credentials: false

      # SECURITY: Gitleaks scans for hardcoded secrets in commit history
      - name: Run gitleaks secret scan
        uses: gitleaks/gitleaks-action@cb7149a9c69f0f7c6a0c5b7b094889a91831ff7f  # v2.3.2
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
          # SECURITY: Don't expose findings in PR comments (use Security tab)
          GITLEAKS_ENABLE_COMMENTS: false
          # SECURITY: Fail on secret detection
          GITLEAKS_ENABLE_UPLOAD_ARTIFACT: true
```

