---
name: third-party-action-risk-assessment
description: >-
  Structured framework for evaluating GitHub Actions security before adoption. Trust tiers, risk assessment checklist, and decision tree for action evaluation.
---

# Third-Party Action Risk Assessment

## When to Use This Skill

Trust but verify. Every third-party action you adopt into your workflows executes with access to your secrets, code, and deployment infrastructure. Know what you're trusting.

> **The Risk**
>
>
> Third-party actions run arbitrary code inside your CI/CD pipeline with full access to repository secrets, cloud credentials, and source code. A malicious or compromised action can exfiltrate everything, deploy backdoors, or modify your codebase.



## Implementation


See the full implementation guide in the source documentation.


## Examples

See [examples.md](examples.md) for code examples.






## References

- [Source Documentation](https://adaptive-enforcement-lab.com/secure/github-actions-security/)
- [AEL Secure](https://adaptive-enforcement-lab.com/secure/)
