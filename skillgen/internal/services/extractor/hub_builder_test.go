package extractor

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/adaptive-enforcement-lab/claude-skills/skillgen/internal/adapters/parser"
	"github.com/adaptive-enforcement-lab/claude-skills/skillgen/internal/domain"
)

func newTestHubBuilder() *HubBuilder {
	return NewHubBuilder(NewTopicExtractor(), parser.NewAdmonitionConverter())
}

func doc(pathSegments []string, title, description, introduction string) *domain.Document {
	return &domain.Document{
		Path: filepath.Join(pathSegments...),
		Frontmatter: domain.Frontmatter{
			Title:       title,
			Description: description,
		},
		Introduction: introduction,
	}
}

// docWithBody is like doc but also sets RawContent, the source ReferenceBody
// is built from.
func docWithBody(pathSegments []string, title, description, introduction, rawContent string) *domain.Document {
	d := doc(pathSegments, title, description, introduction)
	d.RawContent = rawContent
	return d
}

func TestHubBuilderBuildsOverviewFromCategoryRoot(t *testing.T) {
	docs := []*domain.Document{
		doc([]string{"docs", "patterns", "index.md"}, "Patterns", "",
			"First sentence. Second sentence. Third sentence. Fourth sentence."),
	}

	hub, err := newTestHubBuilder().Build("patterns", docs, domain.PluginConfig{Description: "Curated description."})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := "First sentence. Second sentence. Third sentence."
	if hub.Metadata.Overview != want {
		t.Errorf("Overview = %q, want %q", hub.Metadata.Overview, want)
	}
	if hub.Metadata.Name != "patterns" {
		t.Errorf("Name = %q, want %q", hub.Metadata.Name, "patterns")
	}
	if hub.Metadata.Description != "Curated description." {
		t.Errorf("Description = %q, want the plugin-metadata description", hub.Metadata.Description)
	}
}

