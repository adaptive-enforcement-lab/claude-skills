package extractor

import (
	"fmt"
	"strings"

	"github.com/adaptive-enforcement-lab/claude-skills/skillgen/internal/domain"
)

// NameDeriver implements ports.NameDeriver.
type NameDeriver struct{}

// NewNameDeriver creates a new name deriver.
func NewNameDeriver() *NameDeriver {
	return &NameDeriver{}
}

// DeriveSkillName converts a title to kebab-case skill name.
// Example: "GitHub Actions Integration" -> "github-actions-integration"
func (d *NameDeriver) DeriveSkillName(title string) string {
	// Convert to lowercase
	name := strings.ToLower(title)

	// Replace spaces and special characters with hyphens
	name = strings.Map(func(r rune) rune {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			return r
		}
		if r == ' ' || r == '-' || r == '_' || r == '/' {
			return '-'
		}
		return -1 // Remove character
	}, name)

	// Remove consecutive hyphens
	for strings.Contains(name, "--") {
		name = strings.ReplaceAll(name, "--", "-")
	}

	// Trim hyphens from start and end
	name = strings.Trim(name, "-")

	return name
}

// DeriveFilename generates a filename from a code block.
func (d *NameDeriver) DeriveFilename(block domain.CodeBlock, index int) string {
	// If block already has a filename, use it
	if block.Filename != "" {
		return block.Filename
	}

	// Try to infer from language
	ext := languageToExtension(block.Language)

	// Generate default filename
	if ext != "" {
		return fmt.Sprintf("example-%d.%s", index+1, ext)
	}

	return fmt.Sprintf("code-%d.txt", index+1)
}

// languageToExtension maps language names to file extensions.
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
	case "dockerfile":
		return "Dockerfile"
	case "makefile":
		return "Makefile"
	case "toml":
		return "toml"
	case "ini":
		return "ini"
	case "sql":
		return "sql"
	default:
		if language != "" {
			return language
		}
		return ""
	}
}
