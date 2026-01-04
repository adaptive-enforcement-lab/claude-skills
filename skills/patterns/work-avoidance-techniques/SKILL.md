---
name: work-avoidance-techniques
description: >-
  Layer work avoidance checks from existence to content to semantic comparison. Each technique catches different skip scenarios for maximum automation efficiency.
---

# Work Avoidance Techniques

## When to Use This Skill

Each technique answers a specific question:

| Technique | Question | Best For |
| ----------- | ---------- | ---------- |
| [Content Hashing](content-hashing.md) | "Is the content different?" | File comparisons, config sync |
| [Volatile Field Exclusion](volatile-field-exclusion.md) | "Did anything meaningful change?" | Version bumps, timestamps |
| [Existence Checks](existence-checks.md) | "Does it already exist?" | Resource creation (PRs, branches) |
| [Cache-Based Skip](cache-based-skip.md) | "Is the output already built?" | Build artifacts, dependencies |
| [Queue Cleanup](queue-cleanup.md) | "Should queued work execute?" | Mutex-locked workflows |

---



## Implementation


See the full implementation guide in the source documentation.







## References

- [Source Documentation](https://adaptive-enforcement-lab.com/patterns/efficiency/)
- [AEL Patterns](https://adaptive-enforcement-lab.com/patterns/)
