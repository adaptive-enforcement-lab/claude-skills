package validator

import (
	"strings"
	"testing"

	"github.com/adaptive-enforcement-lab/claude-skills/skillgen/internal/domain"
	"github.com/adaptive-enforcement-lab/claude-skills/skillgen/internal/ports"
)

func validSkill() *domain.Skill {
	return &domain.Skill{
		Metadata: domain.SkillMetadata{
			Name:        "patterns",
			Title:       "Patterns",
			Description: "Build automation that survives reruns.",
			Category:    "patterns",
			SourceURL:   "https://adaptive-enforcement-lab.com/patterns/",
			Overview:    "Battle-tested automation patterns for GitHub Actions and Kubernetes.",
		},
	}
}

func messages(errs []ports.ValidationError, sev ports.Severity) []string {
	var out []string
	for _, e := range errs {
		if e.Severity == sev {
			out = append(out, e.Message)
		}
	}
	return out
}

func TestValidateAcceptsWellFormedSkill(t *testing.T) {
	if errs := NewSkillValidator().Validate(validSkill()); len(errs) != 0 {
		t.Errorf("expected no findings, got %d: %v", len(errs), errs)
	}
}

func TestValidateRequiredFields(t *testing.T) {
	tests := []struct {
		name    string
		mutate  func(*domain.Skill)
		wantSub string
	}{
		{"empty name", func(s *domain.Skill) { s.Metadata.Name = "" }, "name is required"},
		{"empty description", func(s *domain.Skill) { s.Metadata.Description = "" }, "description is required"},
		{"unknown category", func(s *domain.Skill) { s.Metadata.Category = "wat" }, "unknown category"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			skill := validSkill()
			tt.mutate(skill)

			got := messages(NewSkillValidator().Validate(skill), ports.SeverityError)
			for _, m := range got {
				if strings.Contains(m, tt.wantSub) {
					return
				}
			}
			t.Errorf("expected an error containing %q, got %v", tt.wantSub, got)
		})
	}
}

func TestValidateNameFormat(t *testing.T) {
	bad := []string{"Idempotency", "idem potency", "idem_potency", "-idempotency", "idempotency-", "idem--potency"}

	for _, name := range bad {
		t.Run(name, func(t *testing.T) {
			skill := validSkill()
			skill.Metadata.Name = name

			if len(messages(NewSkillValidator().Validate(skill), ports.SeverityError)) == 0 {
				t.Errorf("expected name %q to be rejected", name)
			}
		})
	}
}

func TestValidateLengthLimits(t *testing.T) {
	t.Run("description over limit is an error", func(t *testing.T) {
		skill := validSkill()
		skill.Metadata.Description = strings.Repeat("a", MaxDescriptionLength+1)

		if len(messages(NewSkillValidator().Validate(skill), ports.SeverityError)) == 0 {
			t.Error("expected over-length description to be rejected")
		}
	})

	t.Run("name over limit is an error", func(t *testing.T) {
		skill := validSkill()
		skill.Metadata.Name = strings.Repeat("a", MaxNameLength+1)

		if len(messages(NewSkillValidator().Validate(skill), ports.SeverityError)) == 0 {
			t.Error("expected over-length name to be rejected")
		}
	})

	t.Run("very short description warns", func(t *testing.T) {
		skill := validSkill()
		skill.Metadata.Description = "Short."

		if len(messages(NewSkillValidator().Validate(skill), ports.SeverityWarning)) == 0 {
			t.Error("expected a warning for a description too short to route on")
		}
	})
}

func TestValidateWarnsWhenNoBodyExtracted(t *testing.T) {
	skill := validSkill()
	skill.Metadata.Overview = ""

	if len(messages(NewSkillValidator().Validate(skill), ports.SeverityWarning)) == 0 {
		t.Error("expected a warning when neither an overview nor topic groups were extracted")
	}
}

func TestValidateAcceptsGroupsWithoutOverview(t *testing.T) {
	// MainContent is empty until the template renderer runs, so an extracted
	// body must be recognised from either the overview or the topic groups.
	skill := validSkill()
	skill.Metadata.Overview = ""
	skill.Groups = []domain.TopicGroup{{Title: "Architecture", Topics: []domain.Topic{{Title: "Hub and Spoke"}}}}

	if len(NewSkillValidator().Validate(skill)) != 0 {
		t.Error("topic groups alone should count as an extracted body")
	}
}

func TestValidateWarnsOnMissingSourceURL(t *testing.T) {
	skill := validSkill()
	skill.Metadata.SourceURL = ""

	if len(messages(NewSkillValidator().Validate(skill), ports.SeverityWarning)) == 0 {
		t.Error("expected a warning for missing source URL")
	}
}

func TestValidateReportsFileContext(t *testing.T) {
	skill := validSkill()
	skill.Metadata.Description = ""

	errs := NewSkillValidator().Validate(skill)
	if len(errs) == 0 {
		t.Fatal("expected findings")
	}
	for _, e := range errs {
		if e.File == "" {
			t.Errorf("finding %q has no File context", e.Message)
		}
	}
}

func TestSatisfiesPortInterface(t *testing.T) {
	var _ ports.SkillValidator = NewSkillValidator()
}
