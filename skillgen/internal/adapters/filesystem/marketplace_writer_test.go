package filesystem

import (
	"encoding/json"
	"testing"

	"github.com/adaptive-enforcement-lab/claude-skills/skillgen/internal/domain"
)

func TestMarketplaceWriter_Read(t *testing.T) {
	tests := []struct {
		name        string
		setupFile   string
		content     string
		path        string
		wantErr     bool
		errContains string
		validate    func(*testing.T, *domain.Marketplace)
	}{
		{
			name: "valid marketplace file",
			setupFile: "marketplace.json",
			content: `{
  "name": "test-marketplace",
  "owner": {
    "name": "Test Owner",
    "email": "test@example.com"
  },
  "metadata": {
    "description": "Test marketplace",
    "version": "1.0.0",
    "pluginRoot": "./skills"
  },
  "plugins": [
    {
      "name": "patterns",
      "source": "./skills/patterns",
      "description": "Pattern skills",
      "version": "0.2.1",
      "category": "development",
      "tags": ["patterns"]
    }
  ]
}`,
			path:    "marketplace.json",
			wantErr: false,
			validate: func(t *testing.T, m *domain.Marketplace) {
				if m.Name != "test-marketplace" {
					t.Errorf("expected name 'test-marketplace', got %q", m.Name)
				}
				if m.Metadata.Version != "1.0.0" {
					t.Errorf("expected version '1.0.0', got %q", m.Metadata.Version)
				}
				if len(m.Plugins) != 1 {
					t.Errorf("expected 1 plugin, got %d", len(m.Plugins))
				}
			},
		},
		{
			name: "malformed JSON",
			setupFile: "malformed.json",
			content: `{"name": invalid}`,
			path:    "malformed.json",
			wantErr: true,
			errContains: "failed to parse marketplace.json",
		},
		{
			name:        "file not found",
			path:        "nonexistent.json",
			wantErr:     true,
			errContains: "failed to read marketplace.json",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockFS := NewMockFileSystem()

			if tt.setupFile != "" && tt.content != "" {
				mockFS.AddFile(tt.setupFile, []byte(tt.content))
			}

			writer := NewMarketplaceWriter(mockFS)
			marketplace, err := writer.Read(tt.path)

			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error containing %q, got nil", tt.errContains)
					return
				}
				if tt.errContains != "" && !contains(err.Error(), tt.errContains) {
					t.Errorf("expected error containing %q, got %q", tt.errContains, err.Error())
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if tt.validate != nil {
				tt.validate(t, marketplace)
			}
		})
	}
}

