---
name: modular-release-pipelines
description: >-
  Automate version management and changelog generation with smart builds. Only build changed components using GitHub App tokens and release-please integration.
---

# Modular Release Pipelines

## When to Use This Skill

This guide covers implementing release automation with:

- **Release-please** for version bumping and changelog generation
- **GitHub App authentication** for proper workflow triggering
- **Change detection** to skip unnecessary builds
- **Cascade rebuilds** when shared dependencies change

```mermaid
flowchart LR
    subgraph release[Release Pipeline]
        Main[Push to Main] --> AppToken[GitHub App Token]
        AppToken --> RP[Release-Please]
        RP --> PR[Creates PR]
    end

    subgraph build[Build Pipeline]
        PR -->|pull_request event| DC[Detect Changes]
        DC --> Test[Test]
        Test --> Build[Build]
        Build --> Status[Build Status]
    end

    %% Ghostty Hardcore Theme
    style Main fill:#65d9ef,color:#1b1d1e
    style AppToken fill:#9e6ffe,color:#1b1d1e
    style RP fill:#9e6ffe,color:#1b1d1e
    style PR fill:#fd971e,color:#1b1d1e
    style DC fill:#fd971e,color:#1b1d1e
    style Test fill:#9e6ffe,color:#1b1d1e
    style Build fill:#a7e22e,color:#1b1d1e
    style Status fill:#5e7175,color:#f8f8f3

```

---


## Prerequisites

Before implementing release pipelines, set up a GitHub App for your organization:

- [GitHub App Setup](../../secure/github-apps/index.md) - Create and configure the App
- [Token Generation](../../patterns/github-actions/actions-integration/token-generation/index.md) - Generate tokens in workflows

---


## Implementation


See the full implementation guide in the source documentation.


## Examples

See [examples.md](examples.md) for code examples.






## References

- [Source Documentation](https://adaptive-enforcement-lab.com/build/release-pipelines/)
- [AEL Build](https://adaptive-enforcement-lab.com/build/)
