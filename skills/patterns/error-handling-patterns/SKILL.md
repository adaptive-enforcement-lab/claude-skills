---
name: error-handling-patterns
description: >-
  Master when to fail fast vs degrade gracefully. Production-tested error handling strategies for GitHub Actions, CI/CD pipelines, and platform automation.
---

# Error Handling Patterns

## When to Use This Skill

Error handling is about **when** and **how** your automation responds to problems.

| Pattern | When to Use | Strategy |
| --------- | ------------- | ---------- |
| [Fail Fast](fail-fast/index.md) | Invalid input, missing config | Stop immediately, report clearly |
| [Prerequisite Checks](prerequisite-checks/index.md) | Complex preconditions | Validate all upfront before work |
| [Graceful Degradation](graceful-degradation/index.md) | Fallbacks exist | Degrade to safer state, continue |

---



## Implementation


See the full implementation guide in the source documentation.







## References

- [Source Documentation](https://adaptive-enforcement-lab.com/patterns/error-handling/)
- [AEL Patterns](https://adaptive-enforcement-lab.com/patterns/)
