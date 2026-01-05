---
name: slsa-provenance-toolchain-integration - Examples
description: Code examples for SLSA Provenance: Toolchain Integration
---

# SLSA Provenance: Toolchain Integration - Examples


## Example 1: example-1.yaml


```yaml
jobs:
  build:
    outputs:
      hashes: ${{ steps.hash.outputs.hashes }}
    steps:
      - name: Build artifacts
        run: |
          # Toolchain-specific build commands

      - name: Generate hashes
        id: hash
        run: |
          sha256sum artifacts/* | base64 -w0 > hashes.txt
          echo "hashes=$(cat hashes.txt)" >> "$GITHUB_OUTPUT"

  provenance:
    needs: [build]
    permissions:
      actions: read
      id-token: write
      contents: write
    uses: slsa-framework/slsa-github-generator/.github/workflows/generator_generic_slsa3.yml@v2.1.0
    with:
      base64-subjects: "${{ needs.build.outputs.hashes }}"
      upload-assets: true
```



## Example 2: example-2.yaml


```yaml
jobs:
  build-image:
    outputs:
      digest: ${{ steps.build.outputs.digest }}
    steps:
      - name: Build container image
        id: build
        run: |
          # Toolchain-specific container build
          podman build -t myapp:latest .
          DIGEST=$(podman inspect myapp:latest --format='{{.Id}}')
          echo "digest=${DIGEST}" >> "$GITHUB_OUTPUT"

  provenance:
    needs: [build-image]
    permissions:
      actions: read
      id-token: write
      packages: write
    uses: slsa-framework/slsa-github-generator/.github/workflows/generator_container_slsa3.yml@v2.1.0
    with:
      image: ghcr.io/org/myapp
      digest: "${{ needs.build-image.outputs.digest }}"
```



