---
name: secure - Examples
description: Code examples for Secure
---

# Secure - Examples


## Example 1: example-1.yaml


```yaml
# .github/workflows/security-scan.yml
name: Security Scan
on:
  push:
  schedule:
    - cron: '0 2 * * *'  # Daily at 2 AM

jobs:
  scan:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - name: Scan dependencies
        run: trivy fs --scanners vuln .
      - name: Scan containers
        run: trivy image myapp:latest
```



## Example 2: example-2.yaml


```yaml
# Generate SBOM on every release
- name: Generate SBOM
  run: syft packages . -o spdx-json > sbom.spdx.json
- name: Attach to release
  run: gh release upload ${{ github.ref_name }} sbom.spdx.json
```



## Example 3: example-3.yaml


```yaml
# Secure authentication for cross-repo workflows
- uses: actions/create-github-app-token@v1
  id: app-token
  with:
    app-id: ${{ vars.APP_ID }}
    private-key: ${{ secrets.APP_PRIVATE_KEY }}
    owner: adaptive-enforcement-lab
- uses: actions/checkout@v4
  with:
    token: ${{ steps.app-token.outputs.token }}
```



