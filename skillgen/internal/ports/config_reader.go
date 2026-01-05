package ports

import "github.com/adaptive-enforcement-lab/claude-skills/skillgen/internal/domain"

// ConfigReader reads plugin metadata and release manifest configuration.
type ConfigReader interface {
	// ReadPluginMetadata reads and parses plugin-metadata.json.
	// This contains static plugin metadata (descriptions, categories, tags).
	ReadPluginMetadata(path string) (*domain.PluginMetadata, error)

	// ReadReleaseManifest reads and parses .release-please-manifest.json.
	// Returns a map of path -> version (e.g., "skills/patterns" -> "0.2.1").
	ReadReleaseManifest(path string) (map[string]string, error)
}
