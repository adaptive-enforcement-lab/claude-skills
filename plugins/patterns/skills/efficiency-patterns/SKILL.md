---
name: efficiency-patterns
description: >-
  Optimize automation with idempotency and work avoidance. Safe retries plus skipping unnecessary operations maximize efficiency in CI/CD and platform engineering.
---

# Efficiency Patterns

## When to Use This Skill

Efficiency patterns optimize **what** your automation does and **whether** it needs to do it.

| Pattern | When to Use | Strategy |
| --------- | ------------- | ---------- |
| [Idempotency](idempotency/index.md) | Operations may be retried | Same input = same result |
| [Work Avoidance](work-avoidance/index.md) | Results can be cached | Skip if already done |

---


## Implementation

See the full implementation guide in the [source documentation](https://adaptive-enforcement-lab.com/patterns/efficiency/).


## Techniques


### Overview

Efficiency patterns optimize **what** your automation does and **whether** it needs to do it.

| Pattern | When to Use | Strategy |
| --------- | ------------- | ---------- |
| [Idempotency](idempotency/index.md) | Operations may be retried | Same input = same result |
| [Work Avoidance](work-avoidance/index.md) | Results can be cached | Skip if already done |

---


### Decision Flow

```mermaid
flowchart TD
    A[Operation Request] --> B{Already done?}
    B -->|Yes| C[Work Avoidance: Skip]
    B -->|No| D{May be retried?}
    D -->|Yes| E[Idempotency: Safe retry]
    D -->|No| F[Execute normally]

    %% Ghostty Hardcore Theme
    style A fill:#65d9ef,color:#1b1d1e
    style B fill:#fd971e,color:#1b1d1e
    style C fill:#a7e22e,color:#1b1d1e
    style D fill:#fd971e,color:#1b1d1e
    style E fill:#a7e22e,color:#1b1d1e
    style F fill:#9e6ffe,color:#1b1d1e

```

---


### Quick Reference

| Scenario | Pattern | Reasoning |
| ---------- | --------- | ----------- |
| Re-running same operation | Idempotency | Same result every time |
| Resource already exists | Idempotency | Create-or-update safely |
| Content unchanged | Work Avoidance | Skip unnecessary work |
| Build artifact cached | Work Avoidance | Reuse previous results |

---


### Key Difference

| Aspect | Idempotency | Work Avoidance |
| -------- | ------------- | ---------------- |
| Goal | Safe to retry | Avoid doing work |
| Mechanism | Deterministic result | Change detection |
| Trade-off | Complexity vs reliability | Cache invalidation vs speed |

---

*Idempotency makes retries safe. Work avoidance makes retries unnecessary.*


## Comparison

| Aspect | Idempotency | Work Avoidance |
| -------- | ------------- | ---------------- |
| Goal | Safe to retry | Avoid doing work |
| Mechanism | Deterministic result | Change detection |
| Trade-off | Complexity vs reliability | Cache invalidation vs speed |

---

*Idempotency makes retries safe. Work avoidance makes retries unnecessary.*
## References

- [Source Documentation](https://adaptive-enforcement-lab.com/patterns/efficiency/)
- [AEL Patterns](https://adaptive-enforcement-lab.com/patterns/)
