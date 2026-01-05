# Config-Driven Marketplace Generation

**Date:** 2026-01-05
**Status:** Approved for Implementation
**Author:** Adaptive Enforcement Lab

## Problem Statement

The skillgen component currently has partial, inconsistent marketplace.json management:

**Current Issues:**
- Only handles the "secure" plugin automatically via hardcoded `NewSecurePlugin()`
- Plugin metadata (descriptions, tags, categories) is hardcoded in Go code
- Versions are synced AFTER releases via separate `marketplace-version-sync.yml` workflow
- Manual marketplace.json maintenance is error-prone
- Per-collection `plugin.json` files exist in auto-generated `skills/` directory (violates "don't edit skills/" principle)
- Plugin name mismatch requires special handling ("enforce" → "enforce")

**Impact:**
- Maintenance burden for updating descriptions/metadata
- Risk of version drift between manifest and marketplace.json
- Confusion about what's auto-generated vs manually maintained
- Additional workflow complexity

## Solution Overview

Implement **full automation** for marketplace.json and per-collection plugin.json generation using:

1. **`plugin-metadata.json`** (repo root) - Static plugin metadata configuration
2. **`.release-please-manifest.json`** - Dynamic version data (existing)
3. **skillgen** reads both and generates all marketplace files

This eliminates manual maintenance, removes the version sync workflow, and ensures all files in `skills/` are truly auto-generated.

## Architecture

### Single Source of Truth Principle

Two sources of truth that skillgen combines:

1. **`plugin-metadata.json`** (repo root) - Static metadata
   - Descriptions, categories, tags, author info
   - Marketplace-level configuration
   - Manually maintained, version-controlled

2. **`.release-please-manifest.json`** - Dynamic versions
   - Current version for each skill collection
   - Managed by release-please
   - Never manually edited

### Data Flow

```
┌─────────────────────────┐
│ plugin-metadata.json    │ (Static metadata)
│ - descriptions          │
│ - categories/tags       │
│ - common config         │
└───────────┬─────────────┘
            │
            ├──────────────────────────┐
            │                          │
            ▼                          ▼
┌─────────────────────────┐  ┌──────────────────────────┐
│ .release-please-        │  │ skillgen reads both,     │
│ manifest.json           │─▶│ generates:               │
│ - skills/patterns: 0.2.1│  │                          │
│ - skills/enforce: 0.3.0 │  │ 1. marketplace.json      │
│ - skills/build: 0.1.4   │  │ 2. plugin.json per       │
│ - skills/secure: 0.2.2  │  │    collection            │
└─────────────────────────┘  └──────────────────────────┘
```

### Generated Artifacts

1. **`.claude-plugin/marketplace.json`** - Complete catalog with all plugins
2. **`skills/patterns/.claude-plugin/plugin.json`** - Per-collection manifest
3. **`skills/enforce/.claude-plugin/plugin.json`**
4. **`skills/build/.claude-plugin/plugin.json`**
5. **`skills/secure/.claude-plugin/plugin.json`**

All versions come from `.release-please-manifest.json`, all other metadata from `plugin-metadata.json`.

## Configuration File Structure

### plugin-metadata.json

Located at repo root: `plugin-metadata.json`

