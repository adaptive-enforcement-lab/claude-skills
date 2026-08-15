package services

import (
	"fmt"
	"testing"

	"github.com/adaptive-enforcement-lab/claude-skills/skillgen/internal/domain"
)

type MockTemplateRenderer struct {
	readmeData    *domain.ReadmeData
	renderContent string
	renderError   error
}

func (m *MockTemplateRenderer) RenderSkill(skill *domain.Skill) (string, error) {
	return "", fmt.Errorf("not implemented in mock")
}

func (m *MockTemplateRenderer) RenderReference(skill *domain.Skill) (string, error) {
	return "", fmt.Errorf("not implemented in mock")
}

func (m *MockTemplateRenderer) RenderReadme(data *domain.ReadmeData) (string, error) {
	m.readmeData = data
	if m.renderError != nil {
		return "", m.renderError
	}
	return m.renderContent, nil
}

type MockFileSystem struct {
	written     map[string][]byte
	writeError  error
	writtenPath string
}

func NewMockFileSystem() *MockFileSystem {
	return &MockFileSystem{written: make(map[string][]byte)}
}

func (m *MockFileSystem) ReadFile(path string) ([]byte, error) {
	return nil, fmt.Errorf("not implemented")
}

func (m *MockFileSystem) WriteFile(path string, data []byte, perm int) error {
	if m.writeError != nil {
		return m.writeError
	}
	m.written[path] = data
	m.writtenPath = path
	return nil
}

func (m *MockFileSystem) MkdirAll(path string, perm int) error  { return nil }
func (m *MockFileSystem) Glob(pattern string) ([]string, error) { return nil, nil }
func (m *MockFileSystem) Exists(path string) bool               { return false }
func (m *MockFileSystem) IsDir(path string) bool                { return false }
func (m *MockFileSystem) RemoveAll(path string) error           { return nil }

func testHub(category, title, description string) *domain.Skill {
	return &domain.Skill{
		Metadata: domain.SkillMetadata{
			Category:    category,
			Title:       title,
			Description: description,
			SourceURL:   "https://adaptive-enforcement-lab.com/" + category + "/",
		},
		Groups:       []domain.TopicGroup{{Title: "A Group", Description: "d"}},
		LibraryFiles: []domain.LibraryFile{{RelPath: "index.md"}, {RelPath: "a/index.md"}},
	}
}

func testPluginMetadata() *domain.PluginMetadata {
	return &domain.PluginMetadata{
		Marketplace: domain.MarketplaceConfig{
			Name:  "ael-skills",
			Owner: domain.MarketplaceOwner{Name: "AEL"},
		},
		Plugins: map[string]domain.PluginConfig{
			"patterns": {Category: "development", Description: "d"},
			"enforce":  {Category: "security", Description: "d"},
		},
	}
}

func TestReadmeGenerator_Generate(t *testing.T) {
	renderer := &MockTemplateRenderer{renderContent: "# generated"}
	fs := NewMockFileSystem()
	logger := &MockLogger{}

	hubs := []*domain.Skill{
		testHub("enforce", "Enforce", "Use when enforcing policy."),
		testHub("patterns", "Patterns", "Use when building automation patterns that are reused often."),
	}

	gen := NewReadmeGenerator(renderer, fs, logger)
	if err := gen.Generate(hubs, testPluginMetadata(), map[string]string{
		"plugins/patterns": "1.1.0",
		"plugins/enforce":  "1.2.0",
	}, "README.md"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if fs.writtenPath != "README.md" {
		t.Errorf("wrote to %q, want README.md", fs.writtenPath)
	}
	if string(fs.written["README.md"]) != "# generated" {
		t.Errorf("wrote %q, want rendered template content", fs.written["README.md"])
	}

	data := renderer.readmeData
	if data == nil {
		t.Fatal("RenderReadme was not called with data")
	}
	if len(data.Hubs) != 2 {
		t.Fatalf("expected 2 hubs, got %d", len(data.Hubs))
	}

	// Sorted alphabetically by category: enforce before patterns.
	if data.Hubs[0].Category != "enforce" || data.Hubs[1].Category != "patterns" {
		t.Errorf("hubs not sorted alphabetically: %+v", data.Hubs)
	}

	enforce := data.Hubs[0]
	if enforce.Version != "1.2.0" {
		t.Errorf("enforce version = %q, want 1.2.0", enforce.Version)
	}
	if enforce.TopicCount != 2 {
		t.Errorf("enforce TopicCount = %d, want 2 (len of LibraryFiles)", enforce.TopicCount)
	}
	if enforce.CategoryLabel != "Security" {
		t.Errorf("enforce CategoryLabel = %q, want Security", enforce.CategoryLabel)
	}
	if enforce.Focus != "Use when enforcing policy." {
		t.Errorf("enforce Focus = %q, want the untruncated short description", enforce.Focus)
	}
}

func TestReadmeGenerator_TruncatesLongFocus(t *testing.T) {
	renderer := &MockTemplateRenderer{renderContent: "ok"}
	fs := NewMockFileSystem()
	gen := NewReadmeGenerator(renderer, fs, &MockLogger{})

	longDescription := "Use when doing one two three four five six seven eight nine ten eleven twelve thirteen fourteen fifteen sixteen."
	hubs := []*domain.Skill{testHub("patterns", "Patterns", longDescription)}

	if err := gen.Generate(hubs, testPluginMetadata(), nil, "README.md"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	got := renderer.readmeData.Hubs[0].Focus
	if got == longDescription {
		t.Error("expected Focus to be truncated, got the full description")
	}
	wantPrefix := "Use when doing one two three four five six seven eight nine ten eleven"
	if got[:len(got)-len("…")] != wantPrefix {
		t.Errorf("Focus = %q, want the first 14 words of %q", got, longDescription)
	}
}

func TestReadmeGenerator_MissingPluginMetadataErrors(t *testing.T) {
	renderer := &MockTemplateRenderer{}
	fs := NewMockFileSystem()
	gen := NewReadmeGenerator(renderer, fs, &MockLogger{})

	hubs := []*domain.Skill{testHub("build", "Build", "d")}

	err := gen.Generate(hubs, testPluginMetadata(), nil, "README.md")
	if err == nil {
		t.Fatal("expected an error for a category missing from plugin-metadata.json")
	}
}

func TestReadmeGenerator_MissingVersionDefaultsToZero(t *testing.T) {
	renderer := &MockTemplateRenderer{renderContent: "ok"}
	fs := NewMockFileSystem()
	gen := NewReadmeGenerator(renderer, fs, &MockLogger{})

	hubs := []*domain.Skill{testHub("patterns", "Patterns", "d")}

	if err := gen.Generate(hubs, testPluginMetadata(), map[string]string{}, "README.md"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if got := renderer.readmeData.Hubs[0].Version; got != "0.0.0" {
		t.Errorf("Version = %q, want 0.0.0 when absent from the manifest", got)
	}
}
