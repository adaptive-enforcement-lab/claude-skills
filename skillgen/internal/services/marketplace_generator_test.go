package services

import (
	"fmt"
	"strings"
	"testing"

	"github.com/adaptive-enforcement-lab/claude-skills/skillgen/internal/domain"
	"github.com/adaptive-enforcement-lab/claude-skills/skillgen/internal/ports"
)

// Mock implementations for testing

type MockConfigReader struct {
	metadata          *domain.PluginMetadata
	manifest          map[string]string
	metadataError     error
	manifestError     error
	metadataCallCount int
	manifestCallCount int
}

func (m *MockConfigReader) ReadPluginMetadata(path string) (*domain.PluginMetadata, error) {
	m.metadataCallCount++
	if m.metadataError != nil {
		return nil, m.metadataError
	}
	return m.metadata, nil
}

func (m *MockConfigReader) ReadReleaseManifest(path string) (map[string]string, error) {
	m.manifestCallCount++
	if m.manifestError != nil {
		return nil, m.manifestError
	}
	return m.manifest, nil
}

type MockMarketplaceWriter struct {
	generatedMarketplace *domain.Marketplace
	generatedPlugins     map[string]*domain.PluginManifest
	generateError        error
	writePluginError     error
	generateCallCount    int
	writePluginCallCount int
}

func NewMockMarketplaceWriter() *MockMarketplaceWriter {
	return &MockMarketplaceWriter{
		generatedPlugins: make(map[string]*domain.PluginManifest),
	}
}

func (m *MockMarketplaceWriter) Read(path string) (*domain.Marketplace, error) {
	return nil, fmt.Errorf("not implemented in mock")
}

func (m *MockMarketplaceWriter) Write(marketplace *domain.Marketplace, path string) error {
	return fmt.Errorf("not implemented in mock")
}

func (m *MockMarketplaceWriter) PreservePrivateCollection(marketplace *domain.Marketplace) error {
	return fmt.Errorf("not implemented in mock")
}

func (m *MockMarketplaceWriter) GenerateFromConfig(
	metadata *domain.PluginMetadata,
	versions map[string]string,
	outputPath string,
) error {
	m.generateCallCount++
	if m.generateError != nil {
		return m.generateError
	}
	// Store for verification
	m.generatedMarketplace = &domain.Marketplace{
		Name: metadata.Marketplace.Name,
	}
	return nil
}

func (m *MockMarketplaceWriter) WritePluginManifest(
	manifest *domain.PluginManifest,
	outputPath string,
) error {
	m.writePluginCallCount++
	if m.writePluginError != nil {
		return m.writePluginError
	}
	// Store for verification
	m.generatedPlugins[manifest.Name] = manifest
	return nil
}

type MockLogger struct {
	logs []string
}

func (m *MockLogger) Debug(msg string, keysAndValues ...interface{}) {
	m.logs = append(m.logs, fmt.Sprintf("DEBUG: %s %v", msg, keysAndValues))
}

func (m *MockLogger) Info(msg string, keysAndValues ...interface{}) {
	m.logs = append(m.logs, fmt.Sprintf("INFO: %s %v", msg, keysAndValues))
}

func (m *MockLogger) Warn(msg string, keysAndValues ...interface{}) {
	m.logs = append(m.logs, fmt.Sprintf("WARN: %s %v", msg, keysAndValues))
}

func (m *MockLogger) Error(msg string, keysAndValues ...interface{}) {
	m.logs = append(m.logs, fmt.Sprintf("ERROR: %s %v", msg, keysAndValues))
}

func (m *MockLogger) With(keysAndValues ...interface{}) ports.Logger {
	// For testing purposes, just return the same logger
	return m
}

// Tests

