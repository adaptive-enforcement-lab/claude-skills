---
name: prerequisite-checks - Examples
description: Code examples for Prerequisite Checks
---

# Prerequisite Checks - Examples


## Example 1: example-1.mermaid


```mermaid
flowchart LR
    subgraph prereq[Prerequisite Phase]
        A[Check Tools]
        B[Check Access]
        C[Check State]
        D[Check Resources]
    end

    subgraph gate[Gate]
        G{All Pass?}
    end

    subgraph exec[Execution Phase]
        E[Execute Operation]
    end

    A --> B --> C --> D --> G
    G -->|Yes| E
    G -->|No| F[Abort with Report]

    %% Ghostty Hardcore Theme
    style A fill:#65d9ef,color:#1b1d1e
    style B fill:#65d9ef,color:#1b1d1e
    style C fill:#65d9ef,color:#1b1d1e
    style D fill:#65d9ef,color:#1b1d1e
    style G fill:#fd971e,color:#1b1d1e
    style E fill:#a7e22e,color:#1b1d1e
    style F fill:#f92572,color:#1b1d1e
```



## Example 2: example-2.yaml


```yaml
# GitHub Actions prerequisite check
- name: Validate prerequisites
  run: |
    errors=()

    # Environment
    [[ -n "${{ secrets.DEPLOY_TOKEN }}" ]] || errors+=("DEPLOY_TOKEN not set")

    # Tools
    command -v kubectl >/dev/null || errors+=("kubectl not installed")

    # Permissions
    kubectl auth can-i create deployments -n production || errors+=("No deploy permission")

    # State
    kubectl get namespace production >/dev/null || errors+=("Namespace missing")

    # Report
    if [ ${#errors[@]} -gt 0 ]; then
      echo "::error::Prerequisite check failed"
      printf '%s\n' "${errors[@]}"
      exit 1
    fi

    echo "All prerequisites met"
```



