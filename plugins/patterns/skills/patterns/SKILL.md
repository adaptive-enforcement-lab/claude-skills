---
name: patterns
description: >-
  Use when designing or reviewing automation architecture for GitHub Actions, Argo Workflows, or Argo Events — choosing an idempotency or work-avoidance strategy, wiring GitHub App authentication, or picking an error-handling or hub-and-spoke pattern.
---

# Patterns

## Overview

Reusable design patterns for resilient automation.

## Architecture Patterns

Fundamental patterns for building maintainable, scalable systems: separation…
- [Hub and Spoke](https://adaptive-enforcement-lab.com/patterns/architecture/hub-and-spoke/) — Centralized orchestration with distributed execution.
- [Matrix Distribution](https://adaptive-enforcement-lab.com/patterns/architecture/matrix-distribution/) — Parallelize operations across dynamic target lists using GitHub…
- [Separation of Concerns Pattern Overview](https://adaptive-enforcement-lab.com/patterns/architecture/separation-of-concerns/) — Single-responsibility components with clear boundaries.
- [Strangler Fig](https://adaptive-enforcement-lab.com/patterns/architecture/strangler-fig/) — Incremental migration from legacy systems.

## Argo Events

Build event-driven Kubernetes automation with Argo Events.
- [Argo Events Setup Guide](https://adaptive-enforcement-lab.com/patterns/argo-events/setup/) — Deploy event-driven automation with EventSource, EventBus, and Sensor…
- [Event Routing](https://adaptive-enforcement-lab.com/patterns/argo-events/routing/) — Control event flow from EventSources to triggers.
- [Reliability Patterns](https://adaptive-enforcement-lab.com/patterns/argo-events/reliability/) — Build resilient event systems with retry strategies, dead…
- [Troubleshooting](https://adaptive-enforcement-lab.com/patterns/argo-events/troubleshooting/) — Debug event flows systematically from EventSource to Workflow.

## Argo Workflows Patterns

Production Argo Workflows patterns: reusable templates, error handling…
- [Concurrency Control](https://adaptive-enforcement-lab.com/patterns/argo-workflows/concurrency/) — Prevent workflow conflicts with mutex synchronization, semaphores for…
- [Scheduled Workflows](https://adaptive-enforcement-lab.com/patterns/argo-workflows/scheduled/) — CronWorkflow patterns for scheduled automation: time-based execution, concurrency…
- [Workflow Composition](https://adaptive-enforcement-lab.com/patterns/argo-workflows/composition/) — Build complex pipelines from reusable workflow components.
- [WorkflowTemplate Patterns](https://adaptive-enforcement-lab.com/patterns/argo-workflows/templates/) — WorkflowTemplate foundations: versioned, reusable automation building blocks with…

## Efficiency Patterns

Optimize automation with idempotency and work avoidance.
- [Idempotency](https://adaptive-enforcement-lab.com/patterns/efficiency/idempotency/) — Build automation that survives reruns.
- [Implementation Patterns](https://adaptive-enforcement-lab.com/patterns/efficiency/idempotency/patterns/) — Five idempotency patterns for automation: check-before-act, upsert, force…
- [Tombstone/Marker Files](https://adaptive-enforcement-lab.com/patterns/efficiency/idempotency/patterns/tombstone-markers/) — Leave markers indicating operations completed.
- [Work Avoidance](https://adaptive-enforcement-lab.com/patterns/efficiency/work-avoidance/) — Skip work when outcomes won't change.
- [Work Avoidance Techniques](https://adaptive-enforcement-lab.com/patterns/efficiency/work-avoidance/techniques/) — Layer work avoidance checks from existence to content…

## Error Handling Patterns

Master when to fail fast vs degrade gracefully.
- [Fail Fast](https://adaptive-enforcement-lab.com/patterns/error-handling/fail-fast/) — Detect and halt on precondition failures before expensive…
- [Graceful Degradation](https://adaptive-enforcement-lab.com/patterns/error-handling/graceful-degradation/) — Build tiered fallback systems that degrade performance, not…
- [Prerequisite Checks](https://adaptive-enforcement-lab.com/patterns/error-handling/prerequisite-checks/) — Consolidate all precondition validation into a dedicated gate…

## Github Actions
- [Actions Integration](https://adaptive-enforcement-lab.com/patterns/github-actions/actions-integration/) — Integrate GitHub Core Apps with Actions workflows for…
- [Error Handling](https://adaptive-enforcement-lab.com/patterns/github-actions/actions-integration/error-handling/) — Handle token failures, API rate limits, and permission…
- [File Distribution](https://adaptive-enforcement-lab.com/patterns/github-actions/use-cases/file-distribution/) — Automated file distribution across multiple repositories with three-stage…
- [Installation Token Generation](https://adaptive-enforcement-lab.com/patterns/github-actions/actions-integration/token-generation/) — Generate short-lived installation tokens from GitHub App credentials…
- [JWT Authentication](https://adaptive-enforcement-lab.com/patterns/github-actions/actions-integration/jwt-authentication/) — Generate JWTs for GitHub App authentication.
- [Matrix Filtering and Deduplication](https://adaptive-enforcement-lab.com/patterns/github-actions/use-cases/work-avoidance/matrix-patterns/) — Reduce matrix builds from 47 jobs to 3…
- [OAuth User Authentication](https://adaptive-enforcement-lab.com/patterns/github-actions/actions-integration/oauth-authentication/) — OAuth flows for user-context operations.
- [Token Lifecycle Management](https://adaptive-enforcement-lab.com/patterns/github-actions/actions-integration/token-lifecycle/) — Token expiration handling, refresh strategies, and caching patterns…
- [Work Avoidance in GitHub Actions](https://adaptive-enforcement-lab.com/patterns/github-actions/use-cases/work-avoidance/) — Skip unnecessary CI/CD operations before execution.

## Reliability
- [Chaos Engineering for Kubernetes](https://adaptive-enforcement-lab.com/patterns/reliability/chaos-engineering/) — Chaos engineering for Kubernetes with Chaos Mesh and…
- [Chaos Experiment Design](https://adaptive-enforcement-lab.com/patterns/reliability/chaos-engineering/experiment-design/) — Chaos experiment design methodology.

## Security
- [Secure-by-Design Pattern Library](https://adaptive-enforcement-lab.com/patterns/security/secure-by-design/) — Secure-by-design architecture patterns for Kubernetes.

## Full Reference

See [reference.md](reference.md) for the complete content behind every link above, or [the live docs](https://adaptive-enforcement-lab.com/patterns/) on adaptive-enforcement-lab.com.
