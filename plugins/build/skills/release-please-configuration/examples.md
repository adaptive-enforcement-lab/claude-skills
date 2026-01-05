---
name: release-please-configuration - Examples
description: Code examples for Release-Please Configuration
---

# Release-Please Configuration - Examples


## Example 1: example-1.json


```json
{
  "$schema": "https://raw.githubusercontent.com/googleapis/release-please/main/schemas/config.json",
  "include-v-in-tag": false,
  "tag-separator": "-",
  "changelog-sections": [
    { "type": "feat", "section": "Features" },
    { "type": "fix", "section": "Bug Fixes" },
    { "type": "perf", "section": "Performance" },
    { "type": "refactor", "section": "Code Refactoring" },
    { "type": "docs", "section": "Documentation", "hidden": true },
    { "type": "chore", "section": "Maintenance" },
    { "type": "test", "section": "Tests", "hidden": true },
    { "type": "ci", "section": "CI/CD", "hidden": true }
  ],
  "packages": {
    "charts/my-app": {
      "release-type": "helm",
      "component": "my-app",
      "include-component-in-tag": false
    },
    "packages/backend": {
      "release-type": "node",
      "component": "backend",
      "package-name": "my-backend",
      "include-component-in-tag": true
    },
    "packages/frontend": {
      "release-type": "node",
      "component": "frontend",
      "package-name": "my-frontend",
      "include-component-in-tag": true
    }
  },
  "separate-pull-requests": true
}
```



## Example 2: example-2.json


```json
{
  "charts/my-app": "1.0.0",
  "packages/backend": "1.0.0",
  "packages/frontend": "1.0.0"
}
```



## Example 3: example-3.json


```json
{
  "$schema": "https://raw.githubusercontent.com/googleapis/release-please/main/schemas/config.json"
}
```



