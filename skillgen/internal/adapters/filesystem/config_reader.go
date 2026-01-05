package filesystem

import (
	"encoding/json"
	"fmt"

	"github.com/adaptive-enforcement-lab/claude-skills/skillgen/internal/domain"
	"github.com/adaptive-enforcement-lab/claude-skills/skillgen/internal/ports"
)

// ConfigReader implements ports.ConfigReader using the filesystem.
type ConfigReader struct {
	fs ports.FileSystem
}

// NewConfigReader creates a new filesystem-based config reader.
func NewConfigReader(fs ports.FileSystem) *ConfigReader {
	return &ConfigReader{fs: fs}
}

// ReadPluginMetadata reads and parses plugin-metadata.json.
func (r *ConfigReader) ReadPluginMetadata(path string) (*domain.PluginMetadata, error) {
	content, err := r.fs.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read plugin-metadata.json: %w", err)
	}

	var metadata domain.PluginMetadata
	if err := json.Unmarshal(content, &metadata); err != nil {
		return nil, fmt.Errorf("failed to parse plugin-metadata.json: %w", err)
	}

	// Validate required fields
	if metadata.Marketplace.Name == "" {
		return nil, fmt.Errorf("marketplace.name is required in plugin-metadata.json")
	}
	if metadata.Marketplace.Owner.Name == "" {
		return nil, fmt.Errorf("marketplace.owner.name is required in plugin-metadata.json")
	}
	if len(metadata.Plugins) == 0 {
		return nil, fmt.Errorf("plugins map cannot be empty in plugin-metadata.json")
	}

	return &metadata, nil
}

// ReadReleaseManifest reads and parses .release-please-manifest.json.
func (r *ConfigReader) ReadReleaseManifest(path string) (map[string]string, error) {
	content, err := r.fs.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read release manifest: %w", err)
	}

	var manifest map[string]string
	if err := json.Unmarshal(content, &manifest); err != nil {
		return nil, fmt.Errorf("failed to parse release manifest: %w", err)
	}

	if len(manifest) == 0 {
		return nil, fmt.Errorf("release manifest is empty")
	}

	return manifest, nil
}
