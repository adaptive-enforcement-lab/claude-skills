---
name: ci-integration-automated-policy-enforcement
description: >-
  Block non-compliant manifests at merge time using policy-platform container in CI. Automate Kyverno validation across environments with zero configuration drift.
---

# CI Integration: Automated Policy Enforcement

## When to Use This Skill

CI integration enforces policies automatically in every pull request using the **same policy-platform container** developers run locally.

```mermaid
graph LR
    PR[Pull Request] --> ENV[Detect Environment]
    ENV --> LINT[Lint Values]
    LINT --> BUILD[Build Manifests]
    BUILD --> VAL[Validate Policies]
    VAL --> MERGE{All Pass?}
    MERGE -->|Yes| ALLOW[Allow Merge]
    MERGE -->|No| BLOCK[Block Merge]

    %% Ghostty Hardcore Theme
    style ALLOW fill:#a7e22e,color:#1b1d1e
    style BLOCK fill:#f92572,color:#1b1d1e

```

**Key Principle**: CI uses identical validation to local development. No surprises.

---


## Implementation

See the full implementation guide in the [source documentation](https://adaptive-enforcement-lab.com/enforce/policy-as-code/).


## Examples

See [examples.md](examples.md) for code examples.


## Full Reference

See [reference.md](reference.md) for complete documentation.
## References

- [Source Documentation](https://adaptive-enforcement-lab.com/enforce/policy-as-code/)
- [AEL Enforce](https://adaptive-enforcement-lab.com/enforce/)
