---
name: patterns - Reference
description: Complete reference for Patterns
---

# Patterns - Reference

This is the complete reference documentation extracted from the source.

# Patterns

Reusable design patterns for resilient automation.

> **Solve Once, Apply Everywhere**
>
>
> Good patterns solve problems once. Great patterns solve problems across technologies, languages, and platforms.
>

## Overview

This section documents **architecture patterns**, **efficiency patterns**, **error handling patterns**, and **technology-specific patterns** that apply from GitHub Actions to Argo Workflows to Kubernetes operators.

These patterns are battle-tested in production environments, proven to reduce operational toil, and designed to prevent the failure modes that break automation at scale.

## Pattern Categories

### Architecture Patterns

System design patterns for building maintainable, scalable automation.

#### Separation of Concerns

- Split functionality into distinct, composable modules
- Each component does one thing well
- Reduces coupling, improves testability
- Examples: Script-based workflow stages, modular Helm charts

#### Hub-and-Spoke

- Central hub distributes configuration to many spokes
- Enforces consistency across repositories or clusters
- Enables organization-wide policy updates
- Examples: GitHub Apps distributing files, ArgoCD ApplicationSets

#### Strangler Fig

- Incrementally replace legacy systems without Big Bang rewrites
- Route traffic to new implementation while old runs
- Gradual cutover reduces risk
- Examples: Feature flags, API gateway routing, Kubernetes Ingress

#### Three-Stage Design

- Discovery → Execution → Summary pattern for workflows
- Discovery: Find what needs processing
- Execution: Process items in parallel
- Summary: Aggregate results and report
- Examples: Multi-repo file distribution, vulnerability remediation

#### Matrix Distribution

- Parallel execution across multiple dimensions
- Conditional execution based on matrix values
- Template rendering for each matrix combination
- Examples: Multi-environment deployments, multi-arch builds

#### Environment Progression

- Sequential deployment through environments (dev → staging → prod)
- Automated promotion on success
- Rollback on failure
- Examples: GitOps progressive delivery, Argo Rollouts

### Efficiency Patterns

Patterns that reduce runtime, cost, and toil.

#### Idempotency

- Safe to run multiple times without side effects
- Enables retries without duplication
- Critical for automation reliability
- Techniques: Check-before-act, upsert, tombstone markers, unique identifiers

#### Work Avoidance

- Skip unnecessary work when outcomes are unchanged
- Compare content hashes, not timestamps
- Use existence checks before creation
- Examples: Content-based PR creation, cache-based skips, path filtering

### Error Handling Patterns

Patterns that make automation resilient to failure.

#### Fail Fast

- Detect errors early, before expensive operations
- Exit immediately on unrecoverable errors
- Provide clear error messages
- Examples: Prerequisite checks, input validation

#### Prerequisite Checks

- Verify all requirements before starting work
- Check in optimal order (cheap first, expensive last)
- Avoid partial execution states
- Examples: Dependency checks, permission validation

#### Graceful Degradation

- Continue with reduced functionality when components fail
- Fallback to alternative implementations
- Preserve critical path operations
- Examples: Secondary data sources, default values

### Technology-Specific Patterns

#### GitHub Actions Patterns

- **Actions Integration**: Token generation, permission patterns, error handling
- **File Distribution**: Three-stage design for cross-repo operations
- **Release Pipelines**: Release Please integration, change detection
- **Work Avoidance**: Matrix filtering, content comparison, cache strategies
- **Versioned Docs**: Mike integration, version strategies

**Content**: 60+ pages of GitHub Actions patterns

#### Argo Events Patterns

- **Event Routing**: Simple filtering, multi-trigger actions, transformation
- **Conditional Routing**: Route events based on payload content
- **Reliability**: Retry strategies, dead letter queues, backpressure handling
- **Setup and Operations**: EventSource, EventBus, Sensor configuration

**Content**: 12+ pages covering event-driven workflow patterns

#### Argo Workflows Patterns

- **WorkflowTemplate Patterns**: Basic structure, retry strategies, init containers, volumes
- **Concurrency Control**: Mutexes, semaphores, TTL strategies
- **Workflow Composition**: Spawning children, parallel execution, DAG orchestration
- **Scheduled Workflows**: CronWorkflow patterns, concurrency policies, GitHub integration

**Content**: 20+ pages covering orchestration patterns

## Cross-Cutting Concerns

Many patterns apply across multiple technologies:

### Idempotency Across Technologies

- **GitHub Actions**: Content hashing before PR creation
- **Argo Workflows**: Check-before-act in workflow steps
- **Kubernetes**: Declarative desired state (built-in idempotency)
- **Helm**: `--atomic` flag for all-or-nothing deployments

