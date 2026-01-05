package ports

import "github.com/adaptive-enforcement-lab/claude-skills/skillgen/internal/domain"

// SkillValidator validates generated skills against requirements.
type SkillValidator interface {
	// Validate checks if a skill meets all requirements.
	// Returns validation errors (not fatal - skills can still be written).
	Validate(skill *domain.Skill) []ValidationError
}

// MarketplaceValidator validates the marketplace.json structure.
type MarketplaceValidator interface {
	// Validate checks if the marketplace structure is valid.
	Validate(marketplace *domain.Marketplace) []ValidationError
}

// ValidationError represents a validation issue.
type ValidationError struct {
	Severity Severity // error or warning
	Message  string
	File     string
	Line     int
}

// Severity indicates how serious a validation issue is.
type Severity string

const (
	SeverityError   Severity = "error"
	SeverityWarning Severity = "warning"
)

// Error implements the error interface.
func (v ValidationError) Error() string {
	return v.Message
}
