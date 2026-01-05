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

	// AddSecurePlugin ensures the secure plugin exists in the marketplace.
	// Returns true if the plugin was added, false if it already existed.
	AddSecurePlugin(path string) (bool, error)

	// PreservePrivateCollection ensures private-collection is not removed.
	// This plugin is manually maintained and should never be auto-generated.
	PreservePrivateCollection(marketplace *domain.Marketplace) error
}
