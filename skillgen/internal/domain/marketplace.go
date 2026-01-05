package domain

// Marketplace represents the .claude-plugin/marketplace.json catalog structure.
// This file indexes all skill collections (plugins) in the repository.
type Marketplace struct {
	Name     string            `json:"name"`
	Owner    MarketplaceOwner  `json:"owner"`
	Metadata MarketplaceMetadata `json:"metadata"`
	Plugins  []Plugin          `json:"plugins"`
}

// MarketplaceOwner contains contact information for the marketplace maintainer.
type MarketplaceOwner struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// MarketplaceMetadata contains marketplace-level configuration.
type MarketplaceMetadata struct {
	Description string `json:"description"`
	Version     string `json:"version"`
	PluginRoot  string `json:"pluginRoot"`
}

// Plugin represents a skill collection within the marketplace.
// Each plugin corresponds to a category: patterns, enforcement, build, secure.
type Plugin struct {
	Name        string        `json:"name"`
	Source      string        `json:"source"`
	Description string        `json:"description"`
	Version     string        `json:"version"`
	Category    string        `json:"category"`
	Author      *PluginAuthor `json:"author,omitempty"`
	Tags        []string      `json:"tags"`
}

// PluginAuthor contains optional author information for a plugin.
type PluginAuthor struct {
	Name  string `json:"name"`
	Email string `json:"email,omitempty"`
}

// AddPlugin adds a new plugin to the marketplace if it doesn't already exist.
// Returns true if the plugin was added, false if it already existed.
func (m *Marketplace) AddPlugin(plugin Plugin) bool {
	for _, p := range m.Plugins {
		if p.Name == plugin.Name {
			return false // Already exists
		}
	}
	m.Plugins = append(m.Plugins, plugin)
	return true
}

// GetPlugin retrieves a plugin by name, or nil if not found.
func (m *Marketplace) GetPlugin(name string) *Plugin {
	for i := range m.Plugins {
		if m.Plugins[i].Name == name {
			return &m.Plugins[i]
		}
	}
	return nil
}

// UpdatePlugin updates an existing plugin or adds it if it doesn't exist.
func (m *Marketplace) UpdatePlugin(plugin Plugin) {
	for i := range m.Plugins {
		if m.Plugins[i].Name == plugin.Name {
			m.Plugins[i] = plugin
			return
		}
	}
	m.Plugins = append(m.Plugins, plugin)
}

// CategoryToPluginName maps a document category to its plugin name.
// Some mappings are non-obvious: "enforce" -> "enforcement"
func CategoryToPluginName(category string) string {
	switch category {
	case "enforce":
		return "enforcement"
	case "patterns", "build", "secure":
		return category
	default:
		return category
	}
}
