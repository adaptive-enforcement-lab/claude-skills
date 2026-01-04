package ports

import "github.com/adaptive-enforcement-lab/claude-skills/internal/domain"

// SkillGenerator generates complete skills from domain models.
type SkillGenerator interface {
	// Generate creates a skill with all its components from a document.
	Generate(doc *domain.Document) (*domain.Skill, error)
}

// TemplateRenderer renders skill components using Go templates.
type TemplateRenderer interface {
	// RenderSkill renders the main SKILL.md file.
	RenderSkill(skill *domain.Skill) (string, error)

	// RenderExamples renders the examples.md file.
	RenderExamples(skill *domain.Skill) (string, error)

	// RenderTroubleshooting renders the troubleshooting.md file.
	RenderTroubleshooting(skill *domain.Skill) (string, error)

	// RenderReference renders the reference.md file.
	RenderReference(skill *domain.Skill) (string, error)
}

// ScriptGenerator extracts code blocks to script files.
type ScriptGenerator interface {
	// GenerateScripts creates script files from code blocks.
	// Returns the list of scripts that should be written.
	GenerateScripts(codeBlocks []domain.CodeBlock, skillName string) []domain.Script
}
