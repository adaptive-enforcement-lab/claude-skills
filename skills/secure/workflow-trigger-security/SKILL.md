---
name: workflow-trigger-security
description: >-
  Secure GitHub Actions trigger patterns for pull requests, forks, and reusable workflows. Preventing privilege escalation and code injection through trigger misconfiguration.
---

# Workflow Trigger Security

## When to Use This Skill

Workflow triggers control when code executes and in what security context. Choose the wrong trigger and attackers can inject code, escalate privileges, or exfiltrate secrets from forks.

> **The Risk**
>
>
> Trigger misconfiguration is the leading cause of GitHub Actions privilege escalation. `pull_request_target` runs fork code in the base repository context with full secret access. External contributors can weaponize this to steal credentials, push malicious code, or compromise your infrastructure.



## Implementation


See the full implementation guide in the source documentation.


## Examples

See [examples.md](examples.md) for code examples.






## References

- [Source Documentation](https://adaptive-enforcement-lab.com/secure/github-actions-security/)
- [AEL Secure](https://adaptive-enforcement-lab.com/secure/)
