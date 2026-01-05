---
name: slsa-provenance-toolchain-integration - Reference
description: Complete reference for SLSA Provenance: Toolchain Integration
---

# SLSA Provenance: Toolchain Integration - Reference

This is the complete reference documentation extracted from the source.

# SLSA Provenance: Toolchain Integration

Turn language-specific builds into provable pipelines.

<!-- more -->

> **Toolchain Integration Overview**
>
> This section covers SLSA Level 3 provenance generation for Go, Node.js, and Python projects. Each guide includes binary builds, package publishing, container image patterns, and dependency verification workflows.
>

## Overview

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

## Toolchain Guides

### [Go Integration →](go-integration.md)

SLSA provenance for Go binary builds, multi-platform releases, and GoReleaser integration:

- **Binary builds**: Single and multi-platform cross-compilation
- **GoReleaser**: Automated release workflows with provenance
- **Go modules**: Dependency verification with go.sum
- **Container images**: Distroless patterns ([advanced guide](go-advanced.md))

**Key pattern**: Go's reproducible builds + SLSA provenance = non-falsifiable build integrity

### [Node.js Integration →](node-integration.md)

SLSA provenance for npm packages, application artifacts, and container images:

- **npm packages**: Publishing with `npm publish --provenance`
- **Application artifacts**: Bundled JavaScript and TypeScript builds
- **Container images**: Multi-stage builds with Node runtime
- **Dependency lockfiles**: package-lock.json, yarn.lock, pnpm-lock.yaml
- **Registry verification**: npm audit signatures ([advanced guide](node-advanced.md))

**Key pattern**: Lockfile integrity + SLSA provenance = verified supply chain

### [Python Integration →](python-integration.md)

SLSA provenance for PyPI packages, wheels, and container images:

- **PyPI packages**: Publishing wheels and source distributions
- **Application artifacts**: Wheels (.whl), source distributions (.tar.gz)
- **Container images**: Python runtime with application code
- **Dependency lockfiles**: requirements.txt, Pipfile.lock, poetry.lock

**Key pattern**: pip lockfiles + SLSA provenance = provable package publishing

---

## Quick Reference

### Toolchain Comparison

| Toolchain | Primary Artifact | Package Registry | Lockfile | SLSA Tool |
|-----------|------------------|------------------|----------|-----------|
| **Go** | Binary | Go modules | `go.sum` | `slsa-github-generator` |
| **Node.js** | npm package | npm, GitHub Packages | `package-lock.json` | `slsa-github-generator` |
| **Python** | Wheel (.whl) | PyPI, private registries | `poetry.lock` | `slsa-github-generator` |

### Common Patterns

All toolchain guides follow these patterns:

1. Single artifact provenance
2. Multi-platform/multi-artifact builds
3. Package registry integration
4. Container image attestation
5. Verification with slsa-verifier
6. Dependency lockfile verification

### Quick Start Commands

=== "Go"

    ```bash
    # Generate provenance for Go binary
    go build -trimpath -ldflags="-buildid=" -o myapp

    # Verify Go module checksums
    go mod verify

    # Verify SLSA provenance
    slsa-verifier verify-artifact myapp \
      --provenance-path myapp.intoto.jsonl \
      --source-uri github.com/org/repo
    ```

=== "Node.js"

    ```bash
    # Publish npm package with provenance
    npm publish --provenance

    # Verify lockfile integrity
    npm ci --audit

    # Verify SLSA provenance
    slsa-verifier verify-artifact artifact.tgz \
      --provenance-path artifact.tgz.intoto.jsonl \
      --source-uri github.com/org/repo
    ```

=== "Python"

    ```bash
    # Build Python wheel
    python -m build

    # Verify dependency hashes
    pip install --require-hashes -r requirements.txt

    # Verify SLSA provenance
    slsa-verifier verify-artifact dist/mypackage-1.0.0-py3-none-any.whl \
      --provenance-path provenance.intoto.jsonl \
      --source-uri github.com/org/repo
    ```

---

## Common Integration Patterns

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

## Choosing Your Toolchain Guide

Pick the guide matching your project's main language:

- **Go projects** → [Go Integration Guide](go-integration.md)
- **Node.js/TypeScript projects** → [Node.js Integration Guide](node-integration.md)
- **Python projects** → [Python Integration Guide](python-integration.md)

For multi-language projects, start with your primary build artifact's guide, then use multi-artifact provenance pattern to cover all outputs.

---

## Integration Checklist

Use this checklist when integrating SLSA provenance into your toolchain:

- [ ] **Choose toolchain guide** based on primary language
- [ ] **Review build patterns** for your artifact type (binary, package, container)
- [ ] **Implement provenance generation** using `slsa-github-generator`
- [ ] **Test provenance verification** with `slsa-verifier`
- [ ] **Verify dependency lockfiles** in CI/CD
- [ ] **Add deployment gates** requiring provenance verification
- [ ] **Document workflow** for team onboarding
- [ ] **Monitor OpenSSF Scorecard** for Signed-Releases improvement

---

## Advanced Patterns

For advanced integration scenarios, see:

- **[Go Advanced Patterns](go-advanced.md)** - Container images, verification workflows, best practices
- **[Node.js Advanced Patterns](node-advanced.md)** - Registry verification, npm audit signatures, deployment gates
- **[Verification Workflows](../verification-workflows.md)** - Cross-toolchain verification patterns
- **[Policy Templates](../policy-templates.md)** - Kyverno and OPA enforcement for all toolchains

---

## Common Questions

### Do I need separate provenance for each language?

No. Use multi-artifact provenance to cover all build outputs in a single attestation.

### Can I use the same verification workflow for all languages?

Yes. `slsa-verifier` works identically across toolchains regardless of artifact type.

### Should I verify lockfiles or SLSA provenance?

**Both**. Lockfiles verify dependency inputs, SLSA provenance proves build environment integrity.

---

## Next Steps

1. **Choose your toolchain guide**: [Go](go-integration.md), [Node.js](node-integration.md), or [Python](python-integration.md)
2. **Implement provenance generation**: Follow Pattern 1 in your toolchain guide
3. **Add verification workflow**: See toolchain-specific verification sections
4. **Enforce with policy**: Review [Policy Templates](../policy-templates.md)
5. **Scale adoption**: Follow [Adoption Roadmap](../adoption-roadmap.md)

---

## Related Content

- **[SLSA Implementation Playbook](../index.md)** - Complete SLSA adoption guide
- **[SLSA Levels Explained](../slsa-levels.md)** - Understand Level 1-4 requirements
- **[Verification Workflows](../verification-workflows.md)** - CI/CD verification patterns
- **[Runner Configuration](../runner-configuration.md)** - GitHub-hosted vs self-hosted implications
- **[Adoption Roadmap](../adoption-roadmap.md)** - Phased implementation strategy

---

*Language-specific builds become provable pipelines. Choose your toolchain, implement provenance, verify everywhere.*

