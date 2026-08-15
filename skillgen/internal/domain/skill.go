package domain

// Skill represents a single hub skill for a plugin collection: a short
// overview plus a linked index of every topic in that collection. There is
// exactly one Skill per category (patterns, enforce, build, secure).
type Skill struct {
	Metadata     SkillMetadata
	Groups       []TopicGroup
	LibraryFiles []LibraryFile // Every source doc, verbatim, mirroring the docs tree
	MainContent  string        // SKILL.md content (required)
}

// SkillMetadata contains the frontmatter and derived metadata for a hub skill.
type SkillMetadata struct {
	Name          string // Kebab-case name; equal to Category for a hub skill
	Title         string // Display title
	Description   string // Short, curated description from plugin-metadata.json
	Category      string // patterns, enforce, build, secure
	Tags          []string
	Overview      string // Short intro paragraph, from the category root doc
	ReferenceBody string // Full cleaned body of the category root doc, for reference.md
	SourcePath    string // Original document path (category root index.md)
	SourceURL     string // URL to the category root on adaptive-enforcement-lab.com
}

// TopicGroup is a themed cluster of topics within a hub skill (e.g. the
// "Architecture Patterns" group within the "patterns" hub).
type TopicGroup struct {
	Title         string // Group heading
	Description   string // One-line group blurb
	URL           string // Upstream URL to the group's own section page, if any
	ReferenceBody string // Full cleaned body of the group's own doc, if any
	Topics        []Topic
}

// Topic is a single index entry: a title, one-line description, and a link
// to the bundled library/ copy of the doc (SKILL.md links here, not to the
// live site, so the skill is fully usable offline — the library file itself
// carries the upstream URL as its "Source:" line). URL is kept for the
// upstream link but is not rendered in SKILL.md.
type Topic struct {
	Title         string
	Description   string
	URL           string
	LibraryPath   string // Path to the library/ file, relative to SKILL.md
	ReferenceBody string // Full cleaned body of the topic's doc, for reference.md
}

// LibraryFile is a single source doc shipped verbatim (title, source URL
// note, then its full original content — nothing stripped, nothing
// shifted) under a hub's library/ directory, mirroring the doc's path
// under the category. This is the complete unmerged source library,
// shipped in addition to the curated reference.md.
type LibraryFile struct {
	RelPath string // Path relative to the hub's library/ directory
	Content string
}

// Note: CodeBlock, Table, and MermaidDiagram are defined in document.go
// and can be used directly since they're in the same package.
