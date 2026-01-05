---
name: secure-by-design-pattern-library
description: >-
  Secure-by-design architecture patterns for Kubernetes. Zero trust, defense in depth, least privilege, and fail-secure patterns with implementation examples and threat models.
---

# Secure-by-Design Pattern Library

## When to Use This Skill

Building security into architecture from the ground up, not bolting it on afterward. These patterns enforce security properties at the application, network, and admission control layers, making violations visible and costly.


## Implementation

See the full implementation guide in the [source documentation](https://adaptive-enforcement-lab.com/patterns/security/).


## Techniques


### Integration Patterns

### [End-to-End Deployment](integration.md)

Complete example combining all patterns:

- Zero trust mTLS communication
- Defense in depth pod hardening
- Least privilege RBAC configuration
- Fail secure admission controls

### [Security Audit Checklist](integration.md#security-audit-checklist)

Verification checklist before deployment:

- [ ] Zero Trust: mTLS policies in place
- [ ] Defense in Depth: Pod security contexts enforced
- [ ] Network Policies: Default-deny rules configured
- [ ] Least Privilege: Minimal RBAC permissions
- [ ] Fail Secure: Admission webhooks with failurePolicy: Fail
## References

- [Source Documentation](https://adaptive-enforcement-lab.com/patterns/security/)
- [AEL Patterns](https://adaptive-enforcement-lab.com/patterns/)
