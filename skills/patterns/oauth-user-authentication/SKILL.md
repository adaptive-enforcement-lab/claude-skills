---
name: oauth-user-authentication
description: >-
  OAuth flows for user-context operations. Web application patterns, device flow for CLI tools, and token refresh strategies for GitHub Apps.
---

# OAuth User Authentication

## When to Use This Skill

OAuth authentication provides user-context access for GitHub Apps. It enables:

- **User attribution** - Actions appear as the user in audit logs
- **User permissions** - Respect individual user access levels
- **Personal repository access** - Access to user's private repositories
- **Interactive applications** - Web apps and CLI tools requiring user authorization
- **Long-lived sessions** - Tokens valid until revoked

> **OAuth Limitations**
>
>
> - Not suitable for automated workflows (no user present)
> - Requires user consent for each installation
> - Rate limits apply per user (5,000/hour)
> - More complex setup than installation tokens


## Implementation

See the full implementation guide in the [source documentation](https://adaptive-enforcement-lab.com/patterns/github-actions/).


## Techniques


### OAuth vs Other Methods

```mermaid
flowchart TD
    A["Need user context?"] --> B{"Who initiates<br/>the action?"}

    B -->|"Human user<br/>(web app, CLI)"| C["Use OAuth"]
    B -->|"Automated process<br/>(GitHub Actions)"| D["Use Installation Token"]

    C --> C1["User attribution required"]
    C --> C2["Personal repos access"]
    C --> C3["User-level permissions"]

    D --> D1["No user present"]
    D --> D2["Organization repos"]
    D --> D3["App-level permissions"]

    %% Ghostty Hardcore Theme
    style A fill:#515354,stroke:#ccccc7,stroke-width:2px,color:#ccccc7
    style B fill:#fd971e,stroke:#e6db74,stroke-width:2px,color:#1b1d1e
    style C fill:#a7e22e,stroke:#bded5f,stroke-width:2px,color:#1b1d1e
    style D fill:#f92572,stroke:#ff669d,stroke-width:2px,color:#1b1d1e
    style C1 fill:#515354,stroke:#ccccc7,stroke-width:1px,color:#ccccc7
    style C2 fill:#515354,stroke:#ccccc7,stroke-width:1px,color:#ccccc7
    style C3 fill:#515354,stroke:#ccccc7,stroke-width:1px,color:#ccccc7
    style D1 fill:#515354,stroke:#ccccc7,stroke-width:1px,color:#ccccc7
    style D2 fill:#515354,stroke:#ccccc7,stroke-width:1px,color:#ccccc7
    style D3 fill:#515354,stroke:#ccccc7,stroke-width:1px,color:#ccccc7
```

*See [reference.md](reference.md) for additional techniques and detailed examples.*


## Comparison

*See [examples.md](examples.md) for detailed code examples.*


## Examples

See [examples.md](examples.md) for code examples.


## Full Reference

See [reference.md](reference.md) for complete documentation.
## References

- [Source Documentation](https://adaptive-enforcement-lab.com/patterns/github-actions/)
- [AEL Patterns](https://adaptive-enforcement-lab.com/patterns/)
