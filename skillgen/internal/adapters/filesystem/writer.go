package filesystem

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"sort"
	"strings"

	"github.com/adaptive-enforcement-lab/claude-skills/skillgen/internal/domain"
	"github.com/adaptive-enforcement-lab/claude-skills/skillgen/internal/ports"
)

// SkillWriter implements ports.SkillWriter using the filesystem.
type SkillWriter struct {
	fs       ports.FileSystem
	renderer ports.TemplateRenderer
}

// NewSkillWriter creates a new filesystem-based skill writer.
func NewSkillWriter(fs ports.FileSystem, renderer ports.TemplateRenderer) *SkillWriter {
	return &SkillWriter{
		fs:       fs,
		renderer: renderer,
	}
}

// WriteSkill writes the hub skill's SKILL.md and reference.md to the output
// directory. It first removes any stale sibling skill directories left over
// from a previous generation (e.g. the old one-skill-per-doc layout), and
// then wipes the hub's own directory before recreating it — a category's
// root doc can share its name with the new hub (e.g. "patterns"), in which
// case the old per-doc skill's leftover examples.md/scripts/ would
// otherwise survive as a same-named "sibling of itself".
func (w *SkillWriter) WriteSkill(skill *domain.Skill, outputDir string) error {
	skillsDir := filepath.Join(outputDir, skill.Metadata.Category, "skills")
	skillDir := filepath.Join(skillsDir, skill.Metadata.Name)

	if err := w.removeStaleSiblings(skillsDir, skill.Metadata.Name); err != nil {
		return fmt.Errorf("failed to remove stale skill directories in %s: %w", skillsDir, err)
	}

	if err := w.fs.RemoveAll(skillDir); err != nil {
		return fmt.Errorf("failed to clear skill directory %s: %w", skillDir, err)
	}

	// Create skill directory
	if err := w.fs.MkdirAll(skillDir, 0755); err != nil {
		return fmt.Errorf("failed to create skill directory %s: %w", skillDir, err)
	}

	// Write SKILL.md: the lean, scannable index
	skillContent, err := w.renderer.RenderSkill(skill)
	if err != nil {
		return fmt.Errorf("failed to render SKILL.md for %s: %w", skill.Metadata.Name, err)
	}

	skillPath := filepath.Join(skillDir, "SKILL.md")
	if err := w.fs.WriteFile(skillPath, []byte(skillContent), 0644); err != nil {
		return fmt.Errorf("failed to write SKILL.md: %w", err)
	}

	// Write reference.md: the full offline depth behind it
	referenceContent, err := w.renderer.RenderReference(skill)
	if err != nil {
		return fmt.Errorf("failed to render reference.md for %s: %w", skill.Metadata.Name, err)
	}

	referencePath := filepath.Join(skillDir, "reference.md")
	if err := w.fs.WriteFile(referencePath, []byte(referenceContent), 0644); err != nil {
		return fmt.Errorf("failed to write reference.md: %w", err)
	}

	// Write library/: every source doc verbatim, mirroring the docs tree.
	// WriteFile creates parent directories as needed, so no separate
	// MkdirAll per file is required.
	libraryDir := filepath.Join(skillDir, "library")
	for _, lf := range skill.LibraryFiles {
		libraryPath := filepath.Join(libraryDir, lf.RelPath)
		if err := w.fs.WriteFile(libraryPath, []byte(lf.Content), 0644); err != nil {
			return fmt.Errorf("failed to write library file %s: %w", lf.RelPath, err)
		}
	}

	return nil
}

// removeStaleSiblings deletes every entry under skillsDir other than keep,
// so regenerating a category's hub also cleans up skill directories that
// are no longer produced.
func (w *SkillWriter) removeStaleSiblings(skillsDir, keep string) error {
	entries, err := w.fs.Glob(filepath.Join(skillsDir, "*"))
	if err != nil {
		if !w.fs.Exists(skillsDir) {
			return nil
		}
		return err
	}

	for _, entry := range entries {
		if filepath.Base(entry) == keep {
			continue
		}
		if err := w.fs.RemoveAll(entry); err != nil {
			return err
		}
	}

	return nil
}

// MarketplaceWriter implements ports.MarketplaceWriter using the filesystem.
type MarketplaceWriter struct {
	fs ports.FileSystem
}

// NewMarketplaceWriter creates a new filesystem-based marketplace writer.
func NewMarketplaceWriter(fs ports.FileSystem) *MarketplaceWriter {
	return &MarketplaceWriter{fs: fs}
}

