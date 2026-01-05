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
  --release-manifest ./.release-please-manifest.json

# Run with verbose logging
./bin/skillgen \
  --source ../adaptive-enforcement-lab-com/docs \
  --output plugins \
  --plugin-metadata ./plugin-metadata.json \
  --release-manifest ./.release-please-manifest.json \
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
- `--templates`: Path to template directory (default: `./templates`)
- `--marketplace`: Path to marketplace.json (DEPRECATED - now auto-generated)
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
    services/          → Application services (extractor, generator)
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

- `patterns/` → Development patterns (error handling, state management, etc.)
- `enforce/` → Security enforcement automation (pre-commit hooks, policy validation)
- `build/` → Build engineering patterns (CI/CD, release automation)
- `secure/` → Security patterns and practices

Blog posts (detected via frontmatter `date`/`authors` fields) are automatically skipped.

## Templates

Templates live in `templates/` and use Go's text/template syntax:

- `skill.tmpl` - Base SKILL.md template
- `pattern_skill.tmpl` - Pattern-specific skills
- `enforce_skill.tmpl` - Enforcement-specific skills
- `build_skill.tmpl` - Build-specific skills
- `examples.tmpl` - Examples documentation
- `reference.tmpl` - Reference documentation
- `troubleshooting.tmpl` - Troubleshooting guides

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

## Conventional Commits

Use these commit prefixes for release-please automation:

- `feat:` → Minor version bump, new features
- `fix:` → Patch version bump, bug fixes
- `chore:` → Maintenance, no version bump
- `docs:` → Documentation only, hidden from changelog
- `refactor:` → Code refactoring, appears in changelog
- `test:` → Test changes, hidden from changelog
- `perf:` → Performance improvements, appears in changelog

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
The generator logs errors but exits with code 0 even when errors occur. Many errors are expected (missing titles, malformed content) and shouldn't fail CI builds. The generation summary reports error counts for visibility.

## Testing Strategy

- Unit tests for extractors and parsers
- Edge cases: empty content, missing sections, malformed markdown
- Table-driven tests for consistent coverage
- No integration tests yet (future enhancement)
- **CI/CD**: All workflows run `go test ./...` before building binaries

## Dependencies

Go 1.25+ with minimal external dependencies:
- `github.com/yuin/goldmark` - Markdown parsing
- `gopkg.in/yaml.v3` - YAML frontmatter parsing

## Common Pitfalls

1. **Editing plugins/ directly** - These are auto-generated, edits will be overwritten
2. **Editing .claude-plugin/marketplace.json directly** - Auto-generated from plugin-metadata.json
3. **Editing plugins/*/​.claude-plugin/plugin.json directly** - Auto-generated from plugin-metadata.json
4. **Forgetting --plugin-metadata and --release-manifest flags** - Required for marketplace generation
5. **Forgetting --source flag** - Generator requires source docs path
6. **Assuming specific section names** - Source docs vary, extractor uses fuzzy matching
7. **Breaking template syntax** - Go templates are whitespace-sensitive
8. **Not testing with actual docs** - Clone AEL docs repo for realistic testing

## Updating Plugin Metadata

To change plugin descriptions, categories, or tags:
1. Edit `plugin-metadata.json` in the repo root
2. Run skillgen to regenerate marketplace files
3. Commit both plugin-metadata.json and generated files

Versions are automatically synchronized from `.release-please-manifest.json` - never edit them manually.
