package filesystem

import (
	"fmt"

	"github.com/adaptive-enforcement-lab/claude-skills/internal/domain"
	"github.com/adaptive-enforcement-lab/claude-skills/internal/ports"
)

// DocumentReader implements ports.DocumentReader using the filesystem.
type DocumentReader struct {
	fs                ports.FileSystem
	frontmatterParser ports.FrontmatterParser
	sectionParser     ports.SectionParser
	contentExtractor  ports.ContentExtractor
	categories        []string
}

// NewDocumentReader creates a new filesystem-based document reader.
func NewDocumentReader(
	fs ports.FileSystem,
	frontmatterParser ports.FrontmatterParser,
	sectionParser ports.SectionParser,
	contentExtractor ports.ContentExtractor,
	categories []string,
) *DocumentReader {
	return &DocumentReader{
		fs:                fs,
		frontmatterParser: frontmatterParser,
		sectionParser:     sectionParser,
		contentExtractor:  contentExtractor,
		categories:        categories,
	}
}

// ReadDocument reads and parses a single document file.
func (r *DocumentReader) ReadDocument(path string) (*domain.Document, error) {
	// Read file content
	content, err := r.fs.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", path, err)
	}

	// Parse frontmatter
	frontmatter, markdown, err := r.frontmatterParser.Parse(content)
	if err != nil {
		return nil, fmt.Errorf("failed to parse frontmatter in %s: %w", path, err)
	}

	// Parse sections
	sections, err := r.sectionParser.Parse(markdown)
	if err != nil {
		return nil, fmt.Errorf("failed to parse sections in %s: %w", path, err)
	}

	// Extract content components
	introduction := r.sectionParser.ExtractIntroduction(markdown)
	codeBlocks := r.contentExtractor.ExtractCodeBlocks(markdown)
	mermaid := r.contentExtractor.ExtractMermaid(markdown)
	tables := r.contentExtractor.ExtractTables(markdown)
	admonitions := r.contentExtractor.ExtractAdmonitions(markdown)

	// Build document
	doc := &domain.Document{
		Path:         path,
		Frontmatter:  *frontmatter,
		Introduction: introduction,
		Sections:     sections,
		CodeBlocks:   codeBlocks,
		Mermaid:      mermaid,
		Tables:       tables,
		Admonitions:  admonitions,
		RawContent:   markdown,
	}

	return doc, nil
}

// ListIndexFiles finds all index.md files in the specified root path.
func (r *DocumentReader) ListIndexFiles(rootPath string, categories []string) ([]string, error) {
	osFS, ok := r.fs.(*FileSystem)
	if !ok {
		return nil, fmt.Errorf("filesystem type not supported for directory walking")
	}

	return FindIndexFiles(osFS, rootPath, categories)
}
