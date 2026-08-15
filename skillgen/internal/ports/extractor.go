package ports

import "github.com/adaptive-enforcement-lab/claude-skills/skillgen/internal/domain"

// TopicExtractor turns a single parsed document into a lightweight Topic
// link-out entry (title, one-line description, URL). It does not extract
// prose content: fan-out links are the point, not duplication.
type TopicExtractor interface {
	// Extract derives a Topic from a document's frontmatter and path.
	Extract(doc *domain.Document) (*domain.Topic, error)
}

// HubBuilder aggregates every topic in a category into a single hub Skill:
// a short overview plus a grouped index of links to the upstream docs.
type HubBuilder interface {
	// Build assembles the hub skill for a category from its documents.
	Build(category string, docs []*domain.Document, pluginCfg domain.PluginConfig) (*domain.Skill, error)
}
