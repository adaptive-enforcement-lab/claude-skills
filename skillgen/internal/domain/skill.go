package domain

// Skill represents a complete Claude Code skill with all its components.
// A skill may consist of multiple files: SKILL.md (required), examples.md, troubleshooting.md, reference.md, and scripts/.
type Skill struct {
	Metadata        SkillMetadata
	MainContent     string           // SKILL.md content (required)
	Examples        *ExamplesDoc     // nil if <2 code blocks
	Troubleshooting *TroubleshootingDoc // nil if no troubleshooting section
	Reference       *ReferenceDoc    // nil if total content <200 lines
	Scripts         []Script
}

// SkillMetadata contains the frontmatter and derived metadata for a skill.
type SkillMetadata struct {
	Name               string      // Kebab-case name derived from title
	Title              string      // Display title from frontmatter
	Description        string      // From frontmatter
	Category           string      // patterns, enforcement, build, secure
	Tags               []string
	WhenToUse          string      // Extracted from "Why It Matters" or similar sections
	Prerequisites      string      // Extracted from "Prerequisites" section
	ImplementationSteps string     // Extracted from "Implementation" sections
	KeyPrinciples      string      // Extracted from "Key Principles" sections
	WhenToApply        string      // Extracted from "When to Apply" sections (decision matrices)
	Techniques         []Technique // Extracted technique subsections
	Comparison         string      // Extracted comparison/contrast sections
	AntiPatterns       string      // Extracted anti-pattern sections
	RelatedPatterns    []string    // Extracted related pattern links
	SourcePath         string      // Original document path
	SourceURL          string      // URL to source documentation
}

// ExamplesDoc represents the examples.md file for a skill.
// Only generated if the source document contains ≥2 code blocks.
type ExamplesDoc struct {
	Content    string
	CodeBlocks []CodeBlock
}

// ShouldGenerate returns true if there are enough code blocks to warrant an examples file.
func (e *ExamplesDoc) ShouldGenerate() bool {
	return e != nil && len(e.CodeBlocks) >= 2
}

// TroubleshootingDoc represents the troubleshooting.md file for a skill.
// Only generated if the source document has a troubleshooting section.
type TroubleshootingDoc struct {
	Content string
}

// ShouldGenerate returns true if there is troubleshooting content.
func (t *TroubleshootingDoc) ShouldGenerate() bool {
	return t != nil && t.Content != ""
}

// ReferenceDoc represents the reference.md file for a skill.
// Only generated if the source document exceeds 200 lines.
type ReferenceDoc struct {
	Content string
	Tables  []Table
	Diagrams []MermaidDiagram
}

// ShouldGenerate returns true if the source document is large enough for a reference file.
func (r *ReferenceDoc) ShouldGenerate() bool {
	return r != nil && r.Content != ""
}

// Script represents a code block extracted to the scripts/ subdirectory.
type Script struct {
	Filename string // e.g., "install.sh", "config.yaml", "example.go"
	Language string // File extension hint
	Content  string
	Path     string // Relative path within skill directory
}

// Technique represents a subsection technique or method.
type Technique struct {
	Name        string      // Technique name/title
	Description string      // Brief description
	Content     string      // Full technique content
	CodeBlocks  []CodeBlock // Code examples specific to this technique
}

// Note: CodeBlock, Table, and MermaidDiagram are defined in document.go
// and can be used directly since they're in the same package.

// GetOutputPath returns the directory path where this skill should be written.
// Format: skills/{category}/{skill-name}/
func (s *Skill) GetOutputPath(baseDir string) string {
	// Implementation will be in the writer adapter
	return ""
}

// GetSkillFiles returns all files that should be written for this skill.
func (s *Skill) GetSkillFiles() []string {
	files := []string{"SKILL.md"} // Always required

	if s.Examples != nil && s.Examples.ShouldGenerate() {
		files = append(files, "examples.md")
	}

	if s.Troubleshooting != nil && s.Troubleshooting.ShouldGenerate() {
		files = append(files, "troubleshooting.md")
	}

	if s.Reference != nil && s.Reference.ShouldGenerate() {
		files = append(files, "reference.md")
	}

	if len(s.Scripts) > 0 {
		for _, script := range s.Scripts {
			files = append(files, script.Path)
		}
	}

	return files
}
