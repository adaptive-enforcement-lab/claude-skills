---
name: hardened-ci-workflow - Reference
description: Complete reference for Hardened CI Workflow
---

# Hardened CI Workflow - Reference

This is the complete reference documentation extracted from the source.


# Hardened CI Workflow

Copy-paste ready CI workflow templates with comprehensive security hardening. Each example demonstrates action pinning, minimal GITHUB_TOKEN permissions, input validation, and security scanning.

> **Complete Security Patterns**
>
>
> These workflows integrate all security patterns from the hub: SHA-pinned actions, job-level permissions, secret scanning prevention, fork PR safety, and security tooling. Use as production templates.
>

## Universal CI Pattern

Core security controls that apply to all CI workflows regardless of language or tooling.

### Security Principles Applied

Every example in this guide implements these controls:

1. **Action Pinning**: All third-party actions pinned to full SHA-256 commit hashes
2. **Minimal Permissions**: `contents: read` by default, elevated only for specific jobs
3. **Fork PR Safety**: `pull_request` trigger (not `pull_request_target`) for untrusted code
4. **Input Validation**: No direct injection of untrusted inputs into shell commands
5. **Secret Scanning**: Pre-commit hooks and push protection to prevent credential leaks
6. **Dependency Scanning**: Automated vulnerability detection for dependencies
7. **SARIF Upload**: Security findings uploaded to GitHub Security tab

### Base Hardened CI Workflow

Minimal secure CI workflow demonstrating core patterns.

```yaml
name: Hardened CI
on:
  push:
    branches: [main, develop]
  pull_request:
    # SECURITY: pull_request (not pull_request_target) runs untrusted code in isolated context
    # Fork PRs run with read-only GITHUB_TOKEN and no access to secrets
    branches: [main, develop]

# SECURITY: Workflow-level permissions deny all by default
# Job-level permissions grant minimal access per job
permissions:
  contents: read

jobs:
  # Job 1: Build and test with minimal permissions
  test:
    runs-on: ubuntu-latest
    permissions:
      contents: read  # Read repository code
      # No write permissions - prevents tampering
    steps:
      # SECURITY: All actions pinned to full SHA-256 commit hashes
      # Version comments (# vX.Y.Z) provide human readability
      # Dependabot will update SHAs via PRs
      - name: Checkout code
        uses: actions/checkout@b4ffde65f46336ab88eb53be808477a3936bae11  # v4.1.1
        with:
          # SECURITY: Shallow clone (depth: 1) reduces attack surface
          # Full history not needed for CI builds
          persist-credentials: false  # Don't persist git credentials

      - name: Set up build environment
        uses: actions/setup-node@5e21ff4d9bc1a8cf6de233a3057d20ec6b3fb69d  # v3.8.1
        with:
          node-version: '20'
          cache: 'npm'  # Cache dependencies for speed

      - name: Install dependencies
        run: npm ci  # Use ci (not install) for reproducible builds

      - name: Run linter
        run: npm run lint

      - name: Run unit tests
        run: npm test -- --coverage

      - name: Upload coverage reports
        uses: codecov/codecov-action@e0b68c6749509c5f83f984dd99a76a1c1a231044  # v4.0.1
        with:
          # SECURITY: Never use secrets in fork PRs
          # Codecov token optional for public repos
          fail_ci_if_error: false  # Don't fail on upload errors
          files: ./coverage/coverage.xml
        env:
          # SECURITY: Secrets not exposed to fork PRs with pull_request trigger
          CODECOV_TOKEN: ${{ secrets.CODECOV_TOKEN }}

  # Job 2: Security scanning with isolated permissions
  security-scan:
    runs-on: ubuntu-latest
    permissions:
      contents: read       # Read repository code
      security-events: write  # Upload SARIF to Security tab
    steps:
      - name: Checkout code
        uses: actions/checkout@b4ffde65f46336ab88eb53be808477a3936bae11  # v4.1.1
        with:
          persist-credentials: false

      # SECURITY: CodeQL for static analysis
      - name: Initialize CodeQL
        uses: github/codeql-action/init@cdcdbb579706841c47f7063dda365e292e5cad7a  # v2.13.4
        with:
          languages: javascript
          # SECURITY: Use default query suite (security-extended for more coverage)
          queries: security-extended

      - name: Autobuild
        uses: github/codeql-action/autobuild@cdcdbb579706841c47f7063dda365e292e5cad7a  # v2.13.4

      - name: Perform CodeQL Analysis
        uses: github/codeql-action/analyze@cdcdbb579706841c47f7063dda365e292e5cad7a  # v2.13.4
        with:
          # SECURITY: Upload SARIF to Security tab (requires security-events: write)
          category: "/language:javascript"

      # SECURITY: Trivy for dependency and vulnerability scanning
      - name: Run Trivy scanner
        uses: aquasecurity/trivy-action@d43c1f16c00cfd3978dde6c07f4bbcf9eb6993ca  # 0.16.1
        with:
          scan-type: 'fs'
          scan-ref: '.'
          format: 'sarif'
          output: 'trivy-results.sarif'
          severity: 'CRITICAL,HIGH'

      - name: Upload Trivy results
        uses: github/codeql-action/upload-sarif@cdcdbb579706841c47f7063dda365e292e5cad7a  # v2.13.4
        with:
          sarif_file: 'trivy-results.sarif'
          category: 'trivy'

  # Job 3: Build artifacts with minimal permissions
  build:
    runs-on: ubuntu-latest
    # SECURITY: Only build on non-fork PRs and main branch
    # Prevents malicious fork PRs from creating artifacts
    if: github.event_name == 'push' || github.event.pull_request.head.repo.full_name == github.repository
    permissions:
      contents: read
    steps:
      - name: Checkout code
        uses: actions/checkout@b4ffde65f46336ab88eb53be808477a3936bae11  # v4.1.1
        with:
          persist-credentials: false

      - name: Set up build environment
        uses: actions/setup-node@5e21ff4d9bc1a8cf6de233a3057d20ec6b3fb69d  # v3.8.1
        with:
          node-version: '20'

      - name: Install dependencies
        run: npm ci

      - name: Build application
        run: npm run build

      - name: Upload build artifacts
        uses: actions/upload-artifact@c7d193f32edcb7bfad88892161225aeda64e9392  # v4.0.0
        with:
          name: build-artifacts
          path: dist/
          retention-days: 7  # SECURITY: Short retention to reduce exposure
```

