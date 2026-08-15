# Build

Source: https://adaptive-enforcement-lab.com/build/

Development tools and release processes.

> **Build with Intent**
>
>
> Building secure, tested, versioned software requires more than writing code. It requires **architecture**, **testing discipline**, **release automation**, and **documentation workflows** that scale from prototype to production.
>

## Overview

This section covers the development practices, tooling choices, and automation patterns that turn code into deployable, documented, versioned artifacts.

## What You'll Find Here

**[Go CLI Architecture](go-cli-architecture/index.md)**: Production-grade Go CLIs with Kubernetes integration, testing, and packaging (21 pages)

**[Coverage Patterns](coverage-patterns/coverage-patterns.md)**: Testing strategies and coverage enforcement without slowing development

**[Release Pipelines](release-pipelines/index.md)**: Automated releases with Release Please for conventional commits and semantic versioning (8 pages)

**[Versioned Documentation](versioned-docs/index.md)**: Multi-version docs with Mike to prevent user confusion

**[Documentation as Skills](documentation-as-skills/index.md)**: Compile documentation into Claude Code skills so patterns reach engineers inside the agent (5 pages)

## Integration with Secure and Enforce

Build processes integrate with security and enforcement:

1. **Build artifacts** (Build) → **Scan for vulnerabilities** ([Secure](../secure/index.md)) → **Block vulnerable images** ([Enforce](../enforce/index.md))
2. **Run tests** (Build) → **Enforce coverage** ([Enforce](../enforce/index.md)) → **Gate PR merge** ([Enforce](../enforce/index.md))
3. **Generate SBOM** ([Secure](../secure/index.md)) → **Attach to release** (Build) → **Require in deployment** ([Enforce](../enforce/index.md))
4. **Create release** (Build) → **Generate SLSA provenance** ([Enforce](../enforce/index.md)) → **Verify in deployment** ([Enforce](../enforce/index.md))

## Development Workflow

Typical development flow using these patterns:

```mermaid
graph TB
    subgraph dev[Development]
        A[Write Code]
        B[Pre-commit Hooks]
        C[Commit]
        D[Push to PR]
    end

    subgraph ci[CI Validation]
        E[Tests + Coverage]
        F[Security Scan]
        G[Status Checks Pass]
        H[Peer Review]
    end

    subgraph release[Release Automation]
        I[Merge to Main]
        J[Release Please PR]
        K[Merge Release PR]
        L[GitHub Release]
    end

    subgraph deploy[Deployment]
        M[Deploy Artifacts]
    end

    A --> B --> C --> D
    D --> E --> F --> G --> H
    H --> I
    I --> J --> K --> L
    L --> M

    %% Ghostty Hardcore Theme
    style A fill:#5e7175,color:#f8f8f3
    style B fill:#65d9ef,color:#1b1d1e
    style C fill:#a7e22e,color:#1b1d1e
    style D fill:#fd971e,color:#1b1d1e
    style E fill:#65d9ef,color:#1b1d1e
    style F fill:#f92572,color:#1b1d1e
    style G fill:#a7e22e,color:#1b1d1e
    style H fill:#fd971e,color:#1b1d1e
    style I fill:#a7e22e,color:#1b1d1e
    style J fill:#65d9ef,color:#1b1d1e
    style K fill:#fd971e,color:#1b1d1e
    style L fill:#9e6ffe,color:#1b1d1e
    style M fill:#f92572,color:#1b1d1e

```

## Related Content

- [Secure](../secure/index.md): Security scanning and SBOM generation
- [Enforce](../enforce/index.md): Testing enforcement and compliance
- [Patterns](../patterns/index.md): CI/CD patterns and architecture

## Tags

Browse all content tagged with ci-cd, automation, testing, and go on the [Tags](../tags.md) page.
