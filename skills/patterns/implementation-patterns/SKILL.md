---
name: implementation-patterns
description: >-
  Five idempotency patterns for automation: check-before-act, upsert, force overwrite, unique identifiers, and tombstones. Choose based on constraints and APIs.
---

# Implementation Patterns

## When to Use This Skill

| Pattern | Best For | Tradeoff |
| --------- | ---------- | ---------- |
| [Check-Before-Act](check-before-act.md) | Creating resources | Race conditions possible |
| [Upsert](upsert.md) | APIs with atomic operations | Not universally available |
| [Force Overwrite](force-overwrite.md) | Content that can be safely replaced | Destructive if misused |
| [Unique Identifiers](unique-identifiers.md) | Natural deduplication | ID logic can be complex |
| [Tombstone Markers](tombstone-markers/index.md) | Multi-step operations | Markers need cleanup |

---


## Implementation

Five patterns for making operations idempotent. Each has tradeoffs; choose based on your constraints.

---

## Pattern Overview

| Pattern | Best For | Tradeoff |
| --------- | ---------- | ---------- |
| [Check-Before-Act](check-before-act.md) | Creating resources | Race conditions possible |
| [Upsert](upsert.md) | APIs with atomic operations | Not universally available |
| [Force Overwrite](force-overwrite.md) | Content that can be safely replaced | Destructive if misused |
| [Unique Identifiers](unique-identifiers.md) | Natural deduplication | ID logic can be complex |
| [Tombstone Markers](tombstone-markers/index.md) | Multi-step operations | Markers need cleanup |

---

## Quick Reference

### [Check-Before-Act](check-before-act.md)

The most common pattern. Check if the target state exists before attempting to create it.

```bash
if git ls-remote --heads origin "$BRANCH" | grep -q "$BRANCH"; then
  git checkout -B "$BRANCH" "origin/$BRANCH"
else
  git checkout -b "$BRANCH"
fi
```

### [Create-or-Update (Upsert)](upsert.md)

Use APIs or commands that handle both cases atomically.

```bash
gh release create v1.0.0 --notes "Release" || gh release edit v1.0.0 --notes "Release"
```

### [Force Overwrite](force-overwrite.md)

Don't check, just overwrite. Safe when overwriting with identical content is acceptable.

```bash
git push --force-with-lease origin "$BRANCH"
```

### [Unique Identifiers](unique-identifiers.md)

Generate deterministic IDs so duplicate operations target the same resource.

```bash
BRANCH="update-$(sha256sum file.txt | cut -c1-8)"
```

### [Tombstone/Marker Files](tombstone-markers/index.md)

Leave markers indicating operations completed.

```bash
MARKER=".completed-$RUN_ID"
[ -f "$MARKER" ] && exit 0
# Do work...
touch "$MARKER"
```

---

## Choosing a Pattern


*See [examples.md](examples.md) for detailed code examples.*

| Scenario | Recommended Pattern |
| ---------- | ------------------- |
| Creating resources (PRs, branches, files) | [Check-Before-Act](check-before-act.md) |
| Updating existing resources | [Upsert](upsert.md) or [Force Overwrite](force-overwrite.md) |
| Operations with natural keys | [Unique Identifiers](unique-identifiers.md) |
| Complex multi-step operations | [Tombstone Markers](tombstone-markers/index.md) |
| API supports atomic operations | [Upsert](upsert.md) |

> **Combine Patterns**
>
>
> Real-world automation often combines multiple patterns. A workflow might use Check-Before-Act for PR creation, Force Overwrite for branch updates, and Unique Identifiers for naming.

### Pattern Overview

| Pattern | Best For | Tradeoff |
| --------- | ---------- | ---------- |
| [Check-Before-Act](check-before-act.md) | Creating resources | Race conditions possible |
| [Upsert](upsert.md) | APIs with atomic operations | Not universally available |
| [Force Overwrite](force-overwrite.md) | Content that can be safely replaced | Destructive if misused |
| [Unique Identifiers](unique-identifiers.md) | Natural deduplication | ID logic can be complex |
| [Tombstone Markers](tombstone-markers/index.md) | Multi-step operations | Markers need cleanup |

---

### Quick Reference

### [Check-Before-Act](check-before-act.md)

The most common pattern. Check if the target state exists before attempting to create it.

```bash
if git ls-remote --heads origin "$BRANCH" | grep -q "$BRANCH"; then
  git checkout -B "$BRANCH" "origin/$BRANCH"
else
  git checkout -b "$BRANCH"
fi
```

### [Create-or-Update (Upsert)](upsert.md)

Use APIs or commands that handle both cases atomically.

```bash
gh release create v1.0.0 --notes "Release" || gh release edit v1.0.0 --notes "Release"
```

### [Force Overwrite](force-overwrite.md)

Don't check, just overwrite. Safe when overwriting with identical content is acceptable.

```bash
git push --force-with-lease origin "$BRANCH"
```

### [Unique Identifiers](unique-identifiers.md)

Generate deterministic IDs so duplicate operations target the same resource.