```json
{
  "$schema": "https://json-schema.org/draft-07/schema#",
  "marketplace": {
    "name": "ael-skills",
    "owner": {
      "name": "Adaptive Enforcement Lab",
      "email": "contact@adaptive-enforcement-lab.com"
    },
    "description": "Claude Code skills for secure development patterns, enforcement automation, and build engineering",
    "pluginRoot": "./skills"
  },
  "common": {
    "author": {
      "name": "Adaptive Enforcement Lab",
      "email": "contact@adaptive-enforcement-lab.com"
    },
    "homepage": "https://github.com/adaptive-enforcement-lab/claude-skills",
    "repository": "https://github.com/adaptive-enforcement-lab/claude-skills",
    "license": "MIT"
  },
  "plugins": {
    "patterns": {
      "description": "Reusable engineering patterns for error handling, state management, performance optimization, and resilience",
      "category": "development",
      "tags": ["patterns", "engineering", "best-practices"],
      "keywords": ["error-handling", "state-management", "performance", "resilience"]
    },
    "enforce": {
      "marketplaceName": "enforce",
      "description": "Security and compliance enforcement automation including pre-commit hooks, policy checks, and validation",
      "category": "security",
      "tags": ["security", "compliance", "enforcement", "devsecops"],
      "keywords": ["policy-as-code", "kyverno", "opa"]
    },
    "build": {
      "description": "Build engineering patterns for CI/CD pipelines, release automation, and deployment strategies",
      "category": "devops",
      "tags": ["ci-cd", "build", "release", "deployment"],
      "keywords": ["release-please", "github-actions"]
    },
    "secure": {
      "description": "Security patterns and hardening guides for cloud-native applications, GitHub Actions, and supply chain security",
      "category": "security",
      "tags": ["security", "hardening", "devsecops", "supply-chain"],
      "keywords": ["github-actions", "gke", "oidc"]
    }
  }
}
```

### Key Design Decisions

1. **Marketplace-level config** - Name, owner, description, pluginRoot
2. **Common fields** - Applied to all plugin.json files (DRY principle)
3. **Per-plugin config** - Unique descriptions, categories, tags
4. **Optional `marketplaceName`** - Handles edge cases where marketplace name differs from plugin key
5. **Separation of tags vs keywords**:
   - `tags` → Used in marketplace.json for discovery
   - `keywords` → Used in plugin.json for additional SEO/context

### Version Mapping

From `.release-please-manifest.json`:
```json
{
  "skills/patterns": "0.2.1",
  "skills/enforce": "0.3.0",
  "skills/build": "0.1.4",
  "skills/secure": "0.2.2"
}
```

Skillgen extracts the collection name from the path (`patterns`, `enforce`, etc.) and matches it to `plugins` config.

## Implementation Details

### New Domain Models

**`internal/domain/plugin_config.go`**:

```go
// PluginMetadata represents the plugin-metadata.json configuration
type PluginMetadata struct {
    Marketplace MarketplaceConfig         `json:"marketplace"`
    Common      CommonPluginFields        `json:"common"`
    Plugins     map[string]PluginConfig   `json:"plugins"`
}

// MarketplaceConfig defines marketplace-level settings
type MarketplaceConfig struct {
    Name        string            `json:"name"`
    Owner       MarketplaceOwner  `json:"owner"`
    Description string            `json:"description"`
    PluginRoot  string            `json:"pluginRoot"`
}

// CommonPluginFields applied to all plugin.json files
type CommonPluginFields struct {
    Author     *PluginAuthor `json:"author"`
    Homepage   string        `json:"homepage"`
    Repository string        `json:"repository"`
    License    string        `json:"license"`
}

// PluginConfig per-collection configuration
type PluginConfig struct {
    MarketplaceName string   `json:"marketplaceName,omitempty"`
    Description     string   `json:"description"`
    Category        string   `json:"category"`
    Tags            []string `json:"tags"`
    Keywords        []string `json:"keywords"`
}

// PluginManifest represents an individual plugin.json file
type PluginManifest struct {
    Name        string        `json:"name"`
    Description string        `json:"description"`
    Version     string        `json:"version"`
    Author      *PluginAuthor `json:"author"`
    Homepage    string        `json:"homepage"`
    Repository  string        `json:"repository"`
    License     string        `json:"license"`
    Keywords    []string      `json:"keywords"`
}
```

### New Ports

**`internal/ports/config_reader.go`**:

