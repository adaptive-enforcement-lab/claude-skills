---
name: work-avoidance-in-github-actions
description: >-
  Skip unnecessary CI/CD operations before execution. Detect unchanged content, cached builds, and irrelevant paths to reduce workflow costs and execution time.
---

# Work Avoidance in GitHub Actions

## When to Use This Skill

Apply [work avoidance patterns](../../../../patterns/efficiency/work-avoidance/index.md) to skip unnecessary CI/CD operations.

> **Skip Before Execute**
>
> Detect unchanged content, cached builds, and irrelevant paths before running expensive operations.
>

---


## When to Apply

Work avoidance is valuable in GitHub Actions when:

- **Distribution workflows** push files to many repositories
- **Release automation** bumps versions without content changes
- **Scheduled jobs** run regardless of whether work exists
- **Monorepo builds** trigger on any change but only need subset builds

---


## Implementation

| Pattern | Operator Manual | Engineering Pattern |
| --------- | ----------------- | --------------------- |
| Skip version-only changes | [Content Comparison](content-comparison.md) | [Volatile Field Exclusion](../../../../patterns/efficiency/work-avoidance/techniques/volatile-field-exclusion.md) |
| Skip unchanged paths | [Path Filtering](path-filtering.md) | N/A (native GitHub feature) |
| Skip cached builds | [Cache-Based Skip](cache-based-skip.md) | [Cache-Based Skip](../../../../patterns/efficiency/work-avoidance/techniques/cache-based-skip.md) |

---


## Techniques


### Implementation Patterns

| Pattern | Operator Manual | Engineering Pattern |
| --------- | ----------------- | --------------------- |
| Skip version-only changes | [Content Comparison](content-comparison.md) | [Volatile Field Exclusion](../../../../patterns/efficiency/work-avoidance/techniques/volatile-field-exclusion.md) |
| Skip unchanged paths | [Path Filtering](path-filtering.md) | N/A (native GitHub feature) |
| Skip cached builds | [Cache-Based Skip](cache-based-skip.md) | [Cache-Based Skip](../../../../patterns/efficiency/work-avoidance/techniques/cache-based-skip.md) |

---


## Anti-Patterns to Avoid

Apply [work avoidance patterns](../../../../patterns/efficiency/work-avoidance/index.md) to skip unnecessary CI/CD operations.

> **Skip Before Execute**
>
> Detect unchanged content, cached builds, and irrelevant paths before running expensive operations.
>

---

## When to Apply

Work avoidance is valuable in GitHub Actions when:

- **Distribution workflows** push files to many repositories
- **Release automation** bumps versions without content changes
- **Scheduled jobs** run regardless of whether work exists
- **Monorepo builds** trigger on any change but only need subset builds

---

## Implementation Patterns

| Pattern | Operator Manual | Engineering Pattern |
| --------- | ----------------- | --------------------- |
| Skip version-only changes | [Content Comparison](content-comparison.md) | [Volatile Field Exclusion](../../../../patterns/efficiency/work-avoidance/techniques/volatile-field-exclusion.md) |
| Skip unchanged paths | [Path Filtering](path-filtering.md) | N/A (native GitHub feature) |
| Skip cached builds | [Cache-Based Skip](cache-based-skip.md) | [Cache-Based Skip](../../../../patterns/efficiency/work-avoidance/techniques/cache-based-skip.md) |

---

## Quick Reference

### Check for Meaningful Changes


*See [examples.md](examples.md) for detailed code examples.*

### Path-Based Filtering

```yaml
on:
  push:
    paths:
      - 'src/**'
      - 'package.json'
    paths-ignore:
      - '**.md'
      - 'docs/**'
```

### Cache-Based Skip


*See [examples.md](examples.md) for detailed code examples.*

---

## Related

- [Work Avoidance Pattern](../../../../patterns/efficiency/work-avoidance/index.md) - Conceptual pattern and techniques
- [File Distribution](../file-distribution/index.md) - Applies these patterns at scale
- [Idempotency](../file-distribution/idempotency.md) - Complementary pattern for safe reruns

### When to Apply

Work avoidance is valuable in GitHub Actions when:

- **Distribution workflows** push files to many repositories
- **Release automation** bumps versions without content changes
- **Scheduled jobs** run regardless of whether work exists
- **Monorepo builds** trigger on any change but only need subset builds

---

### Implementation Patterns

| Pattern | Operator Manual | Engineering Pattern |
| --------- | ----------------- | --------------------- |
| Skip version-only changes | [Content Comparison](content-comparison.md) | [Volatile Field Exclusion](../../../../patterns/efficiency/work-avoidance/techniques/volatile-field-exclusion.md) |
| Skip unchanged paths | [Path Filtering](path-filtering.md) | N/A (native GitHub feature) |
| Skip cached builds | [Cache-Based Skip](cache-based-skip.md) | [Cache-Based Skip](../../../../patterns/efficiency/work-avoidance/techniques/cache-based-skip.md) |

---

### Quick Reference

### Check for Meaningful Changes


*See [examples.md](examples.md) for detailed code examples.*

### Path-Based Filtering

```yaml
on:
  push:
    paths:
      - 'src/**'
      - 'package.json'
    paths-ignore:
      - '**.md'
      - 'docs/**'
```

### Cache-Based Skip


*See [examples.md](examples.md) for detailed code examples.*

---

### Related

- [Work Avoidance Pattern](../../../../patterns/efficiency/work-avoidance/index.md) - Conceptual pattern and techniques
- [File Distribution](../file-distribution/index.md) - Applies these patterns at scale
- [Idempotency](../file-distribution/idempotency.md) - Complementary pattern for safe reruns


## Examples

See [examples.md](examples.md) for code examples.


## Related Patterns

- Work Avoidance Pattern
- File Distribution
- Idempotency

## References

- [Source Documentation](https://adaptive-enforcement-lab.com/patterns/github-actions/)
- [AEL Patterns](https://adaptive-enforcement-lab.com/patterns/)
