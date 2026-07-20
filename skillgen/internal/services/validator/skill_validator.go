// Package validator checks generated skills against Claude Code's skill
// requirements before they are written to disk.
package validator

import (
	"fmt"
	"path/filepath"
	"regexp"

	"github.com/adaptive-enforcement-lab/claude-skills/skillgen/internal/domain"
	"github.com/adaptive-enforcement-lab/claude-skills/skillgen/internal/ports"
)

const (
	// MaxNameLength is the longest a skill name may be.
	MaxNameLength = 64

	// MaxDescriptionLength is the longest a skill description may be.
	// The description is the only field Claude sees when deciding whether a
	// skill is relevant, so an over-length value risks truncation at the
	// point where it matters most.
	MaxDescriptionLength = 1024

	// MinDescriptionLength is the shortest description considered useful for
	// routing. Below this, a description rarely carries enough signal to
	// distinguish one skill from another.
	MinDescriptionLength = 20
)

// namePattern matches lowercase kebab-case: alphanumeric segments joined by
// single hyphens, with no leading, trailing, or doubled hyphens.
var namePattern = regexp.MustCompile(`^[a-z0-9]+(-[a-z0-9]+)*$`)

// SkillValidator implements ports.SkillValidator.
type SkillValidator struct{}

// NewSkillValidator creates a new skill validator.
func NewSkillValidator() *SkillValidator {
	return &SkillValidator{}
}

// Validate checks a skill against Claude Code's requirements.
// Findings are advisory: errors indicate a skill Claude may reject or fail to
// route to, warnings indicate degraded quality. Neither blocks the write.
func (v *SkillValidator) Validate(skill *domain.Skill) []ports.ValidationError {
	if skill == nil {
		return []ports.ValidationError{{
			Severity: ports.SeverityError,
			Message:  "skill is nil",
		}}
	}

	file := skillFile(skill)
	var findings []ports.ValidationError

	add := func(sev ports.Severity, format string, args ...any) {
		findings = append(findings, ports.ValidationError{
			Severity: sev,
			Message:  fmt.Sprintf(format, args...),
			File:     file,
		})
	}

	findings = append(findings, v.validateName(skill.Metadata.Name, file)...)

	switch desc := skill.Metadata.Description; {
	case desc == "":
		add(ports.SeverityError, "description is required: Claude uses it to decide when the skill applies")
	case len(desc) > MaxDescriptionLength:
		add(ports.SeverityError, "description is %d characters, exceeding the %d limit", len(desc), MaxDescriptionLength)
	case len(desc) < MinDescriptionLength:
		add(ports.SeverityWarning, "description is only %d characters and may be too vague to route on", len(desc))
	}

	// MainContent is deliberately empty until the template renderer runs, so
	// body quality is judged from the extracted sections instead.
	if !hasExtractedBody(&skill.Metadata) {
		add(ports.SeverityWarning, "no section matched a known component; the skill body will be near-empty")
	}

	if !domain.IsCategory(skill.Metadata.Category) {
		add(ports.SeverityError, "unknown category %q", skill.Metadata.Category)
	}

	if skill.Metadata.SourceURL == "" {
		add(ports.SeverityWarning, "no source URL: the skill cannot link back to its documentation")
	}

	return findings
}

// validateName checks the skill name is present, correctly formatted, and
// within length limits. The name becomes the skill's directory, so an invalid
// value produces a skill Claude cannot load.
func (v *SkillValidator) validateName(name, file string) []ports.ValidationError {
	switch {
	case name == "":
		return []ports.ValidationError{{
			Severity: ports.SeverityError,
			Message:  "name is required",
			File:     file,
		}}
	case len(name) > MaxNameLength:
		return []ports.ValidationError{{
			Severity: ports.SeverityError,
			Message:  fmt.Sprintf("name is %d characters, exceeding the %d limit", len(name), MaxNameLength),
			File:     file,
		}}
	case !namePattern.MatchString(name):
		return []ports.ValidationError{{
			Severity: ports.SeverityError,
			Message:  fmt.Sprintf("name %q is not lowercase kebab-case", name),
			File:     file,
		}}
	}

	return nil
}

// hasExtractedBody reports whether section mapping produced any substantive
// content. A skill with none renders to little more than its frontmatter.
func hasExtractedBody(m *domain.SkillMetadata) bool {
	return m.WhenToUse != "" ||
		m.Prerequisites != "" ||
		m.ImplementationSteps != "" ||
		m.KeyPrinciples != "" ||
		m.WhenToApply != "" ||
		m.Comparison != "" ||
		m.AntiPatterns != "" ||
		len(m.Techniques) > 0
}

// skillFile returns the best available identifier for error reporting,
// preferring the source document the skill was generated from.
func skillFile(skill *domain.Skill) string {
	if path := skill.Metadata.SourcePath; path != "" {
		return path
	}
	if name := skill.Metadata.Name; name != "" {
		return filepath.Join(skill.Metadata.Category, "skills", name, "SKILL.md")
	}
	return "<unknown skill>"
}
