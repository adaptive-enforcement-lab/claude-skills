---
name: tombstone-marker-files
description: >-
  Leave markers indicating operations completed. Track progress across complex multi-step workflows. Perfect for resumable operations and distributed systems.
---

# Tombstone/Marker Files

## When to Use This Skill

> **Good Fit**
>


    - Multi-step operations where each step should run once
    - Long-running processes that might be interrupted
    - Operations without natural idempotency (external API calls, emails)
    - Batch processing where items need individual tracking
    - Workflows spanning multiple runs or systems

> **Poor Fit**
>


    - Simple operations where check-before-act suffices
    - When marker storage is unreliable
    - High-frequency operations (marker overhead adds up)
    - When operation result changes over time (markers become stale)

---

##



## Implementation


See the full implementation guide in the source documentation.


## Examples

See [examples.md](examples.md) for code examples.






## References

- [Source Documentation](https://adaptive-enforcement-lab.com/patterns/efficiency/)
- [AEL Patterns](https://adaptive-enforcement-lab.com/patterns/)