func TestMarketplaceWriter_Write(t *testing.T) {
	tests := []struct {
		name        string
		marketplace *domain.Marketplace
		path        string
		wantErr     bool
		errContains string
		validate    func(*testing.T, []byte)
	}{
		{
			name: "successful write",
			marketplace: &domain.Marketplace{
				Name: "test-marketplace",
				Owner: domain.MarketplaceOwner{
					Name:  "Test Owner",
					Email: "test@example.com",
				},
				Metadata: domain.MarketplaceMetadata{
					Version:    "1.0.0",
					PluginRoot: "./skills",
				},
				Plugins: []domain.Plugin{
					{
						Name:        "patterns",
						Source:      "./skills/patterns",
						Description: "Pattern skills",
						Version:     "0.2.1",
						Category:    "development",
					},
				},
			},
			path:    "marketplace.json",
			wantErr: false,
			validate: func(t *testing.T, content []byte) {
				// Verify JSON is valid
				var marketplace domain.Marketplace
				if err := json.Unmarshal(content, &marketplace); err != nil {
					t.Errorf("written content is not valid JSON: %v", err)
					return
				}
				// Verify trailing newline
				if len(content) == 0 || content[len(content)-1] != '\n' {
					t.Error("expected trailing newline")
				}
				// Verify content
				if marketplace.Name != "test-marketplace" {
					t.Errorf("expected name 'test-marketplace', got %q", marketplace.Name)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockFS := NewMockFileSystem()
			writer := NewMarketplaceWriter(mockFS)

			err := writer.Write(tt.marketplace, tt.path)

			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error containing %q, got nil", tt.errContains)
					return
				}
				if tt.errContains != "" && !contains(err.Error(), tt.errContains) {
					t.Errorf("expected error containing %q, got %q", tt.errContains, err.Error())
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			// Verify file was written
			written := mockFS.GetFile(tt.path)
			if written == nil {
				t.Errorf("expected file %q to be written", tt.path)
				return
			}

			if tt.validate != nil {
				tt.validate(t, written)
			}
		})
	}
}

func TestMarketplaceWriter_PreservePrivateCollection(t *testing.T) {
	tests := []struct {
		name        string
		marketplace *domain.Marketplace
		wantErr     bool
	}{
		{
			name: "marketplace with private-collection",
			marketplace: &domain.Marketplace{
				Name: "test",
				Plugins: []domain.Plugin{
					{Name: "patterns", Source: "./skills/patterns"},
					{Name: "private-collection", Source: "./skills/private"},
				},
			},
			wantErr: false,
		},
		{
			name: "marketplace without private-collection",
			marketplace: &domain.Marketplace{
				Name: "test",
				Plugins: []domain.Plugin{
					{Name: "patterns", Source: "./skills/patterns"},
				},
			},
			wantErr: false,
		},
		{
			name: "empty marketplace",
			marketplace: &domain.Marketplace{
				Name:    "test",
				Plugins: []domain.Plugin{},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockFS := NewMockFileSystem()
			writer := NewMarketplaceWriter(mockFS)

			err := writer.PreservePrivateCollection(tt.marketplace)

			if tt.wantErr && err == nil {
				t.Error("expected error, got nil")
			} else if !tt.wantErr && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}

func TestMarketplaceWriter_GenerateFromConfig(t *testing.T) {
	tests := []struct {
		name        string
		metadata    *domain.PluginMetadata
		versions    map[string]string
		outputPath  string
		wantErr     bool
		errContains string
		validate    func(*testing.T, *MockFileSystem)
	}{
		{
			name: "full config with all plugins",
			metadata: &domain.PluginMetadata{
				Marketplace: domain.MarketplaceConfig{
					Name:        "ael-skills",
					Description: "Test marketplace",
					PluginRoot:  "./skills",
					Owner: domain.MarketplaceOwner{
						Name:  "AEL",
						Email: "contact@ael.com",
					},
				},
				Common: domain.CommonPluginFields{
					Author: &domain.PluginAuthor{
						Name:  "Test Author",
						Email: "author@example.com",
					},
				},
				Plugins: map[string]domain.PluginConfig{
					"patterns": {
						Description: "Pattern skills",
						Category:    "development",
						Tags:        []string{"patterns"},
					},
					"enforce": {
						MarketplaceName: "enforcement",
						Description:     "Enforcement skills",
						Category:        "security",
						Tags:            []string{"security"},
					},
				},
			},
			versions: map[string]string{
				".claude-plugin":   "0.2.4",
				"plugins/patterns": "0.2.1",
				"plugins/enforce":  "0.3.0",
			},
			outputPath: "marketplace.json",
			wantErr:    false,
			validate: func(t *testing.T, mockFS *MockFileSystem) {
				content := mockFS.GetFile("marketplace.json")
				if content == nil {
					t.Error("expected marketplace.json to be written")
					return
				}

				var marketplace domain.Marketplace
				if err := json.Unmarshal(content, &marketplace); err != nil {
					t.Errorf("failed to parse written marketplace.json: %v", err)
					return
				}

				if marketplace.Name != "ael-skills" {
					t.Errorf("expected name 'ael-skills', got %q", marketplace.Name)
				}
				if marketplace.Metadata.Version != "0.2.4" {
					t.Errorf("expected version '0.2.4', got %q", marketplace.Metadata.Version)
				}
				if len(marketplace.Plugins) != 2 {
					t.Errorf("expected 2 plugins, got %d", len(marketplace.Plugins))
				}

				// Find enforcement plugin
				var enforcementPlugin *domain.Plugin
				for i := range marketplace.Plugins {
					if marketplace.Plugins[i].Name == "enforcement" {
						enforcementPlugin = &marketplace.Plugins[i]
						break
					}
				}
				if enforcementPlugin == nil {
					t.Error("expected 'enforcement' plugin to exist")
				} else if enforcementPlugin.Version != "0.3.0" {
					t.Errorf("expected enforcement version '0.3.0', got %q", enforcementPlugin.Version)
				}
			},
		},
		{
			name: "minimal config with missing versions",
			metadata: &domain.PluginMetadata{
				Marketplace: domain.MarketplaceConfig{
					Name: "test",
					Owner: domain.MarketplaceOwner{
						Name:  "Test",
						Email: "test@example.com",
					},
				},
				Plugins: map[string]domain.PluginConfig{
					"test-plugin": {
						Description: "Test",
						Category:    "test",
					},
				},
			},
			versions:   map[string]string{}, // Empty versions
			outputPath: "marketplace.json",
			wantErr:    false,
			validate: func(t *testing.T, mockFS *MockFileSystem) {
				content := mockFS.GetFile("marketplace.json")
				if content == nil {
					t.Error("expected marketplace.json to be written")
					return
				}

				var marketplace domain.Marketplace
				if err := json.Unmarshal(content, &marketplace); err != nil {
					t.Errorf("failed to parse written marketplace.json: %v", err)
					return
				}

				// Verify default version is used
				if marketplace.Metadata.Version != "0.0.0" {
					t.Errorf("expected default version '0.0.0', got %q", marketplace.Metadata.Version)
				}
				if len(marketplace.Plugins) != 1 {
					t.Errorf("expected 1 plugin, got %d", len(marketplace.Plugins))
				}
				if marketplace.Plugins[0].Version != "0.0.0" {
					t.Errorf("expected plugin version '0.0.0', got %q", marketplace.Plugins[0].Version)
				}
			},
		},
		{
			name: "plugins are ordered alphabetically for deterministic output",
			metadata: &domain.PluginMetadata{
				Marketplace: domain.MarketplaceConfig{
					Name: "test",
					Owner: domain.MarketplaceOwner{
						Name:  "Test",
						Email: "test@example.com",
					},
				},
				Plugins: map[string]domain.PluginConfig{
					"zebra": {
						Description: "Zebra plugin",
						Category:    "animals",
					},
					"alpha": {
						Description: "Alpha plugin",
						Category:    "letters",
					},
					"middle": {
						Description: "Middle plugin",
						Category:    "position",
					},
					"beta": {
						Description: "Beta plugin",
						Category:    "letters",
					},
				},
			},
			versions: map[string]string{
				"plugins/zebra":  "1.0.0",
				"plugins/alpha":  "2.0.0",
				"plugins/middle": "3.0.0",
				"plugins/beta":   "4.0.0",
			},
			outputPath: "marketplace.json",
			wantErr:    false,
			validate: func(t *testing.T, mockFS *MockFileSystem) {
				content := mockFS.GetFile("marketplace.json")
				if content == nil {
					t.Error("expected marketplace.json to be written")
					return
				}

				var marketplace domain.Marketplace
				if err := json.Unmarshal(content, &marketplace); err != nil {
					t.Errorf("failed to parse written marketplace.json: %v", err)
					return
				}

				// Verify plugins are in alphabetical order
				if len(marketplace.Plugins) != 4 {
					t.Errorf("expected 4 plugins, got %d", len(marketplace.Plugins))
					return
				}

				expectedOrder := []string{"alpha", "beta", "middle", "zebra"}
				for i, expectedName := range expectedOrder {
					if marketplace.Plugins[i].Name != expectedName {
						t.Errorf("plugin at index %d: expected name %q, got %q", i, expectedName, marketplace.Plugins[i].Name)
					}
				}

				// Verify correct versions are assigned
				expectedVersions := map[string]string{
					"alpha":  "2.0.0",
					"beta":   "4.0.0",
					"middle": "3.0.0",
					"zebra":  "1.0.0",
				}
				for _, plugin := range marketplace.Plugins {
					if expectedVersion, ok := expectedVersions[plugin.Name]; ok {
						if plugin.Version != expectedVersion {
							t.Errorf("plugin %q: expected version %q, got %q", plugin.Name, expectedVersion, plugin.Version)
						}
					}
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockFS := NewMockFileSystem()
			writer := NewMarketplaceWriter(mockFS)

			err := writer.GenerateFromConfig(tt.metadata, tt.versions, tt.outputPath)

			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error containing %q, got nil", tt.errContains)
					return
				}
				if tt.errContains != "" && !contains(err.Error(), tt.errContains) {
					t.Errorf("expected error containing %q, got %q", tt.errContains, err.Error())
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if tt.validate != nil {
				tt.validate(t, mockFS)
			}
		})
	}
}

func TestMarketplaceWriter_WritePluginManifest(t *testing.T) {
	tests := []struct {
		name        string
		manifest    *domain.PluginManifest
		outputPath  string
		wantErr     bool
		errContains string
		validate    func(*testing.T, []byte)
	}{
		{
			name: "successful write",
			manifest: &domain.PluginManifest{
				Name:        "patterns",
				Description: "Pattern skills",
				Version:     "0.2.1",
				Author: &domain.PluginAuthor{
					Name:  "Test Author",
					Email: "author@example.com",
				},
				Keywords: []string{"error-handling", "patterns"},
			},
			outputPath: "plugin.json",
			wantErr:    false,
			validate: func(t *testing.T, content []byte) {
				// Verify JSON is valid
				var manifest domain.PluginManifest
				if err := json.Unmarshal(content, &manifest); err != nil {
					t.Errorf("written content is not valid JSON: %v", err)
					return
				}
				// Verify trailing newline
				if len(content) == 0 || content[len(content)-1] != '\n' {
					t.Error("expected trailing newline")
				}
				// Verify content
				if manifest.Name != "patterns" {
					t.Errorf("expected name 'patterns', got %q", manifest.Name)
				}
				if len(manifest.Keywords) != 2 {
					t.Errorf("expected 2 keywords, got %d", len(manifest.Keywords))
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockFS := NewMockFileSystem()
			writer := NewMarketplaceWriter(mockFS)

			err := writer.WritePluginManifest(tt.manifest, tt.outputPath)

			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error containing %q, got nil", tt.errContains)
					return
				}
				if tt.errContains != "" && !contains(err.Error(), tt.errContains) {
					t.Errorf("expected error containing %q, got %q", tt.errContains, err.Error())
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			// Verify file was written
			written := mockFS.GetFile(tt.outputPath)
			if written == nil {
				t.Errorf("expected file %q to be written", tt.outputPath)
				return
			}

			if tt.validate != nil {
				tt.validate(t, written)
			}
		})
	}
}