```go
// ConfigReader reads plugin metadata configuration
type ConfigReader interface {
    // ReadPluginMetadata reads plugin-metadata.json
    ReadPluginMetadata(path string) (*domain.PluginMetadata, error)

    // ReadReleaseManifest reads .release-please-manifest.json
    ReadReleaseManifest(path string) (map[string]string, error)
}
```

**Update `internal/ports/writer.go`**:

```go
type MarketplaceWriter interface {
    // GenerateFromConfig builds marketplace.json from config + versions
    GenerateFromConfig(
        metadata *domain.PluginMetadata,
        versions map[string]string,
        outputPath string,
    ) error

    // WritePluginManifest writes individual plugin.json files
    WritePluginManifest(
        manifest *domain.PluginManifest,
        outputPath string,
    ) error
}
```

### New Service

**`internal/services/marketplace_generator.go`**:

```go
type MarketplaceGenerator struct {
    configReader ConfigReader
    writer       MarketplaceWriter
    logger       Logger
}

func (g *MarketplaceGenerator) Generate(
    metadataPath string,
    manifestPath string,
    outputDir string,
) error {
    // 1. Read plugin-metadata.json
    metadata, err := g.configReader.ReadPluginMetadata(metadataPath)
    if err != nil {
        return fmt.Errorf("failed to read plugin metadata: %w", err)
    }

    // 2. Read .release-please-manifest.json
    versions, err := g.configReader.ReadReleaseManifest(manifestPath)
    if err != nil {
        return fmt.Errorf("failed to read release manifest: %w", err)
    }

    // 3. Generate marketplace.json
    marketplacePath := filepath.Join(".claude-plugin", "marketplace.json")
    err = g.writer.GenerateFromConfig(metadata, versions, marketplacePath)
    if err != nil {
        return fmt.Errorf("failed to generate marketplace.json: %w", err)
    }

    // 4. Generate each plugin.json
    for pluginKey, pluginConfig := range metadata.Plugins {
        version := extractVersionForPlugin(versions, pluginKey)
        if version == "" {
            g.logger.Warn("no version found for plugin", "plugin", pluginKey)
            version = "0.0.0"
        }

        manifest := buildPluginManifest(pluginKey, &pluginConfig, &metadata.Common, version)
        pluginPath := filepath.Join(outputDir, pluginKey, ".claude-plugin/plugin.json")

        // Ensure directory exists
        if err := os.MkdirAll(filepath.Dir(pluginPath), 0755); err != nil {
            return fmt.Errorf("failed to create plugin directory: %w", err)
        }

        if err := g.writer.WritePluginManifest(manifest, pluginPath); err != nil {
            return fmt.Errorf("failed to write plugin manifest for %s: %w", pluginKey, err)
        }

        g.logger.Info("generated plugin manifest", "plugin", pluginKey, "version", version)
    }

    return nil
}

func extractVersionForPlugin(versions map[string]string, pluginKey string) string {
    // Match "skills/{pluginKey}" in manifest
    manifestKey := fmt.Sprintf("skills/%s", pluginKey)
    return versions[manifestKey]
}

func buildPluginManifest(
    name string,
    config *domain.PluginConfig,
    common *domain.CommonPluginFields,
    version string,
) *domain.PluginManifest {
    return &domain.PluginManifest{
        Name:        name,
        Description: config.Description,
        Version:     version,
        Author:      common.Author,
        Homepage:    common.Homepage,
        Repository:  common.Repository,
        License:     common.License,
        Keywords:    config.Keywords,
    }
}
```

### Integration with Existing Code

**Update `cmd/skillgen/main.go`**:

