---
name: release-please-configuration
description: >-
  Configure automated versioning with release-please. Set up semantic versioning, changelog generation, and monorepo support using GitHub App token authentication.
---

# Release-Please Configuration

## When to Use This Skill

Release-please reads your commit history and:

1. Groups changes by type (feat, fix, chore, etc.)
2. Generates changelogs
3. Bumps versions according to semantic versioning
4. Creates pull requests for releases
5. Tags releases when PRs merge

---


## Implementation

[Release-please](https://github.com/marketplace/actions/release-please-action) automates version management based on conventional commits. It creates release PRs with updated changelogs, version bumps, and Git tags.

> **Schema Validation**
>
> Always include the `$schema` property in your config file. It catches invalid options immediately and saves debugging time.
>

---

## Overview

Release-please reads your commit history and:

1. Groups changes by type (feat, fix, chore, etc.)
2. Generates changelogs
3. Bumps versions according to semantic versioning
4. Creates pull requests for releases
5. Tags releases when PRs merge

---

## Configuration Files

### release-please-config.json

The main configuration file defines packages and their versioning behavior:


*See [examples.md](examples.md) for detailed code examples.*

### .release-please-manifest.json

Tracks current versions for each package:

```json
{
  "charts/my-app": "1.0.0",
  "packages/backend": "1.0.0",
  "packages/frontend": "1.0.0"
}
```

---

## Configuration Options

### Global Options

| Option | Description | Example |
| -------- | ------------- | --------- |
| `include-v-in-tag` | Prefix tags with `v` | `true` = `v1.0.0`, `false` = `1.0.0` |
| `tag-separator` | Separator between component and version | `-` = `backend-1.0.0` |
| `separate-pull-requests` | Create one PR per component | Recommended for monorepos |
| `changelog-sections` | How to group commits in changelogs | See example above |

### Package Options

| Option | Description | Values |
| -------- | ------------- | -------- |
| `release-type` | Package ecosystem | `node`, `helm`, `simple`, `python`, `go`, etc. |
| `component` | Component name for tagging | Any string |
| `include-component-in-tag` | Include component in tag | `true` = `backend-1.0.0` |
| `package-name` | Package name (for node, etc.) | Matches package.json name |

---

## Schema Validation

Always validate configuration against the official schema:

```json
{
  "$schema": "https://raw.githubusercontent.com/googleapis/release-please/main/schemas/config.json"
}
```

This catches invalid options immediately. Options like `release-name` don't exist. The schema prevents wasted debugging time.

---

## In This Section

- [Release Types](release-types.md) - Node, Helm, Simple, and changelog customization
- [Extra-Files](extra-files.md) - Version tracking in arbitrary files
- [Workflow Integration](workflow-integration.md) - GitHub Actions setup and outputs
- [Troubleshooting](troubleshooting.md) - Common issues and solutions

---

## Related

- [Change Detection](../change-detection.md) - Skip unnecessary builds
- [Workflow Triggers](../workflow-triggers.md) - GITHUB_TOKEN compatibility
- [Content Comparison](../../../patterns/github-actions/use-cases/work-avoidance/content-comparison.md) - Skip version-only changes

---

## References

- [Release-please Action](https://github.com/marketplace/actions/release-please-action) - GitHub Marketplace
- [Release-please Repository](https://github.com/googleapis/release-please) - googleapis
- [Manifest Releaser Documentation](https://github.com/googleapis/release-please/blob/main/docs/manifest-releaser.md) - Monorepo configuration
- [Configuration Schema](https://raw.githubusercontent.com/googleapis/release-please/main/schemas/config.json) - JSON schema for validation

### Overview

Release-please reads your commit history and:

1. Groups changes by type (feat, fix, chore, etc.)
2. Generates changelogs
3. Bumps versions according to semantic versioning
4. Creates pull requests for releases
5. Tags releases when PRs merge

---

### Configuration Files

### release-please-config.json

The main configuration file defines packages and their versioning behavior:


*See [examples.md](examples.md) for detailed code examples.*

### .release-please-manifest.json

Tracks current versions for each package:

```json
{
  "charts/my-app": "1.0.0",
  "packages/backend": "1.0.0",
  "packages/frontend": "1.0.0"
}
```

---

### Configuration Options

### Global Options

| Option | Description | Example |
| -------- | ------------- | --------- |
| `include-v-in-tag` | Prefix tags with `v` | `true` = `v1.0.0`, `false` = `1.0.0` |
| `tag-separator` | Separator between component and version | `-` = `backend-1.0.0` |
| `separate-pull-requests` | Create one PR per component | Recommended for monorepos |
| `changelog-sections` | How to group commits in changelogs | See example above |

### Package Options

| Option | Description | Values |
| -------- | ------------- | -------- |
| `release-type` | Package ecosystem | `node`, `helm`, `simple`, `python`, `go`, etc. |
| `component` | Component name for tagging | Any string |
| `include-component-in-tag` | Include component in tag | `true` = `backend-1.0.0` |
| `package-name` | Package name (for node, etc.) | Matches package.json name |

---

### Schema Validation

Always validate configuration against the official schema:

```json
{
  "$schema": "https://raw.githubusercontent.com/googleapis/release-please/main/schemas/config.json"
}
```

This catches invalid options immediately. Options like `release-name` don't exist. The schema prevents wasted debugging time.

---

### In This Section

- [Release Types](release-types.md) - Node, Helm, Simple, and changelog customization
- [Extra-Files](extra-files.md) - Version tracking in arbitrary files
- [Workflow Integration](workflow-integration.md) - GitHub Actions setup and outputs
- [Troubleshooting](troubleshooting.md) - Common issues and solutions

---

### Related

- [Change Detection](../change-detection.md) - Skip unnecessary builds
- [Workflow Triggers](../workflow-triggers.md) - GITHUB_TOKEN compatibility
- [Content Comparison](../../../patterns/github-actions/use-cases/work-avoidance/content-comparison.md) - Skip version-only changes

---

### References

- [Release-please Action](https://github.com/marketplace/actions/release-please-action) - GitHub Marketplace
- [Release-please Repository](https://github.com/googleapis/release-please) - googleapis
- [Manifest Releaser Documentation](https://github.com/googleapis/release-please/blob/main/docs/manifest-releaser.md) - Monorepo configuration
- [Configuration Schema](https://raw.githubusercontent.com/googleapis/release-please/main/schemas/config.json) - JSON schema for validation


## Examples

See [examples.md](examples.md) for code examples.


## Related Patterns

- Change Detection
- Workflow Triggers
- Content Comparison

## References

- [Source Documentation](https://adaptive-enforcement-lab.com/build/release-pipelines/)
- [AEL Build](https://adaptive-enforcement-lab.com/build/)