func TestMarketplaceGenerator_Generate(t *testing.T) {
	tests := []struct {
		name         string
		metadata     *domain.PluginMetadata
		manifest     map[string]string
		metadataErr  error
		manifestErr  error
		generateErr  error
		writePluginErr error
		wantErr      bool
		errContains  string
		validate     func(*testing.T, *MockMarketplaceWriter, *MockLogger)
	}{
		{
			name: "successful generation with all plugins",
			metadata: &domain.PluginMetadata{
				Marketplace: domain.MarketplaceConfig{
					Name: "ael-skills",
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
						Keywords:    []string{"patterns"},
					},
					"enforce": {
						Description: "Enforcement skills",
						Category:    "security",
						Keywords:    []string{"security"},
					},
				},
			},
			manifest: map[string]string{
				".claude-plugin":  "0.2.4",
				"plugins/patterns": "0.2.1",
				"plugins/enforce":  "0.3.0",
			},
			wantErr: false,
			validate: func(t *testing.T, writer *MockMarketplaceWriter, logger *MockLogger) {
				// Verify GenerateFromConfig was called
				if writer.generateCallCount != 1 {
					t.Errorf("expected GenerateFromConfig to be called once, got %d", writer.generateCallCount)
				}

				// Verify WritePluginManifest was called for each plugin
				if writer.writePluginCallCount != 2 {
					t.Errorf("expected WritePluginManifest to be called twice, got %d", writer.writePluginCallCount)
				}

				// Verify plugins were generated
				if len(writer.generatedPlugins) != 2 {
					t.Errorf("expected 2 plugins to be generated, got %d", len(writer.generatedPlugins))
				}

				// Verify patterns plugin
				if patterns, ok := writer.generatedPlugins["patterns"]; !ok {
					t.Error("expected 'patterns' plugin to be generated")
				} else {
					if patterns.Version != "0.2.1" {
						t.Errorf("expected patterns version '0.2.1', got %q", patterns.Version)
					}
					if patterns.Description != "Pattern skills" {
						t.Errorf("expected patterns description 'Pattern skills', got %q", patterns.Description)
					}
				}

				// Verify enforce plugin
				if enforce, ok := writer.generatedPlugins["enforce"]; !ok {
					t.Error("expected 'enforce' plugin to be generated")
				} else {
					if enforce.Version != "0.3.0" {
						t.Errorf("expected enforce version '0.3.0', got %q", enforce.Version)
					}
				}

				// Verify logging
				hasCompletionLog := false
				for _, log := range logger.logs {
					if strings.Contains(log, "marketplace generation complete") {
						hasCompletionLog = true
						break
					}
				}
				if !hasCompletionLog {
					t.Error("expected completion log message")
				}
			},
		},
		{
			name: "missing plugin version uses default",
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
			manifest: map[string]string{
				".claude-plugin": "1.0.0",
				// Missing "plugins/test-plugin"
			},
			wantErr: false,
			validate: func(t *testing.T, writer *MockMarketplaceWriter, logger *MockLogger) {
				// Verify plugin was generated with default version
				if plugin, ok := writer.generatedPlugins["test-plugin"]; !ok {
					t.Error("expected 'test-plugin' to be generated")
				} else if plugin.Version != "0.0.0" {
					t.Errorf("expected default version '0.0.0', got %q", plugin.Version)
				}

				// Verify warning was logged
				hasWarning := false
				for _, log := range logger.logs {
					if strings.Contains(log, "WARN") && strings.Contains(log, "no version found") {
						hasWarning = true
						break
					}
				}
				if !hasWarning {
					t.Error("expected warning log for missing version")
				}
			},
		},
		{
			name:        "metadata read error",
			metadataErr: fmt.Errorf("failed to read metadata"),
			wantErr:     true,
			errContains: "failed to read plugin metadata",
		},
		{
			name: "manifest read error",
			metadata: &domain.PluginMetadata{
				Marketplace: domain.MarketplaceConfig{
					Name: "test",
					Owner: domain.MarketplaceOwner{
						Name:  "Test",
						Email: "test@example.com",
					},
				},
				Plugins: map[string]domain.PluginConfig{
					"test": {
						Description: "Test",
						Category:    "test",
					},
				},
			},
			manifestErr: fmt.Errorf("failed to read manifest"),
			wantErr:     true,
			errContains: "failed to read release manifest",
		},
		{
			name: "marketplace generation error",
			metadata: &domain.PluginMetadata{
				Marketplace: domain.MarketplaceConfig{
					Name: "test",
					Owner: domain.MarketplaceOwner{
						Name:  "Test",
						Email: "test@example.com",
					},
				},
				Plugins: map[string]domain.PluginConfig{
					"test": {
						Description: "Test",
						Category:    "test",
					},
				},
			},
			manifest: map[string]string{
				".claude-plugin": "1.0.0",
			},
			generateErr: fmt.Errorf("failed to write marketplace"),
			wantErr:     true,
			errContains: "failed to generate marketplace.json",
		},
		{
			name: "plugin manifest write error",
			metadata: &domain.PluginMetadata{
				Marketplace: domain.MarketplaceConfig{
					Name: "test",
					Owner: domain.MarketplaceOwner{
						Name:  "Test",
						Email: "test@example.com",
					},
				},
				Plugins: map[string]domain.PluginConfig{
					"test": {
						Description: "Test",
						Category:    "test",
					},
				},
			},
			manifest: map[string]string{
				".claude-plugin": "1.0.0",
			},
			writePluginErr: fmt.Errorf("failed to write plugin manifest"),
			wantErr:        true,
			errContains:    "failed to write plugin manifest",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockReader := &MockConfigReader{
				metadata:      tt.metadata,
				manifest:      tt.manifest,
				metadataError: tt.metadataErr,
				manifestError: tt.manifestErr,
			}

			mockWriter := NewMockMarketplaceWriter()
			mockWriter.generateError = tt.generateErr
			mockWriter.writePluginError = tt.writePluginErr

			mockLogger := &MockLogger{
				logs: []string{},
			}

			generator := NewMarketplaceGenerator(mockReader, mockWriter, mockLogger)

			err := generator.Generate(
				"plugin-metadata.json",
				".release-please-manifest.json",
				"./skills",
			)

			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error containing %q, got nil", tt.errContains)
					return
				}
				if !strings.Contains(err.Error(), tt.errContains) {
					t.Errorf("expected error containing %q, got %q", tt.errContains, err.Error())
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if tt.validate != nil {
				tt.validate(t, mockWriter, mockLogger)
			}
		})
	}
}

