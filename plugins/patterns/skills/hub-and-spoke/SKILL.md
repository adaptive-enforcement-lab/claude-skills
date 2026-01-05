---
name: hub-and-spoke
description: >-
  Centralized orchestration with distributed execution. One coordinator spawns many workers in parallel. Scale horizontally without changing hub logic.
---

# Hub and Spoke

## When to Use This Skill

One hub coordinates. Many spokes execute. The hub doesn't do the work. It distributes, tracks, and summarizes.

This pattern scales horizontally. Add workers without touching the orchestrator.


## Implementation

Hub workflow spawns children:


*See [examples.md](examples.md) for detailed code examples.*

Hub discovers repositories, spawns a spoke workflow for each, then summarizes results.


## Examples

See [examples.md](examples.md) for code examples.
## References

- [Source Documentation](https://adaptive-enforcement-lab.com/patterns/architecture/)
- [AEL Patterns](https://adaptive-enforcement-lab.com/patterns/)