```go
var (
    sourcePath          string
    outputPath          string
    marketplacePath     string
    templatesPath       string
    pluginMetadataPath  string
    releaseManifestPath string
    verbose             bool
    showVersion         bool
)

flag.StringVar(&pluginMetadataPath, "plugin-metadata", "./plugin-metadata.json",
    "Path to plugin metadata config")
flag.StringVar(&releaseManifestPath, "release-manifest", "./.release-please-manifest.json",
    "Path to release-please manifest")

// ... after skill generation ...

// Generate marketplace files
logger.Info("generating marketplace files")
configReader := filesystem.NewConfigReader(fs)
marketplaceGen := services.NewMarketplaceGenerator(configReader, marketplaceWriter, logger)

err = marketplaceGen.Generate(
    pluginMetadataPath,
    releaseManifestPath,
    outputPath,
)
if err != nil {
    logger.Error("failed to generate marketplace files", "error", err)
    errors++
} else {
    logger.Info("marketplace files generated successfully")
}
```

### Error Handling

| Condition | Behavior |
|-----------|----------|
| Missing `plugin-metadata.json` | Fatal error with clear message |
| Missing `.release-please-manifest.json` | Fatal error |
| Plugin in config but no version in manifest | Use "0.0.0" with warning |
| Invalid JSON in config | Fatal error with validation details |
| Plugin directory doesn't exist | Create it (normal for new plugins) |

## Workflow Changes

### 1. Update `generate-skills.yml`

**Add new flags:**

```yaml
- name: Build skillgen
  run: |
    cd skillgen
    go build -o ../bin/skillgen ./cmd/skillgen
    cd ..

- name: Generate skills and marketplace
  run: |
    ./bin/skillgen \
      --source ../adaptive-enforcement-lab-com/docs \
      --output skills \
      --plugin-metadata ./plugin-metadata.json \
      --release-manifest ./.release-please-manifest.json \
      --verbose
```

**Update change detection:**

```yaml
if [ -z "$(git status --porcelain skills/ .claude-plugin/)" ]; then
  echo "No changes detected in skills/ or .claude-plugin/, skipping PR creation"
  exit 0
fi
```

**Update PR checklist:**

```markdown
## Generated Changes

- ✅ Skills regenerated from AEL documentation
- ✅ Updated marketplace.json with all plugins
- ✅ Generated plugin.json for each collection
- ✅ Versions synced from .release-please-manifest.json

## Manual Review Required

- [ ] Verify marketplace.json is valid JSON
- [ ] Verify all plugin.json files generated correctly
- [ ] Confirm versions match .release-please-manifest.json
- [ ] Check no manual edits needed in skills/ directory
```

### 2. Delete `marketplace-version-sync.yml`

**Why it's obsolete:**
- skillgen now reads `.release-please-manifest.json` directly during generation
- Versions are always in sync automatically
- No manual syncing needed

**Replacement logic:**
- When docs change → `generate-skills.yml` runs → reads current versions from manifest
- When release PR merges → next `generate-skills.yml` run picks up new versions
- Versions update atomically with skill content

### 3. No changes to `release.yml`

Current workflow already works correctly:
- Release-please updates `.release-please-manifest.json`
- Creates tags for each component
- Next skill generation will pick up new versions automatically

## Migration Plan

### Phase 1: Add Config File
1. Create `plugin-metadata.json` in repo root
2. Populate from existing `marketplace.json` and `skills/*/​.claude-plugin/plugin.json`
3. Validate JSON structure
4. Commit to main branch

### Phase 2: Implement Skillgen Changes
1. Add domain models (`internal/domain/plugin_config.go`)
2. Add ports (`internal/ports/config_reader.go`, update `writer.go`)
3. Implement adapters (`internal/adapters/filesystem/config_reader.go`, update `writer.go`)
4. Implement service (`internal/services/marketplace_generator.go`)
5. Update `cmd/skillgen/main.go` with new flags and generation logic
6. Add unit tests for all new components

### Phase 3: Test Locally
1. Run skillgen with new flags
2. Verify generated `marketplace.json` matches current (except formatting)
3. Verify all `plugin.json` files match current
4. Validate JSON structure against Claude Code requirements
5. Test with different version scenarios

### Phase 4: Update Workflows
1. Update `generate-skills.yml` with new flags
2. Create test PR to verify generation works in CI
3. Review generated output in PR
4. Merge workflow changes

