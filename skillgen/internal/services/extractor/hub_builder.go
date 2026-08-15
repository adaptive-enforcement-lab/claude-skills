package extractor

import (
	"fmt"
	"sort"
	"strings"

	"github.com/adaptive-enforcement-lab/claude-skills/skillgen/internal/domain"
	"github.com/adaptive-enforcement-lab/claude-skills/skillgen/internal/ports"
)

// HubBuilder implements ports.HubBuilder. It aggregates every document in a
// category into one hub Skill: a short overview from the category root doc,
// plus topics grouped by their first path segment under the category, each
// fanning out to the upstream documentation instead of duplicating it. It
// also assembles each doc's full body once for the hub's reference.md, so
// depth is available offline without re-fetching the live docs.
type HubBuilder struct {
	topicExtractor      ports.TopicExtractor
	admonitionConverter ports.AdmonitionConverter
}

// NewHubBuilder creates a new hub builder.
func NewHubBuilder(topicExtractor ports.TopicExtractor, admonitionConverter ports.AdmonitionConverter) *HubBuilder {
	return &HubBuilder{topicExtractor: topicExtractor, admonitionConverter: admonitionConverter}
}

// referenceShift levels: a topic's body is wrapped under "### Title" in
// reference.md, the category root's under "## Overview".
const (
	topicReferenceShift = 2
	rootReferenceShift  = 1
)

// Build assembles the hub skill for a category from its documents.
func (b *HubBuilder) Build(category string, docs []*domain.Document, pluginCfg domain.PluginConfig) (*domain.Skill, error) {
	var rootDoc *domain.Document
	groupRoots := make(map[string]*domain.Document)
	var rest []*domain.Document

	for _, doc := range docs {
		segments := categorySegments(doc.Path, category)
		switch len(segments) {
		case 0:
			rootDoc = doc
		case 1:
			groupRoots[segments[0]] = doc
			rest = append(rest, doc)
		default:
			rest = append(rest, doc)
		}
	}

	if rootDoc == nil {
		return nil, fmt.Errorf("no category root index.md found for %q", category)
	}

	groups := make(map[string]*domain.TopicGroup)
	for _, doc := range rest {
		segments := categorySegments(doc.Path, category)
		groupKey := segments[0]

		group, ok := groups[groupKey]
		if !ok {
			group = &domain.TopicGroup{Title: humanize(groupKey)}
			groups[groupKey] = group
		}

		if groupRootDoc, isRoot := groupRoots[groupKey]; isRoot && doc.Path == groupRootDoc.Path {
			group.Title = doc.Frontmatter.Title
			group.Description = firstSentence(doc.Frontmatter.Description)
			group.URL = buildSourceURL(doc.Path, category)
			// A group's own body sits alongside its child topics under the
			// same "## Group" heading, so it must shift by the same amount
			// as a topic body — otherwise its internal headings can collide
			// with a sibling topic's "### Title" wrapper level.
			group.ReferenceBody = prepareReferenceBody(b.admonitionConverter.Convert(doc.RawContent), topicReferenceShift)
			continue
		}

		topic, err := b.topicExtractor.Extract(doc)
		if err != nil {
			return nil, err
		}
		topic.ReferenceBody = prepareReferenceBody(b.admonitionConverter.Convert(doc.RawContent), topicReferenceShift)
		group.Topics = append(group.Topics, *topic)
	}

	sortedGroups := make([]domain.TopicGroup, 0, len(groups))
	for _, group := range groups {
		sort.Slice(group.Topics, func(i, j int) bool {
			return group.Topics[i].Title < group.Topics[j].Title
		})
		sortedGroups = append(sortedGroups, *group)
	}
	sort.Slice(sortedGroups, func(i, j int) bool {
		return sortedGroups[i].Title < sortedGroups[j].Title
	})

	metadata := domain.SkillMetadata{
		Name:          category,
		Title:         rootDoc.Frontmatter.Title,
		Description:   pluginCfg.Description,
		Category:      category,
		Tags:          pluginCfg.Tags,
		Overview:      firstSentences(rootDoc.Introduction, 3),
		ReferenceBody: prepareReferenceBody(b.admonitionConverter.Convert(rootDoc.RawContent), rootReferenceShift),
		SourcePath:    rootDoc.Path,
		SourceURL:     buildSourceURL(rootDoc.Path, category),
	}

	return &domain.Skill{
		Metadata: metadata,
		Groups:   sortedGroups,
	}, nil
}

// humanize turns a URL slug into a title-cased heading, e.g.
// "github-actions" -> "Github Actions". Used only as a fallback when a
// group has no index.md of its own to supply a proper title.
func humanize(slug string) string {
	words := strings.Split(slug, "-")
	for i, w := range words {
		if w == "" {
			continue
		}
		words[i] = strings.ToUpper(w[:1]) + w[1:]
	}
	return strings.Join(words, " ")
}

// firstSentences returns the first n sentences of the first paragraph of
// text, trimmed and joined back into a single line. Only the first
// paragraph is considered: admonitions and other blocks that follow a blank
// line in a doc's introduction aren't prose meant for a short overview.
func firstSentences(text string, n int) string {
	trimmed := strings.TrimSpace(text)
	if trimmed == "" {
		return ""
	}

	if idx := strings.Index(trimmed, "\n\n"); idx != -1 {
		trimmed = trimmed[:idx]
	}

	var b strings.Builder
	count := 0
	start := 0
	for i, r := range trimmed {
		if r == '.' || r == '!' || r == '?' {
			b.WriteString(trimmed[start : i+1])
			count++
			start = i + 1
			if count >= n {
				break
			}
		}
	}
	if count == 0 {
		return strings.Join(strings.Fields(trimmed), " ")
	}

	return strings.Join(strings.Fields(b.String()), " ")
}
