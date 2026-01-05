package domain

// PluginMetadata represents the plugin-metadata.json configuration file.
// This is the source of truth for plugin descriptions, categories, and tags.
type PluginMetadata struct {
	Marketplace MarketplaceConfig      `json:"marketplace"`
	Common      CommonPluginFields     `json:"common"`
	Plugins     map[string]PluginConfig `json:"plugins"`
}

// MarketplaceConfig defines marketplace-level settings.
type MarketplaceConfig struct {
	Name        string           `json:"name"`
	Owner       MarketplaceOwner `json:"owner"`
	Description string           `json:"description"`
	PluginRoot  string           `json:"pluginRoot"`
}

// CommonPluginFields contains fields applied to all plugin.json files.
// This follows DRY principle - define once, apply everywhere.
type CommonPluginFields struct {
	Author     *PluginAuthor `json:"author,omitempty"`
	Homepage   string        `json:"homepage,omitempty"`
	Repository string        `json:"repository,omitempty"`
	License    string        `json:"license,omitempty"`
}

// PluginConfig defines per-collection configuration.
type PluginConfig struct {
	// MarketplaceName overrides the plugin name in marketplace.json.
	// If empty, uses the plugin key from the map.
	MarketplaceName string `json:"marketplaceName,omitempty"`

	Description string   `json:"description"`
	Category    string   `json:"category"`
	Tags        []string `json:"tags"`
	Keywords    []string `json:"keywords"`
}

// PluginManifest represents an individual plugin.json file.
// This is what gets written to skills/{collection}/.claude-plugin/plugin.json
type PluginManifest struct {
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Version     string        `json:"version"`
	Author      *PluginAuthor `json:"author,omitempty"`
	Homepage    string        `json:"homepage,omitempty"`
	Repository  string        `json:"repository,omitempty"`
	License     string        `json:"license,omitempty"`
	Keywords    []string      `json:"keywords,omitempty"`
}

// GetMarketplaceName returns the effective marketplace name for a plugin.
// Uses MarketplaceName if set, otherwise falls back to the plugin key.
func (pc *PluginConfig) GetMarketplaceName(pluginKey string) string {
	if pc.MarketplaceName != "" {
		return pc.MarketplaceName
	}
	return pluginKey
}