### Phase 5: Remove Obsolete Workflow
1. Delete `.github/workflows/marketplace-version-sync.yml`
2. Document the change in CHANGELOG.md
3. Update CLAUDE.md to remove references to manual version syncing
4. Add note that `plugin-metadata.json` is now the source of truth for metadata

### Rollback Plan

If issues arise:
1. Revert `generate-skills.yml` to previous version
2. Restore `marketplace-version-sync.yml`
3. Manually fix `marketplace.json` if needed
4. Fix skillgen bugs and retry migration

## Testing Strategy

### Unit Tests

1. **Config reading:**
   - Valid plugin-metadata.json parsing
   - Missing required fields
   - Invalid JSON syntax
   - Valid release manifest parsing
   - Version extraction from manifest keys

2. **Marketplace generation:**
   - Complete marketplace.json structure
   - Plugin array population
   - Version mapping from manifest
   - Marketplace metadata preservation

3. **Plugin manifest generation:**
   - Individual plugin.json structure
   - Common fields applied correctly
   - Keywords vs tags separation
   - MarketplaceName override handling

### Integration Tests

1. Run skillgen with real config files
2. Compare generated files with expected output
3. Validate JSON structure
4. Verify all plugins have correct versions
5. Test error conditions (missing files, invalid JSON)

### CI Validation

1. Generate skills in clean environment
2. Validate marketplace.json structure
3. Validate all plugin.json files exist
4. Check versions match manifest
5. Ensure no unexpected files in skills/

## Benefits

1. **Single Command Generation** - One skillgen run creates everything
2. **Version Consistency** - Versions always match release-please manifest
3. **Reduced Complexity** - Removes version sync workflow
4. **True Auto-Generation** - All files in `skills/` are generated, no manual edits
5. **Maintainability** - Metadata in one config file, easy to update
6. **Type Safety** - Go structs enforce structure, catch errors early
7. **Auditability** - Config file is version-controlled, changes tracked in git

## Risks & Mitigations

| Risk | Impact | Mitigation |
|------|--------|------------|
| Config file gets out of sync | Incorrect metadata in generated files | Add validation step in CI, document config structure |
| Version extraction breaks | Plugins get version 0.0.0 | Warn loudly, test with multiple manifest formats |
| Workflow changes break CI | Can't generate skills | Test in feature branch first, keep rollback plan ready |
| Plugin.json format changes | Claude Code can't read files | Follow official schema, validate output |

## Success Criteria

- [ ] `plugin-metadata.json` created and populated
- [ ] Skillgen generates `marketplace.json` with all 4 plugins
- [ ] Skillgen generates `plugin.json` for each collection
- [ ] Versions match `.release-please-manifest.json` exactly
- [ ] `generate-skills.yml` workflow runs successfully
- [ ] `marketplace-version-sync.yml` deleted
- [ ] CLAUDE.md updated with new process
- [ ] All tests pass
- [ ] No manual edits needed in `skills/` directory

## Future Enhancements

1. **Schema Validation** - Add JSON schema for `plugin-metadata.json`
2. **Plugin Discovery** - Auto-detect new plugin directories
3. **Metadata Validation** - Check descriptions meet length requirements
4. **Changelog Integration** - Auto-generate plugin descriptions from CHANGELOG.md
5. **Multi-Marketplace Support** - Generate for different distribution channels

## References

- [Claude Code Plugin Marketplaces Documentation](https://code.claude.com/docs/en/plugin-marketplaces.md)
- [Claude Code Skills Documentation](https://code.claude.com/docs/en/skills.md)
- [Release-Please Manifest Configuration](https://github.com/googleapis/release-please/blob/main/docs/manifest-releaser.md)
- Current implementation: `skillgen/internal/domain/marketplace.go`
- Current workflow: `.github/workflows/marketplace-version-sync.yml`
