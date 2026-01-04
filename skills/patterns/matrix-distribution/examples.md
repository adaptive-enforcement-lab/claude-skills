---
name: matrix-distribution - Examples
description: Code examples for Matrix Distribution
---

# Matrix Distribution - Examples


## Example 1: example-1.yaml


```yaml
strategy:
  matrix:
    target: ${{ fromJson(needs.discover.outputs.targets) }}
  fail-fast: false
  max-parallel: 10
```



## Example 2: example-2.mermaid


```mermaid
flowchart LR
    A[Target List] --> B[Matrix Strategy]
    B --> C1[Job: Target 1]
    B --> C2[Job: Target 2]
    B --> C3[Job: Target N]
    C1 & C2 & C3 --> D[Results]

    %% Ghostty Hardcore Theme
    style A fill:#fd971e,color:#1b1d1e
    style B fill:#65d9ef,color:#1b1d1e
    style C1 fill:#a7e22e,color:#1b1d1e
    style C2 fill:#a7e22e,color:#1b1d1e
    style C3 fill:#a7e22e,color:#1b1d1e
    style D fill:#9e6ffe,color:#1b1d1e
```



## Example 3: example-3.yaml


```yaml
discover:
  outputs:
    targets: ${{ steps.query.outputs.targets }}
  steps:
    - name: Build target list
      id: query
      run: |
        TARGETS='[{"name": "repo-1"}, {"name": "repo-2"}]'
        echo "targets=$TARGETS" >> $GITHUB_OUTPUT

distribute:
  needs: discover
  strategy:
    matrix:
      target: ${{ fromJson(needs.discover.outputs.targets) }}
  steps:
    - run: echo "Processing ${{ matrix.target.name }}"
```



## Example 4: example-4.yaml


```yaml
strategy:
  matrix:
    target: ${{ fromJson(needs.discover.outputs.targets) }}
  fail-fast: false  # Critical: continue processing other targets
```



## Example 5: example-5.yaml


```yaml
strategy:
  matrix:
    target: ${{ fromJson(needs.discover.outputs.targets) }}
  max-parallel: 10  # Limit concurrent jobs
```



