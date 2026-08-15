# Framework Selection

Source: https://adaptive-enforcement-lab.com/build/go-cli-architecture/framework-selection/

Choose the right CLI framework and configuration approach for your Go CLI.

---

## Overview

Building a Kubernetes-native CLI requires thoughtful framework selection. The right choice depends on your complexity needs, ecosystem alignment, and team preferences.

This section covers:

- **[CLI Frameworks](cli-frameworks.md)** - Cobra, urfave/cli, and Kong compared
- **[Configuration with Viper](viper-configuration.md)** - Layered configuration management

---

## Quick Recommendation

> **Default Choice**
>
>
> For Kubernetes-native CLIs, use **Cobra + Viper**. It powers kubectl, docker, gh, and most ecosystem tools. Your users already know the patterns.
>

| Need | Recommendation |
| ------ | ---------------- |
| Kubernetes ecosystem CLI | **Cobra** + Viper |
| Simple utility | **urfave/cli** |
| Type-safe struct definitions | **Kong** |

---

## Decision Matrix

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
