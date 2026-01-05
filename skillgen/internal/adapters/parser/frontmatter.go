package parser

import (
	"bytes"
	"fmt"
	"strings"
	"time"

	"gopkg.in/yaml.v3"

	"github.com/adaptive-enforcement-lab/claude-skills/skillgen/internal/domain"
)

// FrontmatterParser implements ports.FrontmatterParser using gopkg.in/yaml.v3.
type FrontmatterParser struct{}

// NewFrontmatterParser creates a new YAML frontmatter parser.
func NewFrontmatterParser() *FrontmatterParser {
	return &FrontmatterParser{}
}

// Parse extracts YAML frontmatter from markdown content.
// Returns (frontmatter, remaining markdown, error).
func (p *FrontmatterParser) Parse(content []byte) (*domain.Frontmatter, string, error) {
	// Check if content starts with frontmatter delimiter
	if !bytes.HasPrefix(content, []byte("---\n")) && !bytes.HasPrefix(content, []byte("---\r\n")) {
		// No frontmatter - return empty frontmatter and full content as markdown
		return &domain.Frontmatter{}, string(content), nil
	}

	// Find the end of frontmatter (second "---")
	lines := bytes.Split(content, []byte("\n"))
	if len(lines) < 3 {
		return nil, "", fmt.Errorf("invalid frontmatter: too few lines")
	}

	var endLine int
	for i := 1; i < len(lines); i++ {
		trimmed := bytes.TrimSpace(lines[i])
		if bytes.Equal(trimmed, []byte("---")) {
			endLine = i
			break
		}
	}

	if endLine == 0 {
		return nil, "", fmt.Errorf("invalid frontmatter: no closing delimiter found")
	}

	// Extract frontmatter content (between the two "---" lines)
	frontmatterLines := lines[1:endLine]
	frontmatterContent := bytes.Join(frontmatterLines, []byte("\n"))

	// Parse YAML
	var rawData map[string]interface{}
	if err := yaml.Unmarshal(frontmatterContent, &rawData); err != nil {
		return nil, "", fmt.Errorf("failed to parse frontmatter YAML: %w", err)
	}

	// Extract known fields
	frontmatter := &domain.Frontmatter{
		RawData: rawData,
	}

	// Title
	if title, ok := rawData["title"].(string); ok {
		frontmatter.Title = title
	}

	// Description (may be single line or multiline)
	if desc, ok := rawData["description"].(string); ok {
		frontmatter.Description = strings.TrimSpace(desc)
	}

	// Tags
	if tagsRaw, ok := rawData["tags"]; ok {
		if tagsList, ok := tagsRaw.([]interface{}); ok {
			for _, tag := range tagsList {
				if tagStr, ok := tag.(string); ok {
					frontmatter.Tags = append(frontmatter.Tags, tagStr)
				}
			}
		}
	}

	// Date (for blog post detection)
	if dateRaw, ok := rawData["date"]; ok {
		// Try parsing as time.Time first
		if dateTime, ok := dateRaw.(time.Time); ok {
			frontmatter.Date = &dateTime
		} else if dateStr, ok := dateRaw.(string); ok {
			// Try parsing string as date
			parsedDate, err := time.Parse("2006-01-02", dateStr)
			if err == nil {
				frontmatter.Date = &parsedDate
			}
		}
	}

	// Authors (for blog post detection)
	if authorsRaw, ok := rawData["authors"]; ok {
		if authorsList, ok := authorsRaw.([]interface{}); ok {
			for _, author := range authorsList {
				if authorStr, ok := author.(string); ok {
					frontmatter.Authors = append(frontmatter.Authors, authorStr)
				}
			}
		}
	}

	// Extract remaining markdown (everything after the second "---")
	markdownLines := lines[endLine+1:]
	markdown := string(bytes.Join(markdownLines, []byte("\n")))

	return frontmatter, markdown, nil
}
