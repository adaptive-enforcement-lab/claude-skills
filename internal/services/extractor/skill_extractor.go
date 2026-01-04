package extractor

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/adaptive-enforcement-lab/claude-skills/internal/domain"
	"github.com/adaptive-enforcement-lab/claude-skills/internal/ports"
)

// SkillExtractor implements ports.SkillExtractor.
type SkillExtractor struct {
	sectionMapper     ports.SectionMapper
	nameDeriver       ports.NameDeriver
	admonitionConverter ports.AdmonitionConverter
}

// NewSkillExtractor creates a new skill extractor.
func NewSkillExtractor(
	sectionMapper ports.SectionMapper,
	nameDeriver ports.NameDeriver,
	admonitionConverter ports.AdmonitionConverter,
) *SkillExtractor {
	return &SkillExtractor{
		sectionMapper:     sectionMapper,
		nameDeriver:       nameDeriver,
		admonitionConverter: admonitionConverter,
	}
}

// Extract analyzes a document and extracts all components needed for skill generation.
func (e *SkillExtractor) Extract(doc *domain.Document) (*domain.Skill, error) {
	// Derive skill name from title
	skillName := e.nameDeriver.DeriveSkillName(doc.Frontmatter.Title)
	if skillName == "" {
		return nil, fmt.Errorf("cannot derive skill name from empty title")
	}

	// Determine category from path
	category := e.determineCategoryFromPath(doc.Path)
	if category == "" {
		return nil, fmt.Errorf("cannot determine category from path: %s", doc.Path)
	}

	// Build skill metadata
	metadata := domain.SkillMetadata{
		Name:        skillName,
		Title:       doc.Frontmatter.Title,
		Description: doc.Frontmatter.Description,
		Category:    category,
		Tags:        doc.Frontmatter.Tags,
		SourcePath:  doc.Path,
		SourceURL:   e.buildSourceURL(doc.Path, category),
	}

	// Extract mapped sections
	e.extractMappedSections(doc, &metadata)

	// Build skill
	skill := &domain.Skill{
		Metadata:    metadata,
		MainContent: "", // Will be filled by template renderer
	}

	// Check if we should generate examples.md (≥2 code blocks)
	if len(doc.CodeBlocks) >= 2 {
		skill.Examples = &domain.ExamplesDoc{
			CodeBlocks: doc.CodeBlocks,
		}
	}

	// Check if we should generate troubleshooting.md
	troubleshootingSection := e.sectionMapper.FindSection(
		doc.Sections,
		e.sectionMapper.(*SectionMapper).GetKeywordsForComponent("Troubleshooting"),
	)
	if troubleshootingSection != nil {
		content := e.admonitionConverter.Convert(troubleshootingSection.Content)
		skill.Troubleshooting = &domain.TroubleshootingDoc{
			Content: content,
		}
	}

	// Check if we should generate reference.md (>500 lines)
	lineCount := strings.Count(doc.RawContent, "\n")
	if lineCount > 500 {
		skill.Reference = &domain.ReferenceDoc{
			Content:  e.admonitionConverter.Convert(doc.RawContent),
			Tables:   doc.Tables,
			Diagrams: doc.Mermaid,
		}
	}

	// Generate scripts from code blocks
	skill.Scripts = e.generateScripts(doc.CodeBlocks, skillName)

	return skill, nil
}

// extractMappedSections finds and maps sections to skill metadata fields.
func (e *SkillExtractor) extractMappedSections(doc *domain.Document, metadata *domain.SkillMetadata) {
	mapper := e.sectionMapper.(*SectionMapper)

	// Extract WhenToUse
	if section := e.sectionMapper.FindSection(doc.Sections, mapper.GetKeywordsForComponent("WhenToUse")); section != nil {
		metadata.WhenToUse = e.admonitionConverter.Convert(section.Content)
	} else if doc.Introduction != "" {
		// Fallback: use introduction if no "When to Use" section found
		metadata.WhenToUse = e.admonitionConverter.Convert(doc.Introduction)
	}

	// Extract Prerequisites
	if section := e.sectionMapper.FindSection(doc.Sections, mapper.GetKeywordsForComponent("Prerequisites")); section != nil {
		metadata.Prerequisites = e.admonitionConverter.Convert(section.Content)
	}
}

// determineCategoryFromPath extracts the category from the file path.
func (e *SkillExtractor) determineCategoryFromPath(path string) string {
	categories := []string{"patterns", "enforce", "build", "secure"}
	cleanPath := filepath.Clean(path)
	parts := strings.Split(cleanPath, string(filepath.Separator))

	for _, part := range parts {
		for _, category := range categories {
			if part == category {
				return category
			}
		}
	}

	return ""
}

// buildSourceURL constructs the URL to the source documentation.
func (e *SkillExtractor) buildSourceURL(path string, category string) string {
	baseURL := "https://adaptive-enforcement-lab.com"

	// Extract the topic from the path
	// Example: /docs/patterns/idempotency/index.md -> /patterns/idempotency/
	parts := strings.Split(filepath.Clean(path), string(filepath.Separator))

	// Find category in path
	for i, part := range parts {
		if part == category && i+1 < len(parts) {
			topic := parts[i+1]
			return fmt.Sprintf("%s/%s/%s/", baseURL, category, topic)
		}
	}

	return baseURL
}

// generateScripts creates script files from code blocks.
func (e *SkillExtractor) generateScripts(codeBlocks []domain.CodeBlock, skillName string) []domain.Script {
	var scripts []domain.Script

	for i, block := range codeBlocks {
		filename := e.nameDeriver.DeriveFilename(block, i)

		script := domain.Script{
			Filename: filename,
			Language: block.Language,
			Content:  block.Content,
			Path:     filepath.Join("scripts", filename),
		}

		scripts = append(scripts, script)
	}

	return scripts
}
