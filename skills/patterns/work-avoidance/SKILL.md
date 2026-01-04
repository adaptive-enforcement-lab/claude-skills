---
name: work-avoidance
description: >-
  Skip work when outcomes won't change. Detect unchanged content, existing resources, and cached outputs to prevent unnecessary PRs, builds, and processing cycles.
---

# Work Avoidance

## When to Use This Skill

Work avoidance detects when an operation isn't needed and skips it entirely. Unlike [idempotency](../idempotency/index.md) (which makes reruns safe), work avoidance prevents the run from happening at all.

```mermaid
flowchart LR
    subgraph trigger[Trigger]
        Event[Event Received]
    end

    subgraph detect[Detection]
        Check{Work Needed?}
    end

    subgraph action[Action]
        Skip[Skip]
        Execute[Execute]
    end

    Event --> Check
    Check -->|No| Skip
    Check -->|Yes| Execute

    %% Ghostty Hardcore Theme
    style Event fill:#65d9ef,color:#1b1d1e
    style Check fill:#fd971e,color:#1b1d1e
    style Skip fill:#5e7175,color:#f8f8f3
    style Execute fill:#a7e22e,color:#1b1d1e

```

---



## Implementation


See the full implementation guide in the source documentation.


## Examples

See [examples.md](examples.md) for code examples.






## References

- [Source Documentation](https://adaptive-enforcement-lab.com/patterns/efficiency/)
- [AEL Patterns](https://adaptive-enforcement-lab.com/patterns/)
