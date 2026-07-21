package ports

import "github.com/adaptive-enforcement-lab/claude-skills/skillgen/internal/domain"

// SkillWriter writes skill files to the filesystem.
type SkillWriter interface {
	// WriteSkill writes all components of a skill to the output directory.
	// Creates the directory structure: skills/{category}/{skill-name}/
	WriteSkill(skill *domain.Skill, outputDir string) error
}

// MarketplaceWriter manages the marketplace.json file.
type MarketplaceWriter interface {
	// Read reads the current marketplace.json file.
	Read(path string) (*domain.Marketplace, error)

	// Write writes the marketplace.json file atomically.
	Write(marketplace *domain.Marketplace, path string) error

	// GenerateFromConfig builds marketplace.json from config + versions.
	// Every plugin in marketplace.json is rebuilt from plugin-metadata.json on
	// each run, so a plugin absent from that file is dropped.
	GenerateFromConfig(
		metadata *domain.PluginMetadata,
		versions map[string]string,
		outputPath string,
	) error

	// WritePluginManifest writes an individual plugin.json file.
	WritePluginManifest(
		manifest *domain.PluginManifest,
		outputPath string,
	) error
}
