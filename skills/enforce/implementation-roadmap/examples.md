---
name: implementation-roadmap - Examples
description: Code examples for Implementation Roadmap
---

# Implementation Roadmap - Examples


## Example 1: example-1.sh


```bash
gh api repos/org/repo/branches/main/protection \
  | jq '{reviews: .required_pull_request_reviews, admins: .enforce_admins}'
```



## Example 2: example-2.yaml


```yaml
name: Required Checks
on: [pull_request]
jobs:
  tests:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - run: make test
  lint:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - run: make lint
```



## Example 3: example-3.yaml


```yaml
- name: Test app token
  uses: actions/create-github-app-token@v2
  with:
    app-id: ${{ secrets.APP_ID }}
    private-key: ${{ secrets.PRIVATE_KEY }}
```



## Example 4: example-4.yaml


```yaml
name: Monthly Evidence
on:
  schedule:
    - cron: '0 0 1 * *'
jobs:
  archive:
    runs-on: ubuntu-latest
    steps:
      - run: gh api repos/org/repo/branches/main/protection > config.json
      - run: gsutil cp *.json gs://audit-evidence/
```



## Example 5: example-5.yaml


```yaml
repos:
  - repo: https://github.com/trufflesecurity/trufflehog
    rev: v3.63.0
    hooks:
      - id: trufflehog
        entry: trufflehog filesystem --fail --no-update
```



## Example 6: example-6.sh


```bash
git config --global user.signingkey YOUR_GPG_KEY_ID
git config --global commit.gpgsign true
```



## Example 7: example-7.sh


```bash
git log --show-signature | grep "Good signature"
```



## Example 8: example-8.yaml


```yaml
- name: Generate SBOM
  uses: anchore/sbom-action@v0
  with:
    image: app:${{ github.sha }}
    format: cyclonedx-json
    output-file: sbom.json

- name: Upload SBOM
  uses: actions/upload-artifact@v4
  with:
    name: sbom
    path: sbom.json
```



## Example 9: example-9.yaml


```yaml
- name: Scan container
  run: |
    trivy image --severity HIGH,CRITICAL --exit-code 1 \
      gcr.io/project/app:${{ github.sha }}
```



## Example 10: example-10.yaml


```yaml
apiVersion: kyverno.io/v1
kind: ClusterPolicy
metadata:
  name: require-resource-limits
spec:
  validationFailureAction: Enforce
  rules:
    - name: check-limits
      match:
        resources:
          kinds: [Pod]
      validate:
        message: "CPU and memory limits required"
        pattern:
          spec:
            containers:
              - resources:
                  limits:
                    memory: "?*"
                    cpu: "?*"
```



## Example 11: example-11.sh


```bash
# Verify branch protection
gh api repos/org/repo/branches/main/protection

# Sample March PRs
gh api 'repos/org/repo/pulls?state=closed&base=main' \
  --jq '.[] | select(.merged_at | startswith("2025-03"))'

# Check signature coverage
./scripts/signature-coverage.sh 2025-03-01 2025-04-01
```



