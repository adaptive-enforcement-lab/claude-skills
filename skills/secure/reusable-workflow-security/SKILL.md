---
name: reusable-workflow-security
description: >-
  Secure reusable workflow patterns for GitHub Actions. Input validation, secret inheritance, caller restrictions, and SHA pinning for workflow composition.
---

# Reusable Workflow Security

## When to Use This Skill

Reusable workflows centralize logic but inherit the caller's security context. Unvalidated inputs, unrestricted callers, or unpinned workflow references create privilege escalation vectors and supply chain risks.

> **The Risk**
>
>
> Reusable workflows execute with the caller's GITHUB_TOKEN permissions and secret access. An attacker who controls workflow inputs can inject commands, exfiltrate secrets, or escalate privileges. Unpinned workflow references allow supply chain attacks when upstream workflows are compromised.



## Implementation


See the full implementation guide in the source documentation.


## Examples

See [examples.md](examples.md) for code examples.






## References

- [Source Documentation](https://adaptive-enforcement-lab.com/secure/github-actions-security/)
- [AEL Secure](https://adaptive-enforcement-lab.com/secure/)
