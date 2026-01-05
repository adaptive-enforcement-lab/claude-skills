# Contributing

Contributions to Adaptive Enforcement Lab Claude Skills are welcome.

## Quick Start

1. Fork the repository
2. Create a feature branch from `main`
3. Make your changes
4. Submit a pull request

## Development Setup

```bash
# Clone repository
git clone https://github.com/adaptive-enforcement-lab/claude-skills.git
cd claude-skills

# Build generator
cd skillgen && go build -o ../bin/skillgen ./cmd/skillgen && cd ..

# Run tests
cd skillgen && go test ./... && cd ..

# Generate skills from AEL docs
./bin/skillgen \
  --source ../adaptive-enforcement-lab-com/docs \
  --output skills \
  --marketplace .claude-plugin/marketplace.json
```

## Architecture

This project follows Clean/Hexagonal Architecture:

- **Domain** (`skillgen/internal/domain`): Core entities and business logic
- **Ports** (`skillgen/internal/ports`): Interfaces for external dependencies
- **Adapters** (`skillgen/internal/adapters`): Implementations of ports (filesystem, parsers)
- **Services** (`skillgen/internal/services`): Application services (extractors, generators)
- **CMD** (`skillgen/cmd/skillgen`): Entry point and dependency injection

## Code Standards

- Go 1.23+
- Use `gofmt` for formatting
- Follow standard Go conventions
- Keep packages focused and cohesive
- Write tests for business logic

## Commit Messages

Use [Conventional Commits](https://www.conventionalcommits.org/):

- `feat:` New features
- `fix:` Bug fixes
- `docs:` Documentation changes
- `chore:` Maintenance tasks
- `test:` Test additions/changes
- `refactor:` Code refactoring

Examples:
```
feat: add support for code block extraction
fix: handle empty section content gracefully
docs: update README with architecture details
chore: update dependencies to latest versions
```

## Pull Request Process

1. **Branch naming**: `feat/description`, `fix/description`, or `chore/description`
2. **Clear title**: Use conventional commit format
3. **Description**: Explain what and why (not how)
4. **Tests**: Add tests for new functionality
5. **Build**: Ensure `cd skillgen && go build ./cmd/skillgen && go test ./...` succeed
6. **Generated content**: Do not manually edit files in `skills/` - these are auto-generated

## Testing Guidelines

- Unit tests for extractors and parsers
- Integration tests for end-to-end generation
- Test edge cases (empty content, missing sections, malformed markdown)
- Use table-driven tests where appropriate

## Working with Generated Skills

**IMPORTANT**: Never manually edit files in the `skills/` directory. These are automatically generated from [adaptive-enforcement-lab.com](https://github.com/adaptive-enforcement-lab/adaptive-enforcement-lab-com) documentation.

To modify skill content:
1. Edit the source documentation in the AEL docs repository
2. The skills will auto-regenerate via GitHub Actions when docs are updated

## Questions and Support

- **Issues**: Open an issue for bugs or feature requests
- **Discussions**: Use GitHub Discussions for questions
- **Security**: See [SECURITY.md](SECURITY.md) for vulnerability reporting

## License

By contributing, you agree that your contributions will be licensed under the MIT License.
