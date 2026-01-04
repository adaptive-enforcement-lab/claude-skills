package ports

import "github.com/adaptive-enforcement-lab/claude-skills/internal/domain"

// FrontmatterParser extracts YAML frontmatter from markdown content.
type FrontmatterParser interface {
	// Parse extracts frontmatter and returns it along with the remaining markdown content.
	// Returns (frontmatter, remaining markdown, error).
	Parse(content []byte) (*domain.Frontmatter, string, error)
}

// SectionParser parses markdown content into hierarchical sections.
type SectionParser interface {
	// Parse parses markdown and returns a tree of sections.
	Parse(markdown string) ([]domain.Section, error)

	// ExtractIntroduction extracts content before the first heading (after title).
	ExtractIntroduction(markdown string) string
}

// ContentExtractor extracts specific types of content from markdown.
type ContentExtractor interface {
	// ExtractCodeBlocks finds all fenced code blocks in the markdown.
	ExtractCodeBlocks(content string) []domain.CodeBlock

	// ExtractMermaid finds all Mermaid diagram blocks.
	ExtractMermaid(content string) []domain.MermaidDiagram

	// ExtractTables finds all markdown tables.
	ExtractTables(content string) []domain.Table

	// ExtractAdmonitions finds all Material for MkDocs admonition blocks.
	ExtractAdmonitions(content string) []domain.Admonition
}

// AdmonitionConverter converts MkDocs admonitions to standard markdown blockquotes.
type AdmonitionConverter interface {
	// Convert transforms "!!! type 'title'" syntax to "> **title**" blockquotes.
	Convert(content string) string
}
