package extractor

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/adaptive-enforcement-lab/claude-skills/skillgen/internal/domain"
)

// TopicExtractor implements ports.TopicExtractor.
type TopicExtractor struct{}

// NewTopicExtractor creates a new topic extractor.
func NewTopicExtractor() *TopicExtractor {
	return &TopicExtractor{}
}

// Extract derives a Topic from a document's frontmatter and path. Every AEL
// doc carries a hand-written title + description, so no prose extraction is
// needed; the description falls back to the first sentence of the intro only
// when frontmatter omits it.
func (e *TopicExtractor) Extract(doc *domain.Document) (*domain.Topic, error) {
	title := doc.Frontmatter.Title
	if title == "" {
		return nil, fmt.Errorf("cannot derive topic from empty title: %s", doc.Path)
	}

	description := firstSentence(doc.Frontmatter.Description)
	if description == "" {
		description = firstSentence(doc.Introduction)
	}

	category := determineCategoryFromPath(doc.Path)
	if category == "" {
		return nil, fmt.Errorf("cannot determine category from path: %s", doc.Path)
	}

	return &domain.Topic{
		Title:       title,
		Description: description,
		URL:         buildSourceURL(doc.Path, category),
	}, nil
}

// determineCategoryFromPath extracts the category from the file path.
func determineCategoryFromPath(path string) string {
	cleanPath := filepath.Clean(path)
	parts := strings.Split(cleanPath, string(filepath.Separator))

	for _, part := range parts {
		if domain.IsCategory(part) {
			return part
		}
	}

	return ""
}

// buildSourceURL constructs the URL to the source documentation.
//
// The docs site mirrors the directory layout, so every segment from the
// category directory down to the document's parent must be preserved.
// Example: /docs/patterns/efficiency/idempotency/index.md
//
//	-> /patterns/efficiency/idempotency/
func buildSourceURL(path string, category string) string {
	baseURL := "https://adaptive-enforcement-lab.com"

	parts := strings.Split(filepath.Clean(path), string(filepath.Separator))

	// Drop the filename; MkDocs serves index.md as its parent directory.
	if len(parts) > 0 && strings.HasSuffix(parts[len(parts)-1], ".md") {
		parts = parts[:len(parts)-1]
	}

	for i, part := range parts {
		if part == category {
			return fmt.Sprintf("%s/%s/", baseURL, strings.Join(parts[i:], "/"))
		}
	}

	return baseURL
}

// categorySegments returns the path segments strictly between the category
// directory and the document filename, e.g. for
// docs/patterns/architecture/hub-and-spoke/index.md with category "patterns"
// it returns ["architecture", "hub-and-spoke"].
func categorySegments(path, category string) []string {
	parts := strings.Split(filepath.Clean(path), string(filepath.Separator))

	if len(parts) > 0 && strings.HasSuffix(parts[len(parts)-1], ".md") {
		parts = parts[:len(parts)-1]
	}

	for i, part := range parts {
		if part == category {
			return parts[i+1:]
		}
	}

	return nil
}

// maxTopicDescriptionWords caps a topic's one-line description so a hub with
// dozens of topics still fits SKILL.md's word budget with real margin, not
// right at the wire. reference.md and library/ carry the full text
// regardless, so nothing is lost — SKILL.md is purely an index.
const maxTopicDescriptionWords = 6

// firstSentence returns the first sentence of the first paragraph of text,
// trimmed and capped to a short one-liner. Only the first paragraph is
// considered: admonitions and other blocks that follow a blank line in a
// doc's introduction aren't prose meant for a one-line summary.
func firstSentence(text string) string {
	trimmed := strings.TrimSpace(text)
	if trimmed == "" {
		return ""
	}

	if idx := strings.Index(trimmed, "\n\n"); idx != -1 {
		trimmed = trimmed[:idx]
	}

	if idx := strings.IndexAny(trimmed, ".!?"); idx != -1 {
		trimmed = trimmed[:idx+1]
	}

	trimmed = strings.Join(strings.Fields(trimmed), " ")

	return truncateWords(trimmed, maxTopicDescriptionWords)
}

// truncateWords caps text to at most maxWords words, appending an ellipsis
// if it had to cut.
func truncateWords(text string, maxWords int) string {
	words := strings.Fields(text)
	if len(words) <= maxWords {
		return text
	}

	return strings.TrimRight(strings.Join(words[:maxWords], " "), ".,;:!?") + "…"
}