### Three-Stage Design Across Technologies

- **GitHub Actions**: Discovery → Distribution → Summary for file distribution
- **Argo Events**: EventSource → Sensor → Workflow trigger
- **Argo Workflows**: DAG with discovery, parallel execution, aggregation
- **Kubernetes Operators**: List resources → Reconcile → Update status

### Retry Strategies Across Technologies

- **GitHub Actions**: `retry` action with exponential backoff
- **Argo Events**: Sensor retry configuration
- **Argo Workflows**: WorkflowTemplate retry strategy
- **Kubernetes**: Deployment rollout with health checks

## Pattern Selection Guide

**Quick selection**:

- Need cross-repo distribution? → **Three-Stage Design** + **Hub-and-Spoke**
- Need to avoid duplicate work? → **Idempotency** + **Work Avoidance**
- Replacing legacy system? → **Strangler Fig**
- Need error resilience? → **Fail Fast** + **Prerequisite Checks**
- Building event-driven system? → **Argo Events** routing patterns
- Orchestrating complex workflows? → **Argo Workflows** composition patterns

## Common Anti-Patterns

### ❌ Non-Idempotent Operations

```yaml
# BAD: Creates duplicate issues on retry
- name: Create issue
  run: gh issue create --title "Alert" --body "Problem detected"
```

```yaml
# GOOD: Check before creating
- name: Create issue if not exists
  run: |
    existing=$(gh issue list --search "Alert" --state all --json number -q '.[0].number')
    if [ -z "$existing" ]; then
      gh issue create --title "Alert" --body "Problem detected"
    fi
```

### ❌ Unnecessary Work

```yaml
# BAD: Always creates PR even if no changes
- name: Create PR
  run: gh pr create --fill
```

```yaml
# GOOD: Check if changes exist first
- name: Create PR if changes exist
  run: |
    if git diff --quiet HEAD origin/main; then
      echo "No changes, skipping PR"
    else
      gh pr create --fill
    fi
```

### ❌ Late Error Detection

```yaml
# BAD: Deploy first, check permissions later
- name: Deploy
  run: kubectl apply -f manifests/
- name: Check RBAC
  run: kubectl auth can-i create deployments
```

```yaml
# GOOD: Validate before deploying
- name: Prerequisite checks
  run: |
    kubectl auth can-i create deployments || exit 1
    kubectl get namespace production || exit 1
- name: Deploy
  run: kubectl apply -f manifests/
```

## Pattern Layering

Patterns work best when layered:

1. **Architecture pattern** (what to build)
    - Example: Hub-and-Spoke for config distribution

2. **Efficiency pattern** (how to build efficiently)
    - Add: Idempotency for safe retries
    - Add: Work Avoidance to skip unchanged configs

3. **Error handling pattern** (how to fail gracefully)
    - Add: Fail Fast for invalid configs
    - Add: Prerequisite Checks for GitHub API access

**Result**: Hub-and-spoke config distribution that is idempotent, skips unchanged configs, validates early, and handles errors gracefully.

## Getting Started

1. **Read the pattern selection guide**: Understand when to use each pattern
2. **Start with architecture patterns**: Choose the right structure
3. **Add efficiency patterns**: Optimize for scale
4. **Layer error handling**: Make it resilient
5. **Browse technology-specific patterns**: See implementations

## Integration with Other Sections

Patterns are applied throughout the other sections:

- [Secure](../secure/index.md): Security scanning patterns, SBOM generation patterns
- [Enforce](../enforce/index.md): Policy enforcement patterns, admission control patterns
- [Build](../build/index.md): Release automation patterns, testing patterns

## Pattern Documentation Format

Each pattern is documented with:

1. **Intent**: What problem does this solve?
2. **Motivation**: When should you use this?
3. **Structure**: How is it organized?
4. **Implementation**: Code examples
5. **Consequences**: Trade-offs and limitations
6. **Related patterns**: What patterns complement this?

## Contributing Patterns

See CONTRIBUTING.md in the project root for guidelines on documenting new patterns.

**Pattern quality criteria**:

- Solves a recurring problem (not one-off solution)
- Technology-agnostic (or clearly scoped to specific tech)
- Production-tested (not theoretical)
- Documented with real examples (not pseudocode)
- Includes anti-patterns (what not to do)

## Related Content

- [Secure](../secure/index.md): Security patterns
- [Enforce](../enforce/index.md): Enforcement patterns
- [Build](../build/index.md): Build and release patterns

## Tags

Browse all content tagged with patterns, automation, idempotency, three-stage, and hub-and-spoke on the [Tags](../tags.md) page.