func Test_extractVersionForPlugin(t *testing.T) {
	tests := []struct {
		name      string
		versions  map[string]string
		pluginKey string
		want      string
	}{
		{
			name: "version exists",
			versions: map[string]string{
				"plugins/patterns": "0.2.1",
				"plugins/enforce":  "0.3.0",
			},
			pluginKey: "patterns",
			want:      "0.2.1",
		},
		{
			name: "version does not exist",
			versions: map[string]string{
				"plugins/patterns": "0.2.1",
			},
			pluginKey: "enforce",
			want:      "",
		},
		{
			name:      "empty versions map",
			versions:  map[string]string{},
			pluginKey: "patterns",
			want:      "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := extractVersionForPlugin(tt.versions, tt.pluginKey)
			if got != tt.want {
				t.Errorf("extractVersionForPlugin() = %q, want %q", got, tt.want)
			}
		})
	}
}

func Test_buildPluginManifest(t *testing.T) {
	tests := []struct {
		name    string
		key     string
		config  *domain.PluginConfig
		common  *domain.CommonPluginFields
		version string
		validate func(*testing.T, *domain.PluginManifest)
	}{
		{
			name: "full manifest with all fields",
			key:  "patterns",
			config: &domain.PluginConfig{
				Description: "Pattern skills",
				Category:    "development",
				Keywords:    []string{"patterns", "best-practices"},
			},
			common: &domain.CommonPluginFields{
				Author: &domain.PluginAuthor{
					Name:  "Test Author",
					Email: "author@example.com",
				},
				Homepage:   "https://example.com",
				Repository: "https://github.com/test/repo",
				License:    "MIT",
			},
			version: "0.2.1",
			validate: func(t *testing.T, manifest *domain.PluginManifest) {
				if manifest.Name != "patterns" {
					t.Errorf("expected name 'patterns', got %q", manifest.Name)
				}
				if manifest.Version != "0.2.1" {
					t.Errorf("expected version '0.2.1', got %q", manifest.Version)
				}
				if manifest.Description != "Pattern skills" {
					t.Errorf("expected description 'Pattern skills', got %q", manifest.Description)
				}
				if manifest.Author == nil {
					t.Error("expected author to be set")
				} else if manifest.Author.Name != "Test Author" {
					t.Errorf("expected author name 'Test Author', got %q", manifest.Author.Name)
				}
				if manifest.Homepage != "https://example.com" {
					t.Errorf("expected homepage 'https://example.com', got %q", manifest.Homepage)
				}
				if len(manifest.Keywords) != 2 {
					t.Errorf("expected 2 keywords, got %d", len(manifest.Keywords))
				}
			},
		},
		{
			name: "minimal manifest",
			key:  "test",
			config: &domain.PluginConfig{
				Description: "Test",
				Category:    "test",
			},
			common:  &domain.CommonPluginFields{},
			version: "1.0.0",
			validate: func(t *testing.T, manifest *domain.PluginManifest) {
				if manifest.Name != "test" {
					t.Errorf("expected name 'test', got %q", manifest.Name)
				}
				if manifest.Author != nil {
					t.Error("expected author to be nil")
				}
				if manifest.Homepage != "" {
					t.Errorf("expected empty homepage, got %q", manifest.Homepage)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			manifest := buildPluginManifest(tt.key, tt.config, tt.common, tt.version)
			if tt.validate != nil {
				tt.validate(t, manifest)
			}
		})
	}
}
