---
name: required-status-checks - Examples
description: Code examples for Required Status Checks
---

# Required Status Checks - Examples


## Example 1: example-1.yaml


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



## Example 2: example-2.yaml


```yaml
required_status_checks:
  strict: true
  contexts:
    - "tests"
    - "security-scan"
    - "lint"
```



## Example 3: example-3.sh


```bash
# Get workflow runs for a PR
gh api repos/org/repo/actions/runs \
  --jq '.workflow_runs[] | select(.head_branch=="feature-branch") |
    {name: .name, conclusion, created_at}'
```



## Example 4: example-4.yaml


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



## Example 5: example-5.yaml


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



## Example 6: example-6.yaml


```yaml
lint:
  runs-on: ubuntu-latest
  steps:
    - uses: actions/checkout@v4
    - uses: golangci/golangci-lint-action@v4
      with:
        args: --timeout=5m
```



## Example 7: example-7.yaml


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



## Example 8: example-8.yaml


```yaml
forbidden-tech:
  runs-on: ubuntu-latest
  steps:
    - uses: actions/checkout@v4
    - name: Check for forbidden technologies
      run: |
        ./scripts/check-forbidden-tech.sh
```



## Example 9: example-9.yaml


```yaml
required_status_checks:
  strict: true
```



## Example 10: example-10.yaml


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



