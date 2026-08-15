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

// RenderReference renders the reference.md file.
func (r *TemplateRenderer) RenderReference(skill *domain.Skill) (string, error) {
	var buf bytes.Buffer

	if err := r.templates.ExecuteTemplate(&buf, "reference.tmpl", skill); err != nil {
		return "", fmt.Errorf("failed to render reference template: %w", err)
	}

	return buf.String(), nil
}

// RenderReadme renders the repo root README.md.
func (r *TemplateRenderer) RenderReadme(data *domain.ReadmeData) (string, error) {
	var buf bytes.Buffer

	if err := r.templates.ExecuteTemplate(&buf, "readme.tmpl", data); err != nil {
		return "", fmt.Errorf("failed to render readme template: %w", err)
	}

	return buf.String(), nil
}
