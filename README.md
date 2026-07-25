# Adaptive Enforcement Lab - Claude Code Skills

Claude Code skills marketplace for secure development patterns, enforcement automation, and build engineering.

**118 skills** across 4 plugin collections, generated from [AEL documentation](https://adaptive-enforcement-lab.com).

## Installation

```bash
# Add the AEL skills marketplace
/plugin marketplace add adaptive-enforcement-lab/claude-skills

# Install plugin collections (install any subset)
/plugin install patterns@ael-skills
/plugin install enforce@ael-skills
/plugin install secure@ael-skills
/plugin install build@ael-skills
```

## Available Skills

| Collection | Skills | Version | Focus |
| ---------- | -----: | ------- | ----- |
| [`patterns`](plugins/patterns/skills) | 38 | 1.0.2 | Automation and architecture patterns |
| [`enforce`](plugins/enforce/skills) | 34 | 1.0.2 | Policy-as-code and SDLC hardening |
| [`secure`](plugins/secure/skills) | 33 | 1.0.2 | Supply chain and platform security |
| [`build`](plugins/build/skills) | 13 | 1.0.1 | Go CLI and release engineering |

### Patterns (Development)

From [AEL patterns](https://adaptive-enforcement-lab.com/patterns/):

- GitHub Actions automation — matrix distribution and filtering, work avoidance, idempotency, file distribution
- Argo Workflows and Argo Events — WorkflowTemplates, event routing, concurrency control, scheduled workflows
- GitHub App authentication — JWT generation, installation tokens, OAuth user flows, token lifecycle
- Architecture patterns — hub-and-spoke, strangler fig, separation of concerns, secure-by-design
- Reliability — fail-fast, graceful degradation, prerequisite checks, chaos engineering for Kubernetes

### Enforce (Security)

From [AEL enforcement guides](https://adaptive-enforcement-lab.com/enforce/):

- Kyverno policy templates — pod security, image validation, mutation, generation, network, resource governance
- OPA policy templates — pod security, image security, RBAC, resource governance
- Repository controls — branch protection enforcement, required status checks
- Supply chain — SLSA provenance and toolchain integration, policy packaging, multi-source aggregation
- Operations — phased SDLC hardening roadmap, local development with policy-as-code, incident response playbooks

### Secure (Security)

From [AEL security guides](https://adaptive-enforcement-lab.com/secure/):

- GitHub Actions hardening — action pinning, `GITHUB_TOKEN` permissions, workflow trigger security, third-party action risk
- Secrets — storing credentials, rotation patterns, secret scanning integration
- Authentication — OIDC federation, Workload Identity Federation, GitHub core app setup
- Runners — self-hosted hardening, ephemeral runners, runner group management
- Platform — GKE hardening (cluster config, network, IAM, runtime), OpenSSF Scorecard, Go security tooling

### Build (DevOps)

From [AEL build guides](https://adaptive-enforcement-lab.com/build/):

- Go CLI architecture — framework selection, command architecture, Kubernetes client-go integration
- Release automation — release-please configuration, modular release pipelines
- Packaging — distroless container images, static binaries
- Testing strategies and versioned documentation
- Open source project templates (CONTRIBUTING, SECURITY, issue forms)

## Automated Generation

All skills in `plugins/` are generated from AEL documentation — **never edit them by hand**:

- **Source**: [adaptive-enforcement-lab.com](https://github.com/adaptive-enforcement-lab/adaptive-enforcement-lab-com)
- **Generator**: `skillgen`, a Go extraction pipeline in this repo
- **Sync**: the `generate-skills.yml` workflow runs on `repository_dispatch` (`docs-updated`) from the docs repo, on manual `workflow_dispatch`, and on pull requests to `main`. Dispatch runs open or force-push a PR on the `chore/regenerate-skills` branch; on a release-please PR it commits regenerated output directly to that PR's branch. Regeneration is proposed automatically — merging is a human step.

The same run regenerates `.claude-plugin/marketplace.json` and every `plugins/*/.claude-plugin/plugin.json` from `plugin-metadata.json` plus `.release-please-manifest.json`.

## Team Distribution

To auto-register this marketplace for your team, add to `.claude/settings.json` in your project:

```json
{
  "extraKnownMarketplaces": {
    "ael-skills": {
      "source": {
        "source": "github",
        "repo": "adaptive-enforcement-lab/claude-skills"
      }
    }
  },
  "enabledPlugins": {
    "patterns@ael-skills": true,
    "enforce@ael-skills": true,
    "secure@ael-skills": true,
    "build@ael-skills": true
  }
}
```

## Repository Structure

```
.claude-plugin/
└── marketplace.json              # Marketplace catalog (GENERATED)

plugins/                          # Generated plugins (DO NOT EDIT)
├── patterns/
│   ├── .claude-plugin/
│   │   └── plugin.json           # Plugin metadata (GENERATED)
│   └── skills/                   # One directory per skill, each with SKILL.md
├── enforce/
├── secure/
└── build/

plugin-metadata.json              # Source of truth: descriptions, categories, tags
.release-please-manifest.json     # Source of truth: versions (release-please owned)
release-please-config.json        # Multi-component release configuration

skillgen/                         # Generator source
├── cmd/skillgen/                 # Entry point and dependency injection
├── internal/
│   ├── domain/                   # Core entities (Document, Skill, Marketplace)
│   ├── ports/                    # Interfaces
│   ├── adapters/                 # filesystem, parser, logger
│   └── services/                 # extractor, generator, validator
└── templates/                    # skill, examples, reference, troubleshooting

.github/workflows/
├── generate-skills.yml           # Regeneration PR automation
├── build.yml                     # Tests + cross-platform binaries
└── release.yml                   # release-please
```

## Development

Requires Go 1.26+.

```bash
# Build the generator
cd skillgen && go build -o ../bin/skillgen ./cmd/skillgen

# Run tests
cd skillgen && go test ./...

# Run generator (from repo root, needs the docs repo checked out alongside)
./bin/skillgen \
  --source ../adaptive-enforcement-lab-com/docs \
  --output plugins \
  --plugin-metadata ./plugin-metadata.json \
  --release-manifest ./.release-please-manifest.json
```

Add `--verbose` for per-document logging. See [CLAUDE.md](CLAUDE.md) for the full flag reference and [CONTRIBUTING.md](CONTRIBUTING.md) for development guidelines.

## Architecture

`skillgen` follows Clean/Hexagonal Architecture — dependencies point inward:

- **Domain** (`skillgen/internal/domain`): core entities, no external dependencies
- **Ports** (`skillgen/internal/ports`): interfaces for external dependencies
- **Adapters** (`skillgen/internal/adapters`): filesystem, markdown parser, logger
- **Services** (`skillgen/internal/services`): extractor, generator, validator
- **CMD** (`skillgen/cmd/skillgen`): entry point and wiring

Pipeline: read docs → parse frontmatter and sections → extract skills → render templates → write `SKILL.md` (plus optional `examples.md`, `reference.md`, `troubleshooting.md`) → generate marketplace files.

## Releases

Releases are automated with [release-please](https://github.com/googleapis/release-please). Six components version independently, each with its own release PR:

`skillgen`, `.claude-plugin`, `plugins/patterns`, `plugins/enforce`, `plugins/secure`, `plugins/build`

Conventional commits drive version bumps and changelog entries. Generator releases publish binaries for Linux, macOS (amd64 and arm64), and Windows.

## Contributing

Contributions are welcome! See [CONTRIBUTING.md](CONTRIBUTING.md) for:

- Development setup
- Code standards
- Commit message format
- Pull request process

## Security

For security vulnerability reporting, see [SECURITY.md](SECURITY.md).

## Mission

Turn secure development into an enforced standard, not an afterthought.

## Links

- [AEL Documentation](https://adaptive-enforcement-lab.com)
- [GitHub Organization](https://github.com/adaptive-enforcement-lab)
- [Skills Marketplace](.claude-plugin/marketplace.json)

## License

MIT
