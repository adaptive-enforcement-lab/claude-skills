---
name: hardened-release-workflow - Reference
description: Complete reference for Hardened Release Workflow
---

# Hardened Release Workflow - Reference

This is the complete reference documentation extracted from the source.


# Hardened Release Workflow

Copy-paste ready release workflow templates with comprehensive security hardening. Each example demonstrates signed releases, SLSA provenance generation, artifact attestations, minimal permissions, and secure artifact distribution.

> **Complete Security Patterns**
>
>
> These workflows integrate all security patterns from the hub: SHA-pinned actions, minimal GITHUB_TOKEN permissions, SLSA provenance, artifact attestations, signature verification, and secure distribution. Use as production templates for secure software supply chain.
>

## Release Security Principles

Every release workflow in this guide implements these controls:

1. **Action Pinning**: All third-party actions pinned to full SHA-256 commit hashes
2. **Minimal Permissions**: Only required permissions granted per job
3. **SLSA Provenance**: Build provenance attestations for supply chain transparency
4. **Artifact Attestations**: Cryptographic signatures for release artifacts
5. **Signature Verification**: Verifiable release authenticity
6. **Immutable Releases**: Tag protection and commit verification
7. **Approval Gates**: Environment protection for production releases

## GitHub Release Workflow

Secure workflow for creating GitHub releases with signed artifacts and SLSA provenance.

### Basic Signed Release

Minimal secure release workflow with artifact attestations.

```yaml
name: Secure Release
on:
  push:
    tags:
      # SECURITY: Only trigger on semantic version tags to prevent unauthorized releases
      - 'v*.*.*'

# SECURITY: Minimal permissions by default, escalated per job
permissions:
  contents: read

jobs:
  # Job 1: Build artifacts with attestations
  build:
    runs-on: ubuntu-latest
    permissions:
      contents: read      # Read repository code
      id-token: write     # Generate OIDC tokens for signing
      attestations: write # Create artifact attestations
    outputs:
      artifact-id: ${{ steps.upload.outputs.artifact-id }}
    steps:
      # SECURITY: All actions pinned to full SHA-256 commit hashes
      - name: Checkout code
        uses: actions/checkout@b4ffde65f46336ab88eb53be808477a3936bae11  # v4.1.1
        with:
          # SECURITY: Fetch full history to validate tag points to signed commit
          fetch-depth: 0
          persist-credentials: false

      # SECURITY: Verify tag signature if commit signing enforced
      - name: Verify tag signature
        run: |
          git verify-tag ${{ github.ref_name }} || {
            echo "::error::Tag signature verification failed"
            exit 1
          }

      - name: Set up build environment
        uses: actions/setup-node@5e21ff4d9bc1a8cf6de233a3057d20ec6b3fb69d  # v3.8.1
        with:
          node-version: '20'
          cache: 'npm'

      - name: Install dependencies
        run: npm ci  # Reproducible builds from lock file

      # SECURITY: Run tests before building release artifacts
      - name: Run tests
        run: npm test

      - name: Build release artifacts
        run: npm run build

      # SECURITY: Generate checksums for artifact integrity verification
      - name: Generate checksums
        run: |
          cd dist/
          sha256sum * > SHA256SUMS.txt

      # SECURITY: Upload artifacts with attestation
      # Attestation provides cryptographic proof of artifact origin
      - name: Upload artifacts
        id: upload
        uses: actions/upload-artifact@c7d193f32edcb7bfad88892161225aeda64e9392  # v4.0.0
        with:
          name: release-artifacts
          path: |
            dist/*
            dist/SHA256SUMS.txt
          retention-days: 90  # Long retention for releases

      # SECURITY: Attest artifact provenance
      # Creates SLSA provenance linking artifact to source and build
      - name: Attest artifacts
        uses: actions/attest-build-provenance@1c608d11d69870c2092266b3f9a6f3abbf17002c  # v1.4.3
        with:
          subject-path: 'dist/*'

  # Job 2: Create GitHub release with signed artifacts
  release:
    needs: build
    runs-on: ubuntu-latest
    # SECURITY: Environment protection with approval gate
    environment:
      name: production
      url: https://github.com/${{ github.repository }}/releases/tag/${{ github.ref_name }}
    permissions:
      contents: write     # Create release
      attestations: write # Attach attestations to release
    steps:
      - name: Download artifacts
        uses: actions/download-artifact@fa0a91b85d4f404e444e00e005971372dc801d16  # v4.1.8
        with:
          name: release-artifacts
          path: dist/

      # SECURITY: Verify checksums before release
      - name: Verify artifact checksums
        run: |
          cd dist/
          sha256sum -c SHA256SUMS.txt

      # SECURITY: Create release with generated notes and signed artifacts
      - name: Create GitHub Release
        uses: softprops/action-gh-release@de2c0eb89ae2a093876385947365aca7b0e5f844  # v0.1.15
        with:
          # SECURITY: Generate release notes from commits between tags
          generate_release_notes: true
          # Attach signed artifacts
          files: |
            dist/*
            dist/SHA256SUMS.txt
          # SECURITY: Mark pre-releases for non-stable versions
          prerelease: ${{ contains(github.ref_name, '-rc') || contains(github.ref_name, '-beta') }}
          # Fail if release already exists (prevents overwrites)
          fail_on_unmatched_files: true
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}

  # Job 3: Verify release integrity (post-release validation)
  verify:
    needs: release
    runs-on: ubuntu-latest
    permissions:
      contents: read
      attestations: read  # Verify attestations
    steps:
      - name: Download release artifacts
        run: |
          gh release download ${{ github.ref_name }} \
            --repo ${{ github.repository }} \
            --dir verification/
        env:
          GH_TOKEN: ${{ secrets.GITHUB_TOKEN }}

      - name: Verify checksums
        run: |
          cd verification/
          sha256sum -c SHA256SUMS.txt

      # SECURITY: Verify attestations using GitHub CLI
      - name: Verify attestations
        run: |
          cd verification/
          for file in *; do
            [[ "$file" == "SHA256SUMS.txt" ]] && continue
            echo "Verifying attestation for $file"
            gh attestation verify "$file" \
              --repo ${{ github.repository }} \
              --owner ${{ github.repository_owner }}
          done
        env:
          GH_TOKEN: ${{ secrets.GITHUB_TOKEN }}
```

