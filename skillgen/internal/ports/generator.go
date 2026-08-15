package ports

import "github.com/adaptive-enforcement-lab/claude-skills/skillgen/internal/domain"

// SkillGenerator generates complete skills from domain models.
type SkillGenerator interface {
	// Generate creates a skill with all its components from a document.
	Generate(doc *domain.Document) (*domain.Skill, error)
}

// TemplateRenderer renders a hub skill using Go templates.
type TemplateRenderer interface {
	// RenderSkill renders the main SKILL.md file.
	RenderSkill(skill *domain.Skill) (string, error)

	// RenderReference renders the reference.md file: the full, offline
	// depth behind the SKILL.md link index.
	RenderReference(skill *domain.Skill) (string, error)
}
