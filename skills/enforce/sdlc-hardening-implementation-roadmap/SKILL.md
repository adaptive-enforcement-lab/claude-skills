---
name: sdlc-hardening-implementation-roadmap
description: >-
  Deploy defense-in-depth SDLC hardening across four phases: pre-commit hooks, CI/CD gates, runtime enforcement, and continuous audit evidence collection systems.
---

# SDLC Hardening Implementation Roadmap

## When to Use This Skill

This implementation roadmap provides a structured approach to hardening your Software Development Lifecycle (SDLC) across four critical phases:

1. **[Phase 1: Foundation](phase-1/index.md)** - Local enforcement and branch protection
2. **[Phase 2: Automation](phase-2/index.md)** - CI/CD gates and policy automation
3. **[Phase 3: Runtime](phase-3/index.md)** - Production policy enforcement
4. **[Phase 4: Advanced](phase-4/index.md)** - Audit evidence and compliance validation

Each phase builds on the previous one, creating defense-in-depth through multiple enforcement layers.

---


## Prerequisites

Before starting Phase 1, ensure you have:

- [ ] Access to GitHub organization settings
- [ ] Cloud storage bucket for evidence (GCS, S3, Azure Blob)
- [ ] GitHub App or token with appropriate permissions
- [ ] Kubernetes cluster for runtime policies (Phase 3)
- [ ] Team buy-in on enforcement approach

> **Start Small**
>
> Begin with a single repository or team. Validate controls work before scaling organization-wide. Use pilot repository as reference implementation for others.
>

---


## Implementation


See the full implementation guide in the source documentation.


## Examples

See [examples.md](examples.md) for code examples.






## References

- [Source Documentation](https://adaptive-enforcement-lab.com/enforce/implementation-roadmap/)
- [AEL Enforce](https://adaptive-enforcement-lab.com/enforce/)
