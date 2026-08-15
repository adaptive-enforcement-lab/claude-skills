# Strangler Fig

Source: https://adaptive-enforcement-lab.com/patterns/architecture/strangler-fig/

The strangler fig vine grows around a host tree. Eventually the vine takes over completely, and the original tree dies.

> **Migration Pattern**
>
> This guide covers the Strangler Fig pattern for incremental system migration. Review all sections for complete implementation strategy.
>

Your new system gradually replaces the old one. Both run in parallel. Traffic shifts incrementally. When the old system has zero traffic, you remove it.

Zero downtime. Rollback at any point. Migration validated in production.

---

## The Problem with Big Bang Rewrites

Rewriting a monolith all at once:

```mermaid
flowchart LR
    A[Legacy System] --> B[6 Month Rewrite]
    B --> C[Deploy New System]
    C --> D{Works?}
    D -->|No| E[Disaster]
    D -->|Yes| F[Success Maybe]

    %% Ghostty Hardcore Theme
    style A fill:#65d9ef,color:#1b1d1e
    style B fill:#fd971e,color:#1b1d1e
    style C fill:#f92572,color:#1b1d1e
    style D fill:#f92572,color:#1b1d1e
    style E fill:#f92572,color:#1b1d1e
    style F fill:#a7e22e,color:#1b1d1e

```

Six months of development. One deploy. Production traffic hits unknown code. Bugs discovered under real load. No rollback path.

---

## The Strangler Fig Pattern

Incremental replacement:

```mermaid
flowchart TD
    A[Traffic] --> B{Router}
    B -->|90%| C[Legacy System]
    B -->|10%| D[New System]
    C --> E[Legacy Backend]
    D --> F[New Backend]

    %% Ghostty Hardcore Theme
    style A fill:#65d9ef,color:#1b1d1e
    style B fill:#fd971e,color:#1b1d1e
    style C fill:#f92572,color:#1b1d1e
    style D fill:#a7e22e,color:#1b1d1e
    style E fill:#f92572,color:#1b1d1e
    style F fill:#a7e22e,color:#1b1d1e

```

Router directs traffic. Start with 1% to new system. Monitor. Increase gradually. Eventually 100% on new system. Remove legacy.

---

## Two Approaches to Strangler Fig

The strangler fig pattern has two distinct implementation approaches depending on what you're replacing:

### Approach 1: Traffic Routing (User-Facing Systems)

Gradually shift user traffic from old to new system using percentage-based routing.

**Use for**:

- API migrations (REST v1 → v2)
- Feature rollouts (old checkout → new checkout)
- UI rewrites (legacy frontend → modern frontend)
- Application logic changes

**How it works**: Router/proxy directs percentage of traffic to new system. Start at 1%, increase gradually to 100%.

### Approach 2: Component Replacement (Infrastructure)

Replace entire components without routing traffic, including databases, service meshes, operators, and storage.

**Use for**:

- Database migrations (single instance → HA cluster)
- Service mesh replacement (Linkerd → Istio)
- Operator upgrades (CRD v1alpha1 → v1)
- Storage backend changes (EBS → EFS)

**How it works**: Build new component, ensure compatibility, swap references, remove old component. No routing layer needed.

**Key distinction**: Traffic routing = gradual user migration. Component replacement = infrastructure swap with compatibility layer.

---

## Implementation Guides

### Traffic Routing Approach

- **[Implementation Strategies](implementation.md)** - Feature flags, parallel run validation, database migration strategies
- **[Traffic Routing](traffic-routing.md)** - Percentage-based, user-based, and canary deployment patterns
- **[Monitoring and Rollback](monitoring.md)** - Track both systems, compare metrics, instant rollback

### Component Replacement Approach

- **[Platform Component Replacement](platform-component-replacement.md)** - Build-replace-remove pattern for infrastructure, zero downtime component swaps

### Migration Process

- **[Migration Guide](migration-guide.md)** - Eight-phase checklist, common pitfalls, real-world timeline

---

## When to Use This Pattern

**Use when:**

- Replacing critical production systems
- High risk of downtime
- Need gradual validation
- Rollback must be instant

**Don't use when:**

- Small, non-critical systems (just replace)
- No production traffic yet
- Resource cost of running both systems is prohibitive

---

## Related Patterns

- **[Separation of Concerns](../separation-of-concerns/index.md)** - Isolate old and new logic
- **[Graceful Degradation](../../error-handling/graceful-degradation/index.md)** - Fallback to legacy on errors
- **[Environment Progression](../../../blog/posts/2025-12-16-environment-progression-testing.md)** - Test new system in staging first

---

*The new system started at 1% traffic. Mismatches were fixed in shadow mode. Traffic gradually shifted. After 8 weeks, the legacy system handled zero requests. It was decommissioned. The migration completed without a single production incident.*
