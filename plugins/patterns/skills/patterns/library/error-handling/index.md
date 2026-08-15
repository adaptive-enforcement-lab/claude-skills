# Error Handling Patterns

Source: https://adaptive-enforcement-lab.com/patterns/error-handling/

Patterns for detecting, reporting, and recovering from failures.

> **Core Principle**
>
> Fail fast on **precondition failures**. Degrade gracefully on **runtime failures**.
>

---

## Overview

Error handling is about **when** and **how** your automation responds to problems.

| Pattern | When to Use | Strategy |
| --------- | ------------- | ---------- |
| [Fail Fast](fail-fast/index.md) | Invalid input, missing config | Stop immediately, report clearly |
| [Prerequisite Checks](prerequisite-checks/index.md) | Complex preconditions | Validate all upfront before work |
| [Graceful Degradation](graceful-degradation/index.md) | Fallbacks exist | Degrade to safer state, continue |

---

## Decision Flow

```mermaid
flowchart TD
    A[Error Detected] --> B{Can recover?}
    B -->|No| C[Fail Fast]
    B -->|Yes| D{Before work started?}
    D -->|Yes| E[Prerequisite Check]
    D -->|No| F[Graceful Degradation]

    %% Ghostty Hardcore Theme
    style A fill:#f92572,color:#1b1d1e
    style B fill:#fd971e,color:#1b1d1e
    style C fill:#f92572,color:#1b1d1e
    style D fill:#fd971e,color:#1b1d1e
    style E fill:#65d9ef,color:#1b1d1e
    style F fill:#a7e22e,color:#1b1d1e

```

---

## Quick Reference

| Scenario | Pattern | Reasoning |
| ---------- | --------- | ----------- |
| Missing required config | Fail Fast | Can't continue safely |
| Invalid user input | Fail Fast | User error, report immediately |
| Complex deployment requirements | Prerequisite Checks | Validate tools, access, state |
| API timeout | Graceful Degradation | Retry or use backup |
| Service unavailable | Graceful Degradation | Fall back to alternatives |

---

*Fail fast when you can't recover. Degrade gracefully when you can.*
