---
name: secret-scanning-integration - Reference
description: Complete reference for Secret Scanning Integration
---

# Secret Scanning Integration - Reference

This is the complete reference documentation extracted from the source.


# Secret Scanning Integration

Prevention is good. Detection is essential. Assume secrets will leak. Build systems to catch them before damage spreads.

> **The Risk**
>
>
> Secrets leak through commits, workflow logs, artifacts, pull requests, and third-party integrations. Without automated scanning, credentials remain exposed for days or months before detection. By then, your infrastructure is already compromised.
>

## What is Secret Scanning?

GitHub secret scanning automatically detects known secret formats in repositories, workflow logs, and commit history.

**How It Works**:

1. GitHub scans commits, issues, pull requests, and workflow logs
2. Pattern matching identifies known credential formats (API keys, tokens, certificates)
3. Alerts sent to repository administrators and security team
4. Optional push protection blocks commits containing secrets
5. Partner notification for compromised service provider credentials

**Coverage**:

- **Repository scanning**: All commits, branches, and history
- **Push protection**: Block secret commits before they land
- **Workflow logs**: Scan job output for leaked credentials
- **Pull requests**: Scan fork contributions for secret exposure
- **Partner patterns**: 200+ service providers receive breach notifications

## Enabling Secret Scanning

Secret scanning availability depends on repository visibility and GitHub plan.

### Repository Settings

**GitHub Advanced Security (GHAS) Required**:

- Private/internal repositories: GHAS license required
- Public repositories: Free, enabled by default

**Enable via Settings**:

1. Navigate to `Settings → Code security and analysis`
2. Enable **Secret scanning**
3. Enable **Push protection** (recommended)
4. Enable **Non-provider patterns** for generic secrets

```yaml
# .github/workflows/verify-security.yml
# Workflow to enforce security features are enabled

name: Verify Security Configuration
on:
  schedule:
    - cron: '0 8 * * 1'  # Weekly Monday 8 AM
  workflow_dispatch:

permissions:
  contents: read

jobs:
  check-scanning:
    runs-on: ubuntu-latest
    steps:
      - name: Check secret scanning enabled
        uses: actions/github-script@60a0d83039c74a4aee543508d2ffcb1c3799cdea  # v7.0.1
        with:
          script: |
            const { data: repo } = await github.rest.repos.get({
              owner: context.repo.owner,
              repo: context.repo.repo
            });

            const required = [
              { setting: 'security_and_analysis.secret_scanning.status', name: 'Secret Scanning' },
              { setting: 'security_and_analysis.secret_scanning_push_protection.status', name: 'Push Protection' }
            ];

            for (const check of required) {
              const value = check.setting.split('.').reduce((o, k) => o?.[k], repo);
              if (value !== 'enabled') {
                core.setFailed(`${check.name} is not enabled (status: ${value})`);
              } else {
                core.info(`✓ ${check.name} enabled`);
              }
            }
```

### Organization-Level Enablement

Enable secret scanning across all repositories in organization.

**Organization Settings**:

1. Navigate to `Organization Settings → Code security and analysis`
2. Enable **Secret scanning** for all repositories
3. Enable **Push protection** organization-wide
4. Configure **Custom patterns** for org-specific secrets

**Enforcement via API**:

```bash
#!/bin/bash
# enable-secret-scanning.sh
# Enable secret scanning and push protection for all org repos

ORG="your-org"
TOKEN="${GITHUB_TOKEN}"

# Get all repositories in organization
repos=$(gh api \
  --paginate \
  "/orgs/${ORG}/repos" \
  --jq '.[].name')

for repo in $repos; do
  echo "Enabling secret scanning for ${ORG}/${repo}..."

  # Enable secret scanning
  gh api \
    --method PATCH \
    "/repos/${ORG}/${repo}" \
    -f security_and_analysis[secret_scanning][status]=enabled \
    -f security_and_analysis[secret_scanning_push_protection][status]=enabled

  echo "✓ ${repo} configured"
done
```

**Best Practices**:

- Enable organization-wide by default
- Require for all new repositories
- Audit compliance weekly
- Block repository creation without security features

## Push Protection

Push protection blocks commits containing secrets before they reach GitHub.

### How Push Protection Works

```mermaid
sequenceDiagram

%% Ghostty Hardcore Theme
    participant Dev as Developer
    participant Git as Git Client
    participant GH as GitHub
    participant Scan as Secret Scanner

    Dev->>Git: git push origin main
    Git->>GH: Push commit
    GH->>Scan: Scan commit contents

    alt Secret Detected
        Scan-->>GH: Secret found (API key pattern)
        GH-->>Git: ❌ Push rejected
        Git-->>Dev: Error: secret detected<br/>Remove secret and retry
    else No Secret
        Scan-->>GH: No secrets found
        GH-->>Git: ✓ Push accepted
        Git-->>Dev: Push successful
    end

```

### Bypassing Push Protection

Developers can bypass push protection for false positives. Track bypasses for security review.

**Bypass Process**:

1. Developer attempts push with secret
2. GitHub blocks push with bypass option
3. Developer provides justification (false positive, test credential, will revoke)
4. Push allowed with bypass event logged
5. Security team reviews bypass audit trail

**Bypass Workflow**:

```bash
# Developer pushes commit with secret
git push origin main
# > Error: secret detected in commit abc123
# > To bypass, visit: https://github.com/org/repo/security/bypass/abc123

# Developer bypasses with justification
# GitHub logs bypass event

# Security team reviews bypasses
gh api /repos/org/repo/secret-scanning/push-protection-bypasses
```

