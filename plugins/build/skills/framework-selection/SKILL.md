---
name: framework-selection
description: >-
  Choose the right Go CLI framework for Kubernetes-native tools. Decision matrix compares Cobra, urfave/cli, and Kong for ecosystem alignment and features.
---

# Framework Selection

## When to Use This Skill

Building a Kubernetes-native CLI requires thoughtful framework selection. The right choice depends on your complexity needs, ecosystem alignment, and team preferences.

This section covers:

- **[CLI Frameworks](cli-frameworks.md)** - Cobra, urfave/cli, and Kong compared
- **[Configuration with Viper](viper-configuration.md)** - Layered configuration management

---


## When to Apply

| Criteria | Cobra | urfave/cli | Kong |
| ---------- | ------- | ------------ | ------ |
| Ecosystem maturity | High | Medium | Growing |
| Learning curve | Medium | Low | Low |
| Type safety | Low | Low | High |
| Kubernetes alignment | High | Medium | Medium |
| Configuration integration | Excellent (Viper) | Good | Good |
| Shell completion | Built-in | Plugin | Built-in |
| Nested subcommands | Excellent | Good | Good |

---

*Choose tools that match kubectl conventions. Your users already know them.*


## Implementation

See the full implementation guide in the [source documentation](https://adaptive-enforcement-lab.com/build/go-cli-architecture/).
## References

- [Source Documentation](https://adaptive-enforcement-lab.com/build/go-cli-architecture/)
- [AEL Build](https://adaptive-enforcement-lab.com/build/)
