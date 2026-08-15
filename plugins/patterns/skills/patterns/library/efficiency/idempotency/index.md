# Idempotency

Source: https://adaptive-enforcement-lab.com/patterns/efficiency/idempotency/

Build automation that survives reruns.

Workflows fail. Networks timeout. APIs return 500s. Rate limits hit. Runners crash. When failure happens, idempotent operations let you click "rerun" and walk away.

> **Definition**
>
>
> An operation is idempotent if running it multiple times produces the same result as running it once.
>

## Why It Matters

When your workflow fails at step 47 of 50, you have three options:

1. **Rerun from beginning** - Only safe if workflow is idempotent
2. **Manual intervention** - Fix state by hand, then continue
3. **Abandon and start fresh** - Delete partial state, try again later

> **The Scalable Choice**
>
>
> Safe reruns are the only scalable choice. Manual intervention and abandoning runs require human effort, don't scale, and introduce errors.
>

## In This Section

| Page | Description |
| ------ | ------------- |
| [Pros and Cons](pros-and-cons.md) | Tradeoffs of investing in idempotency |
| [Decision Matrix](decision-matrix.md) | When to invest, when to skip |
| [Implementation Patterns](patterns/index.md) | Five patterns with code examples |
| [Real-World Example](real-world-example.md) | File distribution across 40 repositories |
| [Testing](testing.md) | How to verify idempotency |
| [Cache Considerations](caches.md) | The hidden challenge of cached state |

## Quick Reference

```bash
# Idempotent: Running twice produces same result
mkdir -p /tmp/mydir    # Creates dir if missing, no-op if exists

# Not idempotent: Running twice fails or creates duplicates
mkdir /tmp/mydir       # Fails if directory exists
```

For CI/CD pipelines, idempotency means:

- Reruns don't create duplicate PRs
- Reruns don't create duplicate commits
- Reruns don't corrupt data
- Partial failures can be recovered by rerunning
