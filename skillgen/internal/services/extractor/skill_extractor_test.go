package extractor

import "testing"

func TestBuildSourceURL(t *testing.T) {
	e := &SkillExtractor{}

	tests := []struct {
		name     string
		path     string
		category string
		want     string
	}{
		{
			name:     "nested topic preserves full path",
			path:     "ael-docs/docs/patterns/efficiency/idempotency/index.md",
			category: "patterns",
			want:     "https://adaptive-enforcement-lab.com/patterns/efficiency/idempotency/",
		},
		{
			name:     "deeply nested topic preserves every segment",
			path:     "ael-docs/docs/secure/github-actions-security/workflows/environments/index.md",
			category: "secure",
			want:     "https://adaptive-enforcement-lab.com/secure/github-actions-security/workflows/environments/",
		},
		{
			name:     "single level topic",
			path:     "ael-docs/docs/build/packaging/index.md",
			category: "build",
			want:     "https://adaptive-enforcement-lab.com/build/packaging/",
		},
		{
			name:     "category index itself",
			path:     "ael-docs/docs/enforce/index.md",
			category: "enforce",
			want:     "https://adaptive-enforcement-lab.com/enforce/",
		},
		{
			name:     "absolute path",
			path:     "/home/runner/docs/patterns/reliability/chaos-engineering/index.md",
			category: "patterns",
			want:     "https://adaptive-enforcement-lab.com/patterns/reliability/chaos-engineering/",
		},
		{
			name:     "category absent from path falls back to base URL",
			path:     "ael-docs/docs/glossary/index.md",
			category: "patterns",
			want:     "https://adaptive-enforcement-lab.com",
		},
		{
			name:     "repeated segment name does not truncate early",
			path:     "ael-docs/docs/build/build/index.md",
			category: "build",
			want:     "https://adaptive-enforcement-lab.com/build/build/",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := e.buildSourceURL(tt.path, tt.category); got != tt.want {
				t.Errorf("buildSourceURL(%q, %q)\n got: %s\nwant: %s", tt.path, tt.category, got, tt.want)
			}
		})
	}
}

func TestDetermineCategoryFromPath(t *testing.T) {
	e := &SkillExtractor{}

	tests := []struct {
		path string
		want string
	}{
		{"ael-docs/docs/patterns/efficiency/idempotency/index.md", "patterns"},
		{"ael-docs/docs/secure/cloud-native/gke-hardening/index.md", "secure"},
		{"ael-docs/docs/enforce/index.md", "enforce"},
		{"ael-docs/docs/build/packaging/index.md", "build"},
		{"ael-docs/docs/glossary/index.md", ""},
	}

	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			if got := e.determineCategoryFromPath(tt.path); got != tt.want {
				t.Errorf("determineCategoryFromPath(%q) = %q, want %q", tt.path, got, tt.want)
			}
		})
	}
}
