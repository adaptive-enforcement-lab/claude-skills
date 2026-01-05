package parser

import (
	"regexp"
	"strings"
)

// AdmonitionConverter implements ports.AdmonitionConverter.
type AdmonitionConverter struct {
	// Pattern matches: !!! type "title"
	pattern *regexp.Regexp
}

// NewAdmonitionConverter creates a new admonition converter.
func NewAdmonitionConverter() *AdmonitionConverter {
	// Match admonitions like: !!! abstract "What You'll Learn"
	pattern := regexp.MustCompile(`(?m)^!!!\s+(\w+)\s+"([^"]+)"\s*$`)

	return &AdmonitionConverter{
		pattern: pattern,
	}
}

// Convert transforms MkDocs admonition syntax to standard markdown blockquotes.
// Example: !!! tip "Helpful hint" -> > **Helpful hint**
func (c *AdmonitionConverter) Convert(content string) string {
	lines := strings.Split(content, "\n")
	var result []string
	inAdmonition := false
	var admonitionLines []string

	for i, line := range lines {
		// Check if this line starts an admonition
		if match := c.pattern.FindStringSubmatch(line); match != nil {
			// match[1] = type (abstract, tip, warning, etc.)
			// match[2] = title
			inAdmonition = true
			admonitionLines = []string{
				"> **" + match[2] + "**",
				">",
			}
			continue
		}

		// If in admonition, check if line is indented content
		if inAdmonition {
			// Admonition content is indented with 4 spaces
			if strings.HasPrefix(line, "    ") {
				// Strip 4 spaces and add to blockquote
				content := strings.TrimPrefix(line, "    ")
				if strings.TrimSpace(content) == "" {
					admonitionLines = append(admonitionLines, ">")
				} else {
					admonitionLines = append(admonitionLines, "> "+content)
				}
			} else if strings.TrimSpace(line) == "" {
				// Blank line within admonition - add empty blockquote line
				admonitionLines = append(admonitionLines, ">")
			} else {
				// Non-blank, non-indented line - exit admonition
				result = append(result, admonitionLines...)
				result = append(result, "")

				// Don't add the line if it's just a heading marker (##, ###, etc.)
				trimmed := strings.TrimSpace(line)
				if !(strings.HasPrefix(trimmed, "#") && strings.TrimLeft(trimmed, "#") == "") {
					result = append(result, line)
				}

				inAdmonition = false
				admonitionLines = nil
			}
		} else {
			result = append(result, line)
		}

		// Handle end of content while still in admonition
		if i == len(lines)-1 && inAdmonition {
			result = append(result, admonitionLines...)
		}
	}

	return strings.Join(result, "\n")
}
