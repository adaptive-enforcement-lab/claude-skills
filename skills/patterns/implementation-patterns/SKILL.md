---
name: implementation-patterns
description: >-
  Five idempotency patterns for automation: check-before-act, upsert, force overwrite, unique identifiers, and tombstones. Choose based on constraints and APIs.
---

# Implementation Patterns

## When to Use This Skill

| Pattern | Best For | Tradeoff |
| --------- | ---------- | ---------- |
| [Check-Before-Act](check-before-act.md) | Creating resources | Race conditions possible |
| [Upsert](upsert.md) | APIs with atomic operations | Not universally available |
| [Force Overwrite](force-overwrite.md) | Content that can be safely replaced | Destructive if misused |
| [Unique Identifiers](unique-identifiers.md) | Natural deduplication | ID logic can be complex |
| [Tombstone Markers](tombstone-markers/index.md) | Multi-step operations | Markers need cleanup |

---



## Implementation


See the full implementation guide in the source documentation.


## Examples

See [examples.md](examples.md) for code examples.






## References

- [Source Documentation](https://adaptive-enforcement-lab.com/patterns/efficiency/)
- [AEL Patterns](https://adaptive-enforcement-lab.com/patterns/)
