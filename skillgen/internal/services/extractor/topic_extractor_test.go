package extractor

import (
	"path/filepath"
	"testing"

	"github.com/adaptive-enforcement-lab/claude-skills/skillgen/internal/domain"
)

func TestTopicExtractorUsesFrontmatter(t *testing.T) {
	doc := &domain.Document{
		Path: filepath.Join("docs", "patterns", "architecture", "hub-and-spoke", "index.md"),
		Frontmatter: domain.Frontmatter{
			Title:       "Hub and Spoke",
			Description: "Centralized orchestration with distributed execution.",
		},
	}

	topic, err := NewTopicExtractor().Extract(doc)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if topic.Title != "Hub and Spoke" {
		t.Errorf("Title = %q, want %q", topic.Title, "Hub and Spoke")
	}
	if topic.Description != "Centralized orchestration with distributed execution." {
		t.Errorf("Description = %q", topic.Description)
	}

	want := "https://adaptive-enforcement-lab.com/patterns/architecture/hub-and-spoke/"
	if topic.URL != want {
		t.Errorf("URL = %q, want %q", topic.URL, want)
	}
}

func TestTopicExtractorFallsBackToIntroductionSentence(t *testing.T) {
	doc := &domain.Document{
		Path: filepath.Join("docs", "patterns", "architecture", "index.md"),
		Frontmatter: domain.Frontmatter{
			Title: "Architecture Patterns",
		},
		Introduction: "These patterns govern structure. They also govern behavior.",
	}

	topic, err := NewTopicExtractor().Extract(doc)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := "These patterns govern structure."
	if topic.Description != want {
		t.Errorf("Description = %q, want %q", topic.Description, want)
	}
}

func TestTopicExtractorRejectsEmptyTitle(t *testing.T) {
	doc := &domain.Document{Path: filepath.Join("docs", "patterns", "index.md")}

	if _, err := NewTopicExtractor().Extract(doc); err == nil {
		t.Fatal("expected an error for an empty title")
	}
}

func TestTopicExtractorRejectsUnknownCategory(t *testing.T) {
	doc := &domain.Document{
		Path:        filepath.Join("docs", "unknown-category", "index.md"),
		Frontmatter: domain.Frontmatter{Title: "Something"},
	}

	if _, err := NewTopicExtractor().Extract(doc); err == nil {
		t.Fatal("expected an error when the path has no known category segment")
	}
}
