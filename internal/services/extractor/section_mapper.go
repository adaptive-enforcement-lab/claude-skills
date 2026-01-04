package extractor

import (
	"strings"

	"github.com/adaptive-enforcement-lab/claude-skills/internal/domain"
)

// SectionMapper implements ports.SectionMapper with fuzzy matching.
type SectionMapper struct {
	// mappings defines the fuzzy matching rules for section names
	mappings map[string][]string
}

// NewSectionMapper creates a new section mapper with predefined mappings.
func NewSectionMapper() *SectionMapper {
	return &SectionMapper{
		mappings: map[string][]string{
			"WhenToUse": {
				"Why It Matters",
				"When to Use",
				"Use Cases",
				"What You'll Learn",
				"Abstract",
				"Overview",
			},
			"Prerequisites": {
				"Prerequisites",
				"Before You Begin",
				"Requirements",
				"What You Need",
			},
			"ImplementationSteps": {
				"Implementation",
				"Quick Start",
				"Getting Started",
				"How to",
				"Setup",
				"Installation",
				"Configuration",
			},
			"KeyPrinciples": {
				"Key Principles",
				"Core Principles",
				"Best Practices",
				"Guidelines",
				"Principles",
			},
			"Troubleshooting": {
				"Troubleshooting",
				"Common Issues",
				"Debugging",
				"FAQ",
				"Frequently Asked Questions",
				"Problems",
			},
			"References": {
				"References",
				"Related",
				"See Also",
				"Further Reading",
				"Additional Resources",
				"Resources",
			},
		},
	}
}

// MapSection attempts to map a section title to a known skill component name.
// Returns empty string if no mapping is found.
func (m *SectionMapper) MapSection(sectionTitle string) string {
	titleLower := strings.ToLower(strings.TrimSpace(sectionTitle))

	// Try exact match first
	for component, keywords := range m.mappings {
		for _, keyword := range keywords {
			if titleLower == strings.ToLower(keyword) {
				return component
			}
		}
	}

	// Try fuzzy match (contains)
	for component, keywords := range m.mappings {
		for _, keyword := range keywords {
			keywordLower := strings.ToLower(keyword)
			if strings.Contains(titleLower, keywordLower) || strings.Contains(keywordLower, titleLower) {
				return component
			}
		}
	}

	return ""
}

// FindSection searches for a section matching any of the given keywords.
// Returns nil if no match is found.
func (m *SectionMapper) FindSection(sections []domain.Section, keywords []string) *domain.Section {
	for i := range sections {
		section := &sections[i]

		// Check if this section matches any keyword
		for _, keyword := range keywords {
			if m.matchesKeyword(section.Title, keyword) {
				return section
			}
		}

		// Recursively search subsections
		if len(section.SubSections) > 0 {
			if found := m.FindSection(section.SubSections, keywords); found != nil {
				return found
			}
		}
	}

	return nil
}

// matchesKeyword checks if a section title matches a keyword (case-insensitive contains).
func (m *SectionMapper) matchesKeyword(title, keyword string) bool {
	titleLower := strings.ToLower(strings.TrimSpace(title))
	keywordLower := strings.ToLower(strings.TrimSpace(keyword))

	return strings.Contains(titleLower, keywordLower) || strings.Contains(keywordLower, titleLower)
}

// GetKeywordsForComponent returns the list of keywords for a given component.
func (m *SectionMapper) GetKeywordsForComponent(component string) []string {
	if keywords, ok := m.mappings[component]; ok {
		return keywords
	}
	return nil
}
