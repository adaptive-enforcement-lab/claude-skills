---
name: third-party-action-risk-assessment - Examples
description: Code examples for Third-Party Action Risk Assessment
---

# Third-Party Action Risk Assessment - Examples


## Example 1: example-1.mermaid


```mermaid
flowchart TD
    A["Third-Party Action"] --> B["Execution Context"]

    B --> C["Secrets Access"]
    B --> D["Code Access"]
    B --> E["Token Permissions"]

    C --> C1["GITHUB_TOKEN"]
    C --> C2["Cloud Credentials"]
    C --> C3["API Keys"]
    C --> C4["Deployment Tokens"]

    D --> D1["Source Code"]
    D --> D2["Git History"]
    D --> D3["Dependencies"]

    E --> E1["Write to Repo"]
    E --> E2["Create Releases"]
    E --> E3["Modify Workflows"]

    C1 --> F["Attack Vectors"]
    C2 --> F
    C3 --> F
    D1 --> F
    E3 --> F

    F --> G["Exfiltrate Secrets"]
    F --> H["Backdoor Codebase"]
    F --> I["Persistent Access"]
    F --> J["Supply Chain Compromise"]

    %% Ghostty Hardcore Theme
    style A fill:#f92572,color:#1b1d1e
    style B fill:#fd971e,color:#1b1d1e
    style F fill:#e6db74,color:#1b1d1e
    style G fill:#66d9ef,color:#1b1d1e
    style H fill:#66d9ef,color:#1b1d1e
    style I fill:#66d9ef,color:#1b1d1e
    style J fill:#66d9ef,color:#1b1d1e
```



## Example 2: example-2.mermaid


```mermaid
flowchart TD
    Start["New Action Needed"] --> Q1{"GitHub-Maintained<br/>(actions/*, github/*)?"}

    Q1 -->|Yes| Tier1["Tier 1: Low Risk"]
    Q1 -->|No| Q2{"Verified Publisher?"}

    Q2 -->|Yes| Tier2["Tier 2: Medium Risk"]
    Q2 -->|No| Q3{"Community Action<br/>with Good Signals?"}

    Q3 -->|Yes| Tier3["Tier 3: High Risk"]
    Q3 -->|No| Tier4["Tier 4: Critical Risk"]

    Tier1 --> A1["✓ SHA Pin<br/>✓ PR Review<br/>✓ Adopt"]

    Tier2 --> A2["✓ Source Review<br/>✓ Security Approval<br/>✓ SHA Pin<br/>✓ Adopt"]

    Tier3 --> A3["✓ Full Code Audit<br/>✓ Risk Assessment<br/>✓ Alternatives Evaluated"]
    A3 --> Q4{"Risk Acceptable?"}
    Q4 -->|Yes| A4["✓ Security Sign-Off<br/>✓ Monitoring Plan<br/>✓ Fork & Maintain"]
    Q4 -->|No| Reject["❌ Reject - Build Internal"]

    Tier4 --> A5["⚠ Block by Default"]
    A5 --> Q5{"Compelling Business Need?"}
    Q5 -->|Yes| A6["✓ Treat as Tier 3<br/>✓ Enhanced Scrutiny<br/>✓ Fork Required"]
    Q5 -->|No| Reject

    %% Ghostty Hardcore Theme
    style Start fill:#66d9ef,color:#1b1d1e
    style Tier1 fill:#a6e22e,color:#1b1d1e
    style Tier2 fill:#e6db74,color:#1b1d1e
    style Tier3 fill:#fd971e,color:#1b1d1e
    style Tier4 fill:#f92572,color:#1b1d1e
    style Reject fill:#f92572,color:#1b1d1e
    style A1 fill:#a6e22e,color:#1b1d1e
    style A4 fill:#e6db74,color:#1b1d1e
```



## Example 3: example-3.yaml


```yaml
# Bad - tag reference
- uses: community/action@v2

# Good - SHA pinned with version comment
- uses: community/action@a1b2c3d4e5f6...  # v2.1.0
```



