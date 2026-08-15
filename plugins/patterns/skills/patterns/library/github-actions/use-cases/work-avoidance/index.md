# Work Avoidance in GitHub Actions

Source: https://adaptive-enforcement-lab.com/patterns/github-actions/use-cases/work-avoidance/

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

```yaml
- name: Check for meaningful changes
  id: check
  run: |
    # Strip version line before comparing
    strip_version() {
      sed '/^version:.*# x-release-please-version$/d' "$1"
    }

    SOURCE=$(strip_version "source/CONFIG.md")
    TARGET=$(git show HEAD:CONFIG.md 2>/dev/null | \
      sed '/^version:.*# x-release-please-version$/d' || echo "")

    if [ "$SOURCE" = "$TARGET" ]; then
      echo "skip=true" >> $GITHUB_OUTPUT
    else
      echo "skip=false" >> $GITHUB_OUTPUT
    fi

- name: Distribute file
  if: steps.check.outputs.skip != 'true'
  run: ./distribute.sh
```

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

```yaml
- name: Check cache
  id: cache
  uses: actions/cache@v4
  with:
    path: dist/
    key: build-${{ hashFiles('src/**') }}

- name: Build
  if: steps.cache.outputs.cache-hit != 'true'
  run: npm run build
```

---

## Related

- [Work Avoidance Pattern](../../../../patterns/efficiency/work-avoidance/index.md) - Conceptual pattern and techniques
- [File Distribution](../file-distribution/index.md) - Applies these patterns at scale
- [Idempotency](../file-distribution/idempotency.md) - Complementary pattern for safe reruns
