package parser

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"

	"github.com/adaptive-enforcement-lab/claude-skills/skillgen/internal/domain"
)

// ContentExtractor implements ports.ContentExtractor using goldmark.
type ContentExtractor struct {
	markdown goldmark.Markdown
}

// NewContentExtractor creates a new goldmark-based content extractor.
func NewContentExtractor() *ContentExtractor {
	md := goldmark.New()
	return &ContentExtractor{
		markdown: md,
	}
}

// ExtractCodeBlocks finds all fenced code blocks in the markdown.
func (e *ContentExtractor) ExtractCodeBlocks(content string) []domain.CodeBlock {
	source := []byte(content)
	reader := text.NewReader(source)

	// Parse markdown into AST
	doc := e.markdown.Parser().Parse(reader)

	var codeBlocks []domain.CodeBlock
	lineNum := 0

	ast.Walk(doc, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering {
			return ast.WalkContinue, nil
		}

		if codeBlock, ok := n.(*ast.FencedCodeBlock); ok {
			language := ""
			if codeBlock.Info != nil {
				language = string(codeBlock.Info.Text(source))
			}

			// Extract code content
			var content strings.Builder
			for i := 0; i < codeBlock.Lines().Len(); i++ {
				line := codeBlock.Lines().At(i)
				content.Write(line.Value(source))
			}

			// Infer filename from language
			filename := inferFilename(language, len(codeBlocks))

			codeBlocks = append(codeBlocks, domain.CodeBlock{
				Language: language,
				Content:  strings.TrimSpace(content.String()),
				Filename: filename,
				LineNum:  lineNum,
			})

			lineNum++
		}

		return ast.WalkContinue, nil
	})

	return codeBlocks
}

// ExtractMermaid finds all Mermaid diagram blocks.
func (e *ContentExtractor) ExtractMermaid(content string) []domain.MermaidDiagram {
	source := []byte(content)
	reader := text.NewReader(source)

	// Parse markdown into AST
	doc := e.markdown.Parser().Parse(reader)

	var diagrams []domain.MermaidDiagram
	lineNum := 0

	ast.Walk(doc, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering {
			return ast.WalkContinue, nil
		}

		if codeBlock, ok := n.(*ast.FencedCodeBlock); ok {
			language := ""
			if codeBlock.Info != nil {
				language = string(codeBlock.Info.Text(source))
			}

			// Only process mermaid blocks
			if language != "mermaid" {
				return ast.WalkContinue, nil
			}

			// Extract diagram content
			var diagramContent strings.Builder
			for i := 0; i < codeBlock.Lines().Len(); i++ {
				line := codeBlock.Lines().At(i)
				diagramContent.Write(line.Value(source))
			}

			diagrams = append(diagrams, domain.MermaidDiagram{
				Content: strings.TrimSpace(diagramContent.String()),
				Title:   fmt.Sprintf("Diagram %d", lineNum+1),
				LineNum: lineNum,
			})

			lineNum++
		}

		return ast.WalkContinue, nil
	})

	return diagrams
}

// ExtractTables finds all markdown tables.
func (e *ContentExtractor) ExtractTables(content string) []domain.Table {
	// Simple regex-based table extraction
	// Tables in markdown format: | header | header |
	tablePattern := regexp.MustCompile(`(?m)^\|.+\|\s*$`)
	lines := strings.Split(content, "\n")

	var tables []domain.Table
	var currentTable *domain.Table
	lineNum := 0

	for i, line := range lines {
		if tablePattern.MatchString(line) {
			if currentTable == nil {
				// Start new table
				currentTable = &domain.Table{
					LineNum: lineNum,
				}
			}

			// Parse table row
			cells := parseTableRow(line)

			// First row is headers (unless it's a separator row)
			if !isTableSeparator(line) {
				if len(currentTable.Headers) == 0 {
					currentTable.Headers = cells
				} else {
					currentTable.Rows = append(currentTable.Rows, cells)
				}
			}
		} else {
			// End of table
			if currentTable != nil {
				tables = append(tables, *currentTable)
				currentTable = nil
			}
		}

		// Handle end of content while in table
		if i == len(lines)-1 && currentTable != nil {
			tables = append(tables, *currentTable)
		}

		lineNum++
	}

	return tables
}

// ExtractAdmonitions finds all Material for MkDocs admonition blocks.
func (e *ContentExtractor) ExtractAdmonitions(content string) []domain.Admonition {
	// Pattern matches: !!! type "title"
	pattern := regexp.MustCompile(`(?m)^!!!\s+(\w+)\s+"([^"]+)"\s*$`)
	lines := strings.Split(content, "\n")

	var admonitions []domain.Admonition
	var currentAdmonition *domain.Admonition
	lineNum := 0

	for i, line := range lines {
		// Check if this line starts an admonition
		if match := pattern.FindStringSubmatch(line); match != nil {
			// Save previous admonition if exists
			if currentAdmonition != nil {
				admonitions = append(admonitions, *currentAdmonition)
			}

			// Start new admonition
			currentAdmonition = &domain.Admonition{
				Type:    match[1],
				Title:   match[2],
				LineNum: lineNum,
			}
		} else if currentAdmonition != nil {
			// Collect admonition content (indented with 4 spaces)
			if strings.HasPrefix(line, "    ") {
				content := strings.TrimPrefix(line, "    ")
				if currentAdmonition.Content == "" {
					currentAdmonition.Content = content
				} else {
					currentAdmonition.Content += "\n" + content
				}
			} else {
				// End of admonition (no longer indented)
				admonitions = append(admonitions, *currentAdmonition)
				currentAdmonition = nil
			}
		}

		// Handle end of content while in admonition
		if i == len(lines)-1 && currentAdmonition != nil {
			admonitions = append(admonitions, *currentAdmonition)
		}

		lineNum++
	}

	return admonitions
}

// Helper functions

func inferFilename(language string, index int) string {
	ext := languageToExtension(language)
	if ext == "" {
		return fmt.Sprintf("code-%d.txt", index+1)
	}
	return fmt.Sprintf("example-%d.%s", index+1, ext)
}

func languageToExtension(language string) string {
	switch strings.ToLower(language) {
	case "bash", "sh", "shell":
		return "sh"
	case "yaml", "yml":
		return "yaml"
	case "go", "golang":
		return "go"
	case "javascript", "js":
		return "js"
	case "typescript", "ts":
		return "ts"
	case "python", "py":
		return "py"
	case "json":
		return "json"
	case "xml":
		return "xml"
	case "markdown", "md":
		return "md"
	default:
		return language
	}
}

func parseTableRow(line string) []string {
	// Remove leading/trailing pipes
	line = strings.Trim(line, "|")

	// Split by pipe
	cells := strings.Split(line, "|")

	// Trim whitespace from each cell
	for i := range cells {
		cells[i] = strings.TrimSpace(cells[i])
	}

	return cells
}

func isTableSeparator(line string) bool {
	// Table separator row looks like: |---|---|
	return strings.Contains(line, "---")
}
