---
name: slsa-implementation-playbook - Examples
description: Code examples for SLSA Implementation Playbook
---

# SLSA Implementation Playbook - Examples


## Example 1: example-1.yaml


```yaml
jobs:
  provenance:
    permissions:
      actions: read
      id-token: write
      contents: write
    uses: slsa-framework/slsa-github-generator/.github/workflows/generator_generic_slsa3.yml@v2.1.0
    with:
      base64-subjects: "${{ needs.build.outputs.hashes }}"
      upload-assets: true
```



## Example 2: example-2.mermaid


```mermaid
graph TD
    A[Source Code] -->|Branch Protection| B[Protected Branch]
    B -->|Status Checks| C[CI Pipeline]
    C -->|SLSA Provenance| D[Signed Artifact]
    D -->|Verification| E[Runtime Deployment]
    E -->|Policy-as-Code| F[Production]

    %% Ghostty Hardcore Theme
    style A fill:#a7e22e,color:#1b1d1e
    style D fill:#65d9ef,color:#1b1d1e
    style F fill:#f92572,color:#1b1d1e
```



