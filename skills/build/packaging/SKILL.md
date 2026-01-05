---
name: packaging
description: >-
  Package Go CLIs as minimal secure containers with distroless base images. Static binaries, non-root users, read-only filesystems for production.
---

# Packaging

## When to Use This Skill

Packaging a Go CLI involves creating distributable artifacts that run anywhere. This section covers:

- **[Container Builds](container-builds.md)** - Multi-stage Dockerfiles with distroless
- **[Helm Charts](helm-charts.md)** - Deploy your CLI with Helm
- **[Release Automation](release-automation.md)** - Multi-arch builds and GoReleaser
- **[GitHub Actions](github-actions.md)** - Distribute as a reusable GitHub Action
- **[Pre-commit Hooks](pre-commit-hooks.md)** - Distribute as pre-commit hooks

---


## Implementation

See the full implementation guide in the [source documentation](https://adaptive-enforcement-lab.com/build/go-cli-architecture/).


## Key Principles

| Practice | Description |
| ---------- | ------------- |
| **Static binaries** | Use `CGO_ENABLED=0` for portable builds |
| **Non-root user** | Always run as non-root in containers |
| **Read-only filesystem** | Set `readOnlyRootFilesystem: true` |
| **Drop capabilities** | Remove all capabilities with `drop: ALL` |
| **Version in binary** | Inject version at build time |
| **Multi-arch support** | Build for both amd64 and arm64 |

---

*Ship binaries that run anywhere Kubernetes runs.*
## References

- [Source Documentation](https://adaptive-enforcement-lab.com/build/go-cli-architecture/)
- [AEL Build](https://adaptive-enforcement-lab.com/build/)
