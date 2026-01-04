---
name: branch-protection-enforcement-patterns
description: >-
  Comprehensive branch protection configuration patterns with enforcement automation. Security tiers, IaC at scale, GitHub App enforcement, audit reporting, and bypass controls.
---

# Branch Protection Enforcement Patterns

## When to Use This Skill

```mermaid
graph TD
    T[Terraform Module] -->|Applies| BP[Branch Protection Rules]
    GA[GitHub App] -->|Monitors| BP
    GA -->|Detects| DRIFT[Configuration Drift]
    DRIFT -->|Triggers| REM[Automated Remediation]
    REM -->|Restores| BP
    BP -->|Enforces| PR[Pull Requests]
    PR -->|Generates| AUDIT[Audit Evidence]

    %% Ghostty Hardcore Theme
    style T fill:#a7e22e,color:#1b1d1e
    style GA fill:#65d9ef,color:#1b1d1e
    style DRIFT fill:#f92572,color:#1b1d1e
    style BP fill:#fd971e,color:#1b1d1e

```

**Key Components**:

- **Terraform modules** - Declare protection rules as code
- **GitHub Apps** - Monitor and enforce compliance organization-wide
- **Drift detection** - Identify unauthorized changes
- **Automated remediation** - Restore protection without manual intervention
- **Audit collection** - Capture evidence for compliance reporting

---


## Prerequisites

- GitHub organization with admin access
- Terraform or OpenTofu (for IaC deployment)
- GitHub App with appropriate permissions (for automated enforcement)
- Basic understanding of Git workflow and branch protection concepts

---


## Implementation


See the full implementation guide in the source documentation.


## Examples

See [examples.md](examples.md) for code examples.






## References

- [Source Documentation](https://adaptive-enforcement-lab.com/enforce/branch-protection/)
- [AEL Enforce](https://adaptive-enforcement-lab.com/enforce/)
