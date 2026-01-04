---
name: idempotency
description: >-
  Build automation that survives reruns. Idempotent operations let you rerun workflows without fear of duplicates, corruption, or cascading failures in CI/CD.
---

# Idempotency

## When to Use This Skill

When your workflow fails at step 47 of 50, you have three options:

1. **Rerun from beginning** - Only safe if workflow is idempotent
2. **Manual intervention** - Fix state by hand, then continue
3. **Abandon and start fresh** - Delete partial state, try again later

> **The Scalable Choice**
>


    Safe reruns are the only scalable choice. Manual intervention and abandoning runs require human effort, don't scale, and introduce errors.

##



## Implementation


See the full implementation guide in the source documentation.







## References

- [Source Documentation](https://adaptive-enforcement-lab.com/patterns/efficiency/)
- [AEL Patterns](https://adaptive-enforcement-lab.com/patterns/)
