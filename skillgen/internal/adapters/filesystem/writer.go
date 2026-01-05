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

// WriteSkill writes all components of a skill to the output directory.
func (w *SkillWriter) WriteSkill(skill *domain.Skill, outputDir string) error {
	// Determine output path: outputDir/category/skills/skill-name/
	skillDir := filepath.Join(outputDir, skill.Metadata.Category, "skills", skill.Metadata.Name)

	// Create skill directory
	if err := w.fs.MkdirAll(skillDir, 0755); err != nil {
		return fmt.Errorf("failed to create skill directory %s: %w", skillDir, err)
	}

	// Write SKILL.md (always required)
	skillContent, err := w.renderer.RenderSkill(skill)
	if err != nil {
		return fmt.Errorf("failed to render SKILL.md for %s: %w", skill.Metadata.Name, err)
	}

	skillPath := filepath.Join(skillDir, "SKILL.md")
	if err := w.fs.WriteFile(skillPath, []byte(skillContent), 0644); err != nil {
		return fmt.Errorf("failed to write SKILL.md: %w", err)
	}

	// Write examples.md if needed
	if skill.Examples != nil && skill.Examples.ShouldGenerate() {
		examplesContent, err := w.renderer.RenderExamples(skill)
		if err != nil {
			return fmt.Errorf("failed to render examples.md: %w", err)
		}

		examplesPath := filepath.Join(skillDir, "examples.md")
		if err := w.fs.WriteFile(examplesPath, []byte(examplesContent), 0644); err != nil {
			return fmt.Errorf("failed to write examples.md: %w", err)
		}
	}

	// Write troubleshooting.md if needed
	if skill.Troubleshooting != nil && skill.Troubleshooting.ShouldGenerate() {
		troubleshootingContent, err := w.renderer.RenderTroubleshooting(skill)
		if err != nil {
			return fmt.Errorf("failed to render troubleshooting.md: %w", err)
		}

		troubleshootingPath := filepath.Join(skillDir, "troubleshooting.md")
		if err := w.fs.WriteFile(troubleshootingPath, []byte(troubleshootingContent), 0644); err != nil {
			return fmt.Errorf("failed to write troubleshooting.md: %w", err)
		}
	}

	// Write reference.md if needed
	if skill.Reference != nil && skill.Reference.ShouldGenerate() {
		referenceContent, err := w.renderer.RenderReference(skill)
		if err != nil {
			return fmt.Errorf("failed to render reference.md: %w", err)
		}

		referencePath := filepath.Join(skillDir, "reference.md")
		if err := w.fs.WriteFile(referencePath, []byte(referenceContent), 0644); err != nil {
			return fmt.Errorf("failed to write reference.md: %w", err)
		}
	}

	// Write script files if any
	if len(skill.Scripts) > 0 {
		scriptsDir := filepath.Join(skillDir, "scripts")
		if err := w.fs.MkdirAll(scriptsDir, 0755); err != nil {
			return fmt.Errorf("failed to create scripts directory: %w", err)
		}

		for _, script := range skill.Scripts {
			scriptPath := filepath.Join(scriptsDir, script.Filename)
			if err := w.fs.WriteFile(scriptPath, []byte(script.Content), 0755); err != nil {
				return fmt.Errorf("failed to write script %s: %w", script.Filename, err)
			}
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

// PreservePrivateCollection ensures private-collection is not removed.
func (w *MarketplaceWriter) PreservePrivateCollection(marketplace *domain.Marketplace) error {
	// Check if private-collection exists
	if marketplace.GetPlugin("private-collection") == nil {
		// It's okay if it doesn't exist - nothing to preserve
		return nil
	}

	// The plugin exists, so it's already preserved in the marketplace structure
	// No action needed
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
