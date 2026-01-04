---
name: graceful-degradation
description: >-
  Build tiered fallback systems that degrade performance, not availability. Cache layers, API resilience, and automatic failover patterns for platform reliability.
---

# Graceful Degradation

## When to Use This Skill

Graceful degradation is a design principle that ensures systems continue operating when components fail. Rather than crashing or returning errors, the system automatically falls back to slower but working alternatives.

```mermaid
flowchart TD
    subgraph request[Request]
        A[Operation Requested]
    end

    subgraph tiers[Fallback Tiers]
        T1[Tier 1: Optimal]
        T2[Tier 2: Acceptable]
        T3[Tier 3: Guaranteed]
    end

    subgraph result[Result]
        Success[Success]
    end

    A --> T1
    T1 -->|Works| Success
    T1 -->|Fails| T2
    T2 -->|Works| Success
    T2 -->|Fails| T3
    T3 --> Success

    %% Ghostty Hardcore Theme
    style A fill:#65d9ef,color:#1b1d1e
    style T1 fill:#a7e22e,color:#1b1d1e
    style T2 fill:#fd971e,color:#1b1d1e
    style T3 fill:#f92572,color:#1b1d1e
    style Success fill:#a7e22e,color:#1b1d1e

```

The key insight: **degrade performance, not availability**.

---



## Implementation


See the full implementation guide in the source documentation.


## Examples

See [examples.md](examples.md) for code examples.






## References

- [Source Documentation](https://adaptive-enforcement-lab.com/patterns/error-handling/)
- [AEL Patterns](https://adaptive-enforcement-lab.com/patterns/)