**Monitor Bypasses**:

```yaml
# .github/workflows/monitor-bypasses.yml
# Alert security team when push protection bypassed

name: Monitor Push Protection Bypasses
on:
  schedule:
    - cron: '0 */4 * * *'  # Every 4 hours
  workflow_dispatch:

permissions:
  contents: read

jobs:
  check-bypasses:
    runs-on: ubuntu-latest
    steps:
      - name: Get recent bypasses
        uses: actions/github-script@60a0d83039c74a4aee543508d2ffcb1c3799cdea  # v7.0.1
        with:
          script: |
            const bypasses = await github.paginate(
              github.rest.secretScanning.listPushProtectionBypasses,
              {
                owner: context.repo.owner,
                repo: context.repo.repo
              }
            );

            const recent = bypasses.filter(b => {
              const created = new Date(b.created_at);
              const fourHoursAgo = new Date(Date.now() - 4 * 60 * 60 * 1000);
              return created > fourHoursAgo;
            });

            if (recent.length > 0) {
              core.warning(`${recent.length} push protection bypasses in last 4 hours`);
              for (const bypass of recent) {
                core.warning(`Bypass by ${bypass.pusher.login}: ${bypass.token_type}`);
              }
              // Trigger alert to security team (Slack, PagerDuty, etc.)
            }
```

## Custom Pattern Definitions

Define organization-specific secret patterns for internal credentials.

### Creating Custom Patterns

Custom patterns use regular expressions to detect organization-specific secrets.

**Pattern Format**:

```regex
# Pattern components
(?i)                           # Case insensitive
\b                             # Word boundary
(internal_api_key|secret_key)  # Secret identifier
[\s:=]+                        # Separator
([a-f0-9]{64})                 # Secret value pattern
\b                             # Word boundary
```

**Organization-Level Pattern**:

1. Navigate to `Organization Settings → Code security → Secret scanning`
2. Click **New pattern**
3. Define pattern name and regular expression
4. Test against sample secrets
5. Enable for all or selected repositories

**Example Custom Patterns**:

| Secret Type | Pattern | Example Match |
| ----------- | ------- | ------------- |
| Internal API Key | `(?i)\b(internal_api_key\s*[:=]\s*)([a-f0-9]{64})\b` | `INTERNAL_API_KEY=a1b2c3d4...` |
| Service Token | `(?i)\bSVC_TOKEN_([A-Z0-9]{32})\b` | `SVC_TOKEN_AB12CD34EF56GH78...` |
| Database URL | `(?i)postgresql://[^:]+:[^@]+@[^/]+/\w+` | `postgresql://user:pass@host/db` |
| SSH Private Key | `-----BEGIN (RSA\|OPENSSH) PRIVATE KEY-----` | `-----BEGIN RSA PRIVATE KEY-----` |

### Pattern Best Practices

**Effective Patterns**:

- Use word boundaries `\b` to avoid substring false positives
- Include secret identifier context (e.g., `API_KEY=`)
- Match actual secret format (length, character class)
- Test against real examples before deployment
- Document pattern purpose and maintenance owner

**Avoid Common Mistakes**:

- Too broad: `[a-z0-9]+` (matches everything)
- Too narrow: `PROD_KEY=abc123` (only matches one value)
- Missing boundaries: `password.*` (matches variable names)
- No context: `[a-f0-9]{64}` (many false positives)

**Test Pattern**:

```bash
# Test custom pattern against sample file
echo "INTERNAL_API_KEY=a1b2c3d4e5f6..." > test-secret.txt

# GitHub CLI test (pattern must be created first)
gh secret-scanning list --repo org/repo

# Local regex test
grep -P '(?i)\b(internal_api_key\s*[:=]\s*)([a-f0-9]{64})\b' test-secret.txt
```

## Secret Scanning Alerts

Alerts notify repository administrators when secrets are detected.

### Alert Triage Workflow

```mermaid
flowchart TD
    Alert["Secret Detected"] --> Verify["1. Verify Alert<br/>Real or False Positive?"]

    Verify -->|Real Secret| Classify["2. Classify Severity<br/>Production or Test?"]
    Verify -->|False Positive| Dismiss["Dismiss as<br/>False Positive"]

    Classify -->|Production Credential| Critical["🔴 CRITICAL<br/>Immediate Response"]
    Classify -->|Test/Dev Credential| Medium["🟡 MEDIUM<br/>Standard Response"]

    Critical --> Revoke1["3a. Revoke credential<br/>(within 15 minutes)"]
    Medium --> Revoke2["3b. Revoke credential<br/>(within 24 hours)"]

    Revoke1 --> Rotate1["4a. Rotate credential<br/>Update GitHub secret"]
    Revoke2 --> Rotate2["4b. Rotate credential<br/>Update GitHub secret"]

    Rotate1 --> Clean1["5a. Remove from history<br/>BFG Repo-Cleaner"]
    Rotate2 --> Clean2["5b. Remove from history<br/>Git filter-branch"]

    Clean1 --> Document1["6. Document incident"]
    Clean2 --> Document2["6. Document incident"]

    %% Ghostty Hardcore Theme
    style Alert fill:#66d9ef,color:#1b1d1e
    style Critical fill:#f92572,color:#1b1d1e
    style Medium fill:#e6db74,color:#1b1d1e
    style Dismiss fill:#75715e,color:#f8f8f2

```

