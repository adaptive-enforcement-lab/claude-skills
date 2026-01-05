---
name: github-core-app-setup
description: >-
  Configure organization-level GitHub Apps for secure cross-repository automation. Machine identity, audit trails, and enterprise-grade authentication.
---

# GitHub Core App Setup

## When to Use This Skill

This guide describes the concept, setup, and configuration of a GitHub Core App for organization-level automation.


## Prerequisites

### Required Access

> **Required Access**
>
>
> To create a Core App, you need:
>
> - **Organization owner** role
> - Access to organization settings: `https://github.com/organizations/{ORG}/settings/apps`
>

### Planning Considerations

> **Planning Considerations**
>
>
> Before creating the app, determine:
>
> 1. **Permission scope** - Which repository and organization permissions are needed
> 2. **Installation scope** - All repositories or specific teams
> 3. **Token management** - Where secrets will be stored (repository or organization level)
> 4. **Naming convention** - Standard naming (e.g., "CORE App", "Automation Core")


## Implementation

See the full implementation guide in the [source documentation](https://adaptive-enforcement-lab.com/secure/github-apps/).


## Comparison

| Aspect | Core App | Standard App |
| -------- | ---------- | -------------- |
| **Scope** | Organization-wide | Single repository or selected repos |
| **Purpose** | Infrastructure automation | Feature-specific functionality |
| **Permissions** | Broad, covers common operations | Narrow, task-specific |
| **Installation** | All repositories | Selective repositories |
| **Ownership** | Organization-level admin | Project or team |
| **Lifespan** | Permanent infrastructure | Project lifecycle |
## References

- [Source Documentation](https://adaptive-enforcement-lab.com/secure/github-apps/)
- [AEL Secure](https://adaptive-enforcement-lab.com/secure/)
