---
name: github-actions-security-cheat-sheet - Examples
description: Code examples for GitHub Actions Security Cheat Sheet
---

# GitHub Actions Security Cheat Sheet - Examples


## Example 1: example-1.yaml


```yaml
steps:
  # ✅ GOOD: SHA pinned with version comment
  - uses: actions/checkout@b4ffde65f46336ab88eb53be808477a3936bae11  # v4.1.1
  - uses: actions/setup-node@5e21ff4d9bc1a8cf6de233a3057d20ec6b3fb69d  # v3.8.1

  # ❌ BAD: Mutable tag reference
  # - uses: actions/checkout@v4
```



## Example 2: example-2.yaml


```yaml
version: 2
updates:
  - package-ecosystem: "github-actions"
    directory: "/"
    schedule:
      interval: "weekly"
    groups:
      # Group GitHub-maintained actions
      github-actions-core:
        patterns:
          - "actions/*"
          - "github/*"
```



## Example 3: example-3.yaml


```yaml
name: Secure CI
on: [push, pull_request]

# Workflow-level: deny most access
permissions:
  contents: read

jobs:
  test:
    runs-on: ubuntu-latest
    # Job inherits workflow-level permissions
    steps:
      - uses: actions/checkout@b4ffde65f46336ab88eb53be808477a3936bae11  # v4.1.1
      - run: npm test

  publish:
    runs-on: ubuntu-latest
    # Job-level: override for specific needs
    permissions:
      contents: read
      packages: write  # Only this job can publish
    steps:
      - run: npm publish
```



## Example 4: example-4.yaml


```yaml
jobs:
  deploy:
    runs-on: ubuntu-latest
    permissions:
      id-token: write  # Request OIDC token
      contents: read
    environment: production  # Restrict trust to environment
    steps:
      # ✅ GOOD: No stored credentials
      - uses: google-github-actions/auth@55bd3a7c6e2ae7cf1877fd1ccb9d54c0503c457c  # v2.1.2
        with:
          workload_identity_provider: 'projects/123/locations/global/workloadIdentityPools/github/providers/github'
          service_account: 'deploy@project.iam.gserviceaccount.com'

      # ❌ BAD: Stored service account key
      # - run: echo "${{ secrets.GCP_SA_KEY }}" | base64 -d > key.json
```



## Example 5: example-5.yaml


```yaml
# .github/workflows/secret-scan.yml
name: Secret Scanning
on: [push, pull_request]

jobs:
  scan:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@b4ffde65f46336ab88eb53be808477a3936bae11  # v4.1.1
        with:
          fetch-depth: 0  # Full history for scanning

      - uses: gitleaks/gitleaks-action@cb7149a9a1d86f1c2e3ab90ae2f43a75ed56e95a  # v2.3.2
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
```



## Example 6: example-6.yaml


```yaml
# Example policy: Tier 1 + Tier 2 only
Allowed actions and reusable workflows:
  - Allow actions created by GitHub: ✅
  - Allow actions by Marketplace verified creators: ✅
  - Allow specified actions and reusable workflows:
      - aquasecurity/trivy-action@*
      - google-github-actions/*@*
```



## Example 7: example-7.sh


```bash
#!/bin/bash
# Podman ephemeral runner with strict isolation
podman run --rm \
  --security-opt=no-new-privileges:true \
  --cap-drop=ALL \
  --read-only \
  --tmpfs /tmp:rw,noexec,nosuid,size=1g \
  --network=slirp4netns:enable_ipv6=false \
  -e RUNNER_EPHEMERAL=true \
  -e GITHUB_TOKEN="${GITHUB_TOKEN}" \
  ghcr.io/myorg/runner:latest
```



## Example 8: example-8.sh


```bash
# Example: API-based runner group configuration
gh api --method PUT \
  /orgs/ORG/actions/runner-groups/GROUP_ID \
  -f name='production-runners' \
  -f visibility='selected' \
  -F selected_repository_ids='[123,456]' \
  -f allows_public_repositories=false \
  -f restricted_to_workflows=true \
  -F selected_workflows='[".github/workflows/deploy.yml@refs/heads/main"]'
```



## Example 9: example-9.yaml


```yaml
# .github/workflows/ci.yml
name: CI
on:
  pull_request:  # Safe for untrusted code
    branches: [main]

permissions:
  contents: read  # Read-only access

jobs:
```



