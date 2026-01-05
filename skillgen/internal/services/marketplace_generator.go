package services

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/adaptive-enforcement-lab/claude-skills/skillgen/internal/domain"
	"github.com/adaptive-enforcement-lab/claude-skills/skillgen/internal/ports"
)

// MarketplaceGenerator orchestrates marketplace.json and plugin.json generation.
// It reads configuration from plugin-metadata.json and versions from
// .release-please-manifest.json, then generates all marketplace files.
type MarketplaceGenerator struct {
	configReader ports.ConfigReader
	writer       ports.MarketplaceWriter
	logger       ports.Logger
}

// NewMarketplaceGenerator creates a new marketplace generator service.
func NewMarketplaceGenerator(
	configReader ports.ConfigReader,
	writer ports.MarketplaceWriter,
	logger ports.Logger,
) *MarketplaceGenerator {
	return &MarketplaceGenerator{
		configReader: configReader,
		writer:       writer,
		logger:       logger,
	}
}

// Generate creates marketplace.json and all plugin.json files.
func (g *MarketplaceGenerator) Generate(
	metadataPath string,
	manifestPath string,
	outputDir string,
) error {
	g.logger.Info("reading plugin metadata", "path", metadataPath)

	// 1. Read plugin-metadata.json
	metadata, err := g.configReader.ReadPluginMetadata(metadataPath)
	if err != nil {
		return fmt.Errorf("failed to read plugin metadata: %w", err)
	}

	g.logger.Info("reading release manifest", "path", manifestPath)

	// 2. Read .release-please-manifest.json
	versions, err := g.configReader.ReadReleaseManifest(manifestPath)
	if err != nil {
		return fmt.Errorf("failed to read release manifest: %w", err)
	}

	g.logger.Info("generating marketplace.json")

	// 3. Generate marketplace.json
	marketplacePath := filepath.Join(".claude-plugin", "marketplace.json")
	if err := g.writer.GenerateFromConfig(metadata, versions, marketplacePath); err != nil {
		return fmt.Errorf("failed to generate marketplace.json: %w", err)
	}

	g.logger.Info("marketplace.json generated successfully")

	// 4. Generate each plugin.json in deterministic order
	pluginKeys := make([]string, 0, len(metadata.Plugins))
	for key := range metadata.Plugins {
		pluginKeys = append(pluginKeys, key)
	}
	sort.Strings(pluginKeys)

	pluginCount := 0
	for _, pluginKey := range pluginKeys {
		pluginConfig := metadata.Plugins[pluginKey]

		version := extractVersionForPlugin(versions, pluginKey)
		if version == "" {
			g.logger.Warn("no version found for plugin, using 0.0.0", "plugin", pluginKey)
			version = "0.0.0"
		}

		manifest := buildPluginManifest(pluginKey, &pluginConfig, &metadata.Common, version)
		pluginPath := filepath.Join(outputDir, pluginKey, ".claude-plugin", "plugin.json")

		// Ensure directory exists
		pluginDir := filepath.Dir(pluginPath)
		g.logger.Debug("creating plugin directory", "path", pluginDir)

		if err := os.MkdirAll(pluginDir, 0755); err != nil {
			return fmt.Errorf("failed to create plugin directory %s: %w", pluginDir, err)
		}

		if err := g.writer.WritePluginManifest(manifest, pluginPath); err != nil {
			return fmt.Errorf("failed to write plugin manifest for %s: %w", pluginKey, err)
		}

		g.logger.Info("generated plugin manifest", "plugin", pluginKey, "version", version)
		pluginCount++
	}

	g.logger.Info("marketplace generation complete",
		"marketplace", marketplacePath,
		"plugins", pluginCount,
	)

	return nil
}

// extractVersionForPlugin extracts the version for a specific plugin from the manifest.
// Looks for "plugins/{pluginKey}" in the manifest map.
func extractVersionForPlugin(versions map[string]string, pluginKey string) string {
	manifestKey := fmt.Sprintf("plugins/%s", pluginKey)
	return versions[manifestKey]
}

// buildPluginManifest constructs a PluginManifest from config and common fields.
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
