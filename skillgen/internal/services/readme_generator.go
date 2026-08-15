package services

import (
	"fmt"
	"sort"
	"strings"

	"github.com/adaptive-enforcement-lab/claude-skills/skillgen/internal/domain"
	"github.com/adaptive-enforcement-lab/claude-skills/skillgen/internal/ports"
)

// ReadmeGenerator renders README.md from the hub skills already built this
// run, plus the plugin metadata and versions driving them. Because it reads
// the same data the plugins themselves are generated from, README.md cannot
// go stale relative to what was actually generated the way a hand-maintained
// README did.
type ReadmeGenerator struct {
	renderer ports.TemplateRenderer
	fs       ports.FileSystem
	logger   ports.Logger
}

// NewReadmeGenerator creates a new README generator service.
func NewReadmeGenerator(renderer ports.TemplateRenderer, fs ports.FileSystem, logger ports.Logger) *ReadmeGenerator {
	return &ReadmeGenerator{renderer: renderer, fs: fs, logger: logger}
}

// Generate builds README.md from the given hub skills, plugin metadata, and
// release versions, and writes it to outputPath.
func (g *ReadmeGenerator) Generate(
	hubs []*domain.Skill,
	metadata *domain.PluginMetadata,
	versions map[string]string,
	outputPath string,
) error {
	readmeHubs := make([]domain.ReadmeHub, 0, len(hubs))
	for _, hub := range hubs {
		cfg, ok := metadata.Plugins[hub.Metadata.Category]
		if !ok {
			return fmt.Errorf("no plugin-metadata.json entry for category %q", hub.Metadata.Category)
		}

		manifestKey := fmt.Sprintf("plugins/%s", hub.Metadata.Category)
		version := versions[manifestKey]
		if version == "" {
			version = "0.0.0"
		}

		readmeHubs = append(readmeHubs, domain.ReadmeHub{
			Category:      hub.Metadata.Category,
			Title:         hub.Metadata.Title,
			CategoryLabel: titleCategory(cfg.Category),
			Version:       version,
			TopicCount:    len(hub.LibraryFiles),
			Focus:         truncateWords(hub.Metadata.Description, 14),
			SourceURL:     hub.Metadata.SourceURL,
			Groups:        hub.Groups,
		})
	}

	sort.Slice(readmeHubs, func(i, j int) bool { return readmeHubs[i].Category < readmeHubs[j].Category })

	data := &domain.ReadmeData{
		MarketplaceName:        metadata.Marketplace.Name,
		MarketplaceOwnerName:   metadata.Marketplace.Owner.Name,
		MarketplaceDescription: metadata.Marketplace.Description,
		Hubs:                   readmeHubs,
	}

	content, err := g.renderer.RenderReadme(data)
	if err != nil {
		return fmt.Errorf("failed to render README.md: %w", err)
	}

	if err := g.fs.WriteFile(outputPath, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write README.md: %w", err)
	}

	g.logger.Info("generated README.md", "path", outputPath, "hubs", len(readmeHubs))
	return nil
}

// titleCategory formats a plugin-metadata.json category value for display,
// e.g. "devops" -> "DevOps", "security" -> "Security".
func titleCategory(category string) string {
	if category == "devops" {
		return "DevOps"
	}
	if category == "" {
		return ""
	}
	return strings.ToUpper(category[:1]) + category[1:]
}

// truncateWords caps text to at most maxWords words, appending an ellipsis
// if it had to cut. Used to keep the README's Focus column short even
// though the source description (also used as the SKILL.md/marketplace
// description) is often one long "Use when..." sentence with no earlier
// punctuation to split on.
func truncateWords(text string, maxWords int) string {
	words := strings.Fields(text)
	if len(words) <= maxWords {
		return strings.Join(words, " ")
	}

	return strings.TrimRight(strings.Join(words[:maxWords], " "), ".,;:!?") + "…"
}