func TestHubBuilderGroupsByFirstPathSegment(t *testing.T) {
	docs := []*domain.Document{
		doc([]string{"docs", "patterns", "index.md"}, "Patterns", "d", "Intro."),
		doc([]string{"docs", "patterns", "architecture", "index.md"}, "Architecture Patterns", "Structural patterns.", ""),
		doc([]string{"docs", "patterns", "architecture", "hub-and-spoke", "index.md"}, "Hub and Spoke", "Central coordinator.", ""),
		doc([]string{"docs", "patterns", "architecture", "strangler-fig", "index.md"}, "Strangler Fig", "Incremental migration.", ""),
	}

	hub, err := newTestHubBuilder().Build("patterns", docs, domain.PluginConfig{Description: "d"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(hub.Groups) != 1 {
		t.Fatalf("expected 1 group, got %d: %+v", len(hub.Groups), hub.Groups)
	}

	group := hub.Groups[0]
	if group.Title != "Architecture Patterns" {
		t.Errorf("group title = %q, want %q (from architecture/index.md)", group.Title, "Architecture Patterns")
	}
	if group.Description != "Structural patterns." {
		t.Errorf("group description = %q", group.Description)
	}
	if len(group.Topics) != 2 {
		t.Fatalf("expected 2 topics, got %d: %+v", len(group.Topics), group.Topics)
	}
	if group.Topics[0].Title != "Hub and Spoke" || group.Topics[1].Title != "Strangler Fig" {
		t.Errorf("topics not sorted alphabetically: %+v", group.Topics)
	}
}

func TestHubBuilderFallsBackToHumanizedGroupTitle(t *testing.T) {
	docs := []*domain.Document{
		doc([]string{"docs", "enforce", "index.md"}, "Enforce", "d", ""),
		doc([]string{"docs", "enforce", "github-actions", "branch-protection", "index.md"}, "Branch Protection", "Require checks.", ""),
	}

	hub, err := newTestHubBuilder().Build("enforce", docs, domain.PluginConfig{Description: "d"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(hub.Groups) != 1 {
		t.Fatalf("expected 1 group, got %d", len(hub.Groups))
	}
	if hub.Groups[0].Title != "Github Actions" {
		t.Errorf("group title = %q, want humanized slug %q", hub.Groups[0].Title, "Github Actions")
	}
	if len(hub.Groups[0].Topics) != 1 || hub.Groups[0].Topics[0].Title != "Branch Protection" {
		t.Errorf("expected the leaf doc to become a topic, got %+v", hub.Groups[0].Topics)
	}
}

func TestHubBuilderErrorsWithoutCategoryRoot(t *testing.T) {
	docs := []*domain.Document{
		doc([]string{"docs", "patterns", "architecture", "index.md"}, "Architecture Patterns", "d", ""),
	}

	if _, err := newTestHubBuilder().Build("patterns", docs, domain.PluginConfig{}); err == nil {
		t.Fatal("expected an error when no category root index.md is present")
	}
}

func TestHubBuilderSortsGroupsAlphabetically(t *testing.T) {
	docs := []*domain.Document{
		doc([]string{"docs", "patterns", "index.md"}, "Patterns", "d", ""),
		doc([]string{"docs", "patterns", "efficiency", "idempotency", "index.md"}, "Idempotency", "d", ""),
		doc([]string{"docs", "patterns", "architecture", "hub-and-spoke", "index.md"}, "Hub and Spoke", "d", ""),
	}

	hub, err := newTestHubBuilder().Build("patterns", docs, domain.PluginConfig{Description: "d"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(hub.Groups) != 2 || hub.Groups[0].Title != "Architecture" || hub.Groups[1].Title != "Efficiency" {
		t.Errorf("groups not sorted alphabetically: %+v", hub.Groups)
	}
}

func TestHubBuilderPopulatesReferenceBodies(t *testing.T) {
	docs := []*domain.Document{
		docWithBody([]string{"docs", "patterns", "index.md"}, "Patterns", "d", "Intro.",
			"# Patterns\n\nReusable design patterns.\n\n## Overview\n\nMore about patterns."),
		docWithBody([]string{"docs", "patterns", "architecture", "index.md"}, "Architecture Patterns", "Structural patterns.", "",
			"# Architecture Patterns\n\nGroup body content."),
		docWithBody([]string{"docs", "patterns", "architecture", "hub-and-spoke", "index.md"}, "Hub and Spoke", "Central coordinator.", "",
			"# Hub and Spoke\n\nOne coordinator, many workers.\n\n## Trade-offs\n\nScales horizontally."),
	}

	hub, err := newTestHubBuilder().Build("patterns", docs, domain.PluginConfig{Description: "d"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if got, want := hub.Metadata.ReferenceBody, "Reusable design patterns.\n\n### Overview\n\nMore about patterns."; got != want {
		t.Errorf("root ReferenceBody = %q, want %q", got, want)
	}

	group := hub.Groups[0]
	if got, want := group.ReferenceBody, "Group body content."; got != want {
		t.Errorf("group ReferenceBody = %q, want %q", got, want)
	}

	topic := group.Topics[0]
	if got, want := topic.ReferenceBody, "One coordinator, many workers.\n\n#### Trade-offs\n\nScales horizontally."; got != want {
		t.Errorf("topic ReferenceBody = %q, want %q", got, want)
	}

	// Regression guard for the old SectionMapper bug: each doc's body must
	// appear in the reference tree exactly once.
	occurrences := strings.Count(hub.Metadata.ReferenceBody, "More about patterns.") +
		strings.Count(group.ReferenceBody, "More about patterns.") +
		strings.Count(topic.ReferenceBody, "More about patterns.")
	if occurrences != 1 {
		t.Errorf("expected root content to appear exactly once across the reference tree, got %d", occurrences)
	}
}

func TestHubBuilderPopulatesLibraryFiles(t *testing.T) {
	docs := []*domain.Document{
		docWithBody([]string{"docs", "patterns", "index.md"}, "Patterns", "d", "",
			"# Patterns\n\nReusable design patterns."),
		docWithBody([]string{"docs", "patterns", "architecture", "index.md"}, "Architecture Patterns", "d", "",
			"# Architecture Patterns\n\nGroup body content."),
		docWithBody([]string{"docs", "patterns", "architecture", "hub-and-spoke", "index.md"}, "Hub and Spoke", "d", "",
			"# Hub and Spoke\n\nOne coordinator, many workers."),
	}

	hub, err := newTestHubBuilder().Build("patterns", docs, domain.PluginConfig{Description: "d"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(hub.LibraryFiles) != len(docs) {
		t.Fatalf("expected %d library files, got %d: %+v", len(docs), len(hub.LibraryFiles), hub.LibraryFiles)
	}

	byPath := make(map[string]domain.LibraryFile, len(hub.LibraryFiles))
	for _, lf := range hub.LibraryFiles {
		byPath[lf.RelPath] = lf
	}

	root, ok := byPath["index.md"]
	if !ok {
		t.Fatalf("expected a library file at index.md for the category root, got paths: %+v", keysOf(byPath))
	}
	if !strings.HasPrefix(root.Content, "# Patterns\n\nSource: https://adaptive-enforcement-lab.com/patterns/\n\n") {
		t.Errorf("root library content = %q, want title + source note preserved at top", root.Content)
	}
	if !strings.Contains(root.Content, "Reusable design patterns.") {
		t.Errorf("root library content missing original body: %q", root.Content)
	}

	group, ok := byPath["architecture/index.md"]
	if !ok {
		t.Fatalf("expected a library file at architecture/index.md, got paths: %+v", keysOf(byPath))
	}
	if !strings.Contains(group.Content, "Source: https://adaptive-enforcement-lab.com/patterns/architecture/") {
		t.Errorf("group library content missing source note: %q", group.Content)
	}

	topic, ok := byPath["architecture/hub-and-spoke/index.md"]
	if !ok {
		t.Fatalf("expected a library file at architecture/hub-and-spoke/index.md, got paths: %+v", keysOf(byPath))
	}
	want := "# Hub and Spoke\n\nSource: https://adaptive-enforcement-lab.com/patterns/architecture/hub-and-spoke/\n\nOne coordinator, many workers."
	if topic.Content != want {
		t.Errorf("topic library content = %q, want %q", topic.Content, want)
	}
}

func keysOf(m map[string]domain.LibraryFile) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func TestHubBuilderLibraryFilesHandleLeadingBlankLine(t *testing.T) {
	// Real AEL docs have a blank line between frontmatter and the title
	// (RawContent = "\n# Title\n\nBody."), which broke a naive "is line
	// zero a heading" check.
	docs := []*domain.Document{
		docWithBody([]string{"docs", "patterns", "index.md"}, "Patterns", "d", "",
			"\n# Patterns\n\nReusable design patterns."),
		docWithBody([]string{"docs", "patterns", "architecture", "index.md"}, "Architecture Patterns", "d", "",
			"\n# Architecture Patterns\n\nOne hub coordinates."),
	}

	hub, err := newTestHubBuilder().Build("patterns", docs, domain.PluginConfig{Description: "d"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var root domain.LibraryFile
	for _, lf := range hub.LibraryFiles {
		if lf.RelPath == "index.md" {
			root = lf
		}
	}

	want := "# Patterns\n\nSource: https://adaptive-enforcement-lab.com/patterns/\n\nReusable design patterns."
	if root.Content != want {
		t.Errorf("root library content = %q, want %q (title must come before the source note)", root.Content, want)
	}
}
