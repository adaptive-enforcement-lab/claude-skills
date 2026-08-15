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
- `--readme`: Output path for the generated README.md (default: `./README.md`). Same as `--marketplace`: controls where it's written, not whether it's generated.
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

skillgen produces exactly **one hub skill per category** (4 total: patterns, enforce, build, secure): a lean `SKILL.md` (short overview + grouped index of links out to the upstream docs) plus a `reference.md` leaf holding the full offline depth behind it — not one skill per source document, and not a link-only index that depends on live fetches to go deep.

1. **DocumentReader** (adapter) reads every index.md under a category
2. **FrontmatterParser** + **SectionParser** extract title, description, and introduction; **AdmonitionConverter** cleans up VitePress admonitions in each doc's raw body for reference.md
3. **TopicExtractor** (service) turns a single document into a lightweight `Topic` (title, one-line description, URL) for the `SKILL.md` index — no prose extraction
4. **HubBuilder** (service) aggregates a category's documents into one hub `Skill`: the category root doc supplies `Overview`, every other doc is grouped by its first path segment under the category and becomes a linked `Topic`. It also assembles each doc's full body **exactly once** (via `prepareReferenceBody`, a pure heading-shift helper — no fuzzy section matching) into `SkillMetadata.ReferenceBody` / `TopicGroup.ReferenceBody` / `Topic.ReferenceBody` for `reference.md`
5. **TemplateRenderer** (service) applies `skill.tmpl` and `reference.tmpl` to generate the hub's two files
6. **SkillWriter** (adapter) removes stale sibling skill directories, then writes `SKILL.md`, `reference.md`, and the `library/` tree
7. **MarketplaceGenerator** (service) reads plugin-metadata.json and .release-please-manifest.json
8. **MarketplaceWriter** (adapter) generates marketplace.json and all plugin.json files
9. **ReadmeGenerator** (service) builds the repo root `README.md` from the same hub `Skill`s, plugin metadata, and release versions used above — never hand-maintained, so it can't drift the way the old static README did (it said "118 skills" for weeks after the hub-skill rewrite)

### Domain Models

**Document** (`internal/domain/document.go`):
- Represents parsed AEL documentation
- Contains frontmatter, sections, code blocks, mermaid diagrams
- Source for topic extraction

**Skill** (`internal/domain/skill.go`):
- One hub per category: `SkillMetadata` (name, title, description, overview, `ReferenceBody`) plus `Groups []TopicGroup`
- Each `TopicGroup` is a themed cluster of `Topic` entries (title, one-line description, URL), each also carrying a `ReferenceBody` — the topic's full cleaned content for `reference.md`
- A hub skill produces `SKILL.md` (the scannable index, fans out to the upstream docs), `reference.md` (the full depth behind it, offline), and `library/` (every source doc shipped verbatim, one file per doc, mirroring the docs tree) — there are no examples/troubleshooting files or extracted scripts

**Marketplace** (`internal/domain/marketplace.go`):
- Represents .claude-plugin/marketplace.json structure
- Defines available plugin collections (patterns, enforcement, build, secure)

**Readme** (`internal/domain/readme.go`):
- `ReadmeData` (marketplace name/owner/description, `Hubs []ReadmeHub`) is the data passed to `readme.tmpl`
- Each `ReadmeHub` carries the fields that used to be hand-typed and go stale: `Version` (from the release manifest), `TopicCount` (`len(hub.LibraryFiles)`, i.e. the real count of docs shipped), `Focus` (the plugin description, word-truncated for a table cell), and `Groups` (reused directly from the hub `Skill`, so the per-collection bullet lists are the same data `SKILL.md` shows)

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

Templates live in `skillgen/templates/` and use Go's text/template syntax. There are three:

- `skill.tmpl` - SKILL.md: short overview + grouped topic index, links out to the upstream docs
- `reference.tmpl` - reference.md: the full offline depth behind every topic, assembled from each doc's body exactly once
- `readme.tmpl` - the repo root README.md: available-skills table, per-collection topic bullets, and static usage/architecture sections

There are no per-category templates and no examples/troubleshooting templates. `library/` files have no template — each is the doc's own body shipped verbatim (see "Reference Body Assembly" below).

## Configuration Files

