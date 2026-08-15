# Adaptive Enforcement Lab - Claude Code Skills

Claude Code skills marketplace for secure development patterns, enforcement automation, and build engineering.

**4 hub skills** — one per plugin collection, each a lean scannable index plus a full offline reference and the complete source doc library — generated from [AEL documentation](https://adaptive-enforcement-lab.com).

## Installation

```bash
# Add the AEL skills marketplace
/plugin marketplace add adaptive-enforcement-lab/claude-skills

# Install plugin collections (install any subset)
/plugin install build@ael-skills
/plugin install enforce@ael-skills
/plugin install patterns@ael-skills
/plugin install secure@ael-skills
```

## Available Skills

| Collection | Version | Topics indexed | Focus |
| ---------- | ------- | --------------: | ----- |
| [`build`](plugins/build/skills/build) | 1.1.1 | 13 | Use when building a Go CLI, wiring release-please for automated versioning, packaging a container… |
| [`enforce`](plugins/enforce/skills/enforce) | 1.1.1 | 34 | Use when setting up or reviewing policy-as-code enforcement — writing Kyverno or OPA policies… |
| [`patterns`](plugins/patterns/skills/patterns) | 1.1.1 | 38 | Use when designing or reviewing automation architecture for GitHub Actions, Argo Workflows, or Argo… |
| [`secure`](plugins/secure/skills/secure) | 1.1.1 | 33 | Use when hardening GitHub Actions workflows, managing secrets or self-hosted runners, configuring OIDC or… |

Each hub skill ships three files: `SKILL.md` (short overview + grouped link index, under ~500 words), `reference.md` (every topic's full content, concatenated), and `library/` (every source doc shipped verbatim, one file each, mirroring the AEL docs tree).

### Build (DevOps)

From [AEL Build](https://adaptive-enforcement-lab.com/build/):

- **Documentation as Skills** — Compile MkDocs documentation into Claude Code…
- **Go CLI Architecture** — Build Kubernetes-native CLIs in Go with…
- **Modular Release Pipelines** — Automate version management and changelog generation…
- **Open Source Project Templates** — Production-ready templates for CONTRIBUTING.
- **Versioned Documentation** — Deploy version-tagged documentation alongside releases using…

### Enforce (Security)

From [AEL Enforce](https://adaptive-enforcement-lab.com/enforce/):

- **Branch Protection Enforcement Patterns** — Comprehensive branch protection configuration patterns with…
- **Implementation Roadmap** — Phased rollout plan for SDLC hardening.
- **Incident Readiness**
- **Policy-as-Code: End-to-End Enforcement** — Enforce security and compliance policies across…
- **Required Status Checks** — CI/CD pipelines as merge gates.
- **SLSA Implementation Playbook** — Complete SLSA implementation playbook: clarify SLSA…

### Patterns (Development)

From [AEL Patterns](https://adaptive-enforcement-lab.com/patterns/):

- **Architecture Patterns** — Fundamental patterns for building maintainable, scalable…
- **Argo Events** — Build event-driven Kubernetes automation with Argo…
- **Argo Workflows Patterns** — Production Argo Workflows patterns: reusable templates…
- **Efficiency Patterns** — Optimize automation with idempotency and work…
- **Error Handling Patterns** — Master when to fail fast vs…
- **Github Actions**
- **Reliability**
- **Security**

### Secure (Security)

From [AEL Secure](https://adaptive-enforcement-lab.com/secure/):

- **Cloud Native**
- **GitHub Actions Security Patterns Hub** — Complete security patterns for GitHub Actions…
- **GitHub Core App Setup** — Configure organization-level GitHub Apps for secure…
- **Go Security Tooling** — Standard Go security toolkit: race detector…
- **OpenSSF Scorecard Achievement Guide** — Complete OpenSSF Scorecard achievement guide.
- **Risk Management**

## Automated Generation

All skills in `plugins/` — and this README — are generated from AEL documentation — **never edit them by hand**:

- **Source**: [adaptive-enforcement-lab.com](https://github.com/adaptive-enforcement-lab/adaptive-enforcement-lab-com)
- **Generator**: `skillgen`, a Go extraction pipeline in this repo
- **Sync**: the `generate-skills.yml` workflow runs on `repository_dispatch` (`docs-updated`) from the docs repo, on manual `workflow_dispatch`, and on pull requests to `main`. Dispatch runs open or force-push a PR on the `chore/regenerate-skills` branch; on a release-please PR it commits regenerated output directly to that PR's branch. Regeneration is proposed automatically — merging is a human step.

The same run regenerates `.claude-plugin/marketplace.json`, every `plugins/*/.claude-plugin/plugin.json`, and this `README.md` from `plugin-metadata.json` plus `.release-please-manifest.json`.

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
    "build@ael-skills": true,
    "enforce@ael-skills": true,
    "patterns@ael-skills": true,
    "secure@ael-skills": true
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
│   └── skills/patterns/          # One hub skill directory
│       ├── SKILL.md              # Lean overview + grouped link index
│       ├── reference.md          # Full content, every topic, concatenated
│       └── library/              # Every source doc, verbatim, tree-mirrored
├── enforce/
├── secure/
└── build/

plugin-metadata.json              # Source of truth: descriptions, categories, tags
.release-please-manifest.json     # Source of truth: versions (release-please owned)
release-please-config.json        # Multi-component release configuration
README.md                         # This file (GENERATED)

skillgen/                         # Generator source
├── cmd/skillgen/                 # Entry point and dependency injection
├── internal/
│   ├── domain/                   # Core entities (Document, Skill, Marketplace, Readme)
│   ├── ports/                    # Interfaces
│   ├── adapters/                 # filesystem, parser, logger
│   └── services/                 # marketplace_generator.go, readme_generator.go, extractor/, generator/, validator/
└── templates/                    # skill.tmpl, reference.tmpl, readme.tmpl

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
  --release-manifest ./.release-please-manifest.json \
  --templates skillgen/templates
```

`--templates` is required in practice: the flag defaults to `./templates`, which does not exist at the repo root, so omitting it fails with `pattern matches no files`. Add `--verbose` for per-document logging. See [CLAUDE.md](CLAUDE.md) for the full flag reference and [CONTRIBUTING.md](CONTRIBUTING.md) for development guidelines.

## Architecture

`skillgen` follows Clean/Hexagonal Architecture — dependencies point inward:

- **Domain** (`skillgen/internal/domain`): core entities, no external dependencies
- **Ports** (`skillgen/internal/ports`): interfaces for external dependencies
- **Adapters** (`skillgen/internal/adapters`): filesystem, markdown parser, logger
- **Services** (`skillgen/internal/services`): marketplace and README generation, plus extractor/, generator/, validator/ subpackages

Pipeline: read every doc in a category → extract each as a lightweight `Topic` (title, one-line description, URL) → `HubBuilder` aggregates them into one hub `Skill` per category (grouped topics, plus each doc's full body assembled once) → render templates → write `SKILL.md`, `reference.md`, and the `library/` tree → generate marketplace files and this README.

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
