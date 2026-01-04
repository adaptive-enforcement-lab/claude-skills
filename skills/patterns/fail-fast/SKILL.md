---
name: fail-fast
description: >-
  Detect and halt on precondition failures before expensive operations begin. Validate inputs, permissions, and state upfront in CI/CD and automation workflows.
---

# Fail Fast

## When to Use This Skill

Fail fast is a design pattern that validates preconditions before executing expensive or irreversible operations. When validation fails, the system immediately reports the error rather than proceeding and failing later in an unpredictable state.

```mermaid
flowchart TD
    subgraph validation[Validation Phase]
        A[Request] --> B{Preconditions Met?}
    end

    subgraph execution[Execution Phase]
        C[Execute Operation]
        D[Success]
    end

    subgraph failure[Failure Phase]
        E[Report Error]
        F[No Side Effects]
    end

    B -->|Yes| C
    C --> D
    B -->|No| E
    E --> F

    %% Ghostty Hardcore Theme
    style A fill:#65d9ef,color:#1b1d1e
    style B fill:#fd971e,color:#1b1d1e
    style C fill:#a7e22e,color:#1b1d1e
    style D fill:#a7e22e,color:#1b1d1e
    style E fill:#f92572,color:#1b1d1e
    style F fill:#f92572,color:#1b1d1e

```

The key insight: **fail before you start, not in the middle**.

---



## Implementation


See the full implementation guide in the source documentation.


## Examples

See [examples.md](examples.md) for code examples.






## References

- [Source Documentation](https://adaptive-enforcement-lab.com/patterns/error-handling/)
- [AEL Patterns](https://adaptive-enforcement-lab.com/patterns/)
