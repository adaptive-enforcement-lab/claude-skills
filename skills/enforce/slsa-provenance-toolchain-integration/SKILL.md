---
name: slsa-provenance-toolchain-integration
description: >-
  SLSA provenance for Go, Node.js, and Python: binary builds, package publishing, container images, and dependency verification with cryptographic integrity proofs.
---

# SLSA Provenance: Toolchain Integration

## When to Use This Skill

Language-specific toolchains have unique SLSA integration points:

- **Build artifacts**: Binaries, packages, wheels, container images
- **Package registries**: npm, PyPI, Go modules, GitHub Packages
- **Dependency management**: go.sum, package-lock.json, poetry.lock
- **Build tools**: GoReleaser, npm scripts, setuptools, build isolation

Each toolchain guide covers:

- SLSA Level 3 provenance generation patterns
- Multi-platform and cross-compilation workflows
- Package registry integration
- Dependency lockfile verification
- Container image attestation
- Verification workflows
- Common gotchas and troubleshooting

---


## Implementation

See the full implementation guide in the [source documentation](https://adaptive-enforcement-lab.com/enforce/slsa-provenance/).


## Techniques


### Common Integration Patterns

### Pattern: Multi-Artifact Provenance

All toolchains support generating provenance for multiple artifacts in a single build:

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

This pattern works for:

- Go: Multiple binaries, multi-platform builds
- Node.js: Multiple npm packages, container images
- Python: Multiple wheels, source distributions

### Pattern: Container Image Provenance

All toolchains support container image attestation:

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

See toolchain-specific guides for:

- Go: Distroless base images ([go-advanced.md](go-advanced.md))
- Node.js: Multi-stage builds ([node-advanced.md](node-advanced.md))
- Python: Python slim images ([python-integration.md](python-integration.md))

### Pattern: Dependency Lockfile Verification

All toolchains support dependency verification:

=== "Go"

    ```yaml
    - name: Verify Go modules
      run: |
        go mod verify
        go mod download -json | jq -r '.Error' | grep -q '^null$'
    ```

=== "Node.js"

    ```yaml
    - name: Verify npm dependencies
      run: |
        npm ci --audit
        npm audit signatures
    ```

=== "Python"

    ```yaml
    - name: Verify Python dependencies
      run: |
        pip install --require-hashes -r requirements.txt
        pip check
    ```

---

*See [reference.md](reference.md) for additional techniques and detailed examples.*


## Examples

See [examples.md](examples.md) for code examples.


## Full Reference

See [reference.md](reference.md) for complete documentation.


## Related Patterns

- SLSA Implementation Playbook
- SLSA Levels Explained
- Verification Workflows
- Runner Configuration
- Adoption Roadmap

## References

- [Source Documentation](https://adaptive-enforcement-lab.com/enforce/slsa-provenance/)
- [AEL Enforce](https://adaptive-enforcement-lab.com/enforce/)
