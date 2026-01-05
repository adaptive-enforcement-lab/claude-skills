---
name: branch-protection-enforcement-patterns - Examples
description: Code examples for Branch Protection Enforcement Patterns
---

# Branch Protection Enforcement Patterns - Examples


## Example 1: example-1.mermaid


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



## Example 2: example-2.sh


```bash
gh api --method PUT \
  repos/org/repo/branches/main/protection \
  --input protection-config.json
```



