---
name: work-avoidance-techniques
description: >-
  Layer work avoidance checks from existence to content to semantic comparison. Each technique catches different skip scenarios for maximum automation efficiency.
---

# Work Avoidance Techniques

## When to Use This Skill

Each technique answers a specific question:

| Technique | Question | Best For |
| ----------- | ---------- | ---------- |
| [Content Hashing](content-hashing.md) | "Is the content different?" | File comparisons, config sync |
| [Volatile Field Exclusion](volatile-field-exclusion.md) | "Did anything meaningful change?" | Version bumps, timestamps |
| [Existence Checks](existence-checks.md) | "Does it already exist?" | Resource creation (PRs, branches) |
| [Cache-Based Skip](cache-based-skip.md) | "Is the output already built?" | Build artifacts, dependencies |
| [Queue Cleanup](queue-cleanup.md) | "Should queued work execute?" | Mutex-locked workflows |

---


## Implementation

See the full implementation guide in the [source documentation](https://adaptive-enforcement-lab.com/patterns/efficiency/).


## Techniques


### Overview

Each technique answers a specific question:

| Technique | Question | Best For |
| ----------- | ---------- | ---------- |
| [Content Hashing](content-hashing.md) | "Is the content different?" | File comparisons, config sync |
| [Volatile Field Exclusion](volatile-field-exclusion.md) | "Did anything meaningful change?" | Version bumps, timestamps |
| [Existence Checks](existence-checks.md) | "Does it already exist?" | Resource creation (PRs, branches) |
| [Cache-Based Skip](cache-based-skip.md) | "Is the output already built?" | Build artifacts, dependencies |
| [Queue Cleanup](queue-cleanup.md) | "Should queued work execute?" | Mutex-locked workflows |

---


### Combining Techniques

Techniques can be layered for maximum efficiency:

```mermaid
flowchart TD
    subgraph layer1[Layer 1: Existence]
        Exists{Resource exists?}
    end

    subgraph layer2[Layer 2: Content]
        Hash{Content hash matches?}
    end

    subgraph layer3[Layer 3: Semantic]
        Volatile{Only volatile fields changed?}
    end

    subgraph action[Action]
        Skip[Skip]
        Execute[Execute]
    end

    Exists -->|Yes| Hash
    Exists -->|No| Execute
    Hash -->|Yes| Skip
    Hash -->|No| Volatile
    Volatile -->|Yes| Skip
    Volatile -->|No| Execute

    %% Ghostty Hardcore Theme
    style Exists fill:#65d9ef,color:#1b1d1e
    style Hash fill:#fd971e,color:#1b1d1e
    style Volatile fill:#9e6ffe,color:#1b1d1e
    style Skip fill:#5e7175,color:#f8f8f3
    style Execute fill:#a7e22e,color:#1b1d1e

```

---


### Choosing a Technique

| Scenario | Recommended Technique |
| ---------- | ---------------------- |
| File distribution with version bumps | [Volatile Field Exclusion](volatile-field-exclusion.md) |
| OCI image rebuilds | [Content Hashing](content-hashing.md) |
| PR/branch creation | [Existence Checks](existence-checks.md) |
| Dependency installation | [Cache-Based Skip](cache-based-skip.md) |
| API state synchronization | [Content Hashing](content-hashing.md) |
| Generated file deployment | [Volatile Field Exclusion](volatile-field-exclusion.md) |
| Idempotent workflows with mutex locks | [Queue Cleanup](queue-cleanup.md) |

---


### Related

- [Work Avoidance Overview](../index.md) - Pattern introduction
- [Anti-Patterns](../anti-patterns.md) - Common mistakes to avoid


## Anti-Patterns to Avoid

Techniques for detecting when work can be safely skipped.

> **Layer Your Checks**
>
> Start with cheap checks (existence), then content hashes, then semantic comparison. Each layer catches different scenarios.
>

---

## Overview

Each technique answers a specific question:

| Technique | Question | Best For |
| ----------- | ---------- | ---------- |
| [Content Hashing](content-hashing.md) | "Is the content different?" | File comparisons, config sync |
| [Volatile Field Exclusion](volatile-field-exclusion.md) | "Did anything meaningful change?" | Version bumps, timestamps |
| [Existence Checks](existence-checks.md) | "Does it already exist?" | Resource creation (PRs, branches) |
| [Cache-Based Skip](cache-based-skip.md) | "Is the output already built?" | Build artifacts, dependencies |
| [Queue Cleanup](queue-cleanup.md) | "Should queued work execute?" | Mutex-locked workflows |

---

## Combining Techniques

Techniques can be layered for maximum efficiency:


*See [examples.md](examples.md) for detailed code examples.*

---

## Choosing a Technique

| Scenario | Recommended Technique |
| ---------- | ---------------------- |
| File distribution with version bumps | [Volatile Field Exclusion](volatile-field-exclusion.md) |
| OCI image rebuilds | [Content Hashing](content-hashing.md) |
| PR/branch creation | [Existence Checks](existence-checks.md) |
| Dependency installation | [Cache-Based Skip](cache-based-skip.md) |
| API state synchronization | [Content Hashing](content-hashing.md) |
| Generated file deployment | [Volatile Field Exclusion](volatile-field-exclusion.md) |
| Idempotent workflows with mutex locks | [Queue Cleanup](queue-cleanup.md) |

---

## Related

- [Work Avoidance Overview](../index.md) - Pattern introduction
- [Anti-Patterns](../anti-patterns.md) - Common mistakes to avoid

### Overview

Each technique answers a specific question:

| Technique | Question | Best For |
| ----------- | ---------- | ---------- |
| [Content Hashing](content-hashing.md) | "Is the content different?" | File comparisons, config sync |
| [Volatile Field Exclusion](volatile-field-exclusion.md) | "Did anything meaningful change?" | Version bumps, timestamps |
| [Existence Checks](existence-checks.md) | "Does it already exist?" | Resource creation (PRs, branches) |
| [Cache-Based Skip](cache-based-skip.md) | "Is the output already built?" | Build artifacts, dependencies |
| [Queue Cleanup](queue-cleanup.md) | "Should queued work execute?" | Mutex-locked workflows |

---

### Combining Techniques

Techniques can be layered for maximum efficiency:


*See [examples.md](examples.md) for detailed code examples.*

---

### Choosing a Technique

| Scenario | Recommended Technique |
| ---------- | ---------------------- |
| File distribution with version bumps | [Volatile Field Exclusion](volatile-field-exclusion.md) |
| OCI image rebuilds | [Content Hashing](content-hashing.md) |
| PR/branch creation | [Existence Checks](existence-checks.md) |
| Dependency installation | [Cache-Based Skip](cache-based-skip.md) |
| API state synchronization | [Content Hashing](content-hashing.md) |
| Generated file deployment | [Volatile Field Exclusion](volatile-field-exclusion.md) |
| Idempotent workflows with mutex locks | [Queue Cleanup](queue-cleanup.md) |

---

### Related

- [Work Avoidance Overview](../index.md) - Pattern introduction
- [Anti-Patterns](../anti-patterns.md) - Common mistakes to avoid


## Related Patterns

- Work Avoidance Overview
- Anti-Patterns

## References

- [Source Documentation](https://adaptive-enforcement-lab.com/patterns/efficiency/)
- [AEL Patterns](https://adaptive-enforcement-lab.com/patterns/)
