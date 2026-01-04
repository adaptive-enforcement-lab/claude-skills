---
name: jwt-authentication
description: >-
  Generate JWTs for GitHub App authentication. Direct JWT generation for app-level operations, installation discovery, and bootstrapping workflows.
---

# JWT Authentication

## When to Use This Skill

JWTs authenticate your GitHub App itself, not a specific installation. They enable:

- **Installation discovery** - List where your app is installed
- **App metadata retrieval** - Get app configuration and manifest
- **Installation management** - Suspend or configure installations
- **Bootstrap workflows** - Generate installation tokens dynamically

> **JWT Limitations**
>


    - Cannot access repository contents
    - Cannot create issues, pull requests, or commits
    - 10-minute expiration (maximum allowed)
    - App-level permissions only

##



## Implementation


See the full implementation guide in the source documentation.


## Examples

See [examples.md](examples.md) for code examples.






## References

- [Source Documentation](https://adaptive-enforcement-lab.com/patterns/github-actions/)
- [AEL Patterns](https://adaptive-enforcement-lab.com/patterns/)