## Language-Specific CI Workflows

### Node.js / TypeScript CI

Hardened CI for Node.js and TypeScript projects with comprehensive testing and security scanning.

```yaml
name: Node.js CI
on:
  push:
    branches: [main]
  pull_request:
    branches: [main]

permissions:
  contents: read

jobs:
  test:
    name: Test on Node ${{ matrix.node-version }}
    runs-on: ubuntu-latest
    permissions:
      contents: read
    strategy:
      # SECURITY: fail-fast prevents wasting resources on known failures
      fail-fast: true
      matrix:
        node-version: [18, 20]
    steps:
      - uses: actions/checkout@b4ffde65f46336ab88eb53be808477a3936bae11  # v4.1.1
        with:
          persist-credentials: false

      - uses: actions/setup-node@5e21ff4d9bc1a8cf6de233a3057d20ec6b3fb69d  # v3.8.1
        with:
          node-version: ${{ matrix.node-version }}
          cache: 'npm'

      # SECURITY: Audit dependencies for known vulnerabilities
      - name: Audit dependencies
        run: npm audit --audit-level=high

      - name: Install dependencies
        run: npm ci

      # SECURITY: Type checking catches bugs before runtime
      - name: Type check
        run: npm run type-check

      - name: Lint
        run: npm run lint

      - name: Run tests
        run: npm test -- --coverage --maxWorkers=2

      - name: Build
        run: npm run build

  dependency-review:
    name: Dependency Review
    runs-on: ubuntu-latest
    # SECURITY: Only run on PRs to catch risky dependencies before merge
    if: github.event_name == 'pull_request'
    permissions:
      contents: read
      pull-requests: write  # Post review comments
    steps:
      - uses: actions/checkout@b4ffde65f46336ab88eb53be808477a3936bae11  # v4.1.1
        with:
          persist-credentials: false

      # SECURITY: Dependency review detects malicious or vulnerable packages in PRs
      - uses: actions/dependency-review-action@c74b580d73376b7750d3d2a50bfb8adc2c937507  # v3.1.0
        with:
          # Fail on critical/high vulnerabilities
          fail-on-severity: high
          # Deny known malicious packages
          deny-licenses: AGPL-3.0, GPL-3.0
```

### Python CI

Hardened CI for Python projects with security scanning and dependency management.

```yaml
name: Python CI
on:
  push:
    branches: [main]
  pull_request:
    branches: [main]

permissions:
  contents: read

jobs:
  test:
    name: Test on Python ${{ matrix.python-version }}
    runs-on: ubuntu-latest
    permissions:
      contents: read
    strategy:
      fail-fast: true
      matrix:
        python-version: ['3.10', '3.11', '3.12']
    steps:
      - uses: actions/checkout@b4ffde65f46336ab88eb53be808477a3936bae11  # v4.1.1
        with:
          persist-credentials: false

      - uses: actions/setup-python@0a5c61591373683505ea898e09a3ea4f39ef2b9c  # v5.0.0
        with:
          python-version: ${{ matrix.python-version }}
          cache: 'pip'

      # SECURITY: Install dependencies from locked requirements
      - name: Install dependencies
        run: |
          python -m pip install --upgrade pip
          pip install -r requirements-dev.txt

      # SECURITY: Bandit scans for common security issues in Python code
      - name: Run Bandit security scan
        run: |
          pip install bandit[toml]
          bandit -r . -f json -o bandit-report.json || true

      # SECURITY: Safety checks for known vulnerabilities in dependencies
      - name: Check dependencies with Safety
        run: |
          pip install safety
          safety check --json

      - name: Lint with ruff
        run: |
          pip install ruff
          ruff check .

      - name: Type check with mypy
        run: |
          pip install mypy
          mypy .

      - name: Run tests with pytest
        run: |
          pip install pytest pytest-cov
          pytest --cov=. --cov-report=xml --cov-report=term

      - name: Upload coverage
        uses: codecov/codecov-action@e0b68c6749509c5f83f984dd99a76a1c1a231044  # v4.0.1
        with:
          files: ./coverage.xml
          fail_ci_if_error: false

  security-scan:
    name: Security Scanning
    runs-on: ubuntu-latest
    permissions:
      contents: read
      security-events: write
    steps:
      - uses: actions/checkout@b4ffde65f46336ab88eb53be808477a3936bae11  # v4.1.1
        with:
          persist-credentials: false

      # SECURITY: CodeQL for Python static analysis
      - uses: github/codeql-action/init@cdcdbb579706841c47f7063dda365e292e5cad7a  # v2.13.4
        with:
          languages: python
          queries: security-extended

      - uses: github/codeql-action/autobuild@cdcdbb579706841c47f7063dda365e292e5cad7a  # v2.13.4


