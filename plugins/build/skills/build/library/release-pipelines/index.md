# Modular Release Pipelines

Source: https://adaptive-enforcement-lab.com/build/release-pipelines/

Automated version management, changelog generation, and optimized builds for monorepos.

> **Smart Builds**
>
> Only build what changed. GitHub App tokens trigger builds correctly. Release-please handles versions automatically.
>

---

## Overview

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

## The Problem

Traditional CI/CD pipelines rebuild everything on every commit. In a monorepo with multiple components, this means:

- Unnecessary compute time for unchanged components
- Longer feedback loops for developers
- Wasted resources on duplicate work

Additionally, release-please using the default `GITHUB_TOKEN` won't trigger build pipelines on its PRs -- a [GitHub security measure](https://docs.github.com/en/actions/security-guides/automatic-token-authentication#using-the-github_token-in-a-workflow) to prevent infinite loops.

---

## The Solution

A modular pipeline architecture that:

1. Uses a **GitHub App token** for release-please (triggers `pull_request` events correctly)
2. Detects which components changed
3. Only builds affected components
4. Automatically versions and releases based on commits

---

## Guides

| Guide | Description |
| ----- | ----------- |
| [Release-Please Configuration](release-please/index.md) | Setting up automated versioning with GitHub App |
| [Change Detection](change-detection.md) | Detecting and cascading changes |
| [Workflow Triggers](workflow-triggers.md) | GitHub App token vs GITHUB_TOKEN |
| [Protected Branches](protected-branches.md) | Working with branch protection rules |

---

## Prerequisites

Before implementing release pipelines, set up a GitHub App for your organization:

- [GitHub App Setup](../../secure/github-apps/index.md) - Create and configure the App
- [Token Generation](../../patterns/github-actions/actions-integration/token-generation/index.md) - Generate tokens in workflows

---

## Architecture

### Build Pipeline

Runs on pull requests (including release-please PRs with GitHub App token):

```mermaid
flowchart TD
    subgraph detect[Change Detection]
        Contracts[Contracts Changed?]
        Backend[Backend Changed?]
        Frontend[Frontend Changed?]
        Charts[Charts Changed?]
    end

    subgraph cascade[Cascade Logic]
        BNB[Backend Needs Build]
        FNB[Frontend Needs Build]
    end

    subgraph build[Conditional Build]
        Test[Test Node Packages]
        BB[Build Backend]
        BF[Build Frontend]
        HC[Helm Charts]
    end

    Contracts -->|yes| BNB
    Contracts -->|yes| FNB
    Backend -->|yes| BNB
    Frontend -->|yes| FNB

    BNB --> Test
    FNB --> Test
    BNB --> BB
    FNB --> BF
    Charts --> HC

    %% Ghostty Hardcore Theme
    style Contracts fill:#fd971e,color:#1b1d1e
    style Backend fill:#fd971e,color:#1b1d1e
    style Frontend fill:#fd971e,color:#1b1d1e
    style Charts fill:#fd971e,color:#1b1d1e
    style BNB fill:#a7e22e,color:#1b1d1e
    style FNB fill:#a7e22e,color:#1b1d1e
    style Test fill:#9e6ffe,color:#1b1d1e
    style BB fill:#a7e22e,color:#1b1d1e
    style BF fill:#a7e22e,color:#1b1d1e
    style HC fill:#bded5f,color:#1b1d1e

```

### Release Pipeline

Runs on main branch pushes:

```mermaid
flowchart LR
    Main[Push to Main] --> Token[Generate App Token]
    Token --> RP[Release Please]
    RP --> DC[Detect Changes]
    DC --> Test[Test]
    DC --> Build[Build]
    Build --> Scan[Security Scan]
    Scan --> Deploy[Deploy Signal]

    %% Ghostty Hardcore Theme
    style Main fill:#65d9ef,color:#1b1d1e
    style Token fill:#9e6ffe,color:#1b1d1e
    style RP fill:#9e6ffe,color:#1b1d1e
    style DC fill:#fd971e,color:#1b1d1e
    style Test fill:#9e6ffe,color:#1b1d1e
    style Build fill:#a7e22e,color:#1b1d1e
    style Scan fill:#f92572,color:#1b1d1e
    style Deploy fill:#e6db74,color:#1b1d1e

```

---

## Quick Start

1. [Set up GitHub App](../../secure/github-apps/index.md) for your organization
2. [Configure release-please](release-please/index.md) with App token
3. [Set up change detection](change-detection.md) for your components
4. [Handle protected branches](protected-branches.md) if applicable

---

## Related

- [GitHub App Setup](../../secure/github-apps/index.md) - Machine identity for automation
- [Idempotency Patterns](../../patterns/efficiency/idempotency/index.md) - Making reruns safe
- [Three-Stage Design](../../patterns/architecture/three-stage-design.md) - Complex workflows
