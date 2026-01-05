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

	// Check if we should generate reference.md (>200 lines)
	lineCount := strings.Count(doc.RawContent, "\n")
	if lineCount > 200 {
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

	// Extract ImplementationSteps (with code block filtering)
	if section := e.sectionMapper.FindSection(doc.Sections, mapper.GetKeywordsForComponent("ImplementationSteps")); section != nil {
		metadata.ImplementationSteps = e.extractContentWithShortCodeBlocks(section)
	}

	// Extract KeyPrinciples
	if section := e.sectionMapper.FindSection(doc.Sections, mapper.GetKeywordsForComponent("KeyPrinciples")); section != nil {
		metadata.KeyPrinciples = e.admonitionConverter.Convert(section.Content)
	}

	// Extract WhenToApply (with code block filtering)
	if section := e.sectionMapper.FindSection(doc.Sections, mapper.GetKeywordsForComponent("WhenToApply")); section != nil {
		metadata.WhenToApply = e.extractContentWithShortCodeBlocks(section)
	}

	// Extract Techniques (look for parent Techniques section and extract subsections)
	// But skip if the section is actually "Related Patterns"
	if section := e.sectionMapper.FindSection(doc.Sections, mapper.GetKeywordsForComponent("Techniques")); section != nil {
		// Don't extract if this is actually a "Related Patterns" section
		if !strings.Contains(strings.ToLower(section.Title), "related") {
			metadata.Techniques = e.extractTechniques(section, doc.CodeBlocks)
		}
	}

	// Extract Comparison (with code block filtering)
	if section := e.sectionMapper.FindSection(doc.Sections, mapper.GetKeywordsForComponent("Comparison")); section != nil {
		metadata.Comparison = e.extractContentWithShortCodeBlocks(section)
	}

	// Extract AntiPatterns (with code block filtering)
	if section := e.sectionMapper.FindSection(doc.Sections, mapper.GetKeywordsForComponent("AntiPatterns")); section != nil {
		metadata.AntiPatterns = e.extractContentWithShortCodeBlocks(section)
	}

	// Extract RelatedPatterns
	if section := e.sectionMapper.FindSection(doc.Sections, mapper.GetKeywordsForComponent("RelatedPatterns")); section != nil {
		metadata.RelatedPatterns = e.extractRelatedPatterns(section)
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

// extractFullSectionContent extracts section content including all subsections.
func (e *SkillExtractor) extractFullSectionContent(section *domain.Section) string {
	var content strings.Builder

	// Add main section content
	if section.Content != "" {
		content.WriteString(e.admonitionConverter.Convert(section.Content))
		content.WriteString("\n\n")
	}

	// Add subsections
	for _, subsection := range section.SubSections {
		content.WriteString("### ")
		content.WriteString(subsection.Title)
		content.WriteString("\n\n")
		content.WriteString(e.admonitionConverter.Convert(subsection.Content))
		content.WriteString("\n\n")
	}

	return strings.TrimSpace(content.String())
}

// extractTechniques extracts technique subsections from a parent section.
func (e *SkillExtractor) extractTechniques(section *domain.Section, codeBlocks []domain.CodeBlock) []domain.Technique {
	var techniques []domain.Technique

	// If section has subsections, each subsection is a technique
	// But only include first 5 techniques to keep SKILL.md manageable
	maxTechniques := 5
	for i, subsection := range section.SubSections {
		if i >= maxTechniques {
			break
		}
		technique := domain.Technique{
			Name:        subsection.Title,
			Description: extractFirstParagraph(subsection.Content),
			Content:     e.admonitionConverter.Convert(subsection.Content),
			CodeBlocks:  []domain.CodeBlock{}, // Could be enhanced to match code blocks by proximity
		}
		techniques = append(techniques, technique)
	}

	// If no subsections but content exists, treat whole section as one technique
	if len(techniques) == 0 && section.Content != "" {
		technique := domain.Technique{
			Name:        section.Title,
			Description: extractFirstParagraph(section.Content),
			Content:     e.admonitionConverter.Convert(section.Content),
			CodeBlocks:  []domain.CodeBlock{},
		}
		techniques = append(techniques, technique)
	}

	return techniques
}

// extractRelatedPatterns extracts related pattern references from a section.
func (e *SkillExtractor) extractRelatedPatterns(section *domain.Section) []string {
	var patterns []string

	// Look for markdown links in the content
	content := section.Content

	// Simple pattern extraction: look for [Pattern Name](url) format
	// This is a basic implementation - could be enhanced with more sophisticated parsing
	matches := strings.Split(content, "\n")

	for _, line := range matches {
		// Look for lines that look like pattern references
		if strings.Contains(line, "[") && strings.Contains(line, "]") {
			// Extract text between brackets as pattern name
			start := strings.Index(line, "[")
			end := strings.Index(line, "]")
			if start != -1 && end != -1 && end > start {
				patternName := line[start+1 : end]
				if patternName != "" && !strings.Contains(patternName, "!") {
					patterns = append(patterns, patternName)
				}
			}
		}
	}

	return patterns
}

// extractFirstParagraph extracts the first paragraph from content.
func extractFirstParagraph(content string) string {
	lines := strings.Split(content, "\n")
	var paragraph strings.Builder

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" {
			// Empty line marks end of paragraph
			if paragraph.Len() > 0 {
				break
			}
			continue
		}

		// Skip markdown headers
		if strings.HasPrefix(trimmed, "#") {
			continue
		}

		if paragraph.Len() > 0 {
			paragraph.WriteString(" ")
		}
		paragraph.WriteString(trimmed)
	}

	result := paragraph.String()
	// Limit to reasonable description length
	if len(result) > 200 {
		result = result[:197] + "..."
	}

	return result
}

// extractContentWithShortCodeBlocks extracts section content but replaces long code blocks (>10 lines) with references to examples.md
func (e *SkillExtractor) extractContentWithShortCodeBlocks(section *domain.Section) string {
	content := e.extractFullSectionContent(section)

	// Find code blocks in content
	lines := strings.Split(content, "\n")
	var result strings.Builder
	inCodeBlock := false
	codeBlockLines := 0
	codeBlockStart := 0

	for i, line := range lines {
		// Check if starting or ending code block
		if strings.HasPrefix(strings.TrimSpace(line), "```") {
			if !inCodeBlock {
				// Starting code block
				inCodeBlock = true
				codeBlockStart = i
				codeBlockLines = 0
			} else {
				// Ending code block
				inCodeBlock = false

				// If code block is longer than 10 lines, replace with reference
				if codeBlockLines > 10 {
					// Write placeholder instead of the long code block
					result.WriteString("\n*See [examples.md](examples.md) for detailed code examples.*\n")
				} else {
					// Include short code blocks - write buffered lines
					for j := codeBlockStart; j <= i; j++ {
						result.WriteString(lines[j])
						result.WriteString("\n")
					}
				}
				continue
			}
		}

		if inCodeBlock {
			codeBlockLines++
			// Don't write yet - wait to see if block is too long
		} else {
			result.WriteString(line)
			result.WriteString("\n")
		}
	}

	return strings.TrimSpace(result.String())
}
