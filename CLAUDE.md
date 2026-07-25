# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Go-based skill generator that automatically transforms AEL (Adaptive Enforcement Lab) documentation into Claude Code skills. The project follows Clean/Hexagonal Architecture and uses release-please for automated releases.

**CRITICAL**: The `plugins/` directory contains auto-generated files. Never manually edit files in `plugins/` - they are regenerated from source documentation at [adaptive-enforcement-lab.com](https://github.com/adaptive-enforcement-lab/adaptive-enforcement-lab-com).

## Build and Development Commands

### Core Commands

```bash
# Build the generator (from skillgen directory)
cd skillgen && go build -o ../bin/skillgen ./cmd/skillgen && cd ..

# Run the generator (requires AEL docs locally)
./bin/skillgen \
  --source ../adaptive-enforcement-lab-com/docs \
  --output plugins \
  --plugin-metadata ./plugin-metadata.json \
  --release-manifest ./.release-please-manifest.json \
  --templates skillgen/templates

# Run with verbose logging
./bin/skillgen \
  --source ../adaptive-enforcement-lab-com/docs \
  --output plugins \
  --plugin-metadata ./plugin-metadata.json \
  --release-manifest ./.release-please-manifest.json \
  --templates skillgen/templates \
  --verbose

# Run tests (from skillgen directory)
cd skillgen && go test ./... && cd ..

# Format code
gofmt -w skillgen/
```

### Generator Options

- `--source`: Path to AEL documentation source (required)
- `--output`: Output path for generated plugins (default: `./plugins`)
- `--plugin-metadata`: Path to plugin metadata config (default: `./plugin-metadata.json`)
- `--release-manifest`: Path to release-please manifest (default: `./.release-please-manifest.json`)
- `--templates`: Path to template directory (default: `./templates`). **Effectively required**: the default does not exist at the repo root, so omitting the flag fails with `pattern matches no files`. Pass `skillgen/templates`.
- `--marketplace`: Output path for the generated marketplace.json (default: `./.claude-plugin/marketplace.json`). The file's *content* is generated, but this flag still controls where it is written.
- `--verbose`: Enable verbose logging
- `--version`: Show version and exit

## Architecture

### Clean/Hexagonal Architecture

```
skillgen/
  cmd/skillgen/        → Entry point and dependency injection
  internal/
    domain/            → Core entities (Skill, Document, Marketplace)
    ports/             → Interfaces for external dependencies
    adapters/          → Implementations (filesystem, parser, logger)
    services/          → Application services (extractor, generator, validator,
                         marketplace generator)
  templates/           → Go text/template files
```

**Key Principles:**
- Domain layer has no external dependencies
- Ports define interfaces, adapters implement them
- Dependencies point inward (adapters → ports → domain)
- Services orchestrate business logic using ports

### Data Flow

1. **DocumentReader** (adapter) reads index.md files from source docs
2. **FrontmatterParser** + **SectionParser** extract structured content
3. **SkillExtractor** (service) transforms Document → Skill using business rules
4. **TemplateRenderer** (service) applies Go templates to generate markdown
5. **SkillWriter** (adapter) writes SKILL.md files to filesystem
6. **MarketplaceGenerator** (service) reads plugin-metadata.json and .release-please-manifest.json
7. **MarketplaceWriter** (adapter) generates marketplace.json and all plugin.json files

### Domain Models

**Document** (`internal/domain/document.go`):
- Represents parsed AEL documentation
- Contains frontmatter, sections, code blocks, mermaid diagrams
- Source for skill extraction

**Skill** (`internal/domain/skill.go`):
- Output model with metadata, main content, optional examples/reference/troubleshooting
- Each skill may generate multiple files: SKILL.md, examples.md, troubleshooting.md, reference.md

**Marketplace** (`internal/domain/marketplace.go`):
- Represents .claude-plugin/marketplace.json structure
- Defines available plugin collections (patterns, enforcement, build, secure)

**PluginMetadata** (`internal/domain/plugin_config.go`):
- Represents plugin-metadata.json configuration
- Source of truth for plugin descriptions, categories, and tags
- Combined with .release-please-manifest.json for version data

## Skill Generation Categories

The generator processes 4 documentation categories:

- `patterns/` → Automation and architecture patterns (GitHub Actions, Argo, GitHub App auth, idempotency)
- `enforce/` → Policy-as-code enforcement (Kyverno and OPA templates, SLSA, SDLC hardening)
- `build/` → Build engineering (Go CLIs, release-please, packaging, versioned docs)
- `secure/` → Security hardening (GitHub Actions, secrets, runners, OIDC, GKE)

The canonical list lives in `skillgen/internal/domain/category.go` (`domain.Categories`). Note the category is `enforce`, not `enforcement`.

Blog posts (detected via frontmatter `date`/`authors` fields) are automatically skipped.

## Templates

Templates live in `skillgen/templates/` and use Go's text/template syntax. There are four:

- `skill.tmpl` - SKILL.md template, used for every category
- `examples.tmpl` - Examples documentation
- `reference.tmpl` - Reference documentation
- `troubleshooting.tmpl` - Troubleshooting guides

There are no per-category templates; `skill.tmpl` renders all four collections.

## Configuration Files

### `plugin-metadata.json`
**Source of truth for plugin metadata** (descriptions, categories, tags):
- Located at repo root
- Manually maintained, version-controlled
- Defines marketplace-level config (name, owner, description)
- Contains common fields applied to all plugin.json files (author, license, homepage)
- Per-plugin configuration (descriptions, categories, tags, keywords)
- Combined with `.release-please-manifest.json` to generate all marketplace files

### `.release-please-manifest.json`
**Source of truth for versions**:
- Managed by release-please automatically
- Maps skill collections to semantic versions
- Never manually edited
- Read by skillgen to populate version fields in marketplace.json and plugin.json files

## CI/CD Workflows

### `generate-skills.yml`
- Triggers: manual (`workflow_dispatch`) or repository_dispatch from docs repo
- Checks out both claude-skills and AEL docs repos
- Builds generator and runs skill generation with `--plugin-metadata` and `--release-manifest` flags
- Generates all marketplace files automatically:
  - `.claude-plugin/marketplace.json`
  - `plugins/*/​.claude-plugin/plugin.json` for each collection
- Creates idempotent PR with branch `chore/regenerate-skills`
- PR is reused for subsequent runs (force push updates)

### `release.yml`
- Automated releases via release-please
- Conventional commits determine version bumps
- Generates CHANGELOG.md automatically
- Publishes GitHub releases with binaries for Linux, macOS, Windows

### Multi-Component Versioning

Release-please manages 6 independent components:
- `skillgen` (Go binary) - main version
- `marketplace` (.claude-plugin/) - marketplace metadata
- `patterns` (plugins/patterns/) - pattern skills collection
- `enforce` (plugins/enforce/) - enforcement skills collection
- `build` (plugins/build/) - build skills collection
- `secure` (plugins/secure/) - secure skills collection

Each uses separate-pull-requests for independent versioning.

**A commit bumps only the components whose paths it touches.** This is the single most surprising part of the setup: one commit spanning several component directories opens one release PR *per component*. A `chore:` commit that edited `.claude-plugin/marketplace.json`, all four `plugins/*/.claude-plugin/plugin.json`, and two files under `skillgen/` opened **six** release PRs at once (#84-#89), one per package.

`always-update: true` keeps existing release PRs refreshed as `main` moves. It does not create a release PR on its own — a `docs:`-titled commit produced none.

**Merge release PRs one at a time.** All of them modify `.release-please-manifest.json`, so merging one leaves the rest out of date. Let the Release run finish before merging the next; release-please updates the remaining PRs, and `generate-skills.yml` re-syncs their `marketplace.json`. Merge the four collections first, then `marketplace` (its `marketplace.json` aggregates the collection versions), then `skillgen` last, since its release dispatches the binary build.

## Conventional Commits

Use these commit prefixes for release-please automation. The behaviour below comes from `changelog-sections` in `release-please-config.json` — a type listed as visible there is releasable, and a type marked `hidden: true` is not:

| Prefix | Bump | Changelog section |
| ------ | ---- | ----------------- |
| `feat:` | minor | Features |
| `fix:` | patch | Bug Fixes |
| `perf:` | patch | Performance |
| `refactor:` | patch | Code Refactoring |
| `chore:` | **patch** | Maintenance |
| `docs:` | none | hidden |
| `test:` | none | hidden |
| `ci:` | none | hidden |

**`chore:` bumps the version in this repo.** It is configured as a visible "Maintenance" section, so it is releasable — unlike the release-please default, where `chore` is inert. Use `docs:`, `test:` or `ci:` when you want a change to land without cutting a release.

**Squash merge means the PR title is the commit release-please parses.** The repo allows squash merge only, so the individual commits on a branch are discarded and the PR title becomes the conventional commit. A branch of tidy `fix:` commits merged under a `chore:` title releases as a patch chore; the reverse is also true. Title the PR as the release you intend.

## Key Implementation Notes

### Name Derivation
Skill names are auto-generated from document titles:
- Convert to lowercase
- Replace spaces with hyphens
- Remove special characters
- Example: "Error Handling: Fail Fast" → "error-handling-fail-fast"

### Section Mapping
The **SectionMapper** (`internal/services/extractor`) maps source doc sections to skill components:
- "Why It Matters" → `WhenToUse`
- "Prerequisites" → `Prerequisites`
- Custom logic determines what goes into SKILL.md vs examples.md vs reference.md

### Admonition Conversion
Source docs use VitePress admonitions (`::: tip`, `::: warning`). The **AdmonitionConverter** transforms these to standard markdown for Claude Code compatibility.

### Error Handling Philosophy
**Per-document** errors are logged but do not change the exit code. Many are expected (missing titles, malformed content) and shouldn't fail CI builds; the generation summary reports error counts for visibility.

**Startup** failures are fatal (`log.Fatal`, exit 1): a missing `--source`, a template directory that loads no `*.tmpl` files, or a failure to walk the source tree. A run that exits 1 with no summary means the generator never got started, not that documents failed.

## Testing Strategy

- Unit tests for extractors and parsers
- Edge cases: empty content, missing sections, malformed markdown
- Table-driven tests for consistent coverage
- No integration tests yet (future enhancement)
- **CI/CD**: All workflows run `go test ./...` before building binaries

## Dependencies

Go 1.26+ (`skillgen/go.mod`, matching the version pinned in CI) with minimal external dependencies:
- `github.com/yuin/goldmark` - Markdown parsing
- `gopkg.in/yaml.v3` - YAML frontmatter parsing

## Common Pitfalls

1. **Editing plugins/ directly** - These are auto-generated, edits will be overwritten
2. **Editing .claude-plugin/marketplace.json directly** - Auto-generated from plugin-metadata.json
3. **Editing plugins/*/​.claude-plugin/plugin.json directly** - Auto-generated from plugin-metadata.json
4. **Forgetting --plugin-metadata and --release-manifest flags** - Required for marketplace generation
5. **Forgetting --source flag** - Generator requires source docs path
6. **Forgetting --templates skillgen/templates** - The default `./templates` does not exist at the repo root, so the run dies before generating anything
7. **Assuming specific section names** - Source docs vary, extractor uses fuzzy matching
8. **Breaking template syntax** - Go templates are whitespace-sensitive
9. **Not testing with actual docs** - Clone AEL docs repo for realistic testing

## Updating Plugin Metadata

To change plugin descriptions, categories, or tags:
1. Edit `plugin-metadata.json` in the repo root
2. Run skillgen to regenerate marketplace files
3. Commit both plugin-metadata.json and generated files

Versions are automatically synchronized from `.release-please-manifest.json` - never edit them manually.
