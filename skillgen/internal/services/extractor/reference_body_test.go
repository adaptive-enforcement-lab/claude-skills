package extractor

import "testing"

func TestPrepareReferenceBodyDropsLeadingTitle(t *testing.T) {
	body := "# Hub and Spoke\n\nCentral coordinator, many workers.\n\n## Overview\n\nMore detail."
	got := prepareReferenceBody(body, 2)

	want := "Central coordinator, many workers.\n\n#### Overview\n\nMore detail."
	if got != want {
		t.Errorf("got:\n%s\nwant:\n%s", got, want)
	}
}

func TestPrepareReferenceBodyShiftsHeadingsByLevel(t *testing.T) {
	tests := []struct {
		name  string
		shift int
		body  string
		want  string
	}{
		{"shift 1", 1, "# Title\n\n## Section\n\n### Subsection", "### Section\n\n#### Subsection"},
		{"shift 2", 2, "# Title\n\n## Section\n\n### Subsection", "#### Section\n\n##### Subsection"},
		{"no heading beyond title", 2, "# Title\n\nJust a paragraph.", "Just a paragraph."},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := prepareReferenceBody(tt.body, tt.shift)
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}

func TestPrepareReferenceBodyOnlyDropsFirstH1(t *testing.T) {
	// A second H1 later in the body (unusual, but shouldn't be treated as
	// "the title" again) is shifted like any other heading.
	body := "# Title\n\nIntro.\n\n# Another H1\n\nMore."
	got := prepareReferenceBody(body, 1)

	want := "Intro.\n\n## Another H1\n\nMore."
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestPrepareReferenceBodyHandlesEmptyInput(t *testing.T) {
	if got := prepareReferenceBody("", 2); got != "" {
		t.Errorf("got %q, want empty string", got)
	}
}
