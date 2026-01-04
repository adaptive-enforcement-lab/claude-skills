# Adaptive Enforcement Lab - Claude Code Skills

Claude Code skills marketplace for secure development patterns, enforcement automation, and build engineering.

**Status**: 🚧 Under active development

## Installation

```bash
# Add the AEL skills marketplace
/plugin marketplace add adaptive-enforcement-lab/claude-skills

# Install individual plugin collections
/plugin install patterns@ael-skills
/plugin install enforcement@ael-skills
/plugin install build@ael-skills
```

## Available Skills

### Patterns (Development)
Reusable engineering patterns automatically generated from [AEL documentation](https://adaptive-enforcement-lab.com/patterns/):

- Error handling patterns (fail-fast, circuit breakers, retry logic)
- State management patterns
- Performance optimization patterns
- Resilience and fault tolerance patterns

### Enforcement (Security)
Security and compliance enforcement automation from [AEL enforcement guides](https://adaptive-enforcement-lab.com/enforce/):

- Pre-commit hook setup and configuration
- Policy validation automation
- Security scanning integration
- Compliance checking workflows

### Build (DevOps)
Build engineering patterns from [AEL build guides](https://adaptive-enforcement-lab.com/build/):

- CI/CD pipeline patterns
- Release automation strategies
- Deployment patterns
- Build optimization techniques

## Automated Generation

All skills in this repository are automatically generated from AEL documentation:

- **Source**: [adaptive-enforcement-lab.com](https://github.com/adaptive-enforcement-lab/adaptive-enforcement-lab-com)
- **Generator**: Go-based extraction pipeline
- **CI/CD**: GitHub Actions workflows
- **Sync**: Skills update automatically when documentation changes

## Team Distribution

To auto-register this marketplace for your team, add to `.claude/settings.json` in your project:

```json
{
  "extraKnownMarketplaces": {
    "ael-skills": {
      "source": {
        "source": "github",
        "repo": "adaptive-enforcement-lab/claude-skills"
      }
    }
  },
  "enabledPlugins": {
    "patterns@ael-skills": true,
    "enforcement@ael-skills": true,
    "build@ael-skills": true
  }
}
```

## Repository Structure

```
.claude-plugin/
└── marketplace.json              # Marketplace catalog

skills/                           # Generated skills (DO NOT EDIT)
├── .generated                    # Marker file
├── patterns/                     # Pattern-based skills
├── enforcement/                  # Enforcement automation skills
└── build/                        # Build engineering skills

cmd/skillgen/                     # Main application
internal/                         # Internal packages
├── parser/                       # Content parsers
├── extractor/                    # Component extractors
├── generator/                    # Skill generators
└── validator/                    # Validation logic

templates/                        # Go templates

.github/workflows/
└── generate-skills.yml           # CI automation
```

## Development

```bash
# Build the generator
go build -o bin/skillgen ./cmd/skillgen

# Run generator
./bin/skillgen --source ../adaptive-enforcement-lab-com/docs

# Run tests
go test ./...
```

See [CONTRIBUTING.md](CONTRIBUTING.md) for detailed development guidelines.

## Architecture

This project follows Clean/Hexagonal Architecture:

- **Domain** (`internal/domain`): Core entities and business logic
- **Ports** (`internal/ports`): Interfaces for external dependencies
- **Adapters** (`internal/adapters`): Implementations (filesystem, parsers)
- **Services** (`internal/services`): Application services (extractors, generators)
- **CMD** (`cmd/skillgen`): Entry point and dependency injection

## Releases

Releases are automated using [release-please](https://github.com/googleapis/release-please):

- Conventional commits trigger version bumps
- Changelog is auto-generated
- GitHub releases include pre-built binaries for Linux, macOS, and Windows

## Contributing

Contributions are welcome! See [CONTRIBUTING.md](CONTRIBUTING.md) for:

- Development setup
- Code standards
- Commit message format
- Pull request process

## Security

For security vulnerability reporting, see [SECURITY.md](SECURITY.md).

## Mission

Turn secure development into an enforced standard, not an afterthought.

## Links

- [AEL Documentation](https://adaptive-enforcement-lab.com)
- [GitHub Organization](https://github.com/adaptive-enforcement-lab)
- [Skills Marketplace](.claude-plugin/marketplace.json)

## License

MIT
