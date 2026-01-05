package filesystem

import (
	"testing"

	"github.com/adaptive-enforcement-lab/claude-skills/skillgen/internal/domain"
)

func TestConfigReader_ReadPluginMetadata(t *testing.T) {
	tests := []struct {
		name        string
		setupFiles  map[string]string
		path        string
		wantErr     bool
		errContains string
		validate    func(*testing.T, *domain.PluginMetadata)
	}{
		{
			name: "valid config with all fields",
			setupFiles: map[string]string{
				"plugin-metadata.json": "../../services/testdata/valid_metadata.json",
			},
			path:    "plugin-metadata.json",
			wantErr: false,
			validate: func(t *testing.T, meta *domain.PluginMetadata) {
				if meta.Marketplace.Name != "ael-skills" {
					t.Errorf("expected marketplace name 'ael-skills', got %q", meta.Marketplace.Name)
				}
				if meta.Marketplace.Owner.Name != "Adaptive Enforcement Lab" {
					t.Errorf("expected owner name 'Adaptive Enforcement Lab', got %q", meta.Marketplace.Owner.Name)
				}
				if len(meta.Plugins) != 2 {
					t.Errorf("expected 2 plugins, got %d", len(meta.Plugins))
				}
				if _, exists := meta.Plugins["patterns"]; !exists {
					t.Error("expected 'patterns' plugin to exist")
				}
				if _, exists := meta.Plugins["enforce"]; !exists {
					t.Error("expected 'enforce' plugin to exist")
				}
			},
		},
		{
			name: "minimal valid config",
			setupFiles: map[string]string{
				"minimal.json": "../../services/testdata/minimal_metadata.json",
			},
			path:    "minimal.json",
			wantErr: false,
			validate: func(t *testing.T, meta *domain.PluginMetadata) {
				if meta.Marketplace.Name != "test" {
					t.Errorf("expected marketplace name 'test', got %q", meta.Marketplace.Name)
				}
				if len(meta.Plugins) != 1 {
					t.Errorf("expected 1 plugin, got %d", len(meta.Plugins))
				}
				if _, exists := meta.Plugins["test-plugin"]; !exists {
					t.Error("expected 'test-plugin' to exist")
				}
			},
		},
		{
			name: "missing marketplace name",
			setupFiles: map[string]string{
				"invalid.json": "../../services/testdata/invalid_metadata_missing_name.json",
			},
			path:        "invalid.json",
			wantErr:     true,
			errContains: "marketplace.name is required in plugin-metadata.json",
		},
		{
			name: "empty plugins map",
			setupFiles: map[string]string{
				"invalid.json": "../../services/testdata/invalid_metadata_empty_plugins.json",
			},
			path:        "invalid.json",
			wantErr:     true,
			errContains: "plugins map cannot be empty in plugin-metadata.json",
		},
		{
			name: "malformed JSON",
			setupFiles: map[string]string{
				"malformed.json": "../../services/testdata/malformed.json",
			},
			path:        "malformed.json",
			wantErr:     true,
			errContains: "failed to parse plugin-metadata.json",
		},
		{
			name:        "file not found",
			setupFiles:  map[string]string{},
			path:        "nonexistent.json",
			wantErr:     true,
			errContains: "failed to read plugin-metadata.json",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock filesystem
			mockFS := NewMockFileSystem()

			// Load fixture files
			for path, fixturePath := range tt.setupFiles {
				content, err := readTestFixture(fixturePath)
				if err != nil {
					t.Fatalf("failed to read test fixture %s: %v", fixturePath, err)
				}
				mockFS.AddFile(path, content)
			}

			// Create reader with mock filesystem
			reader := NewConfigReader(mockFS)

			// Execute
			meta, err := reader.ReadPluginMetadata(tt.path)

			// Assert error expectations
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

			// Assert success expectations
			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if tt.validate != nil {
				tt.validate(t, meta)
			}
		})
	}
}

func TestConfigReader_ReadReleaseManifest(t *testing.T) {
	tests := []struct {
		name        string
		setupFiles  map[string]string
		path        string
		wantErr     bool
		errContains string
		validate    func(*testing.T, map[string]string)
	}{
		{
			name: "valid manifest with all versions",
			setupFiles: map[string]string{
				"manifest.json": "../../services/testdata/valid_manifest.json",
			},
			path:    "manifest.json",
			wantErr: false,
			validate: func(t *testing.T, versions map[string]string) {
				expected := map[string]string{
					".claude-plugin":  "0.2.4",
					"skills/patterns": "0.2.1",
					"skills/enforce":  "0.3.0",
					"skills/build":    "0.1.4",
					"skills/secure":   "0.2.2",
				}
				if len(versions) != len(expected) {
					t.Errorf("expected %d versions, got %d", len(expected), len(versions))
				}
				for key, expectedVer := range expected {
					if versions[key] != expectedVer {
						t.Errorf("expected version %q for %q, got %q", expectedVer, key, versions[key])
					}
				}
			},
		},
		{
			name: "empty manifest",
			setupFiles: map[string]string{
				"empty.json": "../../services/testdata/empty_manifest.json",
			},
			path:        "empty.json",
			wantErr:     true,
			errContains: "release manifest is empty",
		},
		{
			name: "malformed JSON",
			setupFiles: map[string]string{
				"malformed.json": "../../services/testdata/malformed.json",
			},
			path:        "malformed.json",
			wantErr:     true,
			errContains: "failed to parse release manifest",
		},
		{
			name:        "file not found",
			setupFiles:  map[string]string{},
			path:        "nonexistent.json",
			wantErr:     true,
			errContains: "failed to read release manifest",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock filesystem
			mockFS := NewMockFileSystem()

			// Load fixture files
			for path, fixturePath := range tt.setupFiles {
				content, err := readTestFixture(fixturePath)
				if err != nil {
					t.Fatalf("failed to read test fixture %s: %v", fixturePath, err)
				}
				mockFS.AddFile(path, content)
			}

			// Create reader with mock filesystem
			reader := NewConfigReader(mockFS)

			// Execute
			versions, err := reader.ReadReleaseManifest(tt.path)

			// Assert error expectations
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

			// Assert success expectations
			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if tt.validate != nil {
				tt.validate(t, versions)
			}
		})
	}
}

// Helper functions

func readTestFixture(fixturePath string) ([]byte, error) {
	// Read actual fixture file from testdata directory
	realFS := NewFileSystem()
	return realFS.ReadFile(fixturePath)
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > len(substr) && containsSubstring(s, substr))
}

func containsSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
