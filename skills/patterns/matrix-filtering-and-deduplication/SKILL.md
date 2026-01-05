---
name: matrix-filtering-and-deduplication
description: >-
  Reduce matrix builds from 47 jobs to 3 with path filtering, deduplication, and dynamic generation. Run only what changed and eliminate redundant combinations.
---

# Matrix Filtering and Deduplication

## When to Use This Skill

| Pattern | Use Case | Complexity |
| --------- | ---------- | ------------ |
| **Path Filters** | Single workflow, simple triggers | Low |
| **Dynamic Matrix** | Monorepo, many services | Medium |
| **Dorny Paths Filter** | Shared dependencies, cross-cutting changes | Low |
| **Deduplication** | Overlapping test configurations | Low |
| **Conditional Expansion** | Different rigor per event (push vs PR) | Medium |
| **Directory Discovery** | Auto-scaling as repo grows | Medium |
| **Dependency Tracking** | Expensive vendor/build operations | Low |
| **Fast-Fail** | Critical checks vs optional validations | Low |
| **Caching** | Deterministic builds | Medium |
| **Artifacts** | Build once, test many | Low |
| **Combined Filters** | Maximum work avoidance | High |

---


## Implementation

See the full implementation guide in the [source documentation](https://adaptive-enforcement-lab.com/patterns/github-actions/).


## Comparison

| Scenario | Static Matrix | Dynamic Matrix | Savings |
| ---------- | --------------- | ---------------- | --------- |
| 10 services, 1 changed | 10 jobs | 1 job | 90% |
| 5 charts, 2 changed | 10 jobs (lint+test) | 4 jobs | 60% |
| 3 platforms, code unchanged (cached) | 3 builds | 0 builds | 100% |
| Monorepo with 20 microservices | 20 jobs | 3 jobs (avg) | 85% |

---


## Examples

See [examples.md](examples.md) for code examples.


## Troubleshooting

See [troubleshooting.md](troubleshooting.md) for common issues and solutions.


## Related Patterns

- Work Avoidance
- Hub and Spoke
- Idempotency

## References

- [Source Documentation](https://adaptive-enforcement-lab.com/patterns/github-actions/)
- [AEL Patterns](https://adaptive-enforcement-lab.com/patterns/)
