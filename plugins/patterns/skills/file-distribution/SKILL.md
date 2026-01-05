---
name: file-distribution
description: >-
  Automated file distribution across multiple repositories with three-stage workflow. Discovery, parallel distribution, and summary reporting for consistent updates.
---

# File Distribution

## When to Use This Skill

Automated file distribution across multiple repositories using GitHub Actions and GitHub Apps.

> **Pattern Overview**
>
> A three-stage workflow that discovers targets, distributes files in parallel, and reports results. Idempotent design ensures safe reruns.


## Prerequisites

- [GitHub App Setup](../../../../secure/github-apps/index.md) - Organization-level GitHub App
- [Actions Integration](../../actions-integration/index.md) - Token generation in workflows


## Implementation

### Core Workflow

- [Architecture](architecture.md) - Three-stage workflow overview
- [Stage 1: Discovery](discovery-stage.md) - Query organization for target repositories
- [Stage 2: Distribution](distribution-stage.md) - Parallel distribution to each repository
- [Stage 3: Summary](summary-stage.md) - Aggregate and display results

### Configuration

- [Workflow Configuration](workflow-config.md) - Triggers and permissions
- [Supporting Scripts](supporting-scripts.md) - Branch preparation and helper scripts

### Reliability

- [Idempotency](idempotency.md) - Safe re-execution guarantees
- [Error Handling](error-handling.md) - Failure strategies and reporting
- [Troubleshooting](troubleshooting.md) - Common issues and solutions

### Extensions

- [Extension Patterns](extension-patterns.md) - Multi-file, conditional, and template distribution

### Operations

- [Performance](performance.md) - Parallel processing and rate limits
- [Monitoring](monitoring.md) - Workflow summaries and metrics
- [Security](security.md) - Token scope and audit trails


## Key Principles

1. **Start Small** - Test with 2-3 repositories before full rollout
2. **Monitor First Run** - Watch logs carefully on initial deployment
3. **Gradual Rollout** - Increase `max-parallel` gradually
4. **Clear Documentation** - Document what files are distributed and why
5. **Review Process** - Ensure PRs are reviewed before merging


## Techniques


### Patterns Applied

This workflow implements patterns from the [Developer Guide](../../../../patterns/index.md):

| Pattern | Purpose |
| ------- | ------- |
| [Three-Stage Design](../../../../patterns/architecture/three-stage-design.md) | Separates discovery, execution, and reporting |
| [Matrix Distribution](../../../../patterns/architecture/matrix-distribution/index.md) | Parallelizes operations with conditional logic |
| [Idempotency](../../../../patterns/efficiency/idempotency/index.md) | Ensures safe reruns after partial failures |
| [Work Avoidance](../../../../patterns/efficiency/work-avoidance/index.md) | Skips version-only changes |


## Troubleshooting

See [troubleshooting.md](troubleshooting.md) for common issues and solutions.
## References

- [Source Documentation](https://adaptive-enforcement-lab.com/patterns/github-actions/)
- [AEL Patterns](https://adaptive-enforcement-lab.com/patterns/)
