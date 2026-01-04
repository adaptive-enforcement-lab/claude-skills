---
name: matrix-distribution
description: >-
  Parallelize operations across dynamic target lists using GitHub Actions matrix strategies with failure isolation, rate limiting, and conditional logic.
---

# Matrix Distribution

## When to Use This Skill

> **Good Fit**
>


    - Processing multiple repositories, files, or services
    - Operations that are independent and can run in parallel
    - Workloads that benefit from horizontal scaling
    - Batch operations with predictable per-target runtime

> **Poor Fit**
>


    - Sequential operations where order matters
    - Operations with shared state between targets
    - When total job count would exceed GitHub Actions limits (256)

---

##



## Implementation


See the full implementation guide in the source documentation.


## Examples

See [examples.md](examples.md) for code examples.






## References

- [Source Documentation](https://adaptive-enforcement-lab.com/patterns/architecture/)
- [AEL Patterns](https://adaptive-enforcement-lab.com/patterns/)
