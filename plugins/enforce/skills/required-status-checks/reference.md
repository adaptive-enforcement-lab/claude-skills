---
name: required-status-checks - Reference
description: Complete reference for Required Status Checks
---

# Required Status Checks - Reference

This is the complete reference documentation extracted from the source.


# Required Status Checks

Branch protection can require CI checks to pass before merge.

This turns "you should test" into "untested code cannot merge."

---

## The Contract

GitHub won't allow merge until all required checks report success.

```yaml
# .github/workflows/required-checks.yml
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

## Automatic Audit Trail

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
# Get workflow runs for a PR
gh api repos/org/repo/actions/runs \
  --jq '.workflow_runs[] | select(.head_branch=="feature-branch") |
    {name: .name, conclusion, created_at}'
```

---

## Required Check Types

### Unit and Integration Tests

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

### Security Scanning

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

### Linting

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

### SBOM Generation

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

### Forbidden Technology Check

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

## Strict Mode

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

## Matrix Testing

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

## Next Steps

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

