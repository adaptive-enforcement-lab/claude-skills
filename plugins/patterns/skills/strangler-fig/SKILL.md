---
name: strangler-fig
description: >-
  Incremental migration from legacy systems. Run old and new in parallel, gradually shift traffic, rollback at any point. Zero downtime, production-validated.
---

# Strangler Fig

## When to Use This Skill

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


## Implementation

### Traffic Routing Approach

- **[Implementation Strategies](implementation.md)** - Feature flags, parallel run validation, database migration strategies
- **[Traffic Routing](traffic-routing.md)** - Percentage-based, user-based, and canary deployment patterns
- **[Monitoring and Rollback](monitoring.md)** - Track both systems, compare metrics, instant rollback

### Component Replacement Approach

- **[Platform Component Replacement](platform-component-replacement.md)** - Build-replace-remove pattern for infrastructure, zero downtime component swaps

### Migration Process

- **[Migration Guide](migration-guide.md)** - Eight-phase checklist, common pitfalls, real-world timeline

---


## Techniques


### Two Approaches to Strangler Fig

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


## Examples

See [examples.md](examples.md) for code examples.


## Related Patterns

- Separation of Concerns
- Graceful Degradation
- Environment Progression

## References

- [Source Documentation](https://adaptive-enforcement-lab.com/patterns/architecture/)
- [AEL Patterns](https://adaptive-enforcement-lab.com/patterns/)
