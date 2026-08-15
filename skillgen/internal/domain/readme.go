package domain

// ReadmeHub is one plugin collection's row/section in the generated
// README.md, derived from its already-built hub Skill so the README can
// never drift from what was actually generated.
type ReadmeHub struct {
	Category      string // patterns, enforce, build, secure
	Title         string // Display title, e.g. "Patterns"
	CategoryLabel string // e.g. "Development", "Security", "DevOps"
	Version       string
	TopicCount    int // Number of source docs indexed (len of LibraryFiles)
	Focus         string
	SourceURL     string
	Groups        []TopicGroup
}

// ReadmeData is the top-level data passed to readme.tmpl.
type ReadmeData struct {
	MarketplaceName        string
	MarketplaceOwnerName   string
	MarketplaceDescription string
	Hubs                   []ReadmeHub
}
