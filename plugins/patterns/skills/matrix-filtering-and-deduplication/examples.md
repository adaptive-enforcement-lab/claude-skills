---
name: matrix-filtering-and-deduplication - Examples
description: Code examples for Matrix Filtering and Deduplication
---

# Matrix Filtering and Deduplication - Examples


## Example 1: example-1.yaml


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



## Example 2: example-2.yaml


```yaml
- name: Debug matrix
  run: |
    echo "Matrix JSON: ${{ needs.detect-changes.outputs.matrix }}"
    echo "${{ needs.detect-changes.outputs.matrix }}" | jq .
```



