---
name: ephemeral-runner-patterns
description: >-
  Disposable runner patterns for GitHub Actions. Container-based, VM-based, and ARC deployment strategies with complete state isolation between jobs.
---

# Ephemeral Runner Patterns

## When to Use This Skill

Persistent runners are persistence vectors. Deploy disposable infrastructure instead.

> **The Goal**
>
>
> Every job executes in a fresh environment. Malicious workflows cannot plant backdoors because the execution environment is destroyed after completion. State isolation prevents cross-job contamination.



## Implementation


See the full implementation guide in the source documentation.


## Examples

See [examples.md](examples.md) for code examples.






## References

- [Source Documentation](https://adaptive-enforcement-lab.com/secure/github-actions-security/)
- [AEL Secure](https://adaptive-enforcement-lab.com/secure/)
