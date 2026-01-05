---
name: tombstone-marker-files
description: >-
  Leave markers indicating operations completed. Track progress across complex multi-step workflows. Perfect for resumable operations and distributed systems.
---

# Tombstone/Marker Files

## When to Use This Skill

> **Poor Fit**
>
>
> - Simple operations where check-before-act suffices
> - When marker storage is unreliable
> - High-frequency operations (marker overhead adds up)
> - When operation result changes over time (markers become stale)
>

---


## Implementation

### Basic Marker File


*See [examples.md](examples.md) for detailed code examples.*

### Run-Scoped Markers


*See [examples.md](examples.md) for detailed code examples.*

### Operation-Scoped Markers


*See [examples.md](examples.md) for detailed code examples.*

### Markers with Metadata


*See [examples.md](examples.md) for detailed code examples.*

### Directory-Based Markers


*See [examples.md](examples.md) for detailed code examples.*

---


## Techniques


### Comparison with Other Patterns

| Aspect | [Check-Before-Act](../check-before-act.md) | [Unique Identifiers](../unique-identifiers.md) | Tombstone Markers |
| -------- | ----------------- | ------------------- | ------------------- |
| Tracks completion | No | No | Yes |
| Works for any operation | No | No | Yes |
| Requires storage | No | No | Yes |
| Can track partial progress | No | No | Yes |
| Cleanup required | No | No | Yes |

---


## Comparison

| Aspect | [Check-Before-Act](../check-before-act.md) | [Unique Identifiers](../unique-identifiers.md) | Tombstone Markers |
| -------- | ----------------- | ------------------- | ------------------- |
| Tracks completion | No | No | Yes |
| Works for any operation | No | No | Yes |
| Requires storage | No | No | Yes |
| Can track partial progress | No | No | Yes |
| Cleanup required | No | No | Yes |

---


## Examples

See [examples.md](examples.md) for code examples.
## References

- [Source Documentation](https://adaptive-enforcement-lab.com/patterns/efficiency/)
- [AEL Patterns](https://adaptive-enforcement-lab.com/patterns/)
