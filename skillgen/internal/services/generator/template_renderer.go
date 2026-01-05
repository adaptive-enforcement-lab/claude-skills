package generator

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/adaptive-enforcement-lab/claude-skills/skillgen/internal/domain"
)

// TemplateRenderer implements ports.TemplateRenderer using Go templates.
type TemplateRenderer struct {
	templates *template.Template
}

// NewTemplateRenderer creates a new template renderer.
func NewTemplateRenderer(templatesDir string) (*TemplateRenderer, error) {
	// Define custom template functions
	funcMap := template.FuncMap{
		"add": func(a, b int) int {
			return a + b
		},
		"title": strings.Title,
	}

	// Load all templates
	tmpl, err := template.New("").Funcs(funcMap).ParseGlob(templatesDir + "/*.tmpl")
	if err != nil {
		return nil, fmt.Errorf("failed to load templates from %s: %w", templatesDir, err)
	}

	return &TemplateRenderer{
		templates: tmpl,
	}, nil
}

// RenderSkill renders the main SKILL.md file.
func (r *TemplateRenderer) RenderSkill(skill *domain.Skill) (string, error) {
	var buf bytes.Buffer

	if err := r.templates.ExecuteTemplate(&buf, "skill.tmpl", skill); err != nil {
		return "", fmt.Errorf("failed to render skill template: %w", err)
	}

	return buf.String(), nil
}

// RenderExamples renders the examples.md file.
func (r *TemplateRenderer) RenderExamples(skill *domain.Skill) (string, error) {
	if skill.Examples == nil || !skill.Examples.ShouldGenerate() {
		return "", fmt.Errorf("skill has no examples to render")
	}

	var buf bytes.Buffer

	if err := r.templates.ExecuteTemplate(&buf, "examples.tmpl", skill); err != nil {
		return "", fmt.Errorf("failed to render examples template: %w", err)
	}

	return buf.String(), nil
}

// RenderTroubleshooting renders the troubleshooting.md file.
func (r *TemplateRenderer) RenderTroubleshooting(skill *domain.Skill) (string, error) {
	if skill.Troubleshooting == nil || !skill.Troubleshooting.ShouldGenerate() {
		return "", fmt.Errorf("skill has no troubleshooting content to render")
	}

	var buf bytes.Buffer

	if err := r.templates.ExecuteTemplate(&buf, "troubleshooting.tmpl", skill); err != nil {
		return "", fmt.Errorf("failed to render troubleshooting template: %w", err)
	}

	return buf.String(), nil
}

// RenderReference renders the reference.md file.
func (r *TemplateRenderer) RenderReference(skill *domain.Skill) (string, error) {
	if skill.Reference == nil || !skill.Reference.ShouldGenerate() {
		return "", fmt.Errorf("skill has no reference content to render")
	}

	var buf bytes.Buffer

	if err := r.templates.ExecuteTemplate(&buf, "reference.tmpl", skill); err != nil {
		return "", fmt.Errorf("failed to render reference template: %w", err)
	}

	return buf.String(), nil
}
