---
name: matrix-filtering-and-deduplication - Troubleshooting
description: Troubleshooting guide for Matrix Filtering and Deduplication
---

# Matrix Filtering and Deduplication - Troubleshooting

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