```bash
BRANCH="update-$(sha256sum file.txt | cut -c1-8)"
```

### [Tombstone/Marker Files](tombstone-markers/index.md)

Leave markers indicating operations completed.

```bash
MARKER=".completed-$RUN_ID"
[ -f "$MARKER" ] && exit 0
# Do work...
touch "$MARKER"
```

---

### Choosing a Pattern


*See [examples.md](examples.md) for detailed code examples.*

| Scenario | Recommended Pattern |
| ---------- | ------------------- |
| Creating resources (PRs, branches, files) | [Check-Before-Act](check-before-act.md) |
| Updating existing resources | [Upsert](upsert.md) or [Force Overwrite](force-overwrite.md) |
| Operations with natural keys | [Unique Identifiers](unique-identifiers.md) |
| Complex multi-step operations | [Tombstone Markers](tombstone-markers/index.md) |
| API supports atomic operations | [Upsert](upsert.md) |

> **Combine Patterns**
>
>
> Real-world automation often combines multiple patterns. A workflow might use Check-Before-Act for PR creation, Force Overwrite for branch updates, and Unique Identifiers for naming.


## Techniques


### Pattern Overview

| Pattern | Best For | Tradeoff |
| --------- | ---------- | ---------- |
| [Check-Before-Act](check-before-act.md) | Creating resources | Race conditions possible |
| [Upsert](upsert.md) | APIs with atomic operations | Not universally available |
| [Force Overwrite](force-overwrite.md) | Content that can be safely replaced | Destructive if misused |
| [Unique Identifiers](unique-identifiers.md) | Natural deduplication | ID logic can be complex |
| [Tombstone Markers](tombstone-markers/index.md) | Multi-step operations | Markers need cleanup |

---


### Quick Reference

### [Check-Before-Act](check-before-act.md)

The most common pattern. Check if the target state exists before attempting to create it.

```bash
if git ls-remote --heads origin "$BRANCH" | grep -q "$BRANCH"; then
  git checkout -B "$BRANCH" "origin/$BRANCH"
else
  git checkout -b "$BRANCH"
fi
```

### [Create-or-Update (Upsert)](upsert.md)

Use APIs or commands that handle both cases atomically.

```bash
gh release create v1.0.0 --notes "Release" || gh release edit v1.0.0 --notes "Release"
```

### [Force Overwrite](force-overwrite.md)

Don't check, just overwrite. Safe when overwriting with identical content is acceptable.

```bash
git push --force-with-lease origin "$BRANCH"
```

### [Unique Identifiers](unique-identifiers.md)

Generate deterministic IDs so duplicate operations target the same resource.

```bash
BRANCH="update-$(sha256sum file.txt | cut -c1-8)"
```

### [Tombstone/Marker Files](tombstone-markers/index.md)

Leave markers indicating operations completed.

```bash
MARKER=".completed-$RUN_ID"
[ -f "$MARKER" ] && exit 0
# Do work...
touch "$MARKER"
```

---


### Choosing a Pattern

```mermaid
flowchart TD
    A[Need idempotency] --> B{API has upsert?}
    B -->|Yes| C[Use Upsert]
    B -->|No| D{Safe to overwrite?}
    D -->|Yes| E[Use Force Overwrite]
    D -->|No| F{Natural unique key?}
    F -->|Yes| G[Use Unique Identifiers]
    F -->|No| H{Multi-step operation?}
    H -->|Yes| I[Use Tombstone Markers]
    H -->|No| J[Use Check-Before-Act]

    %% Ghostty Hardcore Theme
    style A fill:#5e7175,color:#f8f8f3
    style B fill:#fd971e,color:#1b1d1e
    style C fill:#a7e22e,color:#1b1d1e
    style D fill:#fd971e,color:#1b1d1e
    style E fill:#e6db74,color:#1b1d1e
    style F fill:#fd971e,color:#1b1d1e
    style G fill:#65d9ef,color:#1b1d1e
    style H fill:#fd971e,color:#1b1d1e
    style I fill:#9e6ffe,color:#1b1d1e
    style J fill:#65d9ef,color:#1b1d1e

```

| Scenario | Recommended Pattern |
| ---------- | ------------------- |
| Creating resources (PRs, branches, files) | [Check-Before-Act](check-before-act.md) |
| Updating existing resources | [Upsert](upsert.md) or [Force Overwrite](force-overwrite.md) |
| Operations with natural keys | [Unique Identifiers](unique-identifiers.md) |
| Complex multi-step operations | [Tombstone Markers](tombstone-markers/index.md) |
| API supports atomic operations | [Upsert](upsert.md) |

> **Combine Patterns**
>
>
> Real-world automation often combines multiple patterns. A workflow might use Check-Before-Act for PR creation, Force Overwrite for branch updates, and Unique Identifiers for naming.


## Examples

See [examples.md](examples.md) for code examples.
## References

- [Source Documentation](https://adaptive-enforcement-lab.com/patterns/efficiency/)
- [AEL Patterns](https://adaptive-enforcement-lab.com/patterns/)
