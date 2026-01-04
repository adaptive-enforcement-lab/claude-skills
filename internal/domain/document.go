package domain

import "time"

// Document represents a parsed documentation file from the AEL repository.
// This is the core domain model for input documents that will be transformed into skills.
type Document struct {
	Path         string
	Frontmatter  Frontmatter
	Introduction string // Content before first heading (after title)
	Sections     []Section
	CodeBlocks   []CodeBlock
	Mermaid      []MermaidDiagram
	Tables       []Table
	Admonitions  []Admonition
	RawContent   string
	RelatedDocs  []string
}

// Frontmatter contains the YAML metadata at the beginning of a markdown document.
type Frontmatter struct {
	Title       string
	Description string
	Tags        []string
	Date        *time.Time // For blog post detection
	Authors     []string   // For blog post detection
	RawData     map[string]interface{}
}

// IsBlogPost returns true if the frontmatter indicates this is a blog post.
// Blog posts have both a date and authors field and should be excluded from skill generation.
func (f Frontmatter) IsBlogPost() bool {
	return f.Date != nil && len(f.Authors) > 0
}

// Section represents a markdown heading and its content.
// Sections are hierarchical, with subsections nested within parent sections.
type Section struct {
	Title       string
	Level       int // H1=1, H2=2, H3=3, etc.
	Content     string
	SubSections []Section
	LineStart   int
	LineEnd     int
}

// CodeBlock represents a fenced code block in markdown.
type CodeBlock struct {
	Language string // e.g., "bash", "yaml", "go", "json"
	Content  string
	Filename string // Inferred from comments or language extension
	LineNum  int
}

// MermaidDiagram represents a Mermaid diagram code block.
// Mermaid diagrams in AEL docs use the Ghostty Hardcore theme colors.
type MermaidDiagram struct {
	Content string
	Title   string
	LineNum int
}

// Table represents a markdown table with headers and rows.
type Table struct {
	Headers []string
	Rows    [][]string
	LineNum int
}

// Admonition represents a Material for MkDocs admonition block.
// These are formatted as "!!! type "title"" and need to be converted to blockquotes.
type Admonition struct {
	Type    string // abstract, tip, warning, success, info, note, danger
	Title   string
	Content string
	LineNum int
}

// DetermineCategory extracts the category from the document's file path.
// Categories map to skill collections: patterns, enforce, build, secure.
func (d *Document) DetermineCategory() string {
	// Extract category from path like "/docs/patterns/..." or "/docs/enforce/..."
	// Implementation will be in the parser adapter
	return ""
}

// GetSkillName derives a kebab-case skill name from the document title.
// For example: "GitHub Actions Integration" -> "github-actions-integration"
func (d *Document) GetSkillName() string {
	// Implementation will be in the extractor service
	return ""
}
