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

// ExtractIntroduction extracts content before the first heading (after title).
func (p *SectionParser) ExtractIntroduction(markdown string) string {
	source := []byte(markdown)
	reader := text.NewReader(source)

	// Parse markdown into AST
	doc := p.markdown.Parser().Parse(reader)

	// Find first heading
	var firstHeadingPos int = -1
	var titleHeadingEnd int = 0

	ast.Walk(doc, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering {
			return ast.WalkContinue, nil
		}

		if heading, ok := n.(*ast.Heading); ok {
			headingStart := heading.Lines().At(0).Start

			// First heading is the title (H1)
			if heading.Level == 1 && titleHeadingEnd == 0 {
				titleHeadingEnd = heading.Lines().At(heading.Lines().Len() - 1).Stop
				return ast.WalkContinue, nil
			}

			// Second heading (first content heading) marks end of introduction
			if firstHeadingPos == -1 && titleHeadingEnd > 0 {
				firstHeadingPos = headingStart
				return ast.WalkStop, nil
			}
		}

		return ast.WalkContinue, nil
	})

	// Extract introduction content
	if titleHeadingEnd > 0 {
		endPos := len(source)
		if firstHeadingPos > 0 {
			endPos = firstHeadingPos
		}

		introduction := strings.TrimSpace(string(source[titleHeadingEnd:endPos]))

		// Remove any trailing heading markers
		introduction = p.removeTrailingHeadings(introduction)

		return introduction
	}

	return ""
}

// extractSections walks the AST and extracts sections.
func (p *SectionParser) extractSections(node ast.Node, source []byte) []domain.Section {
	// First pass: collect all headings with their positions
	type headingInfo struct {
		title     string
		level     int
		startPos  int
		endPos    int
		astNode   ast.Node
	}

	var headings []headingInfo

	ast.Walk(node, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering {
			return ast.WalkContinue, nil
		}

		if heading, ok := n.(*ast.Heading); ok {
			title := p.extractText(heading, source)
			startPos := heading.Lines().At(0).Start
			endPos := heading.Lines().At(heading.Lines().Len() - 1).Stop

			headings = append(headings, headingInfo{
				title:    title,
				level:    heading.Level,
				startPos: startPos,
				endPos:   endPos,
				astNode:  n,
			})
		}

		return ast.WalkContinue, nil
	})

	if len(headings) == 0 {
		return nil
	}

	// Second pass: build sections with content
	var sections []domain.Section

	for i, heading := range headings {
		// Determine content end position (start of next heading of same or higher level)
		contentEnd := len(source)
		for j := i + 1; j < len(headings); j++ {
			if headings[j].level <= heading.level {
				contentEnd = headings[j].startPos
				break
			}
		}

		// Extract content between heading and next section
		contentStart := heading.endPos
		content := strings.TrimSpace(string(source[contentStart:contentEnd]))

		// Remove any trailing heading markers that might have been included
		content = p.removeTrailingHeadings(content)

		// Build section
		section := domain.Section{
			Title:     heading.title,
			Level:     heading.level,
			Content:   content,
			LineStart: heading.startPos,
			LineEnd:   contentEnd,
		}

		// Handle subsections
		for j := i + 1; j < len(headings); j++ {
			if headings[j].level <= heading.level {
				break
			}
			if headings[j].level == heading.level+1 {
				// This is a direct subsection
				subContentEnd := len(source)
				for k := j + 1; k < len(headings); k++ {
					if headings[k].level <= headings[j].level {
						subContentEnd = headings[k].startPos
						break
					}
				}

				subContentStart := headings[j].endPos
				subContent := strings.TrimSpace(string(source[subContentStart:subContentEnd]))

				// Remove any trailing heading markers
				subContent = p.removeTrailingHeadings(subContent)

				subsection := domain.Section{
					Title:     headings[j].title,
					Level:     headings[j].level,
					Content:   subContent,
					LineStart: headings[j].startPos,
					LineEnd:   subContentEnd,
				}

				section.SubSections = append(section.SubSections, subsection)
			}
		}

		// Only add top-level sections (subsections are nested)
		isTopLevel := true
		for j := 0; j < i; j++ {
			if headings[j].level < heading.level {
				// This heading is under a previous heading
				isTopLevel = false
				break
			}
		}

		if isTopLevel {
			sections = append(sections, section)
		}
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

// removeTrailingHeadings removes any heading markers (##, ###, etc.) from the end of content.
// This can happen when section content includes the start of the next section's heading.
func (p *SectionParser) removeTrailingHeadings(content string) string {
	lines := strings.Split(content, "\n")

	// Remove trailing lines that are just heading markers or empty
	for len(lines) > 0 {
		lastLine := strings.TrimSpace(lines[len(lines)-1])

		// Check if line is empty or just heading markers
		if lastLine == "" || (strings.HasPrefix(lastLine, "#") && strings.TrimLeft(lastLine, "#") == "") {
			lines = lines[:len(lines)-1]
		} else {
			break
		}
	}

	return strings.Join(lines, "\n")
}

// extractSectionContent is no longer needed as content is extracted during section parsing.
// Kept for backwards compatibility but returns empty string.
func (p *SectionParser) extractSectionContent(section *domain.Section, source []byte) string {
	return section.Content
}
