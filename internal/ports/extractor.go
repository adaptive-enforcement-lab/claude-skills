package ports

import "github.com/adaptive-enforcement-lab/claude-skills/internal/domain"

// SectionMapper maps document sections to skill components using fuzzy matching.
// For example, "Why It Matters" maps to the "WhenToUse" field in the skill.
type SectionMapper interface {
	// MapSection attempts to map a section title to a known skill component name.
	// Returns empty string if no mapping is found.
	MapSection(sectionTitle string) string

	// FindSection searches for a section matching any of the given keywords.
	// Returns nil if no match is found.
	FindSection(sections []domain.Section, keywords []string) *domain.Section
}

// SkillExtractor transforms a parsed document into skill components.
type SkillExtractor interface {
	// Extract analyzes a document and extracts all components needed for skill generation.
	Extract(doc *domain.Document) (*domain.Skill, error)
}

// NameDeriver generates skill names from document titles.
type NameDeriver interface {
	// DeriveSkillName converts a title to kebab-case skill name.
	// Example: "GitHub Actions Integration" -> "github-actions-integration"
	DeriveSkillName(title string) string

	// DeriveFilename generates a filename from a code block.
	// Uses language extension and any hints in the code content.
	DeriveFilename(block domain.CodeBlock, index int) string
}