// Read reads the current marketplace.json file.
func (w *MarketplaceWriter) Read(path string) (*domain.Marketplace, error) {
	content, err := w.fs.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read marketplace.json: %w", err)
	}

	var marketplace domain.Marketplace
	if err := json.Unmarshal(content, &marketplace); err != nil {
		return nil, fmt.Errorf("failed to parse marketplace.json: %w", err)
	}

	return &marketplace, nil
}

// Write writes the marketplace.json file atomically.
func (w *MarketplaceWriter) Write(marketplace *domain.Marketplace, path string) error {
	// Pretty-print JSON with 2-space indentation
	content, err := json.MarshalIndent(marketplace, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal marketplace.json: %w", err)
	}

	// Add trailing newline
	content = append(content, '\n')

	if err := w.fs.WriteFile(path, content, 0644); err != nil {
		return fmt.Errorf("failed to write marketplace.json: %w", err)
	}

	return nil
}

// GenerateFromConfig builds marketplace.json from config + versions.
func (w *MarketplaceWriter) GenerateFromConfig(
	metadata *domain.PluginMetadata,
	versions map[string]string,
	outputPath string,
) error {
	// Build marketplace structure
	marketplace := &domain.Marketplace{
		Name:  metadata.Marketplace.Name,
		Owner: metadata.Marketplace.Owner,
		Metadata: domain.MarketplaceMetadata{
			Description: metadata.Marketplace.Description,
			Version:     extractMarketplaceVersion(versions),
			PluginRoot:  metadata.Marketplace.PluginRoot,
		},
		Plugins: []domain.Plugin{},
	}

	// Extract and sort plugin keys for deterministic ordering
	pluginKeys := make([]string, 0, len(metadata.Plugins))
	for key := range metadata.Plugins {
		pluginKeys = append(pluginKeys, key)
	}
	sort.Strings(pluginKeys)

	// Build plugin entries in sorted order
	for _, pluginKey := range pluginKeys {
		pluginConfig := metadata.Plugins[pluginKey]

		// Extract version from manifest
		manifestKey := fmt.Sprintf("plugins/%s", pluginKey)
		version := versions[manifestKey]
		if version == "" {
			version = "0.0.0"
		}

		// Determine source path
		source := fmt.Sprintf("./plugins/%s", pluginKey)

		// Build plugin entry
		plugin := domain.Plugin{
			Name:        pluginConfig.GetMarketplaceName(pluginKey),
			Source:      source,
			Description: pluginConfig.Description,
			Version:     version,
			Category:    pluginConfig.Category,
			Tags:        pluginConfig.Tags,
		}

		// Add author if provided in common fields
		if metadata.Common.Author != nil {
			plugin.Author = metadata.Common.Author
		}

		marketplace.Plugins = append(marketplace.Plugins, plugin)
	}

	// Write marketplace.json
	return w.Write(marketplace, outputPath)
}

// WritePluginManifest writes an individual plugin.json file.
func (w *MarketplaceWriter) WritePluginManifest(
	manifest *domain.PluginManifest,
	outputPath string,
) error {
	// Pretty-print JSON with 2-space indentation
	content, err := json.MarshalIndent(manifest, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal plugin manifest: %w", err)
	}

	// Add trailing newline
	content = append(content, '\n')

	if err := w.fs.WriteFile(outputPath, content, 0644); err != nil {
		return fmt.Errorf("failed to write plugin manifest: %w", err)
	}

	return nil
}

// extractMarketplaceVersion extracts the marketplace version from the release manifest.
// Looks for ".claude-plugin" key in the versions map.
// Returns "0.0.0" if the key is not found (following the same convention as extractVersionForPlugin).
func extractMarketplaceVersion(versions map[string]string) string {
	if version, ok := versions[".claude-plugin"]; ok && version != "" {
		return version
	}
	return "0.0.0"
}

// DeriveSkillName converts a title to kebab-case.
func DeriveSkillName(title string) string {
	// Convert to lowercase
	name := strings.ToLower(title)

	// Replace spaces and special characters with hyphens
	name = strings.Map(func(r rune) rune {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			return r
		}
		if r == ' ' || r == '-' || r == '_' {
			return '-'
		}
		return -1 // Remove character
	}, name)

	// Remove consecutive hyphens
	for strings.Contains(name, "--") {
		name = strings.ReplaceAll(name, "--", "-")
	}

	// Trim hyphens from start and end
	name = strings.Trim(name, "-")

	return name
}