### `plugin-metadata.json`
**Source of truth for plugin metadata** (descriptions, categories, tags):
- Located at repo root
- Manually maintained, version-controlled
- Defines marketplace-level config (name, owner, description)
- Contains common fields applied to all plugin.json files (author, license, homepage)
- Per-plugin configuration (descriptions, categories, tags, keywords)
- Combined with `.release-please-manifest.json` to generate all marketplace files
- Each plugin's `description` is triple-purpose: the marketplace catalog blurb, the hub `SKILL.md`'s frontmatter description (via `pluginCfg.Description` in `HubBuilder`), and the source for README.md's truncated "Focus" column (via `ReadmeGenerator`). Per `superpowers:writing-skills`, write it as a triggering condition ("Use when...", concrete situations/symptoms), not a content summary — that's what Claude uses to decide whether to load the skill.

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

### Hub Skill Naming
A hub skill's name is simply its category slug (`patterns`, `enforce`, `build`, `secure`) — there's no per-document name derivation anymore.

### Topic Grouping
The **HubBuilder** (`internal/services/extractor/hub_builder.go`) groups a category's documents by their first path segment under the category (e.g. `patterns/architecture/...` → the "Architecture" group). If that segment has its own `index.md`, its title/description/URL become the group heading; otherwise the slug is humanized as a fallback. No deeper sub-grouping is attempted — this keeps each hub scannable rather than deeply nested. `TopicExtractor` reads each document's frontmatter `title`/`description` directly (falling back to the first sentence of the intro) rather than extracting prose sections.

`TopicExtractor`/`HubBuilder` cap every one-line description (topic and group) to `maxTopicDescriptionWords` (6, in `topic_extractor.go`) so a hub with 20-30 topics fits SKILL.md's ~500-word budget (per `superpowers:writing-skills`) with real margin, not right at the wire — at 6 words the four hubs land at 165-442 words. This is purely a SKILL.md concern — `reference.md` and `library/` always carry the untruncated text, so nothing is lost, just deferred to the on-demand files. If a future doc set pushes a hub back over budget, tighten this constant (or `firstSentences`' sentence count for the root `Overview`) before reaching for deeper structural changes.

### Reference Body Assembly
`reference.md`'s full-depth content comes from `prepareReferenceBody` (`internal/services/extractor/reference_body.go`): a pure line-based helper that drops a doc's leading `# Title` heading and shifts every remaining heading deeper by a fixed number of levels, so it nests under the heading the caller wraps it in. Each doc's body flows through this exactly once, at one shift level per role (category root, group root, or topic) — there is no fuzzy keyword matching and no recursion into already-consumed subsections, which is what caused the old per-doc pipeline to triplicate content into `SKILL.md`. When editing this, keep the invariant: every doc contributes its body to exactly one place in the reference tree.

### Library Files
`HubBuilder.libraryFile` (`internal/services/extractor/hub_builder.go`) ships every doc a hub processes as its own file under `library/`, mirroring the doc's path under the category (e.g. `patterns/architecture/hub-and-spoke/index.md` → `library/architecture/hub-and-spoke/index.md`; the category root doc becomes `library/index.md`). Unlike `reference.md`, nothing is stripped or heading-shifted here — each file keeps its doc's natural title and heading structure, with a `Source: <url>` line inserted right after the title. This exists so the complete, unmerged source library is available in the tree in addition to the curated `SKILL.md` index and the merged `reference.md` — three ways to consume the same content depending on what's needed (routing, curated depth, or the raw original doc).

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
4. **Editing README.md directly** - Auto-generated by `ReadmeGenerator` from the same hub skills, plugin-metadata.json, and release manifest; edits will be overwritten on the next run
5. **Forgetting --plugin-metadata and --release-manifest flags** - Required for marketplace and README generation
6. **Forgetting --source flag** - Generator requires source docs path
7. **Forgetting --templates skillgen/templates** - The default `./templates` does not exist at the repo root, so the run dies before generating anything
8. **Assuming a doc without its own index.md still gets a group heading** - it falls back to a humanized slug instead
9. **Breaking template syntax** - Go templates are whitespace-sensitive
10. **Not testing with actual docs** - Clone AEL docs repo for realistic testing

## Updating Plugin Metadata

To change plugin descriptions, categories, or tags:
1. Edit `plugin-metadata.json` in the repo root
2. Run skillgen to regenerate marketplace files
3. Commit both plugin-metadata.json and generated files

Versions are automatically synchronized from `.release-please-manifest.json` - never edit them manually.
