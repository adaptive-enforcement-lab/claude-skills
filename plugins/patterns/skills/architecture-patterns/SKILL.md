---
name: architecture-patterns
description: >-
  Fundamental patterns for building maintainable, scalable systems: separation of concerns, distributed orchestration, and zero-downtime migration strategies.
---

# Architecture Patterns

## When to Use This Skill

These patterns govern how systems are structured and how components interact.

> **Implementation Guide**
>
> This guide is part of a modular documentation set. Refer to related guides for complete context.
>

---


## Implementation

See the full implementation guide in the [source documentation](https://adaptive-enforcement-lab.com/patterns/architecture/).


## Techniques


### Patterns in This Section

### [Separation of Concerns](separation-of-concerns/index.md)

Single-responsibility components with clear boundaries. Orchestration separate from execution. Testability through isolation.

**Use when:** Building CLIs, microservices, or any system with distinct responsibilities

**Key benefit:** Maintainability – change one thing without breaking everything

---

### [Hub and Spoke](hub-and-spoke/index.md)

Centralized orchestration with distributed execution. One coordinator, many workers. Event-driven task distribution.

**Use when:** Scaling workflows, managing distributed systems, event-driven architectures

**Key benefit:** Scalability – add workers without changing orchestration

---

### [Strangler Fig](strangler-fig/index.md)

Incremental migration from legacy systems. Run old and new in parallel. Gradually shift traffic. Zero downtime transitions.

**Use when:** Replacing monoliths, migrating to new tech, risky system rewrites

**Key benefit:** Risk reduction – rollback at any point, validate in production

---

### [Three-Stage Design](three-stage-design.md)

Separate discovery, execution, and reporting phases. Workflows that fail gracefully and report completely.

**Use when:** Building complex CI/CD workflows, multi-step automation

**Key benefit:** Observability – always know what happened, even on failure

---

### [Matrix Distribution](matrix-distribution/index.md)

Parallel processing of multiple targets. Dynamic matrices for scalability.

**Use when:** Processing many targets, scaling workflows, reducing execution time

**Key benefit:** Performance – parallel execution instead of sequential

---


### How These Patterns Relate

```mermaid
flowchart TD
    A[System Design] --> B[Separation of Concerns]
    B --> C[Clear Boundaries]
    C --> D[Hub and Spoke]
    D --> E[Distributed Execution]
    A --> F[Legacy Migration]
    F --> G[Strangler Fig]
    G --> B

    %% Ghostty Hardcore Theme
    style A fill:#65d9ef,color:#1b1d1e
    style B fill:#a7e22e,color:#1b1d1e
    style C fill:#fd971e,color:#1b1d1e
    style D fill:#9e6ffe,color:#1b1d1e
    style E fill:#a7e22e,color:#1b1d1e
    style F fill:#65d9ef,color:#1b1d1e
    style G fill:#f92572,color:#1b1d1e

```

Separation of Concerns provides the foundation. Hub and Spoke scales it. Strangler Fig migrates to it.

---


### Related Patterns

These architectural patterns complement:

- **[Efficiency Patterns](../efficiency/index.md):** Idempotency, work avoidance
- **[Error Handling](../error-handling/index.md):** Fail fast, graceful degradation
- **[Argo Workflows](../argo-workflows/index.md):** Production workflow orchestration
- **[Argo Events](../argo-events/index.md):** Event-driven automation

---

*Build systems that scale, change, and survive.*


## Related Patterns

- Efficiency Patterns
- Error Handling
- Argo Workflows
- Argo Events

## References

- [Source Documentation](https://adaptive-enforcement-lab.com/patterns/architecture/)
- [AEL Patterns](https://adaptive-enforcement-lab.com/patterns/)
