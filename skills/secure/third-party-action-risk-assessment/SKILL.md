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

See the full implementation guide in the [source documentation](https://adaptive-enforcement-lab.com/secure/github-actions-security/).


## Key Principles

**Always SHA pin third-party actions**: Tag references can be mutated. SHA pins are immutable.

```yaml
# Bad - tag reference
- uses: community/action@v2

# Good - SHA pinned with version comment
- uses: community/action@a1b2c3d4e5f6...  # v2.1.0
```

**Review action source code before first use**: Never trust based on stars or README alone. Read the actual implementation.

**Fork critical actions to organization control**: Removes dependency on external maintainer. Gives you control over updates.

**Monitor for action updates**: Use Dependabot to track new versions. Review changelogs before updating.

**Minimize permissions**: Grant actions only what they need. Use job-level scoping to limit scope.

**Isolate high-risk workflows**: Run untrusted actions in separate jobs with minimal permissions and no secret access.

**Audit action usage quarterly**: Review which actions are in use. Re-assess risk as threat landscape evolves.

**Have an exit strategy**: Know how to replace or remove every action if it becomes compromised or unmaintained.


## Examples

See [examples.md](examples.md) for code examples.


## Full Reference

See [reference.md](reference.md) for complete documentation.
## References

- [Source Documentation](https://adaptive-enforcement-lab.com/secure/github-actions-security/)
- [AEL Secure](https://adaptive-enforcement-lab.com/secure/)
