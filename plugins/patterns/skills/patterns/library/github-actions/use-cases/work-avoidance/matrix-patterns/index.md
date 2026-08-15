# Matrix Filtering and Deduplication

Source: https://adaptive-enforcement-lab.com/patterns/github-actions/use-cases/work-avoidance/matrix-patterns/

GitHub Actions matrix builds run jobs in parallel. By default, they run everything. Every time.

> **Performance Optimization**
>
> These patterns reduce workflow execution time and cost. Combine multiple techniques for maximum efficiency.
>

This wastes compute. Changed one microservice? Don't rebuild all 47. Modified a Helm chart? Don't run Go tests.

Filtering prevents redundant work. Deduplication eliminates duplicate configurations. Dynamic generation builds matrices based on what changed.

Run only what matters. Skip the rest.

---

## The Problem with Static Matrices

Every push triggers all matrix combinations:

```yaml
jobs:
  test:
    strategy:
      matrix:
        service: [api, auth, billing, notifications, scheduler, worker]
        environment: [dev, staging, prod]
    runs-on: ubuntu-latest
    steps:
      - name: Test ${{ matrix.service }} in ${{ matrix.environment }}
        run: make test-${{ matrix.service }}
```

One-line change to `api` service triggers:

- 6 services × 3 environments = **18 jobs**
- Only 3 jobs (api in dev/staging/prod) are relevant
- 15 jobs run unnecessarily

Cost: 15 × average_runtime × compute_rate = wasted money.

---

## Pattern Categories

### Path-Based Filtering

Control when workflows run based on file changes:

- **[Path Filtering Patterns](path-filtering.md)** - Static path filters, dynamic matrices, and Dorny paths filter

### Matrix Optimization

Reduce redundant job combinations:

- **[Matrix Optimization Patterns](matrix-optimization.md)** - Deduplication, conditional expansion, directory discovery

### Caching and Artifacts

Skip work already done:

- **[Caching and Artifact Patterns](caching-artifacts.md)** - Dependency tracking, caching, artifact reuse

### Advanced Techniques

Combine patterns for maximum efficiency:

- **[Advanced Matrix Patterns](advanced-patterns.md)** - Fast-fail strategies, combining filters

---

## Matrix Size Comparison

| Scenario | Static Matrix | Dynamic Matrix | Savings |
| ---------- | --------------- | ---------------- | --------- |
| 10 services, 1 changed | 10 jobs | 1 job | 90% |
| 5 charts, 2 changed | 10 jobs (lint+test) | 4 jobs | 60% |
| 3 platforms, code unchanged (cached) | 3 builds | 0 builds | 100% |
| Monorepo with 20 microservices | 20 jobs | 3 jobs (avg) | 85% |

---

## When to Use Each Pattern

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

## Debugging Matrix Generation

Matrix doesn't run as expected? Debug with:

```yaml
- name: Debug matrix
  run: |
    echo "Matrix JSON: ${{ needs.detect-changes.outputs.matrix }}"
    echo "${{ needs.detect-changes.outputs.matrix }}" | jq .
```

Common issues:

- Empty matrix `{"include":[]}` runs zero jobs (check `if` condition)
- Invalid JSON breaks `fromJson()` (validate with `jq`)
- Missing quotes in shell scripts mangle arrays

---

## Cost Impact

Real-world example from monorepo with 30 microservices:

**Before (static matrix)**:

- 30 services × 5 checks = 150 jobs per push
- Average 3 minutes per job = 450 minutes
- 1000 pushes/month = 450,000 minutes
- At $0.008/minute = **$3,600/month**

**After (dynamic matrix + filtering)**:

- Average 3 services changed per push
- 3 services × 5 checks = 15 jobs per push
- 15 × 3 minutes = 45 minutes
- 1000 pushes/month = 45,000 minutes
- At $0.008/minute = **$360/month**

Savings: **$3,240/month (90% reduction)**

---

## Related Patterns

- **[Work Avoidance](../index.md)** - Overview of efficiency patterns
- **[Hub and Spoke](../../../../../patterns/architecture/hub-and-spoke/index.md)** - Argo Workflows parallel execution
- **[Idempotency](../../../../../patterns/efficiency/idempotency/index.md)** - Re-runnable jobs

---

*Changed one file. Matrix ran one job. The other 29 stayed idle. Compute saved. Time saved. Money saved.*
