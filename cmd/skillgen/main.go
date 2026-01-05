package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/adaptive-enforcement-lab/claude-skills/internal/adapters/filesystem"
	"github.com/adaptive-enforcement-lab/claude-skills/internal/adapters/logger"
	"github.com/adaptive-enforcement-lab/claude-skills/internal/adapters/parser"
	"github.com/adaptive-enforcement-lab/claude-skills/internal/ports"
	"github.com/adaptive-enforcement-lab/claude-skills/internal/services/extractor"
	"github.com/adaptive-enforcement-lab/claude-skills/internal/services/generator"
)

var version = "dev"

func main() {
	var (
		sourcePath      string
		outputPath      string
		marketplacePath string
		templatesPath   string
		verbose         bool
		showVersion     bool
	)

	flag.StringVar(&sourcePath, "source", "", "Path to AEL documentation source (required)")
	flag.StringVar(&outputPath, "output", "./skills", "Path to output generated skills")
	flag.StringVar(&marketplacePath, "marketplace", "./.claude-plugin/marketplace.json", "Path to marketplace.json")
	flag.StringVar(&templatesPath, "templates", "./templates", "Path to template directory")
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
	logger.Info("marketplace", marketplacePath)

	// Initialize filesystem
	fs := filesystem.NewFileSystem()

	// Initialize parsers
	frontmatterParser := parser.NewFrontmatterParser()
	sectionParser := parser.NewSectionParser()
	contentExtractor := parser.NewContentExtractor()
	admonitionConverter := parser.NewAdmonitionConverter()

	// Initialize services
	sectionMapper := extractor.NewSectionMapper()
	nameDeriver := extractor.NewNameDeriver()
	skillExtractor := extractor.NewSkillExtractor(sectionMapper, nameDeriver, admonitionConverter)

	// Initialize template renderer
	templateRenderer, err := generator.NewTemplateRenderer(templatesPath)
	if err != nil {
		log.Fatalf("Failed to load templates: %v", err)
	}

	// Initialize document reader
	categories := []string{"patterns", "enforce", "build", "secure"}
	documentReader := filesystem.NewDocumentReader(fs, frontmatterParser, sectionParser, contentExtractor, categories)

	// Initialize writers
	skillWriter := filesystem.NewSkillWriter(fs, templateRenderer)
	marketplaceWriter := filesystem.NewMarketplaceWriter(fs)

	// Find all index.md files
	logger.Info("discovering index.md files")
	indexFiles, err := documentReader.ListIndexFiles(sourcePath, categories)
	if err != nil {
		log.Fatalf("Failed to discover index.md files: %v", err)
	}

	logger.Info("discovered files", "count", len(indexFiles))

	var (
		processed int
		skipped   int
		errors    int
	)

	// Process each file
	for _, filePath := range indexFiles {
		// Read and parse document
		doc, err := documentReader.ReadDocument(filePath)
		if err != nil {
			logger.Error("failed to read document", "path", filePath, "error", err)
			errors++
			continue
		}

		// Skip blog posts
		if doc.Frontmatter.IsBlogPost() {
			logger.Debug("skipping blog post", "path", filePath)
			skipped++
			continue
		}

		// Extract skill components
		skill, err := skillExtractor.Extract(doc)
		if err != nil {
			logger.Error("failed to extract skill", "path", filePath, "error", err)
			errors++
			continue
		}

		// Write skill to filesystem
		if err := skillWriter.WriteSkill(skill, outputPath); err != nil {
			logger.Error("failed to write skill", "name", skill.Metadata.Name, "error", err)
			errors++
			continue
		}

		logger.Info("generated skill", "name", skill.Metadata.Name, "category", skill.Metadata.Category)
		processed++
	}

	// Add secure plugin to marketplace
	logger.Info("updating marketplace.json")
	added, err := marketplaceWriter.AddSecurePlugin(marketplacePath)
	if err != nil {
		logger.Error("failed to add secure plugin", "error", err)
	} else if added {
		logger.Info("added secure plugin to marketplace")
	} else {
		logger.Debug("secure plugin already exists in marketplace")
	}

	// Summary
	fmt.Println("\n=== Generation Summary ===")
	fmt.Printf("Processed: %d\n", processed)
	fmt.Printf("Skipped:   %d (blog posts)\n", skipped)
	fmt.Printf("Errors:    %d\n", errors)
	fmt.Printf("Output:    %s\n", outputPath)

	// Exit 0 even with errors - most errors are expected (missing titles, etc.)
	// Errors are logged for visibility but don't fail the build
}