**Permissions**: `id-token: write` and `attestations: write` for signing, `contents: write` for release creation.

### Advanced Release with SLSA Provenance

Complete release workflow with SLSA L3 provenance generation using official SLSA generators.

```yaml
name: SLSA L3 Release
on:
  push:
    tags:
      - 'v*.*.*'

permissions:
  contents: read

jobs:
  # Job 1: Build with SLSA provenance generator
  # SECURITY: Uses official SLSA generator (isolated build with provenance)
  build:
    permissions:
      id-token: write   # Generate OIDC tokens
      contents: write   # Upload assets to release
      actions: read     # Read workflow metadata
    uses: slsa-framework/slsa-github-generator/.github/workflows/generator_generic_slsa3.yml@v2.0.0
    with:
      # SECURITY: Build command runs in isolated environment
      compile-generator: true
      # Artifact paths to attest
      base64-subjects: |
        {
          "name": "binary-linux-amd64",
          "digest": {
            "sha256": "${{ needs.build-binary.outputs.hash-linux-amd64 }}"
          }
        }

  # Job 2: Build actual release artifacts
  build-binary:
    runs-on: ubuntu-latest
    permissions:
      contents: read
    outputs:
      hash-linux-amd64: ${{ steps.hash.outputs.hash-linux-amd64 }}
    steps:
      - uses: actions/checkout@b4ffde65f46336ab88eb53be808477a3936bae11  # v4.1.1
        with:
          persist-credentials: false

      - uses: actions/setup-go@93397bea11091df50f3d7e59dc26a7711a8bcfbe  # v4.1.0
        with:
          go-version: '1.22'
          cache: true

      # SECURITY: Reproducible build with -trimpath
      - name: Build binary
        run: |
          go build -trimpath -ldflags="-s -w \
            -X main.version=${{ github.ref_name }} \
            -X main.commit=${{ github.sha }} \
            -X main.date=$(date -u +%Y-%m-%dT%H:%M:%SZ)" \
            -o binary-linux-amd64 .

      # SECURITY: Generate hash for provenance
      - name: Generate hash
        id: hash
        run: |
          echo "hash-linux-amd64=$(sha256sum binary-linux-amd64 | cut -d' ' -f1)" >> "$GITHUB_OUTPUT"

      - name: Upload binary
        uses: actions/upload-artifact@c7d193f32edcb7bfad88892161225aeda64e9392  # v4.0.0
        with:
          name: binary-linux-amd64
          path: binary-linux-amd64
          retention-days: 90

  # Job 3: Create release with SLSA provenance
  release:
    needs: [build, build-binary]
    runs-on: ubuntu-latest
    environment: production
    permissions:
      contents: write
    steps:
      - name: Download binary
        uses: actions/download-artifact@fa0a91b85d4f404e444e00e005971372dc801d16  # v4.1.8
        with:
          name: binary-linux-amd64

      # SECURITY: Download SLSA provenance from generator
      - name: Download provenance
        uses: actions/download-artifact@fa0a91b85d4f404e444e00e005971372dc801d16  # v4.1.8
        with:
          name: binary-linux-amd64.intoto.jsonl

      - name: Create release
        uses: softprops/action-gh-release@de2c0eb89ae2a093876385947365aca7b0e5f844  # v0.1.15
        with:
          generate_release_notes: true
          files: |
            binary-linux-amd64
            binary-linux-amd64.intoto.jsonl
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
```

**SLSA Level**: L3 (isolated build with provenance generation via reusable workflow).

## Container Release Workflow

Secure workflow for building and releasing OCI containers with provenance and SBOM.

### Signed Container Release

Build and push container images with SLSA provenance and SBOM attestations.

```yaml
name: Secure Container Release
on:
  push:
    tags:
      - 'v*.*.*'

permissions:
  contents: read

jobs:
  # Job 1: Build and push container with attestations
  build-container:
    runs-on: ubuntu-latest
    environment: production
    permissions:
      contents: read
      packages: write      # Push to GitHub Container Registry
      id-token: write      # Sign with OIDC
      attestations: write  # Create provenance/SBOM
    steps:
      - uses: actions/checkout@b4ffde65f46336ab88eb53be808477a3936bae11  # v4.1.1
        with:
          persist-credentials: false

      # SECURITY: Log in to GHCR using GITHUB_TOKEN (no long-lived credentials)
      - name: Log in to GitHub Container Registry
        uses: docker/login-action@343f7c4344506bcbf9b4de18042ae17996df046d  # v3.0.0
        with:
          registry: ghcr.io
          username: ${{ github.actor }}
          password: ${{ secrets.GITHUB_TOKEN }}

      # SECURITY: Extract metadata for tags and labels
      - name: Extract container metadata
        id: meta

