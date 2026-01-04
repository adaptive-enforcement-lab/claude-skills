package parser

import (
	"bytes"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"

	"github.com/adaptive-enforcement-lab/claude-skills/internal/domain"
)

// SectionParser implements ports.SectionParser using goldmark.
type SectionParser struct {
	markdown goldmark.Markdown
}

// NewSectionParser creates a new goldmark-based section parser.
func NewSectionParser() *SectionParser {
	md := goldmark.New()
	return &SectionParser{
		markdown: md,
	}
}

// Parse parses markdown and returns a tree of sections.
func (p *SectionParser) Parse(markdown string) ([]domain.Section, error) {
	source := []byte(markdown)
	reader := text.NewReader(source)

	// Parse markdown into AST
	doc := p.markdown.Parser().Parse(reader)

	// Extract sections from AST
	sections := p.extractSections(doc, source)

	return sections, nil
}

// extractSections walks the AST and extracts sections.
func (p *SectionParser) extractSections(node ast.Node, source []byte) []domain.Section {
	var sections []domain.Section
	var currentSection *domain.Section

	ast.Walk(node, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering {
			return ast.WalkContinue, nil
		}

		switch v := n.(type) {
		case *ast.Heading:
			// Create new section
			title := p.extractText(v, source)
			level := v.Level

			section := domain.Section{
				Title:     title,
				Level:     level,
				LineStart: v.Lines().At(0).Start,
				LineEnd:   v.Lines().At(v.Lines().Len() - 1).Stop,
			}

			// If this is a top-level heading or same/higher level as current
			if currentSection == nil || level <= currentSection.Level {
				// Add previous section if exists
				if currentSection != nil {
					sections = append(sections, *currentSection)
				}
				currentSection = &section
			} else {
				// This is a subsection
				if currentSection != nil {
					currentSection.SubSections = append(currentSection.SubSections, section)
				}
			}
		}

		return ast.WalkContinue, nil
	})

	// Add final section
	if currentSection != nil {
		sections = append(sections, *currentSection)
	}

	// Extract content for each section
	for i := range sections {
		sections[i].Content = p.extractSectionContent(&sections[i], source)
	}

	return sections
}

// extractText extracts plain text from an AST node.
func (p *SectionParser) extractText(node ast.Node, source []byte) string {
	var buf bytes.Buffer

	ast.Walk(node, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering {
			return ast.WalkContinue, nil
		}

		switch v := n.(type) {
		case *ast.Text:
			buf.Write(v.Segment.Value(source))
		case *ast.String:
			buf.Write(v.Value)
		}

		return ast.WalkContinue, nil
	})

	return strings.TrimSpace(buf.String())
}

// extractSectionContent extracts the markdown content for a section.
func (p *SectionParser) extractSectionContent(section *domain.Section, source []byte) string {
	// This is a simplified implementation
	// In a more complete version, we'd extract content between this heading and the next
	return section.Title
}
