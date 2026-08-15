package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/adaptive-enforcement-lab/claude-skills/skillgen/internal/adapters/filesystem"
	"github.com/adaptive-enforcement-lab/claude-skills/skillgen/internal/adapters/logger"
	"github.com/adaptive-enforcement-lab/claude-skills/skillgen/internal/adapters/parser"
	"github.com/adaptive-enforcement-lab/claude-skills/skillgen/internal/domain"
	"github.com/adaptive-enforcement-lab/claude-skills/skillgen/internal/ports"
	"github.com/adaptive-enforcement-lab/claude-skills/skillgen/internal/services"
	"github.com/adaptive-enforcement-lab/claude-skills/skillgen/internal/services/extractor"
	"github.com/adaptive-enforcement-lab/claude-skills/skillgen/internal/services/generator"
	"github.com/adaptive-enforcement-lab/claude-skills/skillgen/internal/services/validator"
)

var version = "dev"

func main() {
	var (
		sourcePath          string
		outputPath          string
		marketplacePath     string
		readmePath          string
		templatesPath       string
		pluginMetadataPath  string
		releaseManifestPath string
		verbose             bool
		showVersion         bool
	)

	flag.StringVar(&sourcePath, "source", "", "Path to AEL documentation source (required)")
	flag.StringVar(&outputPath, "output", "./plugins", "Path to output generated plugins")
	flag.StringVar(&marketplacePath, "marketplace", "./.claude-plugin/marketplace.json", "Path to marketplace.json")
	flag.StringVar(&readmePath, "readme", "./README.md", "Path to generated README.md")
	flag.StringVar(&templatesPath, "templates", "./templates", "Path to template directory")
	flag.StringVar(&pluginMetadataPath, "plugin-metadata", "./plugin-metadata.json", "Path to plugin metadata config")
	flag.StringVar(&releaseManifestPath, "release-manifest", "./.release-please-manifest.json", "Path to release-please manifest")
	flag.BoolVar(&verbose, "verbose", false, "Enable verbose logging")
	flag.BoolVar(&showVersion, "version", false, "Show version and exit")
	flag.Parse()

	if showVersion {
		fmt.Printf("skillgen version %s\n", version)
		return
	}

	if sourcePath == "" {
		log.Fatal("--source flag is required")
	}

	// Initialize logger
	logLevel := ports.LogLevelInfo
	if verbose {
		logLevel = ports.LogLevelDebug
	}
	logger := logger.NewLogger(logLevel)

	logger.Info("AEL Claude Skills Generator")
	logger.Info("source", sourcePath)
	logger.Info("output", outputPath)
	logger.Info("plugin-metadata", pluginMetadataPath)
	logger.Info("release-manifest", releaseManifestPath)

	// Initialize filesystem
	fs := filesystem.NewFileSystem()

	// Initialize parsers
	frontmatterParser := parser.NewFrontmatterParser()
	sectionParser := parser.NewSectionParser()
	contentExtractor := parser.NewContentExtractor()
	admonitionConverter := parser.NewAdmonitionConverter()

	// Initialize services
	topicExtractor := extractor.NewTopicExtractor()
	hubBuilder := extractor.NewHubBuilder(topicExtractor, admonitionConverter)

	// Initialize template renderer
	templateRenderer, err := generator.NewTemplateRenderer(templatesPath)
	if err != nil {
		log.Fatalf("Failed to load templates: %v", err)
	}

	// Initialize validator
	skillValidator := validator.NewSkillValidator()

	// Initialize document reader
	categories := domain.Categories
	documentReader := filesystem.NewDocumentReader(fs, frontmatterParser, sectionParser, contentExtractor, categories)

	// Initialize writers
	skillWriter := filesystem.NewSkillWriter(fs, templateRenderer)
	marketplaceWriter := filesystem.NewMarketplaceWriter(fs)
	configReader := filesystem.NewConfigReader(fs)

	// Plugin metadata is the source of truth for each hub's curated
	// description and tags.
	pluginMetadata, err := configReader.ReadPluginMetadata(pluginMetadataPath)
	if err != nil {
		log.Fatalf("Failed to read plugin metadata: %v", err)
	}

	var (
		topics    int
		hubCount  int
		errors    int
		warned    int
		builtHubs []*domain.Skill
	)

	// Build one hub skill per category.
	for _, category := range categories {
		logger.Info("discovering index.md files", "category", category)
		indexFiles, err := documentReader.ListIndexFiles(sourcePath, []string{category})
		if err != nil {
			logger.Error("failed to discover index.md files", "category", category, "error", err)
			errors++
			continue
		}

		var docs []*domain.Document
		for _, filePath := range indexFiles {
			doc, err := documentReader.ReadDocument(filePath)
			if err != nil {
				logger.Error("failed to read document", "path", filePath, "error", err)
				errors++
				continue
			}

			if doc.Frontmatter.IsBlogPost() {
				logger.Debug("skipping blog post", "path", filePath)
				continue
			}

			docs = append(docs, doc)
			topics++
		}

		pluginCfg, ok := pluginMetadata.Plugins[category]
		if !ok {
			logger.Error("no plugin-metadata.json entry for category", "category", category)
			errors++
			continue
		}

		hub, err := hubBuilder.Build(category, docs, pluginCfg)
		if err != nil {
			logger.Error("failed to build hub skill", "category", category, "error", err)
			errors++
			continue
		}

		// Validate the skill. Findings are advisory: a skill that fails
		// validation is still written, but is surfaced so it can be fixed
		// at the source document.
		if findings := skillValidator.Validate(hub); len(findings) > 0 {
			var hasError bool
			for _, f := range findings {
				if f.Severity == ports.SeverityError {
					hasError = true
					logger.Error("skill validation", "name", hub.Metadata.Name, "issue", f.Message)
					continue
				}
				logger.Warn("skill validation", "name", hub.Metadata.Name, "issue", f.Message)
				warned++
			}
			if hasError {
				errors++
			}
		}

		if err := skillWriter.WriteSkill(hub, outputPath); err != nil {
			logger.Error("failed to write hub skill", "category", category, "error", err)
			errors++
			continue
		}

		logger.Info("generated hub skill", "category", category, "groups", len(hub.Groups))
		builtHubs = append(builtHubs, hub)
		hubCount++
	}

	// Generate marketplace files
	logger.Info("generating marketplace files")
	marketplaceGen := services.NewMarketplaceGenerator(configReader, marketplaceWriter, logger)
	err = marketplaceGen.Generate(pluginMetadataPath, releaseManifestPath, outputPath, marketplacePath)
	if err != nil {
		logger.Error("failed to generate marketplace files", "error", err)
		errors++
	} else {
		logger.Info("marketplace files generated successfully")
	}

	// Generate README.md from the same hubs and metadata just used above,
	// so it can never drift from what was actually generated.
	logger.Info("generating README.md")
	versions, err := configReader.ReadReleaseManifest(releaseManifestPath)
	if err != nil {
		logger.Error("failed to read release manifest for README", "error", err)
		errors++
	} else {
		readmeGen := services.NewReadmeGenerator(templateRenderer, fs, logger)
		if err := readmeGen.Generate(builtHubs, pluginMetadata, versions, readmePath); err != nil {
			logger.Error("failed to generate README.md", "error", err)
			errors++
		} else {
			logger.Info("README.md generated successfully")
		}
	}

	// Summary
	fmt.Println("\n=== Generation Summary ===")
	fmt.Printf("Categories:     %d\n", len(categories))
	fmt.Printf("Topics indexed: %d\n", topics)
	fmt.Printf("Hub skills:     %d\n", hubCount)
	fmt.Printf("Warnings:       %d\n", warned)
	fmt.Printf("Errors:         %d\n", errors)
	fmt.Printf("Output:         %s\n", outputPath)

	if errors > 0 {
		logger.Info("completed with errors", "count", errors)
	} else {
		logger.Info("completed successfully")
	}

	// Exit 0 even with errors - most errors are expected (missing titles, etc.)
	// Errors are logged for visibility but don't fail the build
}
